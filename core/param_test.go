package core

import (
	"fmt"
	"testing"
)

func TestGransakParam(t *testing.T) {
	//simple string parameter
	paramStr := "cone"
	wanted := "[cone]"

	param := newGransakParam(paramStr)
	param.parse("", "")

	if toString(param.parts) != wanted {
		t.Errorf("Mismatch, wanted: %s got: %s", wanted, toString(param.parts))
	}

	//ellipsis string parameter
	paramStr = "1..10"
	wanted = "[1 2 3 4 5 6 7 8 9 10]"

	param = newGransakParam(paramStr)
	param.parse("", "")

	if toString(param.parts) != wanted {
		t.Errorf("Mismatch, wanted: %s got: %s", wanted, toString(param.parts))
	}

	//array string parameter
	paramStr = "[1,2,3,4,5,6,7,8,9,10]"

	param = newGransakParam(paramStr)
	param.parse("", "")

	if toString(param.parts) != wanted {
		t.Errorf("Mismatch, wanted: %s got: %s", wanted, toString(param.parts))
	}

	//slice parameter
	paramSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	param = newGransakParam(paramSlice)
	param.parse("", "")

	if toString(param.parts) != wanted {
		t.Errorf("Mismatch, wanted: %s got: %s", wanted, toString(param.parts))
	}

	//word list parameter
	paramStr = "%w(cone gutierrez)"
	wanted = "[cone gutierrez]"

	param = newGransakParam(paramStr)
	param.parse("", "")

	if toString(param.parts) != wanted {
		t.Errorf("Mismatch, wanted: %s got: %s", wanted, toString(param.parts))
	}

	if len(param.parts) < 2 {
		t.Errorf("Mismatch, wanted: %d got: %d", 2, len(param.parts))
		return
	}

	if param.parts[0] != "cone" || param.parts[1] != "gutierrez" {
		t.Errorf("Mismatch, first part (wanted: %s got: %s) second part (wanted: %s got: %s)", "cone", param.parts[0], "gutierrez", param.parts[1])
	}
}

func toString(elem interface{}) string {
	return fmt.Sprintf("%v", elem)
}
