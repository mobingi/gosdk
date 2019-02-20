package dockerhub

type PushData struct {
	Pusher string `json:"pusher"`
	Tag    string `json:"tag"`
}

type Repo struct {
	Name      string `json:"name"`
	Owner     string `json:"owner"`
	Namespace string `json:"namespace"`
	RepoName  string `json:"repo_name"`
}

// WebhookPayload is the representation of DockerHub webhook POST payload.
// We only include the ones we are interested in.
type WebhookPayload struct {
	CallbackUrl string   `json:"callback_url"`
	PushData    PushData `json:"push_data"`
	Repo        Repo     `json:"repository"`
}
