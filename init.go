package crcind

import (
	"github.com/jgolang/config"
	"github.com/jgolang/log"
	"github.com/jhuygens/searcher-engine"
)

var crcindSearcher = Searcher{}

func init() {
	name := config.GetString("searchers.crcind")
	err := searcher.RegisterSearcher(name, crcindSearcher)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Infof("Searcher %v has been register", name)
}
