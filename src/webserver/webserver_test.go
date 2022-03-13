package main

import (
	"errors"
	"testing"

	aero "github.com/aerospike/aerospike-client-go"
)

var client, err = aero.NewClient("192.168.88.190", 3000)

//this ip will be wrong...
//the aerospike docker has to be up for this test to work
func TestUserAccess(t *testing.T) {
	subtests := []struct {
		name            string
		user, key       string
		aerospikeClient *aero.Client
		expectedErr     error
	}{
		{
			name:            "happy path",
			user:            "11",
			key:             "42",
			aerospikeClient: client,
		},
		{
			name:            "error from areoSpikeGet",
			aerospikeClient: client,
			expectedErr:     errors.New("404"),
		},
	}
	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			_, err := getUser(subtest.user, subtest.key, subtest.aerospikeClient)
			if !errors.Is(err, subtest.expectedErr) {
				if err.Error() != err.Error() {
					t.Errorf("expected error (%v), got error (%v)", subtest.expectedErr, err)
				}
			}
		})
	}
}
