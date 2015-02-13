package gransak

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	ellipsisRx = regexp.MustCompile(`^[\d]+[.]{2}[\d]+$`)
	arrayRx    = regexp.MustCompile(`^\[[\d|,]*\]$`)
	wordListRx = regexp.MustCompile(`^\%w\([\w\s]+\)$`)
)

func newGransakParam(param interface{}, kind reflect.Kind) *gransakParam {

	rparam := &gransakParam{
		value: param,
		kind:  kind,
		parts: []interface{}{},
	}

	rparam.findStrRepresentation()

	return rparam
}

type gransakParam struct {
	value interface{}
	kind  reflect.Kind
	parts []interface{}
}

func (this *gransakParam) findStrRepresentation() {
	if this.kind == reflect.Slice {
		t := reflect.ValueOf(this.value)
		length := t.Len()
		//log.Printf("================LENGTH = %d", length)
		for i := 0; i < length; i++ {
			this.parts = append(this.parts, t.Index(i).Interface())
		}
		//this.parts = this.value
		return
	}

	paramStr := fmt.Sprintf("%v", this.value)

	if this.getFromEllipsis(paramStr) {
		return
	}

	if this.getFromArray(paramStr) {
		return
	}

	if this.getFromWordList(paramStr) {
		return
	}
}

func (this *gransakParam) getFromEllipsis(param string) bool {
	if ellipsisRx.MatchString(param) {
		values := strings.Split(param, "..")

		//if it wasn't a number, the regexp would have failed,
		//so we can omit the error
		start, _ := strconv.Atoi(values[0])
		end, _ := strconv.Atoi(values[1])

		for i := start; i <= end; i++ {
			this.parts = append(this.parts, i)
		}
	}
	return false
}

func (this *gransakParam) getFromArray(param string) bool {
	if arrayRx.MatchString(param) {
		r := regexp.MustCompile(`[\[|\]]`)

		param = r.ReplaceAllString(param, "")

		param = strings.Trim(param, ",")

		for _, item := range strings.Split(param, ",") {
			this.parts = append(this.parts, item)
		}

		return true
	}
	return false
}

func (this *gransakParam) getFromWordList(param string) bool {
	if wordListRx.MatchString(param) {
		param = strings.Replace(param, "%w", "", 1)

		r := regexp.MustCompile(`[\(|\)]`)

		param = r.ReplaceAllString(param, "")

		parts := strings.Split(param, " ")

		for _, part := range parts {
			this.parts = append(this.parts, part)
		}
	}
	return false
}