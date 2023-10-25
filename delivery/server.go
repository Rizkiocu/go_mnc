package delivery

import (
	"fmt"
	"test_mnc/config"
	"test_mnc/delivery/controller"
	"test_mnc/delivery/middleware"
	"test_mnc/manager"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	ucManager manager.UseCaseManager
	engine    *gin.Engine
	host      string
	log       *logrus.Logger
}

func (s *Server) Run() {
	s.initMiddlewares()
	s.initControllers()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initMiddlewares() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))
}

func (s *Server) initControllers() {
	rg := s.engine.Group("/api/v1")
	controller.NewAuthController(s.ucManager.UserUseCase(), s.ucManager.AuthUseCase(), rg).Route()
	controller.NewPaymentController(s.ucManager.PaymentUseCase(), rg).Route()
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		fmt.Println(err)
	}
	rm := manager.NewRepoManager(infraManager)
	ucm := manager.NewUseCaseManager(rm)

	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	log := logrus.New()

	engine := gin.Default()
	return &Server{
		ucManager: ucm,
		engine:    engine,
		host:      host,
		log:       log,
	}
}
