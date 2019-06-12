package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/CDcoding2333/pet/internal/app/department"
	"github.com/CDcoding2333/pet/internal/pkg/resp"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// Server ...
type Server interface {
	Start() error
}

type server struct {
	departService department.Service
	port          int
	ch            chan os.Signal
}

// Config ...
type Config struct {
	DB   *gorm.DB
	Port int
	Ch   chan os.Signal
}

// NewAPIServer ...
func NewAPIServer(cnf *Config) (Server, error) {

	resp := resp.NewResponse()

	s := &server{
		port: cnf.Port,
		ch:   cnf.Ch,
	}
	var err error
	s.departService, err = department.NewService(cnf.DB, resp)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return s, nil
}

// Start run http server
func (s server) Start() error {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())

	router.GET("/favicon.ico", func(c *gin.Context) {})
	router.GET("/ping", func(c *gin.Context) { c.Status(http.StatusOK) })

	v2 := router.Group("/poplar/v2")
	{
		v2.POST("/roles", s.departService.HandlerNewDepartment)
	}

	go router.Run(fmt.Sprintf(":%d", 9999))

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: router,
	}

	log.Infoln("Server start:", s.port)

	go func() {
		sig := <-s.ch
		log.Infof(`Received signal "%v", exiting...`, sig)
		if err := httpServer.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
		log.Infoln("Server exited")
	}()

	return httpServer.ListenAndServe()
}
