package assert

import "runtime"
import "io/ioutil"
import "strings"
import "path"

type TestDriver interface {
  Errorf(format string, args ...interface{})
}

func assertEquals(t TestDriver, got, expected interface{}) {
  _, file, line, _ := runtime.Caller(1)
  buf, _ := ioutil.ReadFile(file)
  filename := path.Base(file)
  code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])

  t.Errorf("|||||| %v:%d ----- %s ----- expected: %#v, got: %#v", filename, line, code, expected, got)
}
