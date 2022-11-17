package errors

import (
	"fmt"
	"reflect"

	msm "github.com/creatachain/augusteum/msm/types"
)

const (
	// SuccessMSMCode declares an MSM response use 0 to signal that the
	// processing was successful and no error is returned.
	SuccessMSMCode = 0

	// All unclassified errors that do not provide an MSM code are clubbed
	// under an internal error code and a generic message instead of
	// detailed error string.
	internalMSMCodespace        = UndefinedCodespace
	internalMSMCode      uint32 = 1
)

// MSMInfo returns the MSM error information as consumed by the augusteum
// client. Returned codespace, code, and log message should be used as a MSM response.
// Any error that does not provide MSMCode information is categorized as error
// with code 1, codespace UndefinedCodespace
// When not running in a debug mode all messages of errors that do not provide
// MSMCode information are replaced with generic "internal error". Errors
// without an MSMCode information as considered internal.
func MSMInfo(err error, debug bool) (codespace string, code uint32, log string) {
	if errIsNil(err) {
		return "", SuccessMSMCode, ""
	}

	encode := defaultErrEncoder
	if debug {
		encode = debugErrEncoder
	}

	return msmCodespace(err), msmCode(err), encode(err)
}

// ResponseCheckTx returns an MSM ResponseCheckTx object with fields filled in
// from the given error and gas values.
func ResponseCheckTx(err error, gw, gu uint64, debug bool) msm.ResponseCheckTx {
	space, code, log := MSMInfo(err, debug)
	return msm.ResponseCheckTx{
		Codespace: space,
		Code:      code,
		Log:       log,
		GasWanted: int64(gw),
		GasUsed:   int64(gu),
	}
}

// ResponseDeliverTx returns an MSM ResponseDeliverTx object with fields filled in
// from the given error and gas values.
func ResponseDeliverTx(err error, gw, gu uint64, debug bool) msm.ResponseDeliverTx {
	space, code, log := MSMInfo(err, debug)
	return msm.ResponseDeliverTx{
		Codespace: space,
		Code:      code,
		Log:       log,
		GasWanted: int64(gw),
		GasUsed:   int64(gu),
	}
}

// QueryResult returns a ResponseQuery from an error. It will try to parse MSM
// info from the error.
func QueryResult(err error) msm.ResponseQuery {
	space, code, log := MSMInfo(err, false)
	return msm.ResponseQuery{
		Codespace: space,
		Code:      code,
		Log:       log,
	}
}

// The debugErrEncoder encodes the error with a stacktrace.
func debugErrEncoder(err error) string {
	return fmt.Sprintf("%+v", err)
}

// The defaultErrEncoder applies Redact on the error before encoding it with its internal error message.
func defaultErrEncoder(err error) string {
	return Redact(err).Error()
}

type coder interface {
	MSMCode() uint32
}

// msmCode test if given error contains an MSM code and returns the value of
// it if available. This function is testing for the causer interface as well
// and unwraps the error.
func msmCode(err error) uint32 {
	if errIsNil(err) {
		return SuccessMSMCode
	}

	for {
		if c, ok := err.(coder); ok {
			return c.MSMCode()
		}

		if c, ok := err.(causer); ok {
			err = c.Cause()
		} else {
			return internalMSMCode
		}
	}
}

type codespacer interface {
	Codespace() string
}

// msmCodespace tests if given error contains a codespace and returns the value of
// it if available. This function is testing for the causer interface as well
// and unwraps the error.
func msmCodespace(err error) string {
	if errIsNil(err) {
		return ""
	}

	for {
		if c, ok := err.(codespacer); ok {
			return c.Codespace()
		}

		if c, ok := err.(causer); ok {
			err = c.Cause()
		} else {
			return internalMSMCodespace
		}
	}
}

// errIsNil returns true if value represented by the given error is nil.
//
// Most of the time a simple == check is enough. There is a very narrowed
// spectrum of cases (mostly in tests) where a more sophisticated check is
// required.
func errIsNil(err error) bool {
	if err == nil {
		return true
	}
	if val := reflect.ValueOf(err); val.Kind() == reflect.Ptr {
		return val.IsNil()
	}
	return false
}

var errPanicWithMsg = Wrapf(ErrPanic, "panic message redacted to hide potentially sensitive system info")

// Redact replaces an error that is not initialized as an MSM Error with a
// generic internal error instance. If the error is an MSM Error, that error is
// simply returned.
func Redact(err error) error {
	if ErrPanic.Is(err) {
		return errPanicWithMsg
	}
	if msmCode(err) == internalMSMCode {
		return errInternal
	}

	return err
}
