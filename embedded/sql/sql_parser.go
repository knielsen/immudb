// Code generated by goyacc -l -o sql_parser.go sql_grammar.y. DO NOT EDIT.
package sql

import __yyfmt__ "fmt"

import "fmt"

func setResult(l yyLexer, stmts []SQLStmt) {
	l.(*lexer).result = stmts
}

type yySymType struct {
	yys           int
	stmts         []SQLStmt
	stmt          SQLStmt
	datasource    DataSource
	colsSpec      []*ColSpec
	colSpec       *ColSpec
	cols          []*ColSelector
	rows          []*RowSpec
	row           *RowSpec
	values        []ValueExp
	value         ValueExp
	id            string
	integer       uint64
	float         float64
	str           string
	boolean       bool
	blob          []byte
	sqlType       SQLValueType
	aggFn         AggregateFn
	ids           []string
	col           *ColSelector
	sel           Selector
	sels          []Selector
	distinct      bool
	ds            DataSource
	tableRef      *tableRef
	period        period
	openPeriod    *openPeriod
	periodInstant periodInstant
	joins         []*JoinSpec
	join          *JoinSpec
	joinType      JoinType
	exp           ValueExp
	binExp        ValueExp
	err           error
	ordcols       []*OrdCol
	opt_ord       bool
	logicOp       LogicOperator
	cmpOp         CmpOperator
	pparam        int
	update        *colUpdate
	updates       []*colUpdate
	onConflict    *OnConflictDo
}

const CREATE = 57346
const USE = 57347
const DATABASE = 57348
const SNAPSHOT = 57349
const SINCE = 57350
const AFTER = 57351
const BEFORE = 57352
const UNTIL = 57353
const TX = 57354
const OF = 57355
const TIMESTAMP = 57356
const TABLE = 57357
const UNIQUE = 57358
const INDEX = 57359
const ON = 57360
const ALTER = 57361
const ADD = 57362
const RENAME = 57363
const TO = 57364
const COLUMN = 57365
const PRIMARY = 57366
const KEY = 57367
const BEGIN = 57368
const TRANSACTION = 57369
const COMMIT = 57370
const ROLLBACK = 57371
const INSERT = 57372
const UPSERT = 57373
const INTO = 57374
const VALUES = 57375
const DELETE = 57376
const UPDATE = 57377
const SET = 57378
const CONFLICT = 57379
const DO = 57380
const NOTHING = 57381
const SELECT = 57382
const DISTINCT = 57383
const FROM = 57384
const JOIN = 57385
const HAVING = 57386
const WHERE = 57387
const GROUP = 57388
const BY = 57389
const LIMIT = 57390
const OFFSET = 57391
const ORDER = 57392
const ASC = 57393
const DESC = 57394
const AS = 57395
const UNION = 57396
const ALL = 57397
const NOT = 57398
const LIKE = 57399
const IF = 57400
const EXISTS = 57401
const IN = 57402
const IS = 57403
const AUTO_INCREMENT = 57404
const NULL = 57405
const CAST = 57406
const SCAST = 57407
const NPARAM = 57408
const PPARAM = 57409
const JOINTYPE = 57410
const LOP = 57411
const CMPOP = 57412
const IDENTIFIER = 57413
const TYPE = 57414
const INTEGER = 57415
const FLOAT = 57416
const VARCHAR = 57417
const BOOLEAN = 57418
const BLOB = 57419
const AGGREGATE_FUNC = 57420
const ERROR = 57421
const DOT = 57422
const STMT_SEPARATOR = 57423

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"CREATE",
	"USE",
	"DATABASE",
	"SNAPSHOT",
	"SINCE",
	"AFTER",
	"BEFORE",
	"UNTIL",
	"TX",
	"OF",
	"TIMESTAMP",
	"TABLE",
	"UNIQUE",
	"INDEX",
	"ON",
	"ALTER",
	"ADD",
	"RENAME",
	"TO",
	"COLUMN",
	"PRIMARY",
	"KEY",
	"BEGIN",
	"TRANSACTION",
	"COMMIT",
	"ROLLBACK",
	"INSERT",
	"UPSERT",
	"INTO",
	"VALUES",
	"DELETE",
	"UPDATE",
	"SET",
	"CONFLICT",
	"DO",
	"NOTHING",
	"SELECT",
	"DISTINCT",
	"FROM",
	"JOIN",
	"HAVING",
	"WHERE",
	"GROUP",
	"BY",
	"LIMIT",
	"OFFSET",
	"ORDER",
	"ASC",
	"DESC",
	"AS",
	"UNION",
	"ALL",
	"NOT",
	"LIKE",
	"IF",
	"EXISTS",
	"IN",
	"IS",
	"AUTO_INCREMENT",
	"NULL",
	"CAST",
	"SCAST",
	"NPARAM",
	"PPARAM",
	"JOINTYPE",
	"LOP",
	"CMPOP",
	"IDENTIFIER",
	"TYPE",
	"INTEGER",
	"FLOAT",
	"VARCHAR",
	"BOOLEAN",
	"BLOB",
	"AGGREGATE_FUNC",
	"ERROR",
	"DOT",
	"','",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'.'",
	"STMT_SEPARATOR",
	"'('",
	"')'",
	"'['",
	"']'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int16{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 74,
	57, 139,
	60, 139,
	-2, 127,
	-1, 189,
	43, 103,
	-2, 98,
	-1, 218,
	43, 103,
	-2, 100,
}

const yyPrivate = 57344

const yyLast = 378

var yyAct = [...]int16{
	73, 295, 60, 183, 140, 211, 235, 239, 88, 137,
	146, 175, 217, 106, 234, 98, 176, 157, 45, 6,
	79, 268, 101, 181, 226, 206, 181, 18, 181, 181,
	277, 272, 72, 254, 252, 271, 227, 182, 240, 150,
	255, 253, 222, 76, 205, 203, 78, 195, 194, 180,
	91, 87, 236, 89, 90, 241, 148, 59, 92, 110,
	82, 83, 84, 85, 86, 61, 202, 133, 199, 133,
	77, 159, 132, 117, 103, 81, 76, 128, 129, 78,
	130, 112, 131, 91, 87, 109, 89, 90, 97, 96,
	20, 92, 62, 82, 83, 84, 85, 86, 61, 61,
	142, 99, 110, 77, 62, 57, 124, 139, 81, 294,
	258, 288, 154, 149, 122, 123, 143, 153, 257, 161,
	162, 163, 164, 165, 166, 206, 151, 118, 119, 121,
	120, 196, 124, 174, 177, 124, 181, 144, 105, 251,
	122, 123, 231, 27, 28, 223, 188, 197, 186, 171,
	172, 189, 178, 118, 119, 121, 120, 108, 121, 120,
	173, 62, 138, 192, 204, 193, 190, 187, 191, 198,
	201, 62, 124, 257, 71, 107, 233, 124, 61, 209,
	122, 123, 220, 102, 179, 213, 123, 267, 158, 160,
	215, 155, 158, 118, 119, 121, 120, 152, 118, 119,
	121, 120, 113, 177, 221, 124, 65, 232, 26, 228,
	63, 34, 49, 238, 224, 127, 44, 145, 200, 230,
	266, 242, 229, 93, 126, 237, 118, 119, 121, 120,
	124, 244, 243, 250, 168, 111, 246, 177, 40, 169,
	249, 167, 170, 64, 115, 116, 55, 35, 259, 296,
	297, 260, 280, 212, 149, 264, 263, 184, 287, 275,
	262, 99, 274, 39, 269, 245, 104, 32, 276, 37,
	18, 285, 278, 270, 53, 281, 210, 208, 283, 31,
	30, 21, 247, 286, 135, 289, 134, 41, 42, 207,
	292, 293, 290, 94, 95, 76, 284, 298, 78, 2,
	299, 214, 91, 87, 114, 89, 90, 67, 10, 11,
	92, 66, 82, 83, 84, 85, 86, 61, 185, 147,
	38, 29, 77, 12, 43, 70, 69, 81, 47, 48,
	7, 22, 8, 9, 13, 14, 33, 141, 15, 16,
	23, 25, 24, 19, 18, 256, 100, 125, 248, 265,
	50, 51, 52, 279, 291, 225, 261, 75, 74, 273,
	219, 218, 216, 68, 46, 54, 36, 58, 56, 80,
	282, 136, 156, 17, 5, 4, 3, 1,
}

var yyPact = [...]int16{
	304, -1000, -1000, 3, -1000, -1000, -1000, 254, -1000, -1000,
	325, 137, 306, 248, 247, 225, 140, 193, 228, -1000,
	304, -1000, 180, 180, 180, 307, -1000, 145, 320, 141,
	140, 140, 140, 238, -1000, 191, 21, -1000, -1000, 139,
	187, 135, 293, 180, -1000, -1000, 315, 20, 20, 273,
	1, 0, 216, 112, 230, -1000, 224, -1000, 57, 104,
	-1000, -3, 22, -1000, 176, -7, 131, 286, -1000, 20,
	20, -1000, 239, 45, 159, -1000, 239, 239, -8, -1000,
	-1000, 239, -1000, -1000, -1000, -1000, -1000, -16, -1000, -1000,
	-1000, -1000, -21, -1000, 263, 261, 91, 91, 332, 239,
	56, -1000, 147, -1000, -32, 100, -1000, -1000, 126, 33,
	120, -1000, 117, -17, 118, -1000, -1000, 45, 239, 239,
	239, 239, 239, 239, 178, 182, 77, -1000, 116, 74,
	230, 71, 239, 239, 117, 113, -40, 55, -1000, -52,
	209, 301, 45, 332, 112, 239, 332, 320, 230, 104,
	-19, 104, -1000, -41, -42, -1000, 50, -1000, 75, 91,
	-20, 74, 74, 169, 169, 116, 144, -1000, 155, 239,
	-22, -1000, -44, -1000, 111, -45, 44, 45, -1000, 267,
	244, 108, 243, 204, 239, 283, 209, -1000, 45, 114,
	104, -47, -1000, -1000, -1000, -1000, 121, -66, -53, 91,
	-1000, 116, -13, -1000, 70, -1000, 239, 105, -36, -1000,
	-36, -1000, 239, 45, -33, 204, 216, -1000, 114, 222,
	-1000, -1000, 104, 257, -1000, 177, 66, -1000, -55, -48,
	-56, -49, 45, -1000, 92, -1000, 239, 37, 45, -1000,
	-1000, 91, -1000, 214, -1000, -32, -1000, -33, 158, -1000,
	124, -70, -1000, -1000, -1000, -1000, -1000, -36, 236, -54,
	-58, 218, 212, 332, -59, -1000, -1000, -1000, -1000, -1000,
	234, -1000, -1000, 202, 239, 90, 278, -1000, 232, 209,
	211, 45, 30, -1000, 239, -1000, 204, 90, 90, 45,
	-1000, 28, 198, -1000, 90, -1000, -1000, -1000, 198, -1000,
}

var yyPgo = [...]int16{
	0, 377, 299, 376, 375, 374, 19, 373, 372, 17,
	9, 7, 371, 370, 14, 6, 16, 11, 369, 8,
	20, 368, 367, 2, 366, 365, 10, 319, 18, 364,
	363, 174, 362, 12, 361, 360, 0, 15, 359, 358,
	357, 356, 3, 5, 355, 13, 354, 353, 1, 4,
	263, 349, 348, 347, 22, 346, 345, 343,
}

var yyR1 = [...]int8{
	0, 1, 2, 2, 57, 57, 3, 3, 3, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 50, 50, 11, 11, 5, 5, 5, 5,
	56, 56, 55, 55, 54, 12, 12, 14, 14, 15,
	10, 10, 13, 13, 17, 17, 16, 16, 18, 18,
	18, 18, 18, 18, 18, 18, 18, 18, 19, 8,
	8, 9, 44, 44, 51, 51, 52, 52, 52, 6,
	6, 7, 25, 25, 24, 24, 21, 21, 22, 22,
	20, 20, 20, 23, 23, 26, 26, 26, 27, 28,
	29, 29, 29, 30, 30, 30, 31, 31, 32, 32,
	33, 33, 34, 35, 35, 37, 37, 41, 41, 38,
	38, 42, 42, 43, 43, 47, 47, 49, 49, 46,
	46, 48, 48, 48, 45, 45, 45, 36, 36, 36,
	36, 36, 36, 36, 36, 39, 39, 39, 39, 53,
	53, 40, 40, 40, 40, 40, 40, 40, 40,
}

var yyR2 = [...]int8{
	0, 1, 2, 3, 0, 1, 1, 1, 1, 2,
	1, 1, 1, 4, 2, 3, 3, 11, 8, 9,
	6, 8, 0, 3, 1, 3, 9, 8, 7, 8,
	0, 4, 1, 3, 3, 0, 1, 1, 3, 3,
	1, 3, 1, 3, 0, 1, 1, 3, 1, 1,
	1, 1, 1, 6, 1, 1, 1, 1, 4, 1,
	3, 5, 0, 3, 0, 1, 0, 1, 2, 1,
	4, 13, 0, 1, 0, 1, 1, 1, 2, 4,
	1, 4, 4, 1, 3, 3, 4, 2, 1, 2,
	0, 2, 2, 0, 2, 2, 2, 1, 0, 1,
	1, 2, 6, 0, 1, 0, 2, 0, 3, 0,
	2, 0, 2, 0, 2, 0, 3, 0, 4, 2,
	4, 0, 1, 1, 0, 1, 2, 1, 1, 2,
	2, 4, 4, 6, 6, 1, 1, 3, 3, 0,
	1, 3, 3, 3, 3, 3, 3, 3, 4,
}

var yyChk = [...]int16{
	-1000, -1, -2, -3, -4, -5, -6, 26, 28, 29,
	4, 5, 19, 30, 31, 34, 35, -7, 40, -57,
	87, 27, 6, 15, 17, 16, 71, 6, 7, 15,
	32, 32, 42, -27, 71, 54, -24, 41, -2, -50,
	58, -50, -50, 17, 71, -28, -29, 8, 9, 71,
	-27, -27, -27, 36, -25, 55, -21, 84, -22, -20,
	-23, 78, 71, 71, 56, 71, 18, -50, -30, 11,
	10, -31, 12, -36, -39, -40, 56, 83, 59, -20,
	-18, 88, 73, 74, 75, 76, 77, 64, -19, 66,
	67, 63, 71, -31, 20, 21, 88, 88, -37, 45,
	-55, -54, 71, -6, 42, 81, -45, 71, 53, 88,
	80, 59, 88, 71, 18, -31, -31, -36, 82, 83,
	85, 84, 69, 70, 61, -53, 65, 56, -36, -36,
	88, -36, 88, 88, 23, 23, -12, -10, 71, -10,
	-49, 5, -36, -37, 81, 70, -26, -27, 88, -19,
	71, -20, 71, 84, -23, 71, -8, -9, 71, 88,
	71, -36, -36, -36, -36, -36, -36, 63, 56, 57,
	60, 72, -6, 89, -36, -17, -16, -36, -9, 71,
	89, 81, 89, -42, 48, 17, -49, -54, -36, -49,
	-28, -6, -45, -45, 89, 89, 81, 72, -10, 88,
	63, -36, 88, 89, 53, 89, 81, 22, 33, 71,
	33, -43, 49, -36, 18, -42, -32, -33, -34, -35,
	68, -45, 89, 24, -9, -44, 90, 89, -10, -6,
	-16, 72, -36, 71, -14, -15, 88, -14, -36, -11,
	71, 88, -43, -37, -33, 43, -45, 25, -52, 63,
	56, 73, 89, 89, 89, 89, -56, 81, 18, -17,
	-10, -41, 46, -26, -11, -51, 62, 63, 91, -15,
	37, 89, 89, -38, 44, 47, -49, 89, 38, -47,
	50, -36, -13, -23, 18, 39, -42, 47, 81, -36,
	-43, -46, -23, -23, 81, -48, 51, 52, -23, -48,
}

var yyDef = [...]int16{
	0, -2, 1, 4, 6, 7, 8, 10, 11, 12,
	0, 0, 0, 0, 0, 0, 0, 69, 74, 2,
	5, 9, 22, 22, 22, 0, 14, 0, 90, 0,
	0, 0, 0, 0, 88, 72, 0, 75, 3, 0,
	0, 0, 0, 22, 15, 16, 93, 0, 0, 0,
	0, 0, 105, 0, 0, 73, 0, 76, 77, 124,
	80, 0, 83, 13, 0, 0, 0, 0, 89, 0,
	0, 91, 0, 97, -2, 128, 0, 0, 0, 135,
	136, 0, 48, 49, 50, 51, 52, 0, 54, 55,
	56, 57, 83, 92, 0, 0, 35, 0, 117, 0,
	105, 32, 0, 70, 0, 0, 78, 125, 0, 0,
	0, 23, 0, 0, 0, 94, 95, 96, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 140, 129, 130,
	0, 0, 0, 44, 0, 0, 0, 36, 40, 0,
	111, 0, 106, 117, 0, 0, 117, 90, 0, 124,
	88, 124, 126, 0, 0, 84, 0, 59, 0, 0,
	0, 141, 142, 143, 144, 145, 146, 147, 0, 0,
	0, 138, 0, 137, 0, 0, 45, 46, 20, 0,
	0, 0, 0, 113, 0, 0, 111, 33, 34, -2,
	124, 0, 87, 79, 81, 82, 0, 62, 0, 0,
	148, 131, 0, 132, 0, 58, 0, 0, 0, 41,
	0, 28, 0, 112, 0, 113, 105, 99, -2, 0,
	104, 85, 124, 0, 60, 66, 0, 18, 0, 0,
	0, 0, 47, 21, 30, 37, 44, 27, 114, 118,
	24, 0, 29, 107, 101, 0, 86, 0, 64, 67,
	0, 0, 19, 133, 134, 53, 26, 0, 0, 0,
	0, 109, 0, 117, 0, 61, 65, 68, 63, 38,
	0, 39, 25, 115, 0, 0, 0, 17, 0, 111,
	0, 110, 108, 42, 0, 31, 113, 0, 0, 102,
	71, 116, 121, 43, 0, 119, 122, 123, 121, 120,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	88, 89, 84, 82, 81, 83, 86, 85, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 90, 3, 91,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 87,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmts = yyDollar[1].stmts
			setResult(yylex, yyDollar[1].stmts)
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = []SQLStmt{yyDollar[1].stmt}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmts = append([]SQLStmt{yyDollar[1].stmt}, yyDollar[3].stmts...)
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = &BeginTransactionStmt{}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &BeginTransactionStmt{}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &CommitStmt{}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &RollbackStmt{}
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmt = &CreateDatabaseStmt{ifNotExists: yyDollar[3].boolean, DB: yyDollar[4].id}
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = &UseDatabaseStmt{DB: yyDollar[2].id}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &UseDatabaseStmt{DB: yyDollar[3].id}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &UseSnapshotStmt{period: yyDollar[3].period}
		}
	case 17:
		yyDollar = yyS[yypt-11 : yypt+1]
		{
			yyVAL.stmt = &CreateTableStmt{ifNotExists: yyDollar[3].boolean, table: yyDollar[4].id, colsSpec: yyDollar[6].colsSpec, pkColNames: yyDollar[10].ids}
		}
	case 18:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &CreateIndexStmt{ifNotExists: yyDollar[3].boolean, table: yyDollar[5].id, cols: yyDollar[7].ids}
		}
	case 19:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.stmt = &CreateIndexStmt{unique: true, ifNotExists: yyDollar[4].boolean, table: yyDollar[6].id, cols: yyDollar[8].ids}
		}
	case 20:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.stmt = &AddColumnStmt{table: yyDollar[3].id, colSpec: yyDollar[6].colSpec}
		}
	case 21:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &RenameColumnStmt{table: yyDollar[3].id, oldName: yyDollar[6].id, newName: yyDollar[8].id}
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ids = []string{yyDollar[1].id}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ids = yyDollar[2].ids
		}
	case 26:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.stmt = &UpsertIntoStmt{isInsert: true, tableRef: yyDollar[3].tableRef, cols: yyDollar[5].ids, rows: yyDollar[8].rows, onConflict: yyDollar[9].onConflict}
		}
	case 27:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &UpsertIntoStmt{tableRef: yyDollar[3].tableRef, cols: yyDollar[5].ids, rows: yyDollar[8].rows}
		}
	case 28:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.stmt = &DeleteFromStmt{tableRef: yyDollar[3].tableRef, where: yyDollar[4].exp, indexOn: yyDollar[5].ids, limit: yyDollar[6].exp, offset: yyDollar[7].exp}
		}
	case 29:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &UpdateStmt{tableRef: yyDollar[2].tableRef, updates: yyDollar[4].updates, where: yyDollar[5].exp, indexOn: yyDollar[6].ids, limit: yyDollar[7].exp, offset: yyDollar[8].exp}
		}
	case 30:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.onConflict = nil
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.onConflict = &OnConflictDo{}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.updates = []*colUpdate{yyDollar[1].update}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.updates = append(yyDollar[1].updates, yyDollar[3].update)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.update = &colUpdate{col: yyDollar[1].id, op: yyDollar[2].cmpOp, val: yyDollar[3].exp}
		}
	case 35:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.ids = nil
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ids = yyDollar[1].ids
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.rows = []*RowSpec{yyDollar[1].row}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.rows = append(yyDollar[1].rows, yyDollar[3].row)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.row = &RowSpec{Values: yyDollar[2].values}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ids = []string{yyDollar[1].id}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ids = append(yyDollar[1].ids, yyDollar[3].id)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.cols = []*ColSelector{yyDollar[1].col}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.cols = append(yyDollar[1].cols, yyDollar[3].col)
		}
	case 44:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.values = nil
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.values = yyDollar[1].values
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.values = []ValueExp{yyDollar[1].exp}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].exp)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Integer{val: int64(yyDollar[1].integer)}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Float64{val: float64(yyDollar[1].float)}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Varchar{val: yyDollar[1].str}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Bool{val: yyDollar[1].boolean}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Blob{val: yyDollar[1].blob}
		}
	case 53:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.value = &Cast{val: yyDollar[3].exp, t: yyDollar[5].sqlType}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = yyDollar[1].value
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Param{id: yyDollar[1].id}
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Param{id: fmt.Sprintf("param%d", yyDollar[1].pparam), pos: yyDollar[1].pparam}
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &NullValue{t: AnyType}
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.value = &FnCall{fn: yyDollar[1].id, params: yyDollar[3].values}
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.colsSpec = []*ColSpec{yyDollar[1].colSpec}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.colsSpec = append(yyDollar[1].colsSpec, yyDollar[3].colSpec)
		}
	case 61:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.colSpec = &ColSpec{colName: yyDollar[1].id, colType: yyDollar[2].sqlType, maxLen: int(yyDollar[3].integer), notNull: yyDollar[4].boolean, autoIncrement: yyDollar[5].boolean}
		}
	case 62:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.integer = 0
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.integer = yyDollar[2].integer
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 66:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 70:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmt = &UnionStmt{
				distinct: yyDollar[3].distinct,
				left:     yyDollar[1].stmt.(DataSource),
				right:    yyDollar[4].stmt.(DataSource),
			}
		}
	case 71:
		yyDollar = yyS[yypt-13 : yypt+1]
		{
			yyVAL.stmt = &SelectStmt{
				distinct:  yyDollar[2].distinct,
				selectors: yyDollar[3].sels,
				ds:        yyDollar[5].ds,
				indexOn:   yyDollar[6].ids,
				joins:     yyDollar[7].joins,
				where:     yyDollar[8].exp,
				groupBy:   yyDollar[9].cols,
				having:    yyDollar[10].exp,
				orderBy:   yyDollar[11].ordcols,
				limit:     yyDollar[12].exp,
				offset:    yyDollar[13].exp,
			}
		}
	case 72:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.distinct = true
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.distinct = false
		}
	case 74:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.distinct = false
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.distinct = true
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.sels = nil
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.sels = yyDollar[1].sels
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyDollar[1].sel.setAlias(yyDollar[2].id)
			yyVAL.sels = []Selector{yyDollar[1].sel}
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyDollar[3].sel.setAlias(yyDollar[4].id)
			yyVAL.sels = append(yyDollar[1].sels, yyDollar[3].sel)
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.sel = yyDollar[1].col
		}
	case 81:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.sel = &AggColSelector{aggFn: yyDollar[1].aggFn, col: "*"}
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.sel = &AggColSelector{aggFn: yyDollar[1].aggFn, table: yyDollar[3].col.table, col: yyDollar[3].col.col}
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.col = &ColSelector{col: yyDollar[1].id}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.col = &ColSelector{table: yyDollar[1].id, col: yyDollar[3].id}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyDollar[1].tableRef.period = yyDollar[2].period
			yyDollar[1].tableRef.as = yyDollar[3].id
			yyVAL.ds = yyDollar[1].tableRef
		}
	case 86:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyDollar[2].stmt.(*SelectStmt).as = yyDollar[4].id
			yyVAL.ds = yyDollar[2].stmt.(DataSource)
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.ds = &FnDataSourceStmt{fnCall: yyDollar[1].value.(*FnCall), as: yyDollar[2].id}
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.tableRef = &tableRef{table: yyDollar[1].id}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.period = period{start: yyDollar[1].openPeriod, end: yyDollar[2].openPeriod}
		}
	case 90:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.openPeriod = nil
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.openPeriod = &openPeriod{inclusive: true, instant: yyDollar[2].periodInstant}
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.openPeriod = &openPeriod{instant: yyDollar[2].periodInstant}
		}
	case 93:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.openPeriod = nil
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.openPeriod = &openPeriod{inclusive: true, instant: yyDollar[2].periodInstant}
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.openPeriod = &openPeriod{instant: yyDollar[2].periodInstant}
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.periodInstant = periodInstant{instantType: txInstant, exp: yyDollar[2].exp}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.periodInstant = periodInstant{instantType: timeInstant, exp: yyDollar[1].exp}
		}
	case 98:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.joins = nil
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joins = yyDollar[1].joins
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joins = []*JoinSpec{yyDollar[1].join}
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.joins = append([]*JoinSpec{yyDollar[1].join}, yyDollar[2].joins...)
		}
	case 102:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.join = &JoinSpec{joinType: yyDollar[1].joinType, ds: yyDollar[3].ds, indexOn: yyDollar[4].ids, cond: yyDollar[6].exp}
		}
	case 103:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.joinType = InnerJoin
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joinType = yyDollar[1].joinType
		}
	case 105:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exp = nil
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = yyDollar[2].exp
		}
	case 107:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.cols = nil
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.cols = yyDollar[3].cols
		}
	case 109:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exp = nil
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = yyDollar[2].exp
		}
	case 111:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exp = nil
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = yyDollar[2].exp
		}
	case 113:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exp = nil
		}
	case 114:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = yyDollar[2].exp
		}
	case 115:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.ordcols = nil
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ordcols = yyDollar[3].ordcols
		}
	case 117:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.ids = nil
		}
	case 118:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.ids = yyDollar[4].ids
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.ordcols = []*OrdCol{{sel: yyDollar[1].col, descOrder: yyDollar[2].opt_ord}}
		}
	case 120:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.ordcols = append(yyDollar[1].ordcols, &OrdCol{sel: yyDollar[3].col, descOrder: yyDollar[4].opt_ord})
		}
	case 121:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.opt_ord = false
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.opt_ord = false
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.opt_ord = true
		}
	case 124:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.id = ""
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.id = yyDollar[1].id
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.id = yyDollar[2].id
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exp = yyDollar[1].exp
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exp = yyDollar[1].binExp
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = &NotBoolExp{exp: yyDollar[2].exp}
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.exp = &NumExp{left: &Integer{val: 0}, op: SUBSOP, right: yyDollar[2].exp}
		}
	case 131:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exp = &LikeBoolExp{val: yyDollar[1].exp, notLike: yyDollar[2].boolean, pattern: yyDollar[4].exp}
		}
	case 132:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exp = &ExistsBoolExp{q: (yyDollar[3].stmt).(*SelectStmt)}
		}
	case 133:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.exp = &InSubQueryExp{val: yyDollar[1].exp, notIn: yyDollar[2].boolean, q: yyDollar[5].stmt.(*SelectStmt)}
		}
	case 134:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.exp = &InListExp{val: yyDollar[1].exp, notIn: yyDollar[2].boolean, values: yyDollar[5].values}
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exp = yyDollar[1].sel
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exp = yyDollar[1].value
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.exp = yyDollar[2].exp
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.exp = &Cast{val: yyDollar[1].exp, t: yyDollar[3].sqlType}
		}
	case 139:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].exp, op: ADDOP, right: yyDollar[3].exp}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].exp, op: SUBSOP, right: yyDollar[3].exp}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].exp, op: DIVOP, right: yyDollar[3].exp}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].exp, op: MULTOP, right: yyDollar[3].exp}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &BinBoolExp{left: yyDollar[1].exp, op: yyDollar[2].logicOp, right: yyDollar[3].exp}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &CmpBoolExp{left: yyDollar[1].exp, op: yyDollar[2].cmpOp, right: yyDollar[3].exp}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &CmpBoolExp{left: yyDollar[1].exp, op: EQ, right: &NullValue{t: AnyType}}
		}
	case 148:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.binExp = &CmpBoolExp{left: yyDollar[1].exp, op: NE, right: &NullValue{t: AnyType}}
		}
	}
	goto yystack /* stack new state and value */
}
