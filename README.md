# testutil

> Utilities around cmp to help testing go code

## Install

```shell
go get github.com/avakarev/go-testutil
```

## Usage

```go
package hello_test

import (
	"fmt"
	"testing"

	"github.com/avakarev/go-testutil"
)

func Hello() (string, error) {
	return "Hello World", nil
}

func TestHello(t *testing.T) {
	s, err := Hello()
	testutil.MustNoErr(err, t)
	testutil.Diff("Hello World", s, t)
}
```


## License

`go-testutil` is licensed under MIT license. (see [LICENSE](./LICENSE))
