package crcind

import (
	"reflect"
	"testing"

	"github.com/jhuygens/searcher-engine"
)

func TestSearcher_Search(t *testing.T) {
	type args struct {
		filter searcher.Filter
	}
	tests := []struct {
		name    string
		s       Searcher
		args    args
		want    []searcher.Item
		wantErr bool
	}{
		{
			name: "prueba inicial",
			s:    Searcher{},
			args: args{
				filter: searcher.Filter{
					Name:   []searcher.FieldValue{{Value: "john"}, {Value: "como"}, {Value: "mono"}},
					Artist: []searcher.FieldValue{{Value: "nirvana"}},
					Types:  []string{"people", "track"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Searcher{}
			got, err := s.Search(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Searcher.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Searcher.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
