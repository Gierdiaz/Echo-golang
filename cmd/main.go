package main

import (
	"github.com/Gierdiaz/Echo-golang/config"
	"github.com/Gierdiaz/Echo-golang/database"
	"github.com/Gierdiaz/Echo-golang/routers"
)

func main() {
	config := config.LoadConfig()

	database.Connect(config)

	e := routers.NewRouter()
	e.Logger.Fatal(e.Start(":8080"))
}
