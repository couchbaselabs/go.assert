package assert

import "runtime"
import "io/ioutil"
import "strings"
import "path"
import "reflect"
import "shellcolors"

type TestDriver interface {
  Errorf(format string, args ...interface{})
}

func auxiliaryInfo() (filename string, line int, code string) {
  _, file, line, _ := runtime.Caller(3)
  buf, _ := ioutil.ReadFile(file)
  filename = path.Base(file)
  code = strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
  return
}

func codeErrorString() string {
  filename, line, code := auxiliaryInfo()
  return col.Red.Fmt("\n%v:%d\n\n%s", filename, line, code)
}

func Equals(t TestDriver, got, expected interface{}) {
  if got != expected {
    t.Errorf(codeErrorString() + col.Purple.Fmt("\n\n\texpected: %#v\n\t     got: %#v", expected, got))
  }
}

func DeepEquals(t TestDriver, got, expected interface{}) {
  if !reflect.DeepEqual(got, expected) {
    t.Errorf(codeErrorString() + col.Purple.Fmt("\n\n\texpected: %#v\n\t     got: %#v", expected, got))
  }
}

func NotEquals(t TestDriver, got, expected interface{}) {
  if got == expected {
    t.Errorf(codeErrorString() + col.Purple.Fmt("\n\n\tunexpectedly got: %#v", got))
  }
}

func True(t TestDriver, got bool) {
  if got != true {
    t.Errorf(codeErrorString())
  }
}

func False(t TestDriver, got bool) {
  if got != false {
    t.Errorf(codeErrorString())
  }
}
