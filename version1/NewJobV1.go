package version1

type NewJobV1 struct {
	Type   string `json:"type"`
	RefId  string `json:"ref_id"`
	Ttl    int    `json:"ttl"` // time to live job in ms
	Params any    `json:"params"`
}
