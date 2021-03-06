package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/identity/v3/groups"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

// ListOutput provides a single page of Group results.
const ListOutput = `
{
    "links": {
        "next": null,
        "previous": null,
        "self": "http://example.com/identity/v3/groups"
    },
    "groups": [
        {
            "domain_id": "default",
            "id": "2844b2a08be147a08ef58317d6471f1f",
            "description": "group for internal support users",
            "links": {
                "self": "http://example.com/identity/v3/groups/2844b2a08be147a08ef58317d6471f1f"
            },
            "name": "internal support",
            "extra": {
                "email": "support@localhost"
            }
        },
        {
            "domain_id": "1789d1",
            "id": "9fe1d3",
            "description": "group for support users",
            "links": {
                "self": "https://example.com/identity/v3/groups/9fe1d3"
            },
            "name": "support",
            "extra": {
                "email": "support@example.com"
            }
        }
    ]
}
`

// GetOutput provides a Get result.
const GetOutput = `
{
    "group": {
        "domain_id": "1789d1",
        "id": "9fe1d3",
        "description": "group for support users",
        "links": {
            "self": "https://example.com/identity/v3/groups/9fe1d3"
        },
        "name": "support",
        "extra": {
            "email": "support@example.com"
        }
    }
}
`

// FirstGroup is the first group in the List request.
var FirstGroup = groups.Group{
	DomainID: "default",
	ID:       "2844b2a08be147a08ef58317d6471f1f",
	Links: map[string]interface{}{
		"self": "http://example.com/identity/v3/groups/2844b2a08be147a08ef58317d6471f1f",
	},
	Name:        "internal support",
	Description: "group for internal support users",
	Extra: map[string]interface{}{
		"email": "support@localhost",
	},
}

// SecondGroup is the second group in the List request.
var SecondGroup = groups.Group{
	DomainID: "1789d1",
	ID:       "9fe1d3",
	Links: map[string]interface{}{
		"self": "https://example.com/identity/v3/groups/9fe1d3",
	},
	Name:        "support",
	Description: "group for support users",
	Extra: map[string]interface{}{
		"email": "support@example.com",
	},
}

// ExpectedGroupsSlice is the slice of groups expected to be returned from ListOutput.
var ExpectedGroupsSlice = []groups.Group{FirstGroup, SecondGroup}

// HandleListGroupsSuccessfully creates an HTTP handler at `/groups` on the
// test handler mux that responds with a list of two groups.
func HandleListGroupsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ListOutput)
	})
}

// HandleGetGroupSuccessfully creates an HTTP handler at `/groups` on the
// test handler mux that responds with a single group.
func HandleGetGroupSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/groups/9fe1d3", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, GetOutput)
	})
}
