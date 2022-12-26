package main

// import xml, fmt, io, http, os
import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}

// we are making this a struct because its a user defined data type
// This type has multiple fields of other types
type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

// We seem to use a slice alot because arrays are a useful data type but rigid
type News struct {
	Headline     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

func main() {
	var r RSS

	data := readGoogleTrends()

	err := xml.Unmarshal(data, &r)

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Println("Below are all the Google Search Trends for Today in Canada")
	fmt.Println("----------------------------------------------------------")

	for i := range r.Channel.ItemList {
		rank := (i + 1)
		fmt.Println("#", rank)
		fmt.Println("Search Term: ", r.Channel.ItemList[i].Title)
		fmt.Println("Search Traffic: ", r.Channel.ItemList[i].Traffic)
		fmt.Println("Search Link: ", r.Channel.ItemList[i].Link)
		fmt.Println("News Headlines:")
		fmt.Println("----------------------------------------------------------")

		for j := range r.Channel.ItemList[i].NewsItems {
			fmt.Println("Headline: ", r.Channel.ItemList[i].NewsItems[j].Headline)
			fmt.Println("Headline Link: ", r.Channel.ItemList[i].NewsItems[j].HeadlineLink)
		}
		fmt.Println("----------------------------------------------------------")
		fmt.Printf("\n")
	}

}

func readGoogleTrends() []byte {
	resp := getGoogleTrends()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return data
}

func getGoogleTrends() *http.Response {
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=CA")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return resp
}
