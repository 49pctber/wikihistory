package wikihistory

import "fmt"

type HistoryEntry struct {
	Url       string
	Title     string
	LastVisit int64
}

func (he HistoryEntry) String() string {
	return fmt.Sprintf("%s (%s)", he.Title, he.Url)
}
