package gitutil

import (
	"errors"
)

var (
	ErrFailedToEnsureGitRepoClone = errors.New("failed to ensure clone of Git repo")
)
