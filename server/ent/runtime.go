// Code generated by ent, DO NOT EDIT.

package ent

import (
	"server/ent/schema"
	"server/ent/todo"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescTitle is the schema descriptor for title field.
	todoDescTitle := todoFields[1].Descriptor()
	// todo.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	todo.TitleValidator = todoDescTitle.Validators[0].(func(string) error)
	// todoDescName is the schema descriptor for name field.
	todoDescName := todoFields[3].Descriptor()
	// todo.NameValidator is a validator for the "name" field. It is called by the builders before save.
	todo.NameValidator = todoDescName.Validators[0].(func(string) error)
}