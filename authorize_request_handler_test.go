package fosite_test

import (
	"github.com/golang/mock/gomock"
	"github.com/ory-am/common/pkg"
	. "github.com/ory-am/fosite"
	. "github.com/ory-am/fosite/client"
	. "github.com/ory-am/fosite/internal"
	"github.com/stretchr/testify/assert"
	"github.com/vektra/errors"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"testing"
)

// Should pass
//
// * https://openid.net/specs/oauth-v2-multiple-response-types-1_0.html#Terminology
//   The OAuth 2.0 specification allows for registration of space-separated response_type parameter values.
//   If a Response Type contains one of more space characters (%20), it is compared as a space-delimited list of
//   values in which the order of values does not matter.
func TestNewAuthorizeRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := NewMockStorage(ctrl)
	defer ctrl.Finish()

	redir, _ := url.Parse("https://foo.bar/cb")
	for k, c := range []struct {
		desc          string
		conf          *Fosite
		r             *http.Request
		query         url.Values
		expectedError error
		mock          func()
		expect        *AuthorizeRequest
	}{
		/* empty request */
		{
			desc:          "empty request fails",
			conf:          &Fosite{Store: store},
			r:             &http.Request{},
			expectedError: ErrInvalidClient,
			mock: func() {
				store.EXPECT().GetClient(gomock.Any()).Return(nil, errors.New("foo"))
			},
		},
		/* invalid redirect uri */
		{
			desc:          "invalid redirect uri fails",
			conf:          &Fosite{Store: store},
			query:         url.Values{"redirect_uri": []string{"invalid"}},
			expectedError: ErrInvalidClient,
			mock: func() {
				store.EXPECT().GetClient(gomock.Any()).Return(nil, errors.New("foo"))
			},
		},
		/* invalid client */
		{
			desc:          "invalid client fails",
			conf:          &Fosite{Store: store},
			query:         url.Values{"redirect_uri": []string{"https://foo.bar/cb"}},
			expectedError: ErrInvalidClient,
			mock: func() {
				store.EXPECT().GetClient(gomock.Any()).Return(nil, errors.New("foo"))
			},
		},
		/* redirect client mismatch */
		{
			desc: "client and request redirects mismatch",
			conf: &Fosite{Store: store},
			query: url.Values{
				"client_id": []string{"1234"},
			},
			expectedError: ErrInvalidRequest,
			mock: func() {
				store.EXPECT().GetClient("1234").Return(&SecureClient{RedirectURIs: []string{"invalid"}}, nil)
			},
		},
		/* redirect client mismatch */
		{
			desc: "client and request redirects mismatch",
			conf: &Fosite{Store: store},
			query: url.Values{
				"redirect_uri": []string{""},
				"client_id":    []string{"1234"},
			},
			expectedError: ErrInvalidRequest,
			mock: func() {
				store.EXPECT().GetClient("1234").Return(&SecureClient{RedirectURIs: []string{"invalid"}}, nil)
			},
		},
		/* redirect client mismatch */
		{
			desc: "client and request redirects mismatch",
			conf: &Fosite{Store: store},
			query: url.Values{
				"redirect_uri": []string{"https://foo.bar/cb"},
				"client_id":    []string{"1234"},
			},
			expectedError: ErrInvalidRequest,
			mock: func() {
				store.EXPECT().GetClient("1234").Return(&SecureClient{RedirectURIs: []string{"invalid"}}, nil)
			},
		},
		/* no state */
		{
			desc: "no state",
			conf: &Fosite{Store: store},
			query: url.Values{
				"redirect_uri":  []string{"https://foo.bar/cb"},
				"client_id":     []string{"1234"},
				"response_type": []string{"code"},
			},
			expectedError: ErrInvalidState,
			mock: func() {
				store.EXPECT().GetClient("1234").Return(&SecureClient{RedirectURIs: []string{"https://foo.bar/cb"}}, nil)
			},
		},
		/* short state */
		{
			desc: "short state",
			conf: &Fosite{Store: store},
			query: url.Values{
				"redirect_uri":  {"https://foo.bar/cb"},
				"client_id":     {"1234"},
				"response_type": {"code"},
				"state":         {"short"},
			},
			expectedError: ErrInvalidState,
			mock: func() {
				store.EXPECT().GetClient("1234").Return(&SecureClient{RedirectURIs: []string{"https://foo.bar/cb"}}, nil)
			},
		},
		/* success case */
		{
			desc: "should pass",
			conf: &Fosite{Store: store},
			query: url.Values{
				"redirect_uri":  {"https://foo.bar/cb"},
				"client_id":     {"1234"},
				"response_type": {"code"},
				"state":         {"strong-state"},
				"scope":         {"foo bar"},
			},
			mock: func() {
				store.EXPECT().GetClient("1234").Return(&SecureClient{RedirectURIs: []string{"https://foo.bar/cb"}}, nil)
			},
			expectedError: ErrInvalidScope,
		},
		{
			desc: "should not pass because hybrid flow is not active",
			conf: &Fosite{Store: store},
			query: url.Values{
				"redirect_uri":  {"https://foo.bar/cb"},
				"client_id":     {"1234"},
				"response_type": {"code token"},
				"state":         {"strong-state"},
				"scope":         {DefaultRequiredScopeName + " foo bar"},
			},
			mock: func() {
				store.EXPECT().GetClient("1234").Return(&SecureClient{RedirectURIs: []string{"https://foo.bar/cb"}}, nil)
			},
			expectedError: ErrInvalidRequest,
		},
		{
			desc: "should not pass because hybrid flow is not active",
			conf: &Fosite{Store: store},
			query: url.Values{
				"redirect_uri":  {"https://foo.bar/cb"},
				"client_id":     {"1234"},
				"response_type": {"code"},
				"state":         {"strong-state"},
				"scope":         {DefaultRequiredScopeName + " foo bar"},
			},
			mock: func() {
				store.EXPECT().GetClient("1234").Return(&SecureClient{RedirectURIs: []string{"https://foo.bar/cb"}}, nil)
			},
			expect: &AuthorizeRequest{
				RedirectURI:   redir,
				ResponseTypes: []string{"code"},
				State:         "strong-state",
				Request: Request{
					Scopes: []string{DefaultRequiredScopeName, "foo", "bar"},
					Client: &SecureClient{RedirectURIs: []string{"https://foo.bar/cb"}},
				},
			},
		},
		{
			desc: "should pass",
			conf: &Fosite{Store: store, AllowHybridFlow: true},
			query: url.Values{
				"redirect_uri":  {"https://foo.bar/cb"},
				"client_id":     {"1234"},
				"response_type": {"code token"},
				"state":         {"strong-state"},
				"scope":         {DefaultRequiredScopeName + " foo bar"},
			},
			mock: func() {
				store.EXPECT().GetClient("1234").Return(&SecureClient{RedirectURIs: []string{"https://foo.bar/cb"}}, nil)
			},
			expect: &AuthorizeRequest{
				RedirectURI:   redir,
				ResponseTypes: []string{"code", "token"},
				State:         "strong-state",
				Request: Request{
					Client: &SecureClient{RedirectURIs: []string{"https://foo.bar/cb"}},
					Scopes: []string{DefaultRequiredScopeName, "foo", "bar"},
				},
			},
		},
	} {
		t.Logf("Joining test case %d", k)
		c.mock()
		if c.r == nil {
			c.r = &http.Request{Header: http.Header{}}
			if c.query != nil {
				c.r.URL = &url.URL{RawQuery: c.query.Encode()}
			}
		}

		ar, err := c.conf.NewAuthorizeRequest(context.Background(), c.r)
		assert.Equal(t, c.expectedError == nil, err == nil, "%d: %s\n%s", k, c.desc, err)
		if c.expectedError != nil {
			assert.Equal(t, err.Error(), c.expectedError.Error(), "%d: %s\n%s", k, c.desc, err)
		} else {
			pkg.AssertObjectKeysEqual(t, c.expect, ar, "ResponseTypes", "Scopes", "Client", "RedirectURI", "State")
			assert.NotNil(t, ar.GetRequestedAt())
		}
		t.Logf("Passed test case %d", k)
	}
}
