package slack

import (
	"context"
	"encoding/json"
	"net/url"
)

// AssistantThreadsSetStatusParameters contains the parameters for
// SetAssistantThreadStatus.
type AssistantThreadsSetStatusParameters struct {
	ChannelID string `json:"channel_id"`
	ThreadTS  string `json:"thread_ts"`
	Status    string `json:"status"`
}

// AssistantThreadsSetTitleParameters contains the parameters for
// SetAssistantThreadTitle.
type AssistantThreadsSetTitleParameters struct {
	ChannelID string `json:"channel_id"`
	ThreadTS  string `json:"thread_ts"`
	Title     string `json:"title"`
}

// AssistantThreadsSuggestedPrompt represents a suggested prompt shown to the
// user when they open the assistant sidebar.
type AssistantThreadsSuggestedPrompt struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

// AssistantThreadsSetSuggestedPromptsParameters contains the parameters for
// SetAssistantThreadSuggestedPrompts.
type AssistantThreadsSetSuggestedPromptsParameters struct {
	ChannelID string                            `json:"channel_id"`
	ThreadTS  string                            `json:"thread_ts"`
	Title     string                            `json:"title"`
	Prompts   []AssistantThreadsSuggestedPrompt `json:"prompts"`
}

// SetAssistantThreadStatus sets the status (typing indicator) for an assistant
// thread. Pass an empty string to clear the status.
func (api *Client) SetAssistantThreadStatus(params AssistantThreadsSetStatusParameters) error {
	return api.SetAssistantThreadStatusContext(context.Background(), params)
}

// SetAssistantThreadStatusContext sets the status (typing indicator) for an
// assistant thread with a custom context.
func (api *Client) SetAssistantThreadStatusContext(ctx context.Context, params AssistantThreadsSetStatusParameters) error {
	values := url.Values{
		"token":      {api.token},
		"channel_id": {params.ChannelID},
		"thread_ts":  {params.ThreadTS},
		"status":     {params.Status},
	}

	response := &SlackResponse{}
	if err := api.postMethod(ctx, "assistant.threads.setStatus", values, response); err != nil {
		return err
	}

	return response.Err()
}

// SetAssistantThreadTitle sets the title for an assistant thread, shown in the
// conversation history sidebar.
func (api *Client) SetAssistantThreadTitle(params AssistantThreadsSetTitleParameters) error {
	return api.SetAssistantThreadTitleContext(context.Background(), params)
}

// SetAssistantThreadTitleContext sets the title for an assistant thread with a
// custom context.
func (api *Client) SetAssistantThreadTitleContext(ctx context.Context, params AssistantThreadsSetTitleParameters) error {
	values := url.Values{
		"token":      {api.token},
		"channel_id": {params.ChannelID},
		"thread_ts":  {params.ThreadTS},
		"title":      {params.Title},
	}

	response := &SlackResponse{}
	if err := api.postMethod(ctx, "assistant.threads.setTitle", values, response); err != nil {
		return err
	}

	return response.Err()
}

// SetAssistantThreadSuggestedPrompts sets the suggested prompts shown to the
// user in the assistant sidebar.
func (api *Client) SetAssistantThreadSuggestedPrompts(params AssistantThreadsSetSuggestedPromptsParameters) error {
	return api.SetAssistantThreadSuggestedPromptsContext(context.Background(), params)
}

// SetAssistantThreadSuggestedPromptsContext sets the suggested prompts with a
// custom context.
func (api *Client) SetAssistantThreadSuggestedPromptsContext(ctx context.Context, params AssistantThreadsSetSuggestedPromptsParameters) error {
	values := url.Values{
		"token":      {api.token},
		"channel_id": {params.ChannelID},
		"thread_ts":  {params.ThreadTS},
	}

	if params.Title != "" {
		values.Set("title", params.Title)
	}

	if len(params.Prompts) > 0 {
		promptsJSON, err := json.Marshal(params.Prompts)
		if err != nil {
			return err
		}
		values.Set("prompts", string(promptsJSON))
	}

	response := &SlackResponse{}
	if err := api.postMethod(ctx, "assistant.threads.setSuggestedPrompts", values, response); err != nil {
		return err
	}

	return response.Err()
}
