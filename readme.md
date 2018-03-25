## do async jobs concurently and return results with order
### Install
```bash
go get github.com/bowin/async
``` 

### Usage
```go
package main
import "github.com/bowin/async"
func main()
	f1 := func()interface{} {
		return "hell1"
	}
	f2 := func()interface{} {
		time.Sleep(time.Second)
		return "hell2"
	}
	f3 := func()interface{} {
		time.Sleep(time.Millisecond * 10)
		return "hell3"
	}
    r, _ := DoJobs(f1, f2, f3)
    // r will be array like {"hell1", "hell2", "hell3"}
}    
```
