package types

type IMetadata interface {
	Decode(data []byte) (err error)
	Encode() (data []byte, err error)
	Sign() (data []byte, err error)
}
