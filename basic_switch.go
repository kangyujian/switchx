package switchx

import (
	"log"
	"math/rand"
	"sync"

	"github.com/kangyujian/switchx/core"
)

type BasicSwitch struct {
	nodeMap sync.Map
}

func (b *BasicSwitch) Register(key string, maxThreshold int) {
	b.nodeMap.Store(key, core.NewNode(key, maxThreshold))
}

func (b *BasicSwitch) IsOn(key string, currentCursor int) bool {
	val, ok := b.nodeMap.Load(key)
	if !ok {
		log.Printf("[BasicSwitch] basic switch load key %v , not exist!\n", key)
	}

	node, ok := val.(*core.Node)
	if !ok {
		log.Printf("[BasicSwitch] trans register node failed !\n")
	}

	randCursor := rand.Intn(currentCursor) % node.MaxThreshold
	if randCursor <= currentCursor {
		return true
	}
	return false
}
