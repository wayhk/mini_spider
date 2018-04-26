/* execute.go - task dispatcher and thread control file  */
/*
modification history
--------------------
2016/06/24, by Han Kai, created
2016/08/05,by Han Kai,modify,sealing cocurrent task map
2016/08/19,by Han Kai,modify,sealing cocurrent host map
*/
/*
DESCRIPTION
This file contains execute control func
*/

package spider

import (
	"sync"
)

import (
	"config"
)

var MAX_TASK_COUNT = 1000

//task url struct
type UrlTask struct {
	Url   string
	Depth int64
}

//spider struct
type Spider struct {
	Urls      []string
	Conf      *config.SpiderConfig
	Tasks     chan UrlTask
	Done      chan bool
	TaskCache *MyConcurrentTask
	Wg        sync.WaitGroup
	HostCache *MyConcurrentHost
}

//create new spider
func NewSpider(urls []string, conf *config.SpiderConfig) *Spider {
	return &Spider{
		Urls:      urls,
		Conf:      conf,
		Tasks:     make(chan UrlTask, MAX_TASK_COUNT),
		Done:      make(chan bool, 1),
		TaskCache: NewConcurrentTask(),
		HostCache: NewConcurrentHost(),
	}
}

//spider controller,open exec thread,add basic task
func (sp *Spider) Start() {

	//open exec thread
	for i := int64(0); i < sp.Conf.ThreadCount; i++ {
		go sp.Crawling()
	}

	//add basic task
	for _, url := range sp.Urls {
		task := UrlTask{}
		task.Url = url
		task.Depth = 0

		//add task to task chan
		sp.AddTask(task)

		//backup url
		sp.TaskCache.Put(url, true)
	}
	//wait for done
	sp.Wg.Wait()
	sp.Done <- true
}

//add task to spider tasks chan and backup url
func (sp *Spider) AddTask(task UrlTask) {
	sp.Wg.Add(1)
	go func(t UrlTask) {
		sp.Tasks <- t
		sp.TaskCache.Put(t.Url, true)
	}(task)
}
