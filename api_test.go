package goapitest

import (
	"testing"

	"api.test.go/go-api-test/client"
)

func TestGet(t *testing.T) {
	apiClientGet := client.NewApiClient().WithUrl("https://reqres.in/api/users/2")
	ApiGet(apiClientGet, []byte(``))
	t.Log("Get is successful")
}

func TestPost(t *testing.T) {
	apiClientPost := client.NewApiClient().WithUrl("https://reqres.in/api/users")
	postReqBody := `{"name": "morpheus", "job": "leader}`
	ApiPost(apiClientPost, []byte(postReqBody))
	t.Log("Post is successful")
}

func TestPut(t *testing.T) {
	apiClientPut := client.NewApiClient().WithUrl("https://reqres.in/api/users/2")
	putReqBody := `{"name": "morpheus", "job": "zion resident"}`
	ApiPut(apiClientPut, []byte(putReqBody))
	t.Log("Put is successful")
}

func TestDelete(t *testing.T) {
	apiClientDelete := client.NewApiClient().WithUrl("https://reqres.in/api/users/2")
	ApiDelete(apiClientDelete, []byte(``))
	t.Log("Delete is successful")
}
