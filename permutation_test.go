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
			b = New(42, 13, 7, 17)
		})
		Describe("Map", func() {
			It("should not map to itself", func() {
				for i := uint64(0); i < runs; i++ {
					Expect(b.Map(i)).NotTo(BeNumerically("==", i))
				}
			})
			It("should be the reverse of Unmap", func() {
				for i := uint64(0); i < runs; i++ {
					Expect(b.Map(b.Unmap(42))).To(BeNumerically("==", 42))
				}
			})
		})
		Describe("Unmap", func() {
			It("should be the reverse of Map", func() {
				for i := uint64(0); i < runs; i++ {
					Expect(b.Unmap(b.Map(42))).To(BeNumerically("==", 42))
				}
			})
		})
	})
})
