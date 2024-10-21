package git

type Repository struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Deployment struct {
	Ref                   string      `json:"ref,omitempty"`
	Task                  string      `json:"task,omitempty"`
	AutoMerge             bool        `json:"auto_merge,omitempty"`
	RequiredContexts      []string    `json:"required_contexts,omitempty"`
	Payload               interface{} `json:"payload,omitempty"`
	Environment           string      `json:"environment,omitempty"`
	Description           string      `json:"description,omitempty"`
	TransientEnvironment  bool        `json:"transient_environment,omitempty"`
	ProductionEnvironment bool        `json:"production_environment,omitempty"`
}
