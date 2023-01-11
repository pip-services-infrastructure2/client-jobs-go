package version1

import (
	"context"
	"reflect"
	"time"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	refl "github.com/pip-services3-gox/pip-services3-commons-gox/reflect"
)

type JobsMockClientV1 struct {
	maxPageSize int64
	maxRetries  int
	items       []*JobV1
}

func NewJobsMockClientV1() *JobsMockClientV1 {
	return &JobsMockClientV1{
		maxPageSize: 100,
		maxRetries:  10,
		items:       make([]*JobV1, 0),
	}
}

func (c *JobsMockClientV1) composeFilter(filter *data.FilterParams) func(*JobV1) bool {
	if filter == nil {
		filter = data.NewEmptyFilterParams()
	}

	id, idOk := filter.GetAsNullableString("id")
	jobType, jobTypeOk := filter.GetAsNullableString("type")
	refId, refIdOk := filter.GetAsNullableString("ref_id")

	created, createdOk := filter.GetAsNullableDateTime("created")
	createdFrom, createdFromOk := filter.GetAsNullableDateTime("created_from")
	createdTo, createdToOk := filter.GetAsNullableDateTime("created_to")

	started, startedOk := filter.GetAsNullableDateTime("started")
	startedFrom, startedFromOk := filter.GetAsNullableDateTime("started_from")
	startedTo, startedToOk := filter.GetAsNullableDateTime("started_to")

	lockedUntil, lockedUntilOk := filter.GetAsNullableDateTime("locked_until")
	lockedFrom, lockedFromOk := filter.GetAsNullableDateTime("locked_from")
	lockedTo, lockedToOk := filter.GetAsNullableDateTime("locked_to")

	executeUntil, executeUntilOk := filter.GetAsNullableDateTime("execute_until")
	executeFrom, executeFromOk := filter.GetAsNullableDateTime("execute_from")
	executeTo, executeToOk := filter.GetAsNullableDateTime("execute_to")

	completed, completedOk := filter.GetAsNullableDateTime("completed")
	completedFrom, completedFromOk := filter.GetAsNullableDateTime("completed_from")
	completedTo, completedToOk := filter.GetAsNullableDateTime("completed_to")

	retries, retriesOk := filter.GetAsNullableInteger("retries")
	minRetries, minRetriesOk := filter.GetAsNullableInteger("min_retries")

	return func(item *JobV1) bool {
		if idOk && item.Id != id {
			return false
		}
		if jobTypeOk && item.Type != jobType {
			return false
		}
		if refIdOk && item.RefId != refId {
			return false
		}
		if createdOk && item.Created.UnixMilli() != created.UnixMilli() {
			return false
		}
		if createdFromOk && item.Created.UnixMilli() < createdFrom.UnixMilli() {
			return false
		}
		if createdToOk && item.Created.UnixMilli() > createdTo.UnixMilli() {
			return false
		}
		if startedOk && (item.Started.IsZero() || item.Started.UnixMilli() != started.UnixMilli()) {
			return false
		}
		if startedFromOk && (item.Started.IsZero() || item.Started.UnixMilli() < startedFrom.UnixMilli()) {
			return false
		}
		if startedToOk && (item.Started.IsZero() || item.Started.UnixMilli() > startedTo.UnixMilli()) {
			return false
		}
		if lockedUntilOk && (item.LockedUntil.IsZero() || item.LockedUntil.UnixMilli() != lockedUntil.UnixMilli()) {
			return false
		}
		if lockedFromOk && (item.LockedUntil.IsZero() || item.LockedUntil.UnixMilli() < lockedFrom.UnixMilli()) {
			return false
		}
		if lockedToOk && (item.LockedUntil.IsZero() || item.LockedUntil.UnixMilli() > lockedTo.UnixMilli()) {
			return false
		}
		if executeUntilOk && (item.ExecuteUntil.IsZero() || item.ExecuteUntil.UnixMilli() != executeUntil.UnixMilli()) {
			return false
		}
		if executeFromOk && (item.ExecuteUntil.IsZero() || item.ExecuteUntil.UnixMilli() < executeFrom.UnixMilli()) {
			return false
		}
		if executeToOk && (item.ExecuteUntil.IsZero() || item.ExecuteUntil.UnixMilli() > executeTo.UnixMilli()) {
			return false
		}
		if completedOk && (item.ExecuteUntil.IsZero() || item.Completed.UnixMilli() != completed.UnixMilli()) {
			return false
		}
		if completedFromOk && (item.ExecuteUntil.IsZero() || item.Completed.UnixMilli() < completedFrom.UnixMilli()) {
			return false
		}
		if completedToOk && (item.ExecuteUntil.IsZero() || item.Completed.UnixMilli() > completedTo.UnixMilli()) {
			return false
		}
		if retriesOk && item.Retries != retries {
			return false
		}
		if minRetriesOk && item.Retries <= minRetries {
			return false
		}
		return true
	}
}

func (c *JobsMockClientV1) create(ctx context.Context, correlationId string, job *JobV1) *JobV1 {
	if job == nil {
		return nil
	}

	if job.Id == "" {
		job.Id = data.IdGenerator.NextLong()
	}

	buf := *job

	c.items = append(c.items, &buf)

	return job
}

// Add new job
func (c *JobsMockClientV1) AddJob(ctx context.Context, correlationId string, newJob *NewJobV1) (*JobV1, error) {
	job := NewJobV1FromValue(newJob)
	return c.create(ctx, correlationId, job), nil
}

// Add new job if not exist with same type and ref_id
func (c *JobsMockClientV1) AddUniqJob(ctx context.Context, correlationId string, newJob *NewJobV1) (*JobV1, error) {
	filter := data.NewFilterParamsFromTuples(
		"type", newJob.Type,
		"ref_id", newJob.RefId,
	)

	paging := data.NewEmptyPagingParams()
	page, err := c.getPageByFilter(ctx, correlationId, filter, paging)
	if err != nil {
		return nil, nil
	}

	if len(page.Data) > 0 {
		return page.Data[0], nil
	} else {
		job := NewJobV1FromValue(newJob)
		return c.create(ctx, correlationId, job), nil
	}
}

func (c *JobsMockClientV1) getPageByFilter(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*JobV1], error) {
	filterFunc := c.composeFilter(filter)

	items := make([]*JobV1, 0)
	for _, v := range c.items {
		item := v
		if filterFunc(item) {
			items = append(items, item)
		}
	}

	if paging == nil {
		paging = data.NewEmptyPagingParams()
	}

	skip := paging.GetSkip(-1)
	take := paging.GetTake(c.maxPageSize)

	total := 0

	if paging.Total {
		total = len(items)
	}

	if skip > 0 {
		if int(skip) > len(items) {
			skip = int64(len(items))
		}
		items = items[skip:]
	}

	if take > int64(len(items)) {
		take = int64(len(items))
	}

	items = items[0:take]

	return *data.NewDataPage(items, total), nil
}

// Get list of all jobs
func (c *JobsMockClientV1) GetJobs(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*JobV1], error) {
	return c.getPageByFilter(ctx, correlationId, filter, paging)
}

// Get job by Id
func (c *JobsMockClientV1) GetJobById(ctx context.Context, correlationId string, jobId string) (result *JobV1, err error) {
	for _, v := range c.items {
		if v.Id == jobId {
			buf := *v
			result = &buf
			break
		}
	}
	return result, nil
}

// Start job
func (c *JobsMockClientV1) StartJobById(ctx context.Context, correlationId string, jobId string, timeout int) (*JobV1, error) {
	var item *JobV1
	for _, v := range c.items {
		if v.Id == jobId {
			buf := *v
			item = &buf
			break
		}
	}
	if item == nil {
		return nil, nil
	}

	now := time.Now()
	if item.Completed.IsZero() && (item.LockedUntil.IsZero() || item.LockedUntil.UnixMilli() <= now.UnixMilli()) {
		item.Started = now
		item.LockedUntil = now.Add(time.Duration(timeout) * time.Millisecond)
		item.Retries++
		return item, nil
	}

	return nil, nil
}

// Start fist free job by type
func (c *JobsMockClientV1) StartJobByType(ctx context.Context, correlationId string, jobType string, timeout int) (*JobV1, error) {
	now := time.Now()

	var item *JobV1
	for _, v := range c.items {
		if v.Type == jobType && v.Completed.IsZero() && v.Retries < c.maxRetries && (v.LockedUntil.IsZero() || v.LockedUntil.UnixMilli() <= now.UnixMilli()) {
			buf := *v
			item = &buf
			break
		}
	}

	if item == nil {
		return nil, nil
	}

	item.Started = now
	item.LockedUntil = now.Add(time.Duration(timeout) * time.Millisecond)
	item.Retries++
	return item, nil
}

func (c *JobsMockClientV1) updatePartially(ctx context.Context, correlationId string, jobId string, data *data.AnyValueMap) (*JobV1, error) {
	index := -1
	for i, v := range c.items {
		if v.Id == jobId {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}

	buf := *c.items[index]
	newItem := &buf

	if reflect.ValueOf(newItem).Kind() == reflect.Map {
		refl.ObjectWriter.SetProperties(newItem, data.Value())
	} else {
		var intPointer any = newItem
		if reflect.TypeOf(newItem).Kind() != reflect.Pointer {
			objPointer := reflect.New(reflect.TypeOf(newItem))
			objPointer.Elem().Set(reflect.ValueOf(newItem))
			intPointer = objPointer.Interface()
		}
		refl.ObjectWriter.SetProperties(intPointer, data.Value())
		if _newItem, ok := reflect.ValueOf(intPointer).Elem().Interface().(*JobV1); ok {
			newItem = _newItem
		}
	}

	c.items[index] = newItem

	buf = *c.items[index]
	return &buf, nil

}

// Extend job execution limit on timeout value
func (c *JobsMockClientV1) ExtendJob(ctx context.Context, correlationId string, jobId string, timeout int) (*JobV1, error) {
	now := time.Now()
	update := data.NewAnyValueMapFromTuples(
		"locked_until", now.Add(time.Duration(timeout)*time.Millisecond),
	)
	return c.updatePartially(ctx, correlationId, jobId, update)
}

// Abort job
func (c *JobsMockClientV1) AbortJob(ctx context.Context, correlationId string, jobId string) (*JobV1, error) {
	update := data.NewAnyValueMapFromTuples(
		"started", time.Time{},
		"locked_until", time.Time{},
	)
	return c.updatePartially(ctx, correlationId, jobId, update)
}

// Compleate job
func (c *JobsMockClientV1) CompleteJob(ctx context.Context, correlationId string, jobId string) (*JobV1, error) {
	update := data.NewAnyValueMapFromTuples(
		"started", time.Time{},
		"locked_until", time.Time{},
		"completed", time.Now(),
	)
	return c.updatePartially(ctx, correlationId, jobId, update)
}

// Delete job by Id
func (c *JobsMockClientV1) DeleteJobById(ctx context.Context, correlationId string, jobId string) (*JobV1, error) {
	var index = -1
	for i, v := range c.items {
		if v.Id == jobId {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}

	var item = c.items[index]
	if index < len(c.items) {
		c.items = append(c.items[:index], c.items[index+1:]...)
	} else {
		c.items = c.items[:index]
	}

	return item, nil
}

// Remove all jobs
func (c *JobsMockClientV1) DeleteJobs(ctx context.Context, correlationId string) error {
	c.items = make([]*JobV1, 0)
	return nil
}
