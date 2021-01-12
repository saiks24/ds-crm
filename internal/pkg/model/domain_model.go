package model

import (
	"time"

	"gorm.io/gorm"
)

type ActionType string

type Task struct {
	gorm.Model
	Name           string
	CreatorID      int64
	ExecutorID     int64
	DeletedAt      time.Time
	FlowTypeID     int64
	CurrentStateID int64
}

type Comment struct {
	gorm.Model
	UserID int64
	TaskID int64
	Text   string
}

//type Flow struct {
//	ID int64
//	Name string
//	FromState string
//	ToState string
//}

type FlowPartTemplate struct {
	gorm.Model
	Name     string
	Settings map[string]interface{}
	UserID   int64
}

type FlowPart struct {
	gorm.Model
	Name   string
	FlowID int64
	From   int64
	To     int64
	UserID int64
}

type Log struct {
	gorm.Model
	Name       string
	ActionType ActionType
	TaskID     int64
	UserID     int64
}

type File struct {
	gorm.Model
	ID     int64
	Name   string
	URL    string
	TaskID int64
	UserID int64
}

type User struct {
	gorm.Model
	Name string
}

type Role struct {
	gorm.Model
	Name string
}
