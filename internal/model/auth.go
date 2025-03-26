// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file is contains gorm model
package model

type Roles []string

type ModelAuth struct {
	*Model
	Email    string     `json:"email" gorm:"unique"`
	Password string     `json:"password"`
	Phone    string     `json:"phone" gorm:"unique"`
	UserId   uint       `json:"user_id"`
	User     ModelUsers `json:"_" gorm:"foreignKey:UserId"`
	Roles    Roles      `json:"roles" gorm:"type:text[];notnull"`
}
