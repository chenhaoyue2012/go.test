package moon

import (
	"github.com/labstack/echo"
	"database/sql"
	"fmt"
	"net/http"
	"go.test/lib"
)

func Hello(c	echo.Context)	error{
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.4:3306)/blog?charset=utf8")
	// query
	rows, err := db.Query("SELECT id,name FROM bl_users")
	lib.Checkerr(err)

	for rows.Next() {
		var uid int
		var username string

		err = rows.Scan(&uid, &username)
		lib.Checkerr(err)
		fmt.Println(uid)
		fmt.Println(username)
	}

	return	c.Render(http.StatusOK,"hello","zhangsan")
}