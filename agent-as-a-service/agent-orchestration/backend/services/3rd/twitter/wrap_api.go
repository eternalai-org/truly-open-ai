package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func NewTwitterWrapClient(accessToken string) *Client {
	return &Client{
		user: &User{
			Authorizer: authorize{
				Token: accessToken,
			},
			Client: http.DefaultClient,
			Host:   "https://api.x.com",
		},
		AppToken: accessToken,
	}
}

// get user info
func (c *Client) GetTwitterByUserName(userName string) (*UserObj, error) {
	var apiResponse TwitterMeResp
	url := fmt.Sprintf("https://api.x.com/2/users/by/username/%s?user.fields=profile_image_url,public_metrics,verified,description,entities", userName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.AppToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to search tweets, status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse.UserObj, nil
}

func (c *Client) GetTwitterByID(userTwitterID string) (*UserObj, error) {
	var apiResponse TwitterMeResp
	url := fmt.Sprintf("https://api.x.com/2/users/%s?user.fields=public_metrics,verified,profile_image_url,created_at,verified_type", userTwitterID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.AppToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to search tweets, status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse.UserObj, nil
}

func (c *Client) GetListFollowing(userTwitterID, paginationToken, accessToken string) (*UserFollowLookup, error) {
	fieldOpts := UserFollowOptions{
		TweetFields: []TweetField{
			TweetFieldCreatedAt,
			// TweetFieldContextAnnotations,
		},
		UserFields: []UserField{
			UserFieldDescription,
			UserFieldPublicMetrics,
			UserFieldCreatedAt, UserFieldVerified,
		},
		MaxResults:      100,
		PaginationToken: paginationToken,
	}
	c.user = &User{
		Authorizer: authorize{
			Token: accessToken,
		},
		Client: http.DefaultClient,
		Host:   "https://api.x.com",
	}

	userFollowLookup, err := c.user.LookupFollowing(context.Background(), userTwitterID, fieldOpts)
	if err != nil {
		return nil, err
	}

	return userFollowLookup, nil
}

func (c *Client) LookupUsername(accessToken string, usernames []string) (*UserLookups, error) {
	fieldOpts := UserFieldOptions{
		UserFields: []UserField{
			UserFieldDescription,
			UserFieldPublicMetrics,
			UserFieldCreatedAt,
			UserFieldVerified,
			UserFieldID,
			UserFieldName,
			UserFieldEntities,
			UserFieldUserName,
			UserFieldProfileImageURL,
		},
	}
	c.user = &User{
		Authorizer: authorize{
			Token: accessToken,
		},
		Client: http.DefaultClient,
		Host:   "https://api.x.com",
	}

	userFollowLookup, err := c.user.LookupUsername(context.Background(), usernames, fieldOpts)
	if err != nil {
		return nil, err
	}

	return &userFollowLookup, nil
}

func (c *Client) GetListUserTweets(userTwitterID, paginationToken, accessToken string, maxResults int) (*UserTimeline, error) {
	c.user = &User{
		Authorizer: authorize{
			Token: accessToken,
		},
		Client: http.DefaultClient,
		Host:   "https://api.x.com",
	}

	tweetOpts := UserTimelineOpts{
		TweetFields: []TweetField{
			TweetFieldAttachments,
			TweetFieldAuthorID,
			TweetFieldConversationID,
			TweetFieldCreatedAt,
			TweetFieldEntities,
			TweetFieldID,
			TweetFieldInReplyToUserID,
			TweetFieldPublicMetrics,
			TweetFieldReferencedTweets,
			TweetFieldSource,
			TweetFieldText,
			TweetFieldNoteText,
		},
		UserFields: []UserField{
			UserFieldCreatedAt,
			UserFieldDescription,
			UserFieldEntities,
			UserFieldName,
			UserFieldUserName,
		},
		Expansions: []Expansion{
			ExpansionAuthorID,
			// ExpansionReferencedTweetsID,
			// ExpansionReferencedTweetsIDAuthorID,
			ExpansionAttachmentsMediaKeys,
			ExpansionInReplyToUserID,
		},
		Excludes: []Exclude{
			ExcludeReplies,
			ExcludeRetweets,
		},
		MediaFields: []MediaField{
			MediaFieldMediaKey,
			MediaFieldURL,
			MediaFieldType,
		},
		MaxResults:      maxResults,
		PaginationToken: paginationToken,
	}

	userTweets, err := c.user.Tweets(context.Background(), *&userTwitterID, tweetOpts)

	if err != nil {
		return nil, err
	}

	return userTweets, nil

}

func (c *Client) GetAllUserTweets(userTwitterID, paginationToken, accessToken string, maxResults int) (*UserTimeline, error) {
	c.user = &User{
		Authorizer: authorize{
			Token: accessToken,
		},
		Client: http.DefaultClient,
		Host:   "https://api.x.com",
	}

	tweetOpts := UserTimelineOpts{
		TweetFields: []TweetField{
			TweetFieldAttachments,
			TweetFieldAuthorID,
			TweetFieldConversationID,
			TweetFieldCreatedAt,
			TweetFieldEntities,
			TweetFieldID,
			TweetFieldInReplyToUserID,
			TweetFieldPublicMetrics,
			TweetFieldReferencedTweets,
			TweetFieldSource,
			TweetFieldText,
			TweetFieldNoteText,
		},
		UserFields: []UserField{
			UserFieldCreatedAt,
			UserFieldDescription,
			UserFieldEntities,
			UserFieldName,
			UserFieldUserName,
		},
		Expansions: []Expansion{
			ExpansionAuthorID,
			ExpansionAttachmentsMediaKeys,
			ExpansionInReplyToUserID,
		},
		MediaFields: []MediaField{
			MediaFieldMediaKey,
			MediaFieldURL,
			MediaFieldType,
		},
		MaxResults:      maxResults,
		PaginationToken: paginationToken,
	}

	userTweets, err := c.user.Tweets(context.Background(), *&userTwitterID, tweetOpts)

	if err != nil {
		return nil, err
	}

	return userTweets, nil

}

func (c *Client) LookupUserTweets(accessToken string, tweetIDs []string) (*TweetLookups, error) {
	c.user = &User{
		Authorizer: authorize{
			Token: accessToken,
		},
		Client: http.DefaultClient,
		Host:   "https://api.x.com",
	}

	fieldOpts := TweetFieldOptions{
		Expansions: []Expansion{
			ExpansionAuthorID,
			// ExpansionReferencedTweetsID,
			// ExpansionReferencedTweetsIDAuthorID,
			ExpansionAttachmentsMediaKeys,
			ExpansionInReplyToUserID,
		},
		TweetFields: []TweetField{
			TweetFieldAttachments,
			TweetFieldAuthorID,
			// TweetFieldContextAnnotations,
			TweetFieldConversationID,
			TweetFieldCreatedAt,
			TweetFieldEntities,
			TweetFieldID,
			TweetFieldInReplyToUserID,
			TweetFieldReferencedTweets,
			TweetFieldText,
			TweetFieldNoteText,
		},
		MediaFields: []MediaField{
			MediaFieldMediaKey,
			MediaFieldURL,
			MediaFieldType,
		},
	}

	tweet := &Tweet{
		Authorizer: authorize{
			Token: accessToken,
		},
		Client: http.DefaultClient,
		Host:   "https://api.x.com",
	}

	lookups, err := tweet.Lookup(context.Background(), tweetIDs, fieldOpts)

	if err != nil {
		return nil, err
	}

	return &lookups, nil
}

func (c *Client) GetListUserMentions(userTwitterID, paginationToken, accessToken string, maxResults int) (*UserTimeline, error) {
	c.user = &User{
		Authorizer: authorize{
			Token: accessToken,
		},
		Client: http.DefaultClient,
		Host:   "https://api.x.com",
	}

	tweetOpts := UserTimelineOpts{
		TweetFields: []TweetField{
			TweetFieldAuthorID,
			TweetFieldConversationID,
			TweetFieldCreatedAt,
			TweetFieldID,
			TweetFieldInReplyToUserID,
			TweetFieldReferencedTweets,
			TweetFieldText,
			TweetFieldEntities,
			TweetFieldNoteText,
		},
		UserFields: []UserField{
			UserFieldCreatedAt,
			UserFieldDescription,
			UserFieldEntities,
			UserFieldName,
			UserFieldUserName,
			UserFieldVerified,
		},
		Expansions: []Expansion{
			ExpansionAuthorID,
			// ExpansionReferencedTweetsID,
			// ExpansionReferencedTweetsIDAuthorID,
			ExpansionAttachmentsMediaKeys,
			ExpansionInReplyToUserID,
		},
		MediaFields: []MediaField{
			MediaFieldMediaKey,
			MediaFieldURL,
			MediaFieldType,
		},
		MaxResults: maxResults,
	}
	if paginationToken != "" {
		tweetOpts.PaginationToken = paginationToken
	}

	userTweets, err := c.user.Mentions(context.Background(), *&userTwitterID, tweetOpts)

	if err != nil {
		return nil, err
	}

	return userTweets, nil
}

func (c *Client) SearchRecentTweet(query, paginationToken, accessToken string, maxResult int) (*TweetRecentSearch, error) {
	tweet := &Tweet{
		Authorizer: authorize{
			Token: accessToken,
		},
		Client: http.DefaultClient,
		Host:   "https://api.x.com",
	}
	fieldOpts := TweetFieldOptions{
		TweetFields: []TweetField{
			TweetFieldAuthorID,
			TweetFieldConversationID,
			TweetFieldCreatedAt,
			TweetFieldID,
			TweetFieldInReplyToUserID,
			TweetFieldPublicMetrics,
			TweetFieldReferencedTweets,
			TweetFieldSource,
			TweetFieldText,
			TweetFieldNoteText,
		},
		UserFields: []UserField{
			UserFieldCreatedAt,
			UserFieldDescription,
			UserFieldEntities,
			UserFieldLocation,
			UserFieldName,
			UserFieldProfileImageURL,
			UserFieldURL,
			UserFieldUserName,
			UserFieldVerified,
		},
		Expansions: []Expansion{
			ExpansionAuthorID,
			ExpansionAttachmentsMediaKeys,
			ExpansionInReplyToUserID,
			ExpansionGeoPlaceID,
		},
		MediaFields: []MediaField{
			MediaFieldMediaKey,
			MediaFieldURL,
			MediaFieldType,
		},
	}
	searchOpts := TweetRecentSearchOptions{
		MaxResult: maxResult,
		NextToken: paginationToken,
	}

	recentSearch, err := tweet.RecentSearch(context.Background(), query, searchOpts, fieldOpts)

	if err != nil {
		return nil, err
	}

	return recentSearch, nil
}

func (c *Client) SearchRecentTweetV1(query, sinceID string, accessToken string, maxResult int) (*TweetRecentSearch, error) {
	tweet := &Tweet{
		Authorizer: authorize{
			Token: accessToken,
		},
		Client: http.DefaultClient,
		Host:   "https://api.x.com",
	}
	fieldOpts := TweetFieldOptions{
		TweetFields: []TweetField{
			TweetFieldAuthorID,
			TweetFieldConversationID,
			TweetFieldCreatedAt,
			TweetFieldID,
			TweetFieldInReplyToUserID,
			TweetFieldPublicMetrics,
			TweetFieldReferencedTweets,
			TweetFieldSource,
			TweetFieldText,
			TweetFieldNoteText,
		},
		UserFields: []UserField{
			UserFieldCreatedAt,
			UserFieldDescription,
			UserFieldEntities,
			UserFieldLocation,
			UserFieldName,
			UserFieldProfileImageURL,
			UserFieldURL,
			UserFieldUserName,
			UserFieldVerified,
		},
		Expansions: []Expansion{
			ExpansionAuthorID,
			ExpansionAttachmentsMediaKeys,
			ExpansionInReplyToUserID,
			ExpansionGeoPlaceID,
		},
		MediaFields: []MediaField{
			MediaFieldMediaKey,
			MediaFieldURL,
			MediaFieldType,
		},
	}
	searchOpts := TweetRecentSearchOptions{
		MaxResult: maxResult,
		SinceID:   sinceID,
	}

	recentSearch, err := tweet.RecentSearch(context.Background(), query, searchOpts, fieldOpts)

	if err != nil {
		return nil, err
	}

	return recentSearch, nil
}

type UserSearchResp struct {
	UserObj []*UserObj `json:"data"`
}

func (c *Client) SearchUsers(query, paginationToken, accessToken string) ([]*UserObj, error) {
	var apiResponse UserSearchResp
	params := url.Values{}
	params.Add("query", query)
	params.Add("user.fields", "id,username,name,public_metrics,verified,profile_image_url,created_at")
	if paginationToken != "" {
		params.Add("next_token", paginationToken)
	}
	url := fmt.Sprintf("https://api.x.com/2/users/search?%s", params.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		e := &TweetErrorResponse{}
		if err := decoder.Decode(e); err != nil {
			return nil, &HTTPError{
				Status:     resp.Status,
				StatusCode: resp.StatusCode,
				URL:        resp.Request.URL.String(),
			}
		}
		e.StatusCode = resp.StatusCode
		return nil, e
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}
	return apiResponse.UserObj, nil
}

type TrendsResp struct {
	TrendName  string `json:"trend_name"`
	TweetCount uint   `json:"tweet_count"`
}

func (c *Client) LookupTweetsByID(accessToken string, tweetID string) (*TweetObj, error) {
	resp := TweetObj{}
	mapResp, err := c.LookupUserTweets(accessToken, []string{tweetID})
	if err != nil {
		return nil, err
	}
	for k, v := range *mapResp {
		if strings.EqualFold(k, tweetID) {
			resp = v.Tweet
		}
	}
	return &resp, err
}
