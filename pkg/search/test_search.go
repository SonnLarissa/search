package search

import (
	"context"
	"reflect"
	"testing"
)

func TestAll_success(t *testing.T) {
	ch := All(context.Background(), "test", []string{"test.txt"})

	_, err := <-ch

	if !err {
		t.Error(err)
	}
}

func TestAll(t *testing.T) {
	type args struct {
		ctx    context.Context
		phrase string
		files  []string
	}
	tests := []struct {
		name string
		args args
		want <-chan []Result
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(tt.args.ctx, tt.args.phrase, tt.args.files); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}
