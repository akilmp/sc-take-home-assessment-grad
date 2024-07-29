package folders

import (
	"github.com/gofrs/uuid"
)

// FetchFolderRequestWithPagination is a struct that represents a request to fetch folders with pagination.
// It contains the organization ID, the number of folders to fetch (Limit), and a token to fetch the next set of folders.
type FetchFolderRequestWithPagination struct {
	OrgID uuid.UUID
	Limit int
	Token string
}

// FetchFolderResponseWithPagination is a struct that represents a response from fetching folders with pagination.
// It contains a slice of Folder pointers and a token to fetch the next set of folders.
type FetchFolderResponseWithPagination struct {
	Folders []*Folder
	Token   string
}

// GetAllFoldersWithPagination fetches folders for a given organization ID with pagination support.
func GetAllFoldersWithPagination(req *FetchFolderRequestWithPagination) (*FetchFolderResponseWithPagination, error) {
	// Fetch all folders for the given organization ID.
	r, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err // Fingers crossed
	}

	// Determine the starting index based on the provided token.
	// Let's find out where to start, using the token as a clue.
	startIndex := 0
	if req.Token != "" {
		for i, folder := range r {
			if folder.Id.String() == req.Token {
				startIndex = i + 1
				break // Found our starting point !!!
			}
		}
	}

	// Calculate the ending index based on the limit.
	endIndex := startIndex + req.Limit
	if endIndex > len(r) {
		endIndex = len(r) // Adjust if we've reached the end of the list.
	}

	// Create a slice to hold the folders for the current page.
	// Collecting our current batch of folders.
	var fp []*Folder
	for i := startIndex; i < endIndex; i++ {
		fp = append(fp, r[i]) // Adding folders like stamps to a collection.
	}

	// Determine the next token, which will be the ID of the folder after the current page.
	// Prepping the next clue for our ongoing scavenger hunt.
	var nextToken string
	if endIndex < len(r) {
		nextToken = r[endIndex].Id.String() // Setting up the next token if there are more folders.
	}

	// Return the response with the current page of folders and the next token.
	return &FetchFolderResponseWithPagination{
		Folders: fp,
		Token:   nextToken,
	}, nil
}
