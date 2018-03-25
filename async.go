package async
import (
	"sort"
)

// Result result of a single async job
type Result struct {
	index int
	value interface{}
}

// Results accumalated result of mulpiple async jobs
type Results []Result

func (cr Results) Len()int {
	return len(cr)
}
func (cr Results) Swap(i, j int) {
	cr[i], cr[j] = cr[j], cr[i]
}

func (cr Results) Less(i, j int)bool {
	return cr[i].index < cr[j].index
}

type fn func()interface{}

// DoJobs run multiple jobs concurrently
func DoJobs(fns ...fn) ([]interface{}, error) {
	c := make(chan Result)
	for index, f := range fns {
		go func(index int, f fn) {
			c<- Result{index: index, value: f()}
		}(index, f)
	}
	r := make(Results, 0)
	for i:=0; i<len(fns); i++ {
		r = append(r, <-c)
	}
	close(c)
	sort.Sort(r)
	arr := make([]interface{}, 0)
	for _, item := range r {
		arr = append(arr, item.value)
	}
	return arr, nil
}



