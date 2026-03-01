package base

//

type (
	Any = interface{}
	Map = map[string]Any

	Res interface {
		OK() bool
		Fail() bool
		Code() int
		Status() string
		Args() []Any
		With(...Any) Res
		Error() string
	}

	Vars map[string]Var
	Var  struct {
		nil       bool
		Type      string
		Required  bool
		Nullable  bool
		Name      string
		Text      string
		Default   Any
		Unique    bool
		Check     string
		Collation string
		Comment   string
		Setting   Map
		Options   Map
		Children  Vars
		Encode    string
		Decode    string
		Empty     Res
		Error     Res
		Valid     func(Any, Var) bool
		Value     func(Any, Var) Any
	}
)

const (
	ASC  = 1
	DESC = -1
)

var (
	Nil = Var{nil: true}
)

func (vvv *Var) Nil() bool {
	if vvv == nil {
		return true
	}
	return vvv.nil
}
