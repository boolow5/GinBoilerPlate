package controllers

import (
	"strconv"

	"github.com/boolow5/GinBoilerPlate/models"
	"github.com/gin-gonic/gin"
)

func AllRoles(this *gin.Context) {
	roles, err := models.GetAllRoles()
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	}
	this.JSON(200, gin.H{"roles": roles})
}

func AddRole(this *gin.Context) {
	role := models.Role{}
	err := this.BindJSON(&role)
	if err != nil {
		this.JSON(404, gin.H{"error": err})
	}
	saved, err := role.Add()
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	} else if saved {
		this.JSON(200, gin.H{"success": "Saved role successfully"})
		return
	}
	this.JSON(200, gin.H{"warning": "Role might not be saved! but there was no errors reported"})
}

func UpdateRole(this *gin.Context) {
	role_id, _ := strconv.Atoi(this.Query("role_id"))
	if role_id < 1 {
		this.JSON(404, gin.H{"error": "Invalid role id"})
		return
	}

	updatedRole, err := models.GetRoleByID(role_id)
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	}
	err = this.BindJSON(&updatedRole)
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if updatedRole.ID != role_id {
		this.JSON(404, gin.H{"error": "there is role id mismatch!"})
		return
	}
	updatedRole.ID = role_id
	updated, err := updatedRole.Update()
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if updated {
		this.JSON(200, gin.H{"success": "updated role successfully"})
		return
	}
	this.JSON(200, gin.H{"warning": "updating role failed"})
}

func DeleteRole(this *gin.Context) {
	role_id, _ := strconv.Atoi(this.Query("role_id"))
	if role_id < 1 {
		this.JSON(404, gin.H{"error": "Invalid role id"})
		return
	} else {
		role, err := models.GetRoleByID(role_id)
		if err != nil {
			this.JSON(404, gin.H{"error": err.Error()})
			return
		} else {
			deleted, err := role.Delete()
			if err != nil {
				this.JSON(404, gin.H{"error": err.Error()})
				return
			} else if deleted {
				this.JSON(200, gin.H{"success": "deleted role successfully"})
				return
			}
		}
	}
	this.JSON(200, gin.H{"warning": "role might not be deleted successfully"})
}
