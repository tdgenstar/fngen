package blockly

import (
	"context"
	"fn/fngen"
	"fn/layer"
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
			&Block{
				Square:  layer.Square{},
				Context: context.Background(),
			},
			false,
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

func TestBlock_Concat(t *testing.T) {
	type fields struct {
		Square layer.Square
	}
	tests := []struct {
		name   string
		fields fields
		want   fngen.Monad
	}{
		// TODO: Add test cases.
		// TODO: 연결 시키는데 성공
		{
			name: "block > concat success",
			fields: fields{
				Square: layer.Square{
					Start: layer.Point{X: 1, Y: 2},
					End:   layer.Point{X: 3, Y: 4},
				},
			},
			want: fngen.Monad{
				Context: nil,
				Di:      nil,
				ConcatFn: fngen.ConcatFn(Block{
					Square: layer.Square{
						Start: layer.Point{X: 1, Y: 2},
						End:   layer.Point{X: 3, Y: 4},
					},
				}),
			},
		},
		// TODO: 연결 시키는데 실패 - validate - 연결이 불가능한 블록
		// TODO: 연결 시키는데 실패 - validate - 순환참조
		// TODO: 연결 시키는데 실패 - validate - 없는 것을 연결
		// TODO: 연결 시키는데 실패 - validate - 자기 자신을 연결
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Block{
				Square: tt.fields.Square,
			}
			if got := r.Concat(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}
