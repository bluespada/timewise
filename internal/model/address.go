// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file is contains gorm model

package model

type ModelAddress struct {
	*Model
	UserId   uint   `json:"user_id"`
	Address  string `json:"address"`
	Province string `json:"province"`
	City     string `json:"city"`
	ZipCode  string `json:"zip_code"`
	Country  string `json:"country"`
}
