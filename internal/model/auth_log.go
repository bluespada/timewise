// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file is contains gorm model

package model

import (
	"time"
)

type ModelAuthLog struct {
	*Model
	AuthId       uint      `json:"auth_id"`
	Auth         ModelAuth `json:"_" gorm:"foreignKey:AuthId"`
	Ip           string    `json:"ip" gorm:"notnull"`
	Time         time.Time `json:"time" gorm:"default:now()"`
	UserAgent    string    `json:"user_agent" gorm:"notnull"`
	Status       string    `json:"status" gorm:"notnull"`
	ErrorMessage string    `json:"error_message"`
}
