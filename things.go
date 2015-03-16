package things

import (
	"fmt"
	"time"
	"bytes"
	"net/http"
	"encoding/json"
)

const BaseUrl = "https://api.thethings.io/v2/things"

var debug = true

func SetDebugMode(b bool) {
	debug = b
}

// Activate a thing through code provided by thethings.io
func Activate(code string) (*http.Response, error) {
	var buf []byte
	var resp *http.Response
	var err error
	
	data := struct {
		ActivationCode string `json:"activationCode"`
	}{
		code,
	}
	
	if buf, err = json.Marshal(data); err != nil {
		return nil, fmt.Errorf("Activate: %v", err)
	}
	
	r := bytes.NewReader(buf)

	if resp, err = http.Post(BaseUrl, "application/json", r); err != nil {
		return nil, fmt.Errorf("Activate: %v", err)
	}
	
	return resp, nil
}

type KeyValue struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type ReadValue struct {
	Value string `json:"value"`
	Datetime string `json:"datetime"`
}

// Handler function to call when there is real-time data update
type SubscribeHandlerFunc func(data []KeyValue)

type Thing struct {
	Token string
}

// Read all values from a thing and return an http.Response
func (t *Thing) Resources() (*http.Response, error) {
	var resp *http.Response
	var err error
	
	if resp, err = http.Get(BaseUrl + "/" + t.Token + "/resources"); err != nil {
		return nil, fmt.Errorf("Thing.Resources: %v", err)
	}
	
	return resp, nil
}

// Read from a thing and return an http.Response
func (t *Thing) Read(key string) (*http.Response, error) {
	var resp *http.Response
	var err error
	
	if resp, err = http.Get(BaseUrl + "/" + t.Token + "/resources/" + key); err != nil {
		return nil, fmt.Errorf("Thing.Read: %v", err)
	}
	
	return resp, nil
}

// Write key-values to a thing and return an http.Response
func (t *Thing) Write(keyValues map[string]string) (*http.Response, error) {
	var buf []byte
	var resp *http.Response
	var err error
	
	data := struct {
		KeyValues []KeyValue `json:"values"`
	}{}

	for k, v := range keyValues {
		data.KeyValues = append(data.KeyValues, KeyValue{Key: k, Value: v})
	}
	
	if buf, err = json.Marshal(data); err != nil {
		return nil, fmt.Errorf("Thing.Write: %v", err)
	}
	
	r := bytes.NewReader(buf)
	
	if debug {
		fmt.Println(string(buf))
	}
	
	if resp, err = http.Post(BaseUrl + "/" + t.Token, "application/json", r); err != nil {
		return nil, fmt.Errorf("Thing.Write: %v", err)
	}
	
	return resp, nil
}

// Read value from a thing as string
func (t *Thing) ReadValue(key string) (string, error) {
	var data []ReadValue

	resp, err := t.Read(key)
	if err != nil {
		return "", fmt.Errorf("Thing.ReadValue: %v", err)
	}
	
	json.NewDecoder(resp.Body).Decode(&data)
	return data[0].Value, nil
}

// Write key-value as strings to a thing
func (t *Thing) WriteValue(key, value string) error {
	_, err := t.Write(map[string]string{key: value})
	if err != nil {
		return fmt.Errorf("Thing.WriteValue: %v", err)
	}
	return nil
}

// Subscribe to a thing's updates in real-time
func (t *Thing) Subscribe(fn SubscribeHandlerFunc) (*http.Response, error) {
	var resp *http.Response
	var err error
	
	if resp, err = http.Get(BaseUrl + "/" + t.Token); err != nil {
		return nil, fmt.Errorf("Thing.Subscribe: %v", err)
	}

	go func() {
		var buf = make([]byte, 4096)
		for {
			n, _ := resp.Body.Read(buf)
			if n > 0 {
				var keyValues []KeyValue
				json.Unmarshal(buf[:n], &keyValues)
				if len(keyValues) > 0 {
					fn(keyValues)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()
	
	return resp, nil
}
