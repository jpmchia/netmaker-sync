// Package optional provides primitives to represent optional values.
package optional

// Interface represents an optional interface value
type Interface struct {
	value interface{}
	set   bool
}

// NewInterface creates a new optional interface value
func NewInterface(value interface{}) Interface {
	return Interface{
		value: value,
		set:   true,
	}
}

// IsSet returns whether the value has been set
func (i Interface) IsSet() bool {
	return i.set
}

// Value returns the value
func (i Interface) Value() interface{} {
	return i.value
}

// Reset resets the value
func (i *Interface) Reset() {
	i.value = nil
	i.set = false
}

// String represents an optional string value
type String struct {
	value string
	set   bool
}

// NewString creates a new optional string value
func NewString(value string) String {
	return String{
		value: value,
		set:   true,
	}
}

// IsSet returns whether the value has been set
func (s String) IsSet() bool {
	return s.set
}

// Value returns the value
func (s String) Value() string {
	return s.value
}

// Reset resets the value
func (s *String) Reset() {
	s.value = ""
	s.set = false
}

// Int represents an optional int value
type Int struct {
	value int
	set   bool
}

// NewInt creates a new optional int value
func NewInt(value int) Int {
	return Int{
		value: value,
		set:   true,
	}
}

// IsSet returns whether the value has been set
func (i Int) IsSet() bool {
	return i.set
}

// Value returns the value
func (i Int) Value() int {
	return i.value
}

// Reset resets the value
func (i *Int) Reset() {
	i.value = 0
	i.set = false
}

// Bool represents an optional bool value
type Bool struct {
	value bool
	set   bool
}

// NewBool creates a new optional bool value
func NewBool(value bool) Bool {
	return Bool{
		value: value,
		set:   true,
	}
}

// IsSet returns whether the value has been set
func (b Bool) IsSet() bool {
	return b.set
}

// Value returns the value
func (b Bool) Value() bool {
	return b.value
}

// Reset resets the value
func (b *Bool) Reset() {
	b.value = false
	b.set = false
}

// Float32 represents an optional float32 value
type Float32 struct {
	value float32
	set   bool
}

// NewFloat32 creates a new optional float32 value
func NewFloat32(value float32) Float32 {
	return Float32{
		value: value,
		set:   true,
	}
}

// IsSet returns whether the value has been set
func (f Float32) IsSet() bool {
	return f.set
}

// Value returns the value
func (f Float32) Value() float32 {
	return f.value
}

// Reset resets the value
func (f *Float32) Reset() {
	f.value = 0
	f.set = false
}

// Float64 represents an optional float64 value
type Float64 struct {
	value float64
	set   bool
}

// NewFloat64 creates a new optional float64 value
func NewFloat64(value float64) Float64 {
	return Float64{
		value: value,
		set:   true,
	}
}

// IsSet returns whether the value has been set
func (f Float64) IsSet() bool {
	return f.set
}

// Value returns the value
func (f Float64) Value() float64 {
	return f.value
}

// Reset resets the value
func (f *Float64) Reset() {
	f.value = 0
	f.set = false
}
