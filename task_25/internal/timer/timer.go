package timer

import "time"

func Sleep(duration time.Duration) {
	done := make(chan struct{})
	timer := time.NewTimer(duration)

	go func() {
		<-timer.C
		close(done)
	}()

	<-done
}
