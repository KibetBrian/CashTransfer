package main

import (
	"github.com/KibetBrian/fisa/configs"
	"github.com/KibetBrian/fisa/routes"
)

func main() {
	routes.Routes()
	configs.ConnectDb() 
}
