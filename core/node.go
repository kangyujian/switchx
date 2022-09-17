package core

type Node struct {
	Key          string
	MaxThreshold int64
}

func NewNode(key string, maxThreshold int64) *Node {
	return &Node{Key: key, MaxThreshold: maxThreshold}
}
