package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/view/", viewHandler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// // take in response and give out request
// func handler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 	http.NotFound(w, r)
// 	return
// 	}
// 	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
// }

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/view/"):]
// 	p, _ := loadPage(title)
// 	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
// }

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()
// 	r.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "Hello, World!")
// 	})
// 	r.Run(":8080")
// }

// func handler(c *gin.Context) {
// 	c.String(http.StatusOK, "Hello, World!")
// }

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8080")
}
