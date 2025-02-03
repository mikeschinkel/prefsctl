package gitutil

import (
	"errors"
)

var (
	ErrInvalidGitRepoURL          = errors.New("invalid Git repo URL")
	ErrInvalidGitRepoPath         = errors.New("invalid Git repo path")
	ErrInvalidGitRepoParentDir    = errors.New("invalid Git repo parent dir")
	ErrUnableToDeleteGitRepoDir   = errors.New("unable to delete Git repo dir")
	ErrFailedToEnsureGitRepoClone = errors.New("failed to ensure clone of Git repo")
)
