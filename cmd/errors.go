package main

import (
	"errors"
)

// ErrCommand is the error used when wrong command instructions are provided.
var ErrCommand = errors.New("expected 'validate' subcommand")

// ErrFileNotFound is the error used when file is not found.
var ErrFileNotFound = errors.New("file is not found or does not exist")
