package applemdm

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mikeschinkel/prefsctl/stdlibex"
)

func mustClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		slog.Error(err.Error())
	}
}

func EnsureGitRepo(repoURL string, parentDir string) (err error) {
	var exists bool
	var cmd *exec.Cmd
	var repoDir string
	var output []byte

	err = os.MkdirAll(parentDir, os.ModePerm)
	if err != nil {
		goto end
	}
	repoDir = filepath.Join(parentDir, filepath.Base(repoURL))
	exists, err = stdlibex.DirExists(repoDir)
	if err != nil {
		goto end
	}
	if !exists {
		cmd = exec.Command("git", "clone", repoURL, repoDir)
	} else {
		cmd = exec.Command("git", "-C", repoDir, "pull")
	}
	output, err = cmd.CombinedOutput()
	if err != nil {
		err = errors.Join(
			ErrFailedToEnsureGitRepoClone,
			fmt.Errorf("%s=%s", RepoURLLogArg, repoURL),
			fmt.Errorf("%s=%s", CloneDirLogArg, GitRepoParentDir),
			fmt.Errorf("%s=%s", CmdOutputLogArg, output),
			err,
		)
		goto end
	}
end:
	return err
}
