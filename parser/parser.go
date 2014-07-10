package parser
import("encoding/xml")

// tt > body > div -> p
type Feed struct {
	XMLName xml.Name `xml:"tt"`
	Screens []*Screen `xml:"body>div>p"`
	Lang string `xml:"lang,attr"`
}

type Screen struct {
	Text string `xml:",innerxml"`
	Begin string `xml:"begin,attr"`
	Duration string `xml:"dur,attr"`
}
