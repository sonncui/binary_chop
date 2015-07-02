package binary_chop_test

import (
	. "github.com/sonncui/binary_chop"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Binary", func() {
    Describe("test binary chop", func(){
         Context("with empty array", func(){
	     It("should return -1", func(){
	         Expect(Chop(3, []int{})).To(Equal(-1))
	     })
	 })

	 Context("with single element array", func(){
	     It("should return 0 when given 1 and [1]", func(){
	         Expect(Chop(1,[]int{1})).To(Equal(0))
	     })

	     It("should return -1 when given 3 and [1]", func(){
	         Expect(Chop(3, []int{1})).To(Equal(-1))
	     })
	 })

	 Context("when given multiple numbers", func(){
	     It("should return 0 when given 1 and [1,3,5]", func(){
	         Expect(Chop(1, []int{1,3,5})).To(Equal(0))
	     })

	     It("should return 1 when given 3 and [1,3,5]", func(){
	         Expect(Chop(3, []int{1,3,5})).To(Equal(1))
	     })
	 })
    })

})
