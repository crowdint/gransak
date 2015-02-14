package gransak

import (
	//"net/http"
	//"net/http/httptest"
	"fmt"
	"testing"
)

func TestGransak(t *testing.T) {
	//cont / or / and
	expected := "first_name LIKE {{v}} OR last_name LIKE {{v}}"
	sql, params := Gransak.ToSql("first_name_or_last_name_cont", "cone")
	strParams := toString(params)

	if sql != expected || strParams != "[%cone%]" {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s with Params: %s", sql, expected, strParams)
	}

	expected = "first_name LIKE {{v}} AND last_name LIKE {{v}}"
	sql, params = Gransak.ToSql("first_name_and_last_name_cont", "cone")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_cont / or / and
	expected = "first_name NOT LIKE {{v}} OR last_name NOT LIKE {{v}}"
	sql, params = Gransak.ToSql("first_name_or_last_name_not_cont", "cone")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//matches / or
	expected = "first_name LIKE {{v}} OR last_name LIKE {{v}}"
	sql, params = Gransak.ToSql("first_name_or_last_name_matches", "cone")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//does_not_match / or
	expected = "first_name NOT LIKE {{v}} OR last_name NOT LIKE {{v}}"
	sql, params = Gransak.ToSql("first_name_or_last_name_does_not_match", "cone")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//lt
	expected = "age < {{v}}"
	sql, params = Gransak.ToSql("age_lt", 29)
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//gt
	expected = "age > {{v}}"
	sql, params = Gransak.ToSql("age_gt", 29)
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//lteq
	expected = "age <= {{v}}"
	sql, params = Gransak.ToSql("age_lteq", 29)
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//gteq
	expected = "age >= {{v}}"
	sql, params = Gransak.ToSql("age_gteq", 29)
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//eq / or / and
	expected = "first_name = {{v}} AND last_name = {{v}}"
	sql, params = Gransak.ToSql("first_name_and_last_name_eq", "cone")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	expected = "age = {{v}}"
	sql, params = Gransak.ToSql("age_eq", 29)
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	expected = "age = {{v}} OR years = {{v}}"
	sql, params = Gransak.ToSql("age_or_years_eq", 29)
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_eq / or
	expected = "age <> {{v}} OR years <> {{v}}"
	sql, params = Gransak.ToSql("age_or_years_not_eq", 29)
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//start
	expected = "name LIKE {{v}}"
	sql, params = Gransak.ToSql("name_start", "cone")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_start
	expected = "name NOT LIKE {{v}}"
	sql, params = Gransak.ToSql("name_not_start", "cone")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//end
	expected = "name LIKE {{v}}"
	sql, params = Gransak.ToSql("name_end", "cone")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_end
	expected = "name NOT LIKE {{v}}"
	sql, params = Gransak.ToSql("name_not_end", "cone")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//true
	expected = "is_programmer = 't'"
	sql, params = Gransak.ToSql("is_programmer_true", "1")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_true
	expected = "is_programmer <> 't'"
	sql, params = Gransak.ToSql("is_programmer_not_true", "1")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//false
	expected = "is_programmer = 'f'"
	sql, params = Gransak.ToSql("is_programmer_false", "1")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_false
	expected = "is_programmer <> 'f'"
	sql, params = Gransak.ToSql("is_programmer_not_false", "1")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//present
	expected = "required_field IS NOT NULL AND required_field <> ''"
	sql, params = Gransak.ToSql("required_field_present", "1")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//blank
	expected = "required_field IS NULL OR required_field = ''"
	sql, params = Gransak.ToSql("required_field_blank", "1")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//null
	expected = "required_field IS NULL"
	sql, params = Gransak.ToSql("required_field_null", "1")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_null
	expected = "required_field IS NOT NULL"
	sql, params = Gransak.ToSql("required_field_not_null", "1")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//in
	expected = "age IN ({{v}},{{v}},{{v}})"
	sql, params = Gransak.ToSql("age_in", "28..30")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	sql, params = Gransak.ToSql("age_in", "[28,29,30]")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	sql, params = Gransak.ToSql("age_in", []int{28, 29, 30})
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_in
	expected = "age NOT IN ({{v}},{{v}},{{v}})"
	sql, params = Gransak.ToSql("age_not_in", "28..30")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	sql, params = Gransak.ToSql("age_not_in", "[28,29,30]")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//cont_any
	expected = "user_name LIKE {{v}} OR user_name LIKE {{v}}"
	sql, params = Gransak.ToSql("user_name_cont_any", "%w(cone carlos)")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_cont_any
	expected = "user_name NOT LIKE {{v}} AND user_name NOT LIKE {{v}}"
	sql, params = Gransak.ToSql("user_name_not_cont_any", "%w(cone carlos)")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//we can even do this
	expected = "user_name LIKE {{v}} AND last_name = {{v}}"
	sql, params = Gransak.ToSql("user_name_cont_and_last_name_eq", "%w(cone gutierrez)")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	expected = "user_name LIKE {{v}} AND last_name IS NOT NULL AND last_name <> ''"
	sql, params = Gransak.ToSql("user_name_cont_and_last_name_present", "cone")
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//this doesn't work very well
	//expected = "user_name = '{{v}}' AND last_name LIKE '%gutierrez%' OR last_name LIKE '%gutierrez%'"
	//sql = Gransak.ToSql("user_name_eq_and_last_name_cont_any", "%w(cone gutierrez)")

	//Adding a select statement (only if a table name was specified)
	//expected = "SELECT * FROM conejo WHERE user_name LIKE '%{{v}}%' AND last_name = 'gutierrez'"
	//sql, params = Gransak.Table("conejo").ToSql("user_name_cont_and_last_name_eq", "%w(cone gutierrez)")

	//if sql != expected {
	//t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	//}

	//Has word "not" but is not "not_equal" nor "not_in"
	//so it must be part of the field's name
	expected = "field_not_operator = {{v}}"
	sql, params = Gransak.ToSql("field_not_operator_eq", 29)
	strParams = toString(params)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//t.Error("stop")
}

//func TestFromRequest(t *testing.T) {
//var sql string

//handler := func(w http.ResponseWriter, r *http.Request) {
//sql = Gransak.FromRequest(r)
//}

//req, err := http.NewRequest("GET", "http://gransak.com/params?q[name_eq]=cone&q[last_name_eq]=Gutierrez", nil)
//if err != nil {
//panic(err)
//}

//w := httptest.NewRecorder()
//handler(w, req)

//expected := "name = '{{v}}' AND last_name = 'Gutierrez'"

//if sql != expected {
//t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
//}
//}

func toString(elem interface{}) string {
	return fmt.Sprintf("%v", elem)
}
