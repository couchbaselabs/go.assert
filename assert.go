package assert

import "runtime"
import "io/ioutil"
import "strings"
import "path"

type TestDriver interface {
  Errorf(format string, args ...interface{})
}

func Equals(t TestDriver, got, expected interface{}) {
  if got != expected {
    _, file, line, _ := runtime.Caller(1)
    buf, _ := ioutil.ReadFile(file)
    filename := path.Base(file)
    code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
    t.Errorf("|||||| %v:%d ----- %s ----- expected: %#v ----- got: %#v", filename, line, code, expected, got)
  }
}

func True(t TestDriver, got bool) {
  if got != true {
    _, file, line, _ := runtime.Caller(1)
    buf, _ := ioutil.ReadFile(file)
    filename := path.Base(file)
    code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
    t.Errorf("|||||| %v:%d ----- %s", filename, line, code)
  }
}
