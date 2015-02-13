package core

import (
	"reflect"
	"testing"
)

func TestGransakParam(t *testing.T) {
	//simple string parameter
	paramStr := "cone"
	wanted := paramStr

	param := newGransakParam(paramStr, reflect.String)

	if param.StrRepresentation != wanted {
		t.Errorf("Mismatch, wanted: %s got: %s", wanted, param.StrRepresentation)
	}

	//ellipsis string parameter
	paramStr = "1..10"
	wanted = "1,2,3,4,5,6,7,8,9,10"

	param = newGransakParam(paramStr, reflect.String)

	if param.StrRepresentation != wanted {
		t.Errorf("Mismatch, wanted: %s got: %s", wanted, param.StrRepresentation)
	}

	//array string parameter
	paramStr = "[1,2,3,4,5,6,7,8,9,10]"
	wanted = "1,2,3,4,5,6,7,8,9,10"

	param = newGransakParam(paramStr, reflect.String)

	if param.StrRepresentation != wanted {
		t.Errorf("Mismatch, wanted: %s got: %s", wanted, param.StrRepresentation)
	}

	//slice parameter
	paramSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wanted = "1,2,3,4,5,6,7,8,9,10"

	param = newGransakParam(paramSlice, reflect.Slice)

	if param.StrRepresentation != wanted {
		t.Errorf("Mismatch, wanted: %s got: %s", wanted, param.StrRepresentation)
	}

	//word list parameter
	paramStr = "%w(cone gutierrez)"
	wanted = "w(cone gutierrez)"

	param = newGransakParam(paramStr, reflect.String)

	if param.StrRepresentation != wanted {
		t.Errorf("Mismatch, wanted: %s got: %s", wanted, param.StrRepresentation)
	}

	if len(param.parts) < 2 {
		t.Errorf("Mismatch, wanted: %d got: %d", 2, len(param.parts))
		return
	}

	if param.parts[0] != "cone" || param.parts[1] != "gutierrez" {
		t.Errorf("Mismatch, first part (wanted: %s got: %s) second part (wanted: %s got: %s)", "cone", param.parts[0], "gutierrez", param.parts[1])
	}
}
