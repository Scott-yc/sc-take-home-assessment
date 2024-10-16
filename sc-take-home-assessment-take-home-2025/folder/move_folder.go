package folder

import "errors"

// MoveFolder moves a folder from one parent folder to another.
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// // Get all folders
	folders := f.folders
	//Create two empty folders, make sure they start out empty, and make substitutions afterward
	var srcFolder, dstFolder Folder
	srcIndex, dstIndex := -1, -1

	// Iterate to find the source and target folders
	for i := 0; i < len(folders); i++ {
		if folders[i].Name == name {
			srcFolder = folders[i]
			srcIndex = i
		}
		if folders[i].Name == dst {
			dstFolder = folders[i]
			dstIndex = i
		}
	}

	// Error handling: source folder does not exist
	if srcIndex == -1 {
		return nil, errors.New("source folder does not exist")
	}

	// Error handling: target folder does not exist
	if dstIndex == -1 {
		return nil, errors.New("destination folder does not exist")
	}

	// Error handling: Cannot move to a folder in a different organization
	if srcFolder.OrgId != dstFolder.OrgId {
		return nil, errors.New("cannot move folder to a different organization")
	}

	// Error handling: can't move to subfolders
	if isSubclass(srcFolder.Paths, dstFolder.Paths) {
		return nil, errors.New("cannot move a folder to a child of itself")
	}

	// Error handling: can't move to itself
	if name == dst {
		return nil, errors.New("cannot move folder to itself")
	}

	// Generate new path prefix: destination folder path + “.” + source folder name
	newPrefix := dstFolder.Paths + "." + srcFolder.Name

	// Iterate through all folders and then update the paths of the source folder and its subfolders
	for i := 0; i < len(folders); i++ {
		// If the current folder path begins with the source folder path, it is a subfolder
		if isSubclass(folders[i].Paths, srcFolder.Paths) {
			if len(folders[i].Paths) >= len(srcFolder.Paths) {
				newPath := newPrefix + folders[i].Paths[len(srcFolder.Paths):]
				folders[i].Paths = newPath
			}

		}
	}

	return folders, nil
}

// isSubclass Checks if the destination folder path is a subclass of the source folder path.
func isSubclass(originPath, targetPath string) bool {
	// Make sure that targetPath is longer than originPath and that targetPath is prefixed with originPath.
	if len(targetPath) > len(originPath) && targetPath[:len(originPath)] == originPath {
		// Make sure originPath is immediately followed by “.”
		return targetPath[len(originPath)] == '.'
	}
	return false
}