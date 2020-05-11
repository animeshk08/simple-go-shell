package main

import (
	"os"
	"testing"
)

func TestExecInput(t *testing.T) {
	tests := []struct {
		name        string
		givenInput  string
		expectedDir string // Test `cd`
		expectedErr error
	}{
		{
			name:        "pasing: lk",
			givenInput:  "ls",
			expectedErr: nil,
		},
		{
			name:        "passing: ls -l -a",
			givenInput:  "ls -l -a",
			expectedErr: nil,
		},
		{
			name:        "passing: cd",
			givenInput:  "cd",
			expectedErr: ErrNoPath,
		},
		{
			name:        "passing; cd /",
			givenInput:  "cd /",
			expectedErr: nil,
			expectedDir: "/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// test command
			if err := execInput(tt.givenInput); err != tt.expectedErr {
				t.Errorf("execInput() error = %v, wantErr %v", err, tt.expectedErr)
			}

			// test for changed directory only
			if tt.expectedDir != "" {
				curDir, err := os.Getwd()
				if err != nil {
					t.Errorf("Failed to get new directory: %v", err)
				}
				if tt.expectedDir != curDir {
					t.Errorf("Failed to change to desired directory.")
				}
			}
		})
	}
}
