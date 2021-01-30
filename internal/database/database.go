package database

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDatabase struct {
	DB *gorm.DB
}

func New(dsn string, gormConfig *gorm.Config) (*GormDatabase, error) {
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}

	return &GormDatabase{DB: db}, nil
}

func (d *GormDatabase) AutoMigrate() {
	// enable format UUID as PK
	d.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	if err := d.DB.AutoMigrate(
		//&model.Todo{},
	); err != nil {
		log.Fatal(err.Error())
	}
}

func (d *GormDatabase) DropAllTables() {
	if err := d.DB.Migrator().DropTable(
		"todos",
	); err != nil {
		log.Fatal(err.Error())
	}
}
