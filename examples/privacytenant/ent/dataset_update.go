// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/examples/privacytenant/ent/dataset"
	"github.com/facebook/ent/examples/privacytenant/ent/predicate"
	"github.com/facebook/ent/examples/privacytenant/ent/tenant"
	"github.com/facebook/ent/schema/field"
)

// DatasetUpdate is the builder for updating Dataset entities.
type DatasetUpdate struct {
	config
	hooks    []Hook
	mutation *DatasetMutation
}

// Where adds a new predicate for the DatasetUpdate builder.
func (du *DatasetUpdate) Where(ps ...predicate.Dataset) *DatasetUpdate {
	du.mutation.predicates = append(du.mutation.predicates, ps...)
	return du
}

// SetName sets the "name" field.
func (du *DatasetUpdate) SetName(s string) *DatasetUpdate {
	du.mutation.SetName(s)
	return du
}

// SetTenantID sets the "tenant" edge to the Tenant entity by ID.
func (du *DatasetUpdate) SetTenantID(id int) *DatasetUpdate {
	du.mutation.SetTenantID(id)
	return du
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (du *DatasetUpdate) SetTenant(t *Tenant) *DatasetUpdate {
	return du.SetTenantID(t.ID)
}

// Mutation returns the DatasetMutation object of the builder.
func (du *DatasetUpdate) Mutation() *DatasetMutation {
	return du.mutation
}

// ClearTenant clears the "tenant" edge to the Tenant entity.
func (du *DatasetUpdate) ClearTenant() *DatasetUpdate {
	du.mutation.ClearTenant()
	return du
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DatasetUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		if err = du.check(); err != nil {
			return 0, err
		}
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DatasetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = du.check(); err != nil {
				return 0, err
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DatasetUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DatasetUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DatasetUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DatasetUpdate) check() error {
	if _, ok := du.mutation.TenantID(); du.mutation.TenantCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"tenant\"")
	}
	return nil
}

func (du *DatasetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dataset.Table,
			Columns: dataset.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dataset.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dataset.FieldName,
		})
	}
	if du.mutation.TenantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dataset.TenantTable,
			Columns: []string{dataset.TenantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tenant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dataset.TenantTable,
			Columns: []string{dataset.TenantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tenant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dataset.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DatasetUpdateOne is the builder for updating a single Dataset entity.
type DatasetUpdateOne struct {
	config
	hooks    []Hook
	mutation *DatasetMutation
}

// SetName sets the "name" field.
func (duo *DatasetUpdateOne) SetName(s string) *DatasetUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetTenantID sets the "tenant" edge to the Tenant entity by ID.
func (duo *DatasetUpdateOne) SetTenantID(id int) *DatasetUpdateOne {
	duo.mutation.SetTenantID(id)
	return duo
}

// SetTenant sets the "tenant" edge to the Tenant entity.
func (duo *DatasetUpdateOne) SetTenant(t *Tenant) *DatasetUpdateOne {
	return duo.SetTenantID(t.ID)
}

// Mutation returns the DatasetMutation object of the builder.
func (duo *DatasetUpdateOne) Mutation() *DatasetMutation {
	return duo.mutation
}

// ClearTenant clears the "tenant" edge to the Tenant entity.
func (duo *DatasetUpdateOne) ClearTenant() *DatasetUpdateOne {
	duo.mutation.ClearTenant()
	return duo
}

// Save executes the query and returns the updated Dataset entity.
func (duo *DatasetUpdateOne) Save(ctx context.Context) (*Dataset, error) {
	var (
		err  error
		node *Dataset
	)
	if len(duo.hooks) == 0 {
		if err = duo.check(); err != nil {
			return nil, err
		}
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DatasetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duo.check(); err != nil {
				return nil, err
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DatasetUpdateOne) SaveX(ctx context.Context) *Dataset {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DatasetUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DatasetUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DatasetUpdateOne) check() error {
	if _, ok := duo.mutation.TenantID(); duo.mutation.TenantCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"tenant\"")
	}
	return nil
}

func (duo *DatasetUpdateOne) sqlSave(ctx context.Context) (_node *Dataset, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dataset.Table,
			Columns: dataset.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dataset.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Dataset.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dataset.FieldName,
		})
	}
	if duo.mutation.TenantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dataset.TenantTable,
			Columns: []string{dataset.TenantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tenant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.TenantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   dataset.TenantTable,
			Columns: []string{dataset.TenantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tenant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Dataset{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dataset.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
