//line hawk.y:2

// Package hawkc is the compiler for the Hawk language.
package hawkc

import __yyfmt__ "fmt"

//line hawk.y:3
import (
	"bufio"
	"fmt"
	"io"

	"github.com/mibk/hawk/scan"
	"github.com/mibk/hawk/value"
)

var (
	progName string
	ast      *Program

	defaultAction = &BlockStmt{[]Stmt{&PrintStmt{Fun: "print"}}}
)

//line hawk.y:22
type yySymType struct {
	yys       int
	sym       string
	val       value.Value
	symlist   []string
	decl      Decl
	decllist  []Decl
	expr      Expr
	exprlist  []Expr
	stmt      Stmt
	stmtlist  []Stmt
	blockstmt *BlockStmt
}

const IDENT = 57346
const BOOL = 57347
const STRING = 57348
const PRINT = 57349
const NUM = 57350
const BEGIN = 57351
const END = 57352
const IF = 57353
const ELSE = 57354
const FOR = 57355
const IN = 57356
const BREAK = 57357
const CONTINUE = 57358
const INC = 57359
const DEC = 57360
const ADDEQ = 57361
const SUBEQ = 57362
const MULEQ = 57363
const DIVEQ = 57364
const MODEQ = 57365
const CONCATEQ = 57366
const FUNC = 57367
const RETURN = 57368
const OROR = 57369
const ANDAND = 57370
const EQ = 57371
const NE = 57372
const LE = 57373
const GE = 57374
const NOTMATCH = 57375

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"BOOL",
	"STRING",
	"PRINT",
	"NUM",
	"BEGIN",
	"END",
	"IF",
	"ELSE",
	"FOR",
	"IN",
	"BREAK",
	"CONTINUE",
	"INC",
	"DEC",
	"ADDEQ",
	"SUBEQ",
	"MULEQ",
	"DIVEQ",
	"MODEQ",
	"CONCATEQ",
	"FUNC",
	"RETURN",
	"'?'",
	"':'",
	"OROR",
	"ANDAND",
	"EQ",
	"NE",
	"LE",
	"GE",
	"'<'",
	"'>'",
	"'~'",
	"NOTMATCH",
	"'.'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"';'",
	"'('",
	"')'",
	"','",
	"'{'",
	"'}'",
	"'|'",
	"'='",
	"'['",
	"']'",
	"'$'",
	"'!'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line hawk.y:495

// Compile compiles a Hawk program (name) from src. It is not safe
// for concurrent use.
func Compile(name string, src io.Reader) (*Program, error) {
	progName = name
	sc := new(scan.Scanner)
	ast = NewProgram(sc)
	lexlineno = 1
	nlsemi = false
	l := &yyLex{reader: bufio.NewReader(src)}
	yyParse(l)
	analyse(ast, sc)
	return ast, l.err
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 7,
	49, 74,
	-2, 7,
	-1, 23,
	49, 73,
	-2, 1,
	-1, 64,
	17, 41,
	18, 41,
	19, 41,
	20, 41,
	21, 41,
	22, 41,
	23, 41,
	24, 41,
	52, 41,
	-2, 82,
	-1, 65,
	17, 42,
	18, 42,
	19, 42,
	20, 42,
	21, 42,
	22, 42,
	23, 42,
	24, 42,
	52, 42,
	-2, 87,
	-1, 74,
	49, 73,
	-2, 43,
	-1, 120,
	17, 41,
	18, 41,
	19, 41,
	20, 41,
	21, 41,
	22, 41,
	23, 41,
	24, 41,
	52, 41,
	-2, 82,
	-1, 122,
	49, 74,
	-2, 21,
}

const yyPrivate = 57344

const yyLast = 588

var yyAct = [...]int{

	55, 7, 61, 60, 118, 66, 54, 97, 128, 24,
	6, 43, 44, 59, 20, 15, 14, 113, 13, 22,
	50, 56, 154, 153, 7, 102, 62, 75, 76, 77,
	78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
	88, 89, 90, 91, 131, 65, 25, 51, 51, 147,
	16, 17, 143, 96, 112, 52, 19, 99, 95, 73,
	150, 151, 98, 21, 142, 12, 18, 35, 36, 37,
	38, 39, 115, 152, 117, 122, 92, 167, 121, 116,
	169, 51, 114, 148, 145, 119, 101, 23, 112, 64,
	15, 14, 72, 13, 65, 3, 73, 25, 74, 129,
	69, 70, 62, 126, 134, 135, 136, 137, 138, 139,
	140, 71, 156, 96, 99, 132, 37, 38, 39, 57,
	144, 65, 146, 133, 149, 16, 17, 161, 159, 125,
	45, 19, 11, 100, 25, 1, 58, 68, 21, 67,
	12, 18, 164, 155, 129, 46, 115, 63, 158, 47,
	48, 49, 2, 160, 162, 163, 157, 124, 5, 166,
	4, 168, 165, 0, 0, 0, 0, 0, 62, 0,
	171, 121, 170, 172, 173, 0, 0, 64, 15, 14,
	72, 13, 0, 0, 73, 0, 74, 65, 69, 70,
	41, 42, 40, 35, 36, 37, 38, 39, 0, 71,
	40, 35, 36, 37, 38, 39, 0, 0, 0, 0,
	0, 0, 0, 16, 17, 0, 0, 0, 0, 19,
	0, 0, 0, 0, 0, 0, 21, 0, 12, 18,
	26, 0, 27, 28, 29, 30, 31, 32, 33, 34,
	41, 42, 40, 35, 36, 37, 38, 39, 0, 0,
	0, 0, 0, 0, 0, 0, 26, 130, 27, 28,
	29, 30, 31, 32, 33, 34, 41, 42, 40, 35,
	36, 37, 38, 39, 0, 120, 15, 14, 72, 13,
	0, 0, 73, 127, 74, 0, 69, 70, 0, 0,
	0, 0, 0, 0, 20, 15, 14, 71, 13, 8,
	9, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 16, 17, 0, 0, 10, 0, 19, 0, 0,
	0, 0, 0, 0, 21, 0, 12, 18, 0, 0,
	16, 17, 0, 0, 0, 0, 19, 20, 15, 14,
	0, 13, 0, 21, 0, 12, 18, 26, 0, 27,
	28, 29, 30, 31, 32, 33, 34, 41, 42, 40,
	35, 36, 37, 38, 39, 0, 20, 15, 14, 25,
	13, 0, 0, 16, 17, 0, 0, 0, 0, 19,
	0, 0, 0, 0, 0, 0, 21, 141, 12, 18,
	20, 15, 14, 0, 13, 0, 0, 0, 0, 0,
	0, 0, 16, 17, 0, 0, 0, 0, 19, 94,
	20, 15, 14, 0, 13, 21, 0, 12, 18, 0,
	0, 0, 0, 0, 0, 0, 16, 17, 0, 0,
	0, 0, 19, 0, 0, 0, 0, 0, 0, 21,
	53, 12, 18, 0, 0, 0, 16, 17, 0, 20,
	15, 14, 19, 13, 0, 0, 0, 0, 0, 21,
	0, 12, 18, 26, 0, 27, 28, 29, 30, 31,
	32, 33, 34, 41, 42, 40, 35, 36, 37, 38,
	39, 0, 0, 93, 0, 16, 17, 0, 0, 0,
	0, 19, 0, 0, 0, 0, 0, 0, 21, 0,
	0, 18, 26, 123, 27, 28, 29, 30, 31, 32,
	33, 34, 41, 42, 40, 35, 36, 37, 38, 39,
	26, 0, 27, 28, 29, 30, 31, 32, 33, 34,
	41, 42, 40, 35, 36, 37, 38, 39, 110, 111,
	104, 105, 106, 107, 108, 109, 28, 29, 30, 31,
	32, 33, 34, 41, 42, 40, 35, 36, 37, 38,
	39, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 103, 29, 30, 31, 32, 33, 34,
	41, 42, 40, 35, 36, 37, 38, 39,
}
var yyPact = [...]int{

	290, -1000, 42, -1000, -1000, -1000, -3, 493, -3, -3,
	126, -1000, 445, -1000, -1000, -1000, 445, 445, 445, 406,
	2, 386, -32, 290, -1000, 85, 406, 406, 406, 406,
	406, 406, 406, 406, 406, 406, 406, 406, 406, 406,
	406, 406, 406, -1000, -1000, 30, -1000, -1000, -1000, -1000,
	436, 362, 406, -1000, 14, 493, 406, -1000, 41, -26,
	-1000, -1000, 493, 521, 1, -36, -1000, -1000, -1000, -1000,
	-1000, 406, 406, 406, 271, 475, 516, 543, 153, 153,
	153, 153, 153, 153, 74, 74, -1000, -1000, -1000, 27,
	161, 161, 125, -1000, -1000, 14, 229, -46, 406, 203,
	-6, 85, 117, 406, 406, 406, 406, 406, 406, 406,
	-1000, -1000, 333, 10, -1000, 493, 4, 320, 39, -3,
	35, -1000, 493, 406, 13, -1000, 26, -1000, -1000, 493,
	-1000, -1000, -26, -1000, 493, 493, 493, 493, 493, 493,
	493, -29, -30, 406, 100, 406, -1000, 406, 124, 493,
	-3, 123, -1000, 406, 406, -1000, 48, 32, 320, 66,
	-1000, -1000, 493, 493, -1000, -1000, -1000, 173, -1000, 406,
	-3, 320, -1000, -1000,
}
var yyPgo = [...]int{

	0, 95, 160, 158, 157, 152, 0, 10, 132, 19,
	147, 6, 13, 3, 4, 5, 143, 142, 139, 137,
	2, 136, 135, 133, 7,
}
var yyR1 = [...]int{

	0, 22, 5, 5, 1, 1, 2, 2, 2, 2,
	3, 4, 4, 4, 20, 21, 21, 21, 12, 12,
	12, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 10, 10, 14, 14, 15, 16, 16, 17, 17,
	18, 18, 19, 19, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 7, 7, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 9, 9,
	11, 11, 23, 23, 24, 24,
}
var yyR2 = [...]int{

	0, 2, 1, 3, 1, 1, 2, 1, 2, 2,
	6, 0, 1, 3, 4, 0, 1, 3, 1, 1,
	3, 1, 3, 5, 5, 3, 3, 3, 3, 3,
	3, 2, 2, 1, 1, 1, 1, 1, 2, 2,
	1, 1, 1, 0, 1, 4, 0, 2, 1, 1,
	7, 3, 5, 7, 1, 5, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 0, 1, 1, 1, 1, 2, 2,
	2, 3, 1, 3, 5, 2, 4, 1, 4, 4,
	1, 3, 0, 1, 0, 1,
}
var yyChk = [...]int{

	-1000, -22, -5, -1, -2, -3, -7, -6, 9, 10,
	25, -8, 55, 8, 6, 5, 40, 41, 56, 46,
	4, 53, -9, 45, -20, 49, 27, 29, 30, 31,
	32, 33, 34, 35, 36, 40, 41, 42, 43, 44,
	39, 37, 38, -20, -20, 4, -8, -8, -8, -8,
	-6, 46, 53, 54, -11, -6, 53, -1, -21, -12,
	-13, -20, -6, -10, 4, -9, -15, -18, -19, 15,
	16, 26, 7, 11, 13, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, 46, 47, 47, -11, -6, -24, 48, -6,
	-23, 45, 51, 52, 19, 20, 21, 22, 23, 24,
	17, 18, 53, 53, -7, -6, -11, -6, -14, -7,
	4, -13, -6, 28, -4, 4, -24, 54, 54, -6,
	54, 50, -12, 6, -6, -6, -6, -6, -6, -6,
	-6, 54, 54, 48, -20, 45, -20, 14, 48, -6,
	47, 48, 47, 52, 52, -16, 12, -7, -6, 4,
	-20, 4, -6, -6, -17, -15, -20, 45, -20, 14,
	-14, -6, -20, -20,
}
var yyDef = [...]int{

	73, -2, 0, 2, 4, 5, 0, -2, 0, 0,
	0, 54, 0, 75, 76, 77, 0, 0, 0, 0,
	82, 0, 87, -2, 6, 15, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 8, 9, 0, 56, 78, 79, 80,
	0, 0, 0, 85, 94, 90, 0, 3, 92, 16,
	18, 19, 21, 0, -2, -2, 33, 34, 35, 36,
	37, 73, 40, 0, -2, 0, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71, 72, 11, 81, 83, 94, 0, 0, 95, 0,
	0, 93, 0, 0, 0, 0, 0, 0, 0, 0,
	31, 32, 0, 0, 38, 74, 39, 0, 0, 0,
	-2, 44, -2, 0, 0, 12, 0, 88, 86, 91,
	89, 14, 17, 20, 22, 25, 26, 27, 28, 29,
	30, 0, 0, 0, 46, 73, 51, 0, 0, 55,
	0, 0, 84, 0, 0, 45, 0, 0, 0, 0,
	10, 13, 23, 24, 47, 48, 49, 43, 52, 0,
	0, 0, 50, 53,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 56, 3, 3, 55, 44, 3, 3,
	46, 47, 42, 40, 48, 41, 39, 43, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 28, 45,
	35, 52, 36, 27, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 53, 3, 54, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 49, 51, 50, 37,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 29, 30, 31, 32, 33,
	34, 38,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

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
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
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
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
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
	yyn = yyPact[yystate]
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
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
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
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
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
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
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

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:66
		{
			for _, d := range yyDollar[1].decllist {
				switch d := d.(type) {
				case *BeginAction:
					ast.Begins = append(ast.Begins, d)
				case *EndAction:
					ast.Ends = append(ast.Ends, d)
				case *PatternAction:
					ast.Pactions = append(ast.Pactions, d)
				case *FuncDecl:
					ast.funcs[d.Name] = d
				default:
					panic(fmt.Sprintf("unexpected type: %T", d))
				}
			}
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:85
		{
			yyVAL.decllist = []Decl{yyDollar[1].decl}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:89
		{
			yyVAL.decllist = append(yyDollar[1].decllist, yyDollar[3].decl)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:95
		{
			yyVAL.decl = yyDollar[1].decl
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:99
		{
			yyVAL.decl = yyDollar[1].decl
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:105
		{
			yyVAL.decl = &PatternAction{yyDollar[1].expr, yyDollar[2].blockstmt}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:109
		{
			yyVAL.decl = &PatternAction{yyDollar[1].expr, defaultAction}
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:113
		{
			yyVAL.decl = &BeginAction{yyDollar[2].blockstmt}
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:117
		{
			yyVAL.decl = &EndAction{yyDollar[2].blockstmt}
		}
	case 10:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line hawk.y:123
		{
			yyVAL.decl = &FuncDecl{&FuncScope{}, yyDollar[2].sym, yyDollar[4].symlist, yyDollar[6].blockstmt}
		}
	case 11:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:128
		{
			yyVAL.symlist = nil
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:132
		{
			yyVAL.symlist = []string{yyDollar[1].sym}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:136
		{
			yyVAL.symlist = append(yyDollar[1].symlist, yyDollar[3].sym)
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:142
		{
			yyVAL.blockstmt = &BlockStmt{yyDollar[2].stmtlist}
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:147
		{
			yyVAL.stmtlist = nil
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:151
		{
			yyVAL.stmtlist = []Stmt{yyDollar[1].stmt}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:155
		{
			yyVAL.stmtlist = append(yyDollar[1].stmtlist, yyDollar[3].stmt)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:161
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:165
		{
			yyVAL.stmt = yyDollar[1].blockstmt
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:169
		{
			yyVAL.stmt = &PipeStmt{genDebugInfo(), yyDollar[1].stmt, yyDollar[3].sym}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:175
		{
			yyVAL.stmt = &ExprStmt{yyDollar[1].expr}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:179
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:186
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, &IndexExpr{genDebugInfo(), &Ident{Name: yyDollar[1].sym}, nil}, yyDollar[5].expr}
		}
	case 24:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:190
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, &IndexExpr{genDebugInfo(), yyDollar[1].expr, nil}, yyDollar[5].expr}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:195
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, yyDollar[1].expr, &BinaryExpr{genDebugInfo(), Add, yyDollar[1].expr, yyDollar[3].expr}}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:199
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, yyDollar[1].expr, &BinaryExpr{genDebugInfo(), Sub, yyDollar[1].expr, yyDollar[3].expr}}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:203
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, yyDollar[1].expr, &BinaryExpr{genDebugInfo(), Mul, yyDollar[1].expr, yyDollar[3].expr}}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:207
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, yyDollar[1].expr, &BinaryExpr{genDebugInfo(), Div, yyDollar[1].expr, yyDollar[3].expr}}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:211
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, yyDollar[1].expr, &BinaryExpr{genDebugInfo(), Mod, yyDollar[1].expr, yyDollar[3].expr}}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:215
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, yyDollar[1].expr, &BinaryExpr{genDebugInfo(), Concat, yyDollar[1].expr, yyDollar[3].expr}}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:219
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, yyDollar[1].expr, &BinaryExpr{genDebugInfo(), Add, yyDollar[1].expr, BasicLit{value.NewNumber(1)}}}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:223
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, yyDollar[1].expr, &BinaryExpr{genDebugInfo(), Sub, yyDollar[1].expr, BasicLit{value.NewNumber(1)}}}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:227
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:231
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:235
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:239
		{
			yyVAL.stmt = &StatusStmt{StatusBreak}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:243
		{
			yyVAL.stmt = &StatusStmt{StatusContinue}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:247
		{
			yyVAL.stmt = &ReturnStmt{X: yyDollar[2].expr}
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:251
		{
			yyVAL.stmt = &PrintStmt{genDebugInfo(), nil, yyDollar[1].sym, yyDollar[2].exprlist}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:255
		{
			yyVAL.stmt = &PrintStmt{genDebugInfo(), nil, yyDollar[1].sym, nil}
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:261
		{
			yyVAL.expr = &Ident{ast, yyDollar[1].sym}
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:265
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 43:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:270
		{
			yyVAL.stmt = nil
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:274
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:280
		{
			yyVAL.stmt = &IfStmt{genDebugInfo(), yyDollar[2].expr, yyDollar[3].blockstmt, yyDollar[4].stmt}
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:285
		{
			yyVAL.stmt = nil
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:289
		{
			yyVAL.stmt = yyDollar[2].stmt
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:295
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:299
		{
			yyVAL.stmt = yyDollar[1].blockstmt
		}
	case 50:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line hawk.y:305
		{
			yyVAL.stmt = &ForStmt{genDebugInfo(), yyDollar[2].stmt, yyDollar[4].expr, yyDollar[6].stmt, yyDollar[7].blockstmt}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:309
		{
			yyVAL.stmt = &ForStmt{genDebugInfo(), nil, yyDollar[2].expr, nil, yyDollar[3].blockstmt}
		}
	case 52:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:315
		{
			yyVAL.stmt = &ForeachStmt{genDebugInfo(), &Ident{Name: yyDollar[2].sym}, nil, yyDollar[4].expr, yyDollar[5].blockstmt}
		}
	case 53:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line hawk.y:319
		{
			yyVAL.stmt = &ForeachStmt{genDebugInfo(), &Ident{Name: yyDollar[2].sym}, &Ident{Name: yyDollar[4].sym}, yyDollar[6].expr, yyDollar[7].blockstmt}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:326
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 55:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:330
		{
			yyVAL.expr = &TernaryExpr{genDebugInfo(), yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr}
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:334
		{
			yyVAL.expr = &FieldExpr{genDebugInfo(), nil, yyDollar[2].expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:338
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), OrOr, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:342
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), AndAnd, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:346
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Eq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:350
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), NotEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:354
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), LtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:358
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), GtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:362
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Lt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:366
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Gt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:370
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Add, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:374
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Sub, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:378
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Mul, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:382
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Div, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:386
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Mod, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:390
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Concat, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:394
		{
			yyVAL.expr = &MatchExpr{genDebugInfo(), yyDollar[1].expr, yyDollar[3].expr, true}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:398
		{
			yyVAL.expr = &MatchExpr{genDebugInfo(), yyDollar[1].expr, yyDollar[3].expr, false}
		}
	case 73:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:403
		{
			yyVAL.expr = nil
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:407
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:414
		{
			yyVAL.expr = BasicLit{yyDollar[1].val}
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:418
		{
			yyVAL.expr = BasicLit{value.NewString(yyDollar[1].sym)}
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:422
		{
			yyVAL.expr = BasicLit{value.NewBool(yyDollar[1].sym == "true")}
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:426
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:430
		{
			yyVAL.expr = &UnaryExpr{genDebugInfo(), Minus, yyDollar[2].expr}
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:434
		{
			yyVAL.expr = &UnaryExpr{genDebugInfo(), Not, yyDollar[2].expr}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:438
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:442
		{
			yyVAL.expr = &Ident{Name: yyDollar[1].sym}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:446
		{
			yyVAL.expr = &CallExpr{genDebugInfo(), yyDollar[1].sym, nil}
		}
	case 84:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:450
		{
			yyVAL.expr = &CallExpr{genDebugInfo(), yyDollar[1].sym, yyDollar[3].exprlist}
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:454
		{
			yyVAL.expr = &ArrayLit{}
		}
	case 86:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:458
		{
			yyVAL.expr = &ArrayLit{yyDollar[2].exprlist}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:462
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 88:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:469
		{
			yyVAL.expr = &IndexExpr{genDebugInfo(), &Ident{Name: yyDollar[1].sym}, yyDollar[3].expr}
		}
	case 89:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:473
		{
			yyVAL.expr = &IndexExpr{genDebugInfo(), yyDollar[1].expr, yyDollar[3].expr}
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:480
		{
			yyVAL.exprlist = []Expr{yyDollar[1].expr}
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:484
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	}
	goto yystack /* stack new state and value */
}
