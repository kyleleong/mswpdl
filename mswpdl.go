package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type xmldata struct {
	XMLName xml.Name `xml:"images"`
	URL     string   `xml:"image>url"`
}

var (
	market, outname string
	daysago         int
	debug           bool
)

func init() {
	flag.StringVar(&market, "locale", "en-US", "Different locales will return different images.")
	flag.IntVar(&daysago, "daysago", 0, "Retrieve an old image from the given number of days ago.")
	flag.StringVar(&outname, "outfile", "-", "The name of the file to which the image will be written. '-' writes to stdout.")
	flag.BoolVar(&debug, "debug", false, "Use this switch to print debugging information to stderr.")
	flag.CommandLine.SetOutput(os.Stderr)
}

func main() {
	flag.Parse()

	query := fmt.Sprintf("http://www.bing.com/HPImageArchive.aspx?format=xml&idx=%d&n=1&mkt=%s", daysago, market)
	resp, err := http.Get(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(2)
	}

	var xmldat xmldata
	err = xml.Unmarshal(body, &xmldat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(3)
	}

	resp, err = http.Get("http://bing.com" + xmldat.URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(4)
	}

	var outfile *os.File
	if outname == "-" {
		outfile = os.Stdout
	} else {
		outfile, err = os.Create(outname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err.Error())
			os.Exit(5)
		}
	}

	if debug {
		fmt.Fprintf(os.Stderr, "Not Parsed: %s\n", flag.Args())
		fmt.Fprintf(os.Stderr, "Locale: %s\n", market)
		fmt.Fprintf(os.Stderr, "Daysago: %d\n", daysago)
		fmt.Fprintf(os.Stderr, "Outname: %s\n", outname)
		fmt.Fprintf(os.Stderr, "%s\n", query)
		fmt.Fprintf(os.Stderr, "%s\n", "http://bing.com"+xmldat.URL)
	}

	io.Copy(outfile, resp.Body)
}
