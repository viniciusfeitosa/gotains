package gotains

import (
	"errors"
	"reflect"
)

// Contains is to verify if an iten or more exists in a slice or array.
// The first param is an Slice or Array of any type.
// The "second" is a variadic arguments to verify with the content of the first argument
// The return is true if all variadic arguments exists in your slice or array
func Contains(slc interface{}, verItems ...interface{}) (bool, error) {
	s := reflect.ValueOf(slc)
	if s.Kind() != reflect.Slice && s.Kind() != reflect.Array {
		return false, errors.New("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	itensToCheck := len(verItems)
	checks := 0
	for _, r := range ret {
		for _, v := range verItems {
			if v == r {
				checks++
			}
		}
	}

	if checks == itensToCheck {
		return true, nil
	}
	return false, nil
}
