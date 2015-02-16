package gransak

import (
	"strconv"
	"strings"

	"github.com/crowdint/gransak/core"
)

func ReplaceByEngineHolders(template string, parsedParams []interface{}) string {
	if core.ENGINE == core.MYSQL_ENGINE {
		template = strings.Replace(template, core.VALUE_HOLDER, "?", -1)
	} else {
		numParams := len(parsedParams)

		for i := 1; i <= numParams; i++ {
			template = strings.Replace(
				template,
				core.VALUE_HOLDER,
				"$"+strconv.Itoa(i),
				1,
			)
		}
	}

	return template
}
