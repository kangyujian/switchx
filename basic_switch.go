package switchx

type BasicSwitch struct {
}

func (b *BasicSwitch) Register(key string, maxThreshold int64) {
	panic("implement me")
}

func (b *BasicSwitch) IsOn(key string) bool {
	panic("implement me")
}
