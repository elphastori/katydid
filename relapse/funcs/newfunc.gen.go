// Code generated by funcs-gen. DO NOT EDIT.
package funcs

//NewDoubleFunc dynamically creates and asserts the returning function is of type Double.
//This function is used by the compose library to compile functions together.
func NewDoubleFunc(uniq string, values ...interface{}) (Double, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Double), nil
}

//NewIntFunc dynamically creates and asserts the returning function is of type Int.
//This function is used by the compose library to compile functions together.
func NewIntFunc(uniq string, values ...interface{}) (Int, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Int), nil
}

//NewUintFunc dynamically creates and asserts the returning function is of type Uint.
//This function is used by the compose library to compile functions together.
func NewUintFunc(uniq string, values ...interface{}) (Uint, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Uint), nil
}

//NewBoolFunc dynamically creates and asserts the returning function is of type Bool.
//This function is used by the compose library to compile functions together.
func NewBoolFunc(uniq string, values ...interface{}) (Bool, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Bool), nil
}

//NewStringFunc dynamically creates and asserts the returning function is of type String.
//This function is used by the compose library to compile functions together.
func NewStringFunc(uniq string, values ...interface{}) (String, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(String), nil
}

//NewBytesFunc dynamically creates and asserts the returning function is of type Bytes.
//This function is used by the compose library to compile functions together.
func NewBytesFunc(uniq string, values ...interface{}) (Bytes, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Bytes), nil
}

//NewDoublesFunc dynamically creates and asserts the returning function is of type Doubles.
//This function is used by the compose library to compile functions together.
func NewDoublesFunc(uniq string, values ...interface{}) (Doubles, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Doubles), nil
}

//NewIntsFunc dynamically creates and asserts the returning function is of type Ints.
//This function is used by the compose library to compile functions together.
func NewIntsFunc(uniq string, values ...interface{}) (Ints, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Ints), nil
}

//NewUintsFunc dynamically creates and asserts the returning function is of type Uints.
//This function is used by the compose library to compile functions together.
func NewUintsFunc(uniq string, values ...interface{}) (Uints, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Uints), nil
}

//NewBoolsFunc dynamically creates and asserts the returning function is of type Bools.
//This function is used by the compose library to compile functions together.
func NewBoolsFunc(uniq string, values ...interface{}) (Bools, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Bools), nil
}

//NewStringsFunc dynamically creates and asserts the returning function is of type Strings.
//This function is used by the compose library to compile functions together.
func NewStringsFunc(uniq string, values ...interface{}) (Strings, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(Strings), nil
}

//NewListOfBytesFunc dynamically creates and asserts the returning function is of type ListOfBytes.
//This function is used by the compose library to compile functions together.
func NewListOfBytesFunc(uniq string, values ...interface{}) (ListOfBytes, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.(ListOfBytes), nil
}
