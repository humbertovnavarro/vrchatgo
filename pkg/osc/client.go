package osc

import "github.com/hypebeast/go-osc/osc"

type Client struct {
	sender   *osc.Client
	receiver *osc.Client
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
		sender:   osc.NewClient(_opts.Hostname, _opts.SenderPort),
		receiver: osc.NewClient(_opts.Hostname, _opts.ReceiverPort),
	}
}
