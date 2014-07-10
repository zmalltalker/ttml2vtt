package main

import(
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"github.com/zmalltalker/ttml2vtt/parser"
)



func main(){
	bytes, ioerr := ioutil.ReadAll(os.Stdin)
	replacer := strings.NewReplacer("<br />","\n",`<span style="italic">`,"<i>", "</span>","</i>")



	if ioerr != nil {
		panic("ZOMG")
	}

	screens, header := parser.ScreensFromXml(bytes)

	fmt.Printf(header)
	for _, screen := range screens{
		fmt.Printf("%v ---> %v\n", screen.Begin, screen.Duration)
		fmt.Printf("%v\n\n", strings.TrimSpace(replacer.Replace(screen.Text)))
	}

}
