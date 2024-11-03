// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"fmt"
	"time"
)

func (s *ErrorResponseStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Ref: #/components/schemas/CreateTodoInput
type CreateTodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	LabelIDs    []int  `json:"labelIDs"`
	PriorityID  int    `json:"priorityID"`
	StatusID    int    `json:"statusID"`
}

// GetTitle returns the value of Title.
func (s *CreateTodoInput) GetTitle() string {
	return s.Title
}

// GetDescription returns the value of Description.
func (s *CreateTodoInput) GetDescription() string {
	return s.Description
}

// GetLabelIDs returns the value of LabelIDs.
func (s *CreateTodoInput) GetLabelIDs() []int {
	return s.LabelIDs
}

// GetPriorityID returns the value of PriorityID.
func (s *CreateTodoInput) GetPriorityID() int {
	return s.PriorityID
}

// GetStatusID returns the value of StatusID.
func (s *CreateTodoInput) GetStatusID() int {
	return s.StatusID
}

// SetTitle sets the value of Title.
func (s *CreateTodoInput) SetTitle(val string) {
	s.Title = val
}

// SetDescription sets the value of Description.
func (s *CreateTodoInput) SetDescription(val string) {
	s.Description = val
}

// SetLabelIDs sets the value of LabelIDs.
func (s *CreateTodoInput) SetLabelIDs(val []int) {
	s.LabelIDs = val
}

// SetPriorityID sets the value of PriorityID.
func (s *CreateTodoInput) SetPriorityID(val int) {
	s.PriorityID = val
}

// SetStatusID sets the value of StatusID.
func (s *CreateTodoInput) SetStatusID(val int) {
	s.StatusID = val
}

// Ref: #/components/schemas/CreateTodoResponse
type CreateTodoResponse struct {
	ID int `json:"id"`
}

// GetID returns the value of ID.
func (s *CreateTodoResponse) GetID() int {
	return s.ID
}

// SetID sets the value of ID.
func (s *CreateTodoResponse) SetID(val int) {
	s.ID = val
}

// Ref: #/components/schemas/ErrorResponse
type ErrorResponse struct {
	Error string `json:"error"`
}

// GetError returns the value of Error.
func (s *ErrorResponse) GetError() string {
	return s.Error
}

// SetError sets the value of Error.
func (s *ErrorResponse) SetError(val string) {
	s.Error = val
}

// ErrorResponseStatusCode wraps ErrorResponse with StatusCode.
type ErrorResponseStatusCode struct {
	StatusCode int
	Response   ErrorResponse
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorResponseStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorResponseStatusCode) GetResponse() ErrorResponse {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorResponseStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorResponseStatusCode) SetResponse(val ErrorResponse) {
	s.Response = val
}

// Ref: #/components/schemas/Label
type Label struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

// GetID returns the value of ID.
func (s *Label) GetID() int {
	return s.ID
}

// GetValue returns the value of Value.
func (s *Label) GetValue() string {
	return s.Value
}

// SetID sets the value of ID.
func (s *Label) SetID(val int) {
	s.ID = val
}

// SetValue sets the value of Value.
func (s *Label) SetValue(val string) {
	s.Value = val
}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Priority
type Priority struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetID returns the value of ID.
func (s *Priority) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *Priority) GetName() string {
	return s.Name
}

// SetID sets the value of ID.
func (s *Priority) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Priority) SetName(val string) {
	s.Name = val
}

// Ref: #/components/schemas/ResponseSearchTodo
type ResponseSearchTodo struct {
	Labels     []Label    `json:"labels"`
	Priorities []Priority `json:"priorities"`
	Status     []Status   `json:"status"`
}

// GetLabels returns the value of Labels.
func (s *ResponseSearchTodo) GetLabels() []Label {
	return s.Labels
}

// GetPriorities returns the value of Priorities.
func (s *ResponseSearchTodo) GetPriorities() []Priority {
	return s.Priorities
}

// GetStatus returns the value of Status.
func (s *ResponseSearchTodo) GetStatus() []Status {
	return s.Status
}

// SetLabels sets the value of Labels.
func (s *ResponseSearchTodo) SetLabels(val []Label) {
	s.Labels = val
}

// SetPriorities sets the value of Priorities.
func (s *ResponseSearchTodo) SetPriorities(val []Priority) {
	s.Priorities = val
}

// SetStatus sets the value of Status.
func (s *ResponseSearchTodo) SetStatus(val []Status) {
	s.Status = val
}

// Ref: #/components/schemas/Status
type Status struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

// GetID returns the value of ID.
func (s *Status) GetID() int {
	return s.ID
}

// GetValue returns the value of Value.
func (s *Status) GetValue() string {
	return s.Value
}

// SetID sets the value of ID.
func (s *Status) SetID(val int) {
	s.ID = val
}

// SetValue sets the value of Value.
func (s *Status) SetValue(val string) {
	s.Value = val
}

// Ref: #/components/schemas/Todo
type Todo struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description OptString   `json:"description"`
	CreatedAt   time.Time   `json:"createdAt"`
	Labels      []Label     `json:"labels"`
	FinishedAt  OptDateTime `json:"finishedAt"`
	Priority    Priority    `json:"priority"`
	Status      Status      `json:"status"`
}

// GetID returns the value of ID.
func (s *Todo) GetID() int {
	return s.ID
}

// GetTitle returns the value of Title.
func (s *Todo) GetTitle() string {
	return s.Title
}

// GetDescription returns the value of Description.
func (s *Todo) GetDescription() OptString {
	return s.Description
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Todo) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// GetLabels returns the value of Labels.
func (s *Todo) GetLabels() []Label {
	return s.Labels
}

// GetFinishedAt returns the value of FinishedAt.
func (s *Todo) GetFinishedAt() OptDateTime {
	return s.FinishedAt
}

// GetPriority returns the value of Priority.
func (s *Todo) GetPriority() Priority {
	return s.Priority
}

// GetStatus returns the value of Status.
func (s *Todo) GetStatus() Status {
	return s.Status
}

// SetID sets the value of ID.
func (s *Todo) SetID(val int) {
	s.ID = val
}

// SetTitle sets the value of Title.
func (s *Todo) SetTitle(val string) {
	s.Title = val
}

// SetDescription sets the value of Description.
func (s *Todo) SetDescription(val OptString) {
	s.Description = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Todo) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

// SetLabels sets the value of Labels.
func (s *Todo) SetLabels(val []Label) {
	s.Labels = val
}

// SetFinishedAt sets the value of FinishedAt.
func (s *Todo) SetFinishedAt(val OptDateTime) {
	s.FinishedAt = val
}

// SetPriority sets the value of Priority.
func (s *Todo) SetPriority(val Priority) {
	s.Priority = val
}

// SetStatus sets the value of Status.
func (s *Todo) SetStatus(val Status) {
	s.Status = val
}
