package dto

import "kube_web/models"

type LoginRes struct {
	Token string         `json:"token"`
	User *models.SysUser `json:"user"`
}
