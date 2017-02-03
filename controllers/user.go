package controllers

import (
	"strconv"

	"github.com/boolow5/GinBoilerPlate/models"
	"github.com/gin-gonic/gin"
)

func AllUsers(this *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	}
	this.JSON(200, gin.H{"users": users})
}

func GetUser(this *gin.Context) {
	user_id, err := strconv.Atoi(this.Query("user_id"))
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	}
	user, err := models.GetUserByID(user_id)
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	}
	this.JSON(200, gin.H{"user": user})
}

func AddUser(this *gin.Context) {
	user := models.User{}
	err := this.BindJSON(&user)
	if err != nil {
		this.JSON(404, gin.H{"error": err})
		return
	}
	saved, err := user.Add()
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	} else if saved {
		this.JSON(200, gin.H{"success": "Saved user successfully"})
		return
	}
	this.JSON(200, gin.H{"warning": "User might not be saved! but there was no errors reported"})
}

func UpdateUser(this *gin.Context) {
	user_id, _ := strconv.Atoi(this.Query("user_id"))
	if user_id < 1 {
		this.JSON(404, gin.H{"error": "Invalid user id"})
		return
	}

	updatedUser, err := models.GetUserByID(user_id)
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	}
	err = this.BindJSON(&updatedUser)
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if updatedUser.ID != user_id {
		this.JSON(404, gin.H{"error": "there is user id mismatch!"})
		return
	}
	updatedUser.ID = user_id
	updated, err := updatedUser.Update()
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if updated {
		this.JSON(200, gin.H{"success": "updated user successfully"})
		return
	}
	this.JSON(200, gin.H{"warning": "updating user failed"})
}

func UpdateUserPassword(this *gin.Context) {
	user_id, _ := strconv.Atoi(this.Query("user_id"))
	if user_id < 1 {
		this.JSON(404, gin.H{"error": "Invalid user id"})
		return
	}

	oldUser, err := models.GetUserByID(user_id)
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	}
	updatedUser := struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}{}
	err = this.BindJSON(&updatedUser)
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	}

	if oldUser.ID != user_id {
		this.JSON(404, gin.H{"error": "there is user id mismatch!"})
		return
	}
	//PasswordsMatch(unhashedPassword, hashedPassword string)
	if models.PasswordsMatch(updatedUser.OldPassword, oldUser.Password) == false {
		this.JSON(404, gin.H{"error": "Passwords don't match"})
		return
	}
	oldUser.SetPassword(updatedUser.NewPassword)

	updated, err := oldUser.Update()
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if updated {
		this.JSON(200, gin.H{"success": "updated user successfully"})
		return
	}
	this.JSON(200, gin.H{"warning": "updating user failed"})
}

func DeleteUser(this *gin.Context) {
	user_id, _ := strconv.Atoi(this.Query("user_id"))
	if user_id < 1 {
		this.JSON(404, gin.H{"error": "Invalid user id"})
		return
	} else {
		user, err := models.GetUserByID(user_id)
		if err != nil {
			this.JSON(404, gin.H{"error": err.Error()})
			return
		} else {
			deleted, err := user.Delete()
			if err != nil {
				this.JSON(404, gin.H{"error": err.Error()})
				return
			} else if deleted {
				this.JSON(200, gin.H{"success": "deleted user successfully"})
				return
			}
		}
	}
	this.JSON(200, gin.H{"warning": "user might not be deleted successfully"})
}
