package main

import (
	stderrors "errors"
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func TestSentinelError(t *testing.T) {
	// test sentinel error
	err := doStuff()
	if err != nil && stderrors.Is(err, ErrNotFound) {
		fmt.Printf("1- %+v\n", err)
	}

	if err != nil {
		if e := errors.WithMessage(err, "发生了一些坏事情"); stderrors.Is(e, ErrNotFound) {
			fmt.Printf("2- %v\n", e)
		}
	}
}

func TestErrorType(t *testing.T) {
	err := NewMyError("产生 BUG 了", "ErrorCode")

	if IsMyError(err) {
		fmt.Printf("是自定义错误类型: %+v", err)
	}
}

func TestPkgWrap(t *testing.T) {
	err := pkgWrap()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func TestStdWrap(t *testing.T) {
	err := stdWrap()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func TestDoubleWrap(t *testing.T) {
	err := doubleWrap()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func TestNew(t *testing.T) {
	err := newError()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
