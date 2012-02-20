package assert

import "testing"
import "fmt"
import "strings"

type FakeTester struct {
  str string
  count int
}

func (f *FakeTester) Errorf(format string, args ...interface{}) {
  f.str = fmt.Sprintf(format, args...)
  f.count++
}

func TestValidAssert(t *testing.T) {
  var f FakeTester

  Foo := func() string { return "foo" }

  Equals(&f, Foo(), "foo")

  if f.count != 0 {
    t.Errorf("assert equals error; called %d times", f.count)
  }

  // should contain the line that caused the error
  if f.str != "" {
    t.Errorf("assert equals error; got [%v]", f)
  }
}

func TestFaltyAssert(t *testing.T) {
  var f FakeTester

  Foo := func() string { return "foo" }

  Equals(&f, Foo(), "bar")

  if f.count != 1 {
    t.Errorf("assert equals error; called %d times", f.count)
  }

  // should contain the line that caused the error
  if !strings.Contains(f.str, `Equals(&f, Foo(), "bar")`) {
    t.Errorf("assert equals error; got [%v]", f)
  }

  if !strings.Contains(f.str, `expected: "bar"`) {
    t.Errorf("assert equals error; got [%v]", f)
  }

  if !strings.Contains(f.str, `got: "foo"`) {
    t.Errorf("assert equals error; got [%v]", f)
  }

  // should contain no newlines
  if strings.Contains(f.str, "\n") {
    t.Errorf("assert equals error; got [%v]", f)
  }
}

func TestTrue(t *testing.T) {
  {
    var f FakeTester

    falsifier := func() bool { return false }

    True(&f, falsifier())

    Equals(t, f.count, 1)
    Equals(t, strings.Contains(f.str, `True(&f, falsifier())`), true)
    Equals(t, strings.Contains(f.str, "\n"), false)
  }

  {
    var f FakeTester

    truthifier := func() bool { return true }

    True(&f, truthifier())

    Equals(t, f.count, 0)
    Equals(t, f.str, "")
  }
}
