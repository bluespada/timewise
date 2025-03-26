// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file is contains helper for database

package database

import (
	"log"
	"os"

	"github.com/bluespada/timewise/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Initialize() {

	var err error
	var Dsn string

	if os.Getenv("APP_DB") != "" {
		Dsn = os.Getenv("APP_DB")
	} else {
		Dsn = "host=localhost user=postgres password=postgres dbname=timewise port=5432 sslmode=disable"
	}

	log.Println(Dsn)

	Db, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN: Dsn,
		}),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	// register model object here.
	Db.AutoMigrate(
		&model.ModelAuth{},
	)
}
