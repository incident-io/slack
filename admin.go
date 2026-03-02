package slack

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func (api *Client) adminRequest(ctx context.Context, method string, teamName string, values url.Values) error {
	resp := &SlackResponse{}
	err := parseAdminResponse(ctx, api.httpclient, method, teamName, values, resp, api)
	if err != nil {
		return err
	}

	return resp.Err()
}

// DisableUser disabled a user account, given a user ID
func (api *Client) DisableUser(teamName string, uid string) error {
	return api.DisableUserContext(context.Background(), teamName, uid)
}

// DisableUserContext disabled a user account, given a user ID with a custom context
func (api *Client) DisableUserContext(ctx context.Context, teamName string, uid string) error {
	values := url.Values{
		"user":       {uid},
		"token":      {api.token},
		"set_active": {"true"},
		"_attempts":  {"1"},
	}

	if err := api.adminRequest(ctx, "setInactive", teamName, values); err != nil {
		return fmt.Errorf("failed to disable user with id '%s': %s", uid, err)
	}

	return nil
}

// InviteGuest invites a user to Slack as a single-channel guest
func (api *Client) InviteGuest(teamName, channel, firstName, lastName, emailAddress string) error {
	return api.InviteGuestContext(context.Background(), teamName, channel, firstName, lastName, emailAddress)
}

// InviteGuestContext invites a user to Slack as a single-channel guest with a custom context
func (api *Client) InviteGuestContext(ctx context.Context, teamName, channel, firstName, lastName, emailAddress string) error {
	values := url.Values{
		"email":            {emailAddress},
		"channels":         {channel},
		"first_name":       {firstName},
		"last_name":        {lastName},
		"ultra_restricted": {"1"},
		"token":            {api.token},
		"resend":           {"true"},
		"set_active":       {"true"},
		"_attempts":        {"1"},
	}

	err := api.adminRequest(ctx, "invite", teamName, values)
	if err != nil {
		return fmt.Errorf("Failed to invite single-channel guest: %s", err)
	}

	return nil
}

// InviteRestricted invites a user to Slack as a restricted account
func (api *Client) InviteRestricted(teamName, channel, firstName, lastName, emailAddress string) error {
	return api.InviteRestrictedContext(context.Background(), teamName, channel, firstName, lastName, emailAddress)
}

// InviteRestrictedContext invites a user to Slack as a restricted account with a custom context
func (api *Client) InviteRestrictedContext(ctx context.Context, teamName, channel, firstName, lastName, emailAddress string) error {
	values := url.Values{
		"email":      {emailAddress},
		"channels":   {channel},
		"first_name": {firstName},
		"last_name":  {lastName},
		"restricted": {"1"},
		"token":      {api.token},
		"resend":     {"true"},
		"set_active": {"true"},
		"_attempts":  {"1"},
	}

	err := api.adminRequest(ctx, "invite", teamName, values)
	if err != nil {
		return fmt.Errorf("Failed to restricted account: %s", err)
	}

	return nil
}

// InviteToTeam invites a user to a Slack team
func (api *Client) InviteToTeam(teamName, firstName, lastName, emailAddress string) error {
	return api.InviteToTeamContext(context.Background(), teamName, firstName, lastName, emailAddress)
}

// InviteToTeamContext invites a user to a Slack team with a custom context
func (api *Client) InviteToTeamContext(ctx context.Context, teamName, firstName, lastName, emailAddress string) error {
	values := url.Values{
		"email":      {emailAddress},
		"first_name": {firstName},
		"last_name":  {lastName},
		"token":      {api.token},
		"set_active": {"true"},
		"_attempts":  {"1"},
	}

	err := api.adminRequest(ctx, "invite", teamName, values)
	if err != nil {
		return fmt.Errorf("Failed to invite to team: %s", err)
	}

	return nil
}

// SetRegular enables the specified user
func (api *Client) SetRegular(teamName, user string) error {
	return api.SetRegularContext(context.Background(), teamName, user)
}

// SetRegularContext enables the specified user with a custom context
func (api *Client) SetRegularContext(ctx context.Context, teamName, user string) error {
	values := url.Values{
		"user":       {user},
		"token":      {api.token},
		"set_active": {"true"},
		"_attempts":  {"1"},
	}

	err := api.adminRequest(ctx, "setRegular", teamName, values)
	if err != nil {
		return fmt.Errorf("Failed to change the user (%s) to a regular user: %s", user, err)
	}

	return nil
}

// SendSSOBindingEmail sends an SSO binding email to the specified user
func (api *Client) SendSSOBindingEmail(teamName, user string) error {
	return api.SendSSOBindingEmailContext(context.Background(), teamName, user)
}

// SendSSOBindingEmailContext sends an SSO binding email to the specified user with a custom context
func (api *Client) SendSSOBindingEmailContext(ctx context.Context, teamName, user string) error {
	values := url.Values{
		"user":       {user},
		"token":      {api.token},
		"set_active": {"true"},
		"_attempts":  {"1"},
	}

	err := api.adminRequest(ctx, "sendSSOBind", teamName, values)
	if err != nil {
		return fmt.Errorf("Failed to send SSO binding email for user (%s): %s", user, err)
	}

	return nil
}

// SetUltraRestricted converts a user into a single-channel guest
func (api *Client) SetUltraRestricted(teamName, uid, channel string) error {
	return api.SetUltraRestrictedContext(context.Background(), teamName, uid, channel)
}

// SetUltraRestrictedContext converts a user into a single-channel guest with a custom context
func (api *Client) SetUltraRestrictedContext(ctx context.Context, teamName, uid, channel string) error {
	values := url.Values{
		"user":       {uid},
		"channel":    {channel},
		"token":      {api.token},
		"set_active": {"true"},
		"_attempts":  {"1"},
	}

	err := api.adminRequest(ctx, "setUltraRestricted", teamName, values)
	if err != nil {
		return fmt.Errorf("Failed to ultra-restrict account: %s", err)
	}

	return nil
}

// SetRestricted converts a user into a restricted account
func (api *Client) SetRestricted(teamName, uid string, channelIds ...string) error {
	return api.SetRestrictedContext(context.Background(), teamName, uid, channelIds...)
}

// SetRestrictedContext converts a user into a restricted account with a custom context
func (api *Client) SetRestrictedContext(ctx context.Context, teamName, uid string, channelIds ...string) error {
	values := url.Values{
		"user":       {uid},
		"token":      {api.token},
		"set_active": {"true"},
		"_attempts":  {"1"},
		"channels":   {strings.Join(channelIds, ",")},
	}

	err := api.adminRequest(ctx, "setRestricted", teamName, values)
	if err != nil {
		return fmt.Errorf("failed to restrict account: %s", err)
	}

	return nil
}

// AdminConversationsSetTeamsParams contains arguments for AdminConversationsSetTeams.
type AdminConversationsSetTeamsParams struct {
	ChannelID     string
	OrgChannel    *bool
	TargetTeamIDs []string
	TeamID        *string
}

// Set the workspaces in an Enterprise Grid organisation that connect to a public or
// private channel.
// See: https://api.slack.com/methods/admin.conversations.setTeams
func (api *Client) AdminConversationsSetTeams(ctx context.Context, params AdminConversationsSetTeamsParams) error {
	values := url.Values{
		"token":      {api.token},
		"channel_id": {params.ChannelID},
	}

	if params.OrgChannel != nil {
		values.Add("org_channel", strconv.FormatBool(*params.OrgChannel))
	}

	if len(params.TargetTeamIDs) > 0 {
		values.Add("target_team_ids", strings.Join(params.TargetTeamIDs, ",")) // ["T123", "T456"] - > "T123,T456"
	}

	if params.TeamID != nil {
		values.Add("team_id", *params.TeamID)
	}

	response := &SlackResponse{}
	err := api.postMethod(ctx, "admin.conversations.setTeams", values, response)
	if err != nil {
		return err
	}

	return response.Err()
}

// AdminConversationsConvertToPrivate converts a public channel to a private channel.
// To do this, you must have the admin.conversations:write scope.
// See: https://api.slack.com/methods/admin.conversations.convertToPrivate
func (api *Client) AdminConversationsConvertToPrivate(ctx context.Context, channelID string) error {
	values := url.Values{
		"token":      []string{api.token},
		"channel_id": []string{channelID},
	}

	response := &SlackResponse{}
	err := api.postMethod(ctx, "admin.conversations.convertToPrivate", values, response)
	if err != nil {
		return err
	}

	return response.Err()
}

// AdminConversationsConvertToPublic converts a private channel to a public channel.
// To do this, you must have the admin.conversations:write scope.
// See: https://api.slack.com/methods/admin.conversations.convertToPublic
func (api *Client) AdminConversationsConvertToPublic(ctx context.Context, channelID string) error {
	values := url.Values{
		"token":      []string{api.token},
		"channel_id": []string{channelID},
	}

	response := &SlackResponse{}
	err := api.postMethod(ctx, "admin.conversations.convertToPublic", values, response)
	if err != nil {
		return err
	}

	return response.Err()
}
