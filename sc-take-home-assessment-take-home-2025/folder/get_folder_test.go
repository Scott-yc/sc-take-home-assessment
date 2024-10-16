package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

// Test_folder_GetAllChildFolders tests the GetAllChildFolders function
func Test_folder_GetAllChildFolders(t *testing.T) {
	// Creating simulated data for orgID
	org1 := uuid.Must(uuid.NewV4())
	org2 := uuid.Must(uuid.NewV4())

	// Analog folder data
	folders := []folder.Folder{
		{Name: "alpha", Paths: "alpha", OrgId: org1},
		{Name: "bravo", Paths: "alpha.bravo", OrgId: org1},
		{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: org1},
		{Name: "delta", Paths: "alpha.delta", OrgId: org1},
		{Name: "echo", Paths: "echo", OrgId: org1},
		{Name: "foxtrot", Paths: "foxtrot", OrgId: org2},
	}

	// Defining Test Cases
	tests := []struct {
		name    string
		orgID   uuid.UUID
		folder  string
		want    []folder.Folder
	}{
		{
			name:   "Get child folders for alpha",
			orgID:  org1,
			folder: "alpha",
			want: []folder.Folder{
				{Name: "bravo", Paths: "alpha.bravo", OrgId: org1},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: org1},
				{Name: "delta", Paths: "alpha.delta", OrgId: org1},
			},
		},
		{
			name:   "Get child folders for bravo",
			orgID:  org1,
			folder: "bravo",
			want: []folder.Folder{
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: org1},
			},
		},
		{
			name:   "No child folders for echo",
			orgID:  org1,
			folder: "charlie",
			want:   []folder.Folder{}, //empty
		},
		{
			name:   "No child folders for echo",
			orgID:  org1,
			folder: "echo",
			want:   []folder.Folder{}, //empty
		},
		{
			name:   "No child folders for org2",
			orgID:  org2,
			folder: "foxtrot",
			want:   []folder.Folder{}, //empty
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the driver object and call GetAllChildFolders.
			f := folder.NewDriver(folders)
			got := f.GetAllChildFolders(tt.orgID, tt.folder)

			// Checking that the results are as expected
			if len(got) != len(tt.want) {
				t.Errorf("expected %d child folders, got %d", len(tt.want), len(got))
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("expected folder: %+v, got: %+v", tt.want[i], got[i])
				}
			}
		})
	}
}














