package main

import "web-api-gin/router"

const PORTSERVER = ":8180"

func main() {
	router.StartServer().Run(PORTSERVER)
}
