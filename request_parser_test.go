package gransak

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

	got, gparams := parseUrlValues(params)

	substring := "last_name = {{v}}"

	if !strings.Contains(got, substring) {
		t.Errorf("Response: %s, doesn not contain substring: %s", got, substring)
	}

	substring = "name = {{v}}"

	if !strings.Contains(got, substring) {
		t.Errorf("Response: %s, doesn not contain substring: %s", got, substring)
	}

	substring = "cp = {{v}}"

	if !strings.Contains(got, substring) {
		t.Errorf("Response: %s, doesn not contain substring: %s", got, substring)
	}

	if len(gparams) != 3 {
		t.Error("Invalid number of params")
	}
}
