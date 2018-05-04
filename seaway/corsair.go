package seaway

type CorsairInfo struct {
	Msg string
}

type Corsair interface {
	GetName() string
	Wigwag() (spotted bool, err error)
	Msg() (corsairInfo *CorsairInfo, err error)
}
