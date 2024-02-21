package main

import (
	cli "github.com/cleopatrio/cli/cli/cmd"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cli.Execute()
}
