package assist

import (
	"strings"
)

// Errors contains all happened errors
type Errors []error

// HasError return whether any error in Errors
func (sf Errors) HasError() bool {
	return len(sf) > 0
}

// GetErrors gets all errors that have occurred and returns a slice of errors (Error type)
func (errs Errors) GetErrors() []error {
	return errs
}

// Add adds an error to a given slice of errors
func (sf Errors) Add(newErrors ...error) Errors {
	for _, err := range newErrors {
		if err == nil {
			continue
		}

		if errors, ok := err.(Errors); ok {
			sf = sf.Add(errors...)
		} else {
			ok = true
			for _, e := range sf {
				if err == e {
					ok = false
				}
			}
			if ok {
				sf = append(sf, err)
			}
		}
	}
	return sf
}

// Error takes a slice of all errors that have occurred and returns it as a formatted string
func (errs Errors) Error() string {
	var errors = []string{}
	for _, e := range errs {
		errors = append(errors, e.Error())
	}
	return strings.Join(errors, "; ")
}
