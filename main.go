package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed templates
var tmpl embed.FS

//go:embed static
var static embed.FS

func main() {
	r := gin.Default()
	t, _ := template.ParseFS(tmpl, "templates/*.html")
	r.SetHTMLTemplate(t)
	// r.Static("/static", "./static")
	r.StaticFS("/static", http.FS(static))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.Run()
}
