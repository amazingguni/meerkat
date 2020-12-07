package model

import (
	"github.com/jinzhu/gorm"
)

type Metric struct {
	gorm.Model
	Value float32
}
