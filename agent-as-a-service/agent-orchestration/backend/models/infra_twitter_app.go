package models

import (
	"github.com/jinzhu/gorm"
)

type InfraTwitterApp struct {
	gorm.Model
	Address       string
	InstallCode   string
	Signature     string
	TwitterInfoID uint
	TwitterInfo   *TwitterInfo
}
