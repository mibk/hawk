# Hawk

This project started as a reaction to Brian Kernighan's talk on [How to succeed in
language design without really trying](https://www.youtube.com/watch?v=Sg4U4r_AgJU).
I've been always quite passionate about parsers, compilers, languages, but couldn't
really find a reason to actually build my own language. This talk motivated me so much
that I just had to do something. Because of lack of ideas I decided to make an Awk
clone.

It was never meant to be a replacement for Awk. I'm doing it just for fun. Nevertheless
there might be the potential to make it actually useful. Right now, Hawk is unoptimised,
definitely slower than the original Awk, and the design is in some ways probably quite
naive. Anyway, for anyone interested in making a new version of Awk, what would you add,
leave out, simplify?

The language is basically Awk with a Go-like syntax. Some of the features are taken
from PHP (especially arrays/maps).

## Installation

```
go get -u github.com/mibk/hawk
```

## Usage

Running `hawk -help` gives you an extended help message on how to use Hawk, although the
best documentation for the Hawk syntax currently is `test/run` directory.

```
Usage: hawk 'program' [file ...]
  or:  hawk -f progfile [file ...]

Hawk is an Awk clone. Program is a set of pattern {action} pairs. Hawk reads
from all of the present files and for each line of each file executes all the
provided pairs. If no files are present, hawk reads from stdin.

Run hawk -help for a detailed help message.

Flags:
  -F string
        set the field separator, FS
  -f file
        read program from file
  -help
        display an extended help
```

## Examples of Hawk programs

### Emulate wc

```
{ chars += len($0) + 1; words += NF }
END { printf "%7d %7d %7d\n", NR, words, chars }
```

### Filter values of a list.

Having a list like this,

```
   92 ./value/undefined.go
  260 ./value/value.go
  140 ./value/array.go
  136 ./compiler/internal/hawkc/semantic.go
  319 ./compiler/internal/hawkc/expr.go
  339 ./compiler/internal/hawkc/lex.go
  211 ./compiler/internal/hawkc/syntax.go
   71 ./compiler/internal/hawkc/builtin.go
  277 ./compiler/internal/hawkc/stmt.go
 1214 ./compiler/internal/hawkc/hawk.go
   35 ./compiler/compiler.go
  332 ./scan/scan.go
   91 ./scan/scan_test.go
   82 ./main.go
  105 ./test/errcheck/short_test.go
   47 ./test/run/run_test.go
```

the whole program to print only the lines with the first field greater than
200 is:

```
$1 > 200
```

### Summarise a shopping list.

```
lunch    4.99
lunch    3.99
snack    1
lunch    2.49
snack    0.89
coffee   2.49
lunch    3.20
```

You can pipe any statement to an external program.

```
{ m[$1] += $2 }
END {
    for k, v in m {
        print v, k
    } | "sort -n"
}
```

The output is:

```
1.89 snack
2.49 coffee
14.67 lunch
```
