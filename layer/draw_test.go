package layer

import (
	"reflect"
	"testing"
)

func TestLine_Name(t *testing.T) {
	type fields struct {
		name  string
		Start Point
		End   Point
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			"line > name fail",
			fields{
				name:  "",
				Start: Point{1, 2},
				End:   Point{3, 4},
			},
			"",
			true,
		},
		{
			"line > name success",
			fields{
				name:  "line test",
				Start: Point{},
				End:   Point{},
			},
			"line test",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Line{
				name:  tt.fields.name,
				Start: tt.fields.Start,
				End:   tt.fields.End,
			}
			got, err := r.Name()
			if (err != nil) != tt.wantErr {
				t.Errorf("Name() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Name() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine_Map(t *testing.T) {
	type fields struct {
		name  string
		Start Point
		End   Point
	}
	tests := []struct {
		name   string
		fields fields
		want   []Line
	}{
		{
			"Line > Map() : empty name",
			fields{
				name:  "",
				Start: Point{},
				End:   Point{},
			},
			[]Line{},
		},
		{
			"Line > Map() : success",
			fields{
				name:  "test",
				Start: Point{1, 2},
				End:   Point{3, 4},
			},
			[]Line{
				{"test", Point{1, 2}, Point{3, 4}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Line{
				name:  tt.fields.name,
				Start: tt.fields.Start,
				End:   tt.fields.End,
			}
			if got := r.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine_LengthToPoint(t *testing.T) {
	type fields struct {
		name  string
		Start Point
		End   Point
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			"line > lengthToPoint : not impl",
			fields{
				name:  "",
				Start: Point{},
				End:   Point{},
			},
			0.0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Line{
				name:  tt.fields.name,
				Start: tt.fields.Start,
				End:   tt.fields.End,
			}
			got, err := r.LengthToPoint()
			if (err != nil) != tt.wantErr {
				t.Errorf("LengthToPoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LengthToPoint() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare_Name(t *testing.T) {
	type fields struct {
		name  string
		Start Point
		End   Point
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			"square > name : success",
			fields{
				name:  "test",
				Start: Point{1, 2},
				End:   Point{3, 4},
			},
			"test",
			false,
		},
		{
			"square > name : empty name",
			fields{
				name:  "",
				Start: Point{1, 2},
				End:   Point{3, 4},
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Square{
				name:  tt.fields.name,
				Start: tt.fields.Start,
				End:   tt.fields.End,
			}
			got, err := r.Name()
			if (err != nil) != tt.wantErr {
				t.Errorf("Name() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Name() got = %v, want %v", got, tt.want)
			}
		})
	}
}
