package core

import (
	"encoding/json"
	"io/ioutil"
)

type Project struct {
	Name      string     `json:"name"`
	Path      string     `json:"-"`
	Sequences []Sequence `json:"sequences"`
}

func NewProject(path string) (*Project, error) {
	project := &Project{
		Path: path,
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, project)
	if err != nil {
		return nil, err
	}

	return project, nil
}
