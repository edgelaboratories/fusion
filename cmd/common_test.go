package cmd

import (
	"io/ioutil"
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
	tmpFile, err := ioutil.TempFile(os.TempDir(), "fusion-")
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
	Age  int
	Key  string
	Tags map[string]int
}

func (suite *ConfigTestSuite) TestMerge() {
	f1 := "testdata/file1.yml"
	f2 := "testdata/file2.yml"

	tmpFilename := suite.tmpFile.Name()
	suite.NoError(merge(suite.config, tmpFilename, f1, f2))

	resFile, err := ioutil.ReadFile(tmpFilename)
	suite.NoError(err)

	var res dummy
	suite.NoError(yml.Unmarshal(resFile, &res))

	suite.Equal(34, res.Age)
	suite.Equal("val2", res.Key)
	suite.Equal(1, res.Tags["key1"])
	suite.Equal(2, res.Tags["key2"])
	suite.Equal(3, res.Tags["key3"])
}

func (suite *ConfigTestSuite) TestWriteToFile() {
	suite.Require().NoError(suite.config.LoadFiles("testdata/file1.yml"))

	tmpFilename := suite.tmpFile.Name()

	suite.NoError(writeToFile(suite.config, tmpFilename))

	resFile, err := ioutil.ReadFile(tmpFilename)
	suite.NoError(err)

	var res dummy
	suite.NoError(yml.Unmarshal(resFile, &res))

	suite.Equal(34, res.Age)
	suite.Equal("val", res.Key)
	suite.Equal(1, res.Tags["key1"])
	suite.Equal(2, res.Tags["key2"])
}
