package assert

import "runtime"
import "io/ioutil"
import "strings"
import "path"

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
    t.Errorf("|||||| %v:%d ----- %s ----- expected: %#v ----- got: %#v", filename, line, code, expected, got)
  }
}

func True(t TestDriver, got bool) {
  if got != true {
    filename, line, code := auxiliaryInfo()
    t.Errorf("|||||| %v:%d ----- %s", filename, line, code)
  }
}

func False(t TestDriver, got bool) {
  if got != false {
    filename, line, code := auxiliaryInfo()
    t.Errorf("|||||| %v:%d ----- %s", filename, line, code)
  }
}
