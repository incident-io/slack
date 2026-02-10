package slackevents

// AssistantThreadStartedEvent fires when a user opens the assistant sidebar and
// starts a new thread.
type AssistantThreadStartedEvent struct {
	Type            string                 `json:"type"`
	AssistantThread AssistantThreadPayload `json:"assistant_thread"`
}

// AssistantThreadPayload contains the thread metadata sent with assistant events.
type AssistantThreadPayload struct {
	UserID    string                        `json:"user_id"`
	ChannelID string                        `json:"channel_id"`
	ThreadTS  string                        `json:"thread_ts"`
	Context   AssistantThreadContextPayload `json:"context"`
}

// AssistantThreadContextPayload carries context about where the user was when
// the assistant event fired.
type AssistantThreadContextPayload struct {
	ChannelID string `json:"channel_id"`
	TeamID    string `json:"team_id"`
}

// AssistantThreadContextChangedEvent fires when the user navigates to a
// different channel while the assistant sidebar is open.
type AssistantThreadContextChangedEvent struct {
	Type            string                 `json:"type"`
	AssistantThread AssistantThreadPayload `json:"assistant_thread"`
}
