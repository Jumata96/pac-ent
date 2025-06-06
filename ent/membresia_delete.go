// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jumata96/pac-ent/ent/membresia"
	"github.com/jumata96/pac-ent/ent/predicate"
)

// MembresiaDelete is the builder for deleting a Membresia entity.
type MembresiaDelete struct {
	config
	hooks    []Hook
	mutation *MembresiaMutation
}

// Where appends a list predicates to the MembresiaDelete builder.
func (md *MembresiaDelete) Where(ps ...predicate.Membresia) *MembresiaDelete {
	md.mutation.Where(ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MembresiaDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, md.sqlExec, md.mutation, md.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MembresiaDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MembresiaDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(membresia.Table, sqlgraph.NewFieldSpec(membresia.FieldID, field.TypeInt))
	if ps := md.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, md.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	md.mutation.done = true
	return affected, err
}

// MembresiaDeleteOne is the builder for deleting a single Membresia entity.
type MembresiaDeleteOne struct {
	md *MembresiaDelete
}

// Where appends a list predicates to the MembresiaDelete builder.
func (mdo *MembresiaDeleteOne) Where(ps ...predicate.Membresia) *MembresiaDeleteOne {
	mdo.md.mutation.Where(ps...)
	return mdo
}

// Exec executes the deletion query.
func (mdo *MembresiaDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{membresia.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MembresiaDeleteOne) ExecX(ctx context.Context) {
	if err := mdo.Exec(ctx); err != nil {
		panic(err)
	}
}
