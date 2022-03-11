package pool

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestMaxQueueSize(t *testing.T) {
	type args struct {
		size uint32
	}
	tests := []struct {
		name string
		args args
		want *pool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxQueueSize(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxQueueSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pool_MaxQueueSize(t *testing.T) {
	type fields struct {
		Procs         uint8
		MaxGoRoutines uint32
		queueSize     uint32
		workChan      chan work
		wg            sync.WaitGroup
		ttl           time.Duration
	}
	type args struct {
		size uint32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *pool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pool{
				Procs:         tt.fields.Procs,
				MaxGoRoutines: tt.fields.MaxGoRoutines,
				queueSize:     tt.fields.queueSize,
				workChan:      tt.fields.workChan,
				wg:            tt.fields.wg,
				ttl:           tt.fields.ttl,
			}
			if got := p.MaxQueueSize(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pool.MaxQueueSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
