package tutum

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
)

var (
	User    string
	ApiKey  string
	BaseUrl = "https://app.tutum.co/api/v1/"
)

//Used to unpack the config file.
type Auth struct {
	User   string
	Apikey string
}
type config map[string]Auth

func init() {
	// Initialize base URL
	if os.Getenv("TUTUM_BASE_URL") != "" {
		BaseUrl = os.Getenv("TUTUM_BASE_URL")
	}

	// Initialize credentials
	LoadAuth()
}

func TutumCall(url string, requestType string) ([]byte, error) {
	if !IsAuthenticated() {
		return nil, fmt.Errorf("Couldn't find any Tutum credentials in ~/.tutum or environment variables TUTUM_USER and TUTUM_APIKEY")
	}
	client := &http.Client{}
	req, err := http.NewRequest(requestType, BaseUrl+url, nil)
	authHeader := fmt.Sprintf("ApiKey %s:%s", User, ApiKey)
	req.Header.Add("Authorization", authHeader)
	req.Header.Add("Accept", "application/json")
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Failed API call: %d ", response.Status)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func LoadAuth() error {
	if User != "" && ApiKey != "" {
		// Configuration already loaded
		log.Printf("Credentials for %s already loaded", User)
		return nil
	}

	// Process ~/.tutum configuration file first
	if usr, err := user.Current(); err == nil {
		var conf config
		confFilePath := usr.HomeDir + "/.tutum"
		if _, err := os.Stat(confFilePath); !os.IsNotExist(err) {
			if _, err := toml.DecodeFile(confFilePath, &conf); err == nil {
				if conf["auth"].User != "" && conf["auth"].Apikey != "" {
					User = conf["auth"].User
					ApiKey = conf["auth"].Apikey
					log.Printf("Loading credentials for %s from config file", User)
					return nil
				}
			} else {
				log.Printf("Malformed Tutum configuration file found at %s: %s", confFilePath, err)
				return fmt.Errorf("Malformed Tutum configuration file found at %s: %s", confFilePath, err)
			}
		}
	}

	// Load environment variables as an alternative option
	if os.Getenv("TUTUM_USER") != "" && os.Getenv("TUTUM_APIKEY") != "" {
		User = os.Getenv("TUTUM_USER")
		ApiKey = os.Getenv("TUTUM_APIKEY")
		log.Printf("Loading credentials for %s from environment", User)
		return nil
	}

	log.Print("Couldn't automatically load credentials")
	return fmt.Errorf("Couldn't find any Tutum credentials in ~/.tutum or environment variables TUTUM_USER and TUTUM_APIKEY")
}

func IsAuthenticated() bool {
	return (User != "" && ApiKey != "")
}
