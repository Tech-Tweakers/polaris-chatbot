package logwrapper

import (
	uuid "github.com/satori/go.uuid"
)

type span struct {
	id     string
	parent *span
}

func createSpan(parent *span) *span {
	s := &span{
		id:     uuid.NewV4().String(),
		parent: parent,
	}
	return s
}
