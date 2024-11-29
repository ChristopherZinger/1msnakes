package arrays

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	testArr := make([]*int, 101)
	controlArray := make([]*int, 101)

	testArrLen := len(testArr)
	for i := 0; i < testArrLen; i++ {
		testArr[i] = &i
		controlArray[testArrLen-i-1] = &i
	}

	Reverse(testArr)

	for i, item := range testArr {
		fmt.Printf("%d : %d\n", item, controlArray[i])
		if item != controlArray[i] {
			fmt.Printf("%d != %d\n", item, controlArray[i])
			t.Error("Failed to reverse slice")
		}
	}

}
