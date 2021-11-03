// Copyright © 2020, 2021 Attestant Limited.
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

package http_test

import (
	"context"
	"os"
	"testing"
	"time"

	client "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/http"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/stretchr/testify/require"
)

func TestBeaconCommittees(t *testing.T) {
	tests := []struct {
		name    string
		stateID string
	}{
		{
			name:    "Good",
			stateID: "head",
		},
	}

	service, err := http.New(context.Background(),
		http.WithTimeout(timeout),
		http.WithAddress(os.Getenv("HTTP_ADDRESS")),
	)
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			beaconCommittees, err := service.(client.BeaconCommitteesProvider).BeaconCommittees(context.Background(), test.stateID)
			require.NoError(t, err)
			require.NotNil(t, beaconCommittees)
		})
	}
}

func TestBeaconCommitteesAtEpoch(t *testing.T) {
	ctx := context.Background()

	service, err := http.New(ctx,
		http.WithTimeout(timeout),
		http.WithAddress(os.Getenv("HTTP_ADDRESS")),
	)
	require.NoError(t, err)

	// Needed to fetch current epoch.
	genesis, err := service.(client.GenesisProvider).Genesis(context.Background())
	require.NoError(t, err)
	slotDuration, err := service.(client.SlotDurationProvider).SlotDuration(context.Background())
	require.NoError(t, err)
	slotsPerEpoch, err := service.(client.SlotsPerEpochProvider).SlotsPerEpoch(context.Background())
	require.NoError(t, err)

	tests := []struct {
		name    string
		stateID string
		epoch   int64 // -1 for current
	}{
		{
			name:    "Good",
			stateID: "head",
			epoch:   -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var epoch phase0.Epoch
			if test.epoch == -1 {
				epoch = phase0.Epoch(uint64(time.Since(genesis.GenesisTime).Seconds()) / (uint64(slotDuration.Seconds()) * slotsPerEpoch))
			} else {
				epoch = phase0.Epoch(test.epoch)
			}
			beaconCommittees, err := service.(client.BeaconCommitteesProvider).BeaconCommitteesAtEpoch(context.Background(), test.stateID, epoch)
			require.NoError(t, err)
			require.NotNil(t, beaconCommittees)
		})
	}
}
