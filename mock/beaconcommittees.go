// Copyright © 2020 Attestant Limited.
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

package mock

import (
	"context"

	"github.com/attestantio/go-eth2-client/api"
	apiv1 "github.com/attestantio/go-eth2-client/api/v1"
)

// BeaconCommittees fetches all beacon committees for the epoch at the given state.
func (*Service) BeaconCommittees(_ context.Context,
	_ *api.BeaconCommitteesOpts,
) (
	*api.Response[[]*apiv1.BeaconCommittee],
	error,
) {
	data := make([]*apiv1.BeaconCommittee, 5)
	for i := 0; i < 5; i++ {
		data[i] = &apiv1.BeaconCommittee{}
	}

	return &api.Response[[]*apiv1.BeaconCommittee]{
		Data:     data,
		Metadata: make(map[string]any),
	}, nil
}
