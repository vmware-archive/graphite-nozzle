package processors_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestProcessors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Processors Suite")
}
