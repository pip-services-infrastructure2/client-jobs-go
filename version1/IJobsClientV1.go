package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IJobsClientV1 interface {
	// Add new job
	AddJob(ctx context.Context, correlationId string, newJob *NewJobV1) (*JobV1, error)
	// Add new job if not exist with same type and ref_id
	AddUniqJob(ctx context.Context, correlationId string, newJob *NewJobV1) (*JobV1, error)
	// Get list of all jobs
	GetJobs(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*JobV1], error)
	// Get job by Id
	GetJobById(ctx context.Context, correlationId string, jobId string) (*JobV1, error)
	// Start job
	StartJobById(ctx context.Context, correlationId string, jobId string, timeout int) (*JobV1, error)
	// Start fist free job by type
	StartJobByType(ctx context.Context, correlationId string, jobType string, timeout int) (*JobV1, error)
	// Extend job execution limit on timeout value
	ExtendJob(ctx context.Context, correlationId string, jobId string, timeout int) (*JobV1, error)
	// Abort job
	AbortJob(ctx context.Context, correlationId string, jobId string) (*JobV1, error)
	// Compleate job
	CompleteJob(ctx context.Context, correlationId string, jobId string) (*JobV1, error)
	// Delete job by Id
	DeleteJobById(ctx context.Context, correlationId string, jobId string) (*JobV1, error)
	// Remove all jobs
	DeleteJobs(ctx context.Context, correlationId string) error
}
