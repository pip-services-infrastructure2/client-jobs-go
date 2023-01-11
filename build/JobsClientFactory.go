package build

import (
	clients1 "github.com/pip-services-infrastructure2/client-jobs-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type JobsClientFactory struct {
	*cbuild.Factory
}

func NewJobsClientFactory() *JobsClientFactory {
	c := &JobsClientFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-jobs", "client", "null", "*", "1.0")
	mockClientDescriptor := cref.NewDescriptor("service-jobs", "client", "mock", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-jobs", "client", "commandable-http", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewJobsNullClientV1)
	c.RegisterType(mockClientDescriptor, clients1.NewJobsMockClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewJobsCommandableHttpClientV1)

	return c
}
