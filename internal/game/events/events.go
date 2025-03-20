package events

import "ecstemplate/pkg/ecs"

const (
	Sample ecs.EventType = "sample"
)

type SampleEvent struct {
	Ent           ecs.Entity
	DataBeingSent string
}

func (e *SampleEvent) Type() ecs.EventType {
	return Sample
}

func (e *SampleEvent) Entity() ecs.Entity {
	return e.Ent
}

func (e *SampleEvent) Data() any {
	return map[string]any{"data": e.DataBeingSent}
}
