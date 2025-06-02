package schema

import "entgo.io/ent"

// Rutina holds the schema definition for the Rutina entity.
type Rutina struct {
	ent.Schema
}

// Fields of the Rutina.
func (Rutina) Fields() []ent.Field {
	return nil
}

// Edges of the Rutina.
func (Rutina) Edges() []ent.Edge {
	return nil
}
