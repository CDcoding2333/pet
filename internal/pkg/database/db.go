package database

import (
	"time"

	"github.com/CDcoding2333/pet/configs"
	"github.com/CDcoding2333/pet/internal/pkg/errs"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// NewDB ...
func NewDB(cnf *configs.DbConfig) (*gorm.DB, error) {
	d, err := gorm.Open(cnf.Driver, cnf.Source)
	if err != nil {
		log.Errorf("open db for %s error %s\n", cnf.Source, err.Error())
		return nil, errs.ErrDBInit
	}
	d.LogMode(cnf.LogEnabled)

	d.DB().SetMaxOpenConns(100)
	d.DB().SetMaxIdleConns(100)
	d.DB().SetConnMaxLifetime(10 * time.Second)

	return d, nil
}
