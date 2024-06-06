package main

import (
	"fmt"
	"testing"
)

func TestPoint1(t *testing.T) {
	arr := []int{12, 54, 87, 33, 14, 125, 65, 1}
	for i := 0; i < 10000; i++ {
		lottery := RandomLottery(arr)
		fmt.Printf("%d ", lottery)
	}
}

func TestPoint2(t *testing.T) {

}
