package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Message string `json:"message"`
}

func sayHello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, Message{
		Message: "Hi, you are using the Echo framework",
	})
}

func sayHelloWithParam(ctx echo.Context) error {
	name := ctx.Param("name")
	return ctx.JSON(http.StatusOK, Message{
		Message: "Hi, my name is " + name,
	})
}

func sayHelloWithQuery(ctx echo.Context) error {
	name := ctx.QueryParam("name")
	if name == "" {
		return ctx.JSON(http.StatusBadRequest, Message{
			Message: "Name query parameter is missing",
		})
	}
	return ctx.JSON(http.StatusOK, Message{
		Message: "Hi, this is Echo framework, and I am " + name,
	})
}

func main() {
	e := echo.New()

	e.GET("/hello", sayHello)
	e.GET("/paramhello/:name", sayHelloWithParam)
	e.GET("/queryhello", sayHelloWithQuery) // Fixed typo

	port := ":8080"
	log.Println("Server is running on port", port)
	if err := e.Start(port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
