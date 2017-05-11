package metrics

import (
	"errors"
            "fmt"
	"time"

            "github.com/pivotal-cf/graphite-nozzle/logging"
)

func Retry(attempts int, sleep time.Duration, op func() error) (err error) {
    for i := 0; ; i++ {
        err = op()
        if err == nil {
             logging.LogStd(fmt.Sprintf("succeeded after %d attempts", i+1))
            return
        }

        if i >= (attempts - 1) {
	break
        }

        time.Sleep(sleep)

        logging.LogError(fmt.Sprintf("failed retrying, attempt %d of %d", i+1, attempts), err)
    }

    return errors.New(fmt.Sprintf("gave up after %d attempts", attempts))
}
