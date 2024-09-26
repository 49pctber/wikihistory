package main

import (
	"fmt"

	"github.com/49pctber/wikihistory"
)

func main() {
	history, _ := wikihistory.GetWikiHistory()
	for _, entry := range history {
		fmt.Println(entry)
	}

	fmt.Printf("Found %d entries from the last year.\n", len(history))
}
