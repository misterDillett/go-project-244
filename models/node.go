package models

type Node struct {
	Key      string
	Type     string
	OldValue interface{}
	NewValue interface{}
	Children []*Node
}
