// Code generated by ogen, DO NOT EDIT.

package ogen

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
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

// Ref: #/components/schemas/Todo
type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description OptString `json:"description"`
	IsCompleted OptBool   `json:"isCompleted"`
}

// GetID returns the value of ID.
func (s *Todo) GetID() string {
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

// GetIsCompleted returns the value of IsCompleted.
func (s *Todo) GetIsCompleted() OptBool {
	return s.IsCompleted
}

// SetID sets the value of ID.
func (s *Todo) SetID(val string) {
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

// SetIsCompleted sets the value of IsCompleted.
func (s *Todo) SetIsCompleted(val OptBool) {
	s.IsCompleted = val
}

// Ref: #/components/schemas/TodoInput
type TodoInput struct {
	Title       string    `json:"title"`
	Description OptString `json:"description"`
	IsCompleted OptBool   `json:"isCompleted"`
}

// GetTitle returns the value of Title.
func (s *TodoInput) GetTitle() string {
	return s.Title
}

// GetDescription returns the value of Description.
func (s *TodoInput) GetDescription() OptString {
	return s.Description
}

// GetIsCompleted returns the value of IsCompleted.
func (s *TodoInput) GetIsCompleted() OptBool {
	return s.IsCompleted
}

// SetTitle sets the value of Title.
func (s *TodoInput) SetTitle(val string) {
	s.Title = val
}

// SetDescription sets the value of Description.
func (s *TodoInput) SetDescription(val OptString) {
	s.Description = val
}

// SetIsCompleted sets the value of IsCompleted.
func (s *TodoInput) SetIsCompleted(val OptBool) {
	s.IsCompleted = val
}
