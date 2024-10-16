package folder

import "github.com/gofrs/uuid"

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}


func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
    folders := f.GetFoldersByOrgID(orgID)
    resp := []Folder{}

    // Find the full path to name
    parentPath := ""
    for _, folder := range folders {
        if folder.Name == name {
            parentPath = folder.Paths
            break
        }
    }

    // If the folder is not found, the empty array is returned.
    if parentPath == "" {
        return resp
    }

    // Iterate over all folders and look for subfolders prefixed with parentPath.
    prefix := parentPath + "."

    for _, folder := range folders {
        if isBelong(folder.Paths, prefix) {
            resp = append(resp, folder)
        }
    }
    return resp
}

// isBelong checks if the path starts with the specified prefix
func isBelong(Paths, prefix string) bool {
    return len(Paths) > len(prefix) && Paths[:len(prefix)] == prefix
}