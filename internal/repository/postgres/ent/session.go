// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/imranzahaev/warehouse/internal/repository/postgres/ent/session"
	"github.com/imranzahaev/warehouse/internal/repository/postgres/ent/user"
)

// Session is the model entity for the Session schema.
type Session struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AccessToken holds the value of the "access_token" field.
	AccessToken string `json:"access_token,omitempty"`
	// RefreshToken holds the value of the "refresh_token" field.
	RefreshToken string `json:"refresh_token,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Disabled holds the value of the "disabled" field.
	Disabled bool `json:"disabled,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SessionQuery when eager-loading is set.
	Edges         SessionEdges `json:"edges"`
	user_sessions *int
}

// SessionEdges holds the relations/edges for other nodes in the graph.
type SessionEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SessionEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Session) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case session.FieldDisabled:
			values[i] = new(sql.NullBool)
		case session.FieldID:
			values[i] = new(sql.NullInt64)
		case session.FieldAccessToken, session.FieldRefreshToken:
			values[i] = new(sql.NullString)
		case session.FieldUpdatedAt, session.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case session.ForeignKeys[0]: // user_sessions
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Session", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Session fields.
func (s *Session) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case session.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case session.FieldAccessToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_token", values[i])
			} else if value.Valid {
				s.AccessToken = value.String
			}
		case session.FieldRefreshToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refresh_token", values[i])
			} else if value.Valid {
				s.RefreshToken = value.String
			}
		case session.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case session.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case session.FieldDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				s.Disabled = value.Bool
			}
		case session.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_sessions", value)
			} else if value.Valid {
				s.user_sessions = new(int)
				*s.user_sessions = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Session entity.
func (s *Session) QueryOwner() *UserQuery {
	return (&SessionClient{config: s.config}).QueryOwner(s)
}

// Update returns a builder for updating this Session.
// Note that you need to call Session.Unwrap() before calling this method if this Session
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Session) Update() *SessionUpdateOne {
	return (&SessionClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Session entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Session) Unwrap() *Session {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Session is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Session) String() string {
	var builder strings.Builder
	builder.WriteString("Session(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", access_token=")
	builder.WriteString(s.AccessToken)
	builder.WriteString(", refresh_token=")
	builder.WriteString(s.RefreshToken)
	builder.WriteString(", updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", disabled=")
	builder.WriteString(fmt.Sprintf("%v", s.Disabled))
	builder.WriteByte(')')
	return builder.String()
}

// Sessions is a parsable slice of Session.
type Sessions []*Session

func (s Sessions) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}