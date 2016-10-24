//line hawk.y:2
package compiler

import __yyfmt__ "fmt"

//line hawk.y:2
import (
	"bufio"
	"fmt"
	"io"

	"github.com/mibk/hawk/scan"
)

var (
	scanner *scan.Scanner
	ast     *Program

	defaultAction *BlockStmt
)

//line hawk.y:20
type yySymType struct {
	yys       int
	num       int
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

const NUM = 57346
const IDENT = 57347
const STRING = 57348
const PRINT = 57349
const BEGIN = 57350
const END = 57351
const IF = 57352
const ELSE = 57353
const FOR = 57354
const IN = 57355
const BREAK = 57356
const CONTINUE = 57357
const INC = 57358
const DEC = 57359
const ADDEQ = 57360
const SUBEQ = 57361
const MULEQ = 57362
const DIVEQ = 57363
const MODEQ = 57364
const FUNC = 57365
const RETURN = 57366
const OROR = 57367
const ANDAND = 57368
const EQ = 57369
const NE = 57370
const LE = 57371
const GE = 57372

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUM",
	"IDENT",
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

//line hawk.y:444

// Compile compiles a Hawk program from src. It is not safe
// for concurrent use.
func Compile(src io.Reader) (*Program, error) {
	scanner = new(scan.Scanner)
	defaultAction = &BlockStmt{[]Stmt{
		&PrintStmt{"print", []Expr{&FieldExpr{scanner, Lit(0)}}},
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
	44, 67,
	-2, 7,
	-1, 21,
	44, 66,
	-2, 1,
	-1, 66,
	44, 66,
	-2, 39,
	-1, 108,
	44, 67,
	-2, 21,
}

const yyNprod = 86
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 508

var yyAct = [...]int{

	50, 7, 55, 54, 104, 58, 86, 49, 114, 22,
	6, 38, 39, 53, 13, 19, 14, 147, 137, 45,
	90, 65, 7, 116, 56, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 23,
	46, 127, 134, 135, 87, 15, 16, 47, 85, 136,
	81, 18, 151, 129, 84, 23, 89, 21, 20, 120,
	12, 17, 153, 140, 101, 3, 103, 108, 118, 145,
	107, 143, 102, 111, 100, 40, 88, 105, 33, 34,
	35, 36, 37, 35, 36, 37, 1, 51, 115, 52,
	56, 112, 119, 121, 122, 123, 124, 125, 126, 60,
	59, 148, 139, 117, 2, 110, 128, 5, 130, 24,
	133, 25, 26, 27, 28, 29, 30, 31, 32, 33,
	34, 35, 36, 37, 4, 0, 0, 0, 115, 0,
	101, 0, 142, 138, 0, 0, 0, 144, 146, 0,
	141, 0, 0, 150, 0, 152, 149, 0, 154, 0,
	0, 0, 56, 0, 156, 107, 155, 0, 157, 158,
	13, 57, 14, 64, 0, 11, 65, 0, 66, 0,
	61, 62, 0, 0, 0, 0, 0, 0, 41, 0,
	63, 42, 43, 44, 0, 0, 0, 0, 0, 0,
	0, 15, 16, 0, 13, 19, 14, 18, 0, 0,
	23, 0, 0, 0, 20, 0, 12, 17, 13, 57,
	14, 64, 0, 0, 65, 0, 66, 0, 61, 62,
	0, 0, 0, 0, 0, 15, 16, 0, 63, 0,
	0, 18, 0, 0, 0, 0, 0, 0, 20, 15,
	16, 17, 0, 0, 0, 18, 0, 0, 13, 19,
	14, 0, 20, 0, 12, 17, 13, 106, 14, 64,
	0, 0, 65, 0, 66, 0, 61, 62, 0, 0,
	0, 0, 0, 0, 0, 0, 63, 0, 0, 15,
	16, 0, 0, 0, 0, 18, 83, 15, 16, 0,
	0, 0, 20, 18, 12, 17, 0, 0, 0, 0,
	20, 0, 12, 17, 24, 0, 25, 26, 27, 28,
	29, 30, 31, 32, 33, 34, 35, 36, 37, 13,
	19, 14, 0, 8, 9, 0, 0, 0, 113, 0,
	13, 19, 14, 0, 0, 0, 0, 0, 10, 0,
	0, 13, 19, 14, 0, 0, 0, 0, 0, 0,
	15, 16, 0, 0, 0, 0, 18, 0, 0, 0,
	0, 15, 16, 20, 0, 12, 17, 18, 0, 0,
	0, 0, 15, 16, 20, 48, 12, 17, 18, 0,
	0, 0, 0, 0, 0, 20, 131, 12, 17, 98,
	99, 93, 94, 95, 96, 97, 24, 0, 25, 26,
	27, 28, 29, 30, 31, 32, 33, 34, 35, 36,
	37, 0, 0, 0, 46, 23, 132, 0, 0, 0,
	91, 92, 98, 99, 93, 94, 95, 96, 97, 24,
	0, 25, 26, 27, 28, 29, 30, 31, 32, 33,
	34, 35, 36, 37, 0, 0, 82, 46, 0, 0,
	0, 0, 0, 91, 92, 24, 109, 25, 26, 27,
	28, 29, 30, 31, 32, 33, 34, 35, 36, 37,
	24, 0, 25, 26, 27, 28, 29, 30, 31, 32,
	33, 34, 35, 36, 37, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37,
}
var yyPact = [...]int{

	315, -1000, 17, -1000, -1000, -1000, -5, 445, -5, -5,
	70, -1000, 190, -1000, -1000, 190, 190, 190, 337, -1,
	326, 315, -1000, 156, 337, 337, 337, 337, 337, 337,
	337, 337, 337, 337, 337, 337, 337, 337, -1000, -1000,
	9, -1000, -1000, -1000, -1000, 404, 244, 337, -1000, 1,
	445, -1000, 16, -26, -1000, -1000, 445, 406, -1000, -1000,
	-1000, -1000, -1000, 337, 337, 337, 252, 430, 457, 468,
	43, 43, 43, 43, 43, 43, 46, 46, -1000, -1000,
	-1000, 68, -1000, -1000, 1, 279, -41, 337, -22, 156,
	62, 337, 10, 337, 337, 337, 337, 337, -1000, -1000,
	-1000, 445, -2, 371, 13, -5, 373, -1000, 445, 337,
	0, -1000, 7, -1000, -1000, 445, -1000, -26, -1000, 445,
	-29, 84, 445, 445, 445, 445, 445, 337, 52, 337,
	-1000, 337, 66, 445, -5, 64, -1000, 337, -30, -1000,
	11, 12, 371, 49, -1000, -1000, 445, 337, -1000, -1000,
	-1000, 204, -1000, 337, 445, -5, 371, -1000, -1000,
}
var yyPgo = [...]int{

	0, 65, 124, 107, 105, 104, 0, 10, 165, 7,
	13, 3, 4, 5, 102, 101, 100, 99, 2, 89,
	86, 76, 6,
}
var yyR1 = [...]int{

	0, 20, 5, 5, 1, 1, 2, 2, 2, 2,
	3, 4, 4, 4, 18, 19, 19, 19, 10, 10,
	10, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 12,
	12, 13, 14, 14, 15, 15, 16, 16, 17, 17,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 7, 7, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	9, 9, 21, 21, 22, 22,
}
var yyR2 = [...]int{

	0, 2, 1, 3, 1, 1, 2, 1, 2, 2,
	6, 0, 1, 3, 4, 0, 1, 3, 1, 1,
	3, 1, 3, 5, 6, 3, 3, 3, 3, 3,
	2, 2, 1, 1, 1, 1, 1, 2, 2, 0,
	1, 4, 0, 2, 1, 1, 7, 3, 5, 7,
	1, 5, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 0, 1, 1, 1,
	2, 2, 2, 3, 1, 3, 5, 2, 4, 4,
	1, 3, 0, 1, 0, 1,
}
var yyChk = [...]int{

	-1000, -20, -5, -1, -2, -3, -7, -6, 8, 9,
	23, -8, 50, 4, 6, 35, 36, 51, 41, 5,
	48, 40, -18, 44, 25, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, -18, -18,
	5, -8, -8, -8, -8, -6, 41, 48, 49, -9,
	-6, -1, -19, -10, -11, -18, -6, 5, -13, -16,
	-17, 14, 15, 24, 7, 10, 12, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, 41, 42, 42, -9, -6, -22, 43, -21, 40,
	46, 47, 48, 18, 19, 20, 21, 22, 16, 17,
	-7, -6, -9, -6, -12, -7, 5, -11, -6, 26,
	-4, 5, -22, 49, 49, -6, 45, -10, 6, -6,
	49, -6, -6, -6, -6, -6, -6, 43, -18, 40,
	-18, 13, 43, -6, 42, 43, 42, 47, 49, -14,
	11, -7, -6, 5, -18, 5, -6, 47, -15, -13,
	-18, 40, -18, 13, -6, -12, -6, -18, -18,
}
var yyDef = [...]int{

	66, -2, 0, 2, 4, 5, 0, -2, 0, 0,
	0, 50, 0, 68, 69, 0, 0, 0, 0, 74,
	0, -2, 6, 15, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 8, 9,
	0, 52, 70, 71, 72, 0, 0, 0, 77, 84,
	80, 3, 82, 16, 18, 19, 21, 74, 32, 33,
	34, 35, 36, 66, 0, 0, -2, 0, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 11, 73, 75, 84, 0, 0, 85, 0, 83,
	0, 0, 0, 0, 0, 0, 0, 0, 30, 31,
	37, 67, 38, 0, 0, 0, 74, 40, -2, 0,
	0, 12, 0, 79, 78, 81, 14, 17, 20, 22,
	0, 0, 25, 26, 27, 28, 29, 0, 42, 66,
	47, 0, 0, 51, 0, 0, 76, 0, 79, 41,
	0, 0, 0, 0, 10, 13, 23, 0, 43, 44,
	45, 39, 48, 0, 24, 0, 0, 46, 49,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 51, 3, 3, 50, 39, 3, 3,
	41, 42, 37, 35, 43, 36, 3, 38, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 26, 40,
	33, 47, 34, 25, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 48, 3, 49, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 44, 46, 45,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 27, 28, 29, 30, 31, 32,
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
		//line hawk.y:62
		{
			for _, d := range yyDollar[1].decllist {
				switch d := d.(type) {
				case *BeginAction:
					ast.begin = append(ast.begin, d)
				case *EndAction:
					ast.end = append(ast.end, d)
				case *PatternAction:
					ast.pActions = append(ast.pActions, d)
				case *FuncDecl:
					ast.funcs[d.name] = d
				default:
					panic(fmt.Sprintf("unexpected type: %T", d))
				}
			}
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:81
		{
			yyVAL.decllist = []Decl{yyDollar[1].decl}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:85
		{
			yyVAL.decllist = append(yyDollar[1].decllist, yyDollar[3].decl)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:91
		{
			yyVAL.decl = yyDollar[1].decl
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:95
		{
			yyVAL.decl = yyDollar[1].decl
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:101
		{
			yyVAL.decl = &PatternAction{yyDollar[1].expr, yyDollar[2].blockstmt}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:105
		{
			yyVAL.decl = &PatternAction{yyDollar[1].expr, defaultAction}
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:109
		{
			yyVAL.decl = &BeginAction{yyDollar[2].blockstmt}
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:113
		{
			yyVAL.decl = &EndAction{yyDollar[2].blockstmt}
		}
	case 10:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line hawk.y:119
		{
			yyVAL.decl = &FuncDecl{&FuncScope{}, yyDollar[2].sym, yyDollar[4].symlist, yyDollar[6].blockstmt}
		}
	case 11:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:124
		{
			yyVAL.symlist = nil
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:128
		{
			yyVAL.symlist = []string{yyDollar[1].sym}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:132
		{
			yyVAL.symlist = append(yyDollar[1].symlist, yyDollar[3].sym)
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:138
		{
			yyVAL.blockstmt = &BlockStmt{yyDollar[2].stmtlist}
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:143
		{
			yyVAL.stmtlist = nil
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:147
		{
			yyVAL.stmtlist = []Stmt{yyDollar[1].stmt}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:151
		{
			yyVAL.stmtlist = append(yyDollar[1].stmtlist, yyDollar[3].stmt)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:157
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:161
		{
			yyVAL.stmt = yyDollar[1].blockstmt
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:165
		{
			yyVAL.stmt = &PipeStmt{yyDollar[1].stmt, yyDollar[3].sym}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:171
		{
			yyVAL.stmt = &ExprStmt{yyDollar[1].expr}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:175
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:179
		{
			yyVAL.stmt = &AssignStmt{ast, &IndexExpr{&Ident{ast, yyDollar[1].sym}, nil}, yyDollar[5].expr}
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line hawk.y:183
		{
			yyVAL.stmt = &AssignStmt{ast, &IndexExpr{&Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}, yyDollar[6].expr}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:187
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Add, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:191
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Sub, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:195
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Mul, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:199
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Div, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:203
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Mod, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:209
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Add, &Ident{ast, yyDollar[1].sym}, Lit(1)}}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:213
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Sub, &Ident{ast, yyDollar[1].sym}, Lit(1)}}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:217
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:221
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:225
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:229
		{
			yyVAL.stmt = &StatusStmt{StatusBreak}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:233
		{
			yyVAL.stmt = &StatusStmt{StatusContinue}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:237
		{
			yyVAL.stmt = &ReturnStmt{ast, yyDollar[2].expr}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:241
		{
			yyVAL.stmt = &PrintStmt{yyDollar[1].sym, yyDollar[2].exprlist}
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:246
		{
			yyVAL.stmt = nil
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:250
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:256
		{
			yyVAL.stmt = &IfStmt{yyDollar[2].expr, yyDollar[3].blockstmt, yyDollar[4].stmt}
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:261
		{
			yyVAL.stmt = nil
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:265
		{
			yyVAL.stmt = yyDollar[2].stmt
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:271
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:275
		{
			yyVAL.stmt = yyDollar[1].blockstmt
		}
	case 46:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line hawk.y:281
		{
			yyVAL.stmt = &ForStmt{yyDollar[2].stmt, yyDollar[4].expr, yyDollar[6].stmt, yyDollar[7].blockstmt}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:285
		{
			yyVAL.stmt = &ForStmt{nil, yyDollar[2].expr, nil, yyDollar[3].blockstmt}
		}
	case 48:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:291
		{
			yyVAL.stmt = &ForeachStmt{&Ident{ast, yyDollar[2].sym}, nil, yyDollar[4].expr, yyDollar[5].blockstmt}
		}
	case 49:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line hawk.y:295
		{
			yyVAL.stmt = &ForeachStmt{&Ident{ast, yyDollar[2].sym}, &Ident{ast, yyDollar[4].sym}, yyDollar[6].expr, yyDollar[7].blockstmt}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:302
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:306
		{
			yyVAL.expr = &TernaryExpr{yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr}
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:310
		{
			yyVAL.expr = &FieldExpr{scanner, yyDollar[2].expr}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:314
		{
			yyVAL.expr = &BinaryExpr{OrOr, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:318
		{
			yyVAL.expr = &BinaryExpr{AndAnd, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:322
		{
			yyVAL.expr = &BinaryExpr{Eq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:326
		{
			yyVAL.expr = &BinaryExpr{NotEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:330
		{
			yyVAL.expr = &BinaryExpr{LtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:334
		{
			yyVAL.expr = &BinaryExpr{GtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:338
		{
			yyVAL.expr = &BinaryExpr{Lt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:342
		{
			yyVAL.expr = &BinaryExpr{Gt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:346
		{
			yyVAL.expr = &BinaryExpr{Add, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:350
		{
			yyVAL.expr = &BinaryExpr{Sub, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:354
		{
			yyVAL.expr = &BinaryExpr{Mul, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:358
		{
			yyVAL.expr = &BinaryExpr{Div, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:362
		{
			yyVAL.expr = &BinaryExpr{Mod, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 66:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:367
		{
			yyVAL.expr = nil
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:371
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:378
		{
			yyVAL.expr = Lit(yyDollar[1].num)
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:382
		{
			yyVAL.expr = StringLit(yyDollar[1].sym)
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:386
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:390
		{
			yyVAL.expr = &UnaryExpr{Minus, yyDollar[2].expr}
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:394
		{
			yyVAL.expr = &UnaryExpr{Not, yyDollar[2].expr}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:398
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:402
		{
			yyVAL.expr = &Ident{ast, yyDollar[1].sym}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:406
		{
			yyVAL.expr = &CallExpr{yyDollar[1].sym, nil}
		}
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:410
		{
			yyVAL.expr = &CallExpr{yyDollar[1].sym, yyDollar[3].exprlist}
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:414
		{
			yyVAL.expr = &ArrayLit{}
		}
	case 78:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:418
		{
			yyVAL.expr = &ArrayLit{yyDollar[2].exprlist}
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:423
		{
			yyVAL.expr = &IndexExpr{&Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:429
		{
			yyVAL.exprlist = []Expr{yyDollar[1].expr}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:433
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	}
	goto yystack /* stack new state and value */
}
