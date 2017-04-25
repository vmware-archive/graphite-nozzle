package metrics_test

import (
	. "github.com/pivotal-cf/graphite-nozzle/metrics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"errors"
)

var _ = Describe("metrics_utils", func() {

	Describe("#Retry", func() {

		It("it executes the function once if it succeds", func() {
			times_called := 0
			op := func () (err error) {
				times_called += 1
				return
			}

			err := Retry(3, 0, op)
			Expect(err).To(BeNil())
			Expect(times_called ).To(Equal(1))

		})

		It("it executes the function the configured amount of time then fails", func() {
			times_called := 0
			op := func () (err error) {
				times_called += 1
				return errors.New("This is an error")
			}

			err := Retry(3, 0, op)
			Expect(err).NotTo(BeNil())
			Expect(times_called ).To(Equal(3))
		})
	})

})
