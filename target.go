package main

type FileList struct {
	Files []File `yaml:"items"`
}

type File struct {
	Link   string `yaml:"link"`
	Owner  string `yaml:"owner"`
	Repo   string `yaml:"repo"`
	File   string `yaml:"file"`
	Branch string `yaml:"branch"`
}
