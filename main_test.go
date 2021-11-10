package fop_test

import (
	"path/filepath"
	"testing"

	"github.com/tenkoh/fop"
)

func TestParentDir(t *testing.T) {
	type args struct {
		path string
	}
	dstAbs, _ := filepath.Abs("./testdata/ParentDir/hoge/")
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"with trailing separator", args{"./testdata/ParentDir/hoge/fuga/"}, dstAbs, false},
		{"without trailing separator", args{"./testdata/ParentDir/hoge/fuga"}, dstAbs, false},
		{"file with extension", args{"./testdata/ParentDir/hoge/piyo.txt"}, dstAbs, false},
		{"file without extension", args{"./testdata/ParentDir/hoge/piyo"}, dstAbs, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fop.ParentDir(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParentDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParentDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
