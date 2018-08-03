package libdon

type Don struct {
}

func NewDon() *Don {
	return &Don{}
}

func (sd *Don) Get() string {
	return donAscii
}
