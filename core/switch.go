package core

type Switch interface {
	Register(key string, maxThreshold int64)
	IsOn(key string) bool
}
