package irror

import (
	"fmt"

	"github.com/go-stack/stack"
)

type IRR struct {
	Message    string   `json:"msg"`
	Details    []string `json:"details"`
	Inner      error    `json:"-"`
	Stacktrace string   `json:"-"`
}

func New(msg string) *IRR {
	err := IRR{}
	err.Message = msg
	return &err
}

func (irr *IRR) Wrap(rootErr error) error {
	err := New(irr.Message)
	err.Inner = rootErr
	err.Stacktrace = irr.Stacktracer().String()
	return err
}

func (irr *IRR) Error() string {
	return fmt.Sprintf("Error: %s", irr.Message)
}

func (irr *IRR) AppendDetail(msg string) *IRR {
	irr.Details = append(irr.Details, msg)
	return irr
}

func (irr *IRR) Stacktracer() stack.CallStack {
	return stack.Trace().TrimBelow(stack.Caller(1)).TrimRuntime()
}
