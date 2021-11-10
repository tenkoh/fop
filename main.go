// Copyright: tenkoh

/*
fop: package for operating files.

This package includes some utility functions to operate files and directories.
*/
package fop

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrFailedCreatingDirectory = errors.New("failed to create a directory")
	ErrInvalidDestination      = errors.New("destination must be a directory")
)

// ParentDir returns the parent directory's path in absolute type.
//
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

// CopyTree copies files keeping directory tee, then returns paths of copied files.
//
// The operation is similar to "cp" in linux, "xcopy" in windows.
// Empty directories are ignored.
func CopyTree(src, dst string) ([]string, error) {
	// check dst is not a file. dst must be a directory.
	info, err := os.Stat(dst)
	if err != nil {
		if !info.IsDir() {
			return nil, ErrInvalidDestination
		}
	}
	// if src is a file, just copy it into dst
	info, err = os.Stat(src)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		savepath := filepath.Join(dst, filepath.Base(src))
		if err := copyFile(src, savepath); err != nil {
			return nil, err
		}
		return []string{savepath}, nil
	}

	// if src is a directory, find all files and copy all.
	files, err := WalkFiles(src)
	if err != nil {
		return nil, err
	}
	copied := []string{}
	for _, file := range files {
		rel, err := filepath.Rel(src, file)
		if err != nil {
			return nil, err
		}
		savepath := filepath.Join(dst, rel)
		if err := copyFile(file, savepath); err != nil {
			return nil, err
		}
		copied = append(copied, savepath)
	}
	return copied, nil
}

// src, dst must be filepath, not directory.
func copyFile(src, dst string) error {
	dir := filepath.Dir(dst)
	_, err := os.Stat(dir)
	if err != nil {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()
	g, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer g.Close()
	if _, err := io.Copy(g, f); err != nil {
		return err
	}
	return nil
}

// WalkFiles returns a file tree under "dir".
//
// Empty directories are ignored.
func WalkFiles(dir string) ([]string, error) {
	// when filepath is passed, return it.
	info, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		// using filepath.Join to fit output path style.
		return []string{filepath.Join(dir)}, nil
	}

	paths := []string{}
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		fullpath := filepath.Join(dir, file.Name())
		if !file.IsDir() {
			paths = append(paths, fullpath)
			continue
		}
		subpaths, err := WalkFiles(fullpath)
		if err != nil {
			return nil, err
		}
		paths = append(paths, subpaths...)
	}
	return paths, nil
}
