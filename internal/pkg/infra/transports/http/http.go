package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/base/netutil"
	"github.com/blackhorseya/todo-app/internal/pkg/infra/transports/http/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// InitHandlers define register handler
type InitHandlers func(r *gin.Engine)

// Options declare http's configuration
type Options struct {
	Host string
	Port int
	Mode string
}

// Server declare http server
type Server struct {
	o          *Options
	app        string
	host       string
	port       int
	logger     *zap.Logger
	router     *gin.Engine
	httpServer http.Server
}

// NewOptions serve caller to create Options
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("http", o); err != nil {
		return nil, err
	}

	return o, nil
}

// NewRouter serve caller to create *gin.Engine
func NewRouter(o *Options, logger *zap.Logger, init InitHandlers) *gin.Engine {
	gin.SetMode(o.Mode)

	r := gin.New()

	defaultCORS := cors.DefaultConfig()
	defaultCORS.AllowAllOrigins = true
	defaultCORS.AddExposeHeaders("X-Total-Count")

	r.Use(cors.New(defaultCORS))
	r.Use(middlewares.ContextMiddleware())
	r.Use(gin.Recovery())
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(middlewares.ErrorMiddleware())
	// r.Use(rate.IPRateLimitMiddleware(1, 60))

	r.Use(static.Serve("/", static.LocalFile("./web/build", true)))
	// r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	init(r)

	return r
}

// New serve caller to create Server
func New(o *Options, logger *zap.Logger, router *gin.Engine) (*Server, error) {
	var s = &Server{
		o:          o,
		logger:     logger.With(zap.String("type", "http.Server")),
		router:     router,
		httpServer: http.Server{},
	}

	return s, nil
}

// Application serve caller set server's name
func (s *Server) Application(name string) {
	s.app = name
}

// Start serve caller to start server
func (s *Server) Start() error {
	s.port = s.o.Port
	if s.port == 0 {
		s.port = netutil.GetAvailablePort()
	}

	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	s.httpServer = http.Server{Addr: addr, Handler: s.router}

	s.logger.Info("http server starting...", zap.String("addr", addr))
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("start http server err", zap.Error(err))
			return
		}
	}()

	return nil
}

// Stop serve caller to stop server
func (s *Server) Stop() error {
	s.logger.Info("stopping http server")

	timeout, cancel := contextx.WithTimeout(contextx.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(timeout); err != nil {
		return errors.Wrap(err, "shutdown http server error")
	}

	return nil
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(New, NewRouter, NewOptions)
