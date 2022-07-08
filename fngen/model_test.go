package fngen

import (
	"context"
	"reflect"
	"testing"
)

func TestMonad_Collect(t *testing.T) {
	type fields struct {
		key     []string
		Context context.Context
		Di      []Monad
	}
	tests := []struct {
		name   string
		fields fields
		want   Monad
	}{
		{
			"Monad Collect init",
			fields{
				key:     nil,
				Context: nil,
				Di:      nil,
			},
			Monad{},
		},
		{
			"Monad Collect",
			fields{
				key:     nil,
				Context: nil,
				Di: []Monad{
					{
						key:     []string{"A", "B"},
						Context: nil,
						Di:      nil,
					},
					{
						key:     []string{"C", "D"},
						Context: nil,
						Di:      nil,
					},
				},
			},
			Monad{
				key: []string{"A", "B", "C", "D"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Monad{
				key:     tt.fields.key,
				Context: tt.fields.Context,
				Di:      tt.fields.Di,
			}
			if got := m.Collect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nCollect() = %v, \nwant %v", got, tt.want)
			}
		})
	}
}

func TestMonad_Map(t *testing.T) {
	type fields struct {
		key     []string
		Context context.Context
		Di      []Monad
	}
	tests := []struct {
		name   string
		fields fields
		want   Monad
	}{
		{
			"Monad Map init",
			fields{
				key:     nil,
				Context: nil,
				Di:      nil,
			},
			Monad{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Monad{
				key:     tt.fields.key,
				Context: tt.fields.Context,
				Di:      tt.fields.Di,
			}
			if got := m.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonad_Reduce(t *testing.T) {
	type fields struct {
		key     []string
		Context context.Context
		Di      []Monad
	}
	tests := []struct {
		name   string
		fields fields
		want   Monad
	}{
		{
			"Monad Reduce init",
			fields{
				key:     nil,
				Context: nil,
				Di:      nil,
			},
			Monad{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Monad{
				key:     tt.fields.key,
				Context: tt.fields.Context,
				Di:      tt.fields.Di,
			}
			if got := m.Reduce(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
