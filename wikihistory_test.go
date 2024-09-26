package wikihistory

import (
	"fmt"
	"testing"
)

func TestWikiHistory(t *testing.T) {
	history, _ := GetWikiHistory()
	for _, entry := range history {
		fmt.Println(entry)
	}
}
