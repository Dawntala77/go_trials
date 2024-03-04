package main

import (
	"example.com/myproject/config"
	"example.com/myproject/routes"
)

func init() {
	config.LoadEnv()
}

func main() {
	routes.Run()
}

