// Copyright: tenkoh

/*
fop: package for operating files.

This package includes some utility functions to operate files and directories.
*/
package fop

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ParentDir returns the parent directory's path in absolute type.
// The trailing separator is ignored.
// For example, when "foo" is a directory, all inputs return the same output "/root/".
// /root/foo, /root/foo/, /root/bar.txt
func ParentDir(path string) (string, error) {
	p, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("invalid path; %w", err)
	}
	// remove the last separator
	if strings.HasSuffix(p, string(os.PathSeparator)) {
		p = filepath.Dir(p)
	}
	parent := filepath.Dir(p)
	return parent, nil
}
