package client

type SlackClient interface {
	SendMessage(channel, message string) error
	GetChannelList() []string
	IsConnected() bool
}

// compile-time check
var _ SlackClient = (*SlackClientImpl)(nil)

type SlackClientImpl struct {
}

func NewSlackClient() *SlackClientImpl {
	return &SlackClientImpl{}
}

func (c *SlackClientImpl) SendMessage(channel, message string) error {
	return nil
}

func (c *SlackClientImpl) GetChannelList() []string {
	return []string{"general", "random"}
}

func (c *SlackClientImpl) IsConnected() bool {
	return true
}
