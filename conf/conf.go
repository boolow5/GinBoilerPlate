package conf

import (
	"log"

	"github.com/boolow5/GinBoilerPlate/models"
)

func InitConfig() {
	_, err1 := models.GetRoleByID(1)
	_, err2 := models.GetUserByID(1)
	if (err1 == nil) && (err2 == nil) {
		return
	}
	r := models.Role{}
	r.Name = "admin"
	r.Level = 1
	if _, err := r.Add(); err != nil {
		log.Println("InitConfig: Add Role error", err)
	}
	u := models.User{}
	u.FirstName = "Mahdi"
	u.MiddleName = "Ahmed"
	u.LastName = "Bolow"
	u.Username = "boolow5"
	u.Password = "sharaf"
	u.Role = r
	u.AppID = "143"
	if _, err := u.Add(); err != nil {
		log.Println("InitConfig: Add User error", err)
	}
}
