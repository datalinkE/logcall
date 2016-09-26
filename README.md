## Purpose
Logcall is a little wrapper which can log debug info together with a caller function name.

## Example

#### Code:

```go
package main

import (
	"github.com/datalinkE/logcall/debug"
)

func A() {
	debug.LogCall()
}

func B() {
	debug.LogCall()
	debug.Logf("Formated string from B (with a function name and some arguments %v)", 12345)
	A()
}

func main() {
	debug.Log("Simple string from main, just as with usual log.Print")
	debug.LogCallEnabled = true
	B()
}
```

#### Output:
```
2016/09/26 17:09:11 Simple string from main, just as with usual log.Print
2016/09/26 17:09:11 main.B
2016/09/26 17:09:11 main.B Formated string from B (with a function name and some arguments 12345)
2016/09/26 17:09:11 main.A
```

