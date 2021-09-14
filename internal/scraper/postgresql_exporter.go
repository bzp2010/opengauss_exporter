// NOTICE: The source code in this file is comes from PostgreSQL Exporter

package scraper

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"regexp"
	"time"

	pq "gitee.com/opengauss/openGauss-connector-go-pq"
	"github.com/blang/semver"
	"github.com/gogf/gf/util/gconv"
	"github.com/prometheus/client_golang/prometheus"

	"opengauss_exporter/internal/core/scrape"
	"opengauss_exporter/internal/utils"
)

// ColumnUsage should be one of several enum values which describe how a
// queried row is to be converted to a Prometheus metric.
// NOTICE: this part of the code comes from PostgreSQL Exporter
type ColumnUsage int

const (
	// DISCARD ignores a column
	DISCARD ColumnUsage = iota
	// LABEL identifies a column as a label
	LABEL ColumnUsage = iota
	// COUNTER identifies a column as a counter
	COUNTER ColumnUsage = iota
	// GAUGE identifies a column as a gauge
	GAUGE ColumnUsage = iota
	// MAPPEDMETRIC identifies a column as a mapping of text values
	MAPPEDMETRIC ColumnUsage = iota
	// DURATION identifies a column as a text duration (and converted to milliseconds)
	DURATION ColumnUsage = iota
	// HISTOGRAM identifies a column as a histogram
	HISTOGRAM ColumnUsage = iota
)

// MappingOptions is a copy of ColumnMapping used only for parsing
// NOTICE: this part of the code comes from PostgreSQL Exporter
type MappingOptions struct {
	Usage             string             `yaml:"usage"`
	Description       string             `yaml:"description"`
	Mapping           map[string]float64 `yaml:"metric_mapping"` // Optional column mapping for MAPPEDMETRIC
	SupportedVersions semver.Range       `yaml:"pg_version"`     // Semantic version ranges which are supported. Unsupported columns are not queried (internally converted to DISCARD).
}

// Mapping represents a set of MappingOptions
// NOTICE: this part of the code comes from PostgreSQL Exporter
type Mapping map[string]MappingOptions

// Regex used to get the "short-version" from the postgres version field.
// NOTICE: this part of the code comes from PostgreSQL Exporter
var versionRegex = regexp.MustCompile(`^\w+ ((\d+)(\.\d+)?(\.\d+)?)`)

// Parses the version of postgres into the short version string we can use to match behaviors.
// NOTICE: this part of the code comes from PostgreSQL Exporter
func parseVersion(versionString string) (semver.Version, error) {
	submatches := versionRegex.FindStringSubmatch(versionString)
	if len(submatches) > 1 {
		return semver.ParseTolerant(submatches[1])
	}
	return semver.Version{},
		errors.New(fmt.Sprintln("Could not find a postgres version in string:", versionString))
}

// ColumnMapping is the user-friendly representation of a prometheus descriptor map
// NOTICE: this part of the code comes from PostgreSQL Exporter
type ColumnMapping struct {
	usage             ColumnUsage        `yaml:"usage"`
	description       string             `yaml:"description"`
	mapping           map[string]float64 `yaml:"metric_mapping"` // Optional column mapping for MAPPEDMETRIC
	supportedVersions semver.Range       `yaml:"og_version"`     // Semantic version ranges which are supported. Unsupported columns are not queried (internally converted to DISCARD).
}

// intermediateMetricMap holds the partially loaded metric map parsing.
// This is mainly so we can parse cacheSeconds around.
// NOTICE: this part of the code comes from PostgreSQL Exporter
type intermediateMetricMap struct {
	columnMappings map[string]ColumnMapping
	master         bool
	cacheSeconds   uint64
}

// MetricMapNamespace groups metric maps under a shared set of labels.
// NOTICE: this part of the code comes from PostgreSQL Exporter
type MetricMapNamespace struct {
	labels         []string             // Label names for this namespace
	columnMappings map[string]MetricMap // Column mappings in this namespace
	master         bool                 // Call query only for master database
	cacheSeconds   uint64               // Number of seconds this metric namespace can be cached. 0 disables.
}

// MetricMap stores the prometheus metric description which a given column will
// be mapped to by the collector
// NOTICE: this part of the code comes from PostgreSQL Exporter
type MetricMap struct {
	discard    bool                              // Should metric be discarded during mapping?
	histogram  bool                              // Should metric be treated as a histogram?
	vtype      prometheus.ValueType              // Prometheus valuetype
	desc       *prometheus.Desc                  // Prometheus descriptor
	conversion func(interface{}) (float64, bool) // Conversion function to turn PG result into float64
}

// OverrideQuery 's are run in-place of simple namespace look ups, and provide
// advanced functionality. But they have a tendency to postgres version specific.
// There aren't too many versions, so we simply store customized versions using
// the semver matching we do for columns.
// NOTICE: this part of the code comes from PostgreSQL Exporter
type OverrideQuery struct {
	versionRange semver.Range
	query        string
}

// Metrics map from PostgreSQL Exporter
// NOTICE: this part of the code comes from PostgreSQL Exporter
var builtinMetricMaps = map[string]intermediateMetricMap{
	"pg_stat_bgwriter": {
		map[string]ColumnMapping{
			"checkpoints_timed":     {COUNTER, "Number of scheduled checkpoints that have been performed", nil, nil},
			"checkpoints_req":       {COUNTER, "Number of requested checkpoints that have been performed", nil, nil},
			"checkpoint_write_time": {COUNTER, "Total amount of time that has been spent in the portion of checkpoint processing where files are written to disk, in milliseconds", nil, nil},
			"checkpoint_sync_time":  {COUNTER, "Total amount of time that has been spent in the portion of checkpoint processing where files are synchronized to disk, in milliseconds", nil, nil},
			"buffers_checkpoint":    {COUNTER, "Number of buffers written during checkpoints", nil, nil},
			"buffers_clean":         {COUNTER, "Number of buffers written by the background writer", nil, nil},
			"maxwritten_clean":      {COUNTER, "Number of times the background writer stopped a cleaning scan because it had written too many buffers", nil, nil},
			"buffers_backend":       {COUNTER, "Number of buffers written directly by a backend", nil, nil},
			"buffers_backend_fsync": {COUNTER, "Number of times a backend had to execute its own fsync call (normally the background writer handles those even when the backend does its own write)", nil, nil},
			"buffers_alloc":         {COUNTER, "Number of buffers allocated", nil, nil},
			"stats_reset":           {COUNTER, "Time at which these statistics were last reset", nil, nil},
		},
		true,
		0,
	},
	"pg_stat_database": {
		map[string]ColumnMapping{
			"datid":          {LABEL, "OID of a database", nil, nil},
			"datname":        {LABEL, "Name of this database", nil, nil},
			"numbackends":    {GAUGE, "Number of backends currently connected to this database. This is the only column in this view that returns a value reflecting current state; all other columns return the accumulated values since the last reset.", nil, nil},
			"xact_commit":    {COUNTER, "Number of transactions in this database that have been committed", nil, nil},
			"xact_rollback":  {COUNTER, "Number of transactions in this database that have been rolled back", nil, nil},
			"blks_read":      {COUNTER, "Number of disk blocks read in this database", nil, nil},
			"blks_hit":       {COUNTER, "Number of times disk blocks were found already in the buffer cache, so that a read was not necessary (this only includes hits in the PostgreSQL buffer cache, not the operating system's file system cache)", nil, nil},
			"tup_returned":   {COUNTER, "Number of rows returned by queries in this database", nil, nil},
			"tup_fetched":    {COUNTER, "Number of rows fetched by queries in this database", nil, nil},
			"tup_inserted":   {COUNTER, "Number of rows inserted by queries in this database", nil, nil},
			"tup_updated":    {COUNTER, "Number of rows updated by queries in this database", nil, nil},
			"tup_deleted":    {COUNTER, "Number of rows deleted by queries in this database", nil, nil},
			"conflicts":      {COUNTER, "Number of queries canceled due to conflicts with recovery in this database. (Conflicts occur only on standby servers; see pg_stat_database_conflicts for details.)", nil, nil},
			"temp_files":     {COUNTER, "Number of temporary files created by queries in this database. All temporary files are counted, regardless of why the temporary file was created (e.g., sorting or hashing), and regardless of the log_temp_files setting.", nil, nil},
			"temp_bytes":     {COUNTER, "Total amount of data written to temporary files by queries in this database. All temporary files are counted, regardless of why the temporary file was created, and regardless of the log_temp_files setting.", nil, nil},
			"deadlocks":      {COUNTER, "Number of deadlocks detected in this database", nil, nil},
			"blk_read_time":  {COUNTER, "Time spent reading data file blocks by backends in this database, in milliseconds", nil, nil},
			"blk_write_time": {COUNTER, "Time spent writing data file blocks by backends in this database, in milliseconds", nil, nil},
			"stats_reset":    {COUNTER, "Time at which these statistics were last reset", nil, nil},
		},
		true,
		0,
	},
	"pg_stat_database_conflicts": {
		map[string]ColumnMapping{
			"datid":            {LABEL, "OID of a database", nil, nil},
			"datname":          {LABEL, "Name of this database", nil, nil},
			"confl_tablespace": {COUNTER, "Number of queries in this database that have been canceled due to dropped tablespaces", nil, nil},
			"confl_lock":       {COUNTER, "Number of queries in this database that have been canceled due to lock timeouts", nil, nil},
			"confl_snapshot":   {COUNTER, "Number of queries in this database that have been canceled due to old snapshots", nil, nil},
			"confl_bufferpin":  {COUNTER, "Number of queries in this database that have been canceled due to pinned buffers", nil, nil},
			"confl_deadlock":   {COUNTER, "Number of queries in this database that have been canceled due to deadlocks", nil, nil},
		},
		true,
		0,
	},
	"pg_locks": {
		map[string]ColumnMapping{
			"datname": {LABEL, "Name of this database", nil, nil},
			"mode":    {LABEL, "Type of Lock", nil, nil},
			"count":   {GAUGE, "Number of locks", nil, nil},
		},
		true,
		0,
	},
	"pg_stat_replication": {
		map[string]ColumnMapping{
			"pid":              {DISCARD, "Process ID of a WAL sender process", nil, semver.MustParseRange(">=9.2.0")},
			"usesysid":         {DISCARD, "OID of the user logged into this WAL sender process", nil, nil},
			"usename":          {DISCARD, "Name of the user logged into this WAL sender process", nil, nil},
			"application_name": {LABEL, "Name of the application that is connected to this WAL sender", nil, nil},
			"client_addr":      {LABEL, "IP address of the client connected to this WAL sender. If this field is null, it indicates that the client is connected via a Unix socket on the server machine.", nil, nil},
			"client_hostname":  {DISCARD, "Host name of the connected client, as reported by a reverse DNS lookup of client_addr. This field will only be non-null for IP connections, and only when log_hostname is enabled.", nil, nil},
			"client_port":      {DISCARD, "TCP port number that the client is using for communication with this WAL sender, or -1 if a Unix socket is used", nil, nil},
			"backend_start": {DISCARD, "with time zone	Time when this process was started, i.e., when the client connected to this WAL sender", nil, nil},
			"backend_xmin":             {DISCARD, "The current backend's xmin horizon.", nil, nil},
			"state":                    {LABEL, "Current WAL sender state", nil, nil},
			"sent_location":            {DISCARD, "Last transaction log position sent on this connection", nil, semver.MustParseRange("<10.0.0")},
			"write_location":           {DISCARD, "Last transaction log position written to disk by this standby server", nil, semver.MustParseRange("<10.0.0")},
			"flush_location":           {DISCARD, "Last transaction log position flushed to disk by this standby server", nil, semver.MustParseRange("<10.0.0")},
			"sync_priority":            {DISCARD, "Priority of this standby server for being chosen as the synchronous standby", nil, nil},
			"sync_state":               {DISCARD, "Synchronous state of this standby server", nil, nil},
			"slot_name":                {LABEL, "A unique, cluster-wide identifier for the replication slot", nil, semver.MustParseRange(">=9.2.0")},
			"plugin":                   {DISCARD, "The base name of the shared object containing the output plugin this logical slot is using, or null for physical slots", nil, nil},
			"slot_type":                {DISCARD, "The slot type - physical or logical", nil, nil},
			"datoid":                   {DISCARD, "The OID of the database this slot is associated with, or null. Only logical slots have an associated database", nil, nil},
			"database":                 {DISCARD, "The name of the database this slot is associated with, or null. Only logical slots have an associated database", nil, nil},
			"active":                   {DISCARD, "True if this slot is currently actively being used", nil, nil},
			"active_pid":               {DISCARD, "Process ID of a WAL sender process", nil, nil},
			"xmin":                     {DISCARD, "The oldest transaction that this slot needs the database to retain. VACUUM cannot remove tuples deleted by any later transaction", nil, nil},
			"catalog_xmin":             {DISCARD, "The oldest transaction affecting the system catalogs that this slot needs the database to retain. VACUUM cannot remove catalog tuples deleted by any later transaction", nil, nil},
			"restart_lsn":              {DISCARD, "The address (LSN) of oldest WAL which still might be required by the consumer of this slot and thus won't be automatically removed during checkpoints", nil, nil},
			"pg_current_xlog_location": {DISCARD, "pg_current_xlog_location", nil, nil},
			"pg_xlog_location_diff":    {GAUGE, "Lag in bytes between master and slave", nil, semver.MustParseRange(">=9.2.0 <10.0.0")},
			"confirmed_flush_lsn":      {DISCARD, "LSN position a consumer of a slot has confirmed flushing the data received", nil, nil},
		},
		true,
		0,
	},
	"pg_replication_slots": {
		map[string]ColumnMapping{
			"slot_name":       {LABEL, "Name of the replication slot", nil, nil},
			"database":        {LABEL, "Name of the database", nil, nil},
			"active":          {GAUGE, "Flag indicating if the slot is active", nil, nil},
			"pg_wal_lsn_diff": {GAUGE, "Replication lag in bytes", nil, nil},
		},
		true,
		0,
	},
	"pg_stat_activity": {
		map[string]ColumnMapping{
			"datname":         {LABEL, "Name of this database", nil, nil},
			"state":           {LABEL, "connection state", nil, semver.MustParseRange(">=9.2.0")},
			"count":           {GAUGE, "number of connections in this state", nil, nil},
			"max_tx_duration": {GAUGE, "max duration in seconds any active transaction has been running", nil, nil},
		},
		true,
		0,
	},
}

// Overriding queries for namespaces above.
// NOTICE: this part of the code comes from PostgreSQL Exporter
var builtinQueryOverrides = map[string][]OverrideQuery{
	"pg_locks": {
		{
			semver.MustParseRange(">0.0.0"),
			`SELECT pg_database.datname,tmp.mode,COALESCE(count,0) as count
			FROM
				(
				  VALUES ('accesssharelock'),
				         ('rowsharelock'),
				         ('rowexclusivelock'),
				         ('shareupdateexclusivelock'),
				         ('sharelock'),
				         ('sharerowexclusivelock'),
				         ('exclusivelock'),
				         ('accessexclusivelock'),
					 ('sireadlock')
				) AS tmp(mode) CROSS JOIN pg_database
			LEFT JOIN
			  (SELECT database, lower(mode) AS mode,count(*) AS count
			  FROM pg_locks WHERE database IS NOT NULL
			  GROUP BY database, lower(mode)
			) AS tmp2
			ON tmp.mode=tmp2.mode and pg_database.oid = tmp2.database ORDER BY 1`,
		},
	},
	"pg_replication_slots": {
		{
			semver.MustParseRange(">=9.4.0 <10.0.0"),
			`
			SELECT slot_name, database, active, pg_xlog_location_diff(pg_current_xlog_location(), restart_lsn)
			FROM pg_replication_slots
			`,
		},
	},
	"pg_stat_archiver": {
		{
			semver.MustParseRange(">=0.0.0"),
			`
			SELECT *,
				extract(epoch from now() - last_archived_time) AS last_archive_age
			FROM pg_stat_archiver
			`,
		},
	},
	"pg_stat_activity": {
		// This query only works
		{
			semver.MustParseRange(">=9.2.0"),
			`
			SELECT
				pg_database.datname,
				tmp.state,
				COALESCE(count,0) as count,
				COALESCE(max_tx_duration,0) as max_tx_duration
			FROM
				(
				  VALUES ('active'),
				  		 ('idle'),
				  		 ('idle in transaction'),
				  		 ('idle in transaction (aborted)'),
				  		 ('fastpath function call'),
				  		 ('disabled')
				) AS tmp(state) CROSS JOIN pg_database
			LEFT JOIN
			(
				SELECT
					datname,
					state,
					count(*) AS count,
					MAX(EXTRACT(EPOCH FROM now() - xact_start))::float AS max_tx_duration
				FROM pg_stat_activity GROUP BY datname,state) AS tmp2
				ON tmp.state = tmp2.state AND pg_database.datname = tmp2.datname
			`,
		},
	},
}

// Turn the MetricMap column mapping into a prometheus descriptor mapping.
// NOTICE: this part of the code comes from PostgreSQL Exporter
func makeDescMap(task *scrape.Task, metricMaps map[string]intermediateMetricMap) map[string]MetricMapNamespace {
	var metricMap = make(map[string]MetricMapNamespace)

	for namespace, intermediateMappings := range metricMaps {
		thisMap := make(map[string]MetricMap)

		// Get the variable labels
		var variableLabels []string
		for columnName, columnMapping := range intermediateMappings.columnMappings {
			if columnMapping.usage == LABEL {
				variableLabels = append(variableLabels, columnName)
			}
		}

		for columnName, columnMapping := range intermediateMappings.columnMappings {
			// Check column version compatibility for the current map
			// Force to discard if not compatible.
			if columnMapping.supportedVersions != nil {
				if !columnMapping.supportedVersions(task.PostgreSQLVersion()) {
					// It's very useful to be able to see what columns are being rejected.
					utils.GetLogger().Warnw("Column is being forced to discard due to version incompatibility", "column", columnName)
					thisMap[columnName] = MetricMap{
						discard: true,
						conversion: func(_ interface{}) (float64, bool) {
							return math.NaN(), true
						},
					}
					continue
				}
			}

			// Determine how to convert the column based on its usage.
			// nolint: dupl
			switch columnMapping.usage {
			case DISCARD, LABEL:
				thisMap[columnName] = MetricMap{
					discard: true,
					conversion: func(_ interface{}) (float64, bool) {
						return math.NaN(), true
					},
				}
			case COUNTER:
				thisMap[columnName] = MetricMap{
					vtype: prometheus.CounterValue,
					desc:  prometheus.NewDesc(fmt.Sprintf("%s_%s", namespace, columnName), columnMapping.description, variableLabels, task.ConstLabels()),
					conversion: func(in interface{}) (float64, bool) {
						return gconv.Float64(in), true
					},
				}
			case GAUGE:
				thisMap[columnName] = MetricMap{
					vtype: prometheus.GaugeValue,
					desc:  prometheus.NewDesc(fmt.Sprintf("%s_%s", namespace, columnName), columnMapping.description, variableLabels, task.ConstLabels()),
					conversion: func(in interface{}) (float64, bool) {
						return gconv.Float64(in), true
					},
				}
			case HISTOGRAM:
				thisMap[columnName] = MetricMap{
					histogram: true,
					vtype:     prometheus.UntypedValue,
					desc:      prometheus.NewDesc(fmt.Sprintf("%s_%s", namespace, columnName), columnMapping.description, variableLabels, task.ConstLabels()),
					conversion: func(in interface{}) (float64, bool) {
						return gconv.Float64(in), true
					},
				}
				thisMap[columnName+"_bucket"] = MetricMap{
					histogram: true,
					discard:   true,
				}
				thisMap[columnName+"_sum"] = MetricMap{
					histogram: true,
					discard:   true,
				}
				thisMap[columnName+"_count"] = MetricMap{
					histogram: true,
					discard:   true,
				}
			case MAPPEDMETRIC:
				thisMap[columnName] = MetricMap{
					vtype: prometheus.GaugeValue,
					desc:  prometheus.NewDesc(fmt.Sprintf("%s_%s", namespace, columnName), columnMapping.description, variableLabels, task.ConstLabels()),
					conversion: func(in interface{}) (float64, bool) {
						text, ok := in.(string)
						if !ok {
							return math.NaN(), false
						}

						val, ok := columnMapping.mapping[text]
						if !ok {
							return math.NaN(), false
						}
						return val, true
					},
				}
			case DURATION:
				thisMap[columnName] = MetricMap{
					vtype: prometheus.GaugeValue,
					desc:  prometheus.NewDesc(fmt.Sprintf("%s_%s_milliseconds", namespace, columnName), columnMapping.description, variableLabels, task.ConstLabels()),
					conversion: func(in interface{}) (float64, bool) {
						var durationString string
						switch t := in.(type) {
						case []byte:
							durationString = string(t)
						case string:
							durationString = t
						default:
							utils.GetLogger().Warn("Duration conversion metric was not a string")
							return math.NaN(), false
						}

						if durationString == "-1" {
							return math.NaN(), false
						}

						d, err := time.ParseDuration(durationString)
						if err != nil {
							utils.GetLogger().Warnw("Failed converting result to metric", "column", columnName, "in", in, "err", err)
							return math.NaN(), false
						}
						return float64(d / time.Millisecond), true
					},
				}
			}
		}

		metricMap[namespace] = MetricMapNamespace{variableLabels, thisMap, intermediateMappings.master, intermediateMappings.cacheSeconds}
	}

	return metricMap
}

// Convert the query override file to the version-specific query override file for the exporter.
// NOTICE: this part of the code comes from PostgreSQL Exporter
func makeQueryOverrideMap(task *scrape.Task, queryOverrides map[string][]OverrideQuery) map[string]string {
	resultMap := make(map[string]string)
	for name, overrideDef := range queryOverrides {
		// Find a matching semver. We make it an error to have overlapping
		// ranges at test-time, so only 1 should ever match.
		matched := false
		for _, queryDef := range overrideDef {
			if queryDef.versionRange(task.PostgreSQLVersion()) {
				resultMap[name] = queryDef.query
				matched = true
				break
			}
		}
		if !matched {
			utils.GetLogger().Warnw("No matched query override, disabling metric space",
				"server", task.DataSource().Fingerprint(),
				"name", name,
			)
			resultMap[name] = ""
		}
	}

	return resultMap
}

type BuiltinSQLScraper struct{}

func NewBuiltinSQLScraper() *BuiltinSQLScraper {
	return &BuiltinSQLScraper{}
}

func (b BuiltinSQLScraper) Scrape(task *scrape.Task) ([]prometheus.Metric, []error, error) {
	allNonfatalErrors := make([]error, 0)
	allMetrics := make([]prometheus.Metric, 0)

	queryOverrides := makeQueryOverrideMap(task, builtinQueryOverrides)

	for namespace, mapping := range makeDescMap(task, builtinMetricMaps) {
		utils.GetLogger().Infow("Querying namespace",
			"server", task.DataSource().Fingerprint(),
			"namespace", namespace,
		)

		if mapping.master && !task.DataSource().Master {
			utils.GetLogger().Infow("Query skipped",
				"server", task.DataSource().Fingerprint(),
				"namespace", namespace,
			)
			continue
		}

		metrics, nonfatalErrors, err := queryNamespaceMapping(task, namespace, mapping, queryOverrides)

		// Serious error - a namespace disappeared
		if err != nil {
			utils.GetLogger().Error("err", err)
			return nil, nil, err
		}

		// Non-serious errors - likely version or parsing problems.
		if len(nonfatalErrors) > 0 {
			for _, err := range nonfatalErrors {
				allNonfatalErrors = append(allNonfatalErrors, err)
				utils.GetLogger().Error("err", err)
			}
		}

		allMetrics = append(allMetrics, metrics...)
	}

	return allMetrics, allNonfatalErrors, nil
}

// Query within a namespace mapping and emit metrics. Returns fatal errors if
// the scrape fails, and a slice of errors if they were non-fatal.
func queryNamespaceMapping(task *scrape.Task, namespace string, mapping MetricMapNamespace, queryOverrides map[string]string) ([]prometheus.Metric, []error, error) {
	// Check for a query override for this namespace
	query, found := queryOverrides[namespace]

	// Was this query disabled (i.e. nothing sensible can be queried on cu
	// version of PostgreSQL?
	if query == "" && found {
		// Return success (no pertinent data)
		return []prometheus.Metric{}, []error{}, nil
	}

	// Don't fail on a bad scrape of one metric
	var rows *sql.Rows
	var err error

	if !found {
		rows, err = task.DB().Query(fmt.Sprintf("SELECT * FROM %s;", namespace)) // nolint: gas
	} else {
		rows, err = task.DB().Query(query)
	}
	if err != nil {
		return []prometheus.Metric{}, []error{}, fmt.Errorf("Error running query on datasource %s: %s %v", task.DataSource().Fingerprint(), namespace, err)
	}
	defer rows.Close()

	var columnNames []string
	columnNames, err = rows.Columns()
	if err != nil {
		return []prometheus.Metric{}, []error{}, errors.New(fmt.Sprintln("error retrieving column list for: ", namespace, err))
	}

	// Make a lookup map for the column indices
	var columnIdx = make(map[string]int, len(columnNames))
	for i, n := range columnNames {
		columnIdx[n] = i
	}

	var columnData = make([]interface{}, len(columnNames))
	var scanArgs = make([]interface{}, len(columnNames))
	for i := range columnData {
		scanArgs[i] = &columnData[i]
	}

	nonfatalErrors := make([]error, 0)

	metrics := make([]prometheus.Metric, 0)

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return []prometheus.Metric{}, []error{}, errors.New(fmt.Sprintln("Error retrieving rows:", namespace, err))
		}

		// Get the label values for this row.
		labels := make([]string, len(mapping.labels))
		for idx, label := range mapping.labels {
			labels[idx] = gconv.String(columnData[columnIdx[label]])
		}

		// Loop over column names, and match to scan data. Unknown columns
		// will be filled with an untyped metric number *if* they can be
		// converted to float64s. NULLs are allowed and treated as NaN.
		for idx, columnName := range columnNames {
			var metric prometheus.Metric
			if metricMapping, ok := mapping.columnMappings[columnName]; ok {
				// Is this a metricy metric?
				if metricMapping.discard {
					continue
				}

				if metricMapping.histogram {
					var keys []float64
					err = pq.Array(&keys).Scan(columnData[idx])
					if err != nil {
						return []prometheus.Metric{}, []error{}, errors.New(fmt.Sprintln("Error retrieving", columnName, "buckets:", namespace, err))
					}

					var values []int64
					valuesIdx, ok := columnIdx[columnName+"_bucket"]
					if !ok {
						nonfatalErrors = append(nonfatalErrors, errors.New(fmt.Sprintln("Missing column: ", namespace, columnName+"_bucket")))
						continue
					}
					err = pq.Array(&values).Scan(columnData[valuesIdx])
					if err != nil {
						return []prometheus.Metric{}, []error{}, errors.New(fmt.Sprintln("Error retrieving", columnName, "bucket values:", namespace, err))
					}

					buckets := make(map[float64]uint64, len(keys))
					for i, key := range keys {
						if i >= len(values) {
							break
						}
						buckets[key] = uint64(values[i])
					}

					idx, ok = columnIdx[columnName+"_sum"]
					if !ok {
						nonfatalErrors = append(nonfatalErrors, errors.New(fmt.Sprintln("Missing column: ", namespace, columnName+"_sum")))
						continue
					}
					sum, ok := gconv.Float64(columnData[idx]), true
					if !ok {
						nonfatalErrors = append(nonfatalErrors, errors.New(fmt.Sprintln("Unexpected error parsing column: ", namespace, columnName+"_sum", columnData[idx])))
						continue
					}

					idx, ok = columnIdx[columnName+"_count"]
					if !ok {
						nonfatalErrors = append(nonfatalErrors, errors.New(fmt.Sprintln("Missing column: ", namespace, columnName+"_count")))
						continue
					}
					count, ok := gconv.Uint64(columnData[idx]), true
					if !ok {
						nonfatalErrors = append(nonfatalErrors, errors.New(fmt.Sprintln("Unexpected error parsing column: ", namespace, columnName+"_count", columnData[idx])))
						continue
					}

					metric = prometheus.MustNewConstHistogram(
						metricMapping.desc,
						count, sum, buckets,
						labels...,
					)
				} else {
					value := gconv.Float64(columnData[idx])

					// Generate the metric
					metric = prometheus.MustNewConstMetric(metricMapping.desc, metricMapping.vtype, value, labels...)
				}
			} else {
				// Unknown metric. Report as untyped if scan to float64 works, else note an error too.
				metricLabel := fmt.Sprintf("%s_%s", namespace, columnName)
				desc := prometheus.NewDesc(metricLabel, fmt.Sprintf("Unknown metric from %s", namespace), mapping.labels, task.ConstLabels())

				// Its not an error to fail here, since the values are
				// unexpected anyway.
				value := gconv.Float64(columnData[idx])
				metric = prometheus.MustNewConstMetric(desc, prometheus.UntypedValue, value, labels...)
			}
			metrics = append(metrics, metric)
		}
	}
	return metrics, nonfatalErrors, nil
}
