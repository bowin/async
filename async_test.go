package async
import (
	"testing"
	"time"
	"fmt"
	"github.com/stretchr/testify/assert"
)


func TestAsyncJobs(t *testing.T) {
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
	for d := range r {
		fmt.Print(d)
	}
	ext := []interface {}([]interface {}{"hell1", "hell2", "hell3"})
	assert.Equal(t, ext, r)
}