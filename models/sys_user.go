package models

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(255);uniqueIndex:idx_name_tenant;comment:用户名"`
	Password string `json:"password" gorm:"type:varchar(255);comment:密码"`
	Avatar   string `json:"avatar" gorm:"type:varchar(255);comment:头像"`
	Nickname string `json:"nickname" gorm:"type:varchar(255);comment:昵称"`
	Status   int8   `json:"status" gorm:"default:1;comment:状态 0:禁用 1:启用"`
	ResModel
}
