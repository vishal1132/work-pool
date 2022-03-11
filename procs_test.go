package pool

import (
	"runtime"
	"testing"
)

func TestSetMaxProcs(t *testing.T) {
	cpu := runtime.NumCPU()
	tests := []struct {
		name   string
		procs  uint8
		result uint8
	}{
		{
			name:   "0 procs",
			procs:  0,
			result: defaultProcs,
		},
		{
			name:   "procs more than the number of cpus",
			procs:  uint8(cpu + 1),
			result: uint8(cpu),
		},
		{
			name:   "procs less than the number of cpus",
			procs:  uint8(cpu - 1),
			result: uint8(cpu - 1),
		},
	}
	for _, v := range tests {
		t.Run("", func(t *testing.T) {
			SetMaxProcs(v.procs)
			if v.result != GetProcs() {
				t.Fatalf("expected %d, got %d", v.result, GetProcs())
			}
		})
	}
}
