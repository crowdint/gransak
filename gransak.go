package gransak

import (
	"net/http"
	"net/url"

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

func (this *GransakFilter) ToSql(r *http.Request) string {
	return this.Parse()
}

func (this *GransakFilter) FromRequest(r *http.Request) string {
	return parseRequest(r)
}

func (this *GransakFilter) FromUrlValues(v url.Values) string {
	return parseUrlValues(v)
}
