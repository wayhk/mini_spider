/* main.go - the main structure of config  */
/*
modification history
--------------------
2016/06/24, by Han Kai, created
2016/08/04,by Han Kai,modify,sealing exit
2016/08/08,by Han Kai,modify,add sleep when exit
*/
/*
DESCRIPTION
program entry
*/

package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

import (
)

import (
	"config"
	"parseurl"
	"spider"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//parse args
	arg := ParseArgs()

	if arg.Help {
		arg.PrintHelp()
		os.Exit(0)
	}

	//print version
	if arg.Version {
		arg.PrintVersion()
		os.Exit(0)
	}

	//init log
	err := log.Init("mini_spider", "INFO", arg.LogDir, true, "midnight", 5)
	if err != nil {
		fmt.Printf("log init failed:%#v\n", err)
		os.Exit(1)
	}
	log.Logger.Info("init log success")
	time.Sleep(100 * time.Millisecond)

	//load config
	spiderConfig, err := config.LoadConfig(arg.ConfigDir)
	if err != nil {
		log.Logger.Error("load spider config failed:%#v", err)
		Exit(-1)
	}
	log.Logger.Info("get config success")

	//deal with unexpected quit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-c
		log.Logger.Info("quit,begin to clean resource")
		Exit(0)
	}()

	//parse urlListFile
	urls, err := parseurl.ParseUrl(spiderConfig.UrlListFile)
	if err != nil {
		log.Logger.Error("parse urlListFile failed:%#v", err)
		Exit(-1)
	}

	//execute crawl task
	log.Logger.Info("start crawling...")
	sp := spider.NewSpider(urls, spiderConfig)
	sp.Start()

	//wait finish,exit
	<-sp.Done
	log.Logger.Info("spider finished,bye")
	Exit(0)

}

//sealing exit
func Exit(code int) {
	log.Logger.Close()
	time.Sleep(100 * time.Millisecond)
	os.Exit(code)
}
