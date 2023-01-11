package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type JobsNullClientV1 struct {
}

func NewJobsNullClientV1() *JobsNullClientV1 {
	return &JobsNullClientV1{}
}

// Add new job
func (c *JobsNullClientV1) AddJob(ctx context.Context, correlationId string, newJob *NewJobV1) (*JobV1, error) {
	job := NewJobV1FromValue(newJob)
	return job, nil
}

// Add new job if not exist with same type and ref_id
func (c *JobsNullClientV1) AddUniqJob(ctx context.Context, correlationId string, newJob *NewJobV1) (*JobV1, error) {
	job := NewJobV1FromValue(newJob)
	return job, nil
}

// Get list of all jobs
func (c *JobsNullClientV1) GetJobs(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*JobV1], error) {
	return *data.NewEmptyDataPage[*JobV1](), nil
}

// Get job by Id
func (c *JobsNullClientV1) GetJobById(ctx context.Context, correlationId string, jobId string) (*JobV1, error) {
	return nil, nil
}

// Start job
func (c *JobsNullClientV1) StartJobById(ctx context.Context, correlationId string, jobId string, timeout int) (*JobV1, error) {
	return nil, nil
}

// Start fist free job by type
func (c *JobsNullClientV1) StartJobByType(ctx context.Context, correlationId string, jobType string, timeout int) (*JobV1, error) {
	return nil, nil
}

// Extend job execution limit on timeout value
func (c *JobsNullClientV1) ExtendJob(ctx context.Context, correlationId string, jobId string, timeout int) (*JobV1, error) {
	return nil, nil
}

// Abort job
func (c *JobsNullClientV1) AbortJob(ctx context.Context, correlationId string, jobId string) (*JobV1, error) {
	return nil, nil
}

// Compleate job
func (c *JobsNullClientV1) CompleteJob(ctx context.Context, correlationId string, jobId string) (*JobV1, error) {
	return nil, nil
}

// Delete job by Id
func (c *JobsNullClientV1) DeleteJobById(ctx context.Context, correlationId string, jobId string) (*JobV1, error) {
	return nil, nil
}

// Remove all jobs
func (c *JobsNullClientV1) DeleteJobs(ctx context.Context, correlationId string) error {
	return nil
}
