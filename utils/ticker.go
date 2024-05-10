package utils

import "time"

func Schedule(f func(), interval time.Duration) chan struct{} {
	done := make(chan struct{})
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				f()
			}
		}
	}()
	return done
}
