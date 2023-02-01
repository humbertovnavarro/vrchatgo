package vrcosc

import "github.com/hypebeast/go-osc/osc"

type Client struct {
	Sender   *osc.Client
	Receiver *osc.Client
}

type ClientOptions struct {
	Hostname     string
	SenderPort   int
	ReceiverPort int
}

var defaultClientOptions = &ClientOptions{
	Hostname:     "localhost",
	SenderPort:   9001,
	ReceiverPort: 9000,
}

func New(opts *ClientOptions) *Client {
	_opts := opts
	if _opts == nil {
		_opts = defaultClientOptions
	}
	return &Client{
		Sender:   osc.NewClient(_opts.Hostname, _opts.SenderPort),
		Receiver: osc.NewClient(_opts.Hostname, _opts.ReceiverPort),
	}
}
