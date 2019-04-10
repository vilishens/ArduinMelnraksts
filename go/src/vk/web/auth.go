package web

import (
	vauth "github.com/apexskier/httpauth"
)

var (
	backend     vauth.LeveldbAuthBackend
	aaa         vauth.Authorizer
	roles       map[string]vauth.Role
	backendfile = "../cfg/factory/passes.leveldb"
)

func init() {
	// create the backend
	backend, err := vauth.NewLeveldbAuthBackend(backendfile)
	if err != nil {
		panic(err)
	}

	roles = make(map[string]vauth.Role)
	roles["user"] = 30
	roles["admin"] = 80
	aaa, err = vauth.NewAuthorizer(backend, []byte("cookie-encryption-key"), "user", roles)

	// create a default user
	username := "admin"
	defaultUser := vauth.UserData{Username: username, Role: "admin"}
	err = backend.SaveUser(defaultUser)
	if err != nil {
		panic(err)
	}
	// Update user with a password and email address
	err = aaa.Update(nil, nil, username, "adminadmin", "admin@localhost.com")
	if err != nil {
		panic(err)
	}
}
