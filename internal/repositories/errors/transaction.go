package errors

import stdErrors "errors"

var ErrTransaction_badStatus = stdErrors.New("bad status")
var ErrTransaction_badAmount = stdErrors.New("bad amount")
var ErrTransaction_invoiceNotFound = stdErrors.New("invoice not found")
