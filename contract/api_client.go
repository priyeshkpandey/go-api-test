package contract

type ApiClient interface {
	Get(body []byte)
	Post(body []byte)
	Put(body []byte)
	Delete(body []byte)
	CreateClient() error
}
