package schema

import "entgo.io/ent"

// Entrenador holds the schema definition for the Entrenador entity.
type Entrenador struct {
	ent.Schema
}

// Fields of the Entrenador.
func (Entrenador) Fields() []ent.Field {
	return nil
}

// Edges of the Entrenador.
func (Entrenador) Edges() []ent.Edge {
	return nil
}
