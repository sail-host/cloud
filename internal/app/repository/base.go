package repository

import "gorm.io/gorm"

type DBOption func(db *gorm.DB) *gorm.DB
