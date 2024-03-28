// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"4th_Assignment/ent/migrate"

	"4th_Assignment/ent/contactsubmission"
	"4th_Assignment/ent/schemamigration"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ContactSubmission is the client for interacting with the ContactSubmission builders.
	ContactSubmission *ContactSubmissionClient
	// SchemaMigration is the client for interacting with the SchemaMigration builders.
	SchemaMigration *SchemaMigrationClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.ContactSubmission = NewContactSubmissionClient(c.config)
	c.SchemaMigration = NewSchemaMigrationClient(c.config)
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
		ctx:               ctx,
		config:            cfg,
		ContactSubmission: NewContactSubmissionClient(cfg),
		SchemaMigration:   NewSchemaMigrationClient(cfg),
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
		ctx:               ctx,
		config:            cfg,
		ContactSubmission: NewContactSubmissionClient(cfg),
		SchemaMigration:   NewSchemaMigrationClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ContactSubmission.
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
	c.ContactSubmission.Use(hooks...)
	c.SchemaMigration.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.ContactSubmission.Intercept(interceptors...)
	c.SchemaMigration.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ContactSubmissionMutation:
		return c.ContactSubmission.mutate(ctx, m)
	case *SchemaMigrationMutation:
		return c.SchemaMigration.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ContactSubmissionClient is a client for the ContactSubmission schema.
type ContactSubmissionClient struct {
	config
}

// NewContactSubmissionClient returns a client for the ContactSubmission from the given config.
func NewContactSubmissionClient(c config) *ContactSubmissionClient {
	return &ContactSubmissionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `contactsubmission.Hooks(f(g(h())))`.
func (c *ContactSubmissionClient) Use(hooks ...Hook) {
	c.hooks.ContactSubmission = append(c.hooks.ContactSubmission, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `contactsubmission.Intercept(f(g(h())))`.
func (c *ContactSubmissionClient) Intercept(interceptors ...Interceptor) {
	c.inters.ContactSubmission = append(c.inters.ContactSubmission, interceptors...)
}

// Create returns a builder for creating a ContactSubmission entity.
func (c *ContactSubmissionClient) Create() *ContactSubmissionCreate {
	mutation := newContactSubmissionMutation(c.config, OpCreate)
	return &ContactSubmissionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ContactSubmission entities.
func (c *ContactSubmissionClient) CreateBulk(builders ...*ContactSubmissionCreate) *ContactSubmissionCreateBulk {
	return &ContactSubmissionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ContactSubmissionClient) MapCreateBulk(slice any, setFunc func(*ContactSubmissionCreate, int)) *ContactSubmissionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ContactSubmissionCreateBulk{err: fmt.Errorf("calling to ContactSubmissionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ContactSubmissionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ContactSubmissionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ContactSubmission.
func (c *ContactSubmissionClient) Update() *ContactSubmissionUpdate {
	mutation := newContactSubmissionMutation(c.config, OpUpdate)
	return &ContactSubmissionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ContactSubmissionClient) UpdateOne(cs *ContactSubmission) *ContactSubmissionUpdateOne {
	mutation := newContactSubmissionMutation(c.config, OpUpdateOne, withContactSubmission(cs))
	return &ContactSubmissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ContactSubmissionClient) UpdateOneID(id int) *ContactSubmissionUpdateOne {
	mutation := newContactSubmissionMutation(c.config, OpUpdateOne, withContactSubmissionID(id))
	return &ContactSubmissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ContactSubmission.
func (c *ContactSubmissionClient) Delete() *ContactSubmissionDelete {
	mutation := newContactSubmissionMutation(c.config, OpDelete)
	return &ContactSubmissionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ContactSubmissionClient) DeleteOne(cs *ContactSubmission) *ContactSubmissionDeleteOne {
	return c.DeleteOneID(cs.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ContactSubmissionClient) DeleteOneID(id int) *ContactSubmissionDeleteOne {
	builder := c.Delete().Where(contactsubmission.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ContactSubmissionDeleteOne{builder}
}

// Query returns a query builder for ContactSubmission.
func (c *ContactSubmissionClient) Query() *ContactSubmissionQuery {
	return &ContactSubmissionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeContactSubmission},
		inters: c.Interceptors(),
	}
}

// Get returns a ContactSubmission entity by its id.
func (c *ContactSubmissionClient) Get(ctx context.Context, id int) (*ContactSubmission, error) {
	return c.Query().Where(contactsubmission.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ContactSubmissionClient) GetX(ctx context.Context, id int) *ContactSubmission {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ContactSubmissionClient) Hooks() []Hook {
	return c.hooks.ContactSubmission
}

// Interceptors returns the client interceptors.
func (c *ContactSubmissionClient) Interceptors() []Interceptor {
	return c.inters.ContactSubmission
}

func (c *ContactSubmissionClient) mutate(ctx context.Context, m *ContactSubmissionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ContactSubmissionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ContactSubmissionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ContactSubmissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ContactSubmissionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ContactSubmission mutation op: %q", m.Op())
	}
}

// SchemaMigrationClient is a client for the SchemaMigration schema.
type SchemaMigrationClient struct {
	config
}

// NewSchemaMigrationClient returns a client for the SchemaMigration from the given config.
func NewSchemaMigrationClient(c config) *SchemaMigrationClient {
	return &SchemaMigrationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `schemamigration.Hooks(f(g(h())))`.
func (c *SchemaMigrationClient) Use(hooks ...Hook) {
	c.hooks.SchemaMigration = append(c.hooks.SchemaMigration, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `schemamigration.Intercept(f(g(h())))`.
func (c *SchemaMigrationClient) Intercept(interceptors ...Interceptor) {
	c.inters.SchemaMigration = append(c.inters.SchemaMigration, interceptors...)
}

// Create returns a builder for creating a SchemaMigration entity.
func (c *SchemaMigrationClient) Create() *SchemaMigrationCreate {
	mutation := newSchemaMigrationMutation(c.config, OpCreate)
	return &SchemaMigrationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SchemaMigration entities.
func (c *SchemaMigrationClient) CreateBulk(builders ...*SchemaMigrationCreate) *SchemaMigrationCreateBulk {
	return &SchemaMigrationCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SchemaMigrationClient) MapCreateBulk(slice any, setFunc func(*SchemaMigrationCreate, int)) *SchemaMigrationCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SchemaMigrationCreateBulk{err: fmt.Errorf("calling to SchemaMigrationClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SchemaMigrationCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SchemaMigrationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SchemaMigration.
func (c *SchemaMigrationClient) Update() *SchemaMigrationUpdate {
	mutation := newSchemaMigrationMutation(c.config, OpUpdate)
	return &SchemaMigrationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SchemaMigrationClient) UpdateOne(sm *SchemaMigration) *SchemaMigrationUpdateOne {
	mutation := newSchemaMigrationMutation(c.config, OpUpdateOne, withSchemaMigration(sm))
	return &SchemaMigrationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SchemaMigrationClient) UpdateOneID(id int) *SchemaMigrationUpdateOne {
	mutation := newSchemaMigrationMutation(c.config, OpUpdateOne, withSchemaMigrationID(id))
	return &SchemaMigrationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SchemaMigration.
func (c *SchemaMigrationClient) Delete() *SchemaMigrationDelete {
	mutation := newSchemaMigrationMutation(c.config, OpDelete)
	return &SchemaMigrationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SchemaMigrationClient) DeleteOne(sm *SchemaMigration) *SchemaMigrationDeleteOne {
	return c.DeleteOneID(sm.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SchemaMigrationClient) DeleteOneID(id int) *SchemaMigrationDeleteOne {
	builder := c.Delete().Where(schemamigration.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SchemaMigrationDeleteOne{builder}
}

// Query returns a query builder for SchemaMigration.
func (c *SchemaMigrationClient) Query() *SchemaMigrationQuery {
	return &SchemaMigrationQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSchemaMigration},
		inters: c.Interceptors(),
	}
}

// Get returns a SchemaMigration entity by its id.
func (c *SchemaMigrationClient) Get(ctx context.Context, id int) (*SchemaMigration, error) {
	return c.Query().Where(schemamigration.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SchemaMigrationClient) GetX(ctx context.Context, id int) *SchemaMigration {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SchemaMigrationClient) Hooks() []Hook {
	return c.hooks.SchemaMigration
}

// Interceptors returns the client interceptors.
func (c *SchemaMigrationClient) Interceptors() []Interceptor {
	return c.inters.SchemaMigration
}

func (c *SchemaMigrationClient) mutate(ctx context.Context, m *SchemaMigrationMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SchemaMigrationCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SchemaMigrationUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SchemaMigrationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SchemaMigrationDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown SchemaMigration mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		ContactSubmission, SchemaMigration []ent.Hook
	}
	inters struct {
		ContactSubmission, SchemaMigration []ent.Interceptor
	}
)
