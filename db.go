package main

import "fmt"

/*
	ID: int
	URL: string
	SHORTURL: string
	CREATEDAT: timestamp
*/

func WriteURLToDB(url string, shorturl string) error {
	insertStatement := `INSERT INTO public."Urls" (url, shorturl, createdat) VALUES ($1, $2, NOW())`
	_, err := DB.Query(insertStatement, url, shorturl)
	return err
}

func WriteIfNotExists(url string) (string, error) {
	exists, err := CheckURLExists(url)
	if err != nil {
		return "", err
	}
	if exists {
		shorturl, err := GetShortURLFromURL(url)
		if err != nil {
			return "", err
		}
		return shorturl, nil
	}
	shorturl, err := GenerateUUID()
	if err != nil {
		return "", err
	}
	err = WriteURLToDB(url, shorturl)
	if err != nil {
		return "", err
	}
	return shorturl, nil
}

func GetURLFromShortURL(shorturl string) (string, error) {
	var url string
	err := DB.QueryRow(`SELECT url FROM public."Urls" WHERE shorturl=$1`, shorturl).Scan(&url)
	return url, err
}

func ReadAllURLsFromDB() (string, error) {

	rows, err := DB.Query(`SELECT * FROM public."Urls"`)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var urls string
	for rows.Next() {
		var id int
		var url string
		var shorturl string
		var createdat string
		err = rows.Scan(&id, &url, &shorturl, &createdat)
		if err != nil {
			return "", err
		}
		urls += fmt.Sprintf("%d %s %s %s\n", id, url, shorturl, createdat)
	}
	return urls, nil

}

func CheckShortURLExists(shorturl string) (bool, error) {
	var exists bool
	err := DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM public."Urls" WHERE shorturl=$1)`, shorturl).Scan(&exists)
	return exists, err
}

func CheckURLExists(url string) (bool, error) {
	var exists bool
	err := DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM public."Urls" WHERE url=$1)`, url).Scan(&exists)
	return exists, err
}

func GetShortURLFromURL(url string) (string, error) {
	var shorturl string
	err := DB.QueryRow(`SELECT shorturl FROM public."Urls" WHERE url=$1`, url).Scan(&shorturl)
	return shorturl, err
}

func RemoveExpiredURLs() error {
	_, err := DB.Exec(`DELETE FROM public."Urls" WHERE createdat < NOW() - INTERVAL '1 hour'`)
	return err
}
