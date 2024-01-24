package unicode_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

var endpoint_url string = "unicode.compressibleflowcalculator.com/Prod/api/v1"
var host_url string = "unicode.compressibleflowcalculator.com"

type UnicodeProviderClient struct {
	Username string
}

type UnicodeAppModel struct {
	Id          string `json:"appid" tfsdk:"id"`
	Name        string `json:"name" tfsdk:"name"`
	Description string `json:"description" tfsdk:"description"`
	Created_at  string `json:"created_at" tfsdk:"created_at"`
	Updated_at  string `json:"updated_at" tfsdk:"updated_at"`
}

type UnicodeAppModelReq struct {
	Id          string       `json:"appid" tfsdk:"id"`
	Name        string       `json:"name" tfsdk:"name"`
	Description string       `json:"description" tfsdk:"description"`
	Created_at  string       `json:"created_at" tfsdk:"created_at"`
	Updated_at  string       `json:"updated_at" tfsdk:"updated_at"`
	Conversions []Conversion `json:"conversions"`
}

type Conversion struct {
	Value string `json:"value"`
}

func NewUnicodeProviderClient(username string) *UnicodeProviderClient {
	return &UnicodeProviderClient{Username: username}
}

type UnicodeData struct {
	Char     string `json:"char" tfsdk:"unicode_char"`
	Category string `tfsdk:"unicode_category"`
	Block    string `tfsdk:"unicode_block"`
	Name     string `tfsdk:"unicode_name"`
}

func (u *UnicodeProviderClient) GetUnicodeCharData(unicodeChar string) (*UnicodeData, error) {
	return &UnicodeData{
		Char:     unicodeChar,
		Category: "Ll",
		Block:    "Basic Latin",
		Name:     "LATIN SMALL LETTER A",
	}, nil
}

func (u *UnicodeProviderClient) GetApplications() (*UnicodeAppModel, error) {

	client := &http.Client{}

	return &UnicodeAppModel{

		Id:          "1",
		Name:        "App 1",
		Description: "App 1 Description",
		Created_at:  "2021-01-01",
		Updated_at:  "2021-01-01",
	}, nil
	// Now Assign Cookie Header of user=Username
	client.Jar.SetCookies(&url.URL{Scheme: "https", Host: host_url}, []*http.Cookie{
		{
			Name:  "user",
			Value: u.Username,
		},
	})

	resp, err := client.Do(&http.Request{
		Method: "GET",
		URL:    &url.URL{Scheme: "https", Host: endpoint_url + "/saved"},
	})

	if err != nil {
		return nil, err
	}

	var apps []UnicodeAppModel

	// Decode JSON
	err = json.NewDecoder(resp.Body).Decode(&apps)

	if err != nil {
		return nil, err
	}

	//return &apps, nil
	return &(apps[0]), nil
}

func (u *UnicodeProviderClient) CreateApplication(model UnicodeAppModel) (*UnicodeAppModel, error) {

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Jar: jar,
	}

	//sET cOOKIE
	client.Jar.SetCookies(&url.URL{Scheme: "https", Host: host_url}, []*http.Cookie{
		{
			Name:  "user",
			Value: u.Username,
		},
	})

	// turn model into model request
	model_req := UnicodeAppModelReq{
		Id:          model.Id,
		Name:        model.Name,
		Description: model.Description,
		Created_at:  model.Created_at,
		Updated_at:  model.Updated_at,
		Conversions: []Conversion{},
	}

	body_bytes, err := json.Marshal(model_req)
	if err != nil {
		return nil, err
	}

	req := &http.Request{
		Method: "POST",
		URL:    &url.URL{Scheme: "https", Host: host_url, Path: "/Prod/api/v1/application"},
		Body:   io.NopCloser(bytes.NewReader(body_bytes)),
		Header: make(http.Header),
	}

	res, err := client.Do(req)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (u *UnicodeProviderClient) DeleteApplication(id string) error {

	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}

	client := &http.Client{
		Jar: jar,
	}

	//sET cOOKIE
	client.Jar.SetCookies(&url.URL{Scheme: "https", Host: host_url}, []*http.Cookie{
		{
			Name:  "user",
			Value: u.Username,
		},
	})

	req := &http.Request{
		Method: "DELETE",
		URL:    &url.URL{Scheme: "https", Host: host_url, Path: "/Prod/api/v1/application/" + id},
		Body:   nil,
		Header: make(http.Header),
	}

	res, err := client.Do(req)

	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

	return nil
}

func (u *UnicodeProviderClient) UpdateApplication(model UnicodeAppModel) (*UnicodeAppModel, error) {

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Jar: jar,
	}

	//sET cOOKIE

	body_bytes, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	req := &http.Request{
		Method: "PUT",
		URL:    &url.URL{Scheme: "https", Host: host_url, Path: "/Prod/api/v1/application/" + model.Id},
		Body:   io.NopCloser(bytes.NewReader(body_bytes)),
		Header: make(http.Header),
	}

	_, err = client.Do(req)

	if err != nil {
		return nil, err
	}

	return &model, nil
}
