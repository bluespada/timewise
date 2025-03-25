// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file contains code for GraphQL Playground
package playground

import (
	"embed"
)

//go:embed static/*
var GraphFsStatic embed.FS
