package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-infrastructure2/client-jobs-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type jobsCommandableHttpClientV1Test struct {
	client  *version1.JobsCommandableHttpClientV1
	fixture *JobsClientV1Fixture
}

func newJobsCommandableHttpClientV1Test() *jobsCommandableHttpClientV1Test {
	return &jobsCommandableHttpClientV1Test{}
}

func (c *jobsCommandableHttpClientV1Test) setup(t *testing.T) {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewJobsCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewJobsClientV1Fixture(c.client)
}

func (c *jobsCommandableHttpClientV1Test) teardown(t *testing.T) {
	c.client.DeleteJobs(context.Background(), "123")
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpCrudOperations(t *testing.T) {
	c := newJobsCommandableHttpClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestCrudOperations(t)
}

func TestCommandableHttpControl(t *testing.T) {
	c := newJobsCommandableHttpClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestControl(t)
}
