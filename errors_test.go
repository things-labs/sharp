package assist

import (
	"errors"
	"testing"
)

func TestErrorsCanBeUsedOutsideGorm(t *testing.T) {
	errs := []error{errors.New("First"), errors.New("Second")}

	gErrs := Errors(errs)
	gErrs = gErrs.Add(nil)
	gErrs = gErrs.Add(errors.New("Third"))
	gErrs = gErrs.Add(gErrs)

	if gErrs.Error() != "First; Second; Third" {
		t.Fatalf("Gave wrong error, got %s", gErrs.Error())
	}
	if length := len(gErrs.GetErrors()); length != 3 {
		t.Fatalf("Length of errors, got %+v", length)
	}
	if !gErrs.HasError() {
		t.Fatalf("Has wrong error in errors, got %+v", gErrs.HasError())
	}
}
