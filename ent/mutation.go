// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/201R/go_api_boilerplate/ent/predicate"
	"github.com/201R/go_api_boilerplate/ent/user"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	firstName     *string
	lastName      *string
	email         *string
	location      *string
	phone         *string
	password      *string
	status        *bool
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetFirstName sets the "firstName" field.
func (m *UserMutation) SetFirstName(s string) {
	m.firstName = &s
}

// FirstName returns the value of the "firstName" field in the mutation.
func (m *UserMutation) FirstName() (r string, exists bool) {
	v := m.firstName
	if v == nil {
		return
	}
	return *v, true
}

// OldFirstName returns the old "firstName" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldFirstName(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFirstName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFirstName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFirstName: %w", err)
	}
	return oldValue.FirstName, nil
}

// ClearFirstName clears the value of the "firstName" field.
func (m *UserMutation) ClearFirstName() {
	m.firstName = nil
	m.clearedFields[user.FieldFirstName] = struct{}{}
}

// FirstNameCleared returns if the "firstName" field was cleared in this mutation.
func (m *UserMutation) FirstNameCleared() bool {
	_, ok := m.clearedFields[user.FieldFirstName]
	return ok
}

// ResetFirstName resets all changes to the "firstName" field.
func (m *UserMutation) ResetFirstName() {
	m.firstName = nil
	delete(m.clearedFields, user.FieldFirstName)
}

// SetLastName sets the "lastName" field.
func (m *UserMutation) SetLastName(s string) {
	m.lastName = &s
}

// LastName returns the value of the "lastName" field in the mutation.
func (m *UserMutation) LastName() (r string, exists bool) {
	v := m.lastName
	if v == nil {
		return
	}
	return *v, true
}

// OldLastName returns the old "lastName" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldLastName(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastName: %w", err)
	}
	return oldValue.LastName, nil
}

// ClearLastName clears the value of the "lastName" field.
func (m *UserMutation) ClearLastName() {
	m.lastName = nil
	m.clearedFields[user.FieldLastName] = struct{}{}
}

// LastNameCleared returns if the "lastName" field was cleared in this mutation.
func (m *UserMutation) LastNameCleared() bool {
	_, ok := m.clearedFields[user.FieldLastName]
	return ok
}

// ResetLastName resets all changes to the "lastName" field.
func (m *UserMutation) ResetLastName() {
	m.lastName = nil
	delete(m.clearedFields, user.FieldLastName)
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetLocation sets the "location" field.
func (m *UserMutation) SetLocation(s string) {
	m.location = &s
}

// Location returns the value of the "location" field in the mutation.
func (m *UserMutation) Location() (r string, exists bool) {
	v := m.location
	if v == nil {
		return
	}
	return *v, true
}

// OldLocation returns the old "location" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldLocation(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLocation is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLocation requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLocation: %w", err)
	}
	return oldValue.Location, nil
}

// ClearLocation clears the value of the "location" field.
func (m *UserMutation) ClearLocation() {
	m.location = nil
	m.clearedFields[user.FieldLocation] = struct{}{}
}

// LocationCleared returns if the "location" field was cleared in this mutation.
func (m *UserMutation) LocationCleared() bool {
	_, ok := m.clearedFields[user.FieldLocation]
	return ok
}

// ResetLocation resets all changes to the "location" field.
func (m *UserMutation) ResetLocation() {
	m.location = nil
	delete(m.clearedFields, user.FieldLocation)
}

// SetPhone sets the "phone" field.
func (m *UserMutation) SetPhone(s string) {
	m.phone = &s
}

// Phone returns the value of the "phone" field in the mutation.
func (m *UserMutation) Phone() (r string, exists bool) {
	v := m.phone
	if v == nil {
		return
	}
	return *v, true
}

// OldPhone returns the old "phone" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPhone(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPhone is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPhone requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPhone: %w", err)
	}
	return oldValue.Phone, nil
}

// ClearPhone clears the value of the "phone" field.
func (m *UserMutation) ClearPhone() {
	m.phone = nil
	m.clearedFields[user.FieldPhone] = struct{}{}
}

// PhoneCleared returns if the "phone" field was cleared in this mutation.
func (m *UserMutation) PhoneCleared() bool {
	_, ok := m.clearedFields[user.FieldPhone]
	return ok
}

// ResetPhone resets all changes to the "phone" field.
func (m *UserMutation) ResetPhone() {
	m.phone = nil
	delete(m.clearedFields, user.FieldPhone)
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// SetStatus sets the "status" field.
func (m *UserMutation) SetStatus(b bool) {
	m.status = &b
}

// Status returns the value of the "status" field in the mutation.
func (m *UserMutation) Status() (r bool, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldStatus(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *UserMutation) ResetStatus() {
	m.status = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.firstName != nil {
		fields = append(fields, user.FieldFirstName)
	}
	if m.lastName != nil {
		fields = append(fields, user.FieldLastName)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.location != nil {
		fields = append(fields, user.FieldLocation)
	}
	if m.phone != nil {
		fields = append(fields, user.FieldPhone)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.status != nil {
		fields = append(fields, user.FieldStatus)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldFirstName:
		return m.FirstName()
	case user.FieldLastName:
		return m.LastName()
	case user.FieldEmail:
		return m.Email()
	case user.FieldLocation:
		return m.Location()
	case user.FieldPhone:
		return m.Phone()
	case user.FieldPassword:
		return m.Password()
	case user.FieldStatus:
		return m.Status()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldFirstName:
		return m.OldFirstName(ctx)
	case user.FieldLastName:
		return m.OldLastName(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldLocation:
		return m.OldLocation(ctx)
	case user.FieldPhone:
		return m.OldPhone(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldStatus:
		return m.OldStatus(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldFirstName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFirstName(v)
		return nil
	case user.FieldLastName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastName(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldLocation:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLocation(v)
		return nil
	case user.FieldPhone:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPhone(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldStatus:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldFirstName) {
		fields = append(fields, user.FieldFirstName)
	}
	if m.FieldCleared(user.FieldLastName) {
		fields = append(fields, user.FieldLastName)
	}
	if m.FieldCleared(user.FieldLocation) {
		fields = append(fields, user.FieldLocation)
	}
	if m.FieldCleared(user.FieldPhone) {
		fields = append(fields, user.FieldPhone)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldFirstName:
		m.ClearFirstName()
		return nil
	case user.FieldLastName:
		m.ClearLastName()
		return nil
	case user.FieldLocation:
		m.ClearLocation()
		return nil
	case user.FieldPhone:
		m.ClearPhone()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldFirstName:
		m.ResetFirstName()
		return nil
	case user.FieldLastName:
		m.ResetLastName()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldLocation:
		m.ResetLocation()
		return nil
	case user.FieldPhone:
		m.ResetPhone()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldStatus:
		m.ResetStatus()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
