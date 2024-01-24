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

type UnicodeAppModelReqString struct {
	Id          string       `json:"appid" tfsdk:"id"`
	Name        string       `json:"name" tfsdk:"name"`
	Description string       `json:"description" tfsdk:"description"`
	Created_at  string       `json:"created_at" tfsdk:"created_at"`
	Updated_at  string       `json:"updated_at" tfsdk:"updated_at"`
	Conversions []Conversion `json:"conversions"`
	User_id     string       `json:"user_id"`
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

func (u *UnicodeProviderClient) GetApplication(id string) (*UnicodeAppModel, error) {

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

	req := &http.Request{
		Method: "GET",
		URL:    &url.URL{Scheme: "https", Host: host_url, Path: "/Prod/api/v1/application/" + id},
		Body:   nil,
		Header: make(http.Header),
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

	var app UnicodeAppModel

	// Decode JSON
	err = json.NewDecoder(res.Body).Decode(&app)

	if err != nil {
		return nil, err
	}

	return &app, nil

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
	model_req := UnicodeAppModelReqString{
		Id:          model.Id,
		Name:        model.Name,
		Description: model.Description,
		Created_at:  model.Created_at,
		Updated_at:  model.Updated_at,
		Conversions: []Conversion{},
		User_id:     u.Username,
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

func (u *UnicodeProviderClient) AddConversionToApplication(plan UnicodeStringModel) (*UnicodeStringModel, []Conversion, error) {

	if u.Username == "" {
		return nil, []Conversion{}, fmt.Errorf("Username is Empty")
	}

	var current_app UnicodeAppModelReqString

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, []Conversion{}, err
	}

	client := &http.Client{
		Jar: jar,
	}

	//Set Cookie
	client.Jar.SetCookies(&url.URL{Scheme: "https", Host: host_url}, []*http.Cookie{
		{
			Name:  "user",
			Value: u.Username,
		},
	})

	//Find Current Application State
	get_app_req := &http.Request{
		Method: "GET",
		URL:    &url.URL{Scheme: "https", Host: host_url, Path: "/Prod/api/v1/application/" + plan.AppId},
		Body:   nil,
		Header: make(http.Header),
	}

	res, err := client.Do(get_app_req)

	if err != nil {
		return nil, []Conversion{}, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, []Conversion{}, fmt.Errorf("unexpected status get code %d", res.StatusCode)
	}

	// Decode JSON
	err = json.NewDecoder(res.Body).Decode(&current_app)

	//Now Append the Conversion to the Application

	var conv Conversion
	conv.Value = plan.Value

	var conversions []Conversion

	conversions = append(current_app.Conversions, conv)

	current_app.Conversions = conversions

	current_app.User_id = u.Username

	//Now, Update the Application
	body_bytes, err := json.Marshal(current_app)
	if err != nil {
		return nil, []Conversion{}, err
	}

	update_req := &http.Request{
		Method: "POST",
		URL:    &url.URL{Scheme: "https", Host: host_url, Path: "/Prod/api/v1/application"},
		Body:   io.NopCloser(bytes.NewReader(body_bytes)),
		Header: make(http.Header),
	}

	res, err = client.Do(update_req)

	if err != nil {
		return nil, []Conversion{}, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, []Conversion{}, fmt.Errorf("unexpected status post code %d", res.StatusCode)
	}

	return &plan, current_app.Conversions, nil
}

func (u *UnicodeProviderClient) RemoveConversionFromApplication(plan UnicodeStringModel) (*UnicodeAppModel, error) {

	var current_app UnicodeAppModelReq

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Jar: jar,
	}

	//Set Cookie
	client.Jar.SetCookies(&url.URL{Scheme: "https", Host: host_url}, []*http.Cookie{
		{
			Name:  "user",
			Value: u.Username,
		},
	})

	//Find Current Application State
	get_app_req := &http.Request{
		Method: "GET",
		URL:    &url.URL{Scheme: "https", Host: host_url, Path: "/Prod/api/v1/application/" + plan.AppId},
		Body:   nil,
		Header: make(http.Header),
	}

	res, err := client.Do(get_app_req)

	if err != nil {
		return nil, err
	}

	// Decode JSON
	err = json.NewDecoder(res.Body).Decode(&current_app)

	// Now  Find Current String With That Value and Remove it
	var new_conversions []Conversion

	for _, v := range current_app.Conversions {
		if v.Value != plan.Value {
			new_conversions = append(new_conversions, v)
		}
	}

	current_app.Conversions = new_conversions

	//Now, Update the Application
	body_bytes, err := json.Marshal(current_app)

	if err != nil {
		return nil, err
	}

	update_req := &http.Request{
		Method: "POST",
		URL:    &url.URL{Scheme: "https", Host: host_url, Path: "/Prod/api/v1/application/" + plan.AppId},
		Body:   io.NopCloser(bytes.NewReader(body_bytes)),
		Header: make(http.Header),
	}

	_, err = client.Do(update_req)

	if err != nil {
		return nil, err
	}

	return nil, nil

}

type UnicodeStringModel struct {
	Value string `json:"value" tfsdk:"value"`
	AppId string `json:"app_id" tfsdk:"app_id"`
}

func (u *UnicodeProviderClient) GetConversionFromApplication(model UnicodeStringModel) (*UnicodeStringModel, error) {

	var plan Conversion

	plan.Value = model.Value

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Jar: jar,
	}

	//Set Cookie
	client.Jar.SetCookies(&url.URL{Scheme: "https", Host: host_url}, []*http.Cookie{
		{
			Name:  "user",
			Value: u.Username,
		},
	})

	//Now, Get The Application and find the Conversions
	get_convo_req := &http.Request{
		Method: "GET",
		URL:    &url.URL{Scheme: "https", Host: host_url, Path: "/Prod/api/v1/application/" + model.AppId},
		Body:   nil,
		Header: make(http.Header),
	}

	res, err := client.Do(get_convo_req)

	if err != nil {
		return nil, err
	}

	var Application UnicodeAppModelReq

	// Decode JSON
	err = json.NewDecoder(res.Body).Decode(&Application)

	if err != nil {
		return nil, err
	}

	//Now Append the Conversion to the Application
	Application.Conversions = append(Application.Conversions, plan)

	//Now, Update the Application
	body_bytes, err := json.Marshal(Application)
	if err != nil {
		return nil, err
	}

	update_req := &http.Request{
		Method: "POST",
		URL:    &url.URL{Scheme: "https", Host: host_url, Path: "/Prod/api/v1/application/" + model.AppId},
		Body:   io.NopCloser(bytes.NewReader(body_bytes)),
		Header: make(http.Header),
	}

	res, err = client.Do(update_req)

	if err != nil {
		return nil, err
	}

	//Now

	return &model, nil
}
