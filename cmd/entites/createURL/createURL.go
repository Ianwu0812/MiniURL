package createURL

import (
	"errors"
	"strings"
)

func CreateURL(originalURL string) (string, error) {
	if isEmpty(originalURL) {
		err := errors.New("originalURL is empty")
		return "", err
	}
	shortedURL := ""
	return shortedURL, nil
}

func isEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
