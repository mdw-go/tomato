# timer
--
    import "github.com/mdwhatcott/tomato"


## Usage

#### type Timer

```go
type Timer struct {
}
```


#### func  SetTimer

```go
func SetTimer(duration time.Duration) *Timer
```

#### func (*Timer) Start

```go
func (this *Timer) Start()
```

#### func (*Timer) String

```go
func (this *Timer) String() string
```
