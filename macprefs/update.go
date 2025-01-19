package macprefs

import (
	"github.com/mikeschinkel/prefsctl/applemdm"
	"github.com/mikeschinkel/prefsctl/config"
)

const (
	profileManifestsRepoURL = applemdm.ProfileManifestsRepoURL
	gitRepoParentDir        = applemdm.GitRepoParentDir
)

type UpdateArgs struct {
	Filename Filename
}

func Update(ctx Context, cfg config.Config, ptr Printer, args UpdateArgs) (result Result) {
	if ptr == nil {
		ptr = StandardPrinter{}
	}
	err := applemdm.EnsureGitRepo(profileManifestsRepoURL, gitRepoParentDir)
	if err != nil {
		goto end
	}
end:
	if err != nil {
		result = Result{Err: err}
	}
	return result
}
