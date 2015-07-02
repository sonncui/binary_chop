package binary_chop_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBinaryChop(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BinaryChop Suite")
}
