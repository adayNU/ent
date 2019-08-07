package gen

// Op is a predicate for the where clause.
type Op int

// List of all builtin predicates.
const (
	EQ        Op = iota // =
	NEQ                 // <>
	GT                  // >
	GTE                 // >=
	LT                  // <
	LTE                 // <=
	IsNull              // IS NULL / has
	NotNull             // IS NOT NULL / hasNot
	In                  // within
	NotIn               // without
	Contains            // containing
	HasPrefix           // startingWith
	HasSuffix           // endingWith
)

// Name returns the string representation of an operator.
func (o Op) Name() string {
	if int(o) < len(opText) {
		return opText[o]
	}
	return "Unknown"
}

// Gremlin returns the gremlin code representation of an operator.
func (o Op) Gremlin() string {
	if code := gremlinCode[o]; code != "" {
		return code
	}
	return o.Name()
}

// Variadic reports if the predicate is a variadic function.
func (o Op) Variadic() bool {
	return o == In || o == NotIn
}

// Niladic reports if the predicate is a niladic predicate.
func (o Op) Niladic() bool {
	return o == IsNull || o == NotNull
}

var (
	// operations text.
	opText = [...]string{
		EQ:        "EQ",
		NEQ:       "NEQ",
		GT:        "GT",
		GTE:       "GTE",
		LT:        "LT",
		LTE:       "LTE",
		IsNull:    "IsNull",
		NotNull:   "NotNull",
		Contains:  "Contains",
		HasPrefix: "HasPrefix",
		HasSuffix: "HasSuffix",
		In:        "In",
		NotIn:     "NotIn",
	}
	// operations code in gremlin.
	gremlinCode = [...]string{
		IsNull:    "HasNot",
		NotNull:   "Has",
		In:        "Within",
		NotIn:     "Without",
		Contains:  "Containing",
		HasPrefix: "StartingWith",
		HasSuffix: "EndingWith",
	}
	// operations per type.
	boolOps     = []Op{EQ, NEQ}
	numericOps  = append(boolOps[:], GT, GTE, LT, LTE, In, NotIn)
	stringOps   = append(numericOps[:], Contains, HasPrefix, HasSuffix)
	nillableOps = []Op{IsNull, NotNull}
)