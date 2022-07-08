package blockly

import (
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name    string
		want    *Block
		wantErr bool
	}{
		{
			"block create",
			nil, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create()
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
