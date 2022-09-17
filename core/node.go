package core

type Node struct {
	Key          string
	MaxThreshold int
}

func NewNode(key string, maxThreshold int) *Node {
	return &Node{Key: key, MaxThreshold: maxThreshold}
}
