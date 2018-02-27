package goods

import (
	"github.com/labstack/echo"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"net/http"
)

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