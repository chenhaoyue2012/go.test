package main

import (
	"net/http"
	"html/template"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io"
)


type	Template	struct{
	templates	*template.Template
}

func(t	*Template)	Render(w	io.Writer,	name	string,	data	interface {},	c	echo.Context)	error{
	return	t.templates.ExecuteTemplate(w,	name,	data)
}

func	Hello(c	echo.Context)	error{
	return	c.Render(http.StatusOK,"hello","zhangsan")
}


func main() {
	t :=&Template{
		templates:	template.Must(template.ParseGlob("public/views/*.html")),
	}

	// Echo instance
	e := echo.New()
	e.Renderer	=	t

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, moon!\n")
	})

	e.GET("/hello",	Hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}