package main

import "github.com/adharshmk96/shAuth/cmd"

// @title			ServiceHub Account API
// @version		1.0
// @description	This is account service for ServiceHub
// @host			localhost:8080
// @BasePath		/
// @schemes		http
// @produce		json
// @consumes		json
func main() {
	cmd.Execute()
}
