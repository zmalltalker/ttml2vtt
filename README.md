# TTML2VTT

Call me ignorant, but in my experience, subtitles are something to be
found on the Internet in the shape of `SRT` files. As it turns out,
there are several file formats, none of them pretty.

This tool aims to convert from a really ugly format -
[TTML](http://www.w3.org/TR/ttaf1-dfxp/) into a slighly less ugly
format, [VTT](http://dev.w3.org/html5/webvtt/).


## Installing it

First of all, install golang and git. Once that's done, simply

```shell
go get github.com/zmalltalker/ttml2vtt
``´

Assuming you have `$GOPATH/bin` in your `$PATH` you should be able to
run the `ttml2vtt` binary now.

## Running it

The `ttml2vtt` binary reads an XML file (TTML) from `STDIN` and emits
WTT (or something similar to it) on `STDOUT`.


If you have a TTML file handy, feel free to pass that to
`ttml2vtt`. If you don't have one (seriously, who does?), feel free to
use the sample file included in this repo:

```
curl -s https://raw.githubusercontent.com/zmalltalker/ttml2vtt/master/data/sample.xml| ttml2vtt
```


## Caveats

This is my first Go project, and my big mouth made me do it. I'm going
to be so embarrassed by this code not long from now.

## TODO

* simplify XMLName structure [DONE]
* parse the text properly
* read from stdin [DONE]
* remove whitespace [DONE]
* add styling [DONE, for now]
* Create module, use in main
* Make installable, put on GH [DONE]
