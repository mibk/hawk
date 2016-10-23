//line hawk.y:2
package compiler

import __yyfmt__ "fmt"

//line hawk.y:2
import (
	"bufio"
	"io"

	"github.com/mibk/hawk/scan"
)

var (
	scanner *scan.Scanner
	ast     *Program

	defaultAction *BlockStmt
)

//line hawk.y:19
type yySymType struct {
	yys      int
	num      int
	sym      string
	symlist  []string
	decl     Decl
	decllist []Decl
	expr     Expr
	exprlist []Expr
	stmt     Stmt
	stmtlist []Stmt
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
const BREAK = 57355
const CONTINUE = 57356
const INC = 57357
const DEC = 57358
const ADDEQ = 57359
const SUBEQ = 57360
const MULEQ = 57361
const DIVEQ = 57362
const MODEQ = 57363
const FUNC = 57364
const RETURN = 57365
const OROR = 57366
const ANDAND = 57367
const EQ = 57368
const NE = 57369
const LE = 57370
const GE = 57371

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

//line hawk.y:423

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
	43, 63,
	-2, 7,
	-1, 21,
	43, 62,
	-2, 1,
	-1, 65,
	43, 62,
	-2, 37,
	-1, 106,
	43, 63,
	-2, 21,
}

const yyNprod = 82
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 399

var yyAct = [...]int{

	50, 7, 55, 54, 6, 103, 58, 49, 112, 22,
	53, 38, 39, 46, 85, 13, 19, 14, 138, 45,
	47, 89, 7, 64, 56, 66, 67, 68, 69, 70,
	71, 72, 73, 74, 75, 76, 77, 78, 79, 114,
	23, 124, 13, 19, 14, 15, 16, 86, 84, 129,
	130, 18, 82, 131, 83, 80, 23, 142, 20, 126,
	12, 17, 88, 100, 21, 102, 106, 99, 3, 105,
	104, 101, 15, 16, 134, 35, 36, 37, 18, 33,
	34, 35, 36, 37, 116, 20, 137, 113, 17, 56,
	51, 117, 118, 119, 120, 121, 122, 123, 110, 115,
	109, 40, 87, 1, 52, 125, 59, 127, 128, 139,
	133, 2, 108, 5, 13, 57, 14, 63, 4, 0,
	64, 0, 65, 60, 61, 113, 0, 100, 0, 0,
	0, 135, 136, 62, 0, 0, 0, 141, 11, 143,
	0, 140, 0, 56, 15, 16, 105, 145, 144, 0,
	18, 41, 0, 23, 42, 43, 44, 20, 0, 12,
	17, 13, 57, 14, 63, 0, 0, 64, 0, 65,
	60, 61, 0, 0, 0, 0, 0, 0, 0, 0,
	62, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 15, 16, 0, 0, 0, 0, 18, 0, 0,
	0, 0, 0, 0, 20, 0, 12, 17, 24, 0,
	25, 26, 27, 28, 29, 30, 31, 32, 33, 34,
	35, 36, 37, 13, 19, 14, 0, 13, 19, 14,
	0, 24, 132, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 13, 19, 14, 0,
	8, 9, 0, 15, 16, 111, 0, 15, 16, 18,
	0, 0, 0, 18, 10, 0, 20, 48, 12, 17,
	20, 0, 12, 17, 0, 0, 15, 16, 0, 0,
	0, 0, 18, 0, 0, 0, 0, 0, 0, 20,
	0, 12, 17, 24, 0, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 0, 0,
	0, 0, 23, 97, 98, 92, 93, 94, 95, 96,
	24, 0, 25, 26, 27, 28, 29, 30, 31, 32,
	33, 34, 35, 36, 37, 0, 0, 81, 46, 0,
	0, 0, 0, 0, 90, 91, 24, 107, 25, 26,
	27, 28, 29, 30, 31, 32, 33, 34, 35, 36,
	37, 24, 0, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 27, 28,
	29, 30, 31, 32, 33, 34, 35, 36, 37,
}
var yyPact = [...]int{

	242, -1000, 25, -1000, -1000, -1000, -3, 337, -3, -3,
	96, -1000, 38, -1000, -1000, 38, 38, 38, 223, -27,
	219, 242, -1000, 110, 223, 223, 223, 223, 223, 223,
	223, 223, 223, 223, 223, 223, 223, 223, -1000, -1000,
	15, -1000, -1000, -1000, -1000, 296, 11, 223, -1000, 5,
	337, -1000, 23, -24, -1000, -1000, 337, 298, -1000, -1000,
	-1000, -1000, 223, 223, 223, 157, 322, 349, 360, 45,
	45, 45, 45, 45, 45, 39, 39, -1000, -1000, -1000,
	95, -1000, -1000, 5, 207, -40, 223, -5, 110, 78,
	223, 223, 223, 223, 223, 223, 223, -1000, -1000, -1000,
	337, -1, 269, 20, -3, -1000, 337, 223, 8, -1000,
	12, -1000, -1000, 337, -1000, -24, -1000, 337, 184, 337,
	337, 337, 337, 337, 223, 63, 223, -1000, 337, -3,
	81, -1000, -28, -1000, 13, 18, -1000, -1000, 223, -1000,
	-1000, -1000, 157, 337, -3, -1000,
}
var yyPgo = [...]int{

	0, 68, 118, 113, 112, 111, 0, 4, 138, 7,
	10, 3, 5, 6, 110, 109, 106, 2, 104, 103,
	102, 14,
}
var yyR1 = [...]int{

	0, 19, 5, 5, 1, 1, 2, 2, 2, 2,
	3, 4, 4, 4, 17, 18, 18, 18, 10, 10,
	10, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 12, 12, 13,
	14, 14, 15, 15, 16, 16, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 7, 7, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 9, 9, 20, 20,
	21, 21,
}
var yyR2 = [...]int{

	0, 2, 1, 3, 1, 1, 2, 1, 2, 2,
	6, 0, 1, 3, 4, 0, 1, 3, 1, 1,
	3, 1, 3, 6, 3, 3, 3, 3, 3, 2,
	2, 1, 1, 1, 1, 2, 2, 0, 1, 4,
	0, 2, 1, 1, 7, 3, 1, 5, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 0, 1, 1, 1, 2, 2, 2, 3,
	1, 3, 5, 2, 4, 4, 1, 3, 0, 1,
	0, 1,
}
var yyChk = [...]int{

	-1000, -19, -5, -1, -2, -3, -7, -6, 8, 9,
	22, -8, 49, 4, 6, 34, 35, 50, 40, 5,
	47, 39, -17, 43, 24, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, -17, -17,
	5, -8, -8, -8, -8, -6, 40, 47, 48, -9,
	-6, -1, -18, -10, -11, -17, -6, 5, -13, -16,
	13, 14, 23, 7, 10, 12, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	40, 41, 41, -9, -6, -21, 42, -20, 39, 45,
	46, 47, 17, 18, 19, 20, 21, 15, 16, -7,
	-6, -9, -6, -12, -7, -11, -6, 25, -4, 5,
	-21, 48, 48, -6, 44, -10, 6, -6, -6, -6,
	-6, -6, -6, -6, 42, -17, 39, -17, -6, 41,
	42, 41, 48, -14, 11, -7, -17, 5, 46, -15,
	-13, -17, 39, -6, -12, -17,
}
var yyDef = [...]int{

	62, -2, 0, 2, 4, 5, 0, -2, 0, 0,
	0, 46, 0, 64, 65, 0, 0, 0, 0, 70,
	0, -2, 6, 15, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 8, 9,
	0, 48, 66, 67, 68, 0, 0, 0, 73, 80,
	76, 3, 78, 16, 18, 19, 21, 70, 31, 32,
	33, 34, 62, 0, 0, -2, 0, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	11, 69, 71, 80, 0, 0, 81, 0, 79, 0,
	0, 0, 0, 0, 0, 0, 0, 29, 30, 35,
	63, 36, 0, 0, 0, 38, -2, 0, 0, 12,
	0, 75, 74, 77, 14, 17, 20, 22, 0, 24,
	25, 26, 27, 28, 0, 40, 62, 45, 47, 0,
	0, 72, 75, 39, 0, 0, 10, 13, 0, 41,
	42, 43, 37, 23, 0, 44,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 50, 3, 3, 49, 38, 3, 3,
	40, 41, 36, 34, 42, 35, 3, 37, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 25, 39,
	32, 46, 33, 24, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 47, 3, 48, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 43, 45, 44,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 26, 27, 28, 29, 30, 31,
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
		//line hawk.y:59
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
					panic("unreachable")
				}
			}
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:78
		{
			yyVAL.decllist = []Decl{yyDollar[1].decl}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:82
		{
			yyVAL.decllist = append(yyDollar[1].decllist, yyDollar[3].decl)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:88
		{
			yyVAL.decl = yyDollar[1].decl
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:92
		{
			yyVAL.decl = yyDollar[1].decl
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:98
		{
			yyVAL.decl = &PatternAction{yyDollar[1].expr, yyDollar[2].stmt}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:102
		{
			yyVAL.decl = &PatternAction{yyDollar[1].expr, defaultAction}
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:106
		{
			yyVAL.decl = &BeginAction{yyDollar[2].stmt}
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:110
		{
			yyVAL.decl = &EndAction{yyDollar[2].stmt}
		}
	case 10:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line hawk.y:116
		{
			yyVAL.decl = &FuncDecl{&FuncScope{}, yyDollar[2].sym, yyDollar[4].symlist, yyDollar[6].stmt}
		}
	case 11:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:121
		{
			yyVAL.symlist = nil
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:125
		{
			yyVAL.symlist = []string{yyDollar[1].sym}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:129
		{
			yyVAL.symlist = append(yyDollar[1].symlist, yyDollar[3].sym)
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:135
		{
			yyVAL.stmt = &BlockStmt{yyDollar[2].stmtlist}
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:140
		{
			yyVAL.stmtlist = nil
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:144
		{
			yyVAL.stmtlist = []Stmt{yyDollar[1].stmt}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:148
		{
			yyVAL.stmtlist = append(yyDollar[1].stmtlist, yyDollar[3].stmt)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:154
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:158
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:162
		{
			yyVAL.stmt = &PipeStmt{yyDollar[1].stmt, yyDollar[3].sym}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:168
		{
			yyVAL.stmt = &ExprStmt{yyDollar[1].expr}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:172
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}
		}
	case 23:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line hawk.y:176
		{
			yyVAL.stmt = &AssignStmt{ast, &IndexExpr{&Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}, yyDollar[6].expr}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:180
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Add, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:184
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Sub, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:188
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Mul, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:192
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Div, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:196
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Mod, &Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:202
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Add, &Ident{ast, yyDollar[1].sym}, Lit(1)}}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:206
		{
			yyVAL.stmt = &AssignStmt{ast, &Ident{ast, yyDollar[1].sym}, &BinaryExpr{Sub, &Ident{ast, yyDollar[1].sym}, Lit(1)}}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:210
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:214
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:218
		{
			yyVAL.stmt = &StatusStmt{StatusBreak}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:222
		{
			yyVAL.stmt = &StatusStmt{StatusContinue}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:226
		{
			yyVAL.stmt = &ReturnStmt{ast, yyDollar[2].expr}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:230
		{
			yyVAL.stmt = &PrintStmt{yyDollar[1].sym, yyDollar[2].exprlist}
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:235
		{
			yyVAL.stmt = nil
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:239
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:245
		{
			yyVAL.stmt = &IfStmt{yyDollar[2].expr, yyDollar[3].stmt, yyDollar[4].stmt}
		}
	case 40:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:250
		{
			yyVAL.stmt = nil
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:254
		{
			yyVAL.stmt = yyDollar[2].stmt
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:260
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:264
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 44:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line hawk.y:270
		{
			yyVAL.stmt = &ForStmt{yyDollar[2].stmt, yyDollar[4].expr, yyDollar[6].stmt, yyDollar[7].stmt}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:274
		{
			yyVAL.stmt = &ForStmt{nil, yyDollar[2].expr, nil, yyDollar[3].stmt}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:281
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 47:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:285
		{
			yyVAL.expr = &TernaryExpr{yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr}
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:289
		{
			yyVAL.expr = &FieldExpr{scanner, yyDollar[2].expr}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:293
		{
			yyVAL.expr = &BinaryExpr{OrOr, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:297
		{
			yyVAL.expr = &BinaryExpr{AndAnd, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:301
		{
			yyVAL.expr = &BinaryExpr{Eq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:305
		{
			yyVAL.expr = &BinaryExpr{NotEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:309
		{
			yyVAL.expr = &BinaryExpr{LtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:313
		{
			yyVAL.expr = &BinaryExpr{GtEq, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:317
		{
			yyVAL.expr = &BinaryExpr{Lt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:321
		{
			yyVAL.expr = &BinaryExpr{Gt, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:325
		{
			yyVAL.expr = &BinaryExpr{Add, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:329
		{
			yyVAL.expr = &BinaryExpr{Sub, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:333
		{
			yyVAL.expr = &BinaryExpr{Mul, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:337
		{
			yyVAL.expr = &BinaryExpr{Div, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:341
		{
			yyVAL.expr = &BinaryExpr{Mod, yyDollar[1].expr, yyDollar[3].expr}
		}
	case 62:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hawk.y:346
		{
			yyVAL.expr = nil
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:350
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:357
		{
			yyVAL.expr = Lit(yyDollar[1].num)
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:361
		{
			yyVAL.expr = StringLit(yyDollar[1].sym)
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:365
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:369
		{
			yyVAL.expr = &UnaryExpr{Minus, yyDollar[2].expr}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:373
		{
			yyVAL.expr = &UnaryExpr{Not, yyDollar[2].expr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:377
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:381
		{
			yyVAL.expr = &Ident{ast, yyDollar[1].sym}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:385
		{
			yyVAL.expr = &CallExpr{yyDollar[1].sym, nil}
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line hawk.y:389
		{
			yyVAL.expr = &CallExpr{yyDollar[1].sym, yyDollar[3].exprlist}
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hawk.y:393
		{
			yyVAL.expr = &ArrayLit{}
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:397
		{
			yyVAL.expr = &ArrayLit{yyDollar[2].exprlist}
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hawk.y:402
		{
			yyVAL.expr = &IndexExpr{&Ident{ast, yyDollar[1].sym}, yyDollar[3].expr}
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hawk.y:408
		{
			yyVAL.exprlist = []Expr{yyDollar[1].expr}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hawk.y:412
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[3].expr)
		}
	}
	goto yystack /* stack new state and value */
}
