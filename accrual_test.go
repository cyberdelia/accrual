package accrual

import (
	"fmt"
	"testing"
	"time"
)

func TestDetection(t *testing.T) {
	d := NewDetector(1, NewMemoryWindow(10))
	for i := 0; i < 12; i++ {
		d.Heartbeat()
		time.Sleep(10 * time.Millisecond)
	}
	if d.Failed() {
		t.Error("early mistaken detection")
	}
	time.Sleep(100 * time.Millisecond)
	if !d.Failed() {
		t.Error("late mistaken detection")
	}
}

func ExampleDetector() {
	w := NewMemoryWindow(10)
	d := NewDetector(1, w)
	d.Heartbeat()
	if d.Failed() {
		fmt.Println("Failure dectected")
	}
}
