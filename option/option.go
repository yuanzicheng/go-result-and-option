package option

type Option[T any] struct {
	value *T
}

func New[T any](v *T) *Option[T] {
	if v == nil {
		return &Option[T]{}
	}
	return &Option[T]{value: v}
}

func Some[T any](v *T) *Option[T] {
	return &Option[T]{value: v}
}

func None[T any]() *Option[T] {
	return &Option[T]{}
}

// Map maps an `Option[T]` to `Option[U]` by applying a function to a contained value (if `Some`) or returns `None` (if `None`).
func Map[T any, U any](o *Option[T], f func(*T) *U) *Option[U] {
	if o.value == nil {
		return &Option[U]{}
	}
	return &Option[U]{value: f(o.value)}
}

// MapOr returns the provided fallback result (if none), or applies a function to the contained value (if any).
func MapOr[T any, U any](o *Option[T], fallback *U, f func(*T) *U) *U {
	if o.value == nil {
		return fallback
	}
	return f(o.value)
}

// MapOrElse computes a default function result (if none), or applies a different function to the contained value (if any).
func MapOrElse[T any, U any](o *Option[T], fallbackFn func() *U, f func(*T) *U) *U {
	if o.value == nil {
		return fallbackFn()
	}
	return f(o.value)
}

// And returns [`None`] if the option is [`None`], otherwise returns `optb`.
func And[T any, U any](in *Option[T], out *Option[U]) *Option[U] {
	if in.value == nil {
		return nil
	}
	return out
}

// AndThen returns [`None`] if the option is [`None`], otherwise calls `f` with the wrapped value and returns the result.
func AndThen[T any, U any](in *Option[T], f func(*T) *Option[U]) *Option[U] {
	if in.value == nil {
		return nil
	}
	return f(in.value)
}

// IsSomeAnd returns `true` if the option is a [`Some`].
func (o *Option[T]) IsSome() bool {
	return o.value != nil
}

// IsSomeAnd returns `true` if the option is a [`Some`] and the value inside of it matches a predicate.
func (o *Option[T]) IsSomeAnd(f func(*T) bool) bool {
	return o.value != nil && f(o.value)
}

// IsSomeAnd returns `true` if the option is a [`None`].
func (o *Option[T]) IsNone() bool {
	return o.value == nil
}

// Expect returns the contained [`Some`] value, consuming the `self` value.
// Panics if the value is a [`None`] with a custom panic message provided by `msg`.
func (o *Option[T]) Expect(msg string) *T {
	if o.value == nil {
		panic(msg)
	}
	return o.value
}

// Unwrap returns the contained [`Some`] value, consuming the `self` value.
// Panics if the self value equals [`None`].
func (o *Option[T]) Unwrap(msg string) *T {
	if o.value == nil {
		panic("called `Option::unwrap()` on a `None` value")
	}
	return o.value
}

// Unwrap returns the contained [`Some`] value, consuming the `self` value.
// Panics if the self value equals [`None`].
func (o *Option[T]) UnwrapOr(v *T) *T {
	if o.value == nil {
		return v
	}
	return o.value
}

// UnwrapOrElse returns the contained [`Some`] value or computes it from a closure.
func (o *Option[T]) UnwrapOrElse(f func() *T) *T {
	if o.value == nil {
		return f()
	}
	return o.value
}

// UnwrapOrElse returns the contained [`Some`] value or a default.
func (o *Option[T]) UnwrapOrDefault() *T {
	if o.value == nil {
		return nil
	}
	return o.value
}

// Inspect calls the provided closure with a reference to the contained value (if [`Some`]).
func (o *Option[T]) Inspect(f func(*T)) *Option[T] {
	if o.value != nil {
		f(o.value)
	}
	return o
}

// Or returns the option if it contains a value, otherwise returns `optb`.
func (o *Option[T]) Or(optb *Option[T]) *Option[T] {
	if o.value != nil {
		return o
	}
	return optb
}

// OrElse returns the option if it contains a value, otherwise calls `f` and returns the result.
func (o *Option[T]) OrElse(f func() *Option[T]) *Option[T] {
	if o.value != nil {
		return o
	}
	return f()
}

// XOr returns [`Some`] if exactly one of `self`, `optb` is [`Some`], otherwise returns [`None`].
func (o *Option[T]) XOr(optb *Option[T]) *Option[T] {
	if o.value != nil {
		return o
	}
	if optb.value != nil {
		return optb
	}
	return nil
}

// Take takes the value out of the option, leaving a [`None`] in its place.
func (o *Option[T]) Take() *T {
	v := o.value
	o.value = nil
	return v
}

// TakeIf takes the value out of the option, but only if the predicate evaluates to `true` to the value.
func (o *Option[T]) TakeIf(f func(*T) bool) *T {
	if f(o.value) {
		v := o.value
		o.value = nil
		return v
	}
	return nil
}

// Replace replaces the actual value in the option by the value given in parameter,
// returning the old value if present,
// leaving a [`Some`] in its place without deinitializing either one.
func (o *Option[T]) Replace(v *T) *T {
	old := o.value
	o.value = v
	return old
}

// OkOr transforms the `Option[T]` into a `Result[T]`, mapping [`Some(v)`] to [`Ok(v)`] and [`None`] to [`Err(err)`].
// func (o *Option[T]) OkOr(err error) *result.Result[T] {
// 	if o.value == nil {
// 		return result.Err[T](err)
// 	}
// 	return result.Ok[T](o.value)
// }

// OkOrElse transforms the `Option[T]` into a `Result[T]`, mapping [`Some(v)`] to [`Ok(v)`] and [`None`] to [`Err(err())`].
// func (o *Option[T]) OkOrElse(errFn func() error) *result.Result[T] {
// 	if o.value == nil {
// 		return result.Err[T](errFn())
// 	}
// 	return result.Ok[T](o.value)
// }
