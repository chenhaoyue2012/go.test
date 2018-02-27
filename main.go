package main

import (
	"net/http"
	"html/template"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type	Template	struct{
	templates	*template.Template
}

func(t	*Template)	Render(w	io.Writer,	name	string,	data	interface {},	c	echo.Context)	error{
	return	t.templates.ExecuteTemplate(w,	name,	data)
}

func	Hello(c	echo.Context)	error{
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.4:3306)/blog?charset=utf8")
	// query
	rows, err := db.Query("SELECT id,name FROM bl_users")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string

		err = rows.Scan(&uid, &username)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
	}

	return	c.Render(http.StatusOK,"hello","zhangsan")
}

func Hey(c	echo.Context)	error{
	r, err := redis.Dial("tcp", "172.17.0.3:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return nil
	} else {
		fmt.Println("Connect to redis success")
	}

	username, err := redis.String(r.Do("GET", "user:name:3"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	defer r.Close()

	return c.String(http.StatusOK, fmt.Sprintf("Hello, %v!\n", username))
}

type	User	struct	{
	Name		string	`json:"name"	xml:"name"`
	Email	string	`json:"email"	xml:"email"`
}

func main() {
	t :=&Template{
		templates:	template.Must(template.ParseGlob("public/views/*.html")),
	}

	// Echo instance
	e := echo.New()
	e.Static("/",	"assets")		//静态文件
	e.Renderer	=	t

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		u	:=	&User{
			Name:		"Jon",
			Email:	"jon@labstack.com",
		}
		return	c.JSON(http.StatusOK,	u)
		//return c.String(http.StatusOK, "Hello, moon!\n")
	})

	e.GET("/hello",	Hello)
	e.GET("/hey",	Hey)


	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}