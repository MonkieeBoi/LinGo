package db

import (
	"bufio"
	"database/sql"
	"errors"
	"io"
	"path"
	"sort"
	"strings"

	"github.com/adrg/xdg"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func getDBPath() (string, error) {
	dbFileName := "db.sqlite"
	path, err := xdg.DataFile(path.Join("Goana", dbFileName))
	if err != nil {
		return "", err
	}
	return path, nil
}

func InitDB() error {
	dbPath, err := getDBPath()
	if err != nil {
		return err
	}
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS words (
        alpha TEXT NOT NULL,
        word TEXT NOT NULL,
		PRIMARY KEY (alpha, word)
        );`,
	)
	if err != nil {
		return err
	}
	return nil
}

func sortString(s string) string {
	sSlice := strings.Split(s, "")
	sort.Strings(sSlice)
	return strings.Join(sSlice, "")
}

func isAlpha(s string) bool {
	for _, c := range s {
		if c < 'a' || c > 'z' {
			return false
		}
	}
	return true
}

func HasWords() (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM words").Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func AddWord(word string) error {
	if !isAlpha(word) || len(word) < 2 || len(word) > 15 {
		return errors.New("Invalid word, must only contain alphabetical characters " +
			"and be between 2 and 15 characters in length.")
	}
	word = strings.TrimSpace(strings.ToLower(word))
	alpha := sortString(word)
	_, err := db.Exec("INSERT INTO words (alpha, word) VALUES (?, ?)", alpha, word)
	if err != nil {
		return err
	}
	return nil
}

func AddWords(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		AddWord(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func DelWord(word string) error {
	_, err := db.Exec("DELETE FROM words WHERE word = ?;", word)
	if err != nil {
		return err
	}
	return nil
}

func GetWords(alpha string) ([]string, error) {
	alpha = strings.ToLower(sortString(alpha))
	var words []string
	rows, err := db.Query("SELECT word FROM words WHERE alpha = ?", alpha)
	if err != nil {
		return nil, nil
	}
	defer rows.Close()
	for rows.Next() {
		var word string
		if err := rows.Scan(&word); err != nil {
			return nil, err
		}
		words = append(words, word)
	}
	if err := rows.Err(); err != nil {
		return nil, nil
	}
	return words, nil
}

func GenAlpha() (string, error) {
	var alpha string
	row := db.QueryRow("SELECT DISTINCT alpha FROM words ORDER BY RANDOM() LIMIT 1")
	if err := row.Scan(&alpha); err != nil {
		return "", err
	}
	return alpha, nil
}
