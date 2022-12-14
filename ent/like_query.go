// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/MangoSteen0903/go-blog-graphql/ent/comment"
	"github.com/MangoSteen0903/go-blog-graphql/ent/like"
	"github.com/MangoSteen0903/go-blog-graphql/ent/post"
	"github.com/MangoSteen0903/go-blog-graphql/ent/predicate"
	"github.com/MangoSteen0903/go-blog-graphql/ent/user"
)

// LikeQuery is the builder for querying Like entities.
type LikeQuery struct {
	config
	limit             *int
	offset            *int
	unique            *bool
	order             []OrderFunc
	fields            []string
	predicates        []predicate.Like
	withPosts         *PostQuery
	withOwner         *UserQuery
	withComments      *CommentQuery
	modifiers         []func(*sql.Selector)
	loadTotal         []func(context.Context, []*Like) error
	withNamedPosts    map[string]*PostQuery
	withNamedOwner    map[string]*UserQuery
	withNamedComments map[string]*CommentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LikeQuery builder.
func (lq *LikeQuery) Where(ps ...predicate.Like) *LikeQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit adds a limit step to the query.
func (lq *LikeQuery) Limit(limit int) *LikeQuery {
	lq.limit = &limit
	return lq
}

// Offset adds an offset step to the query.
func (lq *LikeQuery) Offset(offset int) *LikeQuery {
	lq.offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LikeQuery) Unique(unique bool) *LikeQuery {
	lq.unique = &unique
	return lq
}

// Order adds an order step to the query.
func (lq *LikeQuery) Order(o ...OrderFunc) *LikeQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// QueryPosts chains the current query on the "Posts" edge.
func (lq *LikeQuery) QueryPosts() *PostQuery {
	query := &PostQuery{config: lq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(like.Table, like.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, like.PostsTable, like.PostsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOwner chains the current query on the "owner" edge.
func (lq *LikeQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: lq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(like.Table, like.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, like.OwnerTable, like.OwnerPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryComments chains the current query on the "comments" edge.
func (lq *LikeQuery) QueryComments() *CommentQuery {
	query := &CommentQuery{config: lq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(like.Table, like.FieldID, selector),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, like.CommentsTable, like.CommentsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Like entity from the query.
// Returns a *NotFoundError when no Like was found.
func (lq *LikeQuery) First(ctx context.Context) (*Like, error) {
	nodes, err := lq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{like.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LikeQuery) FirstX(ctx context.Context) *Like {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Like ID from the query.
// Returns a *NotFoundError when no Like ID was found.
func (lq *LikeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{like.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LikeQuery) FirstIDX(ctx context.Context) int {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Like entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Like entity is found.
// Returns a *NotFoundError when no Like entities are found.
func (lq *LikeQuery) Only(ctx context.Context) (*Like, error) {
	nodes, err := lq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{like.Label}
	default:
		return nil, &NotSingularError{like.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LikeQuery) OnlyX(ctx context.Context) *Like {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Like ID in the query.
// Returns a *NotSingularError when more than one Like ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LikeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{like.Label}
	default:
		err = &NotSingularError{like.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LikeQuery) OnlyIDX(ctx context.Context) int {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Likes.
func (lq *LikeQuery) All(ctx context.Context) ([]*Like, error) {
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return lq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (lq *LikeQuery) AllX(ctx context.Context) []*Like {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Like IDs.
func (lq *LikeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := lq.Select(like.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LikeQuery) IDsX(ctx context.Context) []int {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LikeQuery) Count(ctx context.Context) (int, error) {
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return lq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LikeQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LikeQuery) Exist(ctx context.Context) (bool, error) {
	if err := lq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return lq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LikeQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LikeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LikeQuery) Clone() *LikeQuery {
	if lq == nil {
		return nil
	}
	return &LikeQuery{
		config:       lq.config,
		limit:        lq.limit,
		offset:       lq.offset,
		order:        append([]OrderFunc{}, lq.order...),
		predicates:   append([]predicate.Like{}, lq.predicates...),
		withPosts:    lq.withPosts.Clone(),
		withOwner:    lq.withOwner.Clone(),
		withComments: lq.withComments.Clone(),
		// clone intermediate query.
		sql:    lq.sql.Clone(),
		path:   lq.path,
		unique: lq.unique,
	}
}

// WithPosts tells the query-builder to eager-load the nodes that are connected to
// the "Posts" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LikeQuery) WithPosts(opts ...func(*PostQuery)) *LikeQuery {
	query := &PostQuery{config: lq.config}
	for _, opt := range opts {
		opt(query)
	}
	lq.withPosts = query
	return lq
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LikeQuery) WithOwner(opts ...func(*UserQuery)) *LikeQuery {
	query := &UserQuery{config: lq.config}
	for _, opt := range opts {
		opt(query)
	}
	lq.withOwner = query
	return lq
}

// WithComments tells the query-builder to eager-load the nodes that are connected to
// the "comments" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LikeQuery) WithComments(opts ...func(*CommentQuery)) *LikeQuery {
	query := &CommentQuery{config: lq.config}
	for _, opt := range opts {
		opt(query)
	}
	lq.withComments = query
	return lq
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
//	client.Like.Query().
//		GroupBy(like.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lq *LikeQuery) GroupBy(field string, fields ...string) *LikeGroupBy {
	grbuild := &LikeGroupBy{config: lq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return lq.sqlQuery(ctx), nil
	}
	grbuild.label = like.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
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
//	client.Like.Query().
//		Select(like.FieldCreatedAt).
//		Scan(ctx, &v)
func (lq *LikeQuery) Select(fields ...string) *LikeSelect {
	lq.fields = append(lq.fields, fields...)
	selbuild := &LikeSelect{LikeQuery: lq}
	selbuild.label = like.Label
	selbuild.flds, selbuild.scan = &lq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a LikeSelect configured with the given aggregations.
func (lq *LikeQuery) Aggregate(fns ...AggregateFunc) *LikeSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LikeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range lq.fields {
		if !like.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LikeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Like, error) {
	var (
		nodes       = []*Like{}
		_spec       = lq.querySpec()
		loadedTypes = [3]bool{
			lq.withPosts != nil,
			lq.withOwner != nil,
			lq.withComments != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Like).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Like{config: lq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(lq.modifiers) > 0 {
		_spec.Modifiers = lq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lq.withPosts; query != nil {
		if err := lq.loadPosts(ctx, query, nodes,
			func(n *Like) { n.Edges.Posts = []*Post{} },
			func(n *Like, e *Post) { n.Edges.Posts = append(n.Edges.Posts, e) }); err != nil {
			return nil, err
		}
	}
	if query := lq.withOwner; query != nil {
		if err := lq.loadOwner(ctx, query, nodes,
			func(n *Like) { n.Edges.Owner = []*User{} },
			func(n *Like, e *User) { n.Edges.Owner = append(n.Edges.Owner, e) }); err != nil {
			return nil, err
		}
	}
	if query := lq.withComments; query != nil {
		if err := lq.loadComments(ctx, query, nodes,
			func(n *Like) { n.Edges.Comments = []*Comment{} },
			func(n *Like, e *Comment) { n.Edges.Comments = append(n.Edges.Comments, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range lq.withNamedPosts {
		if err := lq.loadPosts(ctx, query, nodes,
			func(n *Like) { n.appendNamedPosts(name) },
			func(n *Like, e *Post) { n.appendNamedPosts(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range lq.withNamedOwner {
		if err := lq.loadOwner(ctx, query, nodes,
			func(n *Like) { n.appendNamedOwner(name) },
			func(n *Like, e *User) { n.appendNamedOwner(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range lq.withNamedComments {
		if err := lq.loadComments(ctx, query, nodes,
			func(n *Like) { n.appendNamedComments(name) },
			func(n *Like, e *Comment) { n.appendNamedComments(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range lq.loadTotal {
		if err := lq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lq *LikeQuery) loadPosts(ctx context.Context, query *PostQuery, nodes []*Like, init func(*Like), assign func(*Like, *Post)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Like)
	nids := make(map[int]map[*Like]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(like.PostsTable)
		s.Join(joinT).On(s.C(post.FieldID), joinT.C(like.PostsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(like.PostsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(like.PostsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Like]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "Posts" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (lq *LikeQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*Like, init func(*Like), assign func(*Like, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Like)
	nids := make(map[int]map[*Like]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(like.OwnerTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(like.OwnerPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(like.OwnerPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(like.OwnerPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Like]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "owner" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (lq *LikeQuery) loadComments(ctx context.Context, query *CommentQuery, nodes []*Like, init func(*Like), assign func(*Like, *Comment)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Like)
	nids := make(map[int]map[*Like]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(like.CommentsTable)
		s.Join(joinT).On(s.C(comment.FieldID), joinT.C(like.CommentsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(like.CommentsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(like.CommentsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Like]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "comments" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (lq *LikeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	if len(lq.modifiers) > 0 {
		_spec.Modifiers = lq.modifiers
	}
	_spec.Node.Columns = lq.fields
	if len(lq.fields) > 0 {
		_spec.Unique = lq.unique != nil && *lq.unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LikeQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := lq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (lq *LikeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   like.Table,
			Columns: like.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: like.FieldID,
			},
		},
		From:   lq.sql,
		Unique: true,
	}
	if unique := lq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := lq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, like.FieldID)
		for i := range fields {
			if fields[i] != like.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LikeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(like.Table)
	columns := lq.fields
	if len(columns) == 0 {
		columns = like.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.unique != nil && *lq.unique {
		selector.Distinct()
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedPosts tells the query-builder to eager-load the nodes that are connected to the "Posts"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (lq *LikeQuery) WithNamedPosts(name string, opts ...func(*PostQuery)) *LikeQuery {
	query := &PostQuery{config: lq.config}
	for _, opt := range opts {
		opt(query)
	}
	if lq.withNamedPosts == nil {
		lq.withNamedPosts = make(map[string]*PostQuery)
	}
	lq.withNamedPosts[name] = query
	return lq
}

// WithNamedOwner tells the query-builder to eager-load the nodes that are connected to the "owner"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (lq *LikeQuery) WithNamedOwner(name string, opts ...func(*UserQuery)) *LikeQuery {
	query := &UserQuery{config: lq.config}
	for _, opt := range opts {
		opt(query)
	}
	if lq.withNamedOwner == nil {
		lq.withNamedOwner = make(map[string]*UserQuery)
	}
	lq.withNamedOwner[name] = query
	return lq
}

// WithNamedComments tells the query-builder to eager-load the nodes that are connected to the "comments"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (lq *LikeQuery) WithNamedComments(name string, opts ...func(*CommentQuery)) *LikeQuery {
	query := &CommentQuery{config: lq.config}
	for _, opt := range opts {
		opt(query)
	}
	if lq.withNamedComments == nil {
		lq.withNamedComments = make(map[string]*CommentQuery)
	}
	lq.withNamedComments[name] = query
	return lq
}

// LikeGroupBy is the group-by builder for Like entities.
type LikeGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LikeGroupBy) Aggregate(fns ...AggregateFunc) *LikeGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the group-by query and scans the result into the given value.
func (lgb *LikeGroupBy) Scan(ctx context.Context, v any) error {
	query, err := lgb.path(ctx)
	if err != nil {
		return err
	}
	lgb.sql = query
	return lgb.sqlScan(ctx, v)
}

func (lgb *LikeGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range lgb.fields {
		if !like.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := lgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lgb *LikeGroupBy) sqlQuery() *sql.Selector {
	selector := lgb.sql.Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(lgb.fields)+len(lgb.fns))
		for _, f := range lgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(lgb.fields...)...)
}

// LikeSelect is the builder for selecting fields of Like entities.
type LikeSelect struct {
	*LikeQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LikeSelect) Aggregate(fns ...AggregateFunc) *LikeSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LikeSelect) Scan(ctx context.Context, v any) error {
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	ls.sql = ls.LikeQuery.sqlQuery(ctx)
	return ls.sqlScan(ctx, v)
}

func (ls *LikeSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(ls.sql))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ls.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ls.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ls.sql.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
