package model

import (
	"flow/enum"
	"github.com/jinzhu/gorm"
	"time"
)

type Section struct {
	gorm.Model
	Id        int    `gorm:"primary_key";"AUTO_INCREMENT"`
	Name      string `gorm:"type:varchar(200)"`
	Status    enum.Status
	IsVisible bool
	TenantId  string
	CreatedOn time.Time
	DeletedOn time.Time
}
