package main

import (
	"T4/cmd"
	"T4/model"
)

func main() {
	model.InitDB()
	cmd.InitRouter()
}
