package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

const DemoTableName = "t_demo"

type Demo struct {
	ID         uint64    `gorm:"primary_key,column:id" json:"id" form:"id"`
	Name       string    `gorm:"column:name" json:"name" form:"name"`
	User       *User     `gorm:"type:json;column:user" json:"user" form:"user"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time" form:"create_time"`
}

type User struct {
	ID   uint64 `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

func (s *User) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *User) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), s)
}
