package main

type Node struct {
	CommittedDate string `json:"committedDate"`
	Message       string `json:"message"`
	URL           string `json:"url"`
}

type Edge struct {
	Node Node `json:"node"`
}

type History struct {
	Edges []Edge `json:"edges"`
}

type Object struct {
	History History `json:"history"`
}

type Repository struct {
	Object Object `json:"object"`
}

type Data struct {
	Repository Repository `json:"repository"`
}

type Response struct {
	Data Data `json:"data"`
}
