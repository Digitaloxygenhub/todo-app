package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Task struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Done bool   `json:"done"`
}

var tasks = []Task{}

func main() {
    router := gin.Default()

    router.GET("/tasks", func(c *gin.Context) {
        c.JSON(http.StatusOK, tasks)
    })

    router.POST("/tasks", func(c *gin.Context) {
        var newTask Task
        if err := c.ShouldBindJSON(&newTask); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        newTask.ID = len(tasks) + 1
        tasks = append(tasks, newTask)
        c.JSON(http.StatusCreated, newTask)
    })

    router.Run(":8080")
}
