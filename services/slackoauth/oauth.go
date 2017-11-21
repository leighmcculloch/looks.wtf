package slackoauth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/leighmcculloch/looks.wtf/shared/secrets"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

type slackOauth struct {
	Ok          bool   `json:"ok"`
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	UserID      string `json:"user_id"`
	TeamName    string `json:"team_name"`
	TeamID      string `json:"team_id"`
	Timestamp   time.Time
}

func persistSlackOauth(c context.Context, so slackOauth) error {
	so.Timestamp = time.Now().UTC()
	key := datastore.NewIncompleteKey(c, "slackOauth", nil)
	_, err := datastore.Put(c, key, &so)
	return err
}

func oauthHandler(c context.Context, w http.ResponseWriter, r *http.Request) error {
	code := r.URL.Query().Get("code")
	log.Infof(c, "Request with code: %#v", code)

	slackClientID := secrets.Get(c, "SLACK_CLIENT_ID")
	slackClientSecret := secrets.Get(c, "SLACK_CLIENT_SECRET")

	client := urlfetch.Client(c)
	resp, err := client.PostForm(
		"https://slack.com/api/oauth.access",
		url.Values{
			"client_id":     {slackClientID},
			"client_secret": {slackClientSecret},
			"code":          {code},
		},
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var msg slackOauth
	err = json.Unmarshal(body, &msg)
	if err != nil || !msg.Ok {
		log.Errorf(c, "Error in response: %#v", string(body))
		return err
	}

	log.Infof(c, "Success, response: %#v", msg)
	fmt.Fprintf(w, "The looks.wtf Slack App has been successfully added!\nGive it a go with the `/look awe` command in Slack.")
	err = persistSlackOauth(c, msg)
	if err != nil {
		log.Infof(c, "Failed to persist: %#v", msg)
		return err
	}
	return nil
}
