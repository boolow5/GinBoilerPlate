package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "auth.db")
	if err != nil {
		panic("Error Connecting Database " + err.Error())
	}
	db.AutoMigrate(&Role{}, &User{})
}

type Todo struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	TimeAdded time.Time `json:"time_added"`
	Done      bool      `json:"done"`
}

func (t *Todo) String() string {
	if t.Done {
		return fmt.Sprintf("Did: %s ", t.Title)
	}
	return fmt.Sprintf("Do: %s ", t.Title)
}

func (t *Todo) Add() (bool, error) {
	t.TimeAdded = time.Now()
	result := db.Create(t)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func GetTodoByID(id int) (*Todo, error) {
	todo := Todo{}
	result := db.Where("id = ?", id).First(&todo)
	if result.Error != nil {
		return &Todo{}, result.Error
	}
	return &todo, nil
}

func GetAllTodos() (*[]Todo, error) {
	todos := []Todo{}
	result := db.Find(&todos)
	if result.Error != nil {
		return &[]Todo{}, result.Error
	}
	return &todos, nil
}

func (t *Todo) Update() (bool, error) {
	if t.Title == "" || t.Title == " " {
		return false, errors.New("Empty Title")
	}
	result := db.Save(t)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (t *Todo) Delete() (bool, error) {
	result := db.Delete(t)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
