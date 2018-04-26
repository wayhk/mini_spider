/* crawling.go - the main file of execute crawling task */
/*
modification history
--------------------
2016/06/24, by Han Kai, created
2016/07/25,by Han Kai,modify,split Crawl func
2016/08/02,by Han Kai,modify,async push taskurl
2016/08/08,by Han Kai,modify,modify rwLock to rLock,when get
*/
/*
DESCRIPTION
This file contains crawling func
*/
package spider

import (
	"sync"
)

type MyConcurrentTask struct {
	m       map[string]bool
	rwmutex sync.RWMutex
}

func NewConcurrentTask() *MyConcurrentTask {
	myTask := map[string]bool{}
	return &MyConcurrentTask{
		m: myTask,
	}
}

func (cmap *MyConcurrentTask) Put(k string, v bool) {
	cmap.rwmutex.Lock()
	defer cmap.rwmutex.Unlock()
	cmap.m[k] = v

	return
}

func (cmap *MyConcurrentTask) Get(k string) (bool, bool) {
	cmap.rwmutex.RLock()
	defer cmap.rwmutex.RUnlock()
	v, ok := cmap.m[k]

	return v, ok
}

func (cmap *MyConcurrentTask) Delete(k string) {
	cmap.rwmutex.Lock()
	defer cmap.rwmutex.Unlock()
	delete(cmap.m, k)
}
