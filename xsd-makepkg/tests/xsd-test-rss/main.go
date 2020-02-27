package main

import (
	"encoding/xml"

	"github.com/wagner-aos/go-xsd/xsd-makepkg/tests"

	udevgo "github.com/wagner-aos/go-util/dev/go"

	rss "github.com/wagner-aos/go-xsd-pkg/thearchitect.co.uk/schemas/rss-2_0.xsd_go"
)

type RssDoc struct {
	XMLName xml.Name `xml:"rss"`
	rss.TxsdRss
}

func main() {
	var (
		dirBasePath  = udevgo.GopathSrcGithub("wagner-aos", "go-xsd", "xsd-makepkg", "tests", "xsd-test-rss")
		makeEmptyDoc = func() interface{} { return &RssDoc{} }
	)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
}
