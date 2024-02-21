package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)


func main(){
    closer, err := OpenDB()
    if err != nil {
        log.Fatal(err)
    }
    defer closer()
    fmt.Println("Database Connected!")

    router := gin.Default()

    router.LoadHTMLGlob("templates/*.html")

    router.GET("/", HomePageHandler)
    router.POST("/", ShortURLHandler)
    router.GET("/:short_url", RedirectToURL)
    router.GET("/track/:short_url", TrackURL)
    router.Run(":8080")
}
