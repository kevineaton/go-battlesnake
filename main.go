package main

import (
	"fmt"
	"net/http"

	"github.com/kevineaton/go-battlesnake/api"
)

func main() {
	r := api.SetupApp()
	fmt.Printf("\nStarting Battlesnake API\n")
	err := http.ListenAndServe(fmt.Sprintf(":%d", api.Config.BS_API_PORT), r)
	if err != nil {
		fmt.Printf("\nServer ended: %+v\n", err)
	}
}
