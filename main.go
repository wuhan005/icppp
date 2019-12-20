package main

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	// Load config
	conf := new(config)
	_, err := toml.DecodeFile("./config/conf.toml", &conf)
	if err != nil{
		panic(err)
	}

	// Web server start
	r := gin.Default()
	r.LoadHTMLGlob("./template/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tpl", gin.H{
			"title": c.Request.Host,
			"icp": conf.getICPNo(c.Request.Host),
		})
	})

	r.Static("/static", "./static")

	// 404 route
	r.NoRoute(func(c *gin.Context) {
		c.Header("Server", "Apache-Coyote/1.1")
		c.HTML(http.StatusOK, "404.tpl", gin.H{
			"path": c.Request.URL,
		})
	})

	panic(r.Run(conf.Server.Port))
}

func (c *config) getICPNo(url string) string{
	for _, icp := range c.ICP{
		if icp.URL == url{
			return icp.No
		}
	}
	return ""
}