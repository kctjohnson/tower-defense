package components

import "ecstemplate/pkg/ecs"

const (
	Sample ecs.ComponentType = "sample"
)

type SampleComponent struct {
	ecs.Component
	Data string
}

func (c SampleComponent) GetType() ecs.ComponentType {
	return Sample
}

var ComponentTypes = []ecs.ComponentType{
	Sample,
}
