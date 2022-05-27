package main

import (
	"encoding/json"
	"os"

	"github.com/gocolly/colly/v2"
)

type Noticia struct {
	Headline string
}

func main() {
	c := colly.NewCollector()

	news := make([]Noticia, 0)

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		if e.Attr("itemprop") == "headline" {
			news = append(news, Noticia{Headline: e.Text})
		}
	})

	c.Visit("https://www.farodevigo.es/")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(news)
}
