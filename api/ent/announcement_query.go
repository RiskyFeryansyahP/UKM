// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/confus1on/UKM/ent/announcement"
	"github.com/confus1on/UKM/ent/predicate"
	"github.com/confus1on/UKM/ent/ukm"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// AnnouncementQuery is the builder for querying Announcement entities.
type AnnouncementQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Announcement
	// eager-loading edges.
	withOwnerAnnouncement *UkmQuery
	withFKs               bool
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (aq *AnnouncementQuery) Where(ps ...predicate.Announcement) *AnnouncementQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AnnouncementQuery) Limit(limit int) *AnnouncementQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AnnouncementQuery) Offset(offset int) *AnnouncementQuery {
	aq.offset = &offset
	return aq
}

// Order adds an order step to the query.
func (aq *AnnouncementQuery) Order(o ...Order) *AnnouncementQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryOwnerAnnouncement chains the current query on the owner_announcement edge.
func (aq *AnnouncementQuery) QueryOwnerAnnouncement() *UkmQuery {
	query := &UkmQuery{config: aq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(announcement.Table, announcement.FieldID, aq.sqlQuery()),
		sqlgraph.To(ukm.Table, ukm.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, announcement.OwnerAnnouncementTable, announcement.OwnerAnnouncementColumn),
	)
	query.sql = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
	return query
}

// First returns the first Announcement entity in the query. Returns *NotFoundError when no announcement was found.
func (aq *AnnouncementQuery) First(ctx context.Context) (*Announcement, error) {
	as, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(as) == 0 {
		return nil, &NotFoundError{announcement.Label}
	}
	return as[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AnnouncementQuery) FirstX(ctx context.Context) *Announcement {
	a, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return a
}

// FirstID returns the first Announcement id in the query. Returns *NotFoundError when no id was found.
func (aq *AnnouncementQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{announcement.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (aq *AnnouncementQuery) FirstXID(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Announcement entity in the query, returns an error if not exactly one entity was returned.
func (aq *AnnouncementQuery) Only(ctx context.Context) (*Announcement, error) {
	as, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(as) {
	case 1:
		return as[0], nil
	case 0:
		return nil, &NotFoundError{announcement.Label}
	default:
		return nil, &NotSingularError{announcement.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AnnouncementQuery) OnlyX(ctx context.Context) *Announcement {
	a, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// OnlyID returns the only Announcement id in the query, returns an error if not exactly one id was returned.
func (aq *AnnouncementQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{announcement.Label}
	default:
		err = &NotSingularError{announcement.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (aq *AnnouncementQuery) OnlyXID(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Announcements.
func (aq *AnnouncementQuery) All(ctx context.Context) ([]*Announcement, error) {
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AnnouncementQuery) AllX(ctx context.Context) []*Announcement {
	as, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return as
}

// IDs executes the query and returns a list of Announcement ids.
func (aq *AnnouncementQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := aq.Select(announcement.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AnnouncementQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AnnouncementQuery) Count(ctx context.Context) (int, error) {
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AnnouncementQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AnnouncementQuery) Exist(ctx context.Context) (bool, error) {
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AnnouncementQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AnnouncementQuery) Clone() *AnnouncementQuery {
	return &AnnouncementQuery{
		config:     aq.config,
		limit:      aq.limit,
		offset:     aq.offset,
		order:      append([]Order{}, aq.order...),
		unique:     append([]string{}, aq.unique...),
		predicates: append([]predicate.Announcement{}, aq.predicates...),
		// clone intermediate query.
		sql: aq.sql.Clone(),
	}
}

//  WithOwnerAnnouncement tells the query-builder to eager-loads the nodes that are connected to
// the "owner_announcement" edge. The optional arguments used to configure the query builder of the edge.
func (aq *AnnouncementQuery) WithOwnerAnnouncement(opts ...func(*UkmQuery)) *AnnouncementQuery {
	query := &UkmQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withOwnerAnnouncement = query
	return aq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Announcement.Query().
//		GroupBy(announcement.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aq *AnnouncementQuery) GroupBy(field string, fields ...string) *AnnouncementGroupBy {
	group := &AnnouncementGroupBy{config: aq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = aq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Announcement.Query().
//		Select(announcement.FieldTitle).
//		Scan(ctx, &v)
//
func (aq *AnnouncementQuery) Select(field string, fields ...string) *AnnouncementSelect {
	selector := &AnnouncementSelect{config: aq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = aq.sqlQuery()
	return selector
}

func (aq *AnnouncementQuery) sqlAll(ctx context.Context) ([]*Announcement, error) {
	var (
		nodes       = []*Announcement{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [1]bool{
			aq.withOwnerAnnouncement != nil,
		}
	)
	if aq.withOwnerAnnouncement != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, announcement.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Announcement{config: aq.config}
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
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := aq.withOwnerAnnouncement; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Announcement)
		for i := range nodes {
			if fk := nodes[i].ukm_announcement; fk != nil {
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
				return nil, fmt.Errorf(`unexpected foreign-key "ukm_announcement" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.OwnerAnnouncement = n
			}
		}
	}

	return nodes, nil
}

func (aq *AnnouncementQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AnnouncementQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (aq *AnnouncementQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   announcement.Table,
			Columns: announcement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: announcement.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AnnouncementQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(announcement.Table)
	selector := builder.Select(t1.Columns(announcement.Columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(announcement.Columns...)...)
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AnnouncementGroupBy is the builder for group-by Announcement entities.
type AnnouncementGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AnnouncementGroupBy) Aggregate(fns ...Aggregate) *AnnouncementGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scan the result into the given value.
func (agb *AnnouncementGroupBy) Scan(ctx context.Context, v interface{}) error {
	return agb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (agb *AnnouncementGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := agb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (agb *AnnouncementGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AnnouncementGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (agb *AnnouncementGroupBy) StringsX(ctx context.Context) []string {
	v, err := agb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (agb *AnnouncementGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AnnouncementGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (agb *AnnouncementGroupBy) IntsX(ctx context.Context) []int {
	v, err := agb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (agb *AnnouncementGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AnnouncementGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (agb *AnnouncementGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := agb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (agb *AnnouncementGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AnnouncementGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (agb *AnnouncementGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := agb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (agb *AnnouncementGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := agb.sqlQuery().Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AnnouncementGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql
	columns := make([]string, 0, len(agb.fields)+len(agb.fns))
	columns = append(columns, agb.fields...)
	for _, fn := range agb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(agb.fields...)
}

// AnnouncementSelect is the builder for select fields of Announcement entities.
type AnnouncementSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (as *AnnouncementSelect) Scan(ctx context.Context, v interface{}) error {
	return as.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (as *AnnouncementSelect) ScanX(ctx context.Context, v interface{}) {
	if err := as.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (as *AnnouncementSelect) Strings(ctx context.Context) ([]string, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AnnouncementSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (as *AnnouncementSelect) StringsX(ctx context.Context) []string {
	v, err := as.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (as *AnnouncementSelect) Ints(ctx context.Context) ([]int, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AnnouncementSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (as *AnnouncementSelect) IntsX(ctx context.Context) []int {
	v, err := as.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (as *AnnouncementSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AnnouncementSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (as *AnnouncementSelect) Float64sX(ctx context.Context) []float64 {
	v, err := as.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (as *AnnouncementSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AnnouncementSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (as *AnnouncementSelect) BoolsX(ctx context.Context) []bool {
	v, err := as.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (as *AnnouncementSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := as.sqlQuery().Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (as *AnnouncementSelect) sqlQuery() sql.Querier {
	selector := as.sql
	selector.Select(selector.Columns(as.fields...)...)
	return selector
}