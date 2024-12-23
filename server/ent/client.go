// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"server/ent/migrate"

	"server/ent/label"
	"server/ent/priority"
	"server/ent/status"
	"server/ent/todo"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Label is the client for interacting with the Label builders.
	Label *LabelClient
	// Priority is the client for interacting with the Priority builders.
	Priority *PriorityClient
	// Status is the client for interacting with the Status builders.
	Status *StatusClient
	// Todo is the client for interacting with the Todo builders.
	Todo *TodoClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Label = NewLabelClient(c.config)
	c.Priority = NewPriorityClient(c.config)
	c.Status = NewStatusClient(c.config)
	c.Todo = NewTodoClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Label:    NewLabelClient(cfg),
		Priority: NewPriorityClient(cfg),
		Status:   NewStatusClient(cfg),
		Todo:     NewTodoClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Label:    NewLabelClient(cfg),
		Priority: NewPriorityClient(cfg),
		Status:   NewStatusClient(cfg),
		Todo:     NewTodoClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Label.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Label.Use(hooks...)
	c.Priority.Use(hooks...)
	c.Status.Use(hooks...)
	c.Todo.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Label.Intercept(interceptors...)
	c.Priority.Intercept(interceptors...)
	c.Status.Intercept(interceptors...)
	c.Todo.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *LabelMutation:
		return c.Label.mutate(ctx, m)
	case *PriorityMutation:
		return c.Priority.mutate(ctx, m)
	case *StatusMutation:
		return c.Status.mutate(ctx, m)
	case *TodoMutation:
		return c.Todo.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// LabelClient is a client for the Label schema.
type LabelClient struct {
	config
}

// NewLabelClient returns a client for the Label from the given config.
func NewLabelClient(c config) *LabelClient {
	return &LabelClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `label.Hooks(f(g(h())))`.
func (c *LabelClient) Use(hooks ...Hook) {
	c.hooks.Label = append(c.hooks.Label, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `label.Intercept(f(g(h())))`.
func (c *LabelClient) Intercept(interceptors ...Interceptor) {
	c.inters.Label = append(c.inters.Label, interceptors...)
}

// Create returns a builder for creating a Label entity.
func (c *LabelClient) Create() *LabelCreate {
	mutation := newLabelMutation(c.config, OpCreate)
	return &LabelCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Label entities.
func (c *LabelClient) CreateBulk(builders ...*LabelCreate) *LabelCreateBulk {
	return &LabelCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *LabelClient) MapCreateBulk(slice any, setFunc func(*LabelCreate, int)) *LabelCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &LabelCreateBulk{err: fmt.Errorf("calling to LabelClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*LabelCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &LabelCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Label.
func (c *LabelClient) Update() *LabelUpdate {
	mutation := newLabelMutation(c.config, OpUpdate)
	return &LabelUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LabelClient) UpdateOne(l *Label) *LabelUpdateOne {
	mutation := newLabelMutation(c.config, OpUpdateOne, withLabel(l))
	return &LabelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LabelClient) UpdateOneID(id int) *LabelUpdateOne {
	mutation := newLabelMutation(c.config, OpUpdateOne, withLabelID(id))
	return &LabelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Label.
func (c *LabelClient) Delete() *LabelDelete {
	mutation := newLabelMutation(c.config, OpDelete)
	return &LabelDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LabelClient) DeleteOne(l *Label) *LabelDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LabelClient) DeleteOneID(id int) *LabelDeleteOne {
	builder := c.Delete().Where(label.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LabelDeleteOne{builder}
}

// Query returns a query builder for Label.
func (c *LabelClient) Query() *LabelQuery {
	return &LabelQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLabel},
		inters: c.Interceptors(),
	}
}

// Get returns a Label entity by its id.
func (c *LabelClient) Get(ctx context.Context, id int) (*Label, error) {
	return c.Query().Where(label.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LabelClient) GetX(ctx context.Context, id int) *Label {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTodos queries the todos edge of a Label.
func (c *LabelClient) QueryTodos(l *Label) *TodoQuery {
	query := (&TodoClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := l.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(label.Table, label.FieldID, id),
			sqlgraph.To(todo.Table, todo.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, label.TodosTable, label.TodosPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(l.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LabelClient) Hooks() []Hook {
	return c.hooks.Label
}

// Interceptors returns the client interceptors.
func (c *LabelClient) Interceptors() []Interceptor {
	return c.inters.Label
}

func (c *LabelClient) mutate(ctx context.Context, m *LabelMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LabelCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LabelUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LabelUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LabelDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Label mutation op: %q", m.Op())
	}
}

// PriorityClient is a client for the Priority schema.
type PriorityClient struct {
	config
}

// NewPriorityClient returns a client for the Priority from the given config.
func NewPriorityClient(c config) *PriorityClient {
	return &PriorityClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `priority.Hooks(f(g(h())))`.
func (c *PriorityClient) Use(hooks ...Hook) {
	c.hooks.Priority = append(c.hooks.Priority, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `priority.Intercept(f(g(h())))`.
func (c *PriorityClient) Intercept(interceptors ...Interceptor) {
	c.inters.Priority = append(c.inters.Priority, interceptors...)
}

// Create returns a builder for creating a Priority entity.
func (c *PriorityClient) Create() *PriorityCreate {
	mutation := newPriorityMutation(c.config, OpCreate)
	return &PriorityCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Priority entities.
func (c *PriorityClient) CreateBulk(builders ...*PriorityCreate) *PriorityCreateBulk {
	return &PriorityCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PriorityClient) MapCreateBulk(slice any, setFunc func(*PriorityCreate, int)) *PriorityCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PriorityCreateBulk{err: fmt.Errorf("calling to PriorityClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PriorityCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PriorityCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Priority.
func (c *PriorityClient) Update() *PriorityUpdate {
	mutation := newPriorityMutation(c.config, OpUpdate)
	return &PriorityUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PriorityClient) UpdateOne(pr *Priority) *PriorityUpdateOne {
	mutation := newPriorityMutation(c.config, OpUpdateOne, withPriority(pr))
	return &PriorityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PriorityClient) UpdateOneID(id int) *PriorityUpdateOne {
	mutation := newPriorityMutation(c.config, OpUpdateOne, withPriorityID(id))
	return &PriorityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Priority.
func (c *PriorityClient) Delete() *PriorityDelete {
	mutation := newPriorityMutation(c.config, OpDelete)
	return &PriorityDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PriorityClient) DeleteOne(pr *Priority) *PriorityDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PriorityClient) DeleteOneID(id int) *PriorityDeleteOne {
	builder := c.Delete().Where(priority.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PriorityDeleteOne{builder}
}

// Query returns a query builder for Priority.
func (c *PriorityClient) Query() *PriorityQuery {
	return &PriorityQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePriority},
		inters: c.Interceptors(),
	}
}

// Get returns a Priority entity by its id.
func (c *PriorityClient) Get(ctx context.Context, id int) (*Priority, error) {
	return c.Query().Where(priority.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PriorityClient) GetX(ctx context.Context, id int) *Priority {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTodo queries the todo edge of a Priority.
func (c *PriorityClient) QueryTodo(pr *Priority) *TodoQuery {
	query := (&TodoClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(priority.Table, priority.FieldID, id),
			sqlgraph.To(todo.Table, todo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, priority.TodoTable, priority.TodoColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PriorityClient) Hooks() []Hook {
	return c.hooks.Priority
}

// Interceptors returns the client interceptors.
func (c *PriorityClient) Interceptors() []Interceptor {
	return c.inters.Priority
}

func (c *PriorityClient) mutate(ctx context.Context, m *PriorityMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PriorityCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PriorityUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PriorityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PriorityDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Priority mutation op: %q", m.Op())
	}
}

// StatusClient is a client for the Status schema.
type StatusClient struct {
	config
}

// NewStatusClient returns a client for the Status from the given config.
func NewStatusClient(c config) *StatusClient {
	return &StatusClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `status.Hooks(f(g(h())))`.
func (c *StatusClient) Use(hooks ...Hook) {
	c.hooks.Status = append(c.hooks.Status, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `status.Intercept(f(g(h())))`.
func (c *StatusClient) Intercept(interceptors ...Interceptor) {
	c.inters.Status = append(c.inters.Status, interceptors...)
}

// Create returns a builder for creating a Status entity.
func (c *StatusClient) Create() *StatusCreate {
	mutation := newStatusMutation(c.config, OpCreate)
	return &StatusCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Status entities.
func (c *StatusClient) CreateBulk(builders ...*StatusCreate) *StatusCreateBulk {
	return &StatusCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *StatusClient) MapCreateBulk(slice any, setFunc func(*StatusCreate, int)) *StatusCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &StatusCreateBulk{err: fmt.Errorf("calling to StatusClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*StatusCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &StatusCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Status.
func (c *StatusClient) Update() *StatusUpdate {
	mutation := newStatusMutation(c.config, OpUpdate)
	return &StatusUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StatusClient) UpdateOne(s *Status) *StatusUpdateOne {
	mutation := newStatusMutation(c.config, OpUpdateOne, withStatus(s))
	return &StatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StatusClient) UpdateOneID(id int) *StatusUpdateOne {
	mutation := newStatusMutation(c.config, OpUpdateOne, withStatusID(id))
	return &StatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Status.
func (c *StatusClient) Delete() *StatusDelete {
	mutation := newStatusMutation(c.config, OpDelete)
	return &StatusDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StatusClient) DeleteOne(s *Status) *StatusDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *StatusClient) DeleteOneID(id int) *StatusDeleteOne {
	builder := c.Delete().Where(status.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StatusDeleteOne{builder}
}

// Query returns a query builder for Status.
func (c *StatusClient) Query() *StatusQuery {
	return &StatusQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeStatus},
		inters: c.Interceptors(),
	}
}

// Get returns a Status entity by its id.
func (c *StatusClient) Get(ctx context.Context, id int) (*Status, error) {
	return c.Query().Where(status.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StatusClient) GetX(ctx context.Context, id int) *Status {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTodo queries the todo edge of a Status.
func (c *StatusClient) QueryTodo(s *Status) *TodoQuery {
	query := (&TodoClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(status.Table, status.FieldID, id),
			sqlgraph.To(todo.Table, todo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, status.TodoTable, status.TodoColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StatusClient) Hooks() []Hook {
	return c.hooks.Status
}

// Interceptors returns the client interceptors.
func (c *StatusClient) Interceptors() []Interceptor {
	return c.inters.Status
}

func (c *StatusClient) mutate(ctx context.Context, m *StatusMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&StatusCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&StatusUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&StatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&StatusDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Status mutation op: %q", m.Op())
	}
}

// TodoClient is a client for the Todo schema.
type TodoClient struct {
	config
}

// NewTodoClient returns a client for the Todo from the given config.
func NewTodoClient(c config) *TodoClient {
	return &TodoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `todo.Hooks(f(g(h())))`.
func (c *TodoClient) Use(hooks ...Hook) {
	c.hooks.Todo = append(c.hooks.Todo, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `todo.Intercept(f(g(h())))`.
func (c *TodoClient) Intercept(interceptors ...Interceptor) {
	c.inters.Todo = append(c.inters.Todo, interceptors...)
}

// Create returns a builder for creating a Todo entity.
func (c *TodoClient) Create() *TodoCreate {
	mutation := newTodoMutation(c.config, OpCreate)
	return &TodoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Todo entities.
func (c *TodoClient) CreateBulk(builders ...*TodoCreate) *TodoCreateBulk {
	return &TodoCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TodoClient) MapCreateBulk(slice any, setFunc func(*TodoCreate, int)) *TodoCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TodoCreateBulk{err: fmt.Errorf("calling to TodoClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TodoCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TodoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Todo.
func (c *TodoClient) Update() *TodoUpdate {
	mutation := newTodoMutation(c.config, OpUpdate)
	return &TodoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TodoClient) UpdateOne(t *Todo) *TodoUpdateOne {
	mutation := newTodoMutation(c.config, OpUpdateOne, withTodo(t))
	return &TodoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TodoClient) UpdateOneID(id int) *TodoUpdateOne {
	mutation := newTodoMutation(c.config, OpUpdateOne, withTodoID(id))
	return &TodoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Todo.
func (c *TodoClient) Delete() *TodoDelete {
	mutation := newTodoMutation(c.config, OpDelete)
	return &TodoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TodoClient) DeleteOne(t *Todo) *TodoDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TodoClient) DeleteOneID(id int) *TodoDeleteOne {
	builder := c.Delete().Where(todo.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TodoDeleteOne{builder}
}

// Query returns a query builder for Todo.
func (c *TodoClient) Query() *TodoQuery {
	return &TodoQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTodo},
		inters: c.Interceptors(),
	}
}

// Get returns a Todo entity by its id.
func (c *TodoClient) Get(ctx context.Context, id int) (*Todo, error) {
	return c.Query().Where(todo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TodoClient) GetX(ctx context.Context, id int) *Todo {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPriority queries the priority edge of a Todo.
func (c *TodoClient) QueryPriority(t *Todo) *PriorityQuery {
	query := (&PriorityClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(todo.Table, todo.FieldID, id),
			sqlgraph.To(priority.Table, priority.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, todo.PriorityTable, todo.PriorityColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStatus queries the status edge of a Todo.
func (c *TodoClient) QueryStatus(t *Todo) *StatusQuery {
	query := (&StatusClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(todo.Table, todo.FieldID, id),
			sqlgraph.To(status.Table, status.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, todo.StatusTable, todo.StatusColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLabels queries the labels edge of a Todo.
func (c *TodoClient) QueryLabels(t *Todo) *LabelQuery {
	query := (&LabelClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(todo.Table, todo.FieldID, id),
			sqlgraph.To(label.Table, label.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, todo.LabelsTable, todo.LabelsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TodoClient) Hooks() []Hook {
	return c.hooks.Todo
}

// Interceptors returns the client interceptors.
func (c *TodoClient) Interceptors() []Interceptor {
	return c.inters.Todo
}

func (c *TodoClient) mutate(ctx context.Context, m *TodoMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TodoCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TodoUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TodoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TodoDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Todo mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Label, Priority, Status, Todo []ent.Hook
	}
	inters struct {
		Label, Priority, Status, Todo []ent.Interceptor
	}
)
