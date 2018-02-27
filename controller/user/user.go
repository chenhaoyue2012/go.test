package user

import (
	"github.com/labstack/echo"
	"net/http"
)

type	User	struct	{
	Name		string	`json:"name"	xml:"name"`
	Email	string	`json:"email"	xml:"email"`
}

func Index(c echo.Context) error {
	u	:=	&User{
		Name:		"Jon",
		Email:	"jon@labstack.com",
	}
	return	c.JSON(http.StatusOK,	u)
	//return c.String(http.StatusOK, "Hello, moon!\n")
}
