package main

import "testing"

func TestAdd(t *testing.T) {
	// Test case 1: Testing add function with integer inputs
	resultInt := add(1, 2, 3)
	if resultInt != 6 {
		t.Errorf("Expected result to be 6, but got %d", resultInt)
	}

	// Test case 2: Testing add function with float inputs
	resultFloat := add(1.0, 2.0, 3.1)
	if resultFloat != 6.1 {
		t.Errorf("Expected result to be 6.1, but got %f", resultFloat)
	}
}
