package main

import (
	"fmt"
	"net/http"
	response "section_api/handlerResponse"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	// Cek apakah panjang users == 0
	if len(users) == 0 {
		return response.Response(c, http.StatusOK, "Success get all users", []User{})
	}

	return response.Response(c, http.StatusOK, "Success get all users", users)
}

// get user by id
func GetUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	// Cek jika param error karena invalid dtype
	if err != nil {
		return response.Response(c, http.StatusBadRequest, "Invalid argument type", nil)
	}

	for _, user := range users {
		// Cek jika id user == id param
		if user.Id == userId {
			return response.Response(c, http.StatusOK, "Success get user", users[userId-1])
		}
	}

	return response.Response(c, http.StatusNotFound, fmt.Sprintf("User id %v not found", userId), nil)
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	// Cek jika param error karena invalid dtype
	if err != nil {
		return response.Response(c, http.StatusBadRequest, "Invalid argument type", nil)
	}

	newUsers := make([]User, 0)

	for i, user := range users {
		// Cek jika user id == param id
		if user.Id == userId {
			// Menambah data baru ke newUsers
			newUsers = append(newUsers, users[i+1:]...)
			newUsers = append(newUsers, users[:i]...)

			users = newUsers

			return response.Response(c, http.StatusOK, fmt.Sprintf("Success delete User id %v", userId), newUsers)
		}
	}

	return response.Response(c, http.StatusNotFound, fmt.Sprintf("User id %v not found", userId), nil)
}

// update user by id
func UpdateUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	// Cek jika param error karena invalid dtype
	if err != nil {
		return response.Response(c, http.StatusBadRequest, "Invalid argument type", nil)
	}

	for i, user := range users {
		// Cek jika user id == param id
		if user.Id == userId {
			var u User

			// Dilakukan binding data ke variable u dengan tipe User
			err = c.Bind(&u)
			if err != nil {
				return response.Response(c, http.StatusBadRequest, err.Error(), nil)
			}

			u.Id = users[i].Id

			// Replace nilai dari user lama dengan data yang baru
			users[i] = u

			return response.Response(c, http.StatusOK, fmt.Sprintf("Success update User id %v", userId), users)
		}
	}

	return response.Response(c, http.StatusNotFound, fmt.Sprintf("User id %v not found", userId), nil)
}

// create new user
func CreateUserController(c echo.Context) error {
	fmt.Println("masuk route post")
	// binding data
	user := User{}
	err := c.Bind(&user)
	fmt.Println("setelah data binding :", user)
	fmt.Println(err)
	if err != nil {
		return response.Response(c, http.StatusBadRequest, err.Error(), nil)
	}

	// Cek jika length users == 0
	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	// Validasi email memerlukan @ pada inputan
	if result := strings.Contains(user.Email, "@"); !result {
		return response.Response(c, http.StatusBadRequest, "Invalid email format", user)
	}

	users = append(users, user)
	return response.Response(c, http.StatusOK, "Success create User", user)
}

// Endpoint yang pertama kali akan dibuka
func root(c echo.Context) error {
	return c.String(http.StatusOK, "OK!")
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/", root)
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
