package engine

import (
	"github.com/juju/errors"
	"strings"
)

type FlowKey struct {
	Components []string
}

func (t FlowKey) String() string {
	return strings.Join(t.Components, ".")
}

func (t FlowKey) ShortString() string {
	return strings.Join(t.Components[1:], ".")
}

func (t FlowKey) Parent() (*FlowKey, error) {
	if len(t.Components) > 1 {
		return &FlowKey{Components: t.Components[:len(t.Components)-1]}, nil
	} else {
		return nil, errors.Errorf("FlowKey %v doesn't have a parent", t)
	}
}