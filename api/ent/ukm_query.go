// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/confus1on/UKM/ent/predicate"
	"github.com/confus1on/UKM/ent/profileukm"
	"github.com/confus1on/UKM/ent/ukm"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UkmQuery is the builder for querying Ukm entities.
type UkmQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Ukm
	// eager-loading edges.
	withProfiles *ProfileUKMQuery
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (uq *UkmQuery) Where(ps ...predicate.Ukm) *UkmQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit adds a limit step to the query.
func (uq *UkmQuery) Limit(limit int) *UkmQuery {
	uq.limit = &limit
	return uq
}

// Offset adds an offset step to the query.
func (uq *UkmQuery) Offset(offset int) *UkmQuery {
	uq.offset = &offset
	return uq
}

// Order adds an order step to the query.
func (uq *UkmQuery) Order(o ...Order) *UkmQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryProfiles chains the current query on the profiles edge.
func (uq *UkmQuery) QueryProfiles() *ProfileUKMQuery {
	query := &ProfileUKMQuery{config: uq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(ukm.Table, ukm.FieldID, uq.sqlQuery()),
		sqlgraph.To(profileukm.Table, profileukm.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ukm.ProfilesTable, ukm.ProfilesColumn),
	)
	query.sql = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
	return query
}

// First returns the first Ukm entity in the query. Returns *NotFoundError when no ukm was found.
func (uq *UkmQuery) First(ctx context.Context) (*Ukm, error) {
	us, err := uq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(us) == 0 {
		return nil, &NotFoundError{ukm.Label}
	}
	return us[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UkmQuery) FirstX(ctx context.Context) *Ukm {
	u, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return u
}

// FirstID returns the first Ukm id in the query. Returns *NotFoundError when no id was found.
func (uq *UkmQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ukm.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (uq *UkmQuery) FirstXID(ctx context.Context) int {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Ukm entity in the query, returns an error if not exactly one entity was returned.
func (uq *UkmQuery) Only(ctx context.Context) (*Ukm, error) {
	us, err := uq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(us) {
	case 1:
		return us[0], nil
	case 0:
		return nil, &NotFoundError{ukm.Label}
	default:
		return nil, &NotSingularError{ukm.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UkmQuery) OnlyX(ctx context.Context) *Ukm {
	u, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// OnlyID returns the only Ukm id in the query, returns an error if not exactly one id was returned.
func (uq *UkmQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ukm.Label}
	default:
		err = &NotSingularError{ukm.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (uq *UkmQuery) OnlyXID(ctx context.Context) int {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Ukms.
func (uq *UkmQuery) All(ctx context.Context) ([]*Ukm, error) {
	return uq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uq *UkmQuery) AllX(ctx context.Context) []*Ukm {
	us, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return us
}

// IDs executes the query and returns a list of Ukm ids.
func (uq *UkmQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uq.Select(ukm.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UkmQuery) IDsX(ctx context.Context) []int {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UkmQuery) Count(ctx context.Context) (int, error) {
	return uq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UkmQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UkmQuery) Exist(ctx context.Context) (bool, error) {
	return uq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UkmQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UkmQuery) Clone() *UkmQuery {
	return &UkmQuery{
		config:     uq.config,
		limit:      uq.limit,
		offset:     uq.offset,
		order:      append([]Order{}, uq.order...),
		unique:     append([]string{}, uq.unique...),
		predicates: append([]predicate.Ukm{}, uq.predicates...),
		// clone intermediate query.
		sql: uq.sql.Clone(),
	}
}

//  WithProfiles tells the query-builder to eager-loads the nodes that are connected to
// the "profiles" edge. The optional arguments used to configure the query builder of the edge.
func (uq *UkmQuery) WithProfiles(opts ...func(*ProfileUKMQuery)) *UkmQuery {
	query := &ProfileUKMQuery{config: uq.config}
	for _, opt := range opts {
		opt(query)
	}
	uq.withProfiles = query
	return uq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Ukm.Query().
//		GroupBy(ukm.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uq *UkmQuery) GroupBy(field string, fields ...string) *UkmGroupBy {
	group := &UkmGroupBy{config: uq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = uq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Ukm.Query().
//		Select(ukm.FieldName).
//		Scan(ctx, &v)
//
func (uq *UkmQuery) Select(field string, fields ...string) *UkmSelect {
	selector := &UkmSelect{config: uq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = uq.sqlQuery()
	return selector
}

func (uq *UkmQuery) sqlAll(ctx context.Context) ([]*Ukm, error) {
	var (
		nodes       = []*Ukm{}
		_spec       = uq.querySpec()
		loadedTypes = [1]bool{
			uq.withProfiles != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Ukm{config: uq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
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
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := uq.withProfiles; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Ukm)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.ProfileUKM(func(s *sql.Selector) {
			s.Where(sql.InValues(ukm.ProfilesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.ukm_profiles
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "ukm_profiles" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "ukm_profiles" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Profiles = append(node.Edges.Profiles, n)
		}
	}

	return nodes, nil
}

func (uq *UkmQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UkmQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (uq *UkmQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ukm.Table,
			Columns: ukm.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ukm.FieldID,
			},
		},
		From:   uq.sql,
		Unique: true,
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UkmQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(ukm.Table)
	selector := builder.Select(t1.Columns(ukm.Columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(ukm.Columns...)...)
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UkmGroupBy is the builder for group-by Ukm entities.
type UkmGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UkmGroupBy) Aggregate(fns ...Aggregate) *UkmGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the group-by query and scan the result into the given value.
func (ugb *UkmGroupBy) Scan(ctx context.Context, v interface{}) error {
	return ugb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ugb *UkmGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ugb *UkmGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UkmGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ugb *UkmGroupBy) StringsX(ctx context.Context) []string {
	v, err := ugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ugb *UkmGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UkmGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ugb *UkmGroupBy) IntsX(ctx context.Context) []int {
	v, err := ugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ugb *UkmGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UkmGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ugb *UkmGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ugb *UkmGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ugb.fields) > 1 {
		return nil, errors.New("ent: UkmGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ugb *UkmGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ugb *UkmGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ugb.sqlQuery().Query()
	if err := ugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ugb *UkmGroupBy) sqlQuery() *sql.Selector {
	selector := ugb.sql
	columns := make([]string, 0, len(ugb.fields)+len(ugb.fns))
	columns = append(columns, ugb.fields...)
	for _, fn := range ugb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(ugb.fields...)
}

// UkmSelect is the builder for select fields of Ukm entities.
type UkmSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (us *UkmSelect) Scan(ctx context.Context, v interface{}) error {
	return us.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (us *UkmSelect) ScanX(ctx context.Context, v interface{}) {
	if err := us.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (us *UkmSelect) Strings(ctx context.Context) ([]string, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UkmSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (us *UkmSelect) StringsX(ctx context.Context) []string {
	v, err := us.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (us *UkmSelect) Ints(ctx context.Context) ([]int, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UkmSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (us *UkmSelect) IntsX(ctx context.Context) []int {
	v, err := us.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (us *UkmSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UkmSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (us *UkmSelect) Float64sX(ctx context.Context) []float64 {
	v, err := us.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (us *UkmSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(us.fields) > 1 {
		return nil, errors.New("ent: UkmSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := us.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (us *UkmSelect) BoolsX(ctx context.Context) []bool {
	v, err := us.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (us *UkmSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := us.sqlQuery().Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (us *UkmSelect) sqlQuery() sql.Querier {
	selector := us.sql
	selector.Select(selector.Columns(us.fields...)...)
	return selector
}
