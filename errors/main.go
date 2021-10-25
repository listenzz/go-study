package main

import (
	"bufio"
	stderrors "errors"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

var (
	ErrNotFound = stderrors.New("找不到对象")
)

func main() {
	fmt.Println("--")
}

func doStuff() error {
	return errors.Wrap(ErrNotFound, "干活的时候")
}

type MyError struct {
	Msg  string
	Code string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Msg: %s, Code: %s", e.Msg, e.Code)
}

func NewMyError(msg string, code string) *MyError {
	return &MyError{Msg: msg, Code: code}
}

func IsMyError(err error) bool {
	_, ok := err.(*MyError)
	return ok
}

func CountLines1(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)

	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}

	if err != io.EOF {
		return 0, err
	}
	return lines, nil
}

func CountLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	lines := 0

	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

type Header struct {
	Key, Value string
}

type Status struct {
	Code   int
	Reason string
}

func WriteResponse1(w io.Writer, st Status, headers []Header, body io.Reader) error {
	_, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
	if err != nil {
		return err
	}

	for _, h := range headers {
		_, err := fmt.Fprintf(w, "%s: %s\r\n", h.Key, h.Value)
		if err != nil {
			return err
		}
	}

	if _, err := fmt.Fprint(w, "\r\n"); err != nil {
		return err
	}

	_, err = io.Copy(w, body)
	return err
}

type errWriter struct {
	io.Writer
	err error
}

func (w *errWriter) Write(p []byte) (int, error) {
	if w.err != nil {
		return 0, w.err
	}
	var n int
	n, w.err = w.Writer.Write(p)
	return n, w.err
}

func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ew := &errWriter{Writer: w}

	fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)

	for _, h := range headers {
		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}

	fmt.Fprint(ew, "\r\n")
	io.Copy(ew, body)

	return ew.err
}
