package client

type Client struct {
	AccessKey string
	SecretKey string
}

func NewClient(accessKey, secretKey string) *Client {
	return &Client{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}

func (c *Client) GetNameByGet(name string) (data interface{}, err error) {
	return "名字是" + name, nil
}
