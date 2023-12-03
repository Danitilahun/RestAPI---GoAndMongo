package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"go-mongo-crud/models"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}


// GET /users/:id - Get a user by ID
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userID := p.ByName("id")
	if !bson.IsObjectIdHex(userID) {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	session := uc.session.Copy()
	defer session.Close()

	collection := session.DB("mydatabase").C("users")

	var user models.User
	err := collection.FindId(bson.ObjectIdHex(userID)).One(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// POST /users - Create a user
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	session := uc.session.Copy()
	defer session.Close()

	collection := session.DB("mydatabase").C("users")

	newUser.ID = bson.NewObjectId()
	err := collection.Insert(newUser)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// PUT /users/:id - Update a user by ID
func (uc UserController) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userID := p.ByName("id")
	if !bson.IsObjectIdHex(userID) {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	session := uc.session.Copy()
	defer session.Close()

	collection := session.DB("mydatabase").C("users")

	err := collection.UpdateId(bson.ObjectIdHex(userID), bson.M{"$set": updatedUser})
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}

// DELETE /users/:id - Delete a user by ID
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userID := p.ByName("id")
	if !bson.IsObjectIdHex(userID) {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	session := uc.session.Copy()
	defer session.Close()

	collection := session.DB("mydatabase").C("users")

	err := collection.RemoveId(bson.ObjectIdHex(userID))
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
