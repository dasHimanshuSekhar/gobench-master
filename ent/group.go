// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/gobench-io/gobench/ent/application"
	"github.com/gobench-io/gobench/ent/group"
)

// Group is the model entity for the Group schema.
type Group struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupQuery when eager-loading is set.
	Edges              GroupEdges `json:"edges"`
	application_groups *int
}

// GroupEdges holds the relations/edges for other nodes in the graph.
type GroupEdges struct {
	// Application holds the value of the application edge.
	Application *Application
	// Graphs holds the value of the graphs edge.
	Graphs []*Graph
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ApplicationOrErr returns the Application value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupEdges) ApplicationOrErr() (*Application, error) {
	if e.loadedTypes[0] {
		if e.Application == nil {
			// The edge application was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: application.Label}
		}
		return e.Application, nil
	}
	return nil, &NotLoadedError{edge: "application"}
}

// GraphsOrErr returns the Graphs value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) GraphsOrErr() ([]*Graph, error) {
	if e.loadedTypes[1] {
		return e.Graphs, nil
	}
	return nil, &NotLoadedError{edge: "graphs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Group) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Group) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // application_groups
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Group fields.
func (gr *Group) assignValues(values ...interface{}) error {
	if m, n := len(values), len(group.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	gr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		gr.Name = value.String
	}
	values = values[1:]
	if len(values) == len(group.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field application_groups", value)
		} else if value.Valid {
			gr.application_groups = new(int)
			*gr.application_groups = int(value.Int64)
		}
	}
	return nil
}

// QueryApplication queries the application edge of the Group.
func (gr *Group) QueryApplication() *ApplicationQuery {
	return (&GroupClient{config: gr.config}).QueryApplication(gr)
}

// QueryGraphs queries the graphs edge of the Group.
func (gr *Group) QueryGraphs() *GraphQuery {
	return (&GroupClient{config: gr.config}).QueryGraphs(gr)
}

// Update returns a builder for updating this Group.
// Note that, you need to call Group.Unwrap() before calling this method, if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return (&GroupClient{config: gr.config}).UpdateOne(gr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Group is not a transactional entity")
	}
	gr.config.driver = tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v", gr.ID))
	builder.WriteString(", name=")
	builder.WriteString(gr.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Groups is a parsable slice of Group.
type Groups []*Group

func (gr Groups) config(cfg config) {
	for _i := range gr {
		gr[_i].config = cfg
	}
}
