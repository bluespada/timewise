// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file is contains type declaration for API.

package types

type ApiResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewApiResponse() *ApiResponse {
	return &ApiResponse{
		Error:   false,
		Message: "",
		Data:    nil,
	}
}
