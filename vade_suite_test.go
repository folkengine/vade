package vade_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestVade(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vade Suite")
}
