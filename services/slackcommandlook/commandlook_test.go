package slackcommandlook

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

func TestCommandLookHandler_NoText(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatal(err)
	}

	const slackVerificationToken = "0123456789"

	form := url.Values{}
	form.Add("token", slackVerificationToken)
	form.Add("text", "")
	r, err := inst.NewRequest("POST", "/command/look", bytes.NewReader([]byte(form.Encode())))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		t.Fatal(err)
	}

	c := appengine.NewContext(r)

	err = setSlackVerificationToken(c, slackVerificationToken)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)
	response := string(w.Body.Bytes())
	wantResponsePrefix := "Try using the /look command with one of these words:"

	if strings.HasPrefix(response, wantResponsePrefix) {
		t.Logf("got %q", response)
	} else {
		t.Errorf("got %q, want prefix %q", response, wantResponsePrefix)
	}
}

func setSlackVerificationToken(c context.Context, token string) error {
	_, err := datastore.Put(
		c,
		datastore.NewKey(c, "secret", "SLACK_VERIFICATION_TOKEN", 0, nil),
		&struct {
			Value string `datastore:"value"`
		}{
			Value: token,
		},
	)
	return err
}
