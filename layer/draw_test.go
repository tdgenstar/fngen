package layer

import "testing"

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
