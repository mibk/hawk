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
	scanner *scan.Scanner
	ast     *Program

	defaultAction *BlockStmt
)

//line hawk.y:21
type yySymType struct {
	yys       int
	sym       string
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
const NUM = 57347
const BOOL = 57348
const STRING = 57349
const PRINT = 57350
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
	"NUM",
	"BOOL",
	"STRING",
	"PRINT",
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

//line hawk.y:461

// Compile compiles a Hawk program from src. It is not safe
// for concurrent use.
func Compile(src io.Reader) (*Program, error) {
	scanner = new(scan.Scanner)
	defaultAction = &BlockStmt{[]Stmt{
		&PrintStmt{Fun: "print", Args: []Expr{&FieldExpr{sc: scanner, X: BasicLit{value.Number, "0"}}}},
	}}
	ast = NewProgram(scanner)
	lexlineno = 1
	nlsemi = false
	l := &yyLex{reader: bufio.NewReader(src)}
	yyParse(l)
	analyse(ast)
	return ast, l.err
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 7,
	48, 70,
	-2, 7,
	-1, 22,
	48, 69,
	-2, 1,
	-1, 70,
	48, 69,
	-2, 39,
	-1, 115,
	48, 70,
	-2, 21,
}

const yyNprod = 90
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 529

var yyAct = [...]int{

	54, 7, 59, 58, 111, 62, 93, 53, 121, 23,
	6, 42, 43, 57, 20, 13, 15, 14, 154, 144,
	49, 97, 69, 7, 123, 60, 71, 72, 73, 74,
	75, 76, 77, 78, 79, 80, 81, 82, 83, 84,
	85, 86, 87, 24, 50, 134, 141, 142, 94, 16,
	17, 51, 92, 143, 88, 19, 158, 136, 91, 24,
	96, 22, 21, 127, 12, 18, 160, 147, 108, 125,
	110, 115, 152, 3, 114, 150, 109, 118, 107, 44,
	95, 112, 40, 41, 39, 34, 35, 36, 37, 38,
	34, 35, 36, 37, 38, 122, 55, 60, 119, 126,
	128, 129, 130, 131, 132, 133, 36, 37, 38, 1,
	124, 56, 64, 135, 63, 137, 25, 140, 26, 27,
	28, 29, 30, 31, 32, 33, 40, 41, 39, 34,
	35, 36, 37, 38, 155, 122, 146, 108, 2, 149,
	117, 5, 4, 145, 151, 153, 0, 148, 0, 0,
	157, 0, 159, 156, 0, 161, 0, 0, 0, 60,
	0, 163, 114, 162, 0, 164, 165, 61, 13, 15,
	14, 68, 0, 0, 69, 0, 70, 0, 65, 66,
	39, 34, 35, 36, 37, 38, 0, 0, 67, 61,
	13, 15, 14, 68, 0, 0, 69, 0, 70, 0,
	65, 66, 16, 17, 0, 0, 0, 0, 19, 11,
	67, 24, 0, 0, 0, 21, 0, 12, 18, 0,
	0, 0, 45, 0, 16, 17, 46, 47, 48, 0,
	19, 0, 0, 0, 0, 0, 0, 21, 0, 12,
	18, 25, 0, 26, 27, 28, 29, 30, 31, 32,
	33, 40, 41, 39, 34, 35, 36, 37, 38, 0,
	113, 13, 15, 14, 68, 0, 0, 69, 120, 70,
	0, 65, 66, 0, 0, 0, 0, 20, 13, 15,
	14, 67, 8, 9, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 16, 17, 10, 0, 0,
	0, 19, 0, 0, 0, 0, 0, 0, 21, 0,
	12, 18, 16, 17, 20, 13, 15, 14, 19, 0,
	0, 0, 0, 0, 0, 21, 0, 12, 18, 25,
	0, 26, 27, 28, 29, 30, 31, 32, 33, 40,
	41, 39, 34, 35, 36, 37, 38, 0, 0, 16,
	17, 24, 0, 0, 0, 19, 90, 20, 13, 15,
	14, 0, 21, 0, 12, 18, 0, 0, 20, 13,
	15, 14, 0, 0, 25, 0, 26, 27, 28, 29,
	30, 31, 32, 33, 40, 41, 39, 34, 35, 36,
	37, 38, 16, 17, 89, 0, 0, 0, 19, 0,
	0, 0, 0, 16, 17, 21, 52, 12, 18, 19,
	20, 13, 15, 14, 0, 0, 21, 138, 12, 18,
	105, 106, 100, 101, 102, 103, 104, 0, 105, 106,
	100, 101, 102, 103, 104, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 16, 17, 0, 50, 0,
	139, 19, 0, 0, 98, 99, 50, 0, 21, 0,
	0, 18, 98, 99, 25, 116, 26, 27, 28, 29,
	30, 31, 32, 33, 40, 41, 39, 34, 35, 36,
	37, 38, 25, 0, 26, 27, 28, 29, 30, 31,
	32, 33, 40, 41, 39, 34, 35, 36, 37, 38,
	27, 28, 29, 30, 31, 32, 33, 40, 41, 39,
	34, 35, 36, 37, 38, 28, 29, 30, 31, 32,
	33, 40, 41, 39, 34, 35, 36, 37, 38,
}
var yyPact = [...]int{

	273, -1000, 17, -1000, -1000, -1000, -5, 456, -5, -5,
	75, -1000, 406, -1000, -1000, -1000, 406, 406, 406, 364,
	-1, 353, 273, -1000, 163, 364, 364, 364, 364, 364,
	364, 364, 364, 364, 364, 364, 364, 364, 364, 364,
	364, 364, -1000, -1000, 9, -1000, -1000, -1000, -1000, 348,
	310, 364, -1000, 1, 456, -1000, 16, -29, -1000, -1000,
	456, 411, -1000, -1000, -1000, -1000, -1000, 364, 364, 364,
	256, 438, 471, 485, 46, 46, 46, 46, 46, 46,
	65, 65, -1000, -1000, -1000, 51, 142, 142, 73, -1000,
	-1000, 1, 215, -45, 364, -25, 163, 62, 364, 10,
	364, 364, 364, 364, 364, -1000, -1000, -1000, 456, -2,
	303, 13, -5, 403, -1000, 456, 364, 0, -1000, 7,
	-1000, -1000, 456, -1000, -29, -1000, 456, -32, 90, 456,
	456, 456, 456, 456, 364, 55, 364, -1000, 364, 71,
	456, -5, 68, -1000, 364, -33, -1000, 11, 12, 303,
	52, -1000, -1000, 456, 364, -1000, -1000, -1000, 185, -1000,
	364, 456, -5, 303, -1000, -1000,
}
var yyPgo = [...]int{

	0, 73, 142, 141, 140, 138, 0, 10, 209, 7,
	13, 3, 4, 5, 136, 134, 114, 112, 2, 111,
	109, 80, 6,
}
var yyR1 = [...]int{

	0, 20, 5, 5, 1, 1, 2, 2, 2, 2,
	3, 4, 4, 4, 18, 19, 19, 19, 10, 10,
	10, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 12,
	12, 13, 14, 14, 15, 15, 16, 16, 17, 17,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 7,
	7, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 9, 9, 21, 21, 22, 22,
}
var yyR2 = [...]int{

	0, 2, 1, 3, 1, 1, 2, 1, 2, 2,
	6, 0, 1, 3, 4, 0, 1, 3, 1, 1,
	3, 1, 3, 5, 6, 3, 3, 3, 3, 3,
	2, 2, 1, 1, 1, 1, 1, 2, 2, 0,
	1, 4, 0, 2, 1, 1, 7, 3, 5, 7,
	1, 5, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 0,
	1, 1, 1, 1, 2, 2, 2, 3, 1, 3,
	5, 2, 4, 4, 1, 3, 0, 1, 0, 1,
}
var yyChk = [...]int{

	-1000, -20, -5, -1, -2, -3, -7, -6, 9, 10,
	24, -8, 54, 5, 7, 6, 39, 40, 55, 45,
	4, 52, 44, -18, 48, 26, 28, 29, 30, 31,
	32, 33, 34, 35, 39, 40, 41, 42, 43, 38,
	36, 37, -18, -18, 4, -8, -8, -8, -8, -6,
	45, 52, 53, -9, -6, -1, -19, -10, -11, -18,
	-6, 4, -13, -16, -17, 15, 16, 25, 8, 11,
	13, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, 45, 46,
	46, -9, -6, -22, 47, -21, 44, 50, 51, 52,
	19, 20, 21, 22, 23, 17, 18, -7, -6, -9,
	-6, -12, -7, 4, -11, -6, 27, -4, 4, -22,
	53, 53, -6, 49, -10, 7, -6, 53, -6, -6,
	-6, -6, -6, -6, 47, -18, 44, -18, 14, 47,
	-6, 46, 47, 46, 51, 53, -14, 12, -7, -6,
	4, -18, 4, -6, 51, -15, -13, -18, 44, -18,
	14, -6, -12, -6, -18, -18,
}
var yyDef = [...]int{

	69, -2, 0, 2, 4, 5, 0, -2, 0, 0,
	0, 50, 0, 71, 72, 73, 0, 0, 0, 0,
	78, 0, -2, 6, 15, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 8, 9, 0, 52, 74, 75, 76, 0,
	0, 0, 81, 88, 84, 3, 86, 16, 18, 19,
	21, 78, 32, 33, 34, 35, 36, 69, 0, 0,
	-2, 0, 53, 54, 55, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 11, 77,
	79, 88, 0, 0, 89, 0, 87, 0, 0, 0,
	0, 0, 0, 0, 0, 30, 31, 37, 70, 38,
	0, 0, 0, 78, 40, -2, 0, 0, 12, 0,
	83, 82, 85, 14, 17, 20, 22, 0, 0, 25,
	26, 27, 28, 29, 0, 42, 69, 47, 0, 0,
	51, 0, 0, 80, 0, 83, 41, 0, 0, 0,
	0, 10, 13, 23, 0, 43, 44, 45, 39, 48,
	0, 24, 0, 0, 46, 49,
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
		//line hawk.y:63
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
		//line hawk.y:82
		{
			yyVAL.decllist = []Decl{yyDollar[1].decl}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:86
		{
			yyVAL.decllist = append(yyDollar[1].decllist, yyDollar[3].decl)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:92
		{
			yyVAL.decl = yyDollar[1].decl
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:96
		{
			yyVAL.decl = yyDollar[1].decl
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:102
		{
			yyVAL.decl = &PatternAction{yyDollar[1].expr, yyDollar[2].blockstmt}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:106
		{
			yyVAL.decl = &PatternAction{yyDollar[1].expr, defaultAction}
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:110
		{
			yyVAL.decl = &BeginAction{yyDollar[2].blockstmt}
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:114
		{
			yyVAL.decl = &EndAction{yyDollar[2].blockstmt}
		}
	case 10:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line hawk.y:120
		{
			yyVAL.decl = &FuncDecl{&FuncScope{}, yyDollar[2].sym, yyDollar[4].symlist, yyDollar[6].blockstmt}
		}
	case 11:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:125
		{
			yyVAL.symlist = nil
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:129
		{
			yyVAL.symlist = []string{yyDollar[1].sym}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:133
		{
			yyVAL.symlist = append(yyDollar[1].symlist, yyDollar[3].sym)
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:139
		{
			yyVAL.blockstmt = &BlockStmt{yyDollar[2].stmtlist}
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:144
		{
			yyVAL.stmtlist = nil
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:148
		{
			yyVAL.stmtlist = []Stmt{yyDollar[1].stmt}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:152
		{
			yyVAL.stmtlist = append(yyDollar[1].stmtlist, yyDollar[3].stmt)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:158
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:162
		{
			yyVAL.stmt = yyDollar[1].blockstmt
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:166
		{
			yyVAL.stmt = &PipeStmt{genDebugInfo(), yyDollar[1].stmt, yyDollar[3].sym}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:172
		{
			yyVAL.stmt = &ExprStmt{yyDollar[1].expr}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:176
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, &Ident{Name: yyDollar[1].sym}, yyDollar[3].expr}
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:180
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, &IndexExpr{genDebugInfo(), &Ident{Name: yyDollar[1].sym}, nil}, yyDollar[5].expr}
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line hawk.y:184
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, &IndexExpr{genDebugInfo(), &Ident{Name: yyDollar[1].sym}, yyDollar[3].expr}, yyDollar[6].expr}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:188
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, &Ident{Name: yyDollar[1].sym}, &BinaryExpr{genDebugInfo(), Add, &Ident{Name: yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:192
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, &Ident{Name: yyDollar[1].sym}, &BinaryExpr{genDebugInfo(), Sub, &Ident{Name: yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:196
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, &Ident{Name: yyDollar[1].sym}, &BinaryExpr{genDebugInfo(), Mul, &Ident{Name: yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:200
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, &Ident{Name: yyDollar[1].sym}, &BinaryExpr{genDebugInfo(), Div, &Ident{Name: yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:204
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, &Ident{Name: yyDollar[1].sym}, &BinaryExpr{genDebugInfo(), Mod, &Ident{Name: yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:210
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, &Ident{Name: yyDollar[1].sym}, &BinaryExpr{genDebugInfo(), Add, &Ident{Name: yyDollar[1].sym}, BasicLit{value.Number, "1"}}}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:214
		{
			yyVAL.stmt = &AssignStmt{genDebugInfo(), nil, &Ident{Name: yyDollar[1].sym}, &BinaryExpr{genDebugInfo(), Sub, &Ident{Name: yyDollar[1].sym}, BasicLit{value.Number, "1"}}}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:218
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:222
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:226
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:230
		{
			yyVAL.stmt = &StatusStmt{StatusBreak}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:234
		{
			yyVAL.stmt = &StatusStmt{StatusContinue}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:238
		{
			yyVAL.stmt = &ReturnStmt{ast, yyDollar[2].expr}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:242
		{
			yyVAL.stmt = &PrintStmt{genDebugInfo(), yyDollar[1].sym, yyDollar[2].exprlist}
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:247
		{
			yyVAL.stmt = nil
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:251
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:257
		{
			yyVAL.stmt = &IfStmt{genDebugInfo(), yyDollar[2].expr, yyDollar[3].blockstmt, yyDollar[4].stmt}
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:262
		{
			yyVAL.stmt = nil
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:266
		{
			yyVAL.stmt = yyDollar[2].stmt
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:272
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:276
		{
			yyVAL.stmt = yyDollar[1].blockstmt
		}
	case 46:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line hawk.y:282
		{
			yyVAL.stmt = &ForStmt{genDebugInfo(), yyDollar[2].stmt, yyDollar[4].expr, yyDollar[6].stmt, yyDollar[7].blockstmt}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:286
		{
			yyVAL.stmt = &ForStmt{genDebugInfo(), nil, yyDollar[2].expr, nil, yyDollar[3].blockstmt}
		}
	case 48:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:292
		{
			yyVAL.stmt = &ForeachStmt{genDebugInfo(), &Ident{Name: yyDollar[2].sym}, nil, yyDollar[4].expr, yyDollar[5].blockstmt}
		}
	case 49:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line hawk.y:296
		{
			yyVAL.stmt = &ForeachStmt{genDebugInfo(), &Ident{Name: yyDollar[2].sym}, &Ident{Name: yyDollar[4].sym}, yyDollar[6].expr, yyDollar[7].blockstmt}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:303
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:307
		{
			yyVAL.expr = &TernaryExpr{genDebugInfo(), yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr}
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:311
		{
			yyVAL.expr = &FieldExpr{genDebugInfo(), scanner, yyDollar[2].expr}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:315
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), OrOr, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:319
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), AndAnd, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:323
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Eq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:327
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), NotEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:331
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), LtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:335
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), GtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:339
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Lt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:343
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Gt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:347
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Add, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:351
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Sub, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:355
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Mul, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:359
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Div, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:363
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Mod, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:367
		{
			yyVAL.expr = &BinaryExpr{genDebugInfo(), Concat, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:371
		{
			yyVAL.expr = &MatchExpr{genDebugInfo(), yyDollar[1].expr, yyDollar[3].expr, true}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:375
		{
			yyVAL.expr = &MatchExpr{genDebugInfo(), yyDollar[1].expr, yyDollar[3].expr, false}
		}
	case 69:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:380
		{
			yyVAL.expr = nil
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:384
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:391
		{
			yyVAL.expr = BasicLit{value.Number, yyDollar[1].sym}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:395
		{
			yyVAL.expr = BasicLit{value.String, yyDollar[1].sym}
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:399
		{
			yyVAL.expr = BasicLit{value.Bool, yyDollar[1].sym}
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:403
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:407
		{
			yyVAL.expr = &UnaryExpr{genDebugInfo(), Minus, yyDollar[2].expr}
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:411
		{
			yyVAL.expr = &UnaryExpr{genDebugInfo(), Not, yyDollar[2].expr}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:415
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:419
		{
			yyVAL.expr = &Ident{Name: yyDollar[1].sym}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:423
		{
			yyVAL.expr = &CallExpr{genDebugInfo(), yyDollar[1].sym, nil}
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:427
		{
			yyVAL.expr = &CallExpr{genDebugInfo(), yyDollar[1].sym, yyDollar[3].exprlist}
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:431
		{
			yyVAL.expr = &ArrayLit{}
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:435
		{
			yyVAL.expr = &ArrayLit{yyDollar[2].exprlist}
		}
	case 83:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:440
		{
			yyVAL.expr = &IndexExpr{genDebugInfo(), &Ident{Name: yyDollar[1].sym}, yyDollar[3].expr}
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:446
		{
			yyVAL.exprlist = []Expr{yyDollar[1].expr}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:450
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	}
	goto yystack /* stack new state and value */
}
