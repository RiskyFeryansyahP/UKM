// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/confus1on/UKM/ent/predicate"
	"github.com/confus1on/UKM/ent/profile"
	"github.com/confus1on/UKM/ent/profileukm"
	"github.com/confus1on/UKM/ent/roleukm"
	"github.com/confus1on/UKM/ent/ukm"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ProfileUKMQuery is the builder for querying ProfileUKM entities.
type ProfileUKMQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.ProfileUKM
	// eager-loading edges.
	withOwnerProfile *ProfileQuery
	withOwnerUkm     *UkmQuery
	withOwnerRole    *RoleUKMQuery
	withFKs          bool
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (puq *ProfileUKMQuery) Where(ps ...predicate.ProfileUKM) *ProfileUKMQuery {
	puq.predicates = append(puq.predicates, ps...)
	return puq
}

// Limit adds a limit step to the query.
func (puq *ProfileUKMQuery) Limit(limit int) *ProfileUKMQuery {
	puq.limit = &limit
	return puq
}

// Offset adds an offset step to the query.
func (puq *ProfileUKMQuery) Offset(offset int) *ProfileUKMQuery {
	puq.offset = &offset
	return puq
}

// Order adds an order step to the query.
func (puq *ProfileUKMQuery) Order(o ...Order) *ProfileUKMQuery {
	puq.order = append(puq.order, o...)
	return puq
}

// QueryOwnerProfile chains the current query on the owner_profile edge.
func (puq *ProfileUKMQuery) QueryOwnerProfile() *ProfileQuery {
	query := &ProfileQuery{config: puq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(profileukm.Table, profileukm.FieldID, puq.sqlQuery()),
		sqlgraph.To(profile.Table, profile.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, profileukm.OwnerProfileTable, profileukm.OwnerProfileColumn),
	)
	query.sql = sqlgraph.SetNeighbors(puq.driver.Dialect(), step)
	return query
}

// QueryOwnerUkm chains the current query on the owner_ukm edge.
func (puq *ProfileUKMQuery) QueryOwnerUkm() *UkmQuery {
	query := &UkmQuery{config: puq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(profileukm.Table, profileukm.FieldID, puq.sqlQuery()),
		sqlgraph.To(ukm.Table, ukm.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, profileukm.OwnerUkmTable, profileukm.OwnerUkmColumn),
	)
	query.sql = sqlgraph.SetNeighbors(puq.driver.Dialect(), step)
	return query
}

// QueryOwnerRole chains the current query on the owner_role edge.
func (puq *ProfileUKMQuery) QueryOwnerRole() *RoleUKMQuery {
	query := &RoleUKMQuery{config: puq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(profileukm.Table, profileukm.FieldID, puq.sqlQuery()),
		sqlgraph.To(roleukm.Table, roleukm.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, profileukm.OwnerRoleTable, profileukm.OwnerRoleColumn),
	)
	query.sql = sqlgraph.SetNeighbors(puq.driver.Dialect(), step)
	return query
}

// First returns the first ProfileUKM entity in the query. Returns *NotFoundError when no profileukm was found.
func (puq *ProfileUKMQuery) First(ctx context.Context) (*ProfileUKM, error) {
	pus, err := puq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(pus) == 0 {
		return nil, &NotFoundError{profileukm.Label}
	}
	return pus[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (puq *ProfileUKMQuery) FirstX(ctx context.Context) *ProfileUKM {
	pu, err := puq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return pu
}

// FirstID returns the first ProfileUKM id in the query. Returns *NotFoundError when no id was found.
func (puq *ProfileUKMQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = puq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{profileukm.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (puq *ProfileUKMQuery) FirstXID(ctx context.Context) int {
	id, err := puq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only ProfileUKM entity in the query, returns an error if not exactly one entity was returned.
func (puq *ProfileUKMQuery) Only(ctx context.Context) (*ProfileUKM, error) {
	pus, err := puq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(pus) {
	case 1:
		return pus[0], nil
	case 0:
		return nil, &NotFoundError{profileukm.Label}
	default:
		return nil, &NotSingularError{profileukm.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (puq *ProfileUKMQuery) OnlyX(ctx context.Context) *ProfileUKM {
	pu, err := puq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return pu
}

// OnlyID returns the only ProfileUKM id in the query, returns an error if not exactly one id was returned.
func (puq *ProfileUKMQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = puq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{profileukm.Label}
	default:
		err = &NotSingularError{profileukm.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (puq *ProfileUKMQuery) OnlyXID(ctx context.Context) int {
	id, err := puq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProfileUKMs.
func (puq *ProfileUKMQuery) All(ctx context.Context) ([]*ProfileUKM, error) {
	return puq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (puq *ProfileUKMQuery) AllX(ctx context.Context) []*ProfileUKM {
	pus, err := puq.All(ctx)
	if err != nil {
		panic(err)
	}
	return pus
}

// IDs executes the query and returns a list of ProfileUKM ids.
func (puq *ProfileUKMQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := puq.Select(profileukm.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (puq *ProfileUKMQuery) IDsX(ctx context.Context) []int {
	ids, err := puq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (puq *ProfileUKMQuery) Count(ctx context.Context) (int, error) {
	return puq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (puq *ProfileUKMQuery) CountX(ctx context.Context) int {
	count, err := puq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (puq *ProfileUKMQuery) Exist(ctx context.Context) (bool, error) {
	return puq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (puq *ProfileUKMQuery) ExistX(ctx context.Context) bool {
	exist, err := puq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (puq *ProfileUKMQuery) Clone() *ProfileUKMQuery {
	return &ProfileUKMQuery{
		config:     puq.config,
		limit:      puq.limit,
		offset:     puq.offset,
		order:      append([]Order{}, puq.order...),
		unique:     append([]string{}, puq.unique...),
		predicates: append([]predicate.ProfileUKM{}, puq.predicates...),
		// clone intermediate query.
		sql: puq.sql.Clone(),
	}
}

//  WithOwnerProfile tells the query-builder to eager-loads the nodes that are connected to
// the "owner_profile" edge. The optional arguments used to configure the query builder of the edge.
func (puq *ProfileUKMQuery) WithOwnerProfile(opts ...func(*ProfileQuery)) *ProfileUKMQuery {
	query := &ProfileQuery{config: puq.config}
	for _, opt := range opts {
		opt(query)
	}
	puq.withOwnerProfile = query
	return puq
}

//  WithOwnerUkm tells the query-builder to eager-loads the nodes that are connected to
// the "owner_ukm" edge. The optional arguments used to configure the query builder of the edge.
func (puq *ProfileUKMQuery) WithOwnerUkm(opts ...func(*UkmQuery)) *ProfileUKMQuery {
	query := &UkmQuery{config: puq.config}
	for _, opt := range opts {
		opt(query)
	}
	puq.withOwnerUkm = query
	return puq
}

//  WithOwnerRole tells the query-builder to eager-loads the nodes that are connected to
// the "owner_role" edge. The optional arguments used to configure the query builder of the edge.
func (puq *ProfileUKMQuery) WithOwnerRole(opts ...func(*RoleUKMQuery)) *ProfileUKMQuery {
	query := &RoleUKMQuery{config: puq.config}
	for _, opt := range opts {
		opt(query)
	}
	puq.withOwnerRole = query
	return puq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Reason string `json:"reason,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProfileUKM.Query().
//		GroupBy(profileukm.FieldReason).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (puq *ProfileUKMQuery) GroupBy(field string, fields ...string) *ProfileUKMGroupBy {
	group := &ProfileUKMGroupBy{config: puq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = puq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Reason string `json:"reason,omitempty"`
//	}
//
//	client.ProfileUKM.Query().
//		Select(profileukm.FieldReason).
//		Scan(ctx, &v)
//
func (puq *ProfileUKMQuery) Select(field string, fields ...string) *ProfileUKMSelect {
	selector := &ProfileUKMSelect{config: puq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = puq.sqlQuery()
	return selector
}

func (puq *ProfileUKMQuery) sqlAll(ctx context.Context) ([]*ProfileUKM, error) {
	var (
		nodes       = []*ProfileUKM{}
		withFKs     = puq.withFKs
		_spec       = puq.querySpec()
		loadedTypes = [3]bool{
			puq.withOwnerProfile != nil,
			puq.withOwnerUkm != nil,
			puq.withOwnerRole != nil,
		}
	)
	if puq.withOwnerProfile != nil || puq.withOwnerUkm != nil || puq.withOwnerRole != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, profileukm.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &ProfileUKM{config: puq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, puq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := puq.withOwnerProfile; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ProfileUKM)
		for i := range nodes {
			if fk := nodes[i].profile_ukms; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(profile.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "profile_ukms" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.OwnerProfile = n
			}
		}
	}

	if query := puq.withOwnerUkm; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ProfileUKM)
		for i := range nodes {
			if fk := nodes[i].ukm_profiles; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(ukm.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "ukm_profiles" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.OwnerUkm = n
			}
		}
	}

	if query := puq.withOwnerRole; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ProfileUKM)
		for i := range nodes {
			if fk := nodes[i].role_ukm_profile_roles; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(roleukm.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "role_ukm_profile_roles" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.OwnerRole = n
			}
		}
	}

	return nodes, nil
}

func (puq *ProfileUKMQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := puq.querySpec()
	return sqlgraph.CountNodes(ctx, puq.driver, _spec)
}

func (puq *ProfileUKMQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := puq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (puq *ProfileUKMQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   profileukm.Table,
			Columns: profileukm.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: profileukm.FieldID,
			},
		},
		From:   puq.sql,
		Unique: true,
	}
	if ps := puq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := puq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := puq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := puq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (puq *ProfileUKMQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(puq.driver.Dialect())
	t1 := builder.Table(profileukm.Table)
	selector := builder.Select(t1.Columns(profileukm.Columns...)...).From(t1)
	if puq.sql != nil {
		selector = puq.sql
		selector.Select(selector.Columns(profileukm.Columns...)...)
	}
	for _, p := range puq.predicates {
		p(selector)
	}
	for _, p := range puq.order {
		p(selector)
	}
	if offset := puq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := puq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProfileUKMGroupBy is the builder for group-by ProfileUKM entities.
type ProfileUKMGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pugb *ProfileUKMGroupBy) Aggregate(fns ...Aggregate) *ProfileUKMGroupBy {
	pugb.fns = append(pugb.fns, fns...)
	return pugb
}

// Scan applies the group-by query and scan the result into the given value.
func (pugb *ProfileUKMGroupBy) Scan(ctx context.Context, v interface{}) error {
	return pugb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pugb *ProfileUKMGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (pugb *ProfileUKMGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pugb.fields) > 1 {
		return nil, errors.New("ent: ProfileUKMGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pugb *ProfileUKMGroupBy) StringsX(ctx context.Context) []string {
	v, err := pugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (pugb *ProfileUKMGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pugb.fields) > 1 {
		return nil, errors.New("ent: ProfileUKMGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pugb *ProfileUKMGroupBy) IntsX(ctx context.Context) []int {
	v, err := pugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (pugb *ProfileUKMGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pugb.fields) > 1 {
		return nil, errors.New("ent: ProfileUKMGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pugb *ProfileUKMGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (pugb *ProfileUKMGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pugb.fields) > 1 {
		return nil, errors.New("ent: ProfileUKMGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pugb *ProfileUKMGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pugb *ProfileUKMGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pugb.sqlQuery().Query()
	if err := pugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pugb *ProfileUKMGroupBy) sqlQuery() *sql.Selector {
	selector := pugb.sql
	columns := make([]string, 0, len(pugb.fields)+len(pugb.fns))
	columns = append(columns, pugb.fields...)
	for _, fn := range pugb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(pugb.fields...)
}

// ProfileUKMSelect is the builder for select fields of ProfileUKM entities.
type ProfileUKMSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (pus *ProfileUKMSelect) Scan(ctx context.Context, v interface{}) error {
	return pus.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pus *ProfileUKMSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pus.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (pus *ProfileUKMSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pus.fields) > 1 {
		return nil, errors.New("ent: ProfileUKMSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pus *ProfileUKMSelect) StringsX(ctx context.Context) []string {
	v, err := pus.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (pus *ProfileUKMSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pus.fields) > 1 {
		return nil, errors.New("ent: ProfileUKMSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pus *ProfileUKMSelect) IntsX(ctx context.Context) []int {
	v, err := pus.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (pus *ProfileUKMSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pus.fields) > 1 {
		return nil, errors.New("ent: ProfileUKMSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pus *ProfileUKMSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pus.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (pus *ProfileUKMSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pus.fields) > 1 {
		return nil, errors.New("ent: ProfileUKMSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pus *ProfileUKMSelect) BoolsX(ctx context.Context) []bool {
	v, err := pus.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pus *ProfileUKMSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pus.sqlQuery().Query()
	if err := pus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pus *ProfileUKMSelect) sqlQuery() sql.Querier {
	selector := pus.sql
	selector.Select(selector.Columns(pus.fields...)...)
	return selector
}