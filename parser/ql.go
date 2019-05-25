// Code generated by goyacc -o ql.go -v ql.output -p ql ql.y. DO NOT EDIT.

//line ql.y:2
package parser

import __yyfmt__ "fmt"

//line ql.y:2

//line ql.y:8
type qlSymType struct {
	yys int
	loc Location

	strVal string
}

const LEX_ERROR = 57346
const SELECT = 57347
const FROM = 57348
const WHERE = 57349

var qlToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"SELECT",
	"FROM",
	"WHERE",
}
var qlStatenames = [...]string{}

const qlEofCode = 1
const qlErrCode = 2
const qlInitialStackSize = 16

//line ql.y:29

//line yacctab:1
var qlExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const qlPrivate = 57344

const qlLast = 7

var qlAct = [...]int{

	7, 6, 4, 2, 3, 5, 1,
}
var qlPact = [...]int{

	-3, -3, -1000, -1000, -5, -1000, -7, -1000,
}
var qlPgo = [...]int{

	0, 6, 3, 4,
}
var qlR1 = [...]int{

	0, 1, 1, 2, 3,
}
var qlR2 = [...]int{

	0, 1, 2, 1, 3,
}
var qlChk = [...]int{

	-1000, -1, -2, -3, 5, -2, 6, 7,
}
var qlDef = [...]int{

	0, -2, 1, 3, 0, 2, 0, 4,
}
var qlTok1 = [...]int{

	1,
}
var qlTok2 = [...]int{

	2, 3, 4, 5, 6, 7,
}
var qlTok3 = [...]int{
	0,
}

var qlErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	qlDebug        = 0
	qlErrorVerbose = false
)

type qlLexer interface {
	Lex(lval *qlSymType) int
	Error(s string)
}

type qlParser interface {
	Parse(qlLexer) int
	Lookahead() int
}

type qlParserImpl struct {
	lval  qlSymType
	stack [qlInitialStackSize]qlSymType
	char  int
}

func (p *qlParserImpl) Lookahead() int {
	return p.char
}

func qlNewParser() qlParser {
	return &qlParserImpl{}
}

const qlFlag = -1000

func qlTokname(c int) string {
	if c >= 1 && c-1 < len(qlToknames) {
		if qlToknames[c-1] != "" {
			return qlToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func qlStatname(s int) string {
	if s >= 0 && s < len(qlStatenames) {
		if qlStatenames[s] != "" {
			return qlStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func qlErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !qlErrorVerbose {
		return "syntax error"
	}

	for _, e := range qlErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + qlTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := qlPact[state]
	for tok := TOKSTART; tok-1 < len(qlToknames); tok++ {
		if n := base + tok; n >= 0 && n < qlLast && qlChk[qlAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if qlDef[state] == -2 {
		i := 0
		for qlExca[i] != -1 || qlExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; qlExca[i] >= 0; i += 2 {
			tok := qlExca[i]
			if tok < TOKSTART || qlExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if qlExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += qlTokname(tok)
	}
	return res
}

func qllex1(lex qlLexer, lval *qlSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = qlTok1[0]
		goto out
	}
	if char < len(qlTok1) {
		token = qlTok1[char]
		goto out
	}
	if char >= qlPrivate {
		if char < qlPrivate+len(qlTok2) {
			token = qlTok2[char-qlPrivate]
			goto out
		}
	}
	for i := 0; i < len(qlTok3); i += 2 {
		token = qlTok3[i+0]
		if token == char {
			token = qlTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = qlTok2[1] /* unknown char */
	}
	if qlDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", qlTokname(token), uint(char))
	}
	return char, token
}

func qlParse(qllex qlLexer) int {
	return qlNewParser().Parse(qllex)
}

func (qlrcvr *qlParserImpl) Parse(qllex qlLexer) int {
	var qln int
	var qlVAL qlSymType
	var qlDollar []qlSymType
	_ = qlDollar // silence set and not used
	qlS := qlrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	qlstate := 0
	qlrcvr.char = -1
	qltoken := -1 // qlrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		qlstate = -1
		qlrcvr.char = -1
		qltoken = -1
	}()
	qlp := -1
	goto qlstack

ret0:
	return 0

ret1:
	return 1

qlstack:
	/* put a state and value onto the stack */
	if qlDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", qlTokname(qltoken), qlStatname(qlstate))
	}

	qlp++
	if qlp >= len(qlS) {
		nyys := make([]qlSymType, len(qlS)*2)
		copy(nyys, qlS)
		qlS = nyys
	}
	qlS[qlp] = qlVAL
	qlS[qlp].yys = qlstate

qlnewstate:
	qln = qlPact[qlstate]
	if qln <= qlFlag {
		goto qldefault /* simple state */
	}
	if qlrcvr.char < 0 {
		qlrcvr.char, qltoken = qllex1(qllex, &qlrcvr.lval)
	}
	qln += qltoken
	if qln < 0 || qln >= qlLast {
		goto qldefault
	}
	qln = qlAct[qln]
	if qlChk[qln] == qltoken { /* valid shift */
		qlrcvr.char = -1
		qltoken = -1
		qlVAL = qlrcvr.lval
		qlstate = qln
		if Errflag > 0 {
			Errflag--
		}
		goto qlstack
	}

qldefault:
	/* default state action */
	qln = qlDef[qlstate]
	if qln == -2 {
		if qlrcvr.char < 0 {
			qlrcvr.char, qltoken = qllex1(qllex, &qlrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if qlExca[xi+0] == -1 && qlExca[xi+1] == qlstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			qln = qlExca[xi+0]
			if qln < 0 || qln == qltoken {
				break
			}
		}
		qln = qlExca[xi+1]
		if qln < 0 {
			goto ret0
		}
	}
	if qln == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			qllex.Error(qlErrorMessage(qlstate, qltoken))
			Nerrs++
			if qlDebug >= 1 {
				__yyfmt__.Printf("%s", qlStatname(qlstate))
				__yyfmt__.Printf(" saw %s\n", qlTokname(qltoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for qlp >= 0 {
				qln = qlPact[qlS[qlp].yys] + qlErrCode
				if qln >= 0 && qln < qlLast {
					qlstate = qlAct[qln] /* simulate a shift of "error" */
					if qlChk[qlstate] == qlErrCode {
						goto qlstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if qlDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", qlS[qlp].yys)
				}
				qlp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if qlDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", qlTokname(qltoken))
			}
			if qltoken == qlEofCode {
				goto ret1
			}
			qlrcvr.char = -1
			qltoken = -1
			goto qlnewstate /* try again in the same state */
		}
	}

	/* reduction by production qln */
	if qlDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", qln, qlStatname(qlstate))
	}

	qlnt := qln
	qlpt := qlp
	_ = qlpt // guard against "declared and not used"

	qlp -= qlR2[qln]
	// qlp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if qlp+1 >= len(qlS) {
		nyys := make([]qlSymType, len(qlS)*2)
		copy(nyys, qlS)
		qlS = nyys
	}
	qlVAL = qlS[qlp+1]

	/* consult goto table to find next state */
	qln = qlR1[qln]
	qlg := qlPgo[qln]
	qlj := qlg + qlS[qlp].yys + 1

	if qlj >= qlLast {
		qlstate = qlAct[qlg]
	} else {
		qlstate = qlAct[qlj]
		if qlChk[qlstate] != -qln {
			qlstate = qlAct[qlg]
		}
	}
	// dummy call; replaced with literal code
	switch qlnt {

	}
	goto qlstack /* stack new state and value */
}