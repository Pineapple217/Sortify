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
	"github.com/Pineapple217/Sortify/web/ent/predicate"
	"github.com/Pineapple217/Sortify/web/ent/track"
	"github.com/Pineapple217/Sortify/web/ent/user"
)

// PlaylistUpdate is the builder for updating Playlist entities.
type PlaylistUpdate struct {
	config
	hooks    []Hook
	mutation *PlaylistMutation
}

// Where appends a list predicates to the PlaylistUpdate builder.
func (pu *PlaylistUpdate) Where(ps ...predicate.Playlist) *PlaylistUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PlaylistUpdate) SetName(s string) *PlaylistUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PlaylistUpdate) SetNillableName(s *string) *PlaylistUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PlaylistUpdate) SetUpdatedAt(t time.Time) *PlaylistUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetDeletedAt sets the "deleted_at" field.
func (pu *PlaylistUpdate) SetDeletedAt(t time.Time) *PlaylistUpdate {
	pu.mutation.SetDeletedAt(t)
	return pu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pu *PlaylistUpdate) SetNillableDeletedAt(t *time.Time) *PlaylistUpdate {
	if t != nil {
		pu.SetDeletedAt(*t)
	}
	return pu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pu *PlaylistUpdate) ClearDeletedAt() *PlaylistUpdate {
	pu.mutation.ClearDeletedAt()
	return pu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pu *PlaylistUpdate) SetUserID(id int) *PlaylistUpdate {
	pu.mutation.SetUserID(id)
	return pu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pu *PlaylistUpdate) SetNillableUserID(id *int) *PlaylistUpdate {
	if id != nil {
		pu = pu.SetUserID(*id)
	}
	return pu
}

// SetUser sets the "user" edge to the User entity.
func (pu *PlaylistUpdate) SetUser(u *User) *PlaylistUpdate {
	return pu.SetUserID(u.ID)
}

// AddTrackIDs adds the "tracks" edge to the Track entity by IDs.
func (pu *PlaylistUpdate) AddTrackIDs(ids ...int) *PlaylistUpdate {
	pu.mutation.AddTrackIDs(ids...)
	return pu
}

// AddTracks adds the "tracks" edges to the Track entity.
func (pu *PlaylistUpdate) AddTracks(t ...*Track) *PlaylistUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.AddTrackIDs(ids...)
}

// Mutation returns the PlaylistMutation object of the builder.
func (pu *PlaylistUpdate) Mutation() *PlaylistMutation {
	return pu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pu *PlaylistUpdate) ClearUser() *PlaylistUpdate {
	pu.mutation.ClearUser()
	return pu
}

// ClearTracks clears all "tracks" edges to the Track entity.
func (pu *PlaylistUpdate) ClearTracks() *PlaylistUpdate {
	pu.mutation.ClearTracks()
	return pu
}

// RemoveTrackIDs removes the "tracks" edge to Track entities by IDs.
func (pu *PlaylistUpdate) RemoveTrackIDs(ids ...int) *PlaylistUpdate {
	pu.mutation.RemoveTrackIDs(ids...)
	return pu
}

// RemoveTracks removes "tracks" edges to Track entities.
func (pu *PlaylistUpdate) RemoveTracks(t ...*Track) *PlaylistUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.RemoveTrackIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PlaylistUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PlaylistUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PlaylistUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PlaylistUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PlaylistUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := playlist.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PlaylistUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := playlist.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Playlist.name": %w`, err)}
		}
	}
	return nil
}

func (pu *PlaylistUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(playlist.Table, playlist.Columns, sqlgraph.NewFieldSpec(playlist.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(playlist.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(playlist.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.DeletedAt(); ok {
		_spec.SetField(playlist.FieldDeletedAt, field.TypeTime, value)
	}
	if pu.mutation.DeletedAtCleared() {
		_spec.ClearField(playlist.FieldDeletedAt, field.TypeTime)
	}
	if pu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.TracksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedTracksIDs(); len(nodes) > 0 && !pu.mutation.TracksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TracksIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{playlist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PlaylistUpdateOne is the builder for updating a single Playlist entity.
type PlaylistUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlaylistMutation
}

// SetName sets the "name" field.
func (puo *PlaylistUpdateOne) SetName(s string) *PlaylistUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PlaylistUpdateOne) SetNillableName(s *string) *PlaylistUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PlaylistUpdateOne) SetUpdatedAt(t time.Time) *PlaylistUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetDeletedAt sets the "deleted_at" field.
func (puo *PlaylistUpdateOne) SetDeletedAt(t time.Time) *PlaylistUpdateOne {
	puo.mutation.SetDeletedAt(t)
	return puo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (puo *PlaylistUpdateOne) SetNillableDeletedAt(t *time.Time) *PlaylistUpdateOne {
	if t != nil {
		puo.SetDeletedAt(*t)
	}
	return puo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (puo *PlaylistUpdateOne) ClearDeletedAt() *PlaylistUpdateOne {
	puo.mutation.ClearDeletedAt()
	return puo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (puo *PlaylistUpdateOne) SetUserID(id int) *PlaylistUpdateOne {
	puo.mutation.SetUserID(id)
	return puo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (puo *PlaylistUpdateOne) SetNillableUserID(id *int) *PlaylistUpdateOne {
	if id != nil {
		puo = puo.SetUserID(*id)
	}
	return puo
}

// SetUser sets the "user" edge to the User entity.
func (puo *PlaylistUpdateOne) SetUser(u *User) *PlaylistUpdateOne {
	return puo.SetUserID(u.ID)
}

// AddTrackIDs adds the "tracks" edge to the Track entity by IDs.
func (puo *PlaylistUpdateOne) AddTrackIDs(ids ...int) *PlaylistUpdateOne {
	puo.mutation.AddTrackIDs(ids...)
	return puo
}

// AddTracks adds the "tracks" edges to the Track entity.
func (puo *PlaylistUpdateOne) AddTracks(t ...*Track) *PlaylistUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.AddTrackIDs(ids...)
}

// Mutation returns the PlaylistMutation object of the builder.
func (puo *PlaylistUpdateOne) Mutation() *PlaylistMutation {
	return puo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (puo *PlaylistUpdateOne) ClearUser() *PlaylistUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// ClearTracks clears all "tracks" edges to the Track entity.
func (puo *PlaylistUpdateOne) ClearTracks() *PlaylistUpdateOne {
	puo.mutation.ClearTracks()
	return puo
}

// RemoveTrackIDs removes the "tracks" edge to Track entities by IDs.
func (puo *PlaylistUpdateOne) RemoveTrackIDs(ids ...int) *PlaylistUpdateOne {
	puo.mutation.RemoveTrackIDs(ids...)
	return puo
}

// RemoveTracks removes "tracks" edges to Track entities.
func (puo *PlaylistUpdateOne) RemoveTracks(t ...*Track) *PlaylistUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.RemoveTrackIDs(ids...)
}

// Where appends a list predicates to the PlaylistUpdate builder.
func (puo *PlaylistUpdateOne) Where(ps ...predicate.Playlist) *PlaylistUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PlaylistUpdateOne) Select(field string, fields ...string) *PlaylistUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Playlist entity.
func (puo *PlaylistUpdateOne) Save(ctx context.Context) (*Playlist, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PlaylistUpdateOne) SaveX(ctx context.Context) *Playlist {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PlaylistUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PlaylistUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PlaylistUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := playlist.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PlaylistUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := playlist.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Playlist.name": %w`, err)}
		}
	}
	return nil
}

func (puo *PlaylistUpdateOne) sqlSave(ctx context.Context) (_node *Playlist, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(playlist.Table, playlist.Columns, sqlgraph.NewFieldSpec(playlist.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Playlist.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, playlist.FieldID)
		for _, f := range fields {
			if !playlist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != playlist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(playlist.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(playlist.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.DeletedAt(); ok {
		_spec.SetField(playlist.FieldDeletedAt, field.TypeTime, value)
	}
	if puo.mutation.DeletedAtCleared() {
		_spec.ClearField(playlist.FieldDeletedAt, field.TypeTime)
	}
	if puo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.TracksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedTracksIDs(); len(nodes) > 0 && !puo.mutation.TracksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TracksIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Playlist{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{playlist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
