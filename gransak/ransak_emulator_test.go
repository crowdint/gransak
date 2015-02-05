package filter

import (
	"testing"
)

func TestQueryParser(t *testing.T) {
	ransak := NewRansakEmulator()

	//cont / or / and
	expected := "first_name LIKE '%cone%' OR last_name LIKE '%cone%'"
	sql := ransak.ToSql("first_name_or_last_name_cont", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	expected = "first_name LIKE '%cone%' AND last_name LIKE '%cone%'"
	sql = ransak.ToSql("first_name_and_last_name_cont", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_cont / or / and
	expected = "first_name NOT LIKE '%cone%' OR last_name NOT LIKE '%cone%'"
	sql = ransak.ToSql("first_name_or_last_name_not_cont", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//matches / or
	expected = "first_name LIKE 'cone' OR last_name LIKE 'cone'"
	sql = ransak.ToSql("first_name_or_last_name_matches", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//does_not_match / or
	expected = "first_name NOT LIKE 'cone' OR last_name NOT LIKE 'cone'"
	sql = ransak.ToSql("first_name_or_last_name_does_not_match", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//lt
	expected = "age < 29"
	sql = ransak.ToSql("age_lt", 29)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//gt
	expected = "age > 29"
	sql = ransak.ToSql("age_gt", 29)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//lteq
	expected = "age <= 29"
	sql = ransak.ToSql("age_lteq", 29)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//gteq
	expected = "age >= 29"
	sql = ransak.ToSql("age_gteq", 29)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//eq / or / and
	expected = "first_name = 'cone' AND last_name = 'cone'"
	sql = ransak.ToSql("first_name_and_last_name_eq", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	expected = "age = 29"
	sql = ransak.ToSql("age_eq", 29)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	expected = "age = 29 OR years = 29"
	sql = ransak.ToSql("age_or_years_eq", 29)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_eq / or
	expected = "age <> 29 OR years <> 29"
	sql = ransak.ToSql("age_or_years_not_eq", 29)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//start
	expected = "name LIKE 'cone%'"
	sql = ransak.ToSql("name_start", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_start
	expected = "name NOT LIKE 'cone%'"
	sql = ransak.ToSql("name_not_start", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//end
	expected = "name LIKE '%cone'"
	sql = ransak.ToSql("name_end", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_end
	expected = "name NOT LIKE '%cone'"
	sql = ransak.ToSql("name_not_end", "cone")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//true
	expected = "is_programmer = 't'"
	sql = ransak.ToSql("is_programmer_true", "1")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_true
	expected = "is_programmer <> 't'"
	sql = ransak.ToSql("is_programmer_not_true", "1")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//false
	expected = "is_programmer = 'f'"
	sql = ransak.ToSql("is_programmer_false", "1")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_false
	expected = "is_programmer <> 'f'"
	sql = ransak.ToSql("is_programmer_not_false", "1")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//present
	expected = "required_field IS NOT NULL AND required_field <> ''"
	sql = ransak.ToSql("required_field_present", "1")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//blank
	expected = "required_field IS NULL OR required_field = ''"
	sql = ransak.ToSql("required_field_blank", "1")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//null
	expected = "required_field IS NULL"
	sql = ransak.ToSql("required_field_null", "1")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_null
	expected = "required_field IS NOT NULL"
	sql = ransak.ToSql("required_field_not_null", "1")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//in
	expected = "age IN (28,29,30)"
	sql = ransak.ToSql("age_in", "28..30")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	expected = "age IN (28,29,30)"
	sql = ransak.ToSql("age_in", "[28,29,30]")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//cont_any
	expected = "user_name LIKE '%cone%' OR user_name LIKE '%carlos%'"
	sql = ransak.ToSql("user_name_cont_any", "%w(cone carlos)")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//not_cont_any
	expected = "user_name NOT LIKE '%cone%' AND user_name NOT LIKE '%carlos%'"
	sql = ransak.ToSql("user_name_not_cont_any", "%w(cone carlos)")

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}

	//Has word "not" but is not "not_equal" nor "not_in"
	//so it must be part of the field's name
	expected = "field_not_operator = 29"
	sql = ransak.ToSql("field_not_operator_eq", 29)

	if sql != expected {
		t.Errorf("Mismatch Error:\nGot: %s \nWanted: %s", sql, expected)
	}
}
