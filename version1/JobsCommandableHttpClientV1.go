package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type JobsCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewJobsCommandableHttpClientV1() *JobsCommandableHttpClientV1 {
	return &JobsCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/jobs"),
	}
}

func (c *JobsCommandableHttpClientV1) fixJob(job *JobV1) *JobV1 {
	if job == nil {
		return nil
	}

	job.Completed, _ = convert.DateTimeConverter.ToNullableDateTime(job.Completed)
	job.Started, _ = convert.DateTimeConverter.ToNullableDateTime(job.Started)
	job.ExecuteUntil = convert.DateTimeConverter.ToDateTime(job.ExecuteUntil)
	job.LockedUntil, _ = convert.DateTimeConverter.ToNullableDateTime(job.LockedUntil)
	job.Created = convert.DateTimeConverter.ToDateTime(job.Created)

	return job
}

// Add new job
func (c *JobsCommandableHttpClientV1) AddJob(ctx context.Context, correlationId string, newJob *NewJobV1) (*JobV1, error) {
	res, err := c.CallCommand(ctx, "add_job", correlationId, data.NewAnyValueMapFromTuples(
		"new_job", newJob,
	))

	if err != nil {
		return nil, err
	}

	job, err := clients.HandleHttpResponse[*JobV1](res, correlationId)

	return c.fixJob(job), err
}

// Add new job if not exist with same type and ref_id
func (c *JobsCommandableHttpClientV1) AddUniqJob(ctx context.Context, correlationId string, newJob *NewJobV1) (*JobV1, error) {
	res, err := c.CallCommand(ctx, "add_uniq_job", correlationId, data.NewAnyValueMapFromTuples(
		"new_job", newJob,
	))

	if err != nil {
		return nil, err
	}

	job, err := clients.HandleHttpResponse[*JobV1](res, correlationId)

	return c.fixJob(job), err
}

// Get list of all jobs
func (c *JobsCommandableHttpClientV1) GetJobs(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*JobV1], error) {
	res, err := c.CallCommand(ctx, "get_jobs", correlationId, data.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	))

	if err != nil {
		return *data.NewEmptyDataPage[*JobV1](), err
	}

	page, err := clients.HandleHttpResponse[data.DataPage[*JobV1]](res, correlationId)

	if len(page.Data) == 0 {
		return page, err
	}

	for i := range page.Data {
		c.fixJob(page.Data[i])
	}

	return page, err
}

// Get job by Id
func (c *JobsCommandableHttpClientV1) GetJobById(ctx context.Context, correlationId string, jobId string) (*JobV1, error) {
	res, err := c.CallCommand(ctx, "get_job_by_id", correlationId, data.NewAnyValueMapFromTuples(
		"job_id", jobId,
	))

	if err != nil {
		return nil, err
	}

	job, err := clients.HandleHttpResponse[*JobV1](res, correlationId)

	return c.fixJob(job), err
}

// Start job
func (c *JobsCommandableHttpClientV1) StartJobById(ctx context.Context, correlationId string, jobId string, timeout int) (*JobV1, error) {
	res, err := c.CallCommand(ctx, "start_job_by_id", correlationId, data.NewAnyValueMapFromTuples(
		"job_id", jobId,
		"timeout", timeout,
	))

	if err != nil {
		return nil, err
	}

	job, err := clients.HandleHttpResponse[*JobV1](res, correlationId)

	return c.fixJob(job), err
}

// Start fist free job by type
func (c *JobsCommandableHttpClientV1) StartJobByType(ctx context.Context, correlationId string, jobType string, timeout int) (*JobV1, error) {
	res, err := c.CallCommand(ctx, "start_job_by_type", correlationId, data.NewAnyValueMapFromTuples(
		"type", jobType,
		"timeout", timeout,
	))

	if err != nil {
		return nil, err
	}

	job, err := clients.HandleHttpResponse[*JobV1](res, correlationId)

	return c.fixJob(job), err
}

// Extend job execution limit on timeout value
func (c *JobsCommandableHttpClientV1) ExtendJob(ctx context.Context, correlationId string, jobId string, timeout int) (*JobV1, error) {
	res, err := c.CallCommand(ctx, "extend_job", correlationId, data.NewAnyValueMapFromTuples(
		"job_id", jobId,
		"timeout", timeout,
	))

	if err != nil {
		return nil, err
	}

	job, err := clients.HandleHttpResponse[*JobV1](res, correlationId)

	return c.fixJob(job), err
}

// Abort job
func (c *JobsCommandableHttpClientV1) AbortJob(ctx context.Context, correlationId string, jobId string) (*JobV1, error) {
	res, err := c.CallCommand(ctx, "abort_job", correlationId, data.NewAnyValueMapFromTuples(
		"job_id", jobId,
	))

	if err != nil {
		return nil, err
	}

	job, err := clients.HandleHttpResponse[*JobV1](res, correlationId)

	return c.fixJob(job), err
}

// Compleate job
func (c *JobsCommandableHttpClientV1) CompleteJob(ctx context.Context, correlationId string, jobId string) (*JobV1, error) {
	res, err := c.CallCommand(ctx, "complete_job", correlationId, data.NewAnyValueMapFromTuples(
		"job_id", jobId,
	))

	if err != nil {
		return nil, err
	}

	job, err := clients.HandleHttpResponse[*JobV1](res, correlationId)

	return c.fixJob(job), err
}

// Delete job by Id
func (c *JobsCommandableHttpClientV1) DeleteJobById(ctx context.Context, correlationId string, jobId string) (*JobV1, error) {
	res, err := c.CallCommand(ctx, "delete_job_by_id", correlationId, data.NewAnyValueMapFromTuples(
		"job_id", jobId,
	))

	if err != nil {
		return nil, err
	}

	job, err := clients.HandleHttpResponse[*JobV1](res, correlationId)

	return c.fixJob(job), err
}

// Remove all jobs
func (c *JobsCommandableHttpClientV1) DeleteJobs(ctx context.Context, correlationId string) error {
	_, err := c.CallCommand(ctx, "delete_jobs", correlationId, nil)
	return err
}
