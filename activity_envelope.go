package tumblr

type ActivityEnvelope struct {
	Id              string   `json:"id"`
	Timestamp       int64    `json:"timestamp"`
	Version         string   `json:"version"`
	ActivityPrivacy string   `json:"activity_privacy"`
	ActivityType    string   `json:"activity_type"`
	Activity        Activity `json:"activity"`
}
