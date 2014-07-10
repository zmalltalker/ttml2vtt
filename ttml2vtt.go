package main

import(
	"encoding/xml"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
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


func main(){
	bytes, ioerr := ioutil.ReadAll(os.Stdin)
	replacer := strings.NewReplacer("<br />","\n",`<span style="italic">`,"<i>", "</span>","</i>")
	if ioerr != nil {
		panic("ZOMG")
	}

	v := Feed{}
	err := xml.Unmarshal([]byte(bytes), &v)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Printf("WEBVTT\n##\nLanguage: %v\n\n", v.Lang)
	for _, screen := range v.Screens{
		fmt.Printf("%v ---> %v\n", screen.Begin, screen.Duration)
		fmt.Printf("%v\n\n", strings.TrimSpace(replacer.Replace(screen.Text)))
	}

}