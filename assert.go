package assert

import "runtime"
import "io/ioutil"
import "strings"
import "path"
import "reflect"

import "github.com/sdegutis/shellcolors"

type TestDriver interface {
  Errorf(format string, args ...interface{})
}

func auxiliaryInfo(extraStacks int) (filename string, line int, code string) {
  _, file, line, _ := runtime.Caller(3 + extraStacks)
  buf, _ := ioutil.ReadFile(file)
  filename = path.Base(file)
  code = strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
  return
}

func codeErrorString(extraStacks int) string {
  filename, line, code := auxiliaryInfo(extraStacks)
  return col.Red.Fmt("\n%v:%d\n%s", filename, line, code)
}

func Equals(t TestDriver, got, expected interface{}) {
  if got != expected {
    t.Errorf(codeErrorString(0) + col.Purple.Fmt("\n\n\texpected: %#v\n\t     got: %#v", expected, got))
  }
}

func DeepEquals(t TestDriver, got, expected interface{}) {
  if !reflect.DeepEqual(got, expected) {
    t.Errorf(codeErrorString(0) + col.Purple.Fmt("\n\n\texpected: %#v\n\t     got: %#v", expected, got))
  }
}

func NotEquals(t TestDriver, got, expected interface{}) {
  if got == expected {
    t.Errorf(codeErrorString(0) + col.Purple.Fmt("\n\n\tunexpectedly got: %#v", got))
  }
}

func True(t TestDriver, got bool) {
  if got != true {
    t.Errorf(codeErrorString(0))
  }
}

func False(t TestDriver, got bool) {
  if got != false {
    t.Errorf(codeErrorString(0))
  }
}

func Errorf(t TestDriver, format string, args ...interface{}) {
  format = "\t" + strings.Replace(format, "\n", "\n\t", -1) // indent every line once
  t.Errorf(codeErrorString(1) + "\n\n" + col.Purple.Fmt(format, args...))
}

func StringContains(t TestDriver, full, fragment string) {
  if !strings.Contains(full, fragment) {
    Errorf(t, "   expected: %#v\n to contain: %#v", full, fragment)
  }
}
