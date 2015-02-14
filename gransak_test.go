package gransak

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGransak(t *testing.T) {
	//cont / or / and
	expected := "first_name LIKE ? OR last_name LIKE ?"
	sql, params := Gransak.ToSql("first_name_or_last_name_cont", "cone")
	strParams := toString(params)

	if sql != expected || strParams != "[%cone% %cone%]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	expected = "first_name LIKE ? AND last_name LIKE ?"
	sql, params = Gransak.ToSql("first_name_and_last_name_cont", "cone")
	strParams = toString(params)

	if sql != expected || strParams != "[%cone% %cone%]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//not_cont / or / and
	expected = "first_name NOT LIKE ? OR last_name NOT LIKE ?"
	sql, params = Gransak.ToSql("first_name_or_last_name_not_cont", "cone")
	strParams = toString(params)

	if sql != expected || strParams != "[%cone% %cone%]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//matches / or
	expected = "first_name LIKE ? OR last_name LIKE ?"
	sql, params = Gransak.ToSql("first_name_or_last_name_matches", "cone")
	strParams = toString(params)

	if sql != expected || strParams != "[cone cone]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//does_not_match / or
	expected = "first_name NOT LIKE ? OR last_name NOT LIKE ?"
	sql, params = Gransak.ToSql("first_name_or_last_name_does_not_match", "cone")
	strParams = toString(params)

	if sql != expected || strParams != "[cone cone]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//lt
	expected = "age < ?"
	sql, params = Gransak.ToSql("age_lt", 29)
	strParams = toString(params)

	if sql != expected || strParams != "[29]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//gt
	expected = "age > ?"
	sql, params = Gransak.ToSql("age_gt", 29)
	strParams = toString(params)

	if sql != expected || strParams != "[29]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//lteq
	expected = "age <= ?"
	sql, params = Gransak.ToSql("age_lteq", 29)
	strParams = toString(params)

	if sql != expected || strParams != "[29]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//gteq
	expected = "age >= ?"
	sql, params = Gransak.ToSql("age_gteq", 29)
	strParams = toString(params)

	if sql != expected || strParams != "[29]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//eq / or / and
	expected = "first_name = ? AND last_name = ?"
	sql, params = Gransak.ToSql("first_name_and_last_name_eq", "cone")
	strParams = toString(params)

	if sql != expected || strParams != "[cone cone]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	expected = "age = ?"
	sql, params = Gransak.ToSql("age_eq", 29)
	strParams = toString(params)

	if sql != expected || strParams != "[29]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	expected = "age = ? OR years = ?"
	sql, params = Gransak.ToSql("age_or_years_eq", 29)
	strParams = toString(params)

	if sql != expected || strParams != "[29 29]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//not_eq / or
	expected = "age <> ? OR years <> ?"
	sql, params = Gransak.ToSql("age_or_years_not_eq", 29)
	strParams = toString(params)

	if sql != expected || strParams != "[29 29]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//start
	expected = "name LIKE ?"
	sql, params = Gransak.ToSql("name_start", "cone")
	strParams = toString(params)

	if sql != expected || strParams != "[cone%]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//not_start
	expected = "name NOT LIKE ?"
	sql, params = Gransak.ToSql("name_not_start", "cone")
	strParams = toString(params)

	if sql != expected || strParams != "[cone%]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//end
	expected = "name LIKE ?"
	sql, params = Gransak.ToSql("name_end", "cone")
	strParams = toString(params)

	if sql != expected || strParams != "[%cone]" {

		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//not_end
	expected = "name NOT LIKE ?"
	sql, params = Gransak.ToSql("name_not_end", "cone")
	strParams = toString(params)

	if sql != expected || strParams != "[%cone]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//true
	expected = "is_programmer = 't'"
	sql, params = Gransak.ToSql("is_programmer_true", "1")
	strParams = toString(params)

	if sql != expected || strParams != "[]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//not_true
	expected = "is_programmer <> 't'"
	sql, params = Gransak.ToSql("is_programmer_not_true", "1")
	strParams = toString(params)

	if sql != expected || strParams != "[]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//false
	expected = "is_programmer = 'f'"
	sql, params = Gransak.ToSql("is_programmer_false", "1")
	strParams = toString(params)

	if sql != expected || strParams != "[]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//not_false
	expected = "is_programmer <> 'f'"
	sql, params = Gransak.ToSql("is_programmer_not_false", "1")
	strParams = toString(params)

	if sql != expected || strParams != "[]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//present
	expected = "required_field IS NOT NULL AND required_field <> ''"
	sql, params = Gransak.ToSql("required_field_present", "1")
	strParams = toString(params)

	if sql != expected || strParams != "[]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//blank
	expected = "required_field IS NULL OR required_field = ''"
	sql, params = Gransak.ToSql("required_field_blank", "1")
	strParams = toString(params)

	if sql != expected || strParams != "[]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//null
	expected = "required_field IS NULL"
	sql, params = Gransak.ToSql("required_field_null", "1")
	strParams = toString(params)

	if sql != expected || strParams != "[]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//not_null
	expected = "required_field IS NOT NULL"
	sql, params = Gransak.ToSql("required_field_not_null", "1")
	strParams = toString(params)

	if sql != expected || strParams != "[]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//in
	expected = "age IN (?,?,?)"
	sql, params = Gransak.ToSql("age_in", "28..30")
	strParams = toString(params)

	if sql != expected || strParams != "[28 29 30]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	sql, params = Gransak.ToSql("age_in", "[28,29,30]")
	strParams = toString(params)

	if sql != expected || strParams != "[28 29 30]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	sql, params = Gransak.ToSql("age_in", []int{28, 29, 30})
	strParams = toString(params)

	if sql != expected || strParams != "[28 29 30]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//not_in
	expected = "age NOT IN (?,?,?)"
	sql, params = Gransak.ToSql("age_not_in", "28..30")
	strParams = toString(params)

	if sql != expected || strParams != "[28 29 30]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	sql, params = Gransak.ToSql("age_not_in", "[28,29,30]")
	strParams = toString(params)

	if sql != expected || strParams != "[28 29 30]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//cont_any
	expected = "user_name LIKE ? OR user_name LIKE ?"
	sql, params = Gransak.ToSql("user_name_cont_any", "%w(cone carlos)")
	strParams = toString(params)

	if sql != expected || strParams != "[%cone %carlos]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//not_cont_any
	expected = "user_name NOT LIKE ? AND user_name NOT LIKE ?"
	sql, params = Gransak.ToSql("user_name_not_cont_any", "%w(cone carlos)")
	strParams = toString(params)

	if sql != expected || strParams != "[%cone% %carlos%]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	//Has word "not" but is not "not_equal" nor "not_in"
	//so it must be part of the field's name
	expected = "field_not_operator = ?"
	sql, params = Gransak.ToSql("field_not_operator_eq", 29)
	strParams = toString(params)

	if sql != expected || strParams != "[29]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

}

func TestFromRequest(t *testing.T) {
	var sql string
	var params []interface{}

	handler := func(w http.ResponseWriter, r *http.Request) {
		sql, params = Gransak.FromRequest(r)
	}

	req, err := http.NewRequest("GET", "http://gransak.com/params?q[name_eq]=cone&q[last_name_eq]=Gutierrez", nil)
	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()
	handler(w, req)

	expected := "name = ? AND last_name = ?"

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	if toString(params) != "[cone Gutierrez]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", toString(params), "[cone Gutierrez]")
	}
}

func toString(elem interface{}) string {
	return fmt.Sprintf("%v", elem)
}
