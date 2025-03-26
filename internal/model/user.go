// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file is contains gorm model

package model

type ModelUsers struct {
	*Model
	Name    string         `json:"name"`
	Picture string         `json:"picture"`
	Phone   string         `json:"phone" gorm:"unique"`
	Email   string         `json:"email" gorm:"unique"`
	Address []ModelAddress `gorm:"foreignKey:UserId"`
}
