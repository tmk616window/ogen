// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// SearchGet implements GET /search operation.
//
// Search for todo items.
//
// GET /search
func (UnimplementedHandler) SearchGet(ctx context.Context) (r *ResponseSearchTodo, _ error) {
	return r, ht.ErrNotImplemented
}

// TodoIDDelete implements DELETE /todo/{id} operation.
//
// Delete a todo item by ID.
//
// DELETE /todo/{id}
func (UnimplementedHandler) TodoIDDelete(ctx context.Context, params TodoIDDeleteParams) (r int, _ error) {
	return r, ht.ErrNotImplemented
}

// TodoIDPut implements PUT /todo/{id} operation.
//
// Update a todo item by ID.
//
// PUT /todo/{id}
func (UnimplementedHandler) TodoIDPut(ctx context.Context, req *WhereTodoInput, params TodoIDPutParams) (r int, _ error) {
	return r, ht.ErrNotImplemented
}

// TodoPost implements POST /todo operation.
//
// Create a new todo item.
//
// POST /todo
func (UnimplementedHandler) TodoPost(ctx context.Context, req *CreateTodoInput) (r *CreateTodoResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// TodosGet implements GET /todos operation.
//
// Get all todo items.
//
// GET /todos
func (UnimplementedHandler) TodosGet(ctx context.Context, params TodosGetParams) (r *TodosGetOK, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorResponseStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorResponseStatusCode) {
	r = new(ErrorResponseStatusCode)
	return r
}
