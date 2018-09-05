package models

import (
	"encoding/xml"
)

type UriResponse struct {
	StatusCode int        `xml:"status"`
	StatusMsg  string     `xml:"statusmsg"`
	URI        string     `xml:"uri"`
	Categories []Category `xml:"categories>cat"`
	Bcri       int        `xml:"bcri"`
	A1cat      int        `xml:"a1cat"`
}

type CategoryResponse struct {
	StatusCode int        `xml:"status"`
	StatusMsg  string     `xml:"statusmsg"`
	Categories []Category `xml:"categories>cat"`
}

type Category struct {
	CatID      int    `xml:"catid"`
	CatName    string `xml:"catname"`
	CatGroup   string `xml:"catgroup"`
	Confidence int    `xml:"conf"`
}

type BCapURI struct {
	Bcap     xml.Name    `xml:"bcap"`
	Response UriResponse `xml:"response"`
}

type BCapCat struct {
	Bcap     xml.Name         `xml:"bcap"`
	Response CategoryResponse `xml:"response"`
}

func DecodeCATResponse(body []byte) (*CategoryResponse, error) {
	var b BCapCat
	if err := xml.Unmarshal(body, &b); err != nil {
		logger.WithError(err).Info("Could not decode cat xml response")
		return nil, err
	}
	//fmt.Println(b.Response)
	return &b.Response, nil
}

func DecodeURIResponse(body []byte) (*UriResponse, error) {
	var b BCapURI
	if err := xml.Unmarshal(body, &b); err != nil {
		logger.WithError(err).Info("Could not decode uri xml response")
		return nil, err
	}
	//fmt.Println(b.Response)
	return &b.Response, nil
}
