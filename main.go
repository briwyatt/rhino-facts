package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"strconv"
	"github.com/gocolly/colly"
)

type Fact struct {
	ID	int 	`json:"id`
	Description string `json:"description"`
}

func main() {
	allFacts := make([]Fact, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)

	// this code traverses the DOM
	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		factId, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			// if we get an error print error message
			log.Println("Could not get id")
		}

		// fact description for each factId
		factDesc := element.Text

		// create a new fact struct for every list item we iterate over
		fact := Fact{
			ID: factId,
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("https://www.factretriever.com/rhino-facts")


	// enc := json.NewEncoder(os.Stdout)

	// enc.SetIndent("", " ")
	// enc.Encode(allFacts)
	writeJSON(allFacts)
}


func writeJSON(data []Fact) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create JSON file")
		return
	}


	_ = ioutil.WriteFile("rhinofacts.json", file, 0644)
}
