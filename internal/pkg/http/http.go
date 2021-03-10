package http

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go-demo/internal/pkg/utils/netutil"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Options struct {
	Port    int
	Mode    string
	AppName string
}

type Server struct {
	o          *Options
	host       string
	port       int
	logger     *zap.Logger
	router     *gin.Engine
	httpServer http.Server
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("http", o); err != nil {
		return nil, err
	}

	o.AppName = v.GetString("app.name")

	return o, err
}

func New(o *Options, logger *zap.Logger, router *gin.Engine) (*Server, error) {
	var s = &Server{
		logger: logger.With(zap.String("type", "http.Server")),
		router: router,
		o:      o,
	}
	return s, nil
}

type InitControllers func(r *gin.Engine)

// NewRouter initialize gin router
func NewRouter(o *Options, logger *zap.Logger, init InitControllers) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Recovery()) // panic之后自动恢复
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://*", "https://*"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origins", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	init(r)

	return r
}

func (s *Server) Start() error {
	s.port = s.o.Port
	if s.port == 0 {
		s.port = netutil.GetAvailablePort()
	}

	s.host = netutil.GetLocalIP4()

	if s.host == "" {
		return errors.New("get local ipv4 error")
	}

	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	s.httpServer = http.Server{Addr: addr, Handler: s.router}

	s.logger.Info("http server starting ...", zap.String("addr", addr))
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("start http server err", zap.Error(err))
			return
		}
	}()

	return nil
}

func (s *Server) Reload() error {
	return nil
}

func (s *Server) Stop() error {
	s.logger.Info("stopping http server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) // 平滑关闭,等待5秒钟处理
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "shutdown http server error")
	}

	return nil
}

var ProviderSet = wire.NewSet(New, NewRouter, NewOptions)
