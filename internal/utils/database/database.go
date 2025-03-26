// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file is contains helper for database

package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

// RegisterModels register models to database.
//
// This function will register models to gorm database and create table if not exists.
// It will use default gorm settings. If you want to customize the table name or
// column name, you can use gorm.Model and gorm.Column functions.
//
// For example:
//
//	type Model struct {
//		gorm.Model
//		Column string `gorm:"column:column_name"`
//	}
//
//	RegisterModels(&Model{})
func RegisterModels(models ...interface{}) {
	Db.AutoMigrate(models...)
}

// Initialize initialize database connection. It will use DSN connection string from
// environment variable APP_DB. If APP_DB not set, it will use default connection string.
// The default connection string is :
// host=localhost user=postgres password=postgres dbname=timewise port=5432 sslmode=disable
func Initialize() {

	var err error
	var Dsn string

	if os.Getenv("APP_DB") != "" {
		Dsn = os.Getenv("APP_DB")
	} else {
		Dsn = "host=localhost user=postgres password=postgres dbname=timewise port=5432 sslmode=disable"
	}

	// TODO: need to add more connection options
	Db, err = gorm.Open(
		postgres.New(postgres.Config{
			DSN: Dsn,
		}),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}
}
