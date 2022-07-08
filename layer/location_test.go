package layer

import "testing"

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
