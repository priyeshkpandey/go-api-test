package main

import (
	"fmt"

	"api.test.go/go-api-test/client"
	"api.test.go/go-api-test/contract"
)

func main() {
	apClient := client.NewApiClient().WithUrl("https://reqres.in/api/users/2")
	ApiGet(apClient)
}

func ApiGet(client contract.ApiClient) {
	fmt.Println("Creating client...")
	client.CreateClient()
	fmt.Println("Calling Get...")
	client.Get([]byte(``))
}
