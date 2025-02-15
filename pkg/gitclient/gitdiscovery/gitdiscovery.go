package gitdiscovery

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/shuttlerock/jx-helpers/v3/pkg/gitclient"
	"github.com/shuttlerock/jx-helpers/v3/pkg/gitclient/gitconfig"
	"github.com/shuttlerock/jx-helpers/v3/pkg/gitclient/giturl"
)

// FindGitURLFromDir tries to find the git clone URL from the given directory
func FindGitURLFromDir(dir string, preferUpstream bool) (string, error) {
	_, gitConfDir, err := gitclient.FindGitConfigDir(dir)
	if err != nil {
		return "", errors.Wrapf(err, "there was a problem obtaining the git config dir of directory %s", dir)
	}

	if gitConfDir == "" {
		// lets use an env var instead
		gitURL := os.Getenv("SOURCE_URL")
		if gitURL != "" {
			return gitURL, nil
		}
		return "", fmt.Errorf("no .git directory could be found from dir %s", dir)
	}
	return gitconfig.DiscoverUpstreamGitURL(gitConfDir, preferUpstream)
}

// FindGitInfoFromDir finds the git info from the given dir
func FindGitInfoFromDir(dir string) (*giturl.GitRepository, error) {
	gitURL, err := FindGitURLFromDir(dir, false)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to discover the git URL")
	}
	if gitURL == "" {
		return nil, errors.Errorf("no git URL could be discovered")
	}

	return giturl.ParseGitURL(gitURL)
}
