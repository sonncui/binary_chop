package binary_chop
func Chop(number int, numbers []int) int {
     for _, n := range numbers {
         if n == number {
	    return 0;
	 } 
     }
     return -1
}