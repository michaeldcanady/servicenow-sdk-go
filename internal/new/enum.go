package internal

// Enum represents a generic enumeration interface.
type Enum[T any] interface {
	GetValue() T            // Returns the value of the current enum instance.
	GetValues() []Enum[T]   // Returns all possible values of the enum.
	GetNames() []string     // Returns all possible names of the enum.
	GetName() string        // Returns the name of the current enum value.
	IsDefined(Enum[T]) bool // Checks if the given value is defined in the enum.
	String() string         // Returns the string representation of the enum value.
}

// FlagEnum represents a bitwise enumeration interface with generics.
type FlagEnum[T any] interface {
	Enum[T]
	HasFlag(Enum[T]) bool // Checks if the flag is set in the current enum value.
}

// EnumFactory creates a new enum value of type T.
type EnumFactory[T any] func(T) Enum[T]

// TryEnumFactory tries to create a new enum value of type T, returning true if successful.
type TryEnumFactory[T any] func(T, Enum[T]) bool
