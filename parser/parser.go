package parser
import(
	"encoding/xml"
	"fmt"
)

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

func ScreensFromXml(bytes []byte) ([]*Screen, string){

	v := Feed{}
	err := xml.Unmarshal([]byte(bytes), &v)

	if err != nil {
		panic("XML parsing failed ZOMG")
	}

	// TODO: Replace with replacer

	return v.Screens, fmt.Sprintf("WEBVTT\n##\nLanguage: %v\n\n", v.Lang)
}
