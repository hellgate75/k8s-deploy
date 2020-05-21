package model

import "strings"

type Request struct {
	Resource ResourceType	`yaml:"resourceType" json:"resourceType" xml:"resource-type"`
	Action	Action			`yaml:"action" json:"action" xml:"action"`
	Payload	interface{}		`yaml:"payload" json:"payload" xml:"payload"`
	Async	bool			`yaml:"async,omitempty" json:"async,omitempty" xml:"async,omitempty"`
	Filters []Filter		`yaml:"filters,omitempty" json:"filters,omitempty" xml:"filter,omitempty"`
	Queries	[]Query			`yaml:"queries,omitempty" json:"queries,omitempty" xml:"query,omitempty"`
}

type Filter struct {
	Field	Field		`yaml:"field" json:"async" xml:"async"`
	Partial	bool		`yaml:"partialMatch" json:"partialMatch" xml:"partial-match"`
	Value	interface{}	`yaml:"value" json:"value" xml:"value"`
}

type Action string

const (
	GetResoource    Action = "GET"
	AddResoource    Action = "ADD"
	UpdateResoource Action = "UPDATE"
	DeleteResoource Action = "DELETE"
)

func (a Action) Equals(act Action) bool {
	return string(act) != "" && strings.ToUpper(string(act)) == strings.ToUpper(string(a))
}

func (a Action) Same(act string) bool {
	return act != "" && strings.ToUpper(act) == strings.ToUpper(string(a))
}

func (a Action) String(act string) string {
	return strings.ToUpper(string(a))
}

type Field string

func (f Field) Equals(field Field) bool {
	return string(field) != "" && strings.ToUpper(string(field)) == strings.ToUpper(string(f))
}

