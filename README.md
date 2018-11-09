# godoc2anki
Turn [godoc](https://godoc.org/) pages into [Anki](https://apps.ankiweb.net/) flashcards, so you can learn and remember useful Go libraries.

## Usage
```bash
godoc2anki -u net -o my-cards.csv
```

This will create a CSV file in the `Question, Answer` format that can be imported into Anki.

```csv
net/http,Package http provides HTTP client and server implementations.
net/cgi,Package cgi implements CGI (Common Gateway Interface) as specified in RFC 3875.
```

## Flags
| Flag | Purpose | Default|
| ----- | ----- | ----- |
|`-o` | output filename | `<packagename>.csv` |
|`-u` | package name to scrape | `net` |
