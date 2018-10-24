# godoc2anki
Turn [godoc](https://godoc.org/) pages into [Anki](https://apps.ankiweb.net/) flashcards, so you can learn and remember useful Go libraries.

## Usage
```bash
godoc2anki -u https://golang.org/pkg/compress/ -o compress-cards.csv
```

This will create a CSV file in the `Question, Answer` format that can be imported into Anki.

## Flags
| Flag | Purpose | Default|
| ----- | ----- | ----- |
|`-o` | output filename | `<packagename>.csv` |
|`-u` | URL of godoc to scrape | `https://golang.org/pkg` |
