package main

import (
	"testing"
	"reflect"
)

/*

Run the following test:

p1: .   .   .        .
         \          /
p2:       .  .     /
              \   /
p3:  . .       . .

*/

func TestProcesses(t *testing.T) {

	var p1, p2, p3 Process

	p1.Start(3,0)
	p2.Start(3,1)
	p3.Start(3,2)

	// TODO: fill in calls to p.Internal, p.Send(), p.Receive() correctly
    // The execution should match the comment above.

	expected_p1 := []int{4,2,4}
	expected_p2 := []int{2,2,0}
	expected_p3 := []int{2,2,4} 

	got1 := p1.clock
	if !reflect.DeepEqual(got1, expected_p1) {
		t.Fatalf(
			"expected %v\ngot      %v", expected_p1, got1,
		)
	}
	got2 := p2.clock
	if !reflect.DeepEqual(got2, expected_p2) {
		t.Fatalf(
			"expected %v\ngot      %v", expected_p2, got2,
		)
	}
	got3 := p3.clock
	if !reflect.DeepEqual(got3, expected_p3) {
		t.Fatalf(
			"expected %v\ngot      %v", expected_p3, got3,
		)
	}

}

func TestCompare3(t *testing.T) {
	t1 := []int{1,2,3}
	t2 := []int{2,1,2}

	expected := 0
	result := Compare(t1,t2)
	if(result != expected) {
		t.Errorf("Expected %v, Got %v\n", expected, result)
	}

	t1 = []int{1,2,3}
	t2 = []int{0,1,0}

	expected = 1
	result = Compare(t1,t2)
	if(result <= 0) {
		t.Errorf("Expected %v, Got %v\n", expected, result)
	}

	expected = -1
	result = Compare(t2,t1)
	if(result >= 0) {
		t.Errorf("Expected %v, Got %v\n", expected, result)
	}
}

func TestCompare2(t *testing.T) {
	if Compare([]int{1, 2}, []int{2, 3}) >= 0 {
		t.Error("expected first timestamp to be less than second")
	}

	if Compare([]int{3, 4}, []int{2, 3}) <= 0 {
		t.Error("expected first timestamp to be greater than second")
	}

	if Compare([]int{2, 1}, []int{1, 2}) != 0 {
		t.Error("expected concurrent timestamps")
	}
}