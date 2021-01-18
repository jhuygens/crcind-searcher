package crcind

import (
	"fmt"
	"time"

	"github.com/jgolang/config"
	"github.com/jgolang/consumer/soap"
	"github.com/jgolang/log"
	"github.com/jhuygens/searcher-engine"
)

// Limit returns limit
var Limit = "20"

// Searcher searcher interface implement
type Searcher struct{}

// Search doc ...
func (s Searcher) Search(filter searcher.Filter) ([]searcher.Item, error) {
	if len(filter.Types) == 0 {
		filter.Types = append(filter.Types, "people")
	}
	var items []searcher.Item
	for _, resource := range filter.Types {
		if resource == "people" || resource == "all" {
			for _, name := range filter.Name {
				results, err := searchCRCINDerviceItems(name.Value, "people")
				if err != nil {
					log.Error(err)
				}
				items = append(items, results...)
			}
		}
		if resource == "author" {
			for _, name := range filter.Name {
				results, err := searchCRCINDerviceItems(name.Value, "people")
				if err != nil {
					log.Error(err)
				}
				items = append(items, results...)
			}
		}
		if resource == "artist" {
			for _, name := range filter.Name {
				results, err := searchCRCINDerviceItems(name.Value, "people")
				if err != nil {
					log.Error(err)
				}
				items = append(items, results...)
			}
		}
	}
	return items, nil
}

func searchCRCINDerviceItems(name, resource string) ([]searcher.Item, error) {
	resquest := soap.RequestInfo{
		URL:     config.GetString("integrations.crcind.endpoint"),
		Timeout: time.Duration(config.GetInt("integrations.crcind.timeout")) * time.Second,
		Action:  config.GetString("integrations.crcind.action"),
		Content: GetByName{
			Name: name,
			Tem:  "http://tempuri.org",
		},
	}
	nodeXML, err := soap.ConsumeSOAP12Service(resquest, "//ListByName")
	if err != nil {
		return nil, err
	}
	if soap.GetStringFromXML(nodeXML, "//Name") == "" {
		return nil, fmt.Errorf("CRCIND not results")
	}
	var items []searcher.Item
	items = append(
		items,
		searcher.Item{
			Type:    resource,
			Library: config.GetString("searchers.crcind"),
			Name:    soap.GetStringFromXML(nodeXML, "//Name"),
			Artwork: soap.GetStringFromXML(nodeXML, "//SSN"),
			Info: searcher.Info{
				Title:       soap.GetStringFromXML(nodeXML, "//Name"),
				Description: fmt.Sprintf("Birthdate: %v", soap.GetStringFromXML(nodeXML, "//DOB")),
			},
		},
	)
	return items, nil
}
