package main

import (
  "fmt"
  "net/http"

  "github.com/gin-gonic/gin"
)

var userList = []User{
  User{
    FirstName: "Roman",
    LastName:  "Morozov",
    Email:     "sublimeye@gmail.com",
  }}

var router *gin.Engine

func initializeRoutes() {
  router.GET("/users", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Found users.",
      "items":   userList,
    })
  })

  router.POST("/login", func(c *gin.Context) {
    var user User
    c.BindJSON(&user)

    if isEmailExists(user.Email) {
      c.JSON(http.StatusOK, gin.H{
        "message": "login ok",
      })
    } else {
      c.JSON(http.StatusNotFound, gin.H{
        "message": "user not found, please sign up!",
      })
    }
  })

  router.POST("/signup", func(c *gin.Context) {
    var user User
    c.BindJSON(&user)

    if isEmailExists(user.Email) {
      c.JSON(http.StatusOK, gin.H{
        "message": "signup ok",
      })
      return
    }

    if isUserValid(user) {
      c.JSON(http.StatusOK, gin.H{
        "message": "signup ok",
      })
    } else {
      c.JSON(http.StatusNotFound, gin.H{
        "message": "user not found, please sign up!",
      })
    }
  })
}

func main() {
  fmt.Println("Starting...")

  router = gin.Default()

  initializeRoutes()

  router.Run()
}
