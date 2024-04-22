package db

import (
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBconn *gorm.DB
)
