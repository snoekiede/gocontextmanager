# Introduction

This is a simple implementation of a context manager in Go, in the vein of Python's 'with' statement and C#'s 'using' statement, providing automatic disposal of objects, using the provided Dispose method.

## Example use
Imagine you are working with an imaginary *Memoryfile* struct which looks like this:
```
type MemoryFile struct {
	name string
	open bool
}

func (mf *MemoryFile) Open() {
	mf.open = true
	fmt.Printf("Opening %s\n", mf.name)
}

func (mf *MemoryFile) Close() {
	mf.open = false
	fmt.Printf("Closing %s\n", mf.name)
}

func (mf *MemoryFile) Dispose() {
	fmt.Printf("Disposing %s\n", mf.name)
	mf.Close()
}

func (mf *MemoryFile) Read() (string, error) {
	fmt.Printf("Reading %s\n", mf.name)
	return mf.name, nil
}

```
As you can see, we have put in a *Dispose()* method, the file gets closed. You can use with our resource manager as follows:
~~~
import (
	"fmt"
	contextmanager "github.com/snoekiede/gocontextmanager"
)

func main() {
	mf := &MemoryFile{name: "test.txt"}
	text := "Nothing read"
	text, err := WithResource(mf, func(a *MemoryFile) (string, error) {
		a.Open()
		result, err := a.Read()
		return result, err
	}, func(element *MemoryFile) {
		element.Dispose()
	})
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	fmt.Printf("Read %s\n", text)
}
~~~

