package interview_cases

import "sync"

func MergeChannels(inChans ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(inChans))
	for _, c := range inChans {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
