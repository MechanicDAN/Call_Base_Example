package main

import (
	"test/service"
)

func main() {
	if err := service.Cmd.Execute();
	err != nil {
		panic("main: Cant start server")
	}
}