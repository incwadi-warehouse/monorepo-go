package validation

import (
	"testing"
    "os"
)

type Product struct {
	Name string `json:"name" validate:"required"`
}

func TestVar(t *testing.T) {
    os.Setenv("BRANCHES", "1")
    os.Setenv("INDEXES", "products")
	type args struct {
		name        interface{}
		constraints string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test products",
			args{
				"products_1",
				"indexName",
			},
			false,
		},
		{
			"test books",
			args{
				"books",
				"indexName",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Var(tt.args.name, tt.args.constraints); (err != nil) != tt.wantErr {
				t.Errorf("Var() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStruct(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test product",
			args{Product{"test"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Struct(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("Struct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
