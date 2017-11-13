package shellexec

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		name string
		line string
		cmd  string
		args []string
		env  []string
	}{
		{
			"all escape chars",
			`  echo  \|\&\;\<\>\(\)\$\\\"\'\ \	\` + "\n\\`\\",
			"echo", []string{`|&;<>()$\"'` + " \t`"}, nil,
		},
		{
			"single-quote strings",
			`foo'bar''boo&;<>'`,
			`foobarboo&;<>`, nil, nil,
		},
		{
			"other special characters",
			`echo \*\?\[\#\~\=\%  =%`,
			"echo", []string{"*?[#~=%", "=%"}, nil,
		},
		{
			"escaped =",
			`weird\=name`,
			"weird=name", nil, nil,
		},
		{
			"env variables",
			` X=3  y=4  _12=5 echo Z=12`,
			"echo", []string{"Z=12"}, []string{"X=3", "y=4", "_12=5"},
		},
		{
			"invalid var",
			`1=1`,
			"1=1", nil, nil,
		},
		{
			"invalid var 2",
			`č=1`,
			"č=1", nil, nil,
		},
		{
			"double-quote string",
			`echo "\\\"\$\` + "\n\\`" + `\G" \e`,
			"echo", []string{`\"$` + "`\\G", "e"}, nil,
		},
		{
			"variable expansion",
			`X=$val  echo "$PATH$_A" $val'x3' $EDITOR`,
			"echo", []string{"/usr/local/bin[A]", "3x3", "syd"}, []string{"X=3"},
		},
		{
			"program from var",
			`$EDITOR file:32`,
			"syd", []string{"file:32"}, nil,
		},
		{
			"line continuation",
			"VA\\\nRIABLE=X ech\\\no \\\n 'fo\\\no' bar\\\n",
			"echo", []string{"foo", "bar"}, []string{"VARIABLE=X"},
		},
	}

	for _, tt := range tests {
		c, err := parseString(tt.line)
		if err != nil {
			t.Errorf("%s: unexpected err: %v", tt.name, err)
			continue
		}
		if c.cmd != tt.cmd {
			t.Errorf("%s: cmd: got %q, want %q", tt.name, c.cmd, tt.cmd)
		}
		if !reflect.DeepEqual(c.args, tt.args) {
			t.Errorf("%s: args: got %q, want %q", tt.name, c.args, tt.args)
		}
		if !reflect.DeepEqual(c.env, tt.env) {
			t.Errorf("%s: env: got %q, want %q", tt.name, c.env, tt.env)
		}

	}
}

func TestParseErrors(t *testing.T) {
	tests := []struct {
		name string
		line string
		err  string
	}{
		{
			"empty",
			`  X=Y`,
			"empty command",
		},
		{
			"unterminated single-quote string",
			`echo 'always'be'closin`,
			"string not terminated",
		},
		{
			"unterminated double-quote string",
			`echo "always"be"closin`,
			"string not terminated",
		},
		{
			"unsupported char in string",
			"echo \"`echo this`\"",
			"unsupported character inside string",
		},
		{
			"command substitution",
			"echo $(cat file)",
			"command substitution '$(command)' or arithmetic expansion '$((expression))'",
		},
		{
			"parameter expansion",
			"echo ${var}",
			"parameter expansion '${expression}' not supported",
		},
		{
			"special parameter",
			"echo $@",
			"special parameters not supported: $@",
		},
		{
			"positional parameter",
			"echo $3",
			"positional parameters not supported: $3",
		},
		{
			"'`' inside string",
			"echo \"`head -1 file`\"",
			"unsupported character inside string: `",
		},
		{
			"unsupported character",
			"cat file | sort",
			"unsupported character: |",
		},
		{
			"invalid UTF-8 string",
			"echo \x80",
			"invalid UTF-8 encoding",
		},
	}

	for _, tt := range tests {
		_, err := parseString(tt.line)
		if err == nil {
			t.Errorf("%s: unexpectadly succeeded", tt.name)
			continue
		}
		if !strings.Contains(err.Error(), tt.err) {
			t.Errorf("%s: got %q, want %q", tt.name, err, tt.err)
		}
	}
}

func TestParseInvalidChars(t *testing.T) {
	invalid := []rune{'|', '&', ';', '<', '>', '(', ')', '`',
		'*', '?', '[', '#', '~'}

	for _, r := range invalid {
		if _, err := parseString(string(r)); err == nil ||
			!strings.Contains(err.Error(), "unsupported") {
			t.Errorf("char %q should be invalid", r)
		}
	}
}

func BenchmarkParsingSimpleLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := parseString(`VAR_A=aaa VAR_B=bbb echo 'all'work"and \"no\" \$play"`)
		if err != nil {
			b.Fatalf("unexpected err: %v", err)
		}
	}
}

func parseString(s string) (cmd, error) {
	getenv := func(key string) string { return _Env[key] }
	p := parser{s: s, getenv: getenv}
	return p.parseLine(), p.err
}

var _Env = map[string]string{
	"PATH":   "/usr/local/bin",
	"EDITOR": "syd",
	"_A":     "[A]",
	"val":    "3",
}
