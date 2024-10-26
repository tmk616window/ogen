// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PrioritiesColumns holds the columns for the "priorities" table.
	PrioritiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// PrioritiesTable holds the schema information for the "priorities" table.
	PrioritiesTable = &schema.Table{
		Name:       "priorities",
		Columns:    PrioritiesColumns,
		PrimaryKey: []*schema.Column{PrioritiesColumns[0]},
	}
	// StatusColumns holds the columns for the "status" table.
	StatusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// StatusTable holds the schema information for the "status" table.
	StatusTable = &schema.Table{
		Name:       "status",
		Columns:    StatusColumns,
		PrimaryKey: []*schema.Column{StatusColumns[0]},
	}
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "finished_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "priority_id", Type: field.TypeInt},
		{Name: "status_id", Type: field.TypeInt, Default: 1},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "todos_priorities_todo",
				Columns:    []*schema.Column{TodosColumns[7]},
				RefColumns: []*schema.Column{PrioritiesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "todos_status_todo",
				Columns:    []*schema.Column{TodosColumns[8]},
				RefColumns: []*schema.Column{StatusColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PrioritiesTable,
		StatusTable,
		TodosTable,
	}
)

func init() {
	TodosTable.ForeignKeys[0].RefTable = PrioritiesTable
	TodosTable.ForeignKeys[1].RefTable = StatusTable
}
