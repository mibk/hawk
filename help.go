package main

const extendedHelp = shortHelp + `

The Hawk language

1. Program structure

A Hawk program is a sequence of pattern {action} pairs and function declarations.
Either pattern, or {action} can be omitted. If pattern is omitted, it matches every
line. If {action} is omitted, the default action, printing the matched line, is
performed.

A pattern can be:
	BEGIN
	END
	expression

The actions of the BEGIN patterns are performed at the beginning, before the input
parsing. The actions of the END patterns are performed at the end of parsing the
input.

The statements are terminated by semicolons. The compiler inserts semicolons after
newlines using the same rules as the Go programming language.

There are two forms of comments: // line comments, and /* block comments */.

2. Statements

	The statement syntax is derived from the Go programming language. The range
	statement syntax has a different keyword (in). There are two additional
	statements for printing (print and printf). The pipe statement pipes the output
	of a statement through an external program given in a string constant.

	if expr { statements }

	if expr { statements } else ...

	for opt_expr; opt_expr; opt_expr { statements }

	for opt_expr { statements } // like C's while

	for var in array { statements }

	continue

	break

	statement | "command" // pipe statement

	print expression_list or printf format, expression_list


	assignment operators: =  +=  -=  *=  /=  %=

	post inc and dec:     ++  --


3. Expressions

	The expression syntax is derived  from the Go programming language and also from
	the Awk language.

	Operators ordered by precedence:

	ternary operator:   ?  :
	logical or:         ||
	logical and:        &&
	relational:         <  >  <=  >=  ==  !=
	regexp matching:    ~  !~
	concatenation:      .
	add operations:     +  -
	mul operations:     *  /  %
	unary:              +  -
	logical not:        !
	field:              $


4. Data types

Hawk has 3 basic data types: strings, numbers and booleans. All arithmetic operations are
done using floating-point arithmetics.

	boolean:  true  false
	number:   12  12.38  0xFF  0Xba
	string:   "double\nquotes"  'single \'quotes\''` + "  `raw strings with ``escaped`` back-quotes`" + `


5. Built-in variables

	FILENAME   name of the current input file

	FNR        current record number in FILENAME

	FS         splits records into fields using FS as a regexp

	NF         number of fields in the current record

	NR         current number of records in the whole input stream

	OFS        output fields separator (default is " ")

	ORS        output record separator (default is "\n")

	RS         splits input into records using RS as a regexp


6. Built-in functions

	len(expr)   returns the length of a string or the count of items in an array

	sprintf(format, ...expr)


	Arithmetic functions:

	atan2(x, y)

	cos(x)

	exp(x)

	log(x)

	sin(x)

	sqrt(x)
`
