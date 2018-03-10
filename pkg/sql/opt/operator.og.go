// Code generated by optgen; DO NOT EDIT.

package opt

const (
	UnknownOp Operator = iota

	// ------------------------------------------------------------
	// Scalar Operators
	// ------------------------------------------------------------

	SubqueryOp

	// VariableOp is the typed scalar value of a column in the query. The private
	// field is a Metadata.ColumnIndex that references the column by index.
	VariableOp

	// ConstOp is a typed scalar constant value. The private field is a tree.Datum
	// value having any datum type that's legal in the expression's context.
	ConstOp

	// NullOp is the constant SQL null value that has "unknown value" semantics. If
	// the Typ field is not types.Unknown, then the value is known to be in the
	// domain of that type. This is important for preserving correct types in
	// replacement patterns. For example:
	//   (Plus (Function ...) (Const 1))
	//
	// If the function in that expression has a static type of Int, but then it gets
	// constant folded to (Null), then its type must remain as Int. Any other type
	// violates logical equivalence of the expression, breaking type inference and
	// possibly changing the results of execution. The solution is to tag the null
	// with the correct type:
	//   (Plus (Null (Int)) (Const 1))
	//
	// Null is its own operator rather than a Const datum in order to make matching
	// and replacement easier and more efficient, as patterns can contain (Null)
	// expressions.
	NullOp

	// TrueOp is the boolean true value that is equivalent to the tree.DBoolTrue datum
	// value. It is a separate operator to make matching and replacement simpler and
	// more efficient, as patterns can contain (True) expressions.
	TrueOp

	// FalseOp is the boolean false value that is equivalent to the tree.DBoolFalse
	// datum value. It is a separate operator to make matching and replacement
	// simpler and more efficient, as patterns can contain (False) expressions.
	FalseOp

	PlaceholderOp

	TupleOp

	// ProjectionsOp is a set of typed scalar expressions that will become output
	// columns for a containing Project operator. The private Cols field contains
	// the list of column indexes returned by the expression, as a *opt.ColList. It
	// is not legal for Cols to be empty.
	ProjectionsOp

	// AggregationsOp is a set of aggregate expressions that will become output
	// columns for a containing GroupBy operator. The private Cols field contains
	// the list of column indexes returned by the expression, as a *ColList. It
	// is legal for Cols to be empty.
	AggregationsOp

	ExistsOp

	// AndOp is the boolean conjunction operator that evalutes to true if all of its
	// conditions evaluate to true. If the conditions list is empty, it evalutes to
	// true.
	AndOp

	// OrOp is the boolean disjunction operator that evalutes to true if any of its
	// conditions evaluate to true. If the conditions list is empty, it evaluates to
	// false.
	OrOp

	// NotOp is the boolean negation operator that evaluates to true if its input
	// evalutes to false.
	NotOp

	EqOp

	LtOp

	GtOp

	LeOp

	GeOp

	NeOp

	InOp

	NotInOp

	LikeOp

	NotLikeOp

	ILikeOp

	NotILikeOp

	SimilarToOp

	NotSimilarToOp

	RegMatchOp

	NotRegMatchOp

	RegIMatchOp

	NotRegIMatchOp

	IsOp

	IsNotOp

	ContainsOp

	BitandOp

	BitorOp

	BitxorOp

	PlusOp

	MinusOp

	MultOp

	DivOp

	FloorDivOp

	ModOp

	PowOp

	ConcatOp

	LShiftOp

	RShiftOp

	FetchValOp

	FetchTextOp

	FetchValPathOp

	FetchTextPathOp

	UnaryMinusOp

	UnaryComplementOp

	CastOp

	// CaseOp is a CASE statement of the form:
	//   CASE [ <Input> ]
	//       WHEN <condval1> THEN <expr1>
	//     [ WHEN <condval2> THEN <expr2> ] ...
	//     [ ELSE <expr> ]
	//   END
	//
	// The Case operator evaluates <Input> (if not provided, Input is set to True),
	// then picks the WHEN branch where <condval> is equal to
	// <cond>, then evaluates and returns the corresponding THEN expression. If no
	// WHEN branch matches, the ELSE expression is evaluated and returned, if any.
	// Otherwise, NULL is returned.
	//
	// Note that the Whens list inside Case is used to represent all the WHEN
	// branches as well as the ELSE statement if it exists. It is of the form:
	// [(When <condval1> <expr1>),(When <condval2> <expr2>),...,<expr>]
	CaseOp

	// WhenOp represents a single WHEN ... THEN ... condition inside a CASE statement.
	// It is the type of each list item in Whens (except for the last item which is
	// a raw expression for the ELSE statement).
	WhenOp

	// FunctionOp invokes a builtin SQL function like CONCAT or NOW, passing the given
	// arguments. The private field is an opt.FuncDef struct that provides the name
	// of the function as well as a pointer to the builtin overload definition.
	FunctionOp

	CoalesceOp

	// UnsupportedExprOp is used for interfacing with the old planner code. It can
	// encapsulate a TypedExpr that is otherwise not supported by the optimizer.
	UnsupportedExprOp

	// ------------------------------------------------------------
	// Relational Operators
	// ------------------------------------------------------------

	// ScanOp returns a result set containing every row in the specified table. Rows
	// and columns are not expected to have any particular ordering. The private
	// Table field is a Metadata.TableIndex that references an opt.Table
	// definition in the query's metadata.
	ScanOp

	// ValuesOp returns a manufactured result set containing a constant number of rows.
	// specified by the Rows list field. Each row must contain the same set of
	// columns in the same order.
	//
	// The Rows field contains a list of Tuples, one for each row. Each tuple has
	// the same length (same with that of Cols).
	//
	// The Cols field contains the set of column indices returned by each row
	// as a *ColList. It is legal for Cols to be empty.
	ValuesOp

	// SelectOp filters rows from its input result set, based on the boolean filter
	// predicate expression. Rows which do not match the filter are discarded.
	SelectOp

	// ProjectOp modifies the set of columns returned by the input result set. Columns
	// can be removed, reordered, or renamed. In addition, new columns can be
	// synthesized. Projections is a scalar Projections list operator that contains
	// the list of expressions that describe the output columns. The Cols field of
	// the Projections operator provides the indexes of each of the output columns.
	ProjectOp

	// InnerJoinOp creates a result set that combines columns from its left and right
	// inputs, based upon its "on" join predicate. Rows which do not match the
	// predicate are filtered. While expressions in the predicate can refer to
	// columns projected by either the left or right inputs, the inputs are not
	// allowed to refer to the other's projected columns.
	InnerJoinOp

	LeftJoinOp

	RightJoinOp

	FullJoinOp

	SemiJoinOp

	AntiJoinOp

	// InnerJoinApplyOp has the same join semantics as InnerJoin. However, unlike
	// InnerJoin, it allows the right input to refer to columns projected by the
	// left input.
	InnerJoinApplyOp

	LeftJoinApplyOp

	RightJoinApplyOp

	FullJoinApplyOp

	SemiJoinApplyOp

	AntiJoinApplyOp

	// GroupByOp is an operator that is used for performing aggregations (for queries
	// with aggregate functions, HAVING clauses and/or group by expressions). It
	// groups results that are equal on the grouping columns and computes
	// aggregations as described by Aggregations (which is always an Aggregations
	// operator). The arguments of the aggregations are columns from the input.
	GroupByOp

	// UnionOp is an operator used to combine the Left and Right input relations into
	// a single set containing rows from both inputs. Duplicate rows are discarded.
	// The private field, ColMap, matches columns from the Left and Right inputs
	// of the Union with the output columns. See the comment above opt.SetOpColMap
	// for more details.
	UnionOp

	// IntersectOp is an operator used to perform an intersection between the Left
	// and Right input relations. The result consists only of rows in the Left
	// relation that are also present in the Right relation. Duplicate rows are
	// discarded.
	// The private field, ColMap, matches columns from the Left and Right inputs
	// of the Intersect with the output columns. See the comment above
	// opt.SetOpColMap for more details.
	IntersectOp

	// ExceptOp is an operator used to perform a set difference between the Left and
	// Right input relations. The result consists only of rows in the Left relation
	// that are not present in the Right relation. Duplicate rows are discarded.
	// The private field, ColMap, matches columns from the Left and Right inputs
	// of the Except with the output columns. See the comment above opt.SetOpColMap
	// for more details.
	ExceptOp

	// UnionAllOp is an operator used to combine the Left and Right input relations
	// into a single set containing rows from both inputs. Duplicate rows are
	// not discarded. For example:
	//   SELECT x FROM xx UNION ALL SELECT y FROM yy
	//     x       y         out
	//   -----   -----      -----
	//     1       1          1
	//     1       2    ->    1
	//     2       3          1
	//                        2
	//                        2
	//                        3
	//
	// The private field, ColMap, matches columns from the Left and Right inputs
	// of the UnionAll with the output columns. See the comment above
	// opt.SetOpColMap for more details.
	UnionAllOp

	// IntersectAllOp is an operator used to perform an intersection between the Left
	// and Right input relations. The result consists only of rows in the Left
	// relation that have a corresponding row in the Right relation. Duplicate rows
	// are not discarded. This effectively creates a one-to-one mapping between the
	// Left and Right rows. For example:
	//   SELECT x FROM xx INTERSECT ALL SELECT y FROM yy
	//     x       y         out
	//   -----   -----      -----
	//     1       1          1
	//     1       1    ->    1
	//     1       2          2
	//     2       2          2
	//     2       3
	//     4
	//
	// The private field, ColMap, matches columns from the Left and Right inputs
	// of the IntersectAll with the output columns. See the comment above
	// opt.SetOpColMap for more details.
	IntersectAllOp

	// ExceptAllOp is an operator used to perform a set difference between the Left
	// and Right input relations. The result consists only of rows in the Left
	// relation that do not have a corresponding row in the Right relation.
	// Duplicate rows are not discarded. This effectively creates a one-to-one
	// mapping between the Left and Right rows. For example:
	//   SELECT x FROM xx EXCEPT ALL SELECT y FROM yy
	//     x       y         out
	//   -----   -----      -----
	//     1       1    ->    1
	//     1       1          4
	//     1       2
	//     2       2
	//     2       3
	//     4
	//
	// The private field, ColMap, matches columns from the Left and Right inputs
	// of the ExceptAll with the output columns. See the comment above
	// opt.SetOpColMap for more details.
	ExceptAllOp

	// ------------------------------------------------------------
	// Enforcer Operators
	// ------------------------------------------------------------

	// SortOp enforces the ordering of rows returned by its input expression. Rows can
	// be sorted by one or more of the input columns, each of which can be sorted in
	// either ascending or descending order. See the Ordering field in the
	// PhysicalProps struct.
	// TODO(andyk): Add the Ordering field.
	SortOp

	// NumOperators tracks the total count of operators.
	NumOperators
)

const opNames = "unknownsubqueryvariableconstnulltruefalseplaceholdertupleprojectionsaggregationsexistsandornoteqltgtlegeneinnot-inlikenot-likei-likenot-i-likesimilar-tonot-similar-toreg-matchnot-reg-matchreg-i-matchnot-reg-i-matchisis-notcontainsbitandbitorbitxorplusminusmultdivfloor-divmodpowconcatl-shiftr-shiftfetch-valfetch-textfetch-val-pathfetch-text-pathunary-minusunary-complementcastcasewhenfunctioncoalesceunsupported-exprscanvaluesselectprojectinner-joinleft-joinright-joinfull-joinsemi-joinanti-joininner-join-applyleft-join-applyright-join-applyfull-join-applysemi-join-applyanti-join-applygroup-byunionintersectexceptunion-allintersect-allexcept-allsort"

var opIndexes = [...]uint32{0, 7, 15, 23, 28, 32, 36, 41, 52, 57, 68, 80, 86, 89, 91, 94, 96, 98, 100, 102, 104, 106, 108, 114, 118, 126, 132, 142, 152, 166, 175, 188, 199, 214, 216, 222, 230, 236, 241, 247, 251, 256, 260, 263, 272, 275, 278, 284, 291, 298, 307, 317, 331, 346, 357, 373, 377, 381, 385, 393, 401, 417, 421, 427, 433, 440, 450, 459, 469, 478, 487, 496, 512, 527, 543, 558, 573, 588, 596, 601, 610, 616, 625, 638, 648, 652}

var ScalarOperators = [...]Operator{
	SubqueryOp,
	VariableOp,
	ConstOp,
	NullOp,
	TrueOp,
	FalseOp,
	PlaceholderOp,
	TupleOp,
	ProjectionsOp,
	AggregationsOp,
	ExistsOp,
	AndOp,
	OrOp,
	NotOp,
	EqOp,
	LtOp,
	GtOp,
	LeOp,
	GeOp,
	NeOp,
	InOp,
	NotInOp,
	LikeOp,
	NotLikeOp,
	ILikeOp,
	NotILikeOp,
	SimilarToOp,
	NotSimilarToOp,
	RegMatchOp,
	NotRegMatchOp,
	RegIMatchOp,
	NotRegIMatchOp,
	IsOp,
	IsNotOp,
	ContainsOp,
	BitandOp,
	BitorOp,
	BitxorOp,
	PlusOp,
	MinusOp,
	MultOp,
	DivOp,
	FloorDivOp,
	ModOp,
	PowOp,
	ConcatOp,
	LShiftOp,
	RShiftOp,
	FetchValOp,
	FetchTextOp,
	FetchValPathOp,
	FetchTextPathOp,
	UnaryMinusOp,
	UnaryComplementOp,
	CastOp,
	CaseOp,
	WhenOp,
	FunctionOp,
	CoalesceOp,
	UnsupportedExprOp,
}

var ConstValueOperators = [...]Operator{
	ConstOp,
	NullOp,
	TrueOp,
	FalseOp,
}

var BooleanOperators = [...]Operator{
	TrueOp,
	FalseOp,
	AndOp,
	OrOp,
	NotOp,
}

var ComparisonOperators = [...]Operator{
	EqOp,
	LtOp,
	GtOp,
	LeOp,
	GeOp,
	NeOp,
	InOp,
	NotInOp,
	LikeOp,
	NotLikeOp,
	ILikeOp,
	NotILikeOp,
	SimilarToOp,
	NotSimilarToOp,
	RegMatchOp,
	NotRegMatchOp,
	RegIMatchOp,
	NotRegIMatchOp,
	IsOp,
	IsNotOp,
	ContainsOp,
}

var BinaryOperators = [...]Operator{
	BitandOp,
	BitorOp,
	BitxorOp,
	PlusOp,
	MinusOp,
	MultOp,
	DivOp,
	FloorDivOp,
	ModOp,
	PowOp,
	ConcatOp,
	LShiftOp,
	RShiftOp,
	FetchValOp,
	FetchTextOp,
	FetchValPathOp,
	FetchTextPathOp,
}

var UnaryOperators = [...]Operator{
	UnaryMinusOp,
	UnaryComplementOp,
}

var RelationalOperators = [...]Operator{
	ScanOp,
	ValuesOp,
	SelectOp,
	ProjectOp,
	InnerJoinOp,
	LeftJoinOp,
	RightJoinOp,
	FullJoinOp,
	SemiJoinOp,
	AntiJoinOp,
	InnerJoinApplyOp,
	LeftJoinApplyOp,
	RightJoinApplyOp,
	FullJoinApplyOp,
	SemiJoinApplyOp,
	AntiJoinApplyOp,
	GroupByOp,
	UnionOp,
	IntersectOp,
	ExceptOp,
	UnionAllOp,
	IntersectAllOp,
	ExceptAllOp,
}

var JoinOperators = [...]Operator{
	InnerJoinOp,
	LeftJoinOp,
	RightJoinOp,
	FullJoinOp,
	SemiJoinOp,
	AntiJoinOp,
	InnerJoinApplyOp,
	LeftJoinApplyOp,
	RightJoinApplyOp,
	FullJoinApplyOp,
	SemiJoinApplyOp,
	AntiJoinApplyOp,
}

var JoinApplyOperators = [...]Operator{
	InnerJoinApplyOp,
	LeftJoinApplyOp,
	RightJoinApplyOp,
	FullJoinApplyOp,
	SemiJoinApplyOp,
	AntiJoinApplyOp,
}

var EnforcerOperators = [...]Operator{
	SortOp,
}
