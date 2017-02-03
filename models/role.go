package models

import (
	"errors"
	"fmt"
	"time"
)

type Role struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Level     int       `json:"level"` // 1. sys-admin 2. app-admin 3. manager 4. employee 5. temp-employee
	TimeAdded time.Time `json:"time_added"`
}

func (this *Role) String() string {
	if len(this.Name) > 0 {
		return fmt.Sprintf("%s", this.Name)
	}
	return fmt.Sprintf("%s", this.Name)
}

func (this *Role) Add() (bool, error) {
	this.TimeAdded = time.Now()
	result := db.Create(this)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func GetRoleByID(id int) (*Role, error) {
	Role := Role{}
	result := db.Where("id = ?", id).First(&Role)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Role, nil
}

func GetAllRoles() (*[]Role, error) {
	todos := []Role{}
	result := db.Find(&todos)
	if result.Error != nil {
		return &[]Role{}, result.Error
	}
	return &todos, nil
}

func (this *Role) Update() (bool, error) {
	if this.Name == "" || this.Name == " " {
		return false, errors.New("Empty Name")
	}
	result := db.Save(this)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (this *Role) Delete() (bool, error) {
	result := db.Delete(this)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
