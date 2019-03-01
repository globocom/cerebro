package modules

// PersistenceClient Interface
type PersistenceClient interface {
	Close()
	AddAttribute(attributeName string, attributeType string) error
	GetAttribute(attributeName string) (*Attribute, error)
	DeleteAttribute(attributeName string) error
}

// ESClient struct
type ESClient struct {
}

// Close connection with persistence interface
func (p *ESClient) Close() {
}

func (p *ESClient) AddAttribute(attributeName string, attributeType string) error {
	return nil
}

func (p *ESClient) GetAttribute(attributeName string) (*Attribute, error) {
	return nil, nil
}

func (p *ESClient) DeleteAttribute(attributeName string) error {
	return nil
}

func (p *ESClient) UpdateAttribute(attributeName string, attributeType string) error {
	return nil
}

// NewESClient build instance from ESClient
func NewESClient(settings Settings) PersistenceClient {
	return &ESClient{}
}
