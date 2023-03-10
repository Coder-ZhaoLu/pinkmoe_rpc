// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/predicate"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/sitemeta"
)

// SitemetaQuery is the builder for querying Sitemeta entities.
type SitemetaQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Sitemeta
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SitemetaQuery builder.
func (sq *SitemetaQuery) Where(ps ...predicate.Sitemeta) *SitemetaQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SitemetaQuery) Limit(limit int) *SitemetaQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SitemetaQuery) Offset(offset int) *SitemetaQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SitemetaQuery) Unique(unique bool) *SitemetaQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SitemetaQuery) Order(o ...OrderFunc) *SitemetaQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// First returns the first Sitemeta entity from the query.
// Returns a *NotFoundError when no Sitemeta was found.
func (sq *SitemetaQuery) First(ctx context.Context) (*Sitemeta, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sitemeta.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SitemetaQuery) FirstX(ctx context.Context) *Sitemeta {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Sitemeta ID from the query.
// Returns a *NotFoundError when no Sitemeta ID was found.
func (sq *SitemetaQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sitemeta.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SitemetaQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Sitemeta entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Sitemeta entity is found.
// Returns a *NotFoundError when no Sitemeta entities are found.
func (sq *SitemetaQuery) Only(ctx context.Context) (*Sitemeta, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sitemeta.Label}
	default:
		return nil, &NotSingularError{sitemeta.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SitemetaQuery) OnlyX(ctx context.Context) *Sitemeta {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Sitemeta ID in the query.
// Returns a *NotSingularError when more than one Sitemeta ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SitemetaQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sitemeta.Label}
	default:
		err = &NotSingularError{sitemeta.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SitemetaQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SitemetaSlice.
func (sq *SitemetaQuery) All(ctx context.Context) ([]*Sitemeta, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Sitemeta, *SitemetaQuery]()
	return withInterceptors[[]*Sitemeta](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SitemetaQuery) AllX(ctx context.Context) []*Sitemeta {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Sitemeta IDs.
func (sq *SitemetaQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err = sq.Select(sitemeta.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SitemetaQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SitemetaQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SitemetaQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SitemetaQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SitemetaQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SitemetaQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SitemetaQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SitemetaQuery) Clone() *SitemetaQuery {
	if sq == nil {
		return nil
	}
	return &SitemetaQuery{
		config:     sq.config,
		ctx:        sq.ctx.Clone(),
		order:      append([]OrderFunc{}, sq.order...),
		inters:     append([]Interceptor{}, sq.inters...),
		predicates: append([]predicate.Sitemeta{}, sq.predicates...),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Sitemeta.Query().
//		GroupBy(sitemeta.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SitemetaQuery) GroupBy(field string, fields ...string) *SitemetaGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SitemetaGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = sitemeta.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Sitemeta.Query().
//		Select(sitemeta.FieldCreatedAt).
//		Scan(ctx, &v)
func (sq *SitemetaQuery) Select(fields ...string) *SitemetaSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SitemetaSelect{SitemetaQuery: sq}
	sbuild.label = sitemeta.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SitemetaSelect configured with the given aggregations.
func (sq *SitemetaQuery) Aggregate(fns ...AggregateFunc) *SitemetaSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SitemetaQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !sitemeta.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SitemetaQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Sitemeta, error) {
	var (
		nodes = []*Sitemeta{}
		_spec = sq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Sitemeta).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Sitemeta{config: sq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sq *SitemetaQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SitemetaQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sitemeta.Table, sitemeta.Columns, sqlgraph.NewFieldSpec(sitemeta.FieldID, field.TypeUint64))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sitemeta.FieldID)
		for i := range fields {
			if fields[i] != sitemeta.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SitemetaQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(sitemeta.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = sitemeta.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SitemetaGroupBy is the group-by builder for Sitemeta entities.
type SitemetaGroupBy struct {
	selector
	build *SitemetaQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SitemetaGroupBy) Aggregate(fns ...AggregateFunc) *SitemetaGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SitemetaGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SitemetaQuery, *SitemetaGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SitemetaGroupBy) sqlScan(ctx context.Context, root *SitemetaQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SitemetaSelect is the builder for selecting fields of Sitemeta entities.
type SitemetaSelect struct {
	*SitemetaQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SitemetaSelect) Aggregate(fns ...AggregateFunc) *SitemetaSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SitemetaSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SitemetaQuery, *SitemetaSelect](ctx, ss.SitemetaQuery, ss, ss.inters, v)
}

func (ss *SitemetaSelect) sqlScan(ctx context.Context, root *SitemetaQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
