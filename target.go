package main

type Target struct {
	Items []Item `yaml:"items"`
}

type Item struct {
	Link   string `yaml:"link"`
	Owner  string `yaml:"owner"`
	Repo   string `yaml:"repo"`
	File   string `yaml:"file"`
	Branch string `yaml:"branch"`
}
