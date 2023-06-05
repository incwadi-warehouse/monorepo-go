package validation

import "testing"

func TestVar(t *testing.T) {
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
				"products",
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
