package xeroerrors

import "fmt"

type CustomError struct {
	Err        error
	Message    string
	ErrorTrace string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("error %v: %s /n Trace: %s", e.Err, e.Message, e.ErrorTrace)
}

func ErrorType(err error) error {
	if e, ok := FromError(err); ok {
		return e.Err
	}
	return err
}

func FullError(err error) error {
	e, ok := FromError(err)
	if ok {
		return fmt.Errorf("%s\n%s", e.Error(), e.ErrorTrace)
	}
	return fmt.Errorf("TraceError: %s,\n root error: %s,\n message: %s", e.ErrorTrace, e.Err, e.Message)
}

func FromError(err error) (*CustomError, bool) {
	if err == nil {
		return nil, false
	}
	if e, ok := err.(*CustomError); ok {
		return e, true
	}
	return &CustomError{Err: err}, true
}

func New(err error, message string) *CustomError {
	return &CustomError{Err: err, Message: message}
}

func AppendErrorTrace(err error, trace string) *CustomError {
	e, ok := FromError(err)
	if !ok {
		return &CustomError{Err: err, ErrorTrace: trace}
	}
	e.ErrorTrace = fmt.Sprintf("%s\n%s", e.ErrorTrace, trace)
	return e
}
