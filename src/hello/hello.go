package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	CreatedAt  time.Time
	CreatedAtS string
	CreatedBy  string
	Content    string
	Status     int
}

func main() {
	//fmt.Printf("hello, world\n")

	todo := []Todo{
		Todo{
			CreatedAt:  time.Now(),
			CreatedAtS: time.Now().Format("2006-01-02 15:04:05"),
			CreatedBy:  "Hiro",
			Content:    "Swift",
			Status:     0,
		},
		Todo{
			CreatedAt:  time.Now(),
			CreatedAtS: time.Now().Format("2006-01-02 15:04:05"),
			CreatedBy:  "Hiro",
			Content:    "Go",
			Status:     1,
		},
		Todo{
			CreatedAt:  time.Now(),
			CreatedAtS: time.Now().Format("2006-01-02 15:04:05"),
			CreatedBy:  "Hiro",
			Content:    "English",
			Status:     0,
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
	r.Run() // listen and serve on 0.0.0.0:8080
}
