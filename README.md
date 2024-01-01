# go-logger
go wrapper for logrus


### Import

```
go get -u github.com/skhatri/go-logger
```

### Using it

```
var LOG = logging.NewLogger("service-1")


func someFunc() {
  LOG.WithTask("do-operation").WithAttribute("key", "value").Info()

}
```
