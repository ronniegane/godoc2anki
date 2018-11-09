package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var pkgName *string

func main() {
	pkgName = flag.String("u", "net", "The package you wish to parse")
	outFilename := flag.String("o", "test.csv", "The file to output to")

	flag.Parse()

	docURL := "https://golang.org/pkg/" + *pkgName

	// Make an HTTP request to get the page
	resp, err := http.Get(docURL)
	if err != nil {
		fmt.Println("An error occurred trying to reach", docURL)
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
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	/*
		Structure of a godoc page
			<div class="pkg-dir">
			<table>
				<tr>
					<th class="pkg-name">Name</th>
					<th class="pkg-synopsis">Synopsis</th>
				</tr>

				<tr>
					<td colspan="2"><a href="..">..</a></td>
				</tr>

				<tr>

					<td class="pkg-name" style="padding-left: 0px;">
						<a href="bzip2/">bzip2</a>
					</td>

					<td class="pkg-synopsis">
						Package bzip2 implements bzip2 decompression.
					</td>
				</tr>
	*/
	nodes := findNodes(doc)

	var ret []Card
	for _, node := range nodes {
		ret = append(ret, nodeToCard(node))
	}
	return ret
}

func findNodes(n *html.Node) []*html.Node {
	var nodeList []*html.Node
	if n.Type == html.ElementNode && n.Data == "td" && hasPkgNameClass(n) {
		// Found the package name node
		return []*html.Node{n}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodeList = append(nodeList, findNodes(c)...)
	}
	return nodeList
}

func nodeToCard(n *html.Node) Card {
	// Issue here: the <td> node has 3 children, the first and last ones are TextNodes with simply data "\n\t\t\t\t\t"
	// so to reach the <a> node below it is actually FirstChild.NextSibling
	// likewise the NextSibling of the pkg-name <td> node is just a whitespace text node.
	// n.NextSibling.NextSibling is the actual next <td> with the synopsis.
	// but this synopsis <td> does contain the synopsis in the FirstChild, just it's padded with \n\t\t\t\t\t at the front.
	// Not very consistent.
	// TODO: fix this ugly, completely not null-safe access
	synopsis := strings.TrimSpace(n.NextSibling.NextSibling.FirstChild.Data)
	name := n.FirstChild.NextSibling.FirstChild.Data
	return Card{
		q: *pkgName + "/" + name,
		a: synopsis,
	}
}

func hasPkgNameClass(n *html.Node) bool {
	for _, attr := range n.Attr {
		if attr.Key == "class" && attr.Val == "pkg-name" {
			return true
		}
	}
	return false
}

func cardsToStrings(cards []Card) [][]string {
	var ret [][]string
	for _, card := range cards {
		ret = append(ret, []string{card.q, card.a})
	}
	return ret
}
