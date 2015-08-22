// Accrual is an implementation of the φ Accrual Failure Detector.
// For details see: https://dspace.jaist.ac.jp/dspace/bitstream/10119/4784/1/IS-RR-2004-010.pdf
package accrual

import "time"

// Detector represents a φ-failure detector.
type Detector struct {
	t float64
	w Window
}

// Window is the interface that wraps interval storages.
type Window interface {
	Record()
	Last() time.Time
	Distribution() []int64
}

// NewDetector creates a new detector with the given threshold
// and a window.
func NewDetector(threshold float64, w Window) *Detector {
	return &Detector{
		t: threshold,
		w: w,
	}
}

// Heartbeat records a heartbeat.
func (d *Detector) Heartbeat() {
	d.w.Record()
}

// Failed returns if there is a potential failure.
func (d *Detector) Failed() bool {
	v := time.Now().Sub(d.w.Last())
	φ := phi(float64(v), d.w.Distribution())
	return φ >= d.t
}
