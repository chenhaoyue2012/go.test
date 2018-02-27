package main

import (
	"html/template"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io"
	_ "github.com/go-sql-driver/mysql"
	"go.test/controller/goods"
	"go.test/controller/moon"
	"go.test/controller/user"
	"go.test/config"
	"fmt"
)



type	Template	struct{
	templates	*template.Template
}

func(t	*Template)	Render(w	io.Writer,	name	string,	data	interface {},	c	echo.Context)	error{
	return	t.templates.ExecuteTemplate(w,	name,	data)
}


func main() {
	t :=&Template{
		templates:	template.Must(template.ParseGlob("views/*.html")),
	}

	// Echo instance
	e := echo.New()
	e.HideBanner = true
	e.Static("/",	"assets")		//静态文件
	e.Renderer	=	t

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", user.Index)

	e.GET("/hello",	moon.Hello)
	e.GET("/hey",	goods.Hey)


	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", config.Host, config.Port)))
}