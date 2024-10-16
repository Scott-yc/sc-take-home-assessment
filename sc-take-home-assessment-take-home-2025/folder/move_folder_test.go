package folder_test

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/georgechieng-sc/interns-2022/folder" 
)

func Test_folder_MoveFolder(t *testing.T) {
	
	folders := []folder.Folder{
		{
			Name:  "alpha",
			Paths: "alpha",
			OrgId: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
		},
		{
			Name:  "bravo",
			Paths: "alpha.bravo",
			OrgId: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
		},
		{
			Name:  "charlie",
			Paths: "alpha.bravo.charlie",
			OrgId: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
		},
		{
			Name:  "delta",
			Paths: "alpha.delta",
			OrgId: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
		},
		{
			Name:  "echo",
			Paths: "alpha.delta.echo",
			OrgId: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
		},
		{
			Name:  "foxtrot",
			Paths: "foxtrot",
			OrgId: uuid.Must(uuid.NewV4()), 
		},
		{
			Name:  "golf",
			Paths: "golf",
			OrgId: uuid.Must(uuid.FromString("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")),
		},
	}

	// test example
	tests := []struct {
		name          string
		srcFolderName string
		dstFolderName string
		expectedError string
	}{
		{
			name:          "Move bravo to delta",
			srcFolderName: "bravo",
			dstFolderName: "delta",
			expectedError: "",
		},
		{
			name:          "Move bravo to itself",
			srcFolderName: "bravo",
			dstFolderName: "bravo",
			expectedError: "cannot move folder to itself",
		},
		{
			name:          "Move bravo to charlie",
			srcFolderName: "bravo",
			dstFolderName: "charlie",
			expectedError: "cannot move a folder to a child of itself",
		},
		{
			name:          "Move bravo to foxtrot",
			srcFolderName: "bravo",
			dstFolderName: "foxtrot",
			expectedError: "cannot move folder to a different organization",
		},
	}

	// Iterating Test Cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize the driver and perform a move operation
			d := folder.NewDriver(folders)
			_, err := d.MoveFolder(tt.srcFolderName, tt.dstFolderName)

			// If the expected error is not null, check for a match
			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}