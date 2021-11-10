package fop

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	copyDst := "./testdata/CopyTree"
	if err := os.RemoveAll(copyDst); err != nil {
		return
	}
	if err := os.Mkdir(copyDst, 0777); err != nil {
		return
	}
	m.Run()
}

func TestParentDir(t *testing.T) {
	type args struct {
		path string
	}
	dst := filepath.Clean("./testdata/ParentDir/hoge/")
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"with trailing separator", args{"./testdata/ParentDir/hoge/fuga/"}, dst, false},
		{"without trailing separator", args{"./testdata/ParentDir/hoge/fuga"}, dst, false},
		{"file with extension", args{"./testdata/ParentDir/hoge/piyo.txt"}, dst, false},
		{"file without extension", args{"./testdata/ParentDir/hoge/piyo"}, dst, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParentDir(tt.args.path)
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

func TestCopyTree(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"without trailing slash",
			args{"./testdata/ParentDir", "./testdata/CopyTree"},
			false,
		},
		{
			"with trailing slash to dst",
			args{"./testdata/ParentDir", "./testdata/CopyTree/"},
			false,
		},
		{
			"with trailing slash to src",
			args{"./testdata/ParentDir/", "./testdata/CopyTree"},
			false,
		},
		{
			"with file",
			args{"./testdata/ParentDir/hoge/piyo.txt", "./testdata/CopyTree"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CopyTree(tt.args.src, tt.args.dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("CopyTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
