package core

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func newGransakParam(param interface{}, kind reflect.Kind) *gransakParam {
	ellipsisRx := regexp.MustCompile(`^[\d]+[.]{2}[\d]+$`)

	arrayRx := regexp.MustCompile(`^\[[\d|,]*\]$`)

	wordListRx := regexp.MustCompile(`^\%w\([\w\s]+\)$`)

	rparam := &gransakParam{
		value:      param,
		kind:       kind,
		ellipsisRx: ellipsisRx,
		arrayRx:    arrayRx,
		wordListRx: wordListRx,
		parts:      []string{},
	}

	rparam.findStrRepresentation()

	return rparam
}

type gransakParam struct {
	value             interface{}
	kind              reflect.Kind
	StrRepresentation string
	ellipsisRx        *regexp.Regexp
	arrayRx           *regexp.Regexp
	wordListRx        *regexp.Regexp
	parts             []string
}

func (this *gransakParam) findStrRepresentation() {
	paramStr := fmt.Sprintf("%v", this.value)

	if str, isEllipsis := this.getFromEllipsis(paramStr); isEllipsis {
		this.StrRepresentation = str
		return
	}

	if str, isArray := this.getFromArray(paramStr); isArray {
		this.StrRepresentation = str
		return
	}

	if str, isWordList := this.getFromWordList(paramStr); isWordList {
		this.StrRepresentation = str
		return
	}

	this.StrRepresentation = paramStr
}

func (this *gransakParam) getFromEllipsis(param string) (string, bool) {
	if this.ellipsisRx.MatchString(param) {
		values := strings.Split(param, "..")

		//if it wasn't a number, the regexp would have failed,
		//so we can omit the error
		start, _ := strconv.Atoi(values[0])
		end, _ := strconv.Atoi(values[1])

		var strValues []string

		for i := start; i <= end; i++ {
			strValues = append(strValues, strconv.Itoa(i))
		}

		return strings.Join(strValues, ","), true
	}
	return "", false
}

func (this *gransakParam) getFromArray(param string) (string, bool) {
	if this.kind == reflect.Slice {
		param = strings.Replace(param, " ", ",", -1)
	}

	if this.arrayRx.MatchString(param) {
		r := regexp.MustCompile(`[\[|\]]`)

		param = r.ReplaceAllString(param, "")

		return strings.Trim(param, ","), true
	}
	return "", false
}

func (this *gransakParam) getFromWordList(param string) (string, bool) {
	if this.wordListRx.MatchString(param) {
		param = strings.Replace(param, "%w", "", 1)

		r := regexp.MustCompile(`[\(|\)]`)

		param = r.ReplaceAllString(param, "")

		this.parts = strings.Split(param, " ")
	}
	return "", false
}
