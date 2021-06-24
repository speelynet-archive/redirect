package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRedirect(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Redirect Suite")
}
