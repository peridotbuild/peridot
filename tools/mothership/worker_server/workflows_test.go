// Copyright 2023 Peridot Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mothership_worker_server

import (
	"database/sql"
	"errors"
	"github.com/stretchr/testify/mock"
	base "go.resf.org/peridot/base/go"
	mothership_db "go.resf.org/peridot/tools/mothership/db"
	mothershippb "go.resf.org/peridot/tools/mothership/pb"
	"time"
)

func (s *UnitTestSuite) TestProcessRPMWorkflow_FullSuccess1() {
	s.env.OnActivity(testW.VerifyResourceExists, "memory://efi-rpm-macros-3-3.el8.src.rpm").Return(nil)
	s.env.OnActivity(testW.SetWorkerLastCheckinTime, mock.Anything).Return(nil)

	entry := (&mothership_db.Entry{
		Name:           base.NameGen("entries"),
		CreateTime:     time.Now(),
		OSRelease:      "Rocky Linux release 8.8 (Green Obsidian)",
		Sha256Sum:      "518a9418fec1deaeb4c636615d8d81fb60146883c431ea15ab1127893d075d28",
		RepositoryName: "BaseOS",
		WorkerID: sql.NullString{
			String: "test-worker",
			Valid:  true,
		},
		State: mothershippb.Entry_ARCHIVING,
	}).ToPB()
	s.env.OnActivity(testW.CreateEntry, mock.Anything).Return(entry, nil)

	entry.EntryId = "efi-rpm-macros-3-3.el8.src"
	entry.Sha256Sum = "518a9418fec1deaeb4c636615d8d81fb60146883c431ea15ab1127893d075d28"
	s.env.OnActivity(testW.SetEntryIDFromRPM, entry.Name, "memory://efi-rpm-macros-3-3.el8.src.rpm", entry.Sha256Sum).Return(entry, nil)

	importRpmRes := &mothershippb.ImportRPMResponse{
		CommitHash:   "4e1243bd22c66e76c2ba9eddc1f91394e57f9f83",
		CommitUri:    testW.forge.GetCommitViewerURL("efi-rpm-macros", "4e1243bd22c66e76c2ba9eddc1f91394e57f9f83"),
		CommitBranch: "el-8.8",
		CommitTag:    "imports/el-8.8/efi-rpm-macros-3-3.el8",
		Nevra:        "efi-rpm-macros-0:3-3.el8.aarch64",
		Pkg:          "efi-rpm-macros",
	}
	s.env.OnActivity(testW.ImportRPM, "memory://efi-rpm-macros-3-3.el8.src.rpm", entry.Sha256Sum, entry.OsRelease).Return(importRpmRes, nil)

	s.env.OnActivity(testW.SetEntryState, entry.Name, mothershippb.Entry_ARCHIVED, importRpmRes).Return(entry, nil)

	args := &mothershippb.ProcessRPMArgs{
		Request: &mothershippb.ProcessRPMRequest{
			RpmUri:     "memory://efi-rpm-macros-3-3.el8.src.rpm",
			OsRelease:  "Rocky Linux release 8.8 (Green Obsidian)",
			Checksum:   entry.Sha256Sum,
			Repository: "BaseOS",
		},
		InternalRequest: &mothershippb.ProcessRPMInternalRequest{
			WorkerId: "test-worker",
		},
	}
	s.env.ExecuteWorkflow(ProcessRPMWorkflow, args)
	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())

	var res mothershippb.ProcessRPMResponse
	s.NoError(s.env.GetWorkflowResult(&res))
	s.Equal(entry.Name, res.Entry.Name)
	s.Equal(entry.EntryId, res.Entry.EntryId)
}

func (s *UnitTestSuite) TestProcessRPMWorkflow_OnHold_Cancel() {
	s.env.OnActivity(testW.VerifyResourceExists, "memory://efi-rpm-macros-3-3.el8.src.rpm").Return(nil)
	s.env.OnActivity(testW.SetWorkerLastCheckinTime, mock.Anything).Return(nil)

	entry := (&mothership_db.Entry{
		Name:           base.NameGen("entries"),
		CreateTime:     time.Now(),
		OSRelease:      "Rocky Linux release 8.8 (Green Obsidian)",
		Sha256Sum:      "518a9418fec1deaeb4c636615d8d81fb60146883c431ea15ab1127893d075d28",
		RepositoryName: "BaseOS",
		WorkerID: sql.NullString{
			String: "test-worker",
			Valid:  true,
		},
		State: mothershippb.Entry_ARCHIVING,
	}).ToPB()
	s.env.OnActivity(testW.CreateEntry, mock.Anything).Return(entry, nil)

	entry.EntryId = "efi-rpm-macros-3-3.el8.src"
	entry.Sha256Sum = "518a9418fec1deaeb4c636615d8d81fb60146883c431ea15ab1127893d075d28"
	s.env.OnActivity(testW.SetEntryIDFromRPM, entry.Name, "memory://efi-rpm-macros-3-3.el8.src.rpm", entry.Sha256Sum).Return(entry, nil)

	importErr := errors.New("import error")
	s.env.OnActivity(testW.ImportRPM, "memory://efi-rpm-macros-3-3.el8.src.rpm", entry.Sha256Sum, entry.OsRelease).Return(nil, importErr)

	s.env.OnActivity(testW.SetEntryState, entry.Name, mothershippb.Entry_ON_HOLD, mock.Anything).Return(entry, nil)
	s.env.OnActivity(testW.SetEntryState, entry.Name, mothershippb.Entry_CANCELLED, mock.Anything).Return(entry, nil)

	s.env.RegisterDelayedCallback(func() {
		s.env.CancelWorkflow()
	}, 500*time.Millisecond)

	args := &mothershippb.ProcessRPMArgs{
		Request: &mothershippb.ProcessRPMRequest{
			RpmUri:     "memory://efi-rpm-macros-3-3.el8.src.rpm",
			OsRelease:  "Rocky Linux release 8.8 (Green Obsidian)",
			Checksum:   entry.Sha256Sum,
			Repository: "BaseOS",
		},
		InternalRequest: &mothershippb.ProcessRPMInternalRequest{
			WorkerId: "test-worker",
		},
	}
	s.env.ExecuteWorkflow(ProcessRPMWorkflow, args)

	s.True(s.env.IsWorkflowCompleted())
	s.ErrorContains(s.env.GetWorkflowError(), "canceled")
}

func (s *UnitTestSuite) TestProcessRPMWorkflow_OnHold_Success() {
	s.env.OnActivity(testW.VerifyResourceExists, "memory://efi-rpm-macros-3-3.el8.src.rpm").Return(nil)
	s.env.OnActivity(testW.SetWorkerLastCheckinTime, mock.Anything).Return(nil)

	entry := (&mothership_db.Entry{
		Name:           base.NameGen("entries"),
		CreateTime:     time.Now(),
		OSRelease:      "Rocky Linux release 8.8 (Green Obsidian)",
		Sha256Sum:      "518a9418fec1deaeb4c636615d8d81fb60146883c431ea15ab1127893d075d28",
		RepositoryName: "BaseOS",
		WorkerID: sql.NullString{
			String: "test-worker",
			Valid:  true,
		},
		State: mothershippb.Entry_ARCHIVING,
	}).ToPB()
	s.env.OnActivity(testW.CreateEntry, mock.Anything).Return(entry, nil)

	entry.EntryId = "efi-rpm-macros-3-3.el8.src"
	entry.Sha256Sum = "518a9418fec1deaeb4c636615d8d81fb60146883c431ea15ab1127893d075d28"
	s.env.OnActivity(testW.SetEntryIDFromRPM, entry.Name, "memory://efi-rpm-macros-3-3.el8.src.rpm", entry.Sha256Sum).Return(&*entry, nil)

	importErr := errors.New("import error")
	importRpmRes := &mothershippb.ImportRPMResponse{
		CommitHash:   "4e1243bd22c66e76c2ba9eddc1f91394e57f9f83",
		CommitUri:    testW.forge.GetCommitViewerURL("efi-rpm-macros", "4e1243bd22c66e76c2ba9eddc1f91394e57f9f83"),
		CommitBranch: "el-8.8",
		CommitTag:    "imports/el-8.8/efi-rpm-macros-3-3.el8",
		Nevra:        "efi-rpm-macros-0:3-3.el8.aarch64",
		Pkg:          "efi-rpm-macros",
	}
	shouldErrImport := true
	s.env.OnActivity(testW.ImportRPM, "memory://efi-rpm-macros-3-3.el8.src.rpm", entry.Sha256Sum, entry.OsRelease).
		Return(func(uri string, checksum string, osRelease string) (*mothershippb.ImportRPMResponse, error) {
			if shouldErrImport {
				return nil, importErr
			}
			return importRpmRes, nil
		})

	entry.State = mothershippb.Entry_ON_HOLD
	s.env.OnActivity(testW.SetEntryState, entry.Name, mothershippb.Entry_ON_HOLD, mock.Anything).Return(&*entry, nil)

	entry.State = mothershippb.Entry_ARCHIVED
	s.env.OnActivity(testW.SetEntryState, entry.Name, mothershippb.Entry_ARCHIVING, mock.Anything).Return(&*entry, nil)

	entry.State = mothershippb.Entry_ARCHIVED
	s.env.OnActivity(testW.SetEntryState, entry.Name, mothershippb.Entry_ARCHIVED, importRpmRes).Return(&*entry, nil)

	s.env.RegisterDelayedCallback(func() {
		shouldErrImport = false
		s.env.SignalWorkflow("rescue", true)
	}, 500*time.Millisecond)

	args := &mothershippb.ProcessRPMArgs{
		Request: &mothershippb.ProcessRPMRequest{
			RpmUri:     "memory://efi-rpm-macros-3-3.el8.src.rpm",
			OsRelease:  "Rocky Linux release 8.8 (Green Obsidian)",
			Checksum:   entry.Sha256Sum,
			Repository: "BaseOS",
		},
		InternalRequest: &mothershippb.ProcessRPMInternalRequest{
			WorkerId: "test-worker",
		},
	}
	s.env.ExecuteWorkflow(ProcessRPMWorkflow, args)

	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())

	var res mothershippb.ProcessRPMResponse
	s.NoError(s.env.GetWorkflowResult(&res))
	s.Equal(entry.Name, res.Entry.Name)
	s.Equal(entry.EntryId, res.Entry.EntryId)
}

func (s *UnitTestSuite) TestProcessRPMWorkflow_OnHold_Error() {
	s.env.OnActivity(testW.VerifyResourceExists, "memory://efi-rpm-macros-3-3.el8.src.rpm").Return(nil)
	s.env.OnActivity(testW.SetWorkerLastCheckinTime, mock.Anything).Return(nil)

	entry := (&mothership_db.Entry{
		Name:           base.NameGen("entries"),
		CreateTime:     time.Now(),
		OSRelease:      "Rocky Linux release 8.8 (Green Obsidian)",
		Sha256Sum:      "518a9418fec1deaeb4c636615d8d81fb60146883c431ea15ab1127893d075d28",
		RepositoryName: "BaseOS",
		WorkerID: sql.NullString{
			String: "test-worker",
			Valid:  true,
		},
		State: mothershippb.Entry_ARCHIVING,
	}).ToPB()
	s.env.OnActivity(testW.CreateEntry, mock.Anything).Return(entry, nil)

	entry.EntryId = "efi-rpm-macros-3-3.el8.src"
	entry.Sha256Sum = "518a9418fec1deaeb4c636615d8d81fb60146883c431ea15ab1127893d075d28"
	s.env.OnActivity(testW.SetEntryIDFromRPM, entry.Name, "memory://efi-rpm-macros-3-3.el8.src.rpm", entry.Sha256Sum).Return(entry, nil)

	importErr := errors.New("import error")
	s.env.OnActivity(testW.ImportRPM, "memory://efi-rpm-macros-3-3.el8.src.rpm", entry.Sha256Sum, entry.OsRelease).Return(nil, importErr)

	s.env.OnActivity(testW.SetEntryState, entry.Name, mothershippb.Entry_ON_HOLD, mock.Anything).Return(entry, nil)

	args := &mothershippb.ProcessRPMArgs{
		Request: &mothershippb.ProcessRPMRequest{
			RpmUri:     "memory://efi-rpm-macros-3-3.el8.src.rpm",
			OsRelease:  "Rocky Linux release 8.8 (Green Obsidian)",
			Checksum:   entry.Sha256Sum,
			Repository: "BaseOS",
		},
		InternalRequest: &mothershippb.ProcessRPMInternalRequest{
			WorkerId: "test-worker",
		},
	}
	s.env.ExecuteWorkflow(ProcessRPMWorkflow, args)

	s.True(s.env.IsWorkflowCompleted())
	s.Error(s.env.GetWorkflowError())
}
