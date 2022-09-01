package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type Spell struct {
	Name        string
	School      string
	CastingTime string
	Range       string
	Duration    string
	Components  string
}

func main() {
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	c.Visit("https://en.wikipedia.org/wiki/Radix_tree")
}
