package asyncerrgroup

import (
	"sync"
)

type asyncGroup struct {
	wg      sync.WaitGroup
	errChan chan error
	errOnce sync.Once
}

func NewAsyncGroup() *asyncGroup {
	return &asyncGroup{
		errChan: make(chan error, 1),
	}
}

func (g *asyncGroup) Run(f func() error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		err := f()
		if err != nil {
			g.syncErr(err)
		}
	}()
}

func (g *asyncGroup) syncErr(err error) {
	g.errOnce.Do(func() {
		g.errChan <- err
		close(g.errChan)
	})
}

func (g *asyncGroup) Wait() error {
	go func() {
		g.wg.Wait()
		g.syncErr(nil)
	}()

	return <-g.errChan
}
