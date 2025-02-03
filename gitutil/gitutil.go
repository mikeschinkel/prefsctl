package gitutil

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/stdlibex"
)

// MinParentDirLen is the minimum length a parent of a Git repo URL can be
// This is a variable so can be changed if needed.
// /Users/xyz
// 1234567890
var MinParentDirLen = 10

// MinRepoURLLen is the minimum length a Git repo URL can be
// This is a variable so can be changed if needed.
// https://github.com/x/x
// 1234567890123456789012
var MinRepoURLLen = 22

// MinRepoPathLen is the minimum length a Git repo path can be
// This is a variable so can be changed if needed.
// https://github.com/x/x => x
//
//	1
var MinRepoPathLen = 1

func EnsureGitRepo(repoURL string, parentDir string) (string, error) {
	return ensureGitRepo(repoURL, parentDir, true)
}

type pathLenArgs struct {
	minLen int
	path   string
	descr  string
	err    error
}

func chkPathLen(args pathLenArgs) (err error) {
	if len(args.path) >= args.minLen {
		goto end
	}
	err = errutil.AnnotateErr(args.err,
		"%s too short (must be >=%d): len=%d",
		args.descr,
		args.minLen,
		len(args.path),
	)
end:
	return err
}

func ensureGitRepo(repoURL string, parentDir string, retry bool) (repoDir string, err error) {
	var exists bool
	var cmd *exec.Cmd
	var repoPath string
	var output []byte
	var paths []pathLenArgs

	err = os.MkdirAll(parentDir, os.ModePerm)
	if err != nil {
		goto end
	}
	repoPath = filepath.Base(repoURL)
	repoDir = filepath.Join(parentDir, repoPath)
	exists, err = stdlibex.DirExists(repoDir)
	if err != nil {
		goto end
	}
	paths = []pathLenArgs{
		{MinRepoURLLen, repoURL, "Git repo URL", ErrInvalidGitRepoURL},
		{MinParentDirLen, parentDir, "Git parent dir", ErrInvalidGitRepoParentDir},
		{MinRepoPathLen, repoPath, "Git repo path dir", ErrInvalidGitRepoPath},
	}
	for _, path := range paths {
		if err = chkPathLen(path); err != nil {
			goto end
		}
	}
	if !retry {
		err = os.RemoveAll(repoDir)
		if err != nil {
			err = errors.Join(
				ErrUnableToDeleteGitRepoDir,
			)
			goto end
		}
		exists = false
	}
	if !exists {
		cmd = exec.Command("git", "clone", repoURL, repoDir)
	} else {
		cmd = exec.Command("git", "-C", repoDir, "pull")
	}
	output, err = cmd.CombinedOutput()
	if err != nil {
		if retry {
			repoDir, err = ensureGitRepo(repoURL, parentDir, false)
			goto end
		}
		err = errors.Join(
			ErrFailedToEnsureGitRepoClone,
			fmt.Errorf("%s=%s", RepoURLLogArg, repoURL),
			fmt.Errorf("%s=%s", CloneDirLogArg, repoDir),
			fmt.Errorf("%s=%s", CmdOutputLogArg, output),
			err,
		)
		goto end
	}
end:
	return repoDir, err
}
