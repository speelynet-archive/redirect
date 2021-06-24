package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	redirect "github.com/speelynet/redirect"
	"net/url"
)

var _ = Describe("the Redirect function", func() {
	It("should return a correctly modified URL string", func() {
		for _, v := range []string{"example.com", "example.com/", "example.com/example/path", "example.com/example/path/"} {
			u, _ := url.Parse("http://" + v)
			Expect(redirect.Redirect(u)).To(Equal("https://" + v))
		}
	})
	It("should retain URL queries", func() {
		const queries = "?example=test&example2=test"
		for _, v := range []string{"example.com/", "example.com/example/path"} {
			u, _ := url.Parse("http://" + v + queries)
			Expect(redirect.Redirect(u)).To(Equal("https://" + v + queries))
		}
	})
})
