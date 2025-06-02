package schema

import "entgo.io/ent"

// Cliente holds the schema definition for the Cliente entity.
type Cliente struct {
	ent.Schema
}

// Fields of the Cliente.
func (Cliente) Fields() []ent.Field {
	return nil
}

// Edges of the Cliente.
func (Cliente) Edges() []ent.Edge {
	return nil
}
