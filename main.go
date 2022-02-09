package main

import (
	"github.com/KibetBrian/fisa/routes"
	"github.com/KibetBrian/fisa/configs"
)




func main (){
	routes.Routes();
	configs.ConnectDb();
}

