package mothership_worker_server

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWorker_VerifyResourceExists(t *testing.T) {
	require.Nil(t, testW.VerifyResourceExists("memory://efi-rpm-macros-3-3.el8.src.rpm"))
}

func TestWorker_VerifyResourceExists_NotFound(t *testing.T) {
	err := testW.VerifyResourceExists("memory://not-found.rpm")
	require.NotNil(t, err)
	require.Equal(t, err.Error(), "resource does not exist")
}

func TestWorker_VerifyResourceExists_CannotRead(t *testing.T) {
	err := testW.VerifyResourceExists("bad-protocol://not-found.rpm")
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "client submitted a resource URI that cannot be read by server")
}
