package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID         int    `form:"ida"`
	CreatedBy  string `form:"createdby"`
	Content    string `form:"content"`
	Status     int    `form:"status"`
	CreatedAt  time.Time
	CreatedAtS string
}

var todo []Todo
var idMax = 0

func Saiban() int {
	idMax = idMax + 1
	return idMax
}

func GetDataTodo(c *gin.Context) {
	var b Todo
	if err := c.Bind(&b); err != nil {
		fmt.Errorf("%#v", err)
	}
	b.ID = Saiban()
	b.Status = 0
	b.CreatedAtS = time.Now().Format("2006-01-02 15:04:05")
	todo = append(todo, b)
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"todo": todo,
	})
}

func GetDoneTodo(c *gin.Context) {
	var b Todo
	if err := c.Bind(&b); err != nil {
		fmt.Errorf("%#v", err)
	}

	var s int

	if b.Status == 0 {
		s = 1
	} else {
		s = 0
	}

	for idx, t := range todo {
		if t.ID == b.ID {
			todo[idx].Status = s
		}
	}
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"todo": todo,
	})
}

func main() {
	//fmt.Printf("hello, world\n")

	todo = []Todo{
		Todo{
			ID:         Saiban(),
			CreatedBy:  "Hiro",
			Content:    "Swift",
			Status:     0,
			CreatedAt:  time.Now(),
			CreatedAtS: time.Now().Format("2006-01-02 15:04:05"),
		},
		Todo{
			ID:         Saiban(),
			CreatedBy:  "Hiro",
			Content:    "Go",
			Status:     1,
			CreatedAt:  time.Now(),
			CreatedAtS: time.Now().Format("2006-01-02 15:04:05"),
		},
		Todo{
			ID:         Saiban(),
			CreatedBy:  "Hiro",
			Content:    "English",
			Status:     0,
			CreatedAt:  time.Now(),
			CreatedAtS: time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	r := gin.Default()
	r.LoadHTMLFiles("./tmpl/index.html")
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "world",
		})
	})
	r.GET("/todo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"todo": todo,
		})
	})
	r.GET("/todo2", GetDataTodo)
	r.GET("/done", GetDoneTodo)
	r.Run() // listen and serve on 0.0.0.0:8080
}
