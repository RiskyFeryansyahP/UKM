// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/confus1on/UKM/ent/predicate"
	"github.com/confus1on/UKM/ent/roleukm"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RoleUKMDelete is the builder for deleting a RoleUKM entity.
type RoleUKMDelete struct {
	config
	hooks      []Hook
	mutation   *RoleUKMMutation
	predicates []predicate.RoleUKM
}

// Where adds a new predicate to the delete builder.
func (rud *RoleUKMDelete) Where(ps ...predicate.RoleUKM) *RoleUKMDelete {
	rud.predicates = append(rud.predicates, ps...)
	return rud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rud *RoleUKMDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rud.hooks) == 0 {
		affected, err = rud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleUKMMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rud.mutation = mutation
			affected, err = rud.sqlExec(ctx)
			return affected, err
		})
		for i := len(rud.hooks) - 1; i >= 0; i-- {
			mut = rud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rud *RoleUKMDelete) ExecX(ctx context.Context) int {
	n, err := rud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rud *RoleUKMDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: roleukm.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roleukm.FieldID,
			},
		},
	}
	if ps := rud.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rud.driver, _spec)
}

// RoleUKMDeleteOne is the builder for deleting a single RoleUKM entity.
type RoleUKMDeleteOne struct {
	rud *RoleUKMDelete
}

// Exec executes the deletion query.
func (rudo *RoleUKMDeleteOne) Exec(ctx context.Context) error {
	n, err := rudo.rud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{roleukm.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rudo *RoleUKMDeleteOne) ExecX(ctx context.Context) {
	rudo.rud.ExecX(ctx)
}
