package folders

import (
	"github.com/gofrs/uuid"
)

// GetAllFolders fetches all folders for a given organization ID.
// It takes a FetchFolderRequest and returns a FetchFolderResponse or an error.
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// Get all folders for the specified organization ID.
	// Fetching folders like a digital librarian.
	r, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		// Return an error if something goes wrong. Fingers crossed
		return nil, err
	}

	// Prepare a slice to hold the folder pointers.
	var fp []*Folder
	// Loop through the fetched folders and add each one to the slice.
	for _, v := range r {
		fp = append(fp, v)
	}

	// Return the response with the folders. Mission accomplished!
	return &FetchFolderResponse{Folders: fp}, nil
}

// FetchAllFoldersByOrgID gets folders from the sample data that match the given organization ID.
// It returns a slice of folder pointers and an error.
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	// Get the sample data which has all the folders.
	// Imagine this as our treasure chest of folders.
	folders := GetSampleData()

	// Create a slice to store the matching folders.
	var resFolder []*Folder
	// Loop through the sample data.
	// We're on a quest to find the right folders.
	for _, folder := range folders {
		// Check if the folder's organization ID matches the one we're looking for.
		if folder.OrgId == orgID {
			// If it matches, add the folder to our result slice.
			// It's a match! Adding to our collection.
			resFolder = append(resFolder, folder)
		}
	}
	// Return the matching folders. Case closed!
	return resFolder, nil
}
