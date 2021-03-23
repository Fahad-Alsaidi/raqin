package model

import "os"

type Book struct {
	Name   string
	Author string
	File   os.File
	Path   string
}
