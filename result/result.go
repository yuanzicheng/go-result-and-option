package result

type Result[T any] struct {
	value *T
	err   error
}

func New[T any](v *T, e error) *Result[T] {
	if e == nil {
		return &Result[T]{value: v}
	}
	return &Result[T]{err: e}
}

func Ok[T any](v *T) *Result[T] {
	return &Result[T]{value: v}
}

func Err[T any](err error) *Result[T] {
	return &Result[T]{err: err}
}

// And returns `out` if the result is [`Ok`], otherwise returns the [`Err`] value of `in`.
func And[T any, U any](in *Result[T], out *Result[U]) *Result[U] {
	if in.IsErr() {
		return Err[U](in.err)
	}
	return out
}

// AndThen calls `op` if the result is [`Ok`], otherwise returns the [`Err`] value of `in`.
func AndThen[T any, U any](in *Result[T], op func(*T) *Result[U]) *Result[U] {
	if in.IsErr() {
		return Err[U](in.err)
	}
	return op(in.value)
}

// Maps a `Result[T]` to `Result[U]` by applying a function to a
// contained [`Ok`] value, leaving an [`Err`] value untouched.
//
// This function can be used to compose the results of two functions.
func Map[T any, U any](r Result[T], f func(*T) *U) *Result[U] {
	if r.IsErr() {
		return Err[U](r.err)
	}
	return Ok(f(r.value))
}

// MapOr returns the provided fallback (if [`Err`]), or
// applies a function to the contained value (if [`Ok`]).
func MapOr[T any, U any](r *Result[T], fallback *U, f func(*T) *U) *U {
	if r.IsErr() {
		return fallback
	}
	return f(r.value)
}

// MapOrElse maps a `Result[T]` to `U` by applying fallback function `fallbackFn` to
// a contained [`Err`] value, or function `f` to a contained [`Ok`] value.
//
// This function can be used to unpack a successful result while handling an error.
func MapOrElse[T any, U any](r *Result[T], fallbackFn func(error) *U, f func(*T) *U) *U {
	if r.IsErr() {
		return fallbackFn(r.err)
	}
	return f(r.value)
}

// MapErr maps the error of a `Result[T]` to another error by applying a function to
// a contained [`Err`] value, leaving an [`Ok`] value untouched.
//
// This function can be used to pass through a successful result while handling an error.
func MapErr[T any](r *Result[T], f func(error) error) *Result[T] {
	if r.IsOk() {
		return r
	}
	return Err[T](f(r.err))
}

// IsOk returns `true` if the result is [`Ok`].
func (r Result[T]) IsOk() bool {
	return r.err == nil
}

// IsOkAnd returns `true` if the result is [`Ok`] and the value inside of it matches a predicate.
func (r *Result[T]) IsOkAnd(f func(*T) bool) bool {
	if r.IsErr() {
		return false
	}
	return f(r.value)
}

// IsOkAndNotNil returns `true` if the result is [`Ok`] and the value inside of it is not nil.
func (r *Result[T]) IsOkAndNotNil() bool {
	return r.IsOkAnd(func(t *T) bool {
		return t != nil
	})
}

// IsErr returns `true` if the result is [`Err`].
func (r *Result[T]) IsErr() bool {
	return r.err != nil
}

// IsErrAnd returns `true` if the result is [`Err`] and the value inside of it matches a predicate.
func (r *Result[T]) IsErrAnd(f func(error) bool) bool {
	if r.IsErr() {
		return true
	}
	return f(r.err)
}

// Inspect calls the provided closure with a reference to the contained value (if [`Ok`]).
func (r *Result[T]) Inspect(f func(*T)) *Result[T] {
	if r.IsOk() {
		f(r.value)
	}
	return r
}

// InspectErr calls the provided closure with a reference to the contained error (if [`Err`]).
func (r *Result[T]) InspectErr(f func(error)) *Result[T] {
	if r.IsErr() {
		f(r.err)
	}
	return r
}

// Expect returns the contained [`Ok`] value, consuming the `self` value.
// Panics if the value is an [`Err`], with a panic message including the passed message.
func (r *Result[T]) Expect(msg string) *T {
	if r.IsErr() {
		panic(msg)
	}
	return r.value
}

// ExpectErr returns the contained [`Err`] value, consuming the `self` value.
// Panics if the value is an [`Ok`], with a panic message including the passed message.
func (r *Result[T]) ExpectErr(msg string) error {
	if r.IsOk() {
		panic(msg)
	}
	return r.err
}

// Unwrap extracts the value from the Result. Panics if the Result is Error.
func (r *Result[T]) Unwrap() *T {
	if r.IsErr() {
		panic("called `Result::unwrap()` on an `Err` value")
	}

	return r.value
}

// UnwrapError extracts the error from the Result. Panics if the Result is Ok.
func (r *Result[T]) UnwrapError() error {
	if r.IsOk() {
		panic("called `Result::unwrap_err()` on an `Ok` value")
	}

	return r.err
}

// UnwrapOr extracts the value from the Result. Returns the provided value if the Result is Error.
func (r *Result[T]) UnwrapOr(v *T) *T {
	if r.IsErr() {
		return v
	}

	return r.value
}

// UnwrapOrElse returns the contained [`Ok`] value or computes it from a closure.
func (r *Result[T]) UnwrapOrElse(f func() *T) *T {
	if r.IsErr() {
		return f()
	}
	return r.value
}

// UnwrapOrDefault returns the contained [`Ok`] value or a default (nil).
func (r *Result[T]) UnwrapOrDefault() *T {
	if r.IsErr() {
		return nil
	}
	return r.value
}

// Or returns `res` if the result is [`Err`], otherwise returns the [`Ok`] value of `self`.
func (r *Result[T]) Or(res *Result[T]) *Result[T] {
	if r.IsOk() {
		return r
	}
	return res
}

// OrElse calls `op` if the result is [`Err`], otherwise returns the [`Ok`] value of `self`.
func (r *Result[T]) OrElse(op func(error) *Result[T]) *Result[T] {
	if r.IsOk() {
		return r
	}
	return op(r.err)
}
