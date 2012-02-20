package assert

import "testing"
import "fmt"
import "strings"

type FakeTester string

func (f *FakeTester) Errorf(format string, args ...interface{}) {
  *f = FakeTester(fmt.Sprintf(format, args...))
}

func TestAssert(t *testing.T) {
  var f FakeTester

  Foo := func() string { return "foo" }

  assertEquals(&f, Foo(), "bar")

  // should contain the line that caused the error
  if !strings.Contains(string(f), `assertEquals(&f, Foo(), "bar")`) {
    t.Errorf("assert equals error; got [%v]", f)
  }

  if !strings.Contains(string(f), `expected: "bar"`) {
    t.Errorf("assert equals error; got [%v]", f)
  }

  if !strings.Contains(string(f), `got: "foo"`) {
    t.Errorf("assert equals error; got [%v]", f)
  }

  // should contain no newlines
  if strings.Contains(string(f), "\n") {
    t.Errorf("assert equals error; got [%v]", f)
  }

  assertEquals(t, Foo(), "bar")
  assertEquals(t, Foo(), "bar")
  assertEquals(t, Foo(), "bar")
  assertEquals(t, Foo(), "bar")
  assertEquals(t, Foo(), "bar")
  assertEquals(t, Foo(), "bar")
}
