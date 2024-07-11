// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Pineapple217/Sortify/web/ent/playlist"
	"github.com/Pineapple217/Sortify/web/ent/track"
	"github.com/Pineapple217/Sortify/web/ent/user"
)

// PlaylistCreate is the builder for creating a Playlist entity.
type PlaylistCreate struct {
	config
	mutation *PlaylistMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (pc *PlaylistCreate) SetName(s string) *PlaylistCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PlaylistCreate) SetCreatedAt(t time.Time) *PlaylistCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PlaylistCreate) SetNillableCreatedAt(t *time.Time) *PlaylistCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PlaylistCreate) SetUpdatedAt(t time.Time) *PlaylistCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PlaylistCreate) SetNillableUpdatedAt(t *time.Time) *PlaylistCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetDeletedAt sets the "deleted_at" field.
func (pc *PlaylistCreate) SetDeletedAt(t time.Time) *PlaylistCreate {
	pc.mutation.SetDeletedAt(t)
	return pc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pc *PlaylistCreate) SetNillableDeletedAt(t *time.Time) *PlaylistCreate {
	if t != nil {
		pc.SetDeletedAt(*t)
	}
	return pc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pc *PlaylistCreate) SetUserID(id int) *PlaylistCreate {
	pc.mutation.SetUserID(id)
	return pc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pc *PlaylistCreate) SetNillableUserID(id *int) *PlaylistCreate {
	if id != nil {
		pc = pc.SetUserID(*id)
	}
	return pc
}

// SetUser sets the "user" edge to the User entity.
func (pc *PlaylistCreate) SetUser(u *User) *PlaylistCreate {
	return pc.SetUserID(u.ID)
}

// AddTrackIDs adds the "tracks" edge to the Track entity by IDs.
func (pc *PlaylistCreate) AddTrackIDs(ids ...int) *PlaylistCreate {
	pc.mutation.AddTrackIDs(ids...)
	return pc
}

// AddTracks adds the "tracks" edges to the Track entity.
func (pc *PlaylistCreate) AddTracks(t ...*Track) *PlaylistCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pc.AddTrackIDs(ids...)
}

// Mutation returns the PlaylistMutation object of the builder.
func (pc *PlaylistCreate) Mutation() *PlaylistMutation {
	return pc.mutation
}

// Save creates the Playlist in the database.
func (pc *PlaylistCreate) Save(ctx context.Context) (*Playlist, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PlaylistCreate) SaveX(ctx context.Context) *Playlist {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PlaylistCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PlaylistCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PlaylistCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := playlist.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := playlist.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PlaylistCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Playlist.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := playlist.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Playlist.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Playlist.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Playlist.updated_at"`)}
	}
	return nil
}

func (pc *PlaylistCreate) sqlSave(ctx context.Context) (*Playlist, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PlaylistCreate) createSpec() (*Playlist, *sqlgraph.CreateSpec) {
	var (
		_node = &Playlist{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(playlist.Table, sqlgraph.NewFieldSpec(playlist.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pc.conflict
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(playlist.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(playlist.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(playlist.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.DeletedAt(); ok {
		_spec.SetField(playlist.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlist.UserTable,
			Columns: []string{playlist.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_playlists = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.TracksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   playlist.TracksTable,
			Columns: playlist.TracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Playlist.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlaylistUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (pc *PlaylistCreate) OnConflict(opts ...sql.ConflictOption) *PlaylistUpsertOne {
	pc.conflict = opts
	return &PlaylistUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Playlist.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PlaylistCreate) OnConflictColumns(columns ...string) *PlaylistUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PlaylistUpsertOne{
		create: pc,
	}
}

type (
	// PlaylistUpsertOne is the builder for "upsert"-ing
	//  one Playlist node.
	PlaylistUpsertOne struct {
		create *PlaylistCreate
	}

	// PlaylistUpsert is the "OnConflict" setter.
	PlaylistUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *PlaylistUpsert) SetName(v string) *PlaylistUpsert {
	u.Set(playlist.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PlaylistUpsert) UpdateName() *PlaylistUpsert {
	u.SetExcluded(playlist.FieldName)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PlaylistUpsert) SetUpdatedAt(v time.Time) *PlaylistUpsert {
	u.Set(playlist.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PlaylistUpsert) UpdateUpdatedAt() *PlaylistUpsert {
	u.SetExcluded(playlist.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PlaylistUpsert) SetDeletedAt(v time.Time) *PlaylistUpsert {
	u.Set(playlist.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PlaylistUpsert) UpdateDeletedAt() *PlaylistUpsert {
	u.SetExcluded(playlist.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *PlaylistUpsert) ClearDeletedAt() *PlaylistUpsert {
	u.SetNull(playlist.FieldDeletedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Playlist.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PlaylistUpsertOne) UpdateNewValues() *PlaylistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(playlist.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Playlist.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PlaylistUpsertOne) Ignore() *PlaylistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlaylistUpsertOne) DoNothing() *PlaylistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlaylistCreate.OnConflict
// documentation for more info.
func (u *PlaylistUpsertOne) Update(set func(*PlaylistUpsert)) *PlaylistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlaylistUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *PlaylistUpsertOne) SetName(v string) *PlaylistUpsertOne {
	return u.Update(func(s *PlaylistUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PlaylistUpsertOne) UpdateName() *PlaylistUpsertOne {
	return u.Update(func(s *PlaylistUpsert) {
		s.UpdateName()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PlaylistUpsertOne) SetUpdatedAt(v time.Time) *PlaylistUpsertOne {
	return u.Update(func(s *PlaylistUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PlaylistUpsertOne) UpdateUpdatedAt() *PlaylistUpsertOne {
	return u.Update(func(s *PlaylistUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PlaylistUpsertOne) SetDeletedAt(v time.Time) *PlaylistUpsertOne {
	return u.Update(func(s *PlaylistUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PlaylistUpsertOne) UpdateDeletedAt() *PlaylistUpsertOne {
	return u.Update(func(s *PlaylistUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *PlaylistUpsertOne) ClearDeletedAt() *PlaylistUpsertOne {
	return u.Update(func(s *PlaylistUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *PlaylistUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlaylistCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlaylistUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PlaylistUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PlaylistUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PlaylistCreateBulk is the builder for creating many Playlist entities in bulk.
type PlaylistCreateBulk struct {
	config
	err      error
	builders []*PlaylistCreate
	conflict []sql.ConflictOption
}

// Save creates the Playlist entities in the database.
func (pcb *PlaylistCreateBulk) Save(ctx context.Context) ([]*Playlist, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Playlist, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlaylistMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PlaylistCreateBulk) SaveX(ctx context.Context) []*Playlist {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PlaylistCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PlaylistCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Playlist.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlaylistUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (pcb *PlaylistCreateBulk) OnConflict(opts ...sql.ConflictOption) *PlaylistUpsertBulk {
	pcb.conflict = opts
	return &PlaylistUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Playlist.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PlaylistCreateBulk) OnConflictColumns(columns ...string) *PlaylistUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PlaylistUpsertBulk{
		create: pcb,
	}
}

// PlaylistUpsertBulk is the builder for "upsert"-ing
// a bulk of Playlist nodes.
type PlaylistUpsertBulk struct {
	create *PlaylistCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Playlist.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PlaylistUpsertBulk) UpdateNewValues() *PlaylistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(playlist.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Playlist.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PlaylistUpsertBulk) Ignore() *PlaylistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlaylistUpsertBulk) DoNothing() *PlaylistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlaylistCreateBulk.OnConflict
// documentation for more info.
func (u *PlaylistUpsertBulk) Update(set func(*PlaylistUpsert)) *PlaylistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlaylistUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *PlaylistUpsertBulk) SetName(v string) *PlaylistUpsertBulk {
	return u.Update(func(s *PlaylistUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PlaylistUpsertBulk) UpdateName() *PlaylistUpsertBulk {
	return u.Update(func(s *PlaylistUpsert) {
		s.UpdateName()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PlaylistUpsertBulk) SetUpdatedAt(v time.Time) *PlaylistUpsertBulk {
	return u.Update(func(s *PlaylistUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PlaylistUpsertBulk) UpdateUpdatedAt() *PlaylistUpsertBulk {
	return u.Update(func(s *PlaylistUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PlaylistUpsertBulk) SetDeletedAt(v time.Time) *PlaylistUpsertBulk {
	return u.Update(func(s *PlaylistUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PlaylistUpsertBulk) UpdateDeletedAt() *PlaylistUpsertBulk {
	return u.Update(func(s *PlaylistUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *PlaylistUpsertBulk) ClearDeletedAt() *PlaylistUpsertBulk {
	return u.Update(func(s *PlaylistUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *PlaylistUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PlaylistCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlaylistCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlaylistUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
