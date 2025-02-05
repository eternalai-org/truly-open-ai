package models

import (
	"github.com/jinzhu/gorm"
)

type SampleTwitterApp struct {
	gorm.Model
	InstallCode   string
	ApiKey        string
	TwitterInfoID uint
	TwitterInfo   *TwitterInfo
}
