package wikihistory

import (
	"io/fs"
	"os"
	"path/filepath"
)

const chromeStatement string = `SELECT DISTINCT
CASE 
    WHEN instr(url, '?') > 0 THEN substr(REPLACE(url, '://en.m.', '://en.'), 1, instr(url, '?') - 1) 
    WHEN instr(url, '#') > 0 THEN substr(REPLACE(url, '://en.m.', '://en.'), 1, instr(url, '#') - 1) 
    ELSE REPLACE(url, '://en.m.', '://en.')
END AS url,
CASE
	WHEN instr(title, ' - Wikipedia') > 0 THEN substr(title, 1, instr(title, ' - Wikipedia') - 1)
	ELSE title
END AS title, last_visit_time
FROM urls
WHERE last_visit_time >= strftime('%s', 'now', '-1 year') * 1000000 + 11644473600000000 -- accounting for Windows epoch
AND url LIKE 'https://%.wikipedia.org/wiki/%'
AND url NOT LIKE 'https://%wikipedia.org/wiki/%:%'
AND title NOT NULL
GROUP BY title`

func GetChromeWindowsHistory() ([]HistoryEntry, error) {
	dir, _ := os.UserCacheDir()
	root := filepath.Join(dir, "Google", "Chrome", "User Data")
	profiles := make([]string, 0)

	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() && path != root {
			profiles = append(profiles, path)
			return filepath.SkipDir
		}
		return nil
	})

	dbs := make([]string, 0)
	for _, profile_dir := range profiles {
		filepath.WalkDir(profile_dir, func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() && path != profile_dir {
				return filepath.SkipDir
			} else if d.Name() == "History" {
				dbs = append(dbs, path)
				return filepath.SkipAll
			}
			return nil
		})
	}

	entries := make([]HistoryEntry, 0)

	for _, path := range dbs {
		new_entries, err := GetHistory(path, chromeStatement)
		if err != nil {
			return nil, err
		}

		entries = append(entries, new_entries...)
	}

	return entries, nil
}
