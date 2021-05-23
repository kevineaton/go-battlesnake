package main

import (
	"fmt"
	"net/http"

	"github.com/kevineaton/go-battlesnake/api"
)

func main() {
	r := api.SetupApp()
	port := fmt.Sprintf(":%d", api.Config.APIPort)
	fmt.Printf("\nStarting Battlesnake API\n")
	fmt.Printf("Port: %s\nAuthor: %s\n", port, api.Config.Author)
	fmt.Printf("Your snake has the following attributes:\n\tColor: %s\n\tHead: %s\n\tTail: %s\n", api.Config.SnakeColor, api.Config.SnakeHead, api.Config.SnakeTail)
	if api.Config.ShowAuth {
		fmt.Printf("\n---------------------------\n Auth Key: %s\n---------------------------\n", api.Config.AuthKey)
	}
	err := http.ListenAndServe(port, r)
	if err != nil {
		fmt.Printf("\nServer ended: %+v\n", err)
	}
}
