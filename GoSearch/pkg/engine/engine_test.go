package engine

import (
	"go-dev-v3/GoSearch/pkg/crawler"
	"reflect"
	"testing"
)

func TestService_Search(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		s    *Service
		args args
		want []crawler.Document
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Search(tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
