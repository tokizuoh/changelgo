package main

type Item struct {
	Link string `yaml:"link"`
}

type Target struct {
	Items []Item `yaml:"items"`
}
