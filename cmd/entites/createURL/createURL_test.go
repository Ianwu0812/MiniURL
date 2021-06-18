package createURL

import (
	"testing"
)

func TestCreateURL(t *testing.T) {
	shortedURL, err := CreateURL("")
	if err == nil {
		t.Fatalf(`CreateURL("") = %q, %v`, shortedURL, err)
	}
}
