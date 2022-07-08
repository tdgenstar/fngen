package layer

import (
	"reflect"
	"testing"
)

func TestLocation_String(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.

		{
			"Location > String, success",
			fields{
				X: 100.54,
				Y: 200.123,
			},
			"100.540000,200.123000",
			false,
		},
		{
			"Location > String, fail : nil error",
			fields{},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Location{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			got, err := r.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("String() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocation_Draw(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Location
		wantErr bool
	}{
		{
			"draw : fail nil",
			fields{
				X: 0,
				Y: 0,
			},
			[]Location{},
			true,
		},
		{
			"draw : success",
			fields{
				X: 123.123,
				Y: 12.1234,
			},
			[]Location{
				{123.123000, 12.1234},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Location{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			got, err := r.Draw()
			if (err != nil) != tt.wantErr {
				t.Errorf("Draw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Draw() got = %v, want %v", got, tt.want)
			}
		})
	}
}
