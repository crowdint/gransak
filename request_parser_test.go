package gransak

import (
	"net/url"
	"strings"
	"testing"

	"github.com/crowdint/gransak/core"
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

	Gransak.SetEngine(core.POSTGRESQL_ENGINE)

	got, gparams := parseUrlValues(params)

	substring := "last_name"

	if !strings.Contains(got, substring) {
		t.Errorf("Response: %s, doesn not contain substring: %s", got, substring)
	}

	substring = "name"

	if !strings.Contains(got, substring) {
		t.Errorf("Response: %s, doesn not contain substring: %s", got, substring)
	}

	substring = "cp"

	if !strings.Contains(got, substring) {
		t.Errorf("Response: %s, doesn not contain substring: %s", got, substring)
	}

	if len(gparams) != 3 {
		t.Error("Invalid number of params")
	}
}
