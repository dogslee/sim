package sim

import (
	"fmt"
	"testing"

	"golang.org/x/sync/errgroup"
)

func Test_Lock(t *testing.T) {
	mlock := &Mutex{}
	eg := &errgroup.Group{}

	a := []int{}
	for i := 0; i < 100; i++ {
		n := i
		eg.Go(func() error {
			mlock.Lock()
			a = append(a, n)
			mlock.Unlock()
			return nil
		})
	}

	eg.Wait()
	fmt.Printf("len=%#v\na=%v", len(a), a)
}
