package server

import (
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"opengauss_exporter/internal/core/cache"
	"opengauss_exporter/internal/utils"
)

type Option func(*options)
type options struct {
	HTTP        *httpOption
	HTTPS       *httpsOption
	Middlewares map[string]interface{}
}

type httpOption struct {
	Host string
	Port string
}

type httpsOption struct {
	Host    string
	Port    string
	SSLCert string
	SSLKey  string
}

type Manager struct {
	options *options
	mux     *chi.Mux
	stopSig chan bool
}

var middlewareMap = map[string]func(next http.Handler) http.Handler{
	"logger":     chiMiddleware.Logger,
	"basic_auth": func(next http.Handler) http.Handler { return nil }, // only placeholder
}

// NewManager create a HTTP server for API endpoints
func NewManager(opts ...Option) *Manager {
	s := &Manager{
		options: &options{},
		stopSig: make(chan bool, 1),
	}

	// setup options
	for _, opt := range opts {
		opt(s.options)
	}

	// initialize http handler
	r := chi.NewRouter()

	// setup server manager
	s.mux = r
	s.setupMiddlewares()
	s.setupRouters()

	return s
}

func (s *Manager) With(o Option) {
	o(s.options)
}

func (s *Manager) setupMiddlewares() {
	for middlewareId, middlewareConfig := range s.options.Middlewares {
		if middleware, ok := middlewareMap[middlewareId]; ok {
			if middlewareId == "basic_auth" {
				config := middlewareConfig.(map[string]interface{})
				creds := map[string]string{}
				for key, val := range config {
					creds[key] = val.(string)
				}
				s.mux.Use(chiMiddleware.BasicAuth("opengauss_exporter", creds))
				return
			}
			s.mux.Use(middleware)
		} else {
			utils.GetLogger().Errorw("HTTP server middleware not registered",
				"middleware", middlewareId,
			)
		}
	}

}

func (s *Manager) setupRouters() {
	s.mux.Get("/metrics", func(w http.ResponseWriter, r *http.Request) {
		promhttp.Handler().ServeHTTP(w, r)
	})

	s.mux.Get("/refresh", func(writer http.ResponseWriter, request *http.Request) {
		err := cache.Metrics.Clear()
		if err != nil {
			_, _ = writer.Write([]byte(fmt.Sprintf("refresh failed: %v", err)))
			return
		}
		_, _ = writer.Write([]byte("refresh success"))
	})
}

func (s *Manager) Start(errSig chan error) {
	if s.options.HTTP == nil {
		errSig <- errors.New("HTTP server disabled")
		return
	}

	// start HTTP server
	if s.options.HTTP != nil {
		addr := net.JoinHostPort(s.options.HTTP.Host, s.options.HTTP.Port)
		utils.GetLogger().Infow("HTTP server listening on address",
			"address", addr,
		)
		go func() {
			err := http.ListenAndServe(addr, s.mux)
			if err != nil {
				utils.GetLogger().Errorf("HTTP server listening failed: %v", err)
				errSig <- err
			}
		}()
	}

	//  start HTTPS server
	if s.options.HTTPS != nil {
		addr := net.JoinHostPort(s.options.HTTPS.Host, s.options.HTTPS.Port)
		utils.GetLogger().Infow("HTTPS server listening on address",
			"address", addr,
		)

		go func() {
			err := http.ListenAndServeTLS(addr, s.options.HTTPS.SSLCert, s.options.HTTPS.SSLKey, s.mux)
			if err != nil {
				utils.GetLogger().Errorf("HTTPS server listening failed: %v", err)
				errSig <- err
			}
		}()
	}
}

func (s *Manager) Stop() {
	s.stopSig <- true
}

func WithHTTPServer(host string, port string) Option {
	return func(options *options) {
		if host == "" {
			host = "127.0.0.1"
		}
		options.HTTP = &httpOption{
			Host: host,
			Port: port,
		}
	}
}

func WithHTTPSServer(host string, port string, certFile string, keyFile string) Option {
	return func(options *options) {
		if host == "" {
			host = "127.0.0.1"
		}
		options.HTTPS = &httpsOption{
			Host:    host,
			Port:    port,
			SSLCert: certFile,
			SSLKey:  keyFile,
		}
	}
}

func WithMiddlewares(middlewares map[string]interface{}) func(options *options) {
	return func(options *options) {
		options.Middlewares = middlewares
	}
}
