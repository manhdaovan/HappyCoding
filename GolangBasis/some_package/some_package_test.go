package somepackage

import (
	"testing"
)

func TestExportedFunc(t *testing.T) {
	type args struct {
		a int
		b uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExportedFunc(tt.args.a, tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("ExportedFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
