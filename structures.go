package crcind

import "encoding/xml"

// GetByName soap request
type GetByName struct {
	XMLName xml.Name `xml:"tem:GetByName"`
	Tem     string   `xml:"xmlns:tem,attr"`
	Name    string   `xml:"tem:name"`
}
