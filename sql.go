package wikihistory

import "database/sql"

func GetHistory(path, query string) ([]HistoryEntry, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entries := make([]HistoryEntry, 0)

	for rows.Next() {
		he := HistoryEntry{}
		err = rows.Scan(&he.Url, &he.Title, &he.LastVisit)
		if err != nil {
			return nil, err
		}
		entries = append(entries, he)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return entries, nil
}
