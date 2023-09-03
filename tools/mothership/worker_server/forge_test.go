package mothership_worker_server

import (
	"fmt"
	transport_http "github.com/go-git/go-git/v5/plumbing/transport/http"
	"go.resf.org/peridot/tools/mothership/worker_server/forge"
	"time"
)

type inMemoryForge struct {
	localTempDir  string
	repos         map[string]bool
	remoteBaseURL string
}

func (f *inMemoryForge) GetAuthenticator() (*forge.Authenticator, error) {
	return &forge.Authenticator{
		AuthMethod: &transport_http.BasicAuth{
			Username: "user",
			Password: "pass",
		},
		AuthorName:  "Test User",
		AuthorEmail: "test@resf.org",
		Expires:     time.Now().Add(time.Hour),
	}, nil
}

func (f *inMemoryForge) GetRemote(repo string) string {
	return fmt.Sprintf("file://%s/%s", f.localTempDir, repo)
}

func (f *inMemoryForge) GetCommitViewerURL(repo string, commit string) string {
	return f.remoteBaseURL + "/" + repo + "/commit/" + commit
}

func (f *inMemoryForge) EnsureRepositoryExists(auth *forge.Authenticator, repo string) error {
	f.repos[repo] = true
	return nil
}
