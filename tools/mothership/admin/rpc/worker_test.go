package mothershipadmin_rpc

import (
	"github.com/stretchr/testify/require"
	base "go.resf.org/peridot/base/go"
	mshipadminpb "go.resf.org/peridot/tools/mothership/admin/pb"
	mothership_db "go.resf.org/peridot/tools/mothership/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestGetWorker_Empty(t *testing.T) {
	require.Nil(t, base.Q[mothership_db.Worker](s.db).Delete())
	worker, err := s.GetWorker(testContext(), &mshipadminpb.GetWorkerRequest{})
	require.NotNil(t, err)
	require.Nil(t, worker)
	expectedErr := status.Error(codes.NotFound, "worker not found")
	require.Equal(t, expectedErr.Error(), err.Error())
}

func TestGetWorker_One(t *testing.T) {
	require.Nil(t, base.Q[mothership_db.Worker](s.db).Delete())
	require.Nil(t, base.Q[mothership_db.Worker](s.db).Create(&mothership_db.Worker{
		Name:      "test",
		WorkerID:  "test-id",
		ApiSecret: "secret",
	}))
	defer func() {
		require.Nil(t, base.Q[mothership_db.Worker](s.db).Delete())
	}()

	worker, err := s.GetWorker(testContext(), &mshipadminpb.GetWorkerRequest{
		Name: "test",
	})
	require.Nil(t, err)
	require.Equal(t, "test", worker.Name)
	require.Equal(t, "test-id", worker.WorkerId)
	require.Empty(t, worker.ApiSecret)
}
