package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFoldersWithPagination(t *testing.T) {
	// Set up the organization ID for testing.
	orgID := uuid.FromStringOrNil(folders.DefaultOrgID)
	// Create a request to fetch the first page of folders.
	req := &folders.FetchFolderRequestWithPagination{
		OrgID: orgID,
		Limit: 5,
		Token: "",
	}

	// Fetch the first page of folders.
	res, err := folders.GetAllFoldersWithPagination(req)
	assert.NoError(t, err)        // No errors allowed. Fingers Crossed
	assert.NotNil(t, res)         // We better have some results.
	assert.Len(t, res.Folders, 5) // I asked for 5 folders, you give me 5.

	// Check that all folders belong to the correct organization.
	for _, folder := range res.Folders {
		assert.Equal(t, orgID, folder.OrgId)
	}

	// Use the token from the first response to fetch the second page.
	req.Token = res.Token
	res, err = folders.GetAllFoldersWithPagination(req)
	assert.NoError(t, err)        // Still no errors allowed! Pretty Please
	assert.NotNil(t, res)         // We need our next batch of folders.
	assert.Len(t, res.Folders, 5) // Again, we should get 5 folders.

	// Check that all folders in the second batch also belong to the correct organization.
	for _, folder := range res.Folders {
		assert.Equal(t, orgID, folder.OrgId)
	}
}
