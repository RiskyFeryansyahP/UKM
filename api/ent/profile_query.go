// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/confus1on/UKM/ent/predicate"
	"github.com/confus1on/UKM/ent/profile"
	"github.com/confus1on/UKM/ent/profileukm"
	"github.com/confus1on/UKM/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ProfileQuery is the builder for querying Profile entities.
type ProfileQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Profile
	// eager-loading edges.
	withOwner *UserQuery
	withUkms  *ProfileUKMQuery
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (pq *ProfileQuery) Where(ps ...predicate.Profile) *ProfileQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *ProfileQuery) Limit(limit int) *ProfileQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *ProfileQuery) Offset(offset int) *ProfileQuery {
	pq.offset = &offset
	return pq
}

// Order adds an order step to the query.
func (pq *ProfileQuery) Order(o ...Order) *ProfileQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryOwner chains the current query on the owner edge.
func (pq *ProfileQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: pq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(profile.Table, profile.FieldID, pq.sqlQuery()),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, profile.OwnerTable, profile.OwnerColumn),
	)
	query.sql = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
	return query
}

// QueryUkms chains the current query on the ukms edge.
func (pq *ProfileQuery) QueryUkms() *ProfileUKMQuery {
	query := &ProfileUKMQuery{config: pq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(profile.Table, profile.FieldID, pq.sqlQuery()),
		sqlgraph.To(profileukm.Table, profileukm.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, profile.UkmsTable, profile.UkmsColumn),
	)
	query.sql = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
	return query
}

// First returns the first Profile entity in the query. Returns *NotFoundError when no profile was found.
func (pq *ProfileQuery) First(ctx context.Context) (*Profile, error) {
	prs, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(prs) == 0 {
		return nil, &NotFoundError{profile.Label}
	}
	return prs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *ProfileQuery) FirstX(ctx context.Context) *Profile {
	pr, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return pr
}

// FirstID returns the first Profile id in the query. Returns *NotFoundError when no id was found.
func (pq *ProfileQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{profile.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (pq *ProfileQuery) FirstXID(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Profile entity in the query, returns an error if not exactly one entity was returned.
func (pq *ProfileQuery) Only(ctx context.Context) (*Profile, error) {
	prs, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(prs) {
	case 1:
		return prs[0], nil
	case 0:
		return nil, &NotFoundError{profile.Label}
	default:
		return nil, &NotSingularError{profile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *ProfileQuery) OnlyX(ctx context.Context) *Profile {
	pr, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// OnlyID returns the only Profile id in the query, returns an error if not exactly one id was returned.
func (pq *ProfileQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{profile.Label}
	default:
		err = &NotSingularError{profile.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (pq *ProfileQuery) OnlyXID(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Profiles.
func (pq *ProfileQuery) All(ctx context.Context) ([]*Profile, error) {
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *ProfileQuery) AllX(ctx context.Context) []*Profile {
	prs, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return prs
}

// IDs executes the query and returns a list of Profile ids.
func (pq *ProfileQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pq.Select(profile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *ProfileQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *ProfileQuery) Count(ctx context.Context) (int, error) {
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *ProfileQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *ProfileQuery) Exist(ctx context.Context) (bool, error) {
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *ProfileQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *ProfileQuery) Clone() *ProfileQuery {
	return &ProfileQuery{
		config:     pq.config,
		limit:      pq.limit,
		offset:     pq.offset,
		order:      append([]Order{}, pq.order...),
		unique:     append([]string{}, pq.unique...),
		predicates: append([]predicate.Profile{}, pq.predicates...),
		// clone intermediate query.
		sql: pq.sql.Clone(),
	}
}

//  WithOwner tells the query-builder to eager-loads the nodes that are connected to
// the "owner" edge. The optional arguments used to configure the query builder of the edge.
func (pq *ProfileQuery) WithOwner(opts ...func(*UserQuery)) *ProfileQuery {
	query := &UserQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withOwner = query
	return pq
}

//  WithUkms tells the query-builder to eager-loads the nodes that are connected to
// the "ukms" edge. The optional arguments used to configure the query builder of the edge.
func (pq *ProfileQuery) WithUkms(opts ...func(*ProfileUKMQuery)) *ProfileQuery {
	query := &ProfileUKMQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withUkms = query
	return pq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FirstName string `json:"firstName,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Profile.Query().
//		GroupBy(profile.FieldFirstName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pq *ProfileQuery) GroupBy(field string, fields ...string) *ProfileGroupBy {
	group := &ProfileGroupBy{config: pq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = pq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		FirstName string `json:"firstName,omitempty"`
//	}
//
//	client.Profile.Query().
//		Select(profile.FieldFirstName).
//		Scan(ctx, &v)
//
func (pq *ProfileQuery) Select(field string, fields ...string) *ProfileSelect {
	selector := &ProfileSelect{config: pq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = pq.sqlQuery()
	return selector
}

func (pq *ProfileQuery) sqlAll(ctx context.Context) ([]*Profile, error) {
	var (
		nodes       = []*Profile{}
		_spec       = pq.querySpec()
		loadedTypes = [2]bool{
			pq.withOwner != nil,
			pq.withUkms != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Profile{config: pq.config}
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
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pq.withOwner; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Profile)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.User(func(s *sql.Selector) {
			s.Where(sql.InValues(profile.OwnerColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.user_profile
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "user_profile" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_profile" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Owner = append(node.Edges.Owner, n)
		}
	}

	if query := pq.withUkms; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Profile)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.ProfileUKM(func(s *sql.Selector) {
			s.Where(sql.InValues(profile.UkmsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.profile_ukms
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "profile_ukms" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "profile_ukms" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Ukms = append(node.Edges.Ukms, n)
		}
	}

	return nodes, nil
}

func (pq *ProfileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *ProfileQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (pq *ProfileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   profile.Table,
			Columns: profile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: profile.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *ProfileQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(profile.Table)
	selector := builder.Select(t1.Columns(profile.Columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(profile.Columns...)...)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProfileGroupBy is the builder for group-by Profile entities.
type ProfileGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *ProfileGroupBy) Aggregate(fns ...Aggregate) *ProfileGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scan the result into the given value.
func (pgb *ProfileGroupBy) Scan(ctx context.Context, v interface{}) error {
	return pgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pgb *ProfileGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (pgb *ProfileGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: ProfileGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pgb *ProfileGroupBy) StringsX(ctx context.Context) []string {
	v, err := pgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (pgb *ProfileGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: ProfileGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pgb *ProfileGroupBy) IntsX(ctx context.Context) []int {
	v, err := pgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (pgb *ProfileGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: ProfileGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pgb *ProfileGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (pgb *ProfileGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: ProfileGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pgb *ProfileGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pgb *ProfileGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pgb.sqlQuery().Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *ProfileGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql
	columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
	columns = append(columns, pgb.fields...)
	for _, fn := range pgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(pgb.fields...)
}

// ProfileSelect is the builder for select fields of Profile entities.
type ProfileSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (ps *ProfileSelect) Scan(ctx context.Context, v interface{}) error {
	return ps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ps *ProfileSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ps *ProfileSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: ProfileSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ps *ProfileSelect) StringsX(ctx context.Context) []string {
	v, err := ps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ps *ProfileSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: ProfileSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ps *ProfileSelect) IntsX(ctx context.Context) []int {
	v, err := ps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ps *ProfileSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: ProfileSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ps *ProfileSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ps *ProfileSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: ProfileSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ps *ProfileSelect) BoolsX(ctx context.Context) []bool {
	v, err := ps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ps *ProfileSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sqlQuery().Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ps *ProfileSelect) sqlQuery() sql.Querier {
	selector := ps.sql
	selector.Select(selector.Columns(ps.fields...)...)
	return selector
}
