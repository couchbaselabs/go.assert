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
  _, file, line, _ := runtime.Caller(2)
  buf, _ := ioutil.ReadFile(file)
  filename = path.Base(file)
  code = strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
  return
}

func Equals(t TestDriver, got, expected interface{}) {
  if got != expected {
    filename, line, code := auxiliaryInfo()
    t.Errorf(col.Red.Fmt("\n%v:%d\n\n%s\n\n\texpected: %#v\n\t     got: %#v", filename, line, code, expected, got))
  }
}

func DeepEquals(t TestDriver, got, expected interface{}) {
  if !reflect.DeepEqual(got, expected) {
    filename, line, code := auxiliaryInfo()
    t.Errorf(col.Red.Fmt("\n%v:%d\n\n%s\n\n\texpected: %#v\n\t     got: %#v", filename, line, code, expected, got))
  }
}

func NotEquals(t TestDriver, got, expected interface{}) {
  if got == expected {
    filename, line, code := auxiliaryInfo()
    t.Errorf(col.Red.Fmt("\n%v:%d\n\n%s\n\n\tunexpectedly got: %#v", filename, line, code, got))
  }
}

func True(t TestDriver, got bool) {
  if got != true {
    filename, line, code := auxiliaryInfo()
    t.Errorf(col.Red.Fmt("\n%v:%d\n\n%s", filename, line, code))
  }
}

func False(t TestDriver, got bool) {
  if got != false {
    filename, line, code := auxiliaryInfo()
    t.Errorf(col.Red.Fmt("\n%v:%d\n\n%s", filename, line, code))
  }
}
