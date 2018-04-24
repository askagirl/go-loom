package plugin

import (
	"time"

	loom "github.com/loomnetwork/loom-plugin"
	"github.com/loomnetwork/loom-plugin/types"
)

type (
	Request     = types.Request
	Response    = types.Response
	Meta        = types.ContractMeta
	BlockHeader = types.BlockHeader
)

var (
	EncodingType_JSON      = types.EncodingType_JSON
	EncodingType_PROTOBUF3 = types.EncodingType_PROTOBUF3
)

type Message struct {
	Sender loom.Address
}

type StaticAPI interface {
	StaticCall(addr loom.Address, input []byte) ([]byte, error)
	Emit(event []byte)
}

type VolatileAPI interface {
	Call(addr loom.Address, input []byte) ([]byte, error)
}

type API interface {
	StaticAPI
	VolatileAPI
}

type StaticContext interface {
	StaticAPI
	Get(key []byte) []byte
	Has(key []byte) bool
	Block() BlockHeader
	Now() time.Time
	Message() Message
	ContractAddress() loom.Address
}

type Context interface {
	StaticContext
	VolatileAPI
	Set(key, value []byte)
	Delete(key []byte)
}

type Contract interface {
	Meta() (Meta, error)
	Init(ctx Context, req *Request) error
	Call(ctx Context, req *Request) (*Response, error)
	StaticCall(ctx StaticContext, req *Request) (*Response, error)
}