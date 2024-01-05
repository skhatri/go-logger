# go-logger
go wrapper for logrus


### Import

```
go get -u github.com/skhatri/go-logger
```

### Using it

```go
var LOG = logging.NewLogger("service-1")


func someFunc() {
  LOG.WithTask("do-operation").WithAttribute("key", "value").Info()

}
```

### Full example

```go
package main

import (
	"github.com/skhatri/go-logger/logging"
)

var logger = logging.NewLogger("app")

func main() {
	person := make(map[string]interface{})
	person["name"] = "John"
	person["dateOfBirth"] = "1990-10-01"
	person["token"] = "abc"
	logger.WithTask("start-app").WithAttribute("name", "elastic").
		WithAttributes(person).
		WithAttribute("user_token", "xusud93").
		WithAttribute("dob", "").
		Info()
}
```

Expected Result:

```json
{"dateOfBirth":"***","dob":"***","level":"info","msg":"","name":"John","source":"app","task":"start-app","time":"2024-01-05T19:43:49+11:00","token":"***","user_token":"***"}
```

