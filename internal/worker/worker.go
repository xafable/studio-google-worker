package worker

import (
	"sync"

	"github.com/xafable/studio-google-worker/internal/interfaces"
)

func Run(jobs []interfaces.Job) []error {
	var wg sync.WaitGroup
	wg.Add(len(jobs))

	errChn := make(chan error, len(jobs))

	for _, j := range jobs {
		go func(f func() error) {
			err := f()
			defer wg.Done()
			if err != nil {
				errChn <- err
			}
		}(j.Do)
	}

	wg.Wait()
	close(errChn)

	var errors []error
	for erc := range errChn {
		if erc != nil {
			errors = append(errors, erc)
		}
	}

	return errors
}
