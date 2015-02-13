package gransak

import (
	//"net/http"
	//"net/url"
	"reflect"

	"github.com/crowdint/gransak/core"
)

var Gransak *GransakFilter

type GransakFilter struct {
	*core.GransakCore
}

func init() {
	if Gransak == nil {
		core := core.NewGransak()
		Gransak = &GransakFilter{
			core,
		}
	}
}

func (this *GransakFilter) ToSql(input string, param interface{}) (string, []interface{}) {
	paramSlice := newGransakParam(param, reflect.TypeOf(param).Kind())
	return this.Parse(input, len(paramSlice.parts)), paramSlice.parts
}

//func (this *GransakFilter) FromRequest(r *http.Request) string {
//return parseRequest(r)
//}

//func (this *GransakFilter) FromUrlValues(v url.Values) string {
//return parseUrlValues(v)
//}
