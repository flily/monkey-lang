monkey-lang
===========

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/flily/monkey-lang)
![GitHub top language](https://img.shields.io/github/languages/top/flily/monkey-lang)

[![CI](https://github.com/flily/monkey-lang/actions/workflows/ci.yaml/badge.svg)](https://github.com/flily/monkey-lang/actions/workflows/ci.yaml)
[![Compatibility](https://github.com/flily/monkey-lang/actions/workflows/compatibility.yaml/badge.svg)](https://github.com/flily/monkey-lang/actions/workflows/compatibility.yaml)
[![codecov](https://codecov.io/gh/flily/monkey-lang/branch/main/graph/badge.svg?token=AQjSwtMbAE)](https://codecov.io/gh/flily/monkey-lang)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/53a4d78c3c2b4fd68c46f72ee55343f4)](https://www.codacy.com/gh/flily/monkey-lang/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=flily/monkey-lang&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/flily/monkey-lang)](https://goreportcard.com/report/github.com/flily/monkey-lang)

Implement of monkey language by Go, designed by Thorsten Ball in books Writing An Interpreter in Go and Writing An Compiler in Go.

Compatibility test run on Go versions from 1.12 to 1.20.


Branches
---------

Lastest finished features are committed to `main` branch as development branch.
`main` branch may not be stable, but it must be compilable and the new feature must be finished.

Release branches:
  + `v0`, development releases of each chapter in book Writing An Interpreter in Go.
    The final work is released as `v1.0.0` and in `v1` branch.
  + `v1`, development releases of each chapter in book Writing An Compiler in Go.
    The final work is released as `v1.10.1` and in `v1` branch.
  + `v2`, final work of monkey-lang, followed the steps in books.
    Add full script support, and ONLY bug fixes will be committed to this branch.
  + `v3`, new features or implements for monkey-lang, BUT fully compatible to offical implement.


Unknown bugs
-------------

Monkey-lang has bugs, by original design. Bugs in [BUGS](BUGS.md) are reproduced in offical final work in both WAIIG and WACIG.
