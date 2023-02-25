package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgre struct {
	Conn *gorm.DB
}

var PostgreDB Postgre

func (p *Postgre) Connect() {
	conn, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	PostgreDB.Conn = conn
}
