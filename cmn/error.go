package cmn

import (
	"strings"
	"errors"
	"fmt"
)

var ErrDelim = "--"

func MapString(ds []string, fn func(string) string) []string {
	var vals []string
	for _, d := range ds {
		vals = append(vals, fn(d))
	}
	return vals
}

func Errs(errs ...  error) error {
	var vals []string
	for _, err := range errs {
		vals = append(vals, err.Error())
	}
	return errors.New(strings.Join(vals, ErrDelim))
}

func ErrPipeWithMsg(msg string, errs ... error) error {
	for _, err := range errs {
		if err != nil {
			return errors.New(fmt.Sprintf("%s --> %s", msg, err.Error()))
		}
	}
	return nil
}
