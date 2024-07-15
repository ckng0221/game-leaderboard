package main

import (
	"api/initializers"
	"api/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.ConnectToRedis()
	initializers.SynDatabase()
}

func main() {
	r := routes.SetupRouter()

	r.Run() // listen and serve on 0.0.0.0:8080
}
