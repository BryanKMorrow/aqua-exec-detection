package aqua

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
	"os"
	"strings"
)

type Client struct {
	user     string `json:"user"`
	url      string `json:"url"`
	password string `json:"password"`
	token    string `json:"token"`
}

// NewClient - initialize and return the Client
func NewClient() *Client {
	c := &Client{
		url:      strings.TrimSpace(os.Getenv("AQUA_URL")),
		user:     strings.TrimSpace(os.Getenv("AQUA_USER")),
		password: strings.TrimSpace(os.Getenv("AQUA_PASSWORD")),
	}
	return c
}

// GetAuthToken - Connect to Aqua and return a JWT bearerToken (string)
// Return: bool - successfully connected?
func (cli *Client) GetAuthToken() bool {
	var connected bool
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, body, err := request.Post(cli.url+"/api/v1/login").Param("abilities", "1").
		Send(`{"id":"` + cli.user + `", "password":"` + cli.password + `"}`).End()
	if err != nil {
		connected = false
		return connected
	}

	if resp.StatusCode == 200 {
		var raw map[string]interface{}
		_ = json.Unmarshal([]byte(body), &raw)
		cli.token = raw["token"].(string)
		connected = true
	} else {
		log.Printf("Failed with status: %s", resp.Status)
		connected = false
	}
	return connected
}

type ImageDigest []struct {
	Name       string `json:"name"`
	Digest     string `json:"digest"`
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
}

// GetImageByDigest retrieves an image metadata based on provided digest
// returns the image name to be used in the runtime profile creation.
func (cli *Client) GetImageByDigest(digest string) (ImageDigest, error) {
	var response ImageDigest
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/images/details/%s", digest)
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetImage from %s%s, %v ", cli.url, apiPath, err)
		}
	} else {
		return response, errors.New("image not found by provided digest")
	}
	return response, nil
}

func (cli *Client) CreateRuntimePolicy(policy Policy) string {
	var response = ""
	payload, err := json.Marshal(policy)
	if err != nil {
		log.Println(err)
	}
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/runtime_policies")
	resp, _, errs := request.Clone().Post(cli.url + apiPath).Send(string(payload)).End()
	if errs != nil {
		log.Printf("Failed creating runtime policy: %s \n  Status Code: %d", cli.url+apiPath, resp.StatusCode)
	}
	if resp.StatusCode == 204 {
		response = `{"message": "Runtime Policy created successfully"}`
	} else if resp.StatusCode == 400 {
		response = fmt.Sprintf(`{"message": "Runtime Policy creation failed reading body", "status_code": %v}`, resp.StatusCode)
		log.Println(string(payload))
	} else if resp.StatusCode == 500 {
		response = fmt.Sprintf(`{"message": "Runtime Policy creation failed, policy already exists", "status_code": %v}`, resp.StatusCode)
		log.Println(string(payload))
	} else {
		response = fmt.Sprintf(`{"message": "Runtime Policy creation failed", "status_code": %v}`, resp.StatusCode)
		log.Println(string(payload))
	}
	return response
}
