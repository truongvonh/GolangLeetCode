package main

import (
	"./utils"
	"reflect"
	"testing"
)

func TestSumUp(t *testing.T) {
	resultFunc := utils.TwoSum([]int{2, 7, 11, 15}, 9)
	resultExpected := []int{0,1}

	if !reflect.DeepEqual(resultFunc, resultExpected) {
		t.Errorf("Test Sum up failed")
	}
}
