package filter

import (
	"net/url"
	"strings"
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
		"q[cp_eq]": []string{
			"10289",
		},
	}

	got := parseUrlValues(params)

	substring := "last_name = 'Gutierrez'"

	if !strings.Contains(got, substring) {
		t.Errorf("Response: %s, doesn not contain substring: %s", got, substring)
	}

	substring = "name = 'cone'"

	if !strings.Contains(got, substring) {
		t.Errorf("Response: %s, doesn not contain substring: %s", got, substring)
	}

	substring = "cp = 10289"

	if !strings.Contains(got, substring) {
		t.Errorf("Response: %s, doesn not contain substring: %s", got, substring)
	}
}
