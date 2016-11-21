//line hawk.y:2
package hawkc

import __yyfmt__ "fmt"

//line hawk.y:2
import (
	"bufio"
	"fmt"
	"io"

	"github.com/mibk/hawk/scan"
	"github.com/mibk/hawk/value"
)

var (
	ast *Program

	defaultAction = &BlockStmt{[]Stmt{&PrintStmt{Fun: "print"}}}
)

//line hawk.y:20
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
const FUNC = 57366
const RETURN = 57367
const OROR = 57368
const ANDAND = 57369
const EQ = 57370
const NE = 57371
const LE = 57372
const GE = 57373
const NOTMATCH = 57374

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

//line hawk.y:489

// Compile compiles a Hawk program from src. It is not safe
// for concurrent use.
func Compile(src io.Reader) (*Program, error) {
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
	48, 73,
	-2, 7,
	-1, 23,
	48, 72,
	-2, 1,
	-1, 64,
	17, 40,
	18, 40,
	19, 40,
	20, 40,
	21, 40,
	22, 40,
	23, 40,
	51, 40,
	-2, 81,
	-1, 65,
	17, 41,
	18, 41,
	19, 41,
	20, 41,
	21, 41,
	22, 41,
	23, 41,
	51, 41,
	-2, 86,
	-1, 74,
	48, 72,
	-2, 42,
	-1, 119,
	17, 40,
	18, 40,
	19, 40,
	20, 40,
	21, 40,
	22, 40,
	23, 40,
	51, 40,
	-2, 81,
	-1, 121,
	48, 73,
	-2, 21,
}

const yyNprod = 95
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 554

var yyAct = [...]int{

	55, 7, 61, 60, 117, 66, 54, 97, 59, 24,
	6, 43, 44, 20, 15, 14, 127, 13, 22, 112,
	50, 56, 152, 151, 7, 102, 62, 75, 76, 77,
	78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
	88, 89, 90, 91, 65, 51, 130, 25, 16, 17,
	145, 141, 111, 96, 19, 73, 150, 99, 95, 51,
	98, 21, 140, 12, 18, 92, 52, 148, 149, 37,
	38, 39, 114, 165, 116, 121, 143, 101, 120, 115,
	167, 51, 113, 146, 23, 118, 154, 3, 111, 132,
	159, 157, 25, 65, 35, 36, 37, 38, 39, 128,
	100, 124, 62, 125, 133, 134, 135, 136, 137, 138,
	131, 57, 96, 99, 45, 1, 58, 68, 67, 142,
	65, 144, 26, 147, 27, 28, 29, 30, 31, 32,
	33, 34, 41, 42, 40, 35, 36, 37, 38, 39,
	162, 153, 128, 63, 114, 11, 156, 2, 123, 129,
	5, 158, 160, 161, 155, 4, 0, 164, 46, 166,
	163, 0, 47, 48, 49, 0, 62, 0, 169, 120,
	168, 170, 171, 0, 64, 15, 14, 72, 13, 0,
	0, 73, 0, 74, 65, 69, 70, 41, 42, 40,
	35, 36, 37, 38, 39, 71, 64, 15, 14, 72,
	13, 0, 0, 73, 0, 74, 0, 69, 70, 16,
	17, 0, 0, 0, 0, 19, 0, 71, 25, 0,
	0, 0, 21, 0, 12, 18, 0, 0, 0, 0,
	0, 16, 17, 0, 0, 0, 0, 19, 40, 35,
	36, 37, 38, 39, 21, 0, 12, 18, 26, 0,
	27, 28, 29, 30, 31, 32, 33, 34, 41, 42,
	40, 35, 36, 37, 38, 39, 0, 119, 15, 14,
	72, 13, 0, 0, 73, 126, 74, 0, 69, 70,
	0, 0, 0, 0, 0, 20, 15, 14, 71, 13,
	8, 9, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 16, 17, 0, 10, 0, 0, 19, 0,
	0, 0, 0, 0, 0, 21, 0, 12, 18, 0,
	16, 17, 0, 0, 0, 0, 19, 20, 15, 14,
	0, 13, 0, 21, 0, 12, 18, 26, 0, 27,
	28, 29, 30, 31, 32, 33, 34, 41, 42, 40,
	35, 36, 37, 38, 39, 0, 20, 15, 14, 25,
	13, 0, 16, 17, 0, 0, 0, 0, 19, 0,
	0, 0, 0, 0, 0, 21, 139, 12, 18, 20,
	15, 14, 0, 13, 109, 110, 104, 105, 106, 107,
	108, 16, 17, 0, 0, 0, 0, 19, 94, 20,
	15, 14, 0, 13, 21, 0, 12, 18, 0, 0,
	0, 0, 0, 0, 16, 17, 0, 0, 103, 0,
	19, 0, 0, 0, 0, 0, 0, 21, 53, 12,
	18, 0, 0, 0, 16, 17, 0, 20, 15, 14,
	19, 13, 0, 0, 0, 0, 0, 21, 0, 12,
	18, 26, 0, 27, 28, 29, 30, 31, 32, 33,
	34, 41, 42, 40, 35, 36, 37, 38, 39, 0,
	0, 93, 16, 17, 0, 0, 0, 0, 19, 0,
	0, 0, 0, 0, 0, 21, 0, 0, 18, 26,
	122, 27, 28, 29, 30, 31, 32, 33, 34, 41,
	42, 40, 35, 36, 37, 38, 39, 26, 0, 27,
	28, 29, 30, 31, 32, 33, 34, 41, 42, 40,
	35, 36, 37, 38, 39, 28, 29, 30, 31, 32,
	33, 34, 41, 42, 40, 35, 36, 37, 38, 39,
	29, 30, 31, 32, 33, 34, 41, 42, 40, 35,
	36, 37, 38, 39,
}
var yyPact = [...]int{

	281, -1000, 40, -1000, -1000, -1000, -1, 481, -1, -1,
	110, -1000, 433, -1000, -1000, -1000, 433, 433, 433, 395,
	14, 375, -31, 281, -1000, 170, 395, 395, 395, 395,
	395, 395, 395, 395, 395, 395, 395, 395, 395, 395,
	395, 395, 395, -1000, -1000, 20, -1000, -1000, -1000, -1000,
	425, 352, 395, -1000, 13, 481, 395, -1000, 33, -25,
	-1000, -1000, 481, 367, 0, -33, -1000, -1000, -1000, -1000,
	-1000, 395, 395, 395, 263, 463, 496, 510, 151, 151,
	151, 151, 151, 151, 28, 28, -1000, -1000, -1000, 55,
	200, 200, 97, -1000, -1000, 13, 222, -37, 395, 96,
	-3, 170, 83, 395, 395, 395, 395, 395, 395, -1000,
	-1000, 323, 9, -1000, 481, 4, 311, 32, -1, 36,
	-1000, 481, 395, 21, -1000, 10, -1000, -1000, 481, -1000,
	-1000, -25, -1000, 481, 481, 481, 481, 481, 481, -28,
	-29, 395, 74, 395, -1000, 395, 87, 481, -1, 86,
	-1000, 395, 395, -1000, 44, 29, 311, 66, -1000, -1000,
	481, 481, -1000, -1000, -1000, 192, -1000, 395, -1, 311,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 87, 155, 150, 148, 147, 0, 10, 145, 18,
	143, 6, 8, 3, 4, 5, 141, 140, 118, 117,
	2, 116, 115, 100, 7,
}
var yyR1 = [...]int{

	0, 22, 5, 5, 1, 1, 2, 2, 2, 2,
	3, 4, 4, 4, 20, 21, 21, 21, 12, 12,
	12, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	10, 10, 14, 14, 15, 16, 16, 17, 17, 18,
	18, 19, 19, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 7, 7, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 9, 9, 11,
	11, 23, 23, 24, 24,
}
var yyR2 = [...]int{

	0, 2, 1, 3, 1, 1, 2, 1, 2, 2,
	6, 0, 1, 3, 4, 0, 1, 3, 1, 1,
	3, 1, 3, 5, 5, 3, 3, 3, 3, 3,
	2, 2, 1, 1, 1, 1, 1, 2, 2, 1,
	1, 1, 0, 1, 4, 0, 2, 1, 1, 7,
	3, 5, 7, 1, 5, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 0, 1, 1, 1, 1, 2, 2, 2,
	3, 1, 3, 5, 2, 4, 1, 4, 4, 1,
	3, 0, 1, 0, 1,
}
var yyChk = [...]int{

	-1000, -22, -5, -1, -2, -3, -7, -6, 9, 10,
	24, -8, 54, 8, 6, 5, 39, 40, 55, 45,
	4, 52, -9, 44, -20, 48, 26, 28, 29, 30,
	31, 32, 33, 34, 35, 39, 40, 41, 42, 43,
	38, 36, 37, -20, -20, 4, -8, -8, -8, -8,
	-6, 45, 52, 53, -11, -6, 52, -1, -21, -12,
	-13, -20, -6, -10, 4, -9, -15, -18, -19, 15,
	16, 25, 7, 11, 13, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, 45, 46, 46, -11, -6, -24, 47, -6,
	-23, 44, 50, 51, 19, 20, 21, 22, 23, 17,
	18, 52, 52, -7, -6, -11, -6, -14, -7, 4,
	-13, -6, 27, -4, 4, -24, 53, 53, -6, 53,
	49, -12, 6, -6, -6, -6, -6, -6, -6, 53,
	53, 47, -20, 44, -20, 14, 47, -6, 46, 47,
	46, 51, 51, -16, 12, -7, -6, 4, -20, 4,
	-6, -6, -17, -15, -20, 44, -20, 14, -14, -6,
	-20, -20,
}
var yyDef = [...]int{

	72, -2, 0, 2, 4, 5, 0, -2, 0, 0,
	0, 53, 0, 74, 75, 76, 0, 0, 0, 0,
	81, 0, 86, -2, 6, 15, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 8, 9, 0, 55, 77, 78, 79,
	0, 0, 0, 84, 93, 89, 0, 3, 91, 16,
	18, 19, 21, 0, -2, -2, 32, 33, 34, 35,
	36, 72, 39, 0, -2, 0, 56, 57, 58, 59,
	60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
	70, 71, 11, 80, 82, 93, 0, 0, 94, 0,
	0, 92, 0, 0, 0, 0, 0, 0, 0, 30,
	31, 0, 0, 37, 73, 38, 0, 0, 0, -2,
	43, -2, 0, 0, 12, 0, 87, 85, 90, 88,
	14, 17, 20, 22, 25, 26, 27, 28, 29, 0,
	0, 0, 45, 72, 50, 0, 0, 54, 0, 0,
	83, 0, 0, 44, 0, 0, 0, 0, 10, 13,
	23, 24, 46, 47, 48, 42, 51, 0, 0, 0,
	49, 52,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 55, 3, 3, 54, 43, 3, 3,
	45, 46, 41, 39, 47, 40, 38, 42, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 27, 44,
	34, 51, 35, 26, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 52, 3, 53, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 48, 50, 49, 36,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 28, 29, 30, 31, 32, 33,
	37,
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
		//line hawk.y:64
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
		//line hawk.y:83
		{
			yyVAL.decllist = []Decl{yyDollar[1].decl}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:87
		{
			yyVAL.decllist = append(yyDollar[1].decllist, yyDollar[3].decl)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:93
		{
			yyVAL.decl = yyDollar[1].decl
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:97
		{
			yyVAL.decl = yyDollar[1].decl
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:103
		{
			yyVAL.decl = &PatternAction{yyDollar[1].expr, yyDollar[2].blockstmt}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:107
		{
			yyVAL.decl = &PatternAction{yyDollar[1].expr, defaultAction}
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:111
		{
			yyVAL.decl = &BeginAction{yyDollar[2].blockstmt}
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:115
		{
			yyVAL.decl = &EndAction{yyDollar[2].blockstmt}
		}
	case 10:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line hawk.y:121
		{
			yyVAL.decl = &FuncDecl{&FuncScope{}, yyDollar[2].sym, yyDollar[4].symlist, yyDollar[6].blockstmt}
		}
	case 11:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:126
		{
			yyVAL.symlist = nil
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:130
		{
			yyVAL.symlist = []string{yyDollar[1].sym}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:134
		{
			yyVAL.symlist = append(yyDollar[1].symlist, yyDollar[3].sym)
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:140
		{
			yyVAL.blockstmt = &BlockStmt{yyDollar[2].stmtlist}
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:145
		{
			yyVAL.stmtlist = nil
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:149
		{
			yyVAL.stmtlist = []Stmt{yyDollar[1].stmt}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:153
		{
			yyVAL.stmtlist = append(yyDollar[1].stmtlist, yyDollar[3].stmt)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:159
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:163
		{
			yyVAL.stmt = yyDollar[1].blockstmt
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:167
		{
			yyVAL.stmt = &PipeStmt{genDebugInfo(), yyDollar[1].stmt, yyDollar[3].sym}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:173
		{
			yyVAL.stmt = &ExprStmt{yyDollar[1].expr}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:177
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:184
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, &IndexExpr{genDebugInfo(), &Ident{Name: yyDollar[1].sym}, nil}, yyDollar[5].expr}
		}
	case 24:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:188
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, &IndexExpr{genDebugInfo(), yyDollar[1].expr, nil}, yyDollar[5].expr}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:193
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, yyDollar[1].expr, &BinaryExpr{genDebugInfo(), Add, yyDollar[1].expr, yyDollar[3].expr}}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:197
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, yyDollar[1].expr, &BinaryExpr{genDebugInfo(), Sub, yyDollar[1].expr, yyDollar[3].expr}}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:201
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, yyDollar[1].expr, &BinaryExpr{genDebugInfo(), Mul, yyDollar[1].expr, yyDollar[3].expr}}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:205
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, yyDollar[1].expr, &BinaryExpr{genDebugInfo(), Div, yyDollar[1].expr, yyDollar[3].expr}}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:209
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, yyDollar[1].expr, &BinaryExpr{genDebugInfo(), Mod, yyDollar[1].expr, yyDollar[3].expr}}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:213
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, yyDollar[1].expr, &BinaryExpr{genDebugInfo(), Add, yyDollar[1].expr, BasicLit{value.NewNumber(1)}}}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:217
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, yyDollar[1].expr, &BinaryExpr{genDebugInfo(), Sub, yyDollar[1].expr, BasicLit{value.NewNumber(1)}}}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:221
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:225
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:229
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:233
		{
			yyVAL.stmt = &StatusStmt{StatusBreak}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:237
		{
			yyVAL.stmt = &StatusStmt{StatusContinue}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:241
		{
			yyVAL.stmt = &ReturnStmt{X: yyDollar[2].expr}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:245
		{
			yyVAL.stmt = &PrintStmt{genDebugInfo(), nil, yyDollar[1].sym, yyDollar[2].exprlist}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:249
		{
			yyVAL.stmt = &PrintStmt{genDebugInfo(), nil, yyDollar[1].sym, nil}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:255
		{
			yyVAL.expr = &Ident{ast, yyDollar[1].sym}
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:259
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:264
		{
			yyVAL.stmt = nil
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:268
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:274
		{
			yyVAL.stmt = &IfStmt{genDebugInfo(), yyDollar[2].expr, yyDollar[3].blockstmt, yyDollar[4].stmt}
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:279
		{
			yyVAL.stmt = nil
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:283
		{
			yyVAL.stmt = yyDollar[2].stmt
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:289
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:293
		{
			yyVAL.stmt = yyDollar[1].blockstmt
		}
	case 49:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line hawk.y:299
		{
			yyVAL.stmt = &ForStmt{genDebugInfo(), yyDollar[2].stmt, yyDollar[4].expr, yyDollar[6].stmt, yyDollar[7].blockstmt}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:303
		{
			yyVAL.stmt = &ForStmt{genDebugInfo(), nil, yyDollar[2].expr, nil, yyDollar[3].blockstmt}
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:309
		{
			yyVAL.stmt = &ForeachStmt{genDebugInfo(), &Ident{Name: yyDollar[2].sym}, nil, yyDollar[4].expr, yyDollar[5].blockstmt}
		}
	case 52:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line hawk.y:313
		{
			yyVAL.stmt = &ForeachStmt{genDebugInfo(), &Ident{Name: yyDollar[2].sym}, &Ident{Name: yyDollar[4].sym}, yyDollar[6].expr, yyDollar[7].blockstmt}
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:320
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 54:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:324
		{
			yyVAL.expr = &TernaryExpr{genDebugInfo(), yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr}
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:328
		{
			yyVAL.expr = &FieldExpr{genDebugInfo(), nil, yyDollar[2].expr}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:332
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), OrOr, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:336
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), AndAnd, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:340
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Eq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:344
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), NotEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:348
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), LtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:352
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), GtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:356
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Lt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:360
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Gt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:364
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Add, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:368
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Sub, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:372
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Mul, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:376
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Div, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:380
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Mod, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:384
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Concat, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:388
		{
			yyVAL.expr = &MatchExpr{genDebugInfo(), yyDollar[1].expr, yyDollar[3].expr, true}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:392
		{
			yyVAL.expr = &MatchExpr{genDebugInfo(), yyDollar[1].expr, yyDollar[3].expr, false}
		}
	case 72:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:397
		{
			yyVAL.expr = nil
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:401
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:408
		{
			yyVAL.expr = BasicLit{yyDollar[1].val}
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:412
		{
			yyVAL.expr = BasicLit{value.NewString(yyDollar[1].sym)}
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:416
		{
			yyVAL.expr = BasicLit{value.NewBool(yyDollar[1].sym == "true")}
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:420
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:424
		{
			yyVAL.expr = &UnaryExpr{genDebugInfo(), Minus, yyDollar[2].expr}
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:428
		{
			yyVAL.expr = &UnaryExpr{genDebugInfo(), Not, yyDollar[2].expr}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:432
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:436
		{
			yyVAL.expr = &Ident{Name: yyDollar[1].sym}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:440
		{
			yyVAL.expr = &CallExpr{genDebugInfo(), yyDollar[1].sym, nil}
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:444
		{
			yyVAL.expr = &CallExpr{genDebugInfo(), yyDollar[1].sym, yyDollar[3].exprlist}
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:448
		{
			yyVAL.expr = &ArrayLit{}
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:452
		{
			yyVAL.expr = &ArrayLit{yyDollar[2].exprlist}
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:456
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 87:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:463
		{
			yyVAL.expr = &IndexExpr{genDebugInfo(), &Ident{Name: yyDollar[1].sym}, yyDollar[3].expr}
		}
	case 88:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:467
		{
			yyVAL.expr = &IndexExpr{genDebugInfo(), yyDollar[1].expr, yyDollar[3].expr}
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:474
		{
			yyVAL.exprlist = []Expr{yyDollar[1].expr}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:478
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	}
	goto yystack /* stack new state and value */
}
