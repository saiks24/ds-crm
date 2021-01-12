package db

import "gorm.io/gorm"

type Storage interface {
}

type impl struct {
	db *gorm.DB
}

func New() Storage {
	return &impl{}
}
