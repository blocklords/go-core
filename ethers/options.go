package ethers

import "strings"

const NilAddress = "0x0000000000000000000000000000000000000000"
const DeadAddress = "0x000000000000000000000000000000000000dEaD"

func BurnAddress() []string {
	return []string{strings.ToLower(NilAddress), strings.ToLower(DeadAddress)}
}

type (
	RequestOptions struct {
		host  string
		key   string
		proxy string
	}

	RequestOptsFunc func(options *RequestOptions)

	ContractOptions struct {
		contract string
		topic    string
		block    BlockInterface
	}

	ContractOptsFunc func(options *ContractOptions)
	Options          struct {
		*RequestOptions
		*ContractOptions
	}
)

func WithRequestHost(host string) RequestOptsFunc {
	return func(options *RequestOptions) {
		options.host = host
	}
}

func WithRequestKey(key string) RequestOptsFunc {
	return func(options *RequestOptions) {
		options.key = key
	}
}
func WithRequestProxy(proxy string) RequestOptsFunc {
	return func(options *RequestOptions) {
		options.proxy = proxy
	}
}

func NewRequestOptions(optFunc ...RequestOptsFunc) *RequestOptions {
	opts := &RequestOptions{}
	for _, of := range optFunc {
		of(opts)
	}
	return opts
}
func (opts *RequestOptions) Host() string {
	return opts.host
}

func (opts *RequestOptions) Key() string {
	return opts.key
}

func (opts *RequestOptions) Proxy() string {
	return opts.proxy
}

func WithContractAddress(address string) ContractOptsFunc {
	return func(options *ContractOptions) {
		options.contract = address
	}
}

func WithContractTopic(topic string) ContractOptsFunc {
	return func(options *ContractOptions) {
		options.topic = topic
	}
}

func WithIBlock(face BlockInterface) ContractOptsFunc {
	return func(options *ContractOptions) {
		options.block = face
	}
}

func NewContractOptions(optFunc ...ContractOptsFunc) *ContractOptions {
	opts := &ContractOptions{}
	for _, of := range optFunc {
		of(opts)
	}
	return opts
}
func (opts *ContractOptions) Contract() string {
	return opts.contract
}

func (opts *ContractOptions) Topic() string {
	return opts.topic
}

func (opts *ContractOptions) IBlock() BlockInterface {
	return opts.block
}
