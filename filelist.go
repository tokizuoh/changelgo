package main

type FileList struct {
	Files []File
}

type File struct {
	Link   string
	Owner  string
	Repo   string
	Name   string
	Branch string
}
