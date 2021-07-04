package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	//fmt.Println("hey")
	fName := "date.csv"
	file, err := os.Create(fName)
	if err != nil {
		fmt.Println("faili ei loodud")
	}
	//_ = file

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("www.postimees.ee", "postimees.ee"),
	)

	c.OnHTML(".list-article__text", func(e *colly.HTMLElement) {
		//fmt.Println(e.)

		writer.Write([]string{

			e.Text,
			//e.ChildAttr("#text"),
		})
	})

	c.Visit("https://www.postimees.ee")

	data, err := os.Open("date.csv")
	if err != nil {
		fmt.Println("ei avatud date.csv")
	}
	r := csv.NewReader(data)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("vahepeal ei toota")
		os.Exit(1)
	}

	for _, v := range lines {
		fmt.Println(v[0])
	}
}
