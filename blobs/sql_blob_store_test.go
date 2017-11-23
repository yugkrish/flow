package blobs

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	// needed to use sqlite in tests
	"github.com/fnproject/flow/persistence"
	_ "github.com/mattn/go-sqlite3"
)

func TestShouldInsertBlobAndGenerateId(t *testing.T) {
	store := givenEmptyBlobStore()

	data := []byte("Some data")
	blob, err := store.Create("graph", "test/type", bytes.NewReader(data))

	assert.NoError(t, err)
	require.NotNil(t, blob)
	assert.NotNil(t, blob.ID)
	assert.Equal(t, "test/type", blob.ContentType)
	assert.Equal(t, int64(len(data)), blob.Length)

}

func TestShouldRetrieveStoredBlob(t *testing.T) {
	store := givenEmptyBlobStore()

	data := []byte("Some data")
	blob, err := store.Create("graph", "test/type", bytes.NewReader(data))
	require.NoError(t, err)

	dataReader, err := store.Read("graph", blob.ID)
	buf := bytes.Buffer{}
	buf.ReadFrom(dataReader)

	assert.NoError(t, err)
	assert.Equal(t, data, buf.Bytes())

}

func TestShouldFailWithUnknownBlob(t *testing.T) {
	store := givenEmptyBlobStore()

	newData, err := store.Read("graph", "foo")
	assert.Nil(t, newData)
	assert.Error(t, err)

}
func TestShouldReadAndWriteEmptyBlob(t *testing.T) {
	store := givenEmptyBlobStore()

	blob, err := store.Create("graph", "test/type", bytes.NewReader([]byte{}))
	require.NoError(t, err)
	assert.Equal(t, int64(0), blob.Length)

	dataReader, err := store.Read("graph", blob.ID)
	buf := bytes.Buffer{}
	buf.ReadFrom(dataReader)

	assert.NoError(t, err)
	assert.Empty(t, buf.Bytes())
}

func givenEmptyBlobStore() Store {
	persistence.ResetTestDB()
	db, err := persistence.CreateDBConnection(persistence.TestDBURL())
	if err != nil {
		panic(err)
	}
	store, err := NewSQLBlobStore(db)
	if err != nil {
		panic(err)
	}
	return store
}
