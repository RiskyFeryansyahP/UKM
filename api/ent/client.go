// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/confus1on/UKM/ent/migrate"

	"github.com/confus1on/UKM/ent/profile"
	"github.com/confus1on/UKM/ent/profileukm"
	"github.com/confus1on/UKM/ent/role"
	"github.com/confus1on/UKM/ent/roleukm"
	"github.com/confus1on/UKM/ent/ukm"
	"github.com/confus1on/UKM/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Profile is the client for interacting with the Profile builders.
	Profile *ProfileClient
	// ProfileUKM is the client for interacting with the ProfileUKM builders.
	ProfileUKM *ProfileUKMClient
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
	// RoleUKM is the client for interacting with the RoleUKM builders.
	RoleUKM *RoleUKMClient
	// Ukm is the client for interacting with the Ukm builders.
	Ukm *UkmClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Profile = NewProfileClient(c.config)
	c.ProfileUKM = NewProfileUKMClient(c.config)
	c.Role = NewRoleClient(c.config)
	c.RoleUKM = NewRoleUKMClient(c.config)
	c.Ukm = NewUkmClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
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

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:     cfg,
		Profile:    NewProfileClient(cfg),
		ProfileUKM: NewProfileUKMClient(cfg),
		Role:       NewRoleClient(cfg),
		RoleUKM:    NewRoleUKMClient(cfg),
		Ukm:        NewUkmClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Profile.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
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
	c.Profile.Use(hooks...)
	c.ProfileUKM.Use(hooks...)
	c.Role.Use(hooks...)
	c.RoleUKM.Use(hooks...)
	c.Ukm.Use(hooks...)
	c.User.Use(hooks...)
}

// ProfileClient is a client for the Profile schema.
type ProfileClient struct {
	config
}

// NewProfileClient returns a client for the Profile from the given config.
func NewProfileClient(c config) *ProfileClient {
	return &ProfileClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `profile.Hooks(f(g(h())))`.
func (c *ProfileClient) Use(hooks ...Hook) {
	c.hooks.Profile = append(c.hooks.Profile, hooks...)
}

// Create returns a create builder for Profile.
func (c *ProfileClient) Create() *ProfileCreate {
	mutation := newProfileMutation(c.config, OpCreate)
	return &ProfileCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Profile.
func (c *ProfileClient) Update() *ProfileUpdate {
	mutation := newProfileMutation(c.config, OpUpdate)
	return &ProfileUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProfileClient) UpdateOne(pr *Profile) *ProfileUpdateOne {
	return c.UpdateOneID(pr.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *ProfileClient) UpdateOneID(id int) *ProfileUpdateOne {
	mutation := newProfileMutation(c.config, OpUpdateOne)
	mutation.id = &id
	return &ProfileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Profile.
func (c *ProfileClient) Delete() *ProfileDelete {
	mutation := newProfileMutation(c.config, OpDelete)
	return &ProfileDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProfileClient) DeleteOne(pr *Profile) *ProfileDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProfileClient) DeleteOneID(id int) *ProfileDeleteOne {
	builder := c.Delete().Where(profile.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProfileDeleteOne{builder}
}

// Create returns a query builder for Profile.
func (c *ProfileClient) Query() *ProfileQuery {
	return &ProfileQuery{config: c.config}
}

// Get returns a Profile entity by its id.
func (c *ProfileClient) Get(ctx context.Context, id int) (*Profile, error) {
	return c.Query().Where(profile.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProfileClient) GetX(ctx context.Context, id int) *Profile {
	pr, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pr
}

// QueryOwner queries the owner edge of a Profile.
func (c *ProfileClient) QueryOwner(pr *Profile) *UserQuery {
	query := &UserQuery{config: c.config}
	id := pr.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(profile.Table, profile.FieldID, id),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, profile.OwnerTable, profile.OwnerColumn),
	)
	query.sql = sqlgraph.Neighbors(pr.driver.Dialect(), step)

	return query
}

// QueryUkms queries the ukms edge of a Profile.
func (c *ProfileClient) QueryUkms(pr *Profile) *ProfileUKMQuery {
	query := &ProfileUKMQuery{config: c.config}
	id := pr.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(profile.Table, profile.FieldID, id),
		sqlgraph.To(profileukm.Table, profileukm.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, profile.UkmsTable, profile.UkmsColumn),
	)
	query.sql = sqlgraph.Neighbors(pr.driver.Dialect(), step)

	return query
}

// Hooks returns the client hooks.
func (c *ProfileClient) Hooks() []Hook {
	return c.hooks.Profile
}

// ProfileUKMClient is a client for the ProfileUKM schema.
type ProfileUKMClient struct {
	config
}

// NewProfileUKMClient returns a client for the ProfileUKM from the given config.
func NewProfileUKMClient(c config) *ProfileUKMClient {
	return &ProfileUKMClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `profileukm.Hooks(f(g(h())))`.
func (c *ProfileUKMClient) Use(hooks ...Hook) {
	c.hooks.ProfileUKM = append(c.hooks.ProfileUKM, hooks...)
}

// Create returns a create builder for ProfileUKM.
func (c *ProfileUKMClient) Create() *ProfileUKMCreate {
	mutation := newProfileUKMMutation(c.config, OpCreate)
	return &ProfileUKMCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for ProfileUKM.
func (c *ProfileUKMClient) Update() *ProfileUKMUpdate {
	mutation := newProfileUKMMutation(c.config, OpUpdate)
	return &ProfileUKMUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProfileUKMClient) UpdateOne(pu *ProfileUKM) *ProfileUKMUpdateOne {
	return c.UpdateOneID(pu.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *ProfileUKMClient) UpdateOneID(id int) *ProfileUKMUpdateOne {
	mutation := newProfileUKMMutation(c.config, OpUpdateOne)
	mutation.id = &id
	return &ProfileUKMUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ProfileUKM.
func (c *ProfileUKMClient) Delete() *ProfileUKMDelete {
	mutation := newProfileUKMMutation(c.config, OpDelete)
	return &ProfileUKMDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProfileUKMClient) DeleteOne(pu *ProfileUKM) *ProfileUKMDeleteOne {
	return c.DeleteOneID(pu.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProfileUKMClient) DeleteOneID(id int) *ProfileUKMDeleteOne {
	builder := c.Delete().Where(profileukm.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProfileUKMDeleteOne{builder}
}

// Create returns a query builder for ProfileUKM.
func (c *ProfileUKMClient) Query() *ProfileUKMQuery {
	return &ProfileUKMQuery{config: c.config}
}

// Get returns a ProfileUKM entity by its id.
func (c *ProfileUKMClient) Get(ctx context.Context, id int) (*ProfileUKM, error) {
	return c.Query().Where(profileukm.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProfileUKMClient) GetX(ctx context.Context, id int) *ProfileUKM {
	pu, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pu
}

// QueryOwnerProfile queries the owner_profile edge of a ProfileUKM.
func (c *ProfileUKMClient) QueryOwnerProfile(pu *ProfileUKM) *ProfileQuery {
	query := &ProfileQuery{config: c.config}
	id := pu.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(profileukm.Table, profileukm.FieldID, id),
		sqlgraph.To(profile.Table, profile.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, profileukm.OwnerProfileTable, profileukm.OwnerProfileColumn),
	)
	query.sql = sqlgraph.Neighbors(pu.driver.Dialect(), step)

	return query
}

// QueryOwnerUkm queries the owner_ukm edge of a ProfileUKM.
func (c *ProfileUKMClient) QueryOwnerUkm(pu *ProfileUKM) *UkmQuery {
	query := &UkmQuery{config: c.config}
	id := pu.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(profileukm.Table, profileukm.FieldID, id),
		sqlgraph.To(ukm.Table, ukm.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, profileukm.OwnerUkmTable, profileukm.OwnerUkmColumn),
	)
	query.sql = sqlgraph.Neighbors(pu.driver.Dialect(), step)

	return query
}

// QueryOwnerRole queries the owner_role edge of a ProfileUKM.
func (c *ProfileUKMClient) QueryOwnerRole(pu *ProfileUKM) *RoleUKMQuery {
	query := &RoleUKMQuery{config: c.config}
	id := pu.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(profileukm.Table, profileukm.FieldID, id),
		sqlgraph.To(roleukm.Table, roleukm.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, profileukm.OwnerRoleTable, profileukm.OwnerRoleColumn),
	)
	query.sql = sqlgraph.Neighbors(pu.driver.Dialect(), step)

	return query
}

// Hooks returns the client hooks.
func (c *ProfileUKMClient) Hooks() []Hook {
	return c.hooks.ProfileUKM
}

// RoleClient is a client for the Role schema.
type RoleClient struct {
	config
}

// NewRoleClient returns a client for the Role from the given config.
func NewRoleClient(c config) *RoleClient {
	return &RoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `role.Hooks(f(g(h())))`.
func (c *RoleClient) Use(hooks ...Hook) {
	c.hooks.Role = append(c.hooks.Role, hooks...)
}

// Create returns a create builder for Role.
func (c *RoleClient) Create() *RoleCreate {
	mutation := newRoleMutation(c.config, OpCreate)
	return &RoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Role.
func (c *RoleClient) Update() *RoleUpdate {
	mutation := newRoleMutation(c.config, OpUpdate)
	return &RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoleClient) UpdateOne(r *Role) *RoleUpdateOne {
	return c.UpdateOneID(r.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *RoleClient) UpdateOneID(id int) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne)
	mutation.id = &id
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Role.
func (c *RoleClient) Delete() *RoleDelete {
	mutation := newRoleMutation(c.config, OpDelete)
	return &RoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RoleClient) DeleteOne(r *Role) *RoleDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RoleClient) DeleteOneID(id int) *RoleDeleteOne {
	builder := c.Delete().Where(role.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoleDeleteOne{builder}
}

// Create returns a query builder for Role.
func (c *RoleClient) Query() *RoleQuery {
	return &RoleQuery{config: c.config}
}

// Get returns a Role entity by its id.
func (c *RoleClient) Get(ctx context.Context, id int) (*Role, error) {
	return c.Query().Where(role.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoleClient) GetX(ctx context.Context, id int) *Role {
	r, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return r
}

// QueryAccess queries the access edge of a Role.
func (c *RoleClient) QueryAccess(r *Role) *UserQuery {
	query := &UserQuery{config: c.config}
	id := r.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(role.Table, role.FieldID, id),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, role.AccessTable, role.AccessColumn),
	)
	query.sql = sqlgraph.Neighbors(r.driver.Dialect(), step)

	return query
}

// Hooks returns the client hooks.
func (c *RoleClient) Hooks() []Hook {
	return c.hooks.Role
}

// RoleUKMClient is a client for the RoleUKM schema.
type RoleUKMClient struct {
	config
}

// NewRoleUKMClient returns a client for the RoleUKM from the given config.
func NewRoleUKMClient(c config) *RoleUKMClient {
	return &RoleUKMClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `roleukm.Hooks(f(g(h())))`.
func (c *RoleUKMClient) Use(hooks ...Hook) {
	c.hooks.RoleUKM = append(c.hooks.RoleUKM, hooks...)
}

// Create returns a create builder for RoleUKM.
func (c *RoleUKMClient) Create() *RoleUKMCreate {
	mutation := newRoleUKMMutation(c.config, OpCreate)
	return &RoleUKMCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for RoleUKM.
func (c *RoleUKMClient) Update() *RoleUKMUpdate {
	mutation := newRoleUKMMutation(c.config, OpUpdate)
	return &RoleUKMUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoleUKMClient) UpdateOne(ru *RoleUKM) *RoleUKMUpdateOne {
	return c.UpdateOneID(ru.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *RoleUKMClient) UpdateOneID(id int) *RoleUKMUpdateOne {
	mutation := newRoleUKMMutation(c.config, OpUpdateOne)
	mutation.id = &id
	return &RoleUKMUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for RoleUKM.
func (c *RoleUKMClient) Delete() *RoleUKMDelete {
	mutation := newRoleUKMMutation(c.config, OpDelete)
	return &RoleUKMDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RoleUKMClient) DeleteOne(ru *RoleUKM) *RoleUKMDeleteOne {
	return c.DeleteOneID(ru.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RoleUKMClient) DeleteOneID(id int) *RoleUKMDeleteOne {
	builder := c.Delete().Where(roleukm.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoleUKMDeleteOne{builder}
}

// Create returns a query builder for RoleUKM.
func (c *RoleUKMClient) Query() *RoleUKMQuery {
	return &RoleUKMQuery{config: c.config}
}

// Get returns a RoleUKM entity by its id.
func (c *RoleUKMClient) Get(ctx context.Context, id int) (*RoleUKM, error) {
	return c.Query().Where(roleukm.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoleUKMClient) GetX(ctx context.Context, id int) *RoleUKM {
	ru, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ru
}

// QueryProfileRoles queries the profile_roles edge of a RoleUKM.
func (c *RoleUKMClient) QueryProfileRoles(ru *RoleUKM) *ProfileUKMQuery {
	query := &ProfileUKMQuery{config: c.config}
	id := ru.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(roleukm.Table, roleukm.FieldID, id),
		sqlgraph.To(profileukm.Table, profileukm.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, roleukm.ProfileRolesTable, roleukm.ProfileRolesColumn),
	)
	query.sql = sqlgraph.Neighbors(ru.driver.Dialect(), step)

	return query
}

// Hooks returns the client hooks.
func (c *RoleUKMClient) Hooks() []Hook {
	return c.hooks.RoleUKM
}

// UkmClient is a client for the Ukm schema.
type UkmClient struct {
	config
}

// NewUkmClient returns a client for the Ukm from the given config.
func NewUkmClient(c config) *UkmClient {
	return &UkmClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ukm.Hooks(f(g(h())))`.
func (c *UkmClient) Use(hooks ...Hook) {
	c.hooks.Ukm = append(c.hooks.Ukm, hooks...)
}

// Create returns a create builder for Ukm.
func (c *UkmClient) Create() *UkmCreate {
	mutation := newUkmMutation(c.config, OpCreate)
	return &UkmCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Ukm.
func (c *UkmClient) Update() *UkmUpdate {
	mutation := newUkmMutation(c.config, OpUpdate)
	return &UkmUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UkmClient) UpdateOne(u *Ukm) *UkmUpdateOne {
	return c.UpdateOneID(u.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *UkmClient) UpdateOneID(id int) *UkmUpdateOne {
	mutation := newUkmMutation(c.config, OpUpdateOne)
	mutation.id = &id
	return &UkmUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Ukm.
func (c *UkmClient) Delete() *UkmDelete {
	mutation := newUkmMutation(c.config, OpDelete)
	return &UkmDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UkmClient) DeleteOne(u *Ukm) *UkmDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UkmClient) DeleteOneID(id int) *UkmDeleteOne {
	builder := c.Delete().Where(ukm.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UkmDeleteOne{builder}
}

// Create returns a query builder for Ukm.
func (c *UkmClient) Query() *UkmQuery {
	return &UkmQuery{config: c.config}
}

// Get returns a Ukm entity by its id.
func (c *UkmClient) Get(ctx context.Context, id int) (*Ukm, error) {
	return c.Query().Where(ukm.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UkmClient) GetX(ctx context.Context, id int) *Ukm {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryProfiles queries the profiles edge of a Ukm.
func (c *UkmClient) QueryProfiles(u *Ukm) *ProfileUKMQuery {
	query := &ProfileUKMQuery{config: c.config}
	id := u.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(ukm.Table, ukm.FieldID, id),
		sqlgraph.To(profileukm.Table, profileukm.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ukm.ProfilesTable, ukm.ProfilesColumn),
	)
	query.sql = sqlgraph.Neighbors(u.driver.Dialect(), step)

	return query
}

// Hooks returns the client hooks.
func (c *UkmClient) Hooks() []Hook {
	return c.hooks.Ukm
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	return c.UpdateOneID(u.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne)
	mutation.id = &id
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryProfile queries the profile edge of a User.
func (c *UserClient) QueryProfile(u *User) *ProfileQuery {
	query := &ProfileQuery{config: c.config}
	id := u.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(user.Table, user.FieldID, id),
		sqlgraph.To(profile.Table, profile.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, user.ProfileTable, user.ProfileColumn),
	)
	query.sql = sqlgraph.Neighbors(u.driver.Dialect(), step)

	return query
}

// QueryRole queries the role edge of a User.
func (c *UserClient) QueryRole(u *User) *RoleQuery {
	query := &RoleQuery{config: c.config}
	id := u.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(user.Table, user.FieldID, id),
		sqlgraph.To(role.Table, role.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, user.RoleTable, user.RoleColumn),
	)
	query.sql = sqlgraph.Neighbors(u.driver.Dialect(), step)

	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
