package story

type Story struct {
	Title    string
	Type     string
	State    string
	Estimate int
}

type SlackFormat struct {
	Attachments []SlackMessage `json:"attachments"`
}

type SlackMessage struct {
}
