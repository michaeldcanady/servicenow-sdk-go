package main

type QueryParam struct {
	Name        string // e.g. "sysparm_limit"
	Type        string // e.g. "int", "string"
	Description string // optional
	Required    bool
}
