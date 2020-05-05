package util

import (
	"github.com/go-errors/errors"
)

var (
	UnimplementedText = "Reached Unimplemented code"
	Shrug             = `¯\_(ツ)_/¯`
)

// duplicate of the generate_function.go file simply nicer to be able to say errors.Check(something)
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func WarnUnimplemented() error {
	Logger().Warnf(UnimplementedText + ": " + ColorInfo(Shrug))
	return errors.New("Unimplemented Error")
}
func ErrorUnimplemented() error {
	Logger().Errorf(UnimplementedText + ": " + ColorInfo(Shrug))
	return errors.New("Unimplemented Error")
}
func FatalUnimplemented() error {
	Logger().Fatalf(UnimplementedText + ": " + ColorInfo(Shrug))
	return errors.New("Unimplemented Error")
}
