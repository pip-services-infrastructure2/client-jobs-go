package version1

import "time"

type JobV1 struct {
	// Job description
	Id     string `json:"id"`
	Type   string `json:"type"`
	RefId  string `json:"ref_id"`
	Params any    `json:"params"`

	// Job control
	Created      time.Time `json:"created"`
	Started      time.Time `json:"started"`
	LockedUntil  time.Time `json:"locked_until"`
	ExecuteUntil time.Time `json:"execute_until"`
	Completed    time.Time `json:"completed"`
	Retries      int       `json:"retries"`
	Ttl          int       `json:"ttl"`
}

func NewJobV1FromValue(newJob *NewJobV1) *JobV1 {
	now := time.Now()
	c := &JobV1{
		Created: now,
		Retries: 0,
	}

	if newJob != nil {
		c.Type = newJob.Type
		c.RefId = newJob.RefId
		c.Params = newJob.Params
		if newJob.Ttl > 0 {
			c.ExecuteUntil = now.Add(time.Duration(newJob.Ttl) * time.Microsecond)
		}
	}

	return c
}
