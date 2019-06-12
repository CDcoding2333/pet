package main

import (
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/CDcoding2333/pet/api"
	"github.com/CDcoding2333/pet/configs"
	"github.com/CDcoding2333/pet/internal/pkg/database"
	"github.com/CDcoding2333/pet/pkg/flagtools"
	"github.com/CDcoding2333/pet/pkg/logrus"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

var (
	dbCnf          configs.DbConfig
	port           int
	logLevel       int
	logFileEnabled bool
)

func init() {
	//api-server
	pflag.IntVar(&port, "port", 8088, "server port")

	//log
	pflag.IntVar(&logLevel, "log_level", 5, "log lever 0-5")
	pflag.BoolVar(&logFileEnabled, "log_file_enabled", false, "enable log to file")

	//db
	pflag.StringVar(&dbCnf.Driver, "db_driver", "mysql", "db driver")
	pflag.StringVar(&dbCnf.Source, "db_con", "root:root@tcp(127.0.0.1:3306)/pet?charset=utf8&parseTime=True&loc=Local", "driver connection string")
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())

	flagtools.InitFlags()

	// init logs
	if err := initLogs(); err != nil {
		return
	}

	db, err := database.NewDB(&dbCnf)
	if err != nil {
		log.Fatalln(err)
		return
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(
		ch,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)

	s, err := api.NewAPIServer(&api.Config{DB: db, Port: port, Ch: ch})
	if err != nil {
		log.Fatalln(err)
		return
	}

	s.Start()
}

func initLogs() error {
	conf := &logrus.Config{
		Format: &log.JSONFormatter{},
		Level:  log.Level(logLevel),
		Tags:   "app",
	}
	return logrus.InitLog(conf)
}
