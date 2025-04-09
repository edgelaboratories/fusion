package cmd

import (
	"os"
	"testing"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/stretchr/testify/suite"
	yml "gopkg.in/yaml.v2"
)

type ConfigTestSuite struct {
	suite.Suite

	config  *config.Config
	tmpFile *os.File
}

func (suite *ConfigTestSuite) SetupSuite() {
	suite.config = config.New("test")

	// Since we are only testing common functions,
	// we only test with YAML but any format is ok
	suite.config.AddDriver(yaml.Driver)
	suite.config.WithOptions(func(opts *config.Options) {
		opts.DumpFormat = config.Yaml
	})
}

func (suite *ConfigTestSuite) TearDownSuite() {
	suite.config.ClearAll()
}

func (suite *ConfigTestSuite) SetupTest() {
	tmpFile, err := os.CreateTemp(os.TempDir(), "fusion-")
	suite.Require().NoError(err)
	suite.tmpFile = tmpFile
}

func (suite *ConfigTestSuite) TearDownTest() {
	os.Remove(suite.tmpFile.Name())
	suite.config.ClearData()
}

func TestConfigTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(ConfigTestSuite))
}

type dummy struct {
	Username string
	Email    string
	Age      int
	Key      string
	Tags     map[string]int
}

func (suite *ConfigTestSuite) TestMerge() {
	f1 := "testdata/file1.yml"
	f2 := "testdata/file2.yml"
	exp := "testdata/expected.yml"

	tmpFilename := suite.tmpFile.Name()
	suite.Require().NoError(merge(suite.config, tmpFilename, f1, f2))

	resultFile, err := os.ReadFile(tmpFilename)
	suite.NoError(err)

	var result dummy
	suite.NoError(yml.Unmarshal(resultFile, &result))

	expectedFile, err := os.ReadFile(exp)
	suite.NoError(err)

	var expected dummy
	suite.NoError(yml.Unmarshal(expectedFile, &expected))

	suite.Equal(expected, result)
}

func (suite *ConfigTestSuite) TestWriteToFile() {
	file := "testdata/file1.yml"
	suite.Require().NoError(suite.config.LoadFiles(file))

	tmpFilename := suite.tmpFile.Name()

	suite.Require().NoError(writeToFile(suite.config, tmpFilename))

	resultFile, err := os.ReadFile(tmpFilename)
	suite.NoError(err)

	var result dummy
	suite.NoError(yml.Unmarshal(resultFile, &result))

	expectedFile, err := os.ReadFile(file)
	suite.NoError(err)

	var expected dummy
	suite.NoError(yml.Unmarshal(expectedFile, &expected))

	suite.Equal(expected, result)
}
