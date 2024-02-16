package contract

type ApiClient interface {
	Get(body []byte)
	//Post()
	//Put()
	//Delete()
	CreateClient() error
}
