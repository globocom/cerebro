package modules

// PersistenceClient Interface
type PersistenceClient interface {
	Close()
}

// ESClient struct
type ESClient struct {
}

// Close connection with persistence interface
func (p *ESClient) Close() {
}

// NewESClient build instance from ESClient
func NewESClient(settings Settings) *ESClient {
	return &ESClient{}
}
