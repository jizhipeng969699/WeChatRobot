package singleton

import (
	"sync"
	"testing"
)

func TestGetSingleton(t *testing.T) {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(group *sync.WaitGroup) {
			defer group.Done()
			s := GetSingleton()
			t.Logf("%p\n", s)
		}(wg)
	}

	wg.Wait()
}
