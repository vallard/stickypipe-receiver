// Package handlers contains the handler logic for processing requests.
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vallard/stickypipe-receiver/app"
	"github.com/vallard/stickypipe-receiver/models"
)

// pipesHandle maintains the set of handlers for the pipes api.
type pipesHandle struct{}

// Pipes fronts the access to the pipes service functionality.
var Pipes pipesHandle

// List returns all the existing pipes in the system.
// 200 Success, 404 Not Found, 500 Internal
func (pipesHandle) List(c *app.Context) error {
	c.Respond(nil, http.StatusOK)
	return nil
}

// Create inserts a new user into the system.
// 200 OK, 400 Bad Request, 500 Internal
// curl -i -X POST -d "{ \"Switch\" : \"lx02\" }" localhost:3000/v1/pipes}
func (pipesHandle) Create(c *app.Context) error {
	var p models.Pipe

	if err := json.NewDecoder(c.Request.Body).Decode(&p); err != nil {
		fmt.Println("Error", err)
		return err
	}
	fmt.Println("This is P:", p)
	c.Respond(p, http.StatusOK)
	return nil
}
