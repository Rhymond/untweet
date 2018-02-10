package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
)

// Client struct extends instance of TwitterGo
type Client struct {
	*twittergo.Client
}

// FollowersList keeps followers ID's list
type FollowersList struct {
	Ids []int64
}

// NewClient creates new Client instance by using
// Customer Key, Customer Secret, Auth Token and Auth Secret
func NewClient(ck, cs, at, as string) *Client {
	oauthConfig := &oauth1a.ClientConfig{
		ConsumerKey:    ck,
		ConsumerSecret: cs,
	}

	user := oauth1a.NewAuthorizedConfig(at, as)
	return &Client{twittergo.NewClient(oauthConfig, user)}
}

// Notify send direct message to given user ID
func (c *Client) Notify(userID, text string) error {
	query := url.Values{}
	query.Set("user_id", userID)
	query.Set("text", text)

	_, err := c.Send("POST", "/1.1/direct_messages/new.json", query)

	return err
}

// GetUserInfo gets user information by given user ID
func (c *Client) GetUserInfo(userID string) (*twittergo.User, error) {
	query := url.Values{}
	query.Set("user_id", userID)

	resp, err := c.Send("GET", "/1.1/users/show.json", query)

	if err != nil {
		return nil, err
	}

	results := &twittergo.User{}
	err = resp.Parse(results)

	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetFollowersList gets followers list by given user ID
func (c *Client) GetFollowersList(userID string) (*FollowersList, error) {
	query := url.Values{}
	query.Set("user_id", userID)

	resp, err := c.Send("GET", "/1.1/followers/ids.json", query)

	if err != nil {
		return nil, err
	}

	results := &FollowersList{}
	err = resp.Parse(results)

	if err != nil {
		return nil, err
	}

	return results, nil
}

// Send sends request to twitter API using twitterGo package
// It logs request and rate limit data
func (c *Client) Send(reqType, endpoint string, query url.Values) (*twittergo.APIResponse, error) {
	url := fmt.Sprintf("%s?%v", endpoint, query.Encode())

	log.Printf(
		"Sending Request: [%s] %s Params: [%v]",
		reqType,
		endpoint,
		query,
	)

	req, err := http.NewRequest(reqType, url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.SendRequest(req)

	log.Printf(
		"Rate Limit: %d/%d, Rate Limit Reset: %d (%s)",
		resp.RateLimitRemaining(),
		resp.RateLimit(),
		resp.RateLimitReset().Unix(),
		resp.RateLimitReset().Format(time.RFC1123),
	)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
