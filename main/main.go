package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func handleGetTasks(c *gin.Context) {
	var loadedTasks, err = GetAllTasks()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": loadedTasks})
}

func handleGetTask(c *gin.Context) {
	var task Task
	if err := c.BindUri(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	var loadedTask, err = GetTaskByID(task.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ID": loadedTask.ID, "Body": loadedTask.Body})
}

func handleCreateTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := Create(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func handleUpdateTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	savedTask, err := Update(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": savedTask})
}

func main() {

	//os.Setenv("PORT", "32432")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := gin.Default()
	r.GET("/tasks/:id", handleGetTask)
	r.GET("/tasks/", handleGetTasks)
	r.PUT("/tasks/", handleCreateTask)
	r.POST("/tasks/", handleUpdateTask)
	r.Run(":8090") // listen and serve on 0.0.0.0:8080

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "okay"})

		return
	})

	r.Run(":" + port)

}
