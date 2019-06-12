package logrus

import (
	"net"
	"os"

	"github.com/CDcoding2333/pet/pkg/logrus/defaulthook"
	"github.com/CDcoding2333/pet/pkg/logrus/filehook"
	"github.com/CDcoding2333/pet/pkg/logrus/filename"
	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	log "github.com/sirupsen/logrus"
)

// Config log config
type Config struct {
	Format         log.Formatter
	Level          log.Level
	Tags           string
	EnableFile     bool
	FileConfig     *filehook.Config
	EnableLogstash bool
	LogstashConf   *LogstashConfig
}

// LogstashConfig ...
type LogstashConfig struct {
	Network string
	Conn    string
}

// InitLog ....
func InitLog(conf *Config) error {

	if conf.Format == nil {
		conf.Format = &log.TextFormatter{ForceColors: true, FullTimestamp: true}
	}

	log.SetLevel(conf.Level)
	log.SetFormatter(conf.Format)
	log.SetOutput(os.Stdout)
	log.AddHook(&defaulthook.DefaultFieldHook{AppName: conf.Tags})
	log.AddHook(filename.NewHook())

	if conf.EnableFile {
		fileHook, err := filehook.NewLfsHook(conf.FileConfig, conf.Format)
		if err != nil {
			return err
		}
		log.AddHook(fileHook)
	}

	if conf.EnableLogstash {
		conn, err := net.Dial(conf.LogstashConf.Network, conf.LogstashConf.Conn)
		if err != nil {
			return err
		}
		log.AddHook(logrustash.New(conn, conf.Format))
	}

	return nil
}
