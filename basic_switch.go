package switchx

import (
	"github.com/kangyujian/switchx/core"
	"sync"
)

type BasicSwitch struct {
	nodeMap sync.Map
}

func (b *BasicSwitch) Register(key string, maxThreshold int64) {
	b.nodeMap.Store(key, core.NewNode(key, maxThreshold))
}

func (b *BasicSwitch) IsOn(key string) bool {
	panic("implement me")
}
