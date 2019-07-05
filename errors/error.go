package errors

import (
	"fmt"

	ierror "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/error"
)

// Errorf creates an error according to a format specifier
//
// By default Errorf return an internal error with code `FF000`
func Errorf(format string, a ...interface{}) *ierror.Error {
	return ierror.New(fmt.Sprintf(format, a...)).SetCode(internalErrCode)
}

// FromError cast a golang error into an internal Error
//
// if `err` is an internal then it is returned
func FromError(err error) *ierror.Error {
	if err == nil {
		return nil
	}

	ierr, ok := err.(*ierror.Error)
	if !ok {
		return Errorf(err.Error())
	}

	return ierr
}

// isErrorClass returns whether code belongs to a base error class
//
// While codes are uint64 for performance purposes they should be seen as 5 nibble codes encoded on 20 bits
//
// For code to be of class base
//  - 2 first nibbles must be identical (e.g. DB300 belongs to class DB000 but DB300 doesn't belong to 9F000)
//  - if base 3rd nibble is non zero then 3rd nibble must be identical (e.g. DB201 belongs to DB200 but DB300 doesn't belong to DB200)
func isErrorClass(code, base uint64) bool {
	// Error codes have a 5 hex reprensentation (<=> 20 bits representation)
	//  - (code^base)&255<<12 compute difference between 2 first nibbles (bits 13 to 20)
	//  - (code^base)&(base&15<<8) compute difference between 3rd nibble in case base 3rd nibble is non zero (bits 9 to 12)
	return (code^base)&(255<<12+base&15<<8) == 0
}
