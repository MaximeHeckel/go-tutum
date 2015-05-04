package tutum

import "encoding/json"

type StackListResponse struct {
	Objects []StackShort `json:"objects"`
}

type StackShort struct {
	Deployed_datetime  string   `json:"deployed_datetime"`
	Destroyed_datetime string   `json:"destroyed_datetime"`
	Name               string   `json:"name"`
	Resource_uri       string   `json:"resource_uri"`
	Service            []string `json:"services"`
	State              string   `json:"state"`
	Synchronized       bool     `json:"synchronized"`
	Uuid               string   `json:"uuid"`
}

type Stack struct {
	Deployed_datetime  string    `json:"deployed_datetime"`
	Destroyed_datetime string    `json:"destroyed_datetime`
	Name               string    `json:"name"`
	Resource_uri       string    `json:"resource_uri`
	Service            []Service `json:"services"`
	State              string    `json:"state"`
	Synchronized       bool      `json:"synchronized"`
	Uuid               string    `json:"uuid"`
}

/*
func ListStacks
Returns : Array of Stack objects
*/
func ListStacks() (StackListResponse, error) {
	url := "stack/"
	request := "GET"

	//Empty Body Request
	body := []byte(`{}`)
	var response StackListResponse

	data, err := TutumCall(url, request, body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}

/*
func GetStack
Argument : uuid
Returns : Stack JSON object
*/
func GetStack(uuid string) (Stack, error) {

	url := ""
	if string(uuid[0]) == "/" {
		url = uuid[8:]
	} else {
		url = "stack/" + uuid + "/"
	}

	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Stack

	data, err := TutumCall(url, request, body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

/*
func Export
Returns : String that contains the Stack details
*/
func (self *Stack) ExportStack() (string, error) {

	url := "stack/" + self.Uuid + "/export/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)

	data, err := TutumCall(url, request, body)
	if err != nil {
		return "", err
	}

	s := string(data)

	return s, nil
}

/*
func CreateStack
Argument : Stack JSON object (see documentation)
*/
func CreateStack(requestBody string) (Stack, error) {
	url := "stack/"
	request := "POST"

	newStack := []byte(requestBody)
	var response Stack

	data, err := TutumCall(url, request, newStack)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

/*
func Update
Argument : a Stack JSON object (see documentation)
*/
func (self *Stack) Update(requestBody string) error {

	url := "stack/" + self.Uuid + "/"
	request := "PATCH"

	updatedStack := []byte(requestBody)

	_, err := TutumCall(url, request, updatedStack)
	if err != nil {
		return err
	}

	return nil
}

func (self *Stack) Start() error {

	url := "stack/" + self.Uuid + "/start/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}

func (self *Stack) Stop() error {

	url := "stack/" + self.Uuid + "/stop/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}

func (self *Stack) Redeploy() error {

	url := "stack/" + self.Uuid + "/redeploy/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}

func (self *Stack) Terminate() error {

	url := "stack/" + self.Uuid + "/"
	request := "DELETE"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}