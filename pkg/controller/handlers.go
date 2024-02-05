package controller

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"github.com/chaitanyakolluru/go-ums-backend/pkg/model"
	"github.com/labstack/echo/v4"
)

func SaveUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	for _, user := range model.Users {
		if user.User.Name == u.Name {
			return c.JSON(http.StatusConflict, "user already exists")
		}
	}

	c.Get("db").(*gorm.DB).Create(&model.UserData{User: *u})
	return c.JSON(http.StatusCreated, "user created")
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	for _, user := range model.Users {
		if id == fmt.Sprintf("%d", user.ID) {
			c.Get("db").(*gorm.DB).Delete(&user, id)
			return c.JSON(http.StatusOK, "user deleted")
		}
	}

	return c.JSON(http.StatusNotFound, "user not found")
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	for _, user := range model.Users {
		if id == fmt.Sprintf("%d", user.ID) {
			c.Get("db").(*gorm.DB).Model(&user).Updates(model.UserData{User: *u})
			return c.JSON(http.StatusOK, "user updated")
		}
	}

	return c.JSON(http.StatusNotFound, "user not found")
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, model.Users)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")

	for _, u := range model.Users {
		if id == fmt.Sprintf("%d", u.ID) {
			c.Get("db").(*gorm.DB).First(&u, id)
			return c.JSON(http.StatusOK, u)
		}
	}

	return c.JSON(http.StatusNotFound, "user not found")
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "status: ok")
}
