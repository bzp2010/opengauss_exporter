---
title: 刮削器
---

## 刮削器列表

### 核心

#### og_up
数据库实例是否正常运行。

#### og_version
数据库版本号，在 label 中进行标记，例如
```text
{og="2.1.0",pg="9.2.4",server="xxxx"}
```

### PostgreSQL Exporter {#postgresql_exporter}
这个刮削器实现了 PostgreSQL 原版 Exporter 的刮削功能，它从 `pg_` 的一些系统视图中获取运行数据。

#### pg_stat_bgwriter
该视图显示关于后端写进程活动的统计信息。  
维度：数据库节点

- pg_stat_bgwriter_buffers_alloc
- pg_stat_bgwriter_buffers_backend
- pg_stat_bgwriter_buffers_backend_fsync
- pg_stat_bgwriter_buffers_checkpoint
- pg_stat_bgwriter_buffers_clean
- pg_stat_bgwriter_checkpoint_sync_time
- pg_stat_bgwriter_checkpoint_write_time
- pg_stat_bgwriter_checkpoints_req
- pg_stat_bgwriter_checkpoints_timed
- pg_stat_bgwriter_maxwritten_clean
- pg_stat_bgwriter_stats_reset

:::note参考文档
[pg_stat_bgwriter](https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/PG_STAT_BGWRITER.html)
:::

#### pg_stat_database
该视图将包含openGauss中每个数据库的数据库统计信息。  
维度：数据库节点 - 数据库

- pg_stat_database_blk_read_time
- pg_stat_database_blk_write_time
- pg_stat_database_blks_hit
- pg_stat_database_blks_read
- pg_stat_database_conflicts
- pg_stat_database_conflicts_confl_bufferpin
- pg_stat_database_conflicts_confl_deadlock
- pg_stat_database_conflicts_confl_lock
- pg_stat_database_conflicts_confl_snapshot
- pg_stat_database_conflicts_confl_tablespace
- pg_stat_database_deadlocks
- pg_stat_database_numbackends
- pg_stat_database_stats_reset
- pg_stat_database_temp_bytes
- pg_stat_database_temp_files
- pg_stat_database_tup_deleted
- pg_stat_database_tup_fetched
- pg_stat_database_tup_inserted
- pg_stat_database_tup_returned
- pg_stat_database_tup_updated
- pg_stat_database_xact_commit
- pg_stat_database_xact_rollback

:::note参考文档
[pg_stat_database](https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/PG_STAT_DATABASE.html)
:::

#### pg_stat_database_conflicts
该视图显示数据库冲突状态的统计信息。  
维度：数据库节点 - 数据库

- pg_stat_database_conflicts
- pg_stat_database_conflicts_confl_bufferpin
- pg_stat_database_conflicts_confl_deadlock
- pg_stat_database_conflicts_confl_lock
- pg_stat_database_conflicts_confl_snapshot
- pg_stat_database_conflicts_confl_tablespace

:::note参考文档
[pg_stat_database_conflicts](https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/PG_STAT_DATABASE_CONFLICTS.html)
:::

#### pg_locks
该视图存储各打开事务所持有的锁信息。  
维度：数据库节点 - 数据库 - 指标类型

- pg_locks_count
  - accessexclusivelock
  - accesssharelock
  - exclusivelock
  - rowexclusivelock
  - rowsharelock
  - sharelock
  - sharerowexclusivelock
  - shareupdateexclusivelock
  - sireadlock

:::note参考文档
[pg_locks](https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/PG_LOCKS.html)
:::

#### pg_stat_activity
该视图显示和当前用户查询相关的信息。  
维度：数据库节点 - 数据库 - 指标类型

- pg_stat_activity_count
  - active
  - disabled
  - fastpath function call
  - idle
  - idle in transaction
  - idle in transaction (aborted)

:::note参考文档
[pg_stat_activity](https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/PG_STAT_ACTIVITY.html)
:::

#### pg_stat_replication
该视图用于描述日志同步状态信息，如发起端发送日志位置，收端接收日志位置等。  
**备注：当前视图在 openGauss 中不可用**

#### pg_replication_slots
该视图查看复制节点的信息。  
**备注：当前视图在 openGauss 中不可用**

#### pg_stat_archiver
**备注：当前视图在 openGauss 中不可用**

### PG_SETTING
这个刮削器实现了从 openGauss 数据库中获取数据库配置的功能，它从 `pg_settings` 视图中抓取全部的数据库配置。

:::note
指标数量过多，此处不进行列举
:::

### GS_OS_RUN_INFO
这个刮削器从 `GS_OS_RUN_INFO` 系统视图中获取数据库服务器运行情况。  
维度：数据库节点

- gs_os_run_info_avg_busy_time
- gs_os_run_info_avg_idle_time 
- gs_os_run_info_avg_iowait_time
- gs_os_run_info_avg_nice_time
- gs_os_run_info_avg_sys_time
- gs_os_run_info_avg_user_time
- gs_os_run_info_busy_time
- gs_os_run_info_idle_time
- gs_os_run_info_iowait_time
- gs_os_run_info_load
- gs_os_run_info_nice_time
- gs_os_run_info_num_cpu_cores
- gs_os_run_info_num_cpu_sockets
- gs_os_run_info_num_cpus
- gs_os_run_info_physical_memory_bytes
- gs_os_run_info_sys_time
- gs_os_run_info_user_time
- gs_os_run_info_vm_page_in_bytes
- gs_os_run_info_vm_page_out_bytes

:::note参考文档
[GS_OS_RUN_INFO](https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/GS_OS_RUN_INFO.html)
:::

### GS_SQL_COUNT
这个刮削器从 `GS_SQL_COUNT` 系统视图中获取数据库中执行的SQL语句计数情况。  
维度：数据库节点 - 数据库

- gs_sql_count_select_count
- gs_sql_count_insert_count
- gs_sql_count_update_count
- gs_sql_count_delete_count
- gs_sql_count_dcl_count
- gs_sql_count_ddl_count
- gs_sql_count_dml_count
- gs_sql_count_mergeinto_count
- gs_sql_count_avg_delete_elapse
- gs_sql_count_avg_insert_elapse
- gs_sql_count_avg_select_elapse
- gs_sql_count_avg_update_elapse
- gs_sql_count_total_delete_elapse
- gs_sql_count_total_insert_elapse
- gs_sql_count_total_select_elapse
- gs_sql_count_total_update_elapse
- gs_sql_count_max_delete_elapse
- gs_sql_count_max_insert_elapse
- gs_sql_count_max_select_elapse
- gs_sql_count_max_update_elapse
- gs_sql_count_min_delete_elapse
- gs_sql_count_min_insert_elapse
- gs_sql_count_min_select_elapse
- gs_sql_count_min_update_elapse

:::note参考文档
[GS_SQL_COUNT](https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/GS_SQL_COUNT.html)
:::

### GS_INSTANCE_TIME
这个刮削器从 `GS_INSTANCE_TIME` 系统视图中获取当前数据库实例个阶段运行耗时状态。  
维度：数据库节点 - 数据库

- gs_instance_time_cpu_time
- gs_instance_time_data_io_time
- gs_instance_time_db_time
- gs_instance_time_execution_time
- gs_instance_time_net_send_time
- gs_instance_time_parse_time
- gs_instance_time_pl_compilation_time
- gs_instance_time_pl_execution_time
- gs_instance_time_plan_time
- gs_instance_time_rewrite_time

:::note参考文档
[GS_INSTANCE_TIME](https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/GS_INSTANCE_TIME.html)
:::

### GS_TOTAL_MEMORY_DETAIL
这个刮削器从 `GS_TOTAL_MEMORY_DETAIL` 系统视图中获取当前数据库实例内存使用状态。  
维度：数据库节点 - 数据库

- gs_total_memory_detail_backend_used_memory
- gs_total_memory_detail_cstore_used_memory
- gs_total_memory_detail_dynamic_peak_memory
- gs_total_memory_detail_dynamic_peak_shrctx
- gs_total_memory_detail_dynamic_used_memory
- gs_total_memory_detail_dynamic_used_shrctx
- gs_total_memory_detail_gpu_dynamic_peak_memory
- gs_total_memory_detail_gpu_dynamic_used_memory
- gs_total_memory_detail_gpu_max_dynamic_memory
- gs_total_memory_detail_max_backend_memory
- gs_total_memory_detail_max_cstore_memory
- gs_total_memory_detail_max_dynamic_memory
- gs_total_memory_detail_max_process_memory
- gs_total_memory_detail_max_sctpcomm_memory
- gs_total_memory_detail_max_shared_memory
- gs_total_memory_detail_other_used_memory
- gs_total_memory_detail_pooler_conn_memory
- gs_total_memory_detail_pooler_freeconn_memory
- gs_total_memory_detail_process_used_memory
- gs_total_memory_detail_sctpcomm_peak_memory
- gs_total_memory_detail_sctpcomm_used_memory
- gs_total_memory_detail_shared_used_memory
- gs_total_memory_detail_storage_compress_memory
- gs_total_memory_detail_udf_reserved_memory

:::note参考文档
[GS_TOTAL_MEMORY_DETAIL](https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/GS_TOTAL_MEMORY_DETAIL.html)
:::