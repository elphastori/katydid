// Code generated by funcs-gen. DO NOT EDIT.
package funcs

type inSetInt struct {
	Elem Int
	List ConstInts
	set  map[int64]struct{}
}

func (this *inSetInt) Init() error {
	l, err := this.List.Eval()
	if err != nil {
		return err
	}
	this.set = make(map[int64]struct{})
	for i := range l {
		this.set[l[i]] = struct{}{}
	}
	return nil
}

func (this *inSetInt) Eval() (bool, error) {
	v, err := this.Elem.Eval()
	if err != nil {
		return false, err
	}
	_, ok := this.set[v]
	return ok, nil
}

func init() {
	Register("contains", new(inSetInt))
}

//ContainsInt returns a function that checks when the element if contained in the list.
func ContainsInt(element Int, list ConstInts) Bool {
	return &inSetInt{element, list, nil}
}

type inSetUint struct {
	Elem Uint
	List ConstUints
	set  map[uint64]struct{}
}

func (this *inSetUint) Init() error {
	l, err := this.List.Eval()
	if err != nil {
		return err
	}
	this.set = make(map[uint64]struct{})
	for i := range l {
		this.set[l[i]] = struct{}{}
	}
	return nil
}

func (this *inSetUint) Eval() (bool, error) {
	v, err := this.Elem.Eval()
	if err != nil {
		return false, err
	}
	_, ok := this.set[v]
	return ok, nil
}

func init() {
	Register("contains", new(inSetUint))
}

//ContainsUint returns a function that checks when the element if contained in the list.
func ContainsUint(element Uint, list ConstUints) Bool {
	return &inSetUint{element, list, nil}
}

type inSetString struct {
	Elem String
	List ConstStrings
	set  map[string]struct{}
}

func (this *inSetString) Init() error {
	l, err := this.List.Eval()
	if err != nil {
		return err
	}
	this.set = make(map[string]struct{})
	for i := range l {
		this.set[l[i]] = struct{}{}
	}
	return nil
}

func (this *inSetString) Eval() (bool, error) {
	v, err := this.Elem.Eval()
	if err != nil {
		return false, err
	}
	_, ok := this.set[v]
	return ok, nil
}

func init() {
	Register("contains", new(inSetString))
}

//ContainsString returns a function that checks when the element if contained in the list.
func ContainsString(element String, list ConstStrings) Bool {
	return &inSetString{element, list, nil}
}
