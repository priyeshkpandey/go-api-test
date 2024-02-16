package goapitest

import (
	"fmt"

	"api.test.go/go-api-test/contract"
)

func ApiGet(client contract.ApiClient, body []byte) {
	fmt.Println("Creating client...")
	client.CreateClient()
	fmt.Println("Calling Get...")
	client.Get(body)
}

func ApiPost(client contract.ApiClient, body []byte) {
	fmt.Println("Creating client...")
	client.CreateClient()
	fmt.Println("Calling Post...")
	client.Post(body)
}

func ApiPut(client contract.ApiClient, body []byte) {
	fmt.Println("Creating client...")
	client.CreateClient()
	fmt.Println("Calling Put...")
	client.Put(body)
}

func ApiDelete(client contract.ApiClient, body []byte) {
	fmt.Println("Creating client...")
	client.CreateClient()
	fmt.Println("Calling Delete...")
	client.Delete(body)
}
