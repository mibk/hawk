package shellexec

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"unicode"
	"unicode/utf8"
)

// Command parses line using a shell-like syntax and returns
// the os/exec.Cmd struct to execute the line.
func Command(line string) (*exec.Cmd, error) {
	p := parser{s: line, getenv: os.Getenv}
	c := p.parseLine()
	if p.err != nil {
		return nil, p.err
	}
	cmd := exec.Command(c.cmd, c.args...)
	cmd.Env = c.env
	return cmd, nil
}

type cmd struct {
	cmd  string
	args []string
	env  []string
}

type parser struct {
	buf    bytes.Buffer
	s      string
	last   rune
	peeked *rune
	getenv func(key string) string
	err    error

	identBuf bytes.Buffer
}

const eof rune = utf8.MaxRune + 1

func (p *parser) next() rune {
	if p.peeked != nil {
		r := *p.peeked
		p.peeked = nil
		return r
	}

	if len(p.s) == 0 {
		p.last = eof
		return eof
	}
	var size int
	p.last, size = utf8.DecodeRuneInString(p.s)
	p.s = p.s[size:]
	if p.last == utf8.RuneError {
		p.errorf("invalid UTF-8 encoding")
		return eof
	}

	if p.last == '\\' && p.s != "" && p.s[0] == '\n' {
		// line continuation; remove it from the input
		p.s = p.s[1:]
		return p.next()
	}

	return p.last
}

func (p *parser) backup() {
	p.peeked = &p.last
}

func (p *parser) token() string {
	t := p.buf.String()
	p.buf.Reset()
	return t
}

func (p *parser) errorf(format string, args ...interface{}) {
	if p.err != nil {
		return
	}
	p.s = p.s[:0]
	p.err = fmt.Errorf(format, args...)
}

func (p *parser) parseLine() cmd {
	var c cmd
loop:
	for {
		r := p.next()
		switch {
		case unicode.IsSpace(r):
			continue
		case r == eof:
			break loop
		}
		p.backup()

		if c.cmd == "" {
			if isVarAssign := p.parseVarAssign(); isVarAssign {
				c.env = append(c.env, p.token())
			} else {
				c.cmd = p.token()
			}
		} else {
			p.parseField()
			c.args = append(c.args, p.token())
		}
	}
	if c.cmd == "" {
		p.errorf("empty command")
	}
	return c
}

func (p *parser) parseVarAssign() (isVarAssign bool) {
	v := p.parseIdent()
	p.buf.WriteString(v)
	if v != "" && p.next() == '=' {
		isVarAssign = true
	}
	p.backup()
	p.parseField()
	return
}

func (p *parser) parseField() {
	esc := false
	for {
		r := p.next()
		if r == eof {
			break
		}

		if esc {
			p.buf.WriteRune(r)
			esc = false
			continue
		}
		if unicode.IsSpace(r) {
			break
		}
		switch r {
		case '\'':
			p.parseSingleQuotes()
		case '"':
			p.parseDoubleQuotes()
		case '\\':
			esc = true
			continue
		case '$':
			p.parseVarExpr()
			p.backup()
			continue
		case '|', '&', ';', '<', '>', '(', ')', '`',
			// Forbid these characters as they may need to be
			// quoted under certain circumstances.
			'*', '?', '[', '#', '~':
			p.errorf("unsupported character: %c", r)
		default:
			p.buf.WriteRune(r)
		}
	}
}

func (p *parser) parseSingleQuotes() {
	for {
		switch r := p.next(); r {
		case '\'':
			return
		case eof:
			p.errorf("string not terminated")
			return
		default:
			p.buf.WriteRune(r)
		}
	}
}

func (p *parser) parseDoubleQuotes() {
	var esc bool
	for {
		r := p.next()
		if r == eof {
			p.errorf("string not terminated")
			return
		}

		if esc {
			switch r {
			default:
				p.buf.WriteRune('\\')
				fallthrough
			case '$', '`', '"', '\\':
				p.buf.WriteRune(r)
			}
			esc = false
			continue
		}
		switch r {
		case '"':
			return
		case '\\':
			esc = true
			continue
		case '$':
			p.parseVarExpr()
			continue
		case '`':
			p.errorf("unsupported character inside string: %c", r)
		}
		p.buf.WriteRune(r)
	}
}

func (p *parser) parseVarExpr() {
	switch r := p.next(); r {
	case '(':
		p.errorf("command substitution '$(command)' or arithmetic expansion '$((expression))' not supported")
	case '{':
		p.errorf("parameter expansion '${expression}' not supported")
	case '@', '*', '#', '?', '-', '$', '!', '0':
		p.errorf("special parameters not supported: $%c", r)
	case '1', '2', '3', '4', '5', '6', '7', '8', '9':
		p.errorf("positional parameters not supported: $%c", r)
	}
	p.backup()

	v := "$"
	if name := p.parseIdent(); name != "" {
		v = p.getenv(name)
	}
	p.buf.WriteString(v)
}

func (p *parser) parseIdent() string {
	p.identBuf.Reset()
	for {
		r := p.next()
		if !(r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' || r == '_' ||
			p.identBuf.Len() > 0 && r >= '0' && r <= '9') {
			p.backup()
			return p.identBuf.String()
		}
		p.identBuf.WriteRune(r)
	}
}
