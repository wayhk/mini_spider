/* currenthost.go - safe host map */
/*
modification history
--------------------
2016/08/28 create by HanKai,safe map
*/
/*
DESCRIPTION
This file contains mycurrentHost
*/
package spider

import (
	"sync"
)

type MyConcurrentHost struct {
	m       map[string]int64
	rwmutex sync.RWMutex
}

func NewConcurrentHost() *MyConcurrentHost {
	myHost := map[string]int64{}
	return &MyConcurrentHost{
		m: myHost,
	}
}

func (cmap *MyConcurrentHost) Put(k string, v int64) {
	cmap.rwmutex.Lock()
	defer cmap.rwmutex.Unlock()
	cmap.m[k] = v

	return
}

func (cmap *MyConcurrentHost) Get(k string) (int64, bool) {
	cmap.rwmutex.RLock()
	defer cmap.rwmutex.RUnlock()
	v, ok := cmap.m[k]

	return v, ok
}

func (cmap *MyConcurrentHost) Delete(k string) {
	cmap.rwmutex.Lock()
	defer cmap.rwmutex.Unlock()
	delete(cmap.m, k)
}
