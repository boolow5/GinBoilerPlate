package models

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         int       `json:"id" gorm:"primary_key"`
	FirstName  string    `json:"first_name" binding:"required"`
	MiddleName string    `json:"middle_name"`
	LastName   string    `json:"last_name" binding:"required"`
	Username   string    `json:"username" binding:"required" gorm:"unique"`
	Password   string    `json:"password" binding:"required"`
	Role       Role      `json:"role"`
	RoleID     int       `json:"role_id" binding:"required"`
	TimeAdded  time.Time `json:"time_added"`
	AppID      string    `json:"app_id" binding:"required"`
	Active     bool      `json:"active"`
}

func (this *User) String() string {
	if len(this.FirstName) > 0 {
		return fmt.Sprintf("%s %s %s", this.FirstName, this.MiddleName, this.LastName)
	}
	return fmt.Sprintf("%s", this.Username)
}

func AuthenticateUser(username, password, app_id string) (*User, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if (user.Username == username) && checkPassword(password, user.Password) && (user.AppID == app_id) {
		return user, nil
	}
	return nil, errors.New("Invalid user")
}

func hashPassword(password string) (string, error) {
	if len(password) > 59 {
		return password, nil
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		return password, err
	}
	return string(hashedPass), nil
}
func checkPassword(yourPassword, ourPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(ourPassword), []byte(yourPassword))
	if err != nil {
		return false
	}
	return true
}

func PasswordsMatch(unhashedPassword, hashedPassword string) bool {
	return checkPassword(unhashedPassword, hashedPassword) == true
}

func (this *User) SetPassword(newPassword string) bool {
	var err error
	this.Password, err = hashPassword(newPassword)
	return (err == nil)
}

func (this *User) DoesExist() bool {
	counter := 0
	db.Where("username = ? ", this.Username).Find(this).Count(&counter)
	if counter > 0 {
		return true
	}
	return false
}

func (this *User) Add() (bool, error) {
	var err error
	this.TimeAdded = time.Now()
	this.Password, err = hashPassword(this.Password)
	if err != nil {
		return false, err
	}
	if this.DoesExist() {
		return false, errors.New("Username already exists")
	}
	result := db.Create(this)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func GetUserByID(id int) (*User, error) {
	user := User{}
	result := db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	role := Role{}
	result = db.Where("id = ?", user.RoleID).First(&role)
	if result.Error != nil {
		return &user, result.Error
	}
	user.Role = role
	return &user, nil
}

func GetUserByUsername(username string) (*User, error) {
	user := User{}
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return &user, result.Error
	}
	role := Role{}
	result = db.Where("id = ?", user.RoleID).First(&role)

	if result.Error != nil {
		return &user, result.Error
	}
	user.Role = role
	return &user, nil
}

func GetAllUsers() (*[]User, error) {
	todos := []User{}
	result := db.Find(&todos)
	if result.Error != nil {
		return &[]User{}, result.Error
	}
	return &todos, nil
}

func (this *User) Update() (bool, error) {
	if this.FirstName == "" || this.FirstName == " " {
		return false, errors.New("Empty First name")
	}
	if this.LastName == "" || this.LastName == " " {
		return false, errors.New("Empty Last name")
	}
	result := db.Save(this)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (this *User) Delete() (bool, error) {
	result := db.Delete(this)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
