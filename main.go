package main

import (
	"fmt"
	"net/http"

	"github.com/kevineaton/go-battlesnake/api"
)

func main() {
	r := api.SetupApp()
	fmt.Printf("\nStarting Battlesnake API\n")
	fmt.Printf("Your snake has the following attributes:\n\tColor: %s\n\tHead: %s\n\tTail: %s\n", api.Config.SnakeColor, api.Config.SnakeHead, api.Config.SnakeTail)
	err := http.ListenAndServe(fmt.Sprintf(":%d", api.Config.APIPort), r)
	if err != nil {
		fmt.Printf("\nServer ended: %+v\n", err)
	}
}
