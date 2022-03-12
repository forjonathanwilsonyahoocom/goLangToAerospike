package main

import (
	"errors"
	"testing"
)

func TestUserAccess(t *testing.T) {
	mockError := errors.New("uh oh")
	subtests := []struct {
		name         string
		user, key    string
		areoSpikeGet func(string, string) (string, error)
		expectedErr  error
	}{
		{
			name: "happy path",
			user: "u",
			key:  "p",
			areoSpikeGet: func(s string, s2 string) (rs string, err error) {
				return s, nil
			},
		},
		{
			name: "error from areoSpikeGet",
			areoSpikeGet: func(s string, s2 string) (rs string, err error) {
				return s, mockError
			},
			//	expectedErr: mockError,
		},
	}
	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			_, err := getUser(subtest.user, subtest.key)
			if !errors.Is(err, subtest.expectedErr) {
				t.Errorf("expected error (%v), got error (%v)", subtest.expectedErr, err)
			}
		})
	}
}
