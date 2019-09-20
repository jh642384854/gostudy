package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)
/**
	定义多服务
 */
var g errgroup.Group

func main() {
	server01 := &http.Server{
		Addr:":8080",
		Handler:router01(),
		ReadTimeout:5*time.Second,
		WriteTimeout:5*time.Second,
	}
	server02 := &http.Server{
		Addr:":8081",
		Handler:router02(),
		ReadTimeout:5*time.Second,
		WriteTimeout:5*time.Second,
	}
	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait();err != nil{
		log.Fatal(err)
	}
}

//定义一个服务1
func router01() http.Handler  {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"code":http.StatusOK,
			"msg":"Welcome server 01",
		})
	})
	return e
}

//定义一个服务2
func router02() http.Handler  {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"code":http.StatusOK,
			"msg":"Welcome server 02",
		})
	})
	return e
}
