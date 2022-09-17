package core

type Switch interface {
	Register(key string, maxThreshold int)
	IsOn(key string, currentCursor int) bool
}
