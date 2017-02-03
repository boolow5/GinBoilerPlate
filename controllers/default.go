package controllers

import (
	"strconv"

	"github.com/boolow5/GinBoilerPlate/models"
	"github.com/gin-gonic/gin"
)

func Index(this *gin.Context) {
	/*this.HTML(200, "index.tmpl", gin.H{
		"title": "Vue + GIN project",
	})*/
	this.JSON(200, gin.H{
		"message":            "Welcome to BolAuth.",
		"description":        "Micro service for authentication and authorization for applications build by BoSS",
		"available_versions": []string{"v1"},
		"last_updated":       "14.01.2017",
	})
}

func AllTodos(this *gin.Context) {
	todos, err := models.GetAllTodos()
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	}
	this.JSON(200, gin.H{"todos": todos})
}

func AddTodo(this *gin.Context) {
	todo := models.Todo{}
	err := this.BindJSON(&todo)
	if err != nil {
		this.JSON(404, gin.H{"error": err})
	}
	saved, err := todo.Add()
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
	} else if saved {
		this.JSON(200, gin.H{"success": "Saved todo successfully"})
	}
	this.JSON(200, gin.H{"warning": "Todo might not be saved! but there was no errors reported"})
}

func UpdateTodo(this *gin.Context) {
	todo_id, _ := strconv.Atoi(this.Query("todo_id"))
	if todo_id < 1 {
		this.JSON(404, gin.H{"error": "Invalid todo id"})
		return
	}

	updatedTodo, err := models.GetTodoByID(todo_id)
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	}
	err = this.BindJSON(&updatedTodo)
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if updatedTodo.ID != todo_id {
		this.JSON(404, gin.H{"error": "there is todo id mismatch!"})
		return
	}
	updatedTodo.ID = todo_id
	updated, err := updatedTodo.Update()
	if err != nil {
		this.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if updated {
		this.JSON(200, gin.H{"success": "updated todo successfully"})
		return
	}
	this.JSON(200, gin.H{"warning": "updating todo failed"})
}

func DeleteTodo(this *gin.Context) {
	todo_id, _ := strconv.Atoi(this.Query("todo_id"))
	if todo_id < 1 {
		this.JSON(404, gin.H{"error": "Invalid todo id"})
		return
	} else {
		todo, err := models.GetTodoByID(todo_id)
		if err != nil {
			this.JSON(404, gin.H{"error": err.Error()})
			return
		} else {
			deleted, err := todo.Delete()
			if err != nil {
				this.JSON(404, gin.H{"error": err.Error()})
				return
			} else if deleted {
				this.JSON(200, gin.H{"success": "deleted todo successfully"})
				return
			}
		}
	}
	this.JSON(200, gin.H{"warning": "todo might not be deleted successfully"})
}
