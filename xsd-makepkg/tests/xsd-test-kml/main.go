package main

import (
	"encoding/xml"

	"github.com/wagner-aos/go-xsd/xsd-makepkg/tests"

	udevgo "github.com/wagner-aos/go-util/dev/go"

	kml "github.com/wagner-aos/go-xsd-pkg/schemas.opengis.net/kml/2.2.0/ogckml22.xsd_go"
)

type KmlDoc struct {
	XMLName xml.Name `xml:"http://www.opengis.net/kml/2.2 kml"`
	kml.TKmlType
}

func main() {
	var (
		dirBasePath  = udevgo.GopathSrcGithub("wagner-aos", "go-xsd", "xsd-makepkg", "tests", "xsd-test-kml")
		makeEmptyDoc = func() interface{} { return &KmlDoc{} }
	)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
}
