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
	"github.com/pkg/errors"
	"go.temporal.io/sdk/temporal"
	"net/url"
	"strings"
)

// VerifyResourceExists verifies that the resource exists.
// This is a Temporal activity.
func (w *Worker) VerifyResourceExists(uri string) error {
	canRead, err := w.storage.CanReadURI(uri)
	if err != nil {
		return errors.Wrap(err, "failed to check if resource URI can be read")
	}

	if !canRead {
		return temporal.NewNonRetryableApplicationError(
			"cannot read resource URI",
			"cannotReadResourceURI",
			errors.New("client submitted a resource URI that cannot be read by server"),
		)
	}

	// Get object name from URI.
	// Check if object exists.
	// If not, return error.
	parsed, err := url.Parse(uri)
	if err != nil {
		return temporal.NewNonRetryableApplicationError(
			"could not parse resource URI",
			"couldNotParseResourceURI",
			errors.Wrap(err, "failed to parse resource URI"),
		)
	}

	split := strings.SplitN(parsed.Path, "/", 2)
	if len(split) < 2 {
		return temporal.NewNonRetryableApplicationError(
			"invalid resource URI",
			"invalidResourceURI",
			errors.New("client submitted an invalid resource URI"),
		)
	}

	object := split[1]
	exists, err := w.storage.Exists(object)
	if err != nil {
		return errors.Wrap(err, "failed to check if resource exists")
	}

	if !exists {
		// Since the client can trigger this activity before uploading the resource,
		// we should not return a non-retryable error.
		// The parent workflow should handle the retry arrangements up to 2 hours
		// per the spec.
		return errors.New("resource does not exist")
	}

	return nil
}
