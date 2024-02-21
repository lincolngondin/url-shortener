package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

var (
	driverName string = "postgres"
	host       string = "localhost"
	port       string = "5432"
	user       string = "postgres"
	password   string = "lincoln"
	dbname     string = "url-shortener"
)

func OpenDB() (func(), error) {
	var err error
	db, err = sql.Open(driverName, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return nil, err
	}
	return func() {
		db.Close()
	}, nil
}

func PutURL(url URL) error {
	query := "INSERT INTO urls (original_url, shortened_url, creation_time, last_click, total_clicks) VALUES ($1, $2, $3, $4, $5);"
	_, err := db.Exec(query, url.OriginalURL, url.ShortenderURL, url.CreationTime, url.LastClick, url.TotalClicks)
	return err
}

func GetOriginalURLAndInclementClicks(shortURL string) (string, error) {
    tx, transactionErr := db.Begin()
    if transactionErr != nil {
        return "", transactionErr
    }
	row := tx.QueryRow("SELECT original_url FROM urls WHERE shortened_url = $1", shortURL)
	var url string
	err := row.Scan(&url)
	if err == sql.ErrNoRows {
        tx.Rollback()
		return "", errors.New("URL Not Found!")
	}
    _, updateErr := tx.Exec("UPDATE urls SET total_clicks = total_clicks+1, last_click = $1 WHERE shortened_url = $2", time.Now().Format(time.DateTime), shortURL)
    if updateErr != nil {
        tx.Rollback()
        return "", updateErr
    }
    commitErr := tx.Commit()
    if commitErr != nil {
        return "", commitErr
    }

	return url, nil
}

func GetURLInfo(shortURL string) (*URL, error) {
    row := db.QueryRow("SELECT * FROM urls WHERE shortened_url = $1", shortURL)
    var url *URL = &URL{}
    err := row.Scan(&url.OriginalURL, &url.ShortenderURL, &url.CreationTime, &url.LastClick, &url.TotalClicks)
    if err == sql.ErrNoRows {
        return nil, err
    }
    return url, nil
}
