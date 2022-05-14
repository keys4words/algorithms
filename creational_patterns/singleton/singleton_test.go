package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("Expected pointer to Singleton after calling GetInstance(), not nill")
	}
	expectedCounter := counter1

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count: %d\n", currentCount)
	}
	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("Expected same instance in counter2 but go different instance")
	}
	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling AddOne() using the second counter, but go %d\n", currentCount)
	}
}
