package orderfile

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// fileRemoverFunc removes pre-targeted file when executed
type fileRemoverFunc func()

// createTestFixture in form of temporary file with given content
// returns path to file and hook-like remover function
func createTestFixture(t *testing.T, content string) (string, fileRemoverFunc) {
	fixtureFile, err := ioutil.TempFile("", "test_fixture_*.yml")
	require.NoError(t, err)

	err = ioutil.WriteFile(fixtureFile.Name(), []byte(content), 0600)
	require.NoError(t, err)

	return fixtureFile.Name(), func() { require.NoError(t, os.Remove(fixtureFile.Name())) }
}

// @ Given - empty file
// @ When  - trying to create new orderfile instance
// @ Then  - returned Orderfile constains no orders and no version info
func TestEmptyOrderfileLoadingFromEmptyFile(t *testing.T) {
	fixture, remover := createTestFixture(t, "")
	defer remover()

	orderFile, err := NewOrderFileFrom(fixture)
	assert.NoError(t, err)
	assert.Equal(t, []string{}, orderFile.ListOrdersNames())
	assert.Equal(t, "", orderFile.Version)
}

// @ Given - syntactically correct YAML file containing wrong fields/content
// @ When  - trying to create new orderfile instance
// @ Then  - returned error informs about ambiguous field
func TestErrorWhenLoadingFromNonOrdrefileYAML(t *testing.T) {
	differentSyntaxedYaml := `x:
  a: 1
  b: 2
  c: 3
`
	fixture, remover := createTestFixture(t, differentSyntaxedYaml)

	defer remover()

	_, err := NewOrderFileFrom(fixture)
	assert.EqualError(t, err, "yaml: unmarshal errors:\n  line 1: field x not found in type orderfile.Orderfile")
}

// @ Given - YAML file with syntactically wrong content
// @ When  - trying to create new orderfile instance
// @ Then  - returned error informs about lack of availability to unmarshal file into Orderfile
func TestErrorForLoadingFromValidOrderfile(t *testing.T) {
	fixture, remover := createTestFixture(t, "InvalidYamlFileContent")
	defer remover()

	_, err := NewOrderFileFrom(fixture)
	assert.EqualError(t, err, "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `Invalid...` into orderfile.Orderfile")
}

// @ Given - non-existing path
// @ When  - trying to create new orderfile instance
// @ Then  - returned error informs about "no such file or directory"
func TestErrorForNonExistingPath(t *testing.T) {
	_, err := NewOrderFileFrom("non-existing-path")

	possibleNonExistingPathsErrors := []string{
		"open non-existing-path: no such file or directory",                  // linux
		"open non-existing-path: The system cannot find the file specified.", // windows
	}

	assert.Contains(t, possibleNonExistingPathsErrors, err.Error())
}
