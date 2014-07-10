package main

import(
	"encoding/xml"
	"fmt"
)

// tt > body > div -> p
type Feed struct {
	XMLName xml.Name `xml:"tt"`
	Body Body `xml:"body"`
}

type Body struct {
	XMLName xml.Name `xml:"body"`
	Movie Movie `xml:"div"`
}

type Screen struct {
	Text string `xml:",chardata"`
	Begin string `xml:"begin,attr"`
	Duration string `xml:"dur,attr"`
}

type Movie struct {
	XMLName xml.Name `xml:"div"`
	Screens []Screen `xml:"p"`
	Title string
}



func main(){
	data := `
<tt xmlns="http://www.w3.org/ns/ttml" xmlns:tts="http://www.w3.org/2006/04/ttaf1#styling" lang="no">
  <head>
    <styling>
      <style id="italic" tts:fontStyle="italic" />
      <style id="left" tts:textAlign="left" />
      <style id="center" tts:textAlign="center" />
      <style id="right" tts:textAlign="right" />
    </styling>
  </head>
  <body>
    <div>
      <p begin="00:00:00.000" dur="00:00:00.000">Copyright (C) NRK</p>
      <p begin="00:00:08.200" dur="00:00:06.660" style="left">-Jeg er ikke middagen din!<br />-Jeg vet ikke hvem jeg er.</p>
      <p begin="00:00:15.160" dur="00:00:06.940" style="left">
	<span style="italic"> -Betjent Jason Stackhouse. </span>
	<br />-Hvor er du? Du må hjelpe meg.</p>
      <p begin="00:00:22.400" dur="00:00:04.220" style="left">Takk, åndefar,<br />for at du skjenker oss nytt liv.</p>
    </div>
  </body>
</tt>`
//	v := Movie{Title: "The great Gatsby"}
	v := Feed{}
	err := xml.Unmarshal([]byte(data), &v)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	for _, screen := range v.Body.Movie.Screens{
		fmt.Printf("%v ---> %v\n", screen.Begin, screen.Duration)
		fmt.Printf("%v\n", screen.Text)
	}

}
