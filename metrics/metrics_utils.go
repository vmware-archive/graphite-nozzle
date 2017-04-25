package metrics

import (
	"fmt"
	"time"

            "github.com/pivotal-cf/graphite-nozzle/logging"
)

func Retry(attempts int, sleep time.Duration, op func() error) (err error) {
	for i := 0; ; i++ {
		err = op()
		if err == nil {
			return
		}

		if i >= (attempts - 1) {
			break
		}

		time.Sleep(sleep)

		logging.LogError("retrying after error:", err)
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
