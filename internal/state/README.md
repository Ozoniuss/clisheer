State
-----

This package provides a solution of adding state to the application. Since this application's binary is meant to be used as a command-line tool to perform actions non-interactively (i.e. one command at a time), state can't really be stored in memory. However, in some cases, keeping track of some background information, e.g. the period you're working with, allows simplifying the executed commands. For this reason, a [lockedfile](https://pkg.go.dev/cmd/go/internal/lockedfile) (which is definitely so overkill but whatever) is provided to allow keeping track of relevant state.