package wikihistory

import (
	"errors"
	"fmt"
	"runtime"
	"sort"

	_ "github.com/mattn/go-sqlite3"
)

func GetWikiHistory() ([]HistoryEntry, error) {
	hes := make(map[string]HistoryEntry, 0)

	switch runtime.GOOS {
	case "windows":
		nhes, err := GetFirefoxWindowsHistory()
		if err != nil {
			fmt.Printf("error getting Firefox history: %v\n", err)
		} else {
			for _, nhe := range nhes {
				hes[nhe.Url] = nhe
			}
		}

		nhes, err = GetChromeWindowsHistory()
		if err != nil {
			fmt.Printf("error getting Chrome history: %v\n", err)
		} else {
			for _, nhe := range nhes {
				hes[nhe.Url] = nhe
			}
		}
	default:
		return nil, errors.New("unsupported operating system")
	}

	ret := make([]HistoryEntry, 0)
	for _, he := range hes {
		ret = append(ret, he)
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].LastVisit < ret[j].LastVisit
	})

	return ret, nil
}
