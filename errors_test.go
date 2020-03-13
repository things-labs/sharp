package assist

import (
	"errors"
	"testing"
)

func TestErrors(t *testing.T) {
	errs := []error{errors.New("First"), errors.New("Second")}

	gErrs := Errors(errs)
	gErrs = gErrs.Combine(nil)
	gErrs = gErrs.Combine(errors.New("Third"))
	gErrs = gErrs.Combine(gErrs)

	if gErrs.Error() != "First; Second; Third" {
		t.Fatalf("Gave wrong error, got %s", gErrs.Error())
	}
	if length := len(gErrs.Errors()); length != 3 {
		t.Fatalf("Length of errors, got %+v", length)
	}
	if len(gErrs.Errors()) == 0 {
		t.Fatalf("Has wrong error in errors, got %+v", gErrs.Errors())
	}
}
