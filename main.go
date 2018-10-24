package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	docURL := flag.String("u", "https://golang.org/pkg", "The URL of the Godoc you wish to parse")
	outFilename := flag.String("o", "test.csv", "The file to output to")

	flag.Parse()

	// Make an HTTP request to get the page
	resp, err := http.Get(*docURL)
	if err != nil {
		fmt.Println("An error occurred trying to reach", *docURL)
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Parse the page to find function names and descriptions
	// Create a list of cards
	cards := parsePage(resp.Body)

	// Output the list of cards to a CSV file
	file, err := os.Create(*outFilename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cardStrings := cardsToStrings(cards)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.WriteAll(cardStrings)

	file.Close()
}

// Card represents a single Q & A flashcard
type Card struct {
	q string // question side of the card
	a string // answer side
}

func parsePage(r io.Reader) []Card {
	var ret []Card
	example := Card{"Test Question", "Example Answer"}
	ret = append(ret, example)
	return ret
}

func cardsToStrings(cards []Card) [][]string {
	var ret [][]string
	for _, card := range cards {
		ret = append(ret, []string{card.q, card.a})
	}
	return ret
}
