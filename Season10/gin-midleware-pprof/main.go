package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	pprof.Register(engine, "devlopment/pprof")
	engine.Run(":8000")
}
