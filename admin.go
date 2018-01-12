package gapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)


func (c *Client) CreateUserForm(email string, login string, name string, password string) error {
	data, err := json.Marshal(map[string]string{
		"email": email,
		"login": login,
		"name": name,
		"password":password,
	})
	if err != nil {
		return err
	}
	req, err := c.newRequest("POST", "/api/admin/users", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return err
}

func (c *Client) DeleteUser(id int64) error {
	req, err := c.newRequest("DELETE", fmt.Sprintf("/api/admin/users/%d", id), nil)
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return err
}
