package slack

import (
	"context"
	"errors"
	"net/url"
)

type Bookmark struct {
	ID                  string  `json:"id"`
	ChannelID           string  `json:"channel_id"`
	Title               string  `json:"title"`
	Link                string  `json:"link"`
	Emoji               string  `json:"emoji,omitempty"`
	IconURL             *string `json:"icon_url"`
	Type                string  `json:"type"`
	DateCreated         uint64  `json:"date_created"`
	DateUpdated         uint64  `json:"date_updated"`
	Rank                string  `json:"rank"`
	LastUpdatedByUserID *string `json:"last_updated_by_user_id"`
	LastUpdatedByTeamID *string `json:"last_updated_by_team_id"`
	ShortcutID          *string `json:"shortcut_id"`
	EntityID            *string `json:"entity_id"`
	AppID               *string `json:"app_id"`
}

type AddBookmarkParameters struct {
	Title     string // A required title for the bookmark
	Type      string // A required type for the bookmark
	Link      string // URL required for type:link
	Emoji     string // An optional emoji
	EntityID  string
	ParentID  string
	ChannelID string `json:"channel_id"`
}

type EditBookmarkParameters struct {
	Title      *string // Change the title. Set to "" to clear
	Emoji      *string // Change the emoji. Set to "" to clear
	Link       string  // Change the link
	ChannelID  string  `json:"channel_id"`
	BookmarkID string  `json:"bookmark_id"`
	Type       string  `json:"type,omitempty"`
}

type addBookmarkResponse struct {
	Bookmark Bookmark `json:"bookmark"`
	SlackResponse
}

type editBookmarkResponse struct {
	Bookmark Bookmark `json:"bookmark"`
	SlackResponse
}

type listBookmarksResponse struct {
	Bookmarks []Bookmark `json:"bookmarks"`
	SlackResponse
}

// AddBookmark adds a bookmark in a channel.
// For more details, see AddBookmarkContext documentation.
func (api *Client) AddBookmark(channelID string, params AddBookmarkParameters) (Bookmark, error) {
	return api.AddBookmarkContext(context.Background(), channelID, params)
}

// AddBookmarkContext adds a bookmark in a channel with a custom context.
// Slack API docs: https://api.slack.com/methods/bookmarks.add
func (api *Client) AddBookmarkContext(ctx context.Context, channelID string, params AddBookmarkParameters) (Bookmark, error) {
	values := url.Values{
		"channel_id": {channelID},
		"token":      {api.token},
		"title":      {params.Title},
		"type":       {params.Type},
	}
	if params.Link != "" {
		values.Set("link", params.Link)
	}
	if params.Emoji != "" {
		values.Set("emoji", params.Emoji)
	}
	if params.EntityID != "" {
		values.Set("entity_id", params.EntityID)
	}
	if params.ParentID != "" {
		values.Set("parent_id", params.ParentID)
	}

	response := &addBookmarkResponse{}
	if err := api.postMethod(ctx, "bookmarks.add", values, response); err != nil {
		return Bookmark{}, err
	}

	return response.Bookmark, response.Err()
}

// RemoveBookmark removes a bookmark from a channel.
// For more details, see RemoveBookmarkContext documentation.
func (api *Client) RemoveBookmark(channelID, bookmarkID string) error {
	return api.RemoveBookmarkContext(context.Background(), channelID, bookmarkID)
}

// RemoveBookmarkContext removes a bookmark from a channel with a custom context.
// Slack API docs: https://api.slack.com/methods/bookmarks.remove
func (api *Client) RemoveBookmarkContext(ctx context.Context, channelID, bookmarkID string) error {
	values := url.Values{
		"channel_id":  {channelID},
		"token":       {api.token},
		"bookmark_id": {bookmarkID},
	}

	response := &SlackResponse{}
	if err := api.postMethod(ctx, "bookmarks.remove", values, response); err != nil {
		return err
	}

	return response.Err()
}

// ListBookmarks returns all bookmarks for a channel.
// For more details, see ListBookmarksContext documentation.
func (api *Client) ListBookmarks(channelID string) ([]Bookmark, error) {
	return api.ListBookmarksContext(context.Background(), channelID)
}

// ListBookmarksContext returns all bookmarks for a channel with a custom context.
// Slack API docs: https://api.slack.com/methods/bookmarks.edit
func (api *Client) ListBookmarksContext(ctx context.Context, channelID string) ([]Bookmark, error) {
	values := url.Values{
		"token":      {api.token},
		"channel_id": {channelID},
	}

	response := &listBookmarksResponseFull{}
	err := api.postMethod(ctx, "bookmarks.list", values, response)
	if err != nil {
		return nil, err
	}
	if !response.Ok {
		return nil, errors.New(response.Error)
	}
	return response.Bookmarks, nil
}

// EditBookmark edits a bookmark in a channel.
// For more details, see EditBookmarkContext documentation.
func (api *Client) EditBookmark(channelID, bookmarkID string, params EditBookmarkParameters) (Bookmark, error) {
	return api.EditBookmarkContext(context.Background(), channelID, bookmarkID, params)
}

// EditBookmarkContext edits a bookmark in a channel with a custom context.
// Slack API docs: https://api.slack.com/methods/bookmarks.edit
func (api *Client) EditBookmarkContext(ctx context.Context, channelID, bookmarkID string, params EditBookmarkParameters) (Bookmark, error) {
	values := url.Values{
		"token":       {api.token},
		"channel_id":  {params.ChannelID},
		"bookmark_id": {params.BookmarkID},
	}

	if params.Type != "" {
		values["type"] = []string{params.Type}
	}

	if params.Emoji != nil {
		values["emoji"] = []string{*params.Emoji}
	}

	if params.Link != "" {
		values["link"] = []string{params.Link}
	}

	if params.Title != nil {
		values["title"] = []string{*params.Title}
	}

	response := &editBookmarkResponse{}
	err := api.postMethod(ctx, "bookmarks.edit", values, &response)

	if err != nil {
		return Bookmark{}, err
	}
	if err := response.Err(); err != nil {
		return Bookmark{}, err
	}

	return response.Bookmark, nil
}

type listBookmarksResponseFull struct {
	Bookmarks []Bookmark
	SlackResponse
}

type singleBookmarkResponse struct {
	Bookmark Bookmark `json:"bookmark"`
	SlackResponse
}
