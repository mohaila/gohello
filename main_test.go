package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHello(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Canary Suite")
}

var _ = Describe("Canary Test", func() {
	It("True is always equal true", func() {
		Expect(true).To(Equal(true))
	})

	It("False and true are always different", func() {
		Expect(false).NotTo(Equal(true))
	})
})
