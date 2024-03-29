package unicode_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var endpoint_url string = "unicode.compressibleflowcalculator.com/Prod/api/v1"
var host_url string = "unicode.compressibleflowcalculator.com"

type UnicodeProviderClient struct {
	Username           string
	mutual_access_test sync.Mutex
	//
	reserved map[string]bool
}

type UnicodeRequestModel struct {
	Name        string `json:"name" tfsdk:"name"`
	Description string `json:"description" tfsdk:"description"`
	Created_at  string `json:"created_at" tfsdk:"created_at"`
	Updated_at  string `json:"updated_at" tfsdk:"updated_at"`
}

type UnicodeAppModel struct {
	Id          basetypes.StringValue `json:"appid" tfsdk:"id"`
	Name        string                `json:"name" tfsdk:"name"`
	Description string                `json:"description" tfsdk:"description"`
	Created_at  string                `json:"created_at" tfsdk:"created_at"`
	Updated_at  string                `json:"updated_at" tfsdk:"updated_at"`
}

type UnicodeAppModelResponse struct {
	Id          string       `json:"appid" tfsdk:"id"`
	Name        string       `json:"name" tfsdk:"name"`
	Description string       `json:"description" tfsdk:"description"`
	Created_at  string       `json:"created_at" tfsdk:"created_at"`
	Updated_at  string       `json:"updated_at" tfsdk:"updated_at"`
	Conversions []Conversion `json:"conversions"`
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
	return &UnicodeProviderClient{Username: username, reserved: make(map[string]bool), mutual_access_test: sync.Mutex{}}
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

	//var app UnicodeAppModel
	var app UnicodeAppModelResponse

	// Decode JSON
	err = json.NewDecoder(res.Body).Decode(&app)

	if err != nil {
		return nil, err
	}

	// Copy app fields into model

	return &UnicodeAppModel{
		Id:          basetypes.NewStringValue(app.Id),
		Name:        app.Name,
		Description: app.Description,
		Created_at:  app.Created_at,
		Updated_at:  app.Updated_at,
	}, nil

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
		//Id:          model.Id,
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

	var response UnicodeAppModelResponse

	// Decode JSON

	err = json.NewDecoder(res.Body).Decode(&response)

	model.Id = basetypes.NewStringValue(response.Id)

	return &model, nil

	//return &response, nil

	//return &model, nil
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

	id := model.Id.ValueString()

	req := &http.Request{
		Method: "PUT",
		URL:    &url.URL{Scheme: "https", Host: host_url, Path: "/Prod/api/v1/application/" + id},
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

	//Check Mutex
	for true {
		u.mutual_access_test.Lock()
		if u.reserved[plan.AppId] == true {
			u.mutual_access_test.Unlock()
			continue
		}
		u.reserved[plan.AppId] = true
		u.mutual_access_test.Unlock()
		break
	}

	if u.Username == "" {
		u.mutual_access_test.Lock()

		u.reserved[plan.AppId] = false
		u.mutual_access_test.Unlock()
		return nil, []Conversion{}, fmt.Errorf("Username is Empty")
	}

	var current_app UnicodeAppModelReqString

	jar, err := cookiejar.New(nil)
	if err != nil {
		u.mutual_access_test.Lock()

		u.reserved[plan.AppId] = false
		u.mutual_access_test.Unlock()
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
		u.mutual_access_test.Lock()

		u.reserved[plan.AppId] = false
		u.mutual_access_test.Unlock()
		return nil, []Conversion{}, err
	}

	if res.StatusCode != http.StatusOK {
		u.mutual_access_test.Lock()
		u.reserved[plan.AppId] = false
		u.mutual_access_test.Unlock()
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
		u.mutual_access_test.Lock()

		u.reserved[plan.AppId] = false
		u.mutual_access_test.Unlock()
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
		u.mutual_access_test.Lock()

		u.reserved[plan.AppId] = false
		u.mutual_access_test.Unlock()
		return nil, []Conversion{}, err
	}

	if res.StatusCode != http.StatusOK {
		u.mutual_access_test.Lock()

		u.reserved[plan.AppId] = false
		u.mutual_access_test.Unlock()
		return nil, []Conversion{}, fmt.Errorf("unexpected status post code %d", res.StatusCode)
	}

	u.mutual_access_test.Lock()
	u.reserved[plan.AppId] = false
	u.mutual_access_test.Unlock()

	return &plan, current_app.Conversions, nil
}

func (u *UnicodeProviderClient) RemoveConversionFromApplication(plan UnicodeStringModel) (*UnicodeAppModel, error) {

	var current_app UnicodeAppModelReqString

	//Check Mutex
	for true {
		u.mutual_access_test.Lock()
		if u.reserved[plan.AppId] == true {
			u.mutual_access_test.Unlock()
			continue
		}
		u.reserved[plan.AppId] = true
		u.mutual_access_test.Unlock()
		break
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		u.mutual_access_test.Lock()

		u.reserved[plan.AppId] = false
		u.mutual_access_test.Unlock()
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

	if res.StatusCode == http.StatusNotFound {
		// Application not found, return nil
		u.mutual_access_test.Lock()

		u.reserved[plan.AppId] = false
		u.mutual_access_test.Unlock()
		return nil, nil //Already Deleted, No Error
	}

	if err != nil {
		u.mutual_access_test.Lock()

		u.reserved[plan.AppId] = false
		u.mutual_access_test.Unlock()
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		u.mutual_access_test.Lock()

		u.reserved[plan.AppId] = false
		u.mutual_access_test.Unlock()
		return nil, fmt.Errorf("unexpected status get code %d", res.StatusCode)
	}

	// Decode JSON
	err = json.NewDecoder(res.Body).Decode(&current_app)

	// Now  Find Current String With That Value and Remove it
	var new_conversions []Conversion
	anymatch := false

	for _, v := range current_app.Conversions {
		if v.Value != plan.Value {
			new_conversions = append(new_conversions, v)
			continue
		}
		anymatch = true
	}

	if !anymatch {
		u.mutual_access_test.Lock()

		u.reserved[plan.AppId] = false
		u.mutual_access_test.Unlock()

		return nil, fmt.Errorf("Conversion Not Found")
	}

	current_app.Conversions = new_conversions

	if len(current_app.Conversions) == 0 {
		current_app.Conversions = []Conversion{}
	}

	current_app.User_id = u.Username

	//Now, Update the Application
	body_bytes, err := json.Marshal(current_app)

	if err != nil {
		u.mutual_access_test.Lock()

		u.reserved[plan.AppId] = false
		u.mutual_access_test.Unlock()
		return nil, err
	}

	update_req := &http.Request{
		Method: "POST",
		URL:    &url.URL{Scheme: "https", Host: host_url, Path: "/Prod/api/v1/application"},
		Body:   io.NopCloser(bytes.NewReader(body_bytes)),
		Header: make(http.Header),
	}

	res, err = client.Do(update_req)

	if err != nil {
		u.mutual_access_test.Lock()

		u.reserved[plan.AppId] = false
		u.mutual_access_test.Unlock()
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		u.mutual_access_test.Lock()

		u.reserved[plan.AppId] = false
		u.mutual_access_test.Unlock()
		return nil, fmt.Errorf("unexpected status post code %d", res.StatusCode)
	}

	u.mutual_access_test.Lock()

	u.reserved[plan.AppId] = false
	u.mutual_access_test.Unlock()

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

	return &UnicodeStringModel{
		Value: plan.Value,
		AppId: model.AppId,
	}, nil

}
