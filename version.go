package jira

import (
	"bytes"
	"encoding/json"
	"github.com/go-jira/jira/jiradata"
)

// https://developer.atlassian.com/cloud/jira/platform/rest/v2#api-api-2-project-projectIdOrKey-versions-get
func (j *Jira) GetProjectVersions(project string) (*jiradata.Versions, error) {
	return GetProjectVersions(j.UA, j.Endpoint, project)
}

func GetProjectVersions(ua HttpClient, endpoint string, project string) (*jiradata.Versions, error) {
	uri := URLJoin(endpoint, "rest/api/2/project", project, "versions")
	resp, err := ua.GetJSON(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		results := jiradata.Versions{}
		return &results, json.NewDecoder(resp.Body).Decode(&results)
	}
	return nil, responseError(resp)
}

func (j *Jira) CreateProjectVersion(vp VersionProvider) (*jiradata.Version, error) {
	return CreateProjectVersion(j.UA, j.Endpoint, vp)
}

func CreateProjectVersion(ua HttpClient, endpoint string, vp VersionProvider) (*jiradata.Version, error) {
	uri := URLJoin(endpoint, "rest/api/2", "version")
	req := vp.ProvideVersion()
	encoded, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := ua.Post(uri, "application/json", bytes.NewBuffer(encoded))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		result := jiradata.Version{}
		return &result, json.NewDecoder(resp.Body).Decode(&result)
	}
	return nil, responseError(resp)
}

type VersionProvider interface {
	ProvideVersion() *jiradata.Version
}
