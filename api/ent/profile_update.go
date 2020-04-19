// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/confus1on/UKM/ent/predicate"
	"github.com/confus1on/UKM/ent/profile"
	"github.com/confus1on/UKM/ent/ukm"
	"github.com/confus1on/UKM/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ProfileUpdate is the builder for updating Profile entities.
type ProfileUpdate struct {
	config
	hooks      []Hook
	mutation   *ProfileMutation
	predicates []predicate.Profile
}

// Where adds a new predicate for the builder.
func (pu *ProfileUpdate) Where(ps ...predicate.Profile) *ProfileUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetFirstName sets the firstName field.
func (pu *ProfileUpdate) SetFirstName(s string) *ProfileUpdate {
	pu.mutation.SetFirstName(s)
	return pu
}

// SetLastName sets the lastName field.
func (pu *ProfileUpdate) SetLastName(s string) *ProfileUpdate {
	pu.mutation.SetLastName(s)
	return pu
}

// SetJk sets the jk field.
func (pu *ProfileUpdate) SetJk(pr profile.Jk) *ProfileUpdate {
	pu.mutation.SetJk(pr)
	return pu
}

// SetAddress sets the address field.
func (pu *ProfileUpdate) SetAddress(s string) *ProfileUpdate {
	pu.mutation.SetAddress(s)
	return pu
}

// SetNillableAddress sets the address field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableAddress(s *string) *ProfileUpdate {
	if s != nil {
		pu.SetAddress(*s)
	}
	return pu
}

// ClearAddress clears the value of address.
func (pu *ProfileUpdate) ClearAddress() *ProfileUpdate {
	pu.mutation.ClearAddress()
	return pu
}

// SetBirthDate sets the birth_date field.
func (pu *ProfileUpdate) SetBirthDate(s string) *ProfileUpdate {
	pu.mutation.SetBirthDate(s)
	return pu
}

// SetNillableBirthDate sets the birth_date field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableBirthDate(s *string) *ProfileUpdate {
	if s != nil {
		pu.SetBirthDate(*s)
	}
	return pu
}

// ClearBirthDate clears the value of birth_date.
func (pu *ProfileUpdate) ClearBirthDate() *ProfileUpdate {
	pu.mutation.ClearBirthDate()
	return pu
}

// SetYearGeneration sets the year_generation field.
func (pu *ProfileUpdate) SetYearGeneration(s string) *ProfileUpdate {
	pu.mutation.SetYearGeneration(s)
	return pu
}

// SetPhone sets the phone field.
func (pu *ProfileUpdate) SetPhone(s string) *ProfileUpdate {
	pu.mutation.SetPhone(s)
	return pu
}

// SetStatus sets the status field.
func (pu *ProfileUpdate) SetStatus(b bool) *ProfileUpdate {
	pu.mutation.SetStatus(b)
	return pu
}

// SetNillableStatus sets the status field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableStatus(b *bool) *ProfileUpdate {
	if b != nil {
		pu.SetStatus(*b)
	}
	return pu
}

// SetImgProfile sets the img_profile field.
func (pu *ProfileUpdate) SetImgProfile(s string) *ProfileUpdate {
	pu.mutation.SetImgProfile(s)
	return pu
}

// SetNillableImgProfile sets the img_profile field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableImgProfile(s *string) *ProfileUpdate {
	if s != nil {
		pu.SetImgProfile(*s)
	}
	return pu
}

// SetCreatedAt sets the created_at field.
func (pu *ProfileUpdate) SetCreatedAt(t time.Time) *ProfileUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableCreatedAt(t *time.Time) *ProfileUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the updated_at field.
func (pu *ProfileUpdate) SetUpdatedAt(t time.Time) *ProfileUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// AddOwnerIDs adds the owner edge to User by ids.
func (pu *ProfileUpdate) AddOwnerIDs(ids ...int) *ProfileUpdate {
	pu.mutation.AddOwnerIDs(ids...)
	return pu
}

// AddOwner adds the owner edges to User.
func (pu *ProfileUpdate) AddOwner(u ...*User) *ProfileUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.AddOwnerIDs(ids...)
}

// AddUkmIDs adds the ukm edge to Ukm by ids.
func (pu *ProfileUpdate) AddUkmIDs(ids ...int) *ProfileUpdate {
	pu.mutation.AddUkmIDs(ids...)
	return pu
}

// AddUkm adds the ukm edges to Ukm.
func (pu *ProfileUpdate) AddUkm(u ...*Ukm) *ProfileUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.AddUkmIDs(ids...)
}

// RemoveOwnerIDs removes the owner edge to User by ids.
func (pu *ProfileUpdate) RemoveOwnerIDs(ids ...int) *ProfileUpdate {
	pu.mutation.RemoveOwnerIDs(ids...)
	return pu
}

// RemoveOwner removes owner edges to User.
func (pu *ProfileUpdate) RemoveOwner(u ...*User) *ProfileUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.RemoveOwnerIDs(ids...)
}

// RemoveUkmIDs removes the ukm edge to Ukm by ids.
func (pu *ProfileUpdate) RemoveUkmIDs(ids ...int) *ProfileUpdate {
	pu.mutation.RemoveUkmIDs(ids...)
	return pu
}

// RemoveUkm removes ukm edges to Ukm.
func (pu *ProfileUpdate) RemoveUkm(u ...*Ukm) *ProfileUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.RemoveUkmIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *ProfileUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.FirstName(); ok {
		if err := profile.FirstNameValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"firstName\": %v", err)
		}
	}
	if v, ok := pu.mutation.LastName(); ok {
		if err := profile.LastNameValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"lastName\": %v", err)
		}
	}
	if v, ok := pu.mutation.Jk(); ok {
		if err := profile.JkValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"jk\": %v", err)
		}
	}
	if v, ok := pu.mutation.Phone(); ok {
		if err := profile.PhoneValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"phone\": %v", err)
		}
	}
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := profile.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProfileUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProfileUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   profile.Table,
			Columns: profile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: profile.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldFirstName,
		})
	}
	if value, ok := pu.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldLastName,
		})
	}
	if value, ok := pu.mutation.Jk(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: profile.FieldJk,
		})
	}
	if value, ok := pu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldAddress,
		})
	}
	if pu.mutation.AddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: profile.FieldAddress,
		})
	}
	if value, ok := pu.mutation.BirthDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldBirthDate,
		})
	}
	if pu.mutation.BirthDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: profile.FieldBirthDate,
		})
	}
	if value, ok := pu.mutation.YearGeneration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldYearGeneration,
		})
	}
	if value, ok := pu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldPhone,
		})
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: profile.FieldStatus,
		})
	}
	if value, ok := pu.mutation.ImgProfile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldImgProfile,
		})
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldCreatedAt,
		})
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldUpdatedAt,
		})
	}
	if nodes := pu.mutation.RemovedOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   profile.OwnerTable,
			Columns: []string{profile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   profile.OwnerTable,
			Columns: []string{profile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := pu.mutation.RemovedUkmIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.UkmTable,
			Columns: profile.UkmPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ukm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UkmIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.UkmTable,
			Columns: profile.UkmPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ukm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProfileUpdateOne is the builder for updating a single Profile entity.
type ProfileUpdateOne struct {
	config
	hooks    []Hook
	mutation *ProfileMutation
}

// SetFirstName sets the firstName field.
func (puo *ProfileUpdateOne) SetFirstName(s string) *ProfileUpdateOne {
	puo.mutation.SetFirstName(s)
	return puo
}

// SetLastName sets the lastName field.
func (puo *ProfileUpdateOne) SetLastName(s string) *ProfileUpdateOne {
	puo.mutation.SetLastName(s)
	return puo
}

// SetJk sets the jk field.
func (puo *ProfileUpdateOne) SetJk(pr profile.Jk) *ProfileUpdateOne {
	puo.mutation.SetJk(pr)
	return puo
}

// SetAddress sets the address field.
func (puo *ProfileUpdateOne) SetAddress(s string) *ProfileUpdateOne {
	puo.mutation.SetAddress(s)
	return puo
}

// SetNillableAddress sets the address field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableAddress(s *string) *ProfileUpdateOne {
	if s != nil {
		puo.SetAddress(*s)
	}
	return puo
}

// ClearAddress clears the value of address.
func (puo *ProfileUpdateOne) ClearAddress() *ProfileUpdateOne {
	puo.mutation.ClearAddress()
	return puo
}

// SetBirthDate sets the birth_date field.
func (puo *ProfileUpdateOne) SetBirthDate(s string) *ProfileUpdateOne {
	puo.mutation.SetBirthDate(s)
	return puo
}

// SetNillableBirthDate sets the birth_date field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableBirthDate(s *string) *ProfileUpdateOne {
	if s != nil {
		puo.SetBirthDate(*s)
	}
	return puo
}

// ClearBirthDate clears the value of birth_date.
func (puo *ProfileUpdateOne) ClearBirthDate() *ProfileUpdateOne {
	puo.mutation.ClearBirthDate()
	return puo
}

// SetYearGeneration sets the year_generation field.
func (puo *ProfileUpdateOne) SetYearGeneration(s string) *ProfileUpdateOne {
	puo.mutation.SetYearGeneration(s)
	return puo
}

// SetPhone sets the phone field.
func (puo *ProfileUpdateOne) SetPhone(s string) *ProfileUpdateOne {
	puo.mutation.SetPhone(s)
	return puo
}

// SetStatus sets the status field.
func (puo *ProfileUpdateOne) SetStatus(b bool) *ProfileUpdateOne {
	puo.mutation.SetStatus(b)
	return puo
}

// SetNillableStatus sets the status field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableStatus(b *bool) *ProfileUpdateOne {
	if b != nil {
		puo.SetStatus(*b)
	}
	return puo
}

// SetImgProfile sets the img_profile field.
func (puo *ProfileUpdateOne) SetImgProfile(s string) *ProfileUpdateOne {
	puo.mutation.SetImgProfile(s)
	return puo
}

// SetNillableImgProfile sets the img_profile field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableImgProfile(s *string) *ProfileUpdateOne {
	if s != nil {
		puo.SetImgProfile(*s)
	}
	return puo
}

// SetCreatedAt sets the created_at field.
func (puo *ProfileUpdateOne) SetCreatedAt(t time.Time) *ProfileUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableCreatedAt(t *time.Time) *ProfileUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the updated_at field.
func (puo *ProfileUpdateOne) SetUpdatedAt(t time.Time) *ProfileUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// AddOwnerIDs adds the owner edge to User by ids.
func (puo *ProfileUpdateOne) AddOwnerIDs(ids ...int) *ProfileUpdateOne {
	puo.mutation.AddOwnerIDs(ids...)
	return puo
}

// AddOwner adds the owner edges to User.
func (puo *ProfileUpdateOne) AddOwner(u ...*User) *ProfileUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.AddOwnerIDs(ids...)
}

// AddUkmIDs adds the ukm edge to Ukm by ids.
func (puo *ProfileUpdateOne) AddUkmIDs(ids ...int) *ProfileUpdateOne {
	puo.mutation.AddUkmIDs(ids...)
	return puo
}

// AddUkm adds the ukm edges to Ukm.
func (puo *ProfileUpdateOne) AddUkm(u ...*Ukm) *ProfileUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.AddUkmIDs(ids...)
}

// RemoveOwnerIDs removes the owner edge to User by ids.
func (puo *ProfileUpdateOne) RemoveOwnerIDs(ids ...int) *ProfileUpdateOne {
	puo.mutation.RemoveOwnerIDs(ids...)
	return puo
}

// RemoveOwner removes owner edges to User.
func (puo *ProfileUpdateOne) RemoveOwner(u ...*User) *ProfileUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.RemoveOwnerIDs(ids...)
}

// RemoveUkmIDs removes the ukm edge to Ukm by ids.
func (puo *ProfileUpdateOne) RemoveUkmIDs(ids ...int) *ProfileUpdateOne {
	puo.mutation.RemoveUkmIDs(ids...)
	return puo
}

// RemoveUkm removes ukm edges to Ukm.
func (puo *ProfileUpdateOne) RemoveUkm(u ...*Ukm) *ProfileUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.RemoveUkmIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *ProfileUpdateOne) Save(ctx context.Context) (*Profile, error) {
	if v, ok := puo.mutation.FirstName(); ok {
		if err := profile.FirstNameValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"firstName\": %v", err)
		}
	}
	if v, ok := puo.mutation.LastName(); ok {
		if err := profile.LastNameValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"lastName\": %v", err)
		}
	}
	if v, ok := puo.mutation.Jk(); ok {
		if err := profile.JkValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"jk\": %v", err)
		}
	}
	if v, ok := puo.mutation.Phone(); ok {
		if err := profile.PhoneValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"phone\": %v", err)
		}
	}
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := profile.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}

	var (
		err  error
		node *Profile
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProfileUpdateOne) SaveX(ctx context.Context) *Profile {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *ProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProfileUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProfileUpdateOne) sqlSave(ctx context.Context) (pr *Profile, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   profile.Table,
			Columns: profile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: profile.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Profile.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldFirstName,
		})
	}
	if value, ok := puo.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldLastName,
		})
	}
	if value, ok := puo.mutation.Jk(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: profile.FieldJk,
		})
	}
	if value, ok := puo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldAddress,
		})
	}
	if puo.mutation.AddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: profile.FieldAddress,
		})
	}
	if value, ok := puo.mutation.BirthDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldBirthDate,
		})
	}
	if puo.mutation.BirthDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: profile.FieldBirthDate,
		})
	}
	if value, ok := puo.mutation.YearGeneration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldYearGeneration,
		})
	}
	if value, ok := puo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldPhone,
		})
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: profile.FieldStatus,
		})
	}
	if value, ok := puo.mutation.ImgProfile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldImgProfile,
		})
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldCreatedAt,
		})
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldUpdatedAt,
		})
	}
	if nodes := puo.mutation.RemovedOwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   profile.OwnerTable,
			Columns: []string{profile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   profile.OwnerTable,
			Columns: []string{profile.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := puo.mutation.RemovedUkmIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.UkmTable,
			Columns: profile.UkmPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ukm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UkmIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.UkmTable,
			Columns: profile.UkmPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ukm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pr = &Profile{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}
