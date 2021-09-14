package server

import (
	"errors"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"opengauss_exporter/internal/utils"
)

type Option func(*options)
type options struct {
	HTTP *httpOption
}

type httpOption struct {
	Host string
	Port string
}

type Manager struct {
	options *options
	mux     *chi.Mux
	stopSig chan bool
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

func (s *Manager) setupMiddlewares() {
	//s.mux.Use(middleware.Logger)
}

func (s *Manager) setupRouters() {
	s.mux.Get("/metrics", func(w http.ResponseWriter, r *http.Request) {
		promhttp.Handler().ServeHTTP(w, r)
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

	//TODO support HTTPS server
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
