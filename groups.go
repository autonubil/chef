package chef

import "fmt"

type GroupsService struct {
	client *Client
}

 
// Group represents the native Go version of the deserialized Group type
type GroupAccess struct {
	Read      bool  `json:"read"`
	Create    bool  `json:"create"`
	Grant     bool  `json:"grant"`
	Update    bool  `json:"update"`
	Delete    bool  `json:"delete"`
}


// Group represents the native Go version of the deserialized Group type
type Group struct {
	Actors    []string      `json:"actors,omitempty"`
	Users     []string      `json:"users,omitempty"`
	Clients   []string      `json:"clients,omitempty"`
	Groups    []string      `json:"groups,omitempty"`
	Orgname   string        `json:"orgname"`
	Name string             `json:"name"`
	Groupname string        `json:"groupname"`
	Access    GroupAccess   `json:"access,omitempty"`
}

type GroupsListResult struct {	
	Groups []Group
} 
 

// String makes GroupsListResult implement the string result
func (e GroupsListResult) String() (out string) {
	out = fmt.Sprintf("%s\n",e.Groups)
	for _, group := range e.Groups {
		out += fmt.Sprintf("%s \n",group.Groupname)
	}
	return 
}

// String makes Group implement the string result
func (group Group) String() (out string) {
	out = fmt.Sprintf("%s [Actors: %v, Users: %v, Client: %v, Groups: %v] \n",group.Groupname, group.Actors, group.Users, group.Clients, group.Groups)
	return out
	
}



func NewGroupAccess(read, create, grant, update, delete  bool) GroupAccess {
    fmt.Printf("chef/groups:NewGroupAccess %v %v %v %v %v\n", read, create, grant, update, delete)
	return GroupAccess{
		Read:   read,
		Create: create,
        Grant:  grant,
        Update: update,
        Delete: delete,
	}
}


func NewGroup(actors, users, clients, groups []string, orgname, name, groupname string, access GroupAccess  ) Group {
    fmt.Printf("chef/groups:NewGroup %s %v %v\n", groupname, actors,access)

	return Group{
		Actors: actors,
		Clients: clients,
		Groups : groups,
		Orgname : orgname,
		Name : name,
		Groupname: groupname,
        Access: access,
	}
}



// Get gets a group from the Chef server.
//
func (e *GroupsService) Get(name string) (group Group, err error) {
	url := fmt.Sprintf("groups/%s", name)
	err = e.client.magicRequestDecoder("GET", url, nil, &group)
	return
}


// List lists the groups in the Chef server.
//
// https://docs.chef.io/api_chef_server.html#groups
func (e *GroupsService) List() (data GroupsListResult , err error) {
	err = e.client.magicRequestDecoder("GET", "groups", nil, &data)
	fmt.Printf("chef/groups:List %v\n", data)
	return
}

