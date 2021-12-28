package tests

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestError(t *testing.T)  {
	err1 := errors.New("test error")
	err2 := fmt.Errorf("wrapped error, %w", err1)
	t.Log(errors.Is(err2, err1))
	t.Log(errors.As(err2, &err1))
	t.Log(err2.Error())
}
