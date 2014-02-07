// build !appengine
package permutation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

const runs = 10000

func TestPermutation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Permutation Suite")
}

var _ = Describe("Permutation", func() {
	var b permutation
	Context("Using a test permutation", func() {
		BeforeEach(func() {
			b = New(42)
		})
		Describe("MapTo", func() {
			It("should work", func() {
				Expect(b.MapTo(42)).To(BeNumerically("==", 4627128764160949907))
			})
			It("should not map to itself", func() {
				for i := uint64(0); i < runs; i++ {
					Expect(b.MapTo(i)).NotTo(BeNumerically("==", i))
				}
			})
		})
		Describe("MapFrom", func() {
			It("should work", func() {
				Expect(b.MapFrom(4627128764160949907)).To(BeNumerically("==", 42))
			})
			It("should be the reverse of MapTo", func() {
				for i := uint64(0); i < runs; i++ {
					Expect(b.MapFrom(b.MapTo(i))).To(BeNumerically("==", i))
				}
			})
		})
	})
})
