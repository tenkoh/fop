package fop

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ParentDir returns the parent directory's path.
// See examples in the test file.
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
