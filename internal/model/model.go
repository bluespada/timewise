// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file is contains base model for timewise

package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	*gorm.Model
	ID uint `json:"id" gorm:"primaryKey"`

	CreatedAt time.Time       `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"default:now()"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}
