package certs

import (
	"strconv"
	"testing"
)

func Test_NewFilesCluster(t *testing.T) {
	testCases := []struct {
		name  string
		files []File
	}{
		{
			name:  "case 0",
			files: NewFilesAPI(TLS{}),
		},
		{
			name:  "case 1",
			files: NewFilesCalicoEtcdClient(TLS{}),
		},
		{
			name:  "case 2",
			files: NewFilesEtcd(TLS{}),
		},
		{
			name:  "case 3",
			files: NewFilesServiceAccount(TLS{}),
		},
		{
			name:  "case 4",
			files: NewFilesWorker(TLS{}),
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// Make sure all paths are unique.
			{
				set := map[string]struct{}{}

				for _, f := range tc.files {
					set[f.AbsolutePath] = struct{}{}
				}
				if len(set) != len(tc.files) {
					t.Errorf("expected len = %d, got %d", len(set), len(tc.files))
				}
			}

			// Make sure all paths are absolute.
			{
				for _, f := range tc.files {
					if f.AbsolutePath[0] != '/' {
						t.Errorf("expected to start with %#q, got %#q", "/", f.AbsolutePath)
					}
				}
			}
		})
	}
}
