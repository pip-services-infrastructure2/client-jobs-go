package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-infrastructure2/client-jobs-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/stretchr/testify/assert"
)

type JobsClientV1Fixture struct {
	Client version1.IJobsClientV1
	JOB1   *version1.NewJobV1
	JOB2   *version1.NewJobV1
	JOB3   *version1.NewJobV1
}

func NewJobsClientV1Fixture(client version1.IJobsClientV1) *JobsClientV1Fixture {
	return &JobsClientV1Fixture{
		Client: client,
		JOB1: &version1.NewJobV1{
			//Id: "Job1_t1_0fsd",
			Type:   "t1",
			RefId:  "obj_0fsd",
			Params: nil,
			Ttl:    1000 * 60 * 60 * 3, // 3 hour
			// Retries: 5
		},
		JOB2: &version1.NewJobV1{
			//Id: "Job2_t1_0fsd",
			Type:   "t1",
			RefId:  "obj_0fsd",
			Params: nil,
			Ttl:    1000 * 60 * 60, // 1 hour
			// Retries: 3
		},
		JOB3: &version1.NewJobV1{
			//Id: "Job2_t1_0fsd",
			Type:   "t2",
			RefId:  "obj_3fsd",
			Params: nil,
			Ttl:    1000 * 60 * 30, // 30 minutes
			// Retries: 2
		},
	}
}

func (c *JobsClientV1Fixture) TestCrudOperations(t *testing.T) {
	// Create the first job
	job1, err := c.Client.AddJob(context.Background(), "123", c.JOB1)
	assert.Nil(t, err)

	assert.NotNil(t, job1)
	assert.NotEmpty(t, job1.Id)
	assert.Equal(t, job1.Type, c.JOB1.Type)
	assert.Equal(t, job1.RefId, c.JOB1.RefId)
	assert.Equal(t, 0, job1.Retries)
	assert.Equal(t, job1.Params, c.JOB1.Params)
	assert.True(t, !job1.Created.IsZero())
	assert.True(t, !job1.ExecuteUntil.IsZero())
	assert.True(t, job1.Started.IsZero())
	assert.True(t, job1.Completed.IsZero())
	assert.True(t, job1.LockedUntil.IsZero())

	// Create the second job
	job2, err := c.Client.AddUniqJob(context.Background(), "123", c.JOB2)
	assert.Nil(t, err)

	assert.NotNil(t, job2)
	assert.NotEmpty(t, job2.Id)
	assert.Equal(t, job2.Type, c.JOB2.Type)
	assert.Equal(t, job2.RefId, c.JOB2.RefId)
	assert.Equal(t, 0, job2.Retries)
	assert.Equal(t, job2.Params, c.JOB2.Params)
	assert.True(t, !job2.Created.IsZero())
	assert.True(t, !job2.ExecuteUntil.IsZero())
	assert.True(t, job2.Started.IsZero())
	assert.True(t, job2.Completed.IsZero())
	assert.True(t, job2.LockedUntil.IsZero())

	// Create the third job
	job3, err := c.Client.AddJob(context.Background(), "123", c.JOB3)
	assert.Nil(t, err)

	assert.NotNil(t, job3)
	assert.NotEmpty(t, job3.Id)
	assert.Equal(t, job3.Type, c.JOB3.Type)
	assert.Equal(t, job3.RefId, c.JOB3.RefId)
	assert.Equal(t, 0, job2.Retries)
	assert.Equal(t, job3.Params, c.JOB3.Params)
	assert.True(t, !job3.Created.IsZero())
	assert.True(t, !job3.ExecuteUntil.IsZero())
	assert.True(t, job3.Started.IsZero())
	assert.True(t, job3.Completed.IsZero())
	assert.True(t, job3.LockedUntil.IsZero())

	// Get one job
	job, err := c.Client.GetJobById(context.Background(), "123", job1.Id)
	assert.Nil(t, err)

	assert.NotNil(t, job)
	assert.Equal(t, job1.Type, job.Type)
	assert.Equal(t, c.JOB1.RefId, job.RefId)
	assert.Equal(t, c.JOB1.Type, job.Type)
	assert.Equal(t, c.JOB1.RefId, job.RefId)
	assert.Equal(t, job1.Retries, job.Retries)
	assert.Equal(t, c.JOB1.Params, job.Params)
	assert.True(t, !job.Created.IsZero())
	assert.True(t, !job.ExecuteUntil.IsZero())
	assert.True(t, job.Started.IsZero())
	assert.True(t, job.Completed.IsZero())
	assert.True(t, job.LockedUntil.IsZero())

	// Get all jobs
	page, err := c.Client.GetJobs(context.Background(), "123", nil, data.NewPagingParams(0, 5, false))
	assert.Nil(t, err)

	assert.Len(t, page.Data, 2)

	job1 = page.Data[0]

	// Delete the job
	job, err = c.Client.DeleteJobById(context.Background(), "123", job1.Id)
	assert.Nil(t, err)

	assert.Equal(t, job1.Id, job.Id)

	// Try to get deleted job
	job, err = c.Client.GetJobById(context.Background(), "123", job1.Id)
	assert.Nil(t, err)
	assert.Nil(t, job)

	// Delete all jobs
	err = c.Client.DeleteJobs(context.Background(), "123")
	assert.Nil(t, err)

	// Try to get jobs after delete
	page, err = c.Client.GetJobs(context.Background(), "123", nil, nil)
	assert.Nil(t, err)
	assert.Len(t, page.Data, 0)
}

func (c *JobsClientV1Fixture) TestControl(t *testing.T) {
	// Create the first job
	job1, err := c.Client.AddJob(context.Background(), "123", c.JOB1)
	assert.Nil(t, err)

	assert.NotNil(t, job1)
	assert.NotEmpty(t, job1.Id)
	assert.Equal(t, job1.Type, c.JOB1.Type)
	assert.Equal(t, job1.RefId, c.JOB1.RefId)
	assert.Equal(t, 0, job1.Retries)
	assert.Equal(t, job1.Params, c.JOB1.Params)
	assert.True(t, !job1.Created.IsZero())
	assert.True(t, !job1.ExecuteUntil.IsZero())
	assert.True(t, job1.Started.IsZero())
	assert.True(t, job1.Completed.IsZero())
	assert.True(t, job1.LockedUntil.IsZero())

	// Create the second job
	job2, err := c.Client.AddUniqJob(context.Background(), "123", c.JOB2)
	assert.Nil(t, err)

	assert.NotNil(t, job2)
	assert.NotEmpty(t, job2.Id)
	assert.Equal(t, job2.Type, c.JOB2.Type)
	assert.Equal(t, job2.RefId, c.JOB2.RefId)
	assert.Equal(t, 0, job2.Retries)
	assert.Equal(t, job2.Params, c.JOB2.Params)
	assert.True(t, !job2.Created.IsZero())
	assert.True(t, !job2.ExecuteUntil.IsZero())
	assert.True(t, job2.Started.IsZero())
	assert.True(t, job2.Completed.IsZero())
	assert.True(t, job2.LockedUntil.IsZero())

	// Create the third job
	job3, err := c.Client.AddJob(context.Background(), "123", c.JOB3)
	assert.Nil(t, err)

	assert.NotNil(t, job3)
	assert.NotEmpty(t, job3.Id)
	assert.Equal(t, job3.Type, c.JOB3.Type)
	assert.Equal(t, job3.RefId, c.JOB3.RefId)
	assert.Equal(t, 0, job2.Retries)
	assert.Equal(t, job3.Params, c.JOB3.Params)
	assert.True(t, !job3.Created.IsZero())
	assert.True(t, !job3.ExecuteUntil.IsZero())
	assert.True(t, job3.Started.IsZero())
	assert.True(t, job3.Completed.IsZero())
	assert.True(t, job3.LockedUntil.IsZero())

	// Get one job
	job, err := c.Client.GetJobById(context.Background(), "123", job1.Id)
	assert.Nil(t, err)

	assert.NotNil(t, job)
	assert.Equal(t, job1.Type, job.Type)
	assert.Equal(t, c.JOB1.RefId, job.RefId)
	assert.Equal(t, c.JOB1.Type, job.Type)
	assert.Equal(t, c.JOB1.RefId, job.RefId)
	assert.Equal(t, job1.Retries, job.Retries)
	assert.Equal(t, c.JOB1.Params, job.Params)
	assert.True(t, !job.Created.IsZero())
	assert.True(t, !job.ExecuteUntil.IsZero())
	assert.True(t, job.Started.IsZero())
	assert.True(t, job.Completed.IsZero())
	assert.True(t, job.LockedUntil.IsZero())

	// Get all jobs
	page, err := c.Client.GetJobs(context.Background(), "123", nil, data.NewPagingParams(0, 5, false))
	assert.Nil(t, err)
	assert.Len(t, page.Data, 2)

	job1 = page.Data[0]
	job2 = page.Data[1]

	// Test start job
	job, err = c.Client.StartJobByType(context.Background(), "123", job1.Type, 1000*60*10)
	assert.Nil(t, err)

	assert.NotNil(t, job)
	assert.True(t, !job.Started.IsZero())
	assert.True(t, !job.LockedUntil.IsZero())
	job1 = job

	// Test extend job
	job, err = c.Client.ExtendJob(context.Background(), "123", job1.Id, 1000*60*2)
	assert.Nil(t, err)

	assert.NotNil(t, job)
	assert.True(t, !job.LockedUntil.IsZero())
	job1 = job

	// Test complete job
	job, err = c.Client.CompleteJob(context.Background(), "123", job1.Id)
	assert.Nil(t, err)

	assert.NotNil(t, job)
	assert.True(t, !job.Completed.IsZero())
	job1 = job

	// Test start job
	job, err = c.Client.StartJobById(context.Background(), "123", job2.Id, 1000*60)
	assert.Nil(t, err)

	assert.NotNil(t, job)
	assert.True(t, !job.Started.IsZero())
	assert.True(t, !job.LockedUntil.IsZero())
	job2 = job

	// Test abort job
	job, err = c.Client.AbortJob(context.Background(), "123", job2.Id)
	assert.Nil(t, err)

	assert.NotNil(t, job)
	assert.True(t, job.Started.IsZero())
	assert.True(t, job.LockedUntil.IsZero())
}
