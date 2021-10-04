package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/cloudogu/cas-oidc-dogu/config"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/caos/oidc/pkg/client/rp"
	"github.com/caos/oidc/pkg/oidc"
	"github.com/caos/oidc/pkg/utils"
)

var (
	callbackPath string = "/cas-oidc-client/auth/callback"
	key          []byte = []byte("test1234test1234")
)

func main() {
	appConfig, err := config.ReadConfig("conf.yaml")
	if err != nil {
		panic(err)
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	clientID := appConfig.ClientID
	clientSecret := appConfig.ClientSecret
	issuer := appConfig.Issuer
	port := appConfig.Port
	fqdn := appConfig.FQDN

	logrus.Info("%+v", appConfig)

	scopes := strings.Split("openid email groups", " ")

	redirectURI := fmt.Sprintf("%s%s", fqdn, callbackPath)
	cookieHandler := utils.NewCookieHandler(key, key, utils.WithUnsecure())

	options := []rp.Option{
		rp.WithCookieHandler(cookieHandler),
		rp.WithVerifierOpts(rp.WithIssuedAtOffset(5 * time.Second)),
	}
	if clientSecret == "" {
		options = append(options, rp.WithPKCE(cookieHandler))
	}
	provider, err := rp.NewRelyingPartyOIDC(issuer, clientID, clientSecret, redirectURI, scopes, options...)
	if err != nil {
		logrus.Fatalf("error creating provider %s", err.Error())
	}

	//generate some state (representing the state of the user in your application,
	//e.g. the page where he was before sending him to login
	state := func() string {
		return uuid.New().String()
	}

	//register the AuthURLHandler at your preferred path
	//the AuthURLHandler creates the auth request and redirects the user to the auth server
	//including state handling with secure cookie and the possibility to use PKCE
	http.Handle("/cas-oidc-client/login", rp.AuthURLHandler(state, provider))

	logrus.Infof("load static content...")
	//for demonstration purposes the returned userinfo response is written as JSON object onto response
	http.HandleFunc("/cas-oidc-client/", serveFiles)

	http.HandleFunc("/cas-oidc-client/logout", logout)

	//for demonstration purposes the returned userinfo response is written as JSON object onto response
	marshalUserinfo := func(w http.ResponseWriter, r *http.Request, tokens *oidc.Tokens, state string, rp rp.RelyingParty, info oidc.UserInfo) {
		data, err := json.Marshal(info)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(data)
	}
	http.Handle(callbackPath, rp.CodeExchangeHandler(rp.UserinfoCallback(marshalUserinfo), provider))

	logrus.Infof("listening on http://localhost:%s/cas-oidc-client/", port)
	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./cas-oidc-client/" {
		p = "./static/index.html"
	}
	http.ServeFile(w, r, p)
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.WriteHeader(200)
}
