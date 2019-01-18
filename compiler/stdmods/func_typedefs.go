package stdmods

import (
	"github.com/d5/tengo/objects"
)

// FuncAR transform a function of 'func()' signature
// into a user function object.
func FuncAR(fn func()) *objects.UserFunction {
	return &objects.UserFunction{
		Value: func(args ...objects.Object) (ret objects.Object, err error) {
			if len(args) != 0 {
				return nil, objects.ErrWrongNumArguments
			}

			fn()

			return objects.UndefinedValue, nil
		},
	}
}

// FuncARF transform a function of 'func() float64' signature
// into a user function object.
func FuncARF(fn func() float64) *objects.UserFunction {
	return &objects.UserFunction{
		Value: func(args ...objects.Object) (ret objects.Object, err error) {
			if len(args) != 0 {
				return nil, objects.ErrWrongNumArguments
			}

			return &objects.Float{Value: fn()}, nil
		},
	}
}

// FuncARSs transform a function of 'func() []string' signature
// into a user function object.
func FuncARSs(fn func() []string) *objects.UserFunction {
	return &objects.UserFunction{
		Value: func(args ...objects.Object) (ret objects.Object, err error) {
			if len(args) != 0 {
				return nil, objects.ErrWrongNumArguments
			}

			arr := &objects.Array{}
			for _, osArg := range fn() {
				arr.Value = append(arr.Value, &objects.String{Value: osArg})
			}

			return arr, nil
		},
	}
}

// FuncAFRF transform a function of 'func(float64) float64' signature
// into a user function object.
func FuncAFRF(fn func(float64) float64) *objects.UserFunction {
	return &objects.UserFunction{
		Value: func(args ...objects.Object) (ret objects.Object, err error) {
			if len(args) != 1 {
				return nil, objects.ErrWrongNumArguments
			}

			f1, ok := objects.ToFloat64(args[0])
			if !ok {
				return nil, objects.ErrInvalidTypeConversion
			}

			return &objects.Float{Value: fn(f1)}, nil
		},
	}
}

// FuncAIRF transform a function of 'func(int) float64' signature
// into a user function object.
func FuncAIRF(fn func(int) float64) *objects.UserFunction {
	return &objects.UserFunction{
		Value: func(args ...objects.Object) (ret objects.Object, err error) {
			if len(args) != 1 {
				return nil, objects.ErrWrongNumArguments
			}

			i1, ok := objects.ToInt(args[0])
			if !ok {
				return nil, objects.ErrInvalidTypeConversion
			}

			return &objects.Float{Value: fn(i1)}, nil
		},
	}
}

// FuncAFRI transform a function of 'func(float64) int' signature
// into a user function object.
func FuncAFRI(fn func(float64) int) *objects.UserFunction {
	return &objects.UserFunction{
		Value: func(args ...objects.Object) (ret objects.Object, err error) {
			if len(args) != 1 {
				return nil, objects.ErrWrongNumArguments
			}

			f1, ok := objects.ToFloat64(args[0])
			if !ok {
				return nil, objects.ErrInvalidTypeConversion
			}

			return &objects.Int{Value: int64(fn(f1))}, nil
		},
	}
}

// FuncAFFRF transform a function of 'func(float64, float64) float64' signature
// into a user function object.
func FuncAFFRF(fn func(float64, float64) float64) *objects.UserFunction {
	return &objects.UserFunction{
		Value: func(args ...objects.Object) (ret objects.Object, err error) {
			if len(args) != 2 {
				return nil, objects.ErrWrongNumArguments
			}

			f1, ok := objects.ToFloat64(args[0])
			if !ok {
				return nil, objects.ErrInvalidTypeConversion
			}

			f2, ok := objects.ToFloat64(args[1])
			if !ok {
				return nil, objects.ErrInvalidTypeConversion
			}

			return &objects.Float{Value: fn(f1, f2)}, nil
		},
	}
}

// FuncAIFRF transform a function of 'func(int, float64) float64' signature
// into a user function object.
func FuncAIFRF(fn func(int, float64) float64) *objects.UserFunction {
	return &objects.UserFunction{
		Value: func(args ...objects.Object) (ret objects.Object, err error) {
			if len(args) != 2 {
				return nil, objects.ErrWrongNumArguments
			}

			i1, ok := objects.ToInt(args[0])
			if !ok {
				return nil, objects.ErrInvalidTypeConversion
			}

			f2, ok := objects.ToFloat64(args[1])
			if !ok {
				return nil, objects.ErrInvalidTypeConversion
			}

			return &objects.Float{Value: fn(i1, f2)}, nil
		},
	}
}

// FuncAFIRF transform a function of 'func(float64, int) float64' signature
// into a user function object.
func FuncAFIRF(fn func(float64, int) float64) *objects.UserFunction {
	return &objects.UserFunction{
		Value: func(args ...objects.Object) (ret objects.Object, err error) {
			if len(args) != 2 {
				return nil, objects.ErrWrongNumArguments
			}

			f1, ok := objects.ToFloat64(args[0])
			if !ok {
				return nil, objects.ErrInvalidTypeConversion
			}

			i2, ok := objects.ToInt(args[1])
			if !ok {
				return nil, objects.ErrInvalidTypeConversion
			}

			return &objects.Float{Value: fn(f1, i2)}, nil
		},
	}
}

// FuncAFIRB transform a function of 'func(float64, int) bool' signature
// into a user function object.
func FuncAFIRB(fn func(float64, int) bool) *objects.UserFunction {
	return &objects.UserFunction{
		Value: func(args ...objects.Object) (ret objects.Object, err error) {
			if len(args) != 2 {
				return nil, objects.ErrWrongNumArguments
			}

			f1, ok := objects.ToFloat64(args[0])
			if !ok {
				return nil, objects.ErrInvalidTypeConversion
			}

			i2, ok := objects.ToInt(args[1])
			if !ok {
				return nil, objects.ErrInvalidTypeConversion
			}

			return &objects.Bool{Value: fn(f1, i2)}, nil
		},
	}
}

// FuncAFRB transform a function of 'func(float64) bool' signature
// into a user function object.
func FuncAFRB(fn func(float64) bool) *objects.UserFunction {
	return &objects.UserFunction{
		Value: func(args ...objects.Object) (ret objects.Object, err error) {
			if len(args) != 1 {
				return nil, objects.ErrWrongNumArguments
			}

			f1, ok := objects.ToFloat64(args[0])
			if !ok {
				return nil, objects.ErrInvalidTypeConversion
			}

			return &objects.Bool{Value: fn(f1)}, nil
		},
	}
}

// FuncASRE transform a function of 'func(string) error' signature into a user function object.
// User function will return 'true' if underlying native function returns nil.
func FuncASRE(fn func(string) error) *objects.UserFunction {
	return &objects.UserFunction{
		Value: func(args ...objects.Object) (objects.Object, error) {
			if len(args) != 1 {
				return nil, objects.ErrWrongNumArguments
			}

			s1, ok := objects.ToString(args[0])
			if !ok {
				return nil, objects.ErrInvalidTypeConversion
			}

			return wrapError(fn(s1)), nil
		},
	}
}
