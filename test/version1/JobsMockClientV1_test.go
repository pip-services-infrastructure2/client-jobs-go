package test_version1

import (
	"testing"

	"github.com/pip-services-infrastructure2/client-jobs-go/version1"
)

type jobsMockClientV1Test struct {
	client  *version1.JobsMockClientV1
	fixture *JobsClientV1Fixture
}

func newJobsMockClientV1Test() *jobsMockClientV1Test {
	return &jobsMockClientV1Test{}
}

func (c *jobsMockClientV1Test) setup(t *testing.T) {
	c.client = version1.NewJobsMockClientV1()
	c.fixture = NewJobsClientV1Fixture(c.client)
}

func (c *jobsMockClientV1Test) teardown(t *testing.T) {
	c.client = nil
}

func TestMockCrudOperations(t *testing.T) {
	c := newJobsMockClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestCrudOperations(t)
}

func TestMockControl(t *testing.T) {
	c := newJobsMockClientV1Test()
	c.setup(t)
	defer c.teardown(t)

	c.fixture.TestControl(t)
}
