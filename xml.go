package main

import (
	"encoding/xml"
	"io/ioutil"
	"regexp"
	"strings"
)

type TstItem struct {
	Case      string `xml:"h"`
	Operation string `xml:"ope"`
	Status    string `xml:"status"`
}

var rxUnIndent = regexp.MustCompile(`(?m)^\s+`)
var rxLineCont = regexp.MustCompile("  \n")

func shrink(s string) string {
	s = rxUnIndent.ReplaceAllString(s, "")
	s = rxLineCont.ReplaceAllString(s, "")
	s = strings.TrimSpace(s)
	return s
}

func readXmlFile(fname string) ([]*TstItem, error) {
	var items struct {
		Items []*TstItem `xml:"item"`
	}

	bin, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(bin, &items)
	if err != nil {
		return nil, err
	}
	for _, item1 := range items.Items {
		item1.Case = shrink(item1.Case)
		item1.Operation = shrink(item1.Operation)
		item1.Status = shrink(item1.Status)
	}
	return items.Items, nil
}

func readXmlFiles(fnames []string) ([]*TstItem, error) {
	var items []*TstItem
	for _, fname := range fnames {
		_items, err := readXmlFile(fname)
		if err != nil {
			return nil, err
		}
		items = append(items, _items...)
	}
	return items, nil
}
