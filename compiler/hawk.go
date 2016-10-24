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

//line hawk.y:449

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
	45, 68,
	-2, 7,
	-1, 21,
	45, 67,
	-2, 1,
	-1, 67,
	45, 67,
	-2, 39,
	-1, 110,
	45, 68,
	-2, 21,
}

const yyNprod = 87
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 474

var yyAct = [...]int{

	51, 7, 56, 55, 106, 59, 88, 50, 116, 22,
	6, 39, 40, 54, 13, 19, 14, 149, 139, 46,
	92, 66, 7, 118, 57, 68, 69, 70, 71, 72,
	73, 74, 75, 76, 77, 78, 79, 80, 81, 82,
	23, 47, 129, 136, 137, 89, 15, 16, 48, 87,
	138, 83, 18, 153, 131, 86, 23, 91, 21, 20,
	122, 12, 17, 155, 142, 103, 120, 105, 110, 3,
	147, 109, 145, 104, 113, 102, 41, 90, 107, 38,
	33, 34, 35, 36, 37, 33, 34, 35, 36, 37,
	117, 52, 57, 114, 121, 123, 124, 125, 126, 127,
	128, 35, 36, 37, 1, 119, 53, 61, 130, 60,
	132, 24, 135, 25, 26, 27, 28, 29, 30, 31,
	32, 38, 33, 34, 35, 36, 37, 150, 141, 2,
	117, 112, 103, 5, 144, 4, 140, 0, 0, 146,
	148, 0, 143, 0, 0, 152, 0, 154, 151, 0,
	156, 13, 19, 14, 57, 0, 158, 109, 157, 0,
	159, 160, 13, 58, 14, 65, 0, 0, 66, 0,
	67, 0, 62, 63, 0, 0, 0, 0, 0, 0,
	0, 0, 64, 15, 16, 0, 0, 0, 11, 18,
	85, 13, 19, 14, 15, 16, 20, 0, 12, 17,
	18, 42, 0, 23, 43, 44, 45, 20, 0, 12,
	17, 13, 58, 14, 65, 0, 0, 66, 0, 67,
	0, 62, 63, 15, 16, 0, 0, 0, 0, 18,
	0, 64, 0, 0, 0, 0, 20, 49, 12, 17,
	13, 19, 14, 15, 16, 0, 0, 0, 0, 18,
	0, 13, 19, 14, 0, 0, 20, 0, 12, 17,
	13, 108, 14, 65, 0, 0, 66, 0, 67, 0,
	62, 63, 15, 16, 0, 0, 0, 0, 18, 0,
	64, 0, 0, 15, 16, 20, 0, 12, 17, 18,
	0, 0, 15, 16, 0, 0, 20, 0, 18, 17,
	0, 0, 0, 0, 0, 20, 0, 12, 17, 24,
	0, 25, 26, 27, 28, 29, 30, 31, 32, 38,
	33, 34, 35, 36, 37, 13, 19, 14, 0, 8,
	9, 0, 0, 133, 115, 0, 100, 101, 95, 96,
	97, 98, 99, 0, 10, 27, 28, 29, 30, 31,
	32, 38, 33, 34, 35, 36, 37, 15, 16, 0,
	0, 0, 47, 18, 134, 0, 0, 0, 93, 94,
	20, 0, 12, 17, 24, 0, 25, 26, 27, 28,
	29, 30, 31, 32, 38, 33, 34, 35, 36, 37,
	0, 0, 0, 0, 23, 100, 101, 95, 96, 97,
	98, 99, 24, 0, 25, 26, 27, 28, 29, 30,
	31, 32, 38, 33, 34, 35, 36, 37, 0, 0,
	84, 47, 0, 0, 0, 0, 0, 93, 94, 24,
	111, 25, 26, 27, 28, 29, 30, 31, 32, 38,
	33, 34, 35, 36, 37, 24, 0, 25, 26, 27,
	28, 29, 30, 31, 32, 38, 33, 34, 35, 36,
	37, 26, 27, 28, 29, 30, 31, 32, 38, 33,
	34, 35, 36, 37,
}
var yyPact = [...]int{

	321, -1000, 17, -1000, -1000, -1000, -5, 420, -5, -5,
	71, -1000, 247, -1000, -1000, 247, 247, 247, 236, -1,
	187, 321, -1000, 158, 236, 236, 236, 236, 236, 236,
	236, 236, 236, 236, 236, 236, 236, 236, 236, -1000,
	-1000, 9, -1000, -1000, -1000, -1000, 377, 147, 236, -1000,
	1, 420, -1000, 16, -27, -1000, -1000, 420, 379, -1000,
	-1000, -1000, -1000, -1000, 236, 236, 236, 256, 404, 433,
	316, 44, 44, 44, 44, 44, 44, 63, 63, -1000,
	-1000, -1000, 49, 69, -1000, -1000, 1, 284, -42, 236,
	-23, 158, 60, 236, 10, 236, 236, 236, 236, 236,
	-1000, -1000, -1000, 420, -2, 349, 13, -5, 320, -1000,
	420, 236, 0, -1000, 7, -1000, -1000, 420, -1000, -27,
	-1000, 420, -30, 86, 420, 420, 420, 420, 420, 236,
	53, 236, -1000, 236, 67, 420, -5, 65, -1000, 236,
	-31, -1000, 11, 12, 349, 50, -1000, -1000, 420, 236,
	-1000, -1000, -1000, 207, -1000, 236, 420, -5, 349, -1000,
	-1000,
}
var yyPgo = [...]int{

	0, 69, 135, 133, 131, 129, 0, 10, 188, 7,
	13, 3, 4, 5, 128, 127, 109, 107, 2, 106,
	104, 77, 6,
}
var yyR1 = [...]int{

	0, 20, 5, 5, 1, 1, 2, 2, 2, 2,
	3, 4, 4, 4, 18, 19, 19, 19, 10, 10,
	10, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 12,
	12, 13, 14, 14, 15, 15, 16, 16, 17, 17,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 7, 7, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 9, 9, 21, 21, 22, 22,
}
var yyR2 = [...]int{

	0, 2, 1, 3, 1, 1, 2, 1, 2, 2,
	6, 0, 1, 3, 4, 0, 1, 3, 1, 1,
	3, 1, 3, 5, 6, 3, 3, 3, 3, 3,
	2, 2, 1, 1, 1, 1, 1, 2, 2, 0,
	1, 4, 0, 2, 1, 1, 7, 3, 5, 7,
	1, 5, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 0, 1, 1,
	1, 2, 2, 2, 3, 1, 3, 5, 2, 4,
	4, 1, 3, 0, 1, 0, 1,
}
var yyChk = [...]int{

	-1000, -20, -5, -1, -2, -3, -7, -6, 8, 9,
	23, -8, 51, 4, 6, 36, 37, 52, 42, 5,
	49, 41, -18, 45, 25, 27, 28, 29, 30, 31,
	32, 33, 34, 36, 37, 38, 39, 40, 35, -18,
	-18, 5, -8, -8, -8, -8, -6, 42, 49, 50,
	-9, -6, -1, -19, -10, -11, -18, -6, 5, -13,
	-16, -17, 14, 15, 24, 7, 10, 12, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, 42, 43, 43, -9, -6, -22, 44,
	-21, 41, 47, 48, 49, 18, 19, 20, 21, 22,
	16, 17, -7, -6, -9, -6, -12, -7, 5, -11,
	-6, 26, -4, 5, -22, 50, 50, -6, 46, -10,
	6, -6, 50, -6, -6, -6, -6, -6, -6, 44,
	-18, 41, -18, 13, 44, -6, 43, 44, 43, 48,
	50, -14, 11, -7, -6, 5, -18, 5, -6, 48,
	-15, -13, -18, 41, -18, 13, -6, -12, -6, -18,
	-18,
}
var yyDef = [...]int{

	67, -2, 0, 2, 4, 5, 0, -2, 0, 0,
	0, 50, 0, 69, 70, 0, 0, 0, 0, 75,
	0, -2, 6, 15, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 8,
	9, 0, 52, 71, 72, 73, 0, 0, 0, 78,
	85, 81, 3, 83, 16, 18, 19, 21, 75, 32,
	33, 34, 35, 36, 67, 0, 0, -2, 0, 53,
	54, 55, 56, 57, 58, 59, 60, 61, 62, 63,
	64, 65, 66, 11, 74, 76, 85, 0, 0, 86,
	0, 84, 0, 0, 0, 0, 0, 0, 0, 0,
	30, 31, 37, 68, 38, 0, 0, 0, 75, 40,
	-2, 0, 0, 12, 0, 80, 79, 82, 14, 17,
	20, 22, 0, 0, 25, 26, 27, 28, 29, 0,
	42, 67, 47, 0, 0, 51, 0, 0, 77, 0,
	80, 41, 0, 0, 0, 0, 10, 13, 23, 0,
	43, 44, 45, 39, 48, 0, 24, 0, 0, 46,
	49,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 52, 3, 3, 51, 40, 3, 3,
	42, 43, 38, 36, 44, 37, 35, 39, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 26, 41,
	33, 48, 34, 25, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 49, 3, 50, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 45, 47, 46,
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
		//line hawk.y:63
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
			yyVAL.stmt = &PipeStmt{yyDollar[1].stmt, yyDollar[3].sym}
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
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:180
		{
			yyVAL.stmt = &AssignStmt{ast, &IndexExpr{&Ident{ast, yyDollar[1].sym}, nil}, yyDollar[5].expr}
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line hawk.y:184
		{
			yyVAL.stmt = &AssignStmt{ast, &IndexExpr{&Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}, yyDollar[6].expr}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:188
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Add, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:192
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Sub, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:196
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Mul, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:200
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Div, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:204
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Mod, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:210
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Add, &Ident{ast, yyDollar[1].sym}, Lit(1)}}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:214
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Sub, &Ident{ast, yyDollar[1].sym}, Lit(1)}}
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
			yyVAL.stmt = &PrintStmt{yyDollar[1].sym, yyDollar[2].exprlist}
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
			yyVAL.stmt = &IfStmt{yyDollar[2].expr, yyDollar[3].blockstmt, yyDollar[4].stmt}
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
			yyVAL.stmt = &ForStmt{yyDollar[2].stmt, yyDollar[4].expr, yyDollar[6].stmt, yyDollar[7].blockstmt}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:286
		{
			yyVAL.stmt = &ForStmt{nil, yyDollar[2].expr, nil, yyDollar[3].blockstmt}
		}
	case 48:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:292
		{
			yyVAL.stmt = &ForeachStmt{&Ident{ast, yyDollar[2].sym}, nil, yyDollar[4].expr, yyDollar[5].blockstmt}
		}
	case 49:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line hawk.y:296
		{
			yyVAL.stmt = &ForeachStmt{&Ident{ast, yyDollar[2].sym}, &Ident{ast, yyDollar[4].sym}, yyDollar[6].expr, yyDollar[7].blockstmt}
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
			yyVAL.expr = &TernaryExpr{yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr}
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:311
		{
			yyVAL.expr = &FieldExpr{scanner, yyDollar[2].expr}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:315
		{
			yyVAL.expr = &BinaryExpr{OrOr, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:319
		{
			yyVAL.expr = &BinaryExpr{AndAnd, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:323
		{
			yyVAL.expr = &BinaryExpr{Eq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:327
		{
			yyVAL.expr = &BinaryExpr{NotEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:331
		{
			yyVAL.expr = &BinaryExpr{LtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:335
		{
			yyVAL.expr = &BinaryExpr{GtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:339
		{
			yyVAL.expr = &BinaryExpr{Lt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:343
		{
			yyVAL.expr = &BinaryExpr{Gt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:347
		{
			yyVAL.expr = &BinaryExpr{Add, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:351
		{
			yyVAL.expr = &BinaryExpr{Sub, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:355
		{
			yyVAL.expr = &BinaryExpr{Mul, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:359
		{
			yyVAL.expr = &BinaryExpr{Div, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:363
		{
			yyVAL.expr = &BinaryExpr{Mod, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:367
		{
			yyVAL.expr = &BinaryExpr{Concat, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:372
		{
			yyVAL.expr = nil
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:376
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:383
		{
			yyVAL.expr = Lit(yyDollar[1].num)
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:387
		{
			yyVAL.expr = StringLit(yyDollar[1].sym)
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:391
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:395
		{
			yyVAL.expr = &UnaryExpr{Minus, yyDollar[2].expr}
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:399
		{
			yyVAL.expr = &UnaryExpr{Not, yyDollar[2].expr}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:403
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:407
		{
			yyVAL.expr = &Ident{ast, yyDollar[1].sym}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:411
		{
			yyVAL.expr = &CallExpr{yyDollar[1].sym, nil}
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:415
		{
			yyVAL.expr = &CallExpr{yyDollar[1].sym, yyDollar[3].exprlist}
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:419
		{
			yyVAL.expr = &ArrayLit{}
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:423
		{
			yyVAL.expr = &ArrayLit{yyDollar[2].exprlist}
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:428
		{
			yyVAL.expr = &IndexExpr{&Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:434
		{
			yyVAL.exprlist = []Expr{yyDollar[1].expr}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:438
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	}
	goto yystack /* stack new state and value */
}
