package evenOdd

import (
	"fmt"
)

func CreateAndPrintEvenOdd() {
	type nb []int 

	iter := make(nb, 10)
    for i := range iter {
        iter[i] = 1 + i
    }

	for _, nmbr := range iter {
		if(nmbr % 2 == 0) {
			fmt.Printf("%v - even\n", nmbr) 
		} else {
			fmt.Printf("%v - odd\n", nmbr) 
		}
	}

}
