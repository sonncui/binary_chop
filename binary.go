package binary_chop
func Chop(number int, numbers []int) int {
     var start int = 0
     var end int = len(numbers)
     for start < end {
     	 if numbers[(end - start)/2] > number {
	     end = (end-start)/2 - 1
	 }

	 if numbers[(end-start)/2] < number {
	     start = (end-start)/2 + 1
	 }
	 
         if numbers [(end - start)/2] == number {
	    return (end - start)/2
	 } 
     }
     return -1
}