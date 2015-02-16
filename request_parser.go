package gransak

import (
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func parseRequest(r *http.Request) (string, []interface{}) {
	params := r.URL.Query()

	return getGransakQuery(&params)
}

func parseUrlValues(params url.Values) (string, []interface{}) {
	return getGransakQuery(&params)
}

func getGransakQuery(params *url.Values) (string, []interface{}) {
	r := regexp.MustCompile(`^q\[[\w]+\]$`)
	var temp, sql string
	statements := []string{}
	parsedParams := []interface{}{}
	gparams := []interface{}{}

	for key, value := range *params {

		if r.MatchString(key) {
			temp = strings.Replace(key, "q[", "", 1)
			temp = strings.Replace(temp, "]", "", 1)

			sql, gparams = getSqlString(temp, value[0])

			parsedParams = append(parsedParams, gparams...)

			if strings.Trim(sql, " ") != "" {
				statements = append(statements, sql)
			}
		}
	}

	result := strings.Join(statements, " AND ")

	result = ReplaceByEngineHolders(result, parsedParams)

	result = Gransak.appendSelectStatement(result)

	Gransak.Table("")

	return result, parsedParams
}

func getSqlString(query, value string) (string, []interface{}) {

	if intVal, err := strconv.ParseInt(value, 0, 64); err == nil {
		return Gransak.core.Parse(query, intVal)
	}

	if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
		return Gransak.core.Parse(query, floatVal)
	}

	return Gransak.core.Parse(query, value)
}
