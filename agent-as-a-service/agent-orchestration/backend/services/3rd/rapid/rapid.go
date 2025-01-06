package rapid

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Rapid struct {
	url    string
	host   string
	apiKey string
}

type RapidUser struct {
	ID uint `json:"id"`
}

type RapidTweetSummaryInfo struct {
	Address               string
	TweetID               string
	ParentTweetID         string
	TwitterID             string
	TwitterName           string
	TwitterUsername       string
	TwitterAvatar         string
	TwitterFollowersCount uint
	FavoriteCount         int
	BookmarkCount         int
	QuoteCount            int
	ReplyCount            int
	RetweetCount          int
	ViewCount             string
	FullText              string
	PostedAt              time.Time
	UserMentions          []string
}

func (m *RapidTweetSummaryInfo) IsMention(username string) bool {
	for _, v := range m.UserMentions {
		if strings.EqualFold(v, username) {
			return true
		}
	}
	return false
}

type RapidSearchResponse struct {
	Data *struct {
		SearchByRawQuery *struct {
			SearchTimeline *struct {
				Timeline *struct {
					Instructions []*struct {
						Type    string `json:"type"`
						Entries []*struct {
							Content *struct {
								EntryType   string `json:"entryType,omitempty"`
								CursorType  string `json:"cursorType,omitempty"`
								Value       string `json:"value,omitempty"`
								ItemContent *struct {
									TweetResults *struct {
										Result *struct {
											TypeName string `json:"__typename,omitempty"`
											Views    *struct {
												Count string `json:"count,omitempty"`
											} `json:"views,omitempty"`
											Core *struct {
												UserResults *struct {
													Results *struct {
														TwitterID      string `json:"rest_id,omitempty"`
														IsBlueVerified bool   `json:"is_blue_verified,omitempty"`
														Legacy         *struct {
															Name           string `json:"name,omitempty"`
															Username       string `json:"screen_name,omitempty"`
															ProfileUrl     string `json:"profile_image_url_https,omitempty"`
															FollowersCount uint   `json:"followers_count,omitempty"`
														} `json:"legacy,omitempty"`
													} `json:"result,omitempty"`
												} `json:"user_results,omitempty"`
											} `json:"core,omitempty"`
											Legacy *struct {
												Entities *struct {
													Hashtags []*struct {
														Text string `json:"text,omitempty"`
													} `json:"hashtags,omitempty"`
													UserMentions []*struct {
														ID         string `json:"id_str,omitempty"`
														ScreenName string `json:"screen_name,omitempty"`
													} `json:"user_mentions,omitempty"`
												} `json:"entities,omitempty"`
												UserID              string `json:"user_id_str,omitempty"`
												ID                  string `json:"id_str,omitempty"`
												CreatedAt           string `json:"created_at,omitempty"`
												FavoriteCount       int    `json:"favorite_count,omitempty"`
												QuoteCount          int    `json:"quote_count,omitempty"`
												ReplyCount          int    `json:"reply_count,omitempty"`
												RetweetCount        int    `json:"retweet_count,omitempty"`
												FullText            string `json:"full_text,omitempty"`
												InReplyToScreenName string `json:"in_reply_to_screen_name,omitempty"`
												InReplyToStatusID   string `json:"in_reply_to_status_id_str,omitempty"`
												InReplyToUserID     string `json:"in_reply_to_user_id_str,omitempty"`
											} `json:"legacy,omitempty"`
										} `json:"result,omitempty"`
									} `json:"tweet_results,omitempty"`
								} `json:"itemContent,omitempty"`
							} `json:"content,omitempty"`
						} `json:"entries,omitempty"`
					} `json:"instructions"`
				} `json:"timeline"`
			} `json:"search_timeline"`
		} `json:"search_by_raw_query"`
	} `json:"data"`
}

type TweetContent struct {
	ItemType     string `json:"itemType"`
	TweetResults *struct {
		Result *struct {
			RestID string `json:"rest_id"`
			Core   *struct {
				UserResults *struct {
					Result *struct {
						ID                         string   `json:"id"`
						RestID                     string   `json:"rest_id"`
						AffiliatesHighlightedLabel struct{} `json:"affiliates_highlighted_label"`
						HasGraduatedAccess         bool     `json:"has_graduated_access"`
						IsBlueVerified             bool     `json:"is_blue_verified"`
						ProfileImageShape          string   `json:"profile_image_shape"`
						Legacy                     *struct {
							CanDM                bool     `json:"can_dm"`
							CanMediaTag          bool     `json:"can_media_tag"`
							CreatedAt            string   `json:"created_at"`
							DefaultProfile       bool     `json:"default_profile"`
							DefaultProfileImage  bool     `json:"default_profile_image"`
							Description          string   `json:"description"`
							FastFollowersCount   int      `json:"fast_followers_count"`
							FavouritesCount      int      `json:"favourites_count"`
							FollowersCount       int      `json:"followers_count"`
							FriendsCount         int      `json:"friends_count"`
							ListedCount          int      `json:"listed_count"`
							MediaCount           int      `json:"media_count"`
							Name                 string   `json:"name"`
							ScreenName           string   `json:"screen_name"`
							StatusesCount        int      `json:"statuses_count"`
							Verified             bool     `json:"verified"`
							WithheldInCountries  []string `json:"withheld_in_countries"`
							ProfileImageURLHTTPS string   `json:"profile_image_url_https"`
						} `json:"legacy"`
					} `json:"result"`
				} `json:"user_results"`
			} `json:"core"`
			Card *struct {
				RestID string `json:"rest_id"`
				Legacy struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"legacy"`
			} `json:"card"`
			Legacy *struct {
				BookmarkCount     int    `json:"bookmark_count"`
				Bookmarked        bool   `json:"bookmarked"`
				CreatedAt         string `json:"created_at"`
				ConversationIDStr string `json:"conversation_id_str"`
				DisplayTextRange  []int  `json:"display_text_range"`
				Entities          *struct {
					UserMentions []struct {
						IDStr      string `json:"id_str"`
						Name       string `json:"name"`
						ScreenName string `json:"screen_name"`
						Indices    []int  `json:"indices"`
					} `json:"user_mentions"`
					URLs []struct {
						DisplayURL  string `json:"display_url"`
						ExpandedURL string `json:"expanded_url"`
						URL         string `json:"url"`
						Indices     []int  `json:"indices"`
					} `json:"urls"`
					Hashtags []struct {
						Indices []int  `json:"indices"`
						Text    string `json:"text"`
					} `json:"hashtags"`
					Symbols []struct{} `json:"symbols"`
				} `json:"entities"`
				FavoriteCount             int    `json:"favorite_count"`
				Favorited                 bool   `json:"favorited"`
				FullText                  string `json:"full_text"`
				IsQuoteStatus             bool   `json:"is_quote_status"`
				Lang                      string `json:"lang"`
				PossiblySensitive         bool   `json:"possibly_sensitive"`
				PossiblySensitiveEditable bool   `json:"possibly_sensitive_editable"`
				QuoteCount                int    `json:"quote_count"`
				ReplyCount                int    `json:"reply_count"`
				RetweetCount              int    `json:"retweet_count"`
				Retweeted                 bool   `json:"retweeted"`
				UserIDStr                 string `json:"user_id_str"`
				IDStr                     string `json:"id_str"`
			} `json:"legacy"`
		} `json:"result"`
	} `json:"tweet_results"`
}
type RapidTweetResponse struct {
	Data *struct {
		User *struct {
			Result *struct {
				TimelineV2 *struct {
					Timeline *struct {
						Instructions []struct {
							Type    string `json:"type"`
							Entries []struct {
								EntryId   string `json:"entryId"`
								SortIndex string `json:"sortIndex"`
								Content   struct {
									EntryType  string `json:"entryType"`
									CursorType string `json:"cursorType,omitempty"`
									Value      string `json:"value,omitempty"`
									Items      []struct {
										EntryId string `json:"entryId"`
										Item    struct {
											ItemContent *TweetContent `json:"itemContent,omitempty"`
										} `json:"item"`
									} `json:"items,omitempty"`
									ItemContent *TweetContent `json:"itemContent,omitempty"`
								} `json:"content"`
							} `json:"entries"`
						} `json:"instructions"`
					} `json:"timeline"`
				} `json:"timeline_v2"`
			} `json:"result"`
		} `json:"user"`
	} `json:"data"`
}

func NewRapid(apiKey string) *Rapid {
	return &Rapid{
		url:    "https://twitter135.p.rapidapi.com/v2",
		host:   "twitter135.p.rapidapi.com",
		apiKey: apiKey,
	}
}

func (m *Rapid) request(fullUrl string, method string, headers map[string]string, reqBody io.Reader) ([]byte, int, error) {

	req, err := http.NewRequest(method, fullUrl, reqBody)
	if err != nil {
		return nil, 0, err
	}

	if len(headers) > 0 {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", m.apiKey)
	req.Header.Add("X-RapidAPI-Host", m.host)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}
	return body, res.StatusCode, nil
}

func (m *Rapid) SearchSpreadTwitterByUserName(cursor, twitterUserName string) ([]*RapidTweetSummaryInfo, string, error) {
	responses := []*RapidTweetSummaryInfo{}
	path := fmt.Sprintf("https://twitter135.p.rapidapi.com/Search/?q=(@%s)&cursor=%s", twitterUserName, cursor)
	fmt.Println(path)
	data, _, err := m.request(path, "GET", nil, nil)
	if err != nil {
		return responses, "", err
	}
	fmt.Println(string(data))
	res := &RapidSearchResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return responses, "", err
	}
	numContent := uint(0)
	cursorTop := ""
	if res != nil && res.Data != nil && res.Data.SearchByRawQuery != nil &&
		res.Data.SearchByRawQuery.SearchTimeline != nil && res.Data.SearchByRawQuery.SearchTimeline.Timeline != nil &&
		res.Data.SearchByRawQuery.SearchTimeline.Timeline.Instructions != nil {
		for _, v := range res.Data.SearchByRawQuery.SearchTimeline.Timeline.Instructions {
			if v.Type == "TimelineAddEntries" {
				for _, item := range v.Entries {
					if item.Content.EntryType == "TimelineTimelineCursor" && item.Content.CursorType == "Top" {
						cursorTop = item.Content.Value
					}
					if item.Content.EntryType == "TimelineTimelineItem" {
						if item.Content != nil && item.Content.ItemContent != nil && item.Content.ItemContent.TweetResults != nil && item.Content.ItemContent.TweetResults.Result != nil && item.Content.ItemContent.TweetResults.Result.Legacy != nil && item.Content.ItemContent.TweetResults.Result.Core != nil && item.Content.ItemContent.TweetResults.Result.Core.UserResults != nil && item.Content.ItemContent.TweetResults.Result.Core.UserResults.Results != nil && item.Content.ItemContent.TweetResults.Result.Core.UserResults.Results.Legacy != nil {
							legacy := item.Content.ItemContent.TweetResults.Result.Legacy
							legacyUser := item.Content.ItemContent.TweetResults.Result.Core.UserResults.Results.Legacy
							isMention := false
							for _, t := range legacy.Entities.UserMentions {
								if strings.EqualFold(t.ScreenName, twitterUserName) {
									isMention = true
								}
							}
							if isMention {
								numContent += 1
								layout := "Mon Jan 02 15:04:05 -0700 2006"
								t, err := time.Parse(layout, legacy.CreatedAt)
								if err == nil {

									responses = append(responses, &RapidTweetSummaryInfo{
										TweetID:         legacy.ID,
										TwitterID:       legacy.UserID,
										ParentTweetID:   legacy.InReplyToStatusID,
										TwitterName:     legacyUser.Name,
										TwitterUsername: legacyUser.Username,
										TwitterAvatar:   legacyUser.ProfileUrl,
										QuoteCount:      legacy.QuoteCount,
										ReplyCount:      legacy.ReplyCount,
										RetweetCount:    legacy.RetweetCount,
										FullText:        legacy.FullText,
										PostedAt:        t,
									})

								}
							}
						}
					}
				}
			}
		}
	}

	return responses, cursorTop, nil
}

type RapidTweetDetailResponse struct {
	Data *struct {
		Conversation *struct {
			Instructions []*struct {
				Type    string `json:"type"`
				Entries []*struct {
					EntryId string `json:"entryId"`
					Content *struct {
						EntryType   string `json:"entryType,omitempty"`
						CursorType  string `json:"cursorType,omitempty"`
						Value       string `json:"value,omitempty"`
						ItemContent *struct {
							TweetResults *struct {
								Result *struct {
									TypeName string `json:"__typename,omitempty"`
									Views    *struct {
										Count string `json:"count,omitempty"`
									} `json:"views,omitempty"`
									Core *struct {
										UserResults *struct {
											Results *struct {
												TwitterID      string `json:"rest_id,omitempty"`
												IsBlueVerified bool   `json:"is_blue_verified,omitempty"`
												Legacy         *struct {
													Name           string `json:"name,omitempty"`
													Username       string `json:"screen_name,omitempty"`
													ProfileUrl     string `json:"profile_image_url_https,omitempty"`
													FollowersCount uint   `json:"followers_count,omitempty"`
												} `json:"legacy,omitempty"`
											} `json:"result,omitempty"`
										} `json:"user_results,omitempty"`
									} `json:"core,omitempty"`
									Legacy *struct {
										Entities *struct {
											Hashtags []*struct {
												Text string `json:"text,omitempty"`
											} `json:"hashtags,omitempty"`
											UserMentions []*struct {
												ID         string `json:"id_str,omitempty"`
												ScreenName string `json:"screen_name,omitempty"`
											} `json:"user_mentions,omitempty"`
										} `json:"entities,omitempty"`
										UserID              string `json:"user_id_str,omitempty"`
										ID                  string `json:"id_str,omitempty"`
										CreatedAt           string `json:"created_at,omitempty"`
										FavoriteCount       int    `json:"favorite_count,omitempty"`
										QuoteCount          int    `json:"quote_count,omitempty"`
										ReplyCount          int    `json:"reply_count,omitempty"`
										RetweetCount        int    `json:"retweet_count,omitempty"`
										BookmarkCount       int    `json:"bookmark_count,omitempty"`
										FullText            string `json:"full_text,omitempty"`
										InReplyToScreenName string `json:"in_reply_to_screen_name,omitempty"`
										InReplyToStatusID   string `json:"in_reply_to_status_id_str,omitempty"`
										InReplyToUserID     string `json:"in_reply_to_user_id_str,omitempty"`
									} `json:"legacy,omitempty"`
								} `json:"result,omitempty"`
							} `json:"tweet_results,omitempty"`
						} `json:"itemContent,omitempty"`
					} `json:"content,omitempty"`
				} `json:"entries,omitempty"`
			} `json:"instructions"`
		} `json:"threaded_conversation_with_injections_v2"`
	} `json:"data"`
}

func (m Rapid) GetTweetDetailByID(tweetID string) (*RapidTweetSummaryInfo, error) {
	path := fmt.Sprintf("%s/TweetDetail/?id=%s", m.url, tweetID)
	data, _, err := m.request(path, "GET", nil, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))
	resp := &RapidTweetDetailResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	if resp != nil && resp.Data != nil && resp.Data.Conversation != nil &&
		resp.Data.Conversation.Instructions != nil && len(resp.Data.Conversation.Instructions) > 0 {
		for _, v := range resp.Data.Conversation.Instructions {
			if v.Type == "TimelineAddEntries" {
				for _, item := range v.Entries {
					if item.Content.EntryType == "TimelineTimelineItem" {
						if item.Content != nil && item.Content.ItemContent != nil &&
							item.EntryId == fmt.Sprintf("tweet-%s", tweetID) &&
							item.Content.ItemContent.TweetResults != nil && item.Content.ItemContent.TweetResults.Result != nil &&
							item.Content.ItemContent.TweetResults.Result.Legacy != nil && item.Content.ItemContent.TweetResults.Result.Core != nil &&
							item.Content.ItemContent.TweetResults.Result.Core.UserResults != nil &&
							item.Content.ItemContent.TweetResults.Result.Core.UserResults.Results != nil &&
							item.Content.ItemContent.TweetResults.Result.Core.UserResults.Results.Legacy != nil {
							legacy := item.Content.ItemContent.TweetResults.Result.Legacy
							legacyUser := item.Content.ItemContent.TweetResults.Result.Core.UserResults.Results.Legacy
							views := item.Content.ItemContent.TweetResults.Result.Views

							viewCount := "0"
							if views != nil {
								viewCount = views.Count
							}
							layout := "Mon Jan 02 15:04:05 -0700 2006"
							t, err := time.Parse(layout, legacy.CreatedAt)
							if err == nil {
								userMentions := []string{}
								if legacy != nil && legacy.Entities != nil && legacy.Entities.UserMentions != nil {
									for _, v := range legacy.Entities.UserMentions {
										userMentions = append(userMentions, v.ScreenName)
									}
								}
								return &RapidTweetSummaryInfo{
									TweetID:               legacy.ID,
									ParentTweetID:         legacy.InReplyToStatusID,
									TwitterID:             legacy.UserID,
									TwitterName:           legacyUser.Name,
									TwitterUsername:       legacyUser.Username,
									TwitterAvatar:         legacyUser.ProfileUrl,
									TwitterFollowersCount: legacyUser.FollowersCount,
									QuoteCount:            legacy.QuoteCount,
									ReplyCount:            legacy.ReplyCount,
									RetweetCount:          legacy.RetweetCount,
									FavoriteCount:         legacy.FavoriteCount,
									BookmarkCount:         legacy.BookmarkCount,
									ViewCount:             viewCount,
									FullText:              legacy.FullText,
									PostedAt:              t,
									UserMentions:          userMentions,
								}, nil
							}
						}
					}
				}
			}
		}
	}
	return nil, nil
}

func (m *Rapid) GetTweetByTwitterID(twitterID, cursor string) ([]*TweetContent, string, error) {
	path := fmt.Sprintf("https://twitter135.p.rapidapi.com/v2/UserTweets/?id=%s&cursor=%s", twitterID, cursor)
	fmt.Println(path)
	data, _, err := m.request(path, "GET", nil, nil)
	if err != nil {
		return nil, "", err
	}
	res := &RapidTweetResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, "", err
	}
	cursorTop := ""
	listTweetContent := []*TweetContent{}
	if res != nil && res.Data != nil {
		for _, v := range res.Data.User.Result.TimelineV2.Timeline.Instructions {
			if v.Type == "TimelineAddEntries" {
				for _, item := range v.Entries {
					if item.Content.EntryType == "TimelineTimelineCursor" && item.Content.CursorType == "Top" {
						cursorTop = item.Content.Value
					} else {
						if item.Content.ItemContent != nil {
							listTweetContent = append(listTweetContent, item.Content.ItemContent)
						} else if len(item.Content.Items) > 0 {
							for _, child := range item.Content.Items {
								listTweetContent = append(listTweetContent, child.Item.ItemContent)
							}

						}
					}
				}
			}
		}
	}

	return listTweetContent, cursorTop, nil
}

func (m *Rapid) GetTweetByTwitterIDV1(twitterID, cursor string) ([]*RapidTweetSummaryInfo, string, error) {
	path := fmt.Sprintf("https://twitter135.p.rapidapi.com/v2/UserTweets/?id=%s&cursor=%s", twitterID, cursor)
	fmt.Println(path)
	data, _, err := m.request(path, "GET", nil, nil)
	if err != nil {
		return nil, "", err
	}
	res := &RapidTweetResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, "", err
	}
	cursorTop := ""
	listTweetContent := []*TweetContent{}
	if res != nil && res.Data != nil {
		for _, v := range res.Data.User.Result.TimelineV2.Timeline.Instructions {
			if v.Type == "TimelineAddEntries" {
				for _, item := range v.Entries {
					if item.Content.EntryType == "TimelineTimelineCursor" && item.Content.CursorType == "Top" {
						cursorTop = item.Content.Value
					} else {
						if item.Content.ItemContent != nil {
							listTweetContent = append(listTweetContent, item.Content.ItemContent)
						} else if len(item.Content.Items) > 0 {
							for _, child := range item.Content.Items {
								listTweetContent = append(listTweetContent, child.Item.ItemContent)
							}

						}
					}
				}
			}
		}
	}

	resp := []*RapidTweetSummaryInfo{}
	for _, item := range listTweetContent {
		if item.TweetResults != nil && item.TweetResults.Result != nil && item.TweetResults.Result.Legacy != nil && item.TweetResults.Result.Core != nil && item.TweetResults.Result.Core.UserResults != nil && item.TweetResults.Result.Core.UserResults.Result != nil && item.TweetResults.Result.Core.UserResults.Result.Legacy != nil {
			legacy := item.TweetResults.Result.Legacy
			legacyUser := item.TweetResults.Result.Core.UserResults.Result.Legacy
			layout := "Mon Jan 02 15:04:05 -0700 2006"
			t, err := time.Parse(layout, legacy.CreatedAt)
			if err == nil {
				resp = append(resp, &RapidTweetSummaryInfo{
					TweetID:      legacy.ConversationIDStr,
					TwitterName:  legacyUser.Name,
					QuoteCount:   legacy.QuoteCount,
					ReplyCount:   legacy.ReplyCount,
					RetweetCount: legacy.RetweetCount,
					FullText:     legacy.FullText,
					PostedAt:     t,
				})
			}
		}
	}
	return resp, cursorTop, nil
}

type Following struct {
	ID              string `json:"rest_id"`
	Username        string `json:"screen_name"`
	Name            string `json:"name"`
	ProfileImageUrl string `json:"profile_image_url_https"`
	FollowersCount  int    `json:"followers_count"`
	FollowingCount  int    `json:"friends_count"`
	IsBlueVerified  bool   `json:"is_blue_verified"`
	CreatedAt       string `json:"created_at"`
}

type ResponseData struct {
	Data struct {
		User struct {
			Result struct {
				Timeline struct {
					Timeline struct {
						Instructions []struct {
							Type    string `json:"type"`
							Entries []struct {
								EntryID   string `json:"entryId"`
								SortIndex string `json:"sortIndex"`
								Content   struct {
									EntryType   string `json:"entryType"`
									Value       string `json:"value"`
									CursorType  string `json:"cursorType"`
									ItemContent struct {
										ItemType    string `json:"itemType"`
										UserResults struct {
											Result struct {
												ID             string `json:"rest_id"`
												IsBlueVerified bool   `json:"is_blue_verified"`
												Legacy         struct {
													ScreenName      string `json:"screen_name"`
													Name            string `json:"name"`
													ProfileImageUrl string `json:"profile_image_url_https"`
													FollowersCount  int    `json:"followers_count"`
													FriendsCount    int    `json:"friends_count"`
													IsVerified      bool   `json:"verified"`
													CreatedAt       string `json:"created_at"`
												} `json:"legacy"`
											} `json:"result"`
										} `json:"user_results"`
									} `json:"itemContent"`
								} `json:"content"`
							} `json:"entries"`
						} `json:"instructions"`
					} `json:"timeline"`
				} `json:"timeline"`
			} `json:"result"`
		} `json:"user"`
	} `json:"data"`
}

func (m *Rapid) GetTwitterFollowings(userID string) ([]Following, error) {
	var followings []Following
	mapFollow := map[string]Following{}
	client := &http.Client{Timeout: 10 * time.Second}

	nextCursor := ""

	for {
		url := fmt.Sprintf("https://twitter135.p.rapidapi.com/v2/Following/?id=%s&count=100", userID)
		if nextCursor != "" {
			url += "&cursor=" + nextCursor
		}
		fmt.Println(len(mapFollow))
		fmt.Println("url : " + url)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("X-RapidAPI-Key", m.apiKey)
		req.Header.Add("X-RapidAPI-Host", "twitter135.p.rapidapi.com")

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to get followings, status: %s, body: %s", resp.Status, body)
		}

		var data ResponseData
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, err
		}

		hasData := false
		// Parse each following entry
		for _, instruction := range data.Data.User.Result.Timeline.Timeline.Instructions {
			if instruction.Type == "TimelineAddEntries" {
				for _, entry := range instruction.Entries {
					if entry.Content.EntryType == "TimelineTimelineItem" && entry.Content.ItemContent.ItemType == "TimelineUser" {
						user := entry.Content.ItemContent.UserResults.Result
						legacy := user.Legacy
						following := Following{
							ID:              user.ID,
							IsBlueVerified:  user.IsBlueVerified,
							Username:        legacy.ScreenName,
							Name:            legacy.Name,
							ProfileImageUrl: legacy.ProfileImageUrl,
							FollowersCount:  legacy.FollowersCount,
							FollowingCount:  legacy.FriendsCount,
							CreatedAt:       legacy.CreatedAt,
						}
						followings = append(followings, following)
						mapFollow[following.ID] = following
						hasData = true
					}
				}
			}
		}

		// Look for the next cursor in the response
		nextCursor = ""
		for _, instruction := range data.Data.User.Result.Timeline.Timeline.Instructions {
			for _, entry := range instruction.Entries {
				if entry.Content.EntryType == "TimelineTimelineCursor" && entry.Content.CursorType == "Bottom" {
					nextCursor = entry.Content.Value
				}
			}
		}

		// Exit if there are no more pages
		if nextCursor == "" || !hasData {
			break
		}
	}

	return followings, nil
}

type TwitterDetail struct {
	TwitterID       string
	TwitterName     string
	TwitterUsername string
	TwitterAvatar   string
}

func (m Rapid) IsPostForFaucetAgent(code string) (*TwitterDetail, error) {
	path := fmt.Sprintf(`https://twitter135.p.rapidapi.com/Search/?q="%s"`, code)

	data, _, err := m.request(path, "GET", nil, nil)
	if err != nil {
		return nil, err
	}
	res := &RapidSearchResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	numContent := uint(0)

	if res != nil && res.Data != nil && res.Data.SearchByRawQuery != nil &&
		res.Data.SearchByRawQuery.SearchTimeline != nil && res.Data.SearchByRawQuery.SearchTimeline.Timeline != nil &&
		res.Data.SearchByRawQuery.SearchTimeline.Timeline.Instructions != nil {
		for _, v := range res.Data.SearchByRawQuery.SearchTimeline.Timeline.Instructions {
			if v.Type == "TimelineAddEntries" {
				for _, item := range v.Entries {
					if item.Content.EntryType == "TimelineTimelineItem" {
						if item.Content != nil && item.Content.ItemContent != nil && item.Content.ItemContent.TweetResults != nil && item.Content.ItemContent.TweetResults.Result != nil && item.Content.ItemContent.TweetResults.Result.Legacy != nil && item.Content.ItemContent.TweetResults.Result.Core != nil && item.Content.ItemContent.TweetResults.Result.Core.UserResults != nil && item.Content.ItemContent.TweetResults.Result.Core.UserResults.Results != nil && item.Content.ItemContent.TweetResults.Result.Core.UserResults.Results.Legacy != nil {
							legacy := item.Content.ItemContent.TweetResults.Result.Legacy
							legacyUser := item.Content.ItemContent.TweetResults.Result.Core.UserResults.Results.Legacy
							isMention := false
							for _, t := range legacy.Entities.UserMentions {
								if strings.EqualFold(t.ScreenName, "CryptoEternalAI") {
									isMention = true
								}
							}
							if isMention {
								numContent += 1
								if err == nil {
									if strings.Contains(legacy.FullText, code) {
										return &TwitterDetail{
											TwitterID:       legacy.UserID,
											TwitterName:     legacyUser.Name,
											TwitterUsername: legacyUser.Username,
											TwitterAvatar:   legacyUser.ProfileUrl,
										}, nil
									}

								}
							}
						}
					}
				}
			}
		}
	}

	return m.IsPostForFaucetAgentBackup(code)
}

func (m Rapid) _isPostForFaucetAgentBackup(code, cursor string) (*TwitterDetail, string, error) {
	path := fmt.Sprintf("https://twitter135.p.rapidapi.com/Search/?q=(@CryptoEternalAI)&cursor=%s", cursor)

	data, _, err := m.request(path, "GET", nil, nil)
	if err != nil {
		return nil, "", err
	}
	res := &RapidSearchResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, "", err
	}
	numContent := uint(0)
	cursorTop := ""
	if res != nil && res.Data != nil && res.Data.SearchByRawQuery != nil &&
		res.Data.SearchByRawQuery.SearchTimeline != nil && res.Data.SearchByRawQuery.SearchTimeline.Timeline != nil &&
		res.Data.SearchByRawQuery.SearchTimeline.Timeline.Instructions != nil {
		for _, v := range res.Data.SearchByRawQuery.SearchTimeline.Timeline.Instructions {
			if v.Type == "TimelineAddEntries" {
				for _, item := range v.Entries {
					if item.Content.EntryType == "TimelineTimelineCursor" && item.Content.CursorType == "Top" {
						cursorTop = item.Content.Value
					}
					if item.Content.EntryType == "TimelineTimelineItem" {
						if item.Content != nil && item.Content.ItemContent != nil && item.Content.ItemContent.TweetResults != nil && item.Content.ItemContent.TweetResults.Result != nil && item.Content.ItemContent.TweetResults.Result.Legacy != nil && item.Content.ItemContent.TweetResults.Result.Core != nil && item.Content.ItemContent.TweetResults.Result.Core.UserResults != nil && item.Content.ItemContent.TweetResults.Result.Core.UserResults.Results != nil && item.Content.ItemContent.TweetResults.Result.Core.UserResults.Results.Legacy != nil {
							legacy := item.Content.ItemContent.TweetResults.Result.Legacy
							legacyUser := item.Content.ItemContent.TweetResults.Result.Core.UserResults.Results.Legacy
							isMention := false
							for _, t := range legacy.Entities.UserMentions {
								if strings.EqualFold(t.ScreenName, "CryptoEternalAI") {
									isMention = true
								}
							}
							if isMention {
								numContent += 1
								if err == nil {
									if strings.Contains(legacy.FullText, code) {
										return &TwitterDetail{
											TwitterID:       legacy.UserID,
											TwitterName:     legacyUser.Name,
											TwitterUsername: legacyUser.Username,
											TwitterAvatar:   legacyUser.ProfileUrl,
										}, cursorTop, nil
									}

								}
							}
						}
					}
				}
			}
		}
	}

	return nil, cursorTop, nil
}

func (m Rapid) IsPostForFaucetAgentBackup(code string) (*TwitterDetail, error) {
	info, cursor, err := m._isPostForFaucetAgentBackup(code, "")
	if err != nil {
		return nil, err
	}
	numRun := 0
	for numRun < 1 {
		numRun++
		if info == nil && cursor != "" {
			info, cursor, err = m._isPostForFaucetAgentBackup(code, cursor)
			if err != nil {
				return nil, err
			}
			if info != nil {
				return info, nil
			}
		}
	}
	return nil, nil
}

func (m Rapid) GetTwitterUserFromTweetID(tweetID string, authPublicCode string) (*TwitterDetail, error) {
	path := fmt.Sprintf("%s/TweetDetail/?id=%s", m.url, tweetID)
	data, _, err := m.request(path, "GET", nil, nil)
	if err != nil {
		return nil, err
	}
	resp := &RapidTweetDetailResponse{}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	if resp != nil && resp.Data != nil && resp.Data.Conversation != nil &&
		resp.Data.Conversation.Instructions != nil && len(resp.Data.Conversation.Instructions) > 0 {
		for _, v := range resp.Data.Conversation.Instructions {
			if v.Type == "TimelineAddEntries" {
				for _, item := range v.Entries {
					if item.Content.EntryType == "TimelineTimelineItem" {
						if item.Content != nil && item.Content.ItemContent != nil && item.Content.ItemContent.TweetResults != nil && item.Content.ItemContent.TweetResults.Result != nil && item.Content.ItemContent.TweetResults.Result.Legacy != nil && item.Content.ItemContent.TweetResults.Result.Core != nil && item.Content.ItemContent.TweetResults.Result.Core.UserResults != nil && item.Content.ItemContent.TweetResults.Result.Core.UserResults.Results != nil && item.Content.ItemContent.TweetResults.Result.Core.UserResults.Results.Legacy != nil {
							legacy := item.Content.ItemContent.TweetResults.Result.Legacy
							legacyUser := item.Content.ItemContent.TweetResults.Result.Core.UserResults.Results.Legacy

							isMention := false
							for _, t := range legacy.Entities.UserMentions {
								if strings.EqualFold(t.ScreenName, "CryptoEternalAI") {
									isMention = true
								}
							}
							if isMention {
								if strings.Contains(legacy.FullText, authPublicCode) {
									return &TwitterDetail{
										TwitterID:       legacy.UserID,
										TwitterName:     legacyUser.Name,
										TwitterUsername: legacyUser.Username,
										TwitterAvatar:   legacyUser.ProfileUrl,
									}, nil
								}
							}
						}
					}
				}
			}
		}
	}
	return nil, nil
}
