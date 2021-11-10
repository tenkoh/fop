// Copyright: tenkoh

/*
fop: package for operating files.

This package includes some utility functions to operate files and directories.
*/
package fop

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

var (
	ErrFailedCreatingDirectory = errors.New("failed to create a directory")
	ErrInvalidDestination      = errors.New("destination must be a directory")
)

// ParentDir returns the parent directory's path in absolute type.
//
// The trailing separator is ignored.
// For example, when "foo" is a directory, all inputs return the same output "/root".
// /root/foo, /root/foo/, /root/bar.txt
//
// Output is cleaned by filepath.Clean.
func ParentDir(path string) (string, error) {
	// remove trailing slash, ./ and so on.
	path = filepath.Clean(path)
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		p, _ := filepath.Split(path)
		p = filepath.Clean(p)
		return p, nil
	}
	return filepath.Dir(path), nil
}

// CopyTree copies files keeping directory tree.
//
// The operation is similar to "cp -r" in linux, "xcopy" in windows.
// Empty directories are copied too.
func CopyTree(src, dst string) error {
	// when src is a file, do copyFile then return.
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		savepath := filepath.Join(dst, filepath.Base(src))
		if err := copyFile(src, savepath); err != nil {
			return err
		}
		return nil
	}

	err = filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		savepath := filepath.Join(dst, rel)
		if !d.IsDir() {
			if err := copyFile(path, savepath); err != nil {
				return err
			}
			return nil
		}
		_, err = os.Stat(savepath)
		if err != nil {
			if err := os.MkdirAll(savepath, 0777); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// src, dst must be filepath, not directory.
func copyFile(src, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	dir := filepath.Dir(dst)
	_, err = os.Stat(dir)
	if err != nil {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}
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
