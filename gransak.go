package gransak

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/crowdint/gransak/core"
)

var (
	Gransak *GransakFilter
)

type GransakFilter struct {
	core      *core.GransakCore
	tableName string
	engine    string
}

func init() {
	if Gransak == nil {
		core := core.NewGransak()
		Gransak = &GransakFilter{
			core: core,
		}
	}
}

func (this *GransakFilter) ToSql(input string, param interface{}) (string, []interface{}) {
	statement, parsedParams := this.core.Parse(input, param)

	result := this.appendSelectStatement(statement)

	this.tableName = ""

	return result, parsedParams
}

func (this *GransakFilter) Table(tableName string) *GransakFilter {
	this.tableName = tableName

	return this
}

func (this *GransakFilter) appendSelectStatement(statement string) string {
	table := strings.Trim(this.tableName, " ")

	if table != "" {
		return "SELECT * FROM " + table + " WHERE " + statement
	}
	return statement
}

func (this *GransakFilter) FromRequest(r *http.Request) (string, []interface{}) {
	return parseRequest(r)
}

func (this *GransakFilter) FromUrlValues(v url.Values) (string, []interface{}) {
	return parseUrlValues(v)
}

func (this *GransakFilter) SetEngine(engine string) {
	core.ENGINE = engine
}
