package tutum

import "encoding/json"

type CListResponse struct {
	Objects []Container `json:"objects"`
}

type Container struct {
	Application            string    `json:"application"`
	Autodestroy            string    `json:"autodestroy"`
	Autoreplace            string    `json:"autoreplace"`
	Autorestart            string    `json:"autorestart"`
	Container_ports        []CCPInfo `json:"container_ports"`
	Container_size         string    `json:"container_size"`
	Current_num_containers int       `json:"current_num_containers"`
	Deployed_datetime      string    `json:"deployed_datetime"`
	Destroyed_datetime     string    `json:"destroyed_datetime"`
	Entrypoint             string    `json:"entrypoint"`
	Exit_code              int       `json:"exit_code"`
	Exit_code_message      string    `json:"exit_code_message"`
	Image_name             string    `json:"image_name"`
	Image_tag              string    `json:"image_tag"`
	Name                   string    `json:"name"`
	Public_dns             string    `json:"public_dns"`
	Resource_uri           string    `json:"resource_uri"`
	Run_command            string    `json:"run_command"`
	Started_datetime       string    `json:"started_datetime"`
	State                  string    `json:"state"`
	Stopped_datetime       string    `json:"stopped_datetime"`
	Unique_name            string    `json:"unique_name"`
	Uuid                   string    `json:"uuid"`
}

type CCPInfo struct {
	Container  string `json:"container"`
	Inner_port int    `json:"inner_port"`
	Outer_port int    `json:"outer_port"`
	Protocol   string `json:"protocol"`
}

/*
func ListContainers
Returns : Array of Container objects
*/
func ListContainers() (CListResponse, error) {

	url := "container/"
	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response CListResponse
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
func GetContainer
Argument : uuid
Returns : Container JSON object
*/
func GetContainer(uuid string) (Container, error) {

	url := ""
	if string(uuid[0]) == "/" {
		url = uuid[8:]
	} else {
		url = "container/" + uuid + "/"
	}

	request := "GET"
	//Empty Body Request
	body := []byte(`{}`)
	var response Container

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
func GetContainerLogs
Argument : uuid
Returns : A string containing the logs of the container
*/
func (self *Container) Logs() (string, error) {

	url := "container/" + self.Uuid + "/logs/"
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
func StartContainer
Argument : uuid
Returns : Container JSON object
*/
func (self *Container) Start() error {

	url := "container/" + self.Uuid + "/start/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)
	var response Container

	data, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return err
	}

	return nil
}

/*
func StopContainer
Argument : uuid
Returns : Container JSON object
*/
func (self *Container) Stop() error {

	url := "container/" + self.Uuid + "/stop/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}

/*
func RedeployContainer
Argument : uuid
Returns : Container JSON object
*/
func (self *Container) Redeploy() error {

	url := "container/" + self.Uuid + "/redeploy/"
	request := "POST"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}

/*
func TerminateContainer
Argument : uuid
Returns : Container JSON object
*/
func (self *Container) Terminate() error {

	url := "container/" + self.Uuid + "/"
	request := "DELETE"
	//Empty Body Request
	body := []byte(`{}`)

	_, err := TutumCall(url, request, body)
	if err != nil {
		return err
	}

	return nil
}