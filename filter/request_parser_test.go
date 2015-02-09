package filter

import (
	"net/url"
	"testing"
)

func TestRequestParser(t *testing.T) {
	params := url.Values{
		"q[name_eq]": []string{
			"cone",
		},
		"q[last_name_eq]": []string{
			"Gutierrez",
		},
	}

	expected := "name = 'cone' AND last_name = 'Gutierrez'"

	got := parseUrlValues(params)

	if expected != got {
		t.Errorf("Mistach, wanted: %s got: %s", expected, got)
	}
}
