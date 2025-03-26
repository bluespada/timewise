// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file is contains type declaration for security access.

package types

type UserAccessRole string

const (
	UserAccessRoleSuperAdmin   UserAccessRole = "super_admin"
	UserAccessRoleAdmin        UserAccessRole = "admin"
	UserAccessRolePayrollAdmin UserAccessRole = "payroll_admin"
	UserAccessRoleManager      UserAccessRole = "manager"
	UserAccessRoleEmployee     UserAccessRole = "employee"
)
