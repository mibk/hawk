// Package shellexec provides cross-platform shell-like command
// execution. It supports a small subset of the Shell Command
// Language. The specification can be found in the Open Group Base
// Specifications Issue 7, IEEE Std 1003.1-2008, 2016 Edition,
// available from
//
//	http://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html
//
// Package shellexec isn't a replacement for os/exec. It's rather
// just a convenience package for use-cases that require/prefer
// having just a single string to define the command to execute.
package shellexec
