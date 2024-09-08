package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func Perform(client *http.Client, req *http.Request, username string, password string) (*http.Response, error) {
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	digest, err := ParseDigestResponse(res.Header.Get("WWW-Authenticate"))

	if err != nil {
		return nil, err
	}

	auth, err := BuildDigestHeader(req.Method, req.URL.Path, digest.Realm, digest.Nonce, username, password)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", auth)
	return client.Do(req)
}

func GenericGet[Res any](host string, path string, username string, password string) (*Res, error) {
	client := &http.Client{}

	url := fmt.Sprintf("http://%s%s", host, path)
	req, _ := http.NewRequest("GET", url, nil)

	res, err := Perform(client, req, username, password)

	if err != nil {
		return nil, err
	}

	if res.StatusCode == 204 {
		return nil, nil
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var response Res
	err = json.Unmarshal(body, &response)
	return &response, err
}

func GenericDelete(host string, path string, username string, password string) error {
	client := &http.Client{}

	url := fmt.Sprintf("http://%s%s", host, path)
	req, _ := http.NewRequest("DELETE", url, nil)

	res, err := Perform(client, req, username, password)

	if err != nil {
		return err
	}

	if res.StatusCode != 204 {
		return errors.New("unexpected http response")
	}

	return nil
}

func GetStatus(host string, username string, password string) (*StatusResponse, error) {
	return GenericGet[StatusResponse](host, "/api/v1/status", username, password)
}

func GetJob(host string, username string, password string) (*JobResponse, error) {
	return GenericGet[JobResponse](host, "/api/v1/job", username, password)
}

func GetStorage(host string, username string, password string) (*StorageResponse, error) {
	return GenericGet[StorageResponse](host, "/api/v1/storage", username, password)
}

func GetFiles(storage string, host string, username string, password string) (*FilesResponse, error) {
	return GenericGet[FilesResponse](host, fmt.Sprintf("/api/v1/files/%s", storage), username, password)
}

func DeleteFile(file string, host string, username string, password string) error {
	return GenericDelete(host, fmt.Sprintf("/api/v1/files/%s", file), username, password)
}
