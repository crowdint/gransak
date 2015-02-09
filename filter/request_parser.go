package filter

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func parseRequest(r *http.Request) string {
	params := r.URL.Query()

	return getGransakQuery(&params)
}

func parseUrlValues(params url.Values) string {
	return getGransakQuery(&params)
}

func getGransakQuery(params *url.Values) string {
	r := regexp.MustCompile(`^q\[[\w]+\]$`)
	var temp, sql string
	statements := []string{}

	for key, value := range *params {

		if r.MatchString(key) {
			temp = strings.Replace(key, "q[", "", 1)
			temp = strings.Replace(temp, "]", "", 1)

			sql = Gransak.ToSql(temp, value[0])

			if strings.Trim(sql, " ") != "" {
				statements = append(statements, sql)
			}
		}
	}

	return strings.Join(statements, " AND ")
}
