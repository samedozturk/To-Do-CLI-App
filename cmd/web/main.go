package main

import (
	"github.com/samedozturk/To-Do-CLI-App/config"
	"github.com/samedozturk/To-Do-CLI-App/internal/storage/mongodb"
)

func main() {
	config.LoadEnv()
	mongodb.SetupDB("exampleDB")
}
