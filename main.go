package main

import (
"encoding/json"
"github.com/gorilla/mux"
"net/http"
"os"
"strconv"
)

type User struct {
	Id   string
	Name string
}

var userMap map[string]User

func main() {
	userMap = make(map[string]User)

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	print("test push")
	print("test push2")
	print("test push3")


	router := mux.NewRouter()
	router.HandleFunc("/user", addUser).Methods(http.MethodPost)
	router.HandleFunc("/user", updateUser).Methods(http.MethodPut)

	router.HandleFunc("/user/info/{id}", getUser).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", deleteUser).Methods(http.MethodDelete)
	http.ListenAndServe(":"+port, router)
}

func deleteUser(writer http.ResponseWriter, request *http.Request) {
	pathVars := mux.Vars(request)
	id := pathVars["id"]

	delete(userMap, id)

	json.NewEncoder(writer).Encode(map[string]string{
		"message": "delete user {" + id + "} success",
	})
}

func updateUser(writer http.ResponseWriter, request *http.Request) {
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	userId := user.Id
	userMap[userId] = User{userId, user.Name}

	json.NewEncoder(writer).Encode(map[string]string{
		"id":   user.Id,
		"name": "update user name:" + user.Name,
	})
}

func getUser(writer http.ResponseWriter, request *http.Request) {
	pathVars := mux.Vars(request)
	id := pathVars["id"]

	user := userMap[id]

	json.NewEncoder(writer).Encode(map[string]string{
		"id":   user.Id,
		"name": user.Name,
	})
}

func addUser(writer http.ResponseWriter, request *http.Request) {
	var newUser User
	err := json.NewDecoder(request.Body).Decode(&newUser)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	userCount := len(userMap) + 1

	newUserId := strconv.FormatInt(int64(userCount), 10)

	newUser = User{newUserId, newUser.Name}

	userMap[newUserId] = newUser

	json.NewEncoder(writer).Encode(map[string]string{
		"id":   newUser.Id,
		"name": newUser.Name,
	})

}

