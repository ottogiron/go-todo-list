package main

import (
	"fmt"

	"github.com/ottogiron/chapi/server"
	"github.com/ottogiron/gotodo/controllers"
)

func main() {
	connectionString := ":3001"
	s := server.NewServer()
	s.Register(&controllers.Todos{})
	runError := s.Run(connectionString)
	if runError != nil {
		fmt.Println("Error when running server", runError)
	}

}
