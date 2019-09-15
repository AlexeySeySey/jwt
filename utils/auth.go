package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	var creds Creds
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !UserContains(Users, creds) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	token, err := GenerateToken(time.Now().Add(5 * time.Minute))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func Home(w http.ResponseWriter, r *http.Request){
	var tknStr AccessToken
	if err := json.NewDecoder(r.Body).Decode(&tknStr); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	claims := &Claims{}
	_, err := ParseToken(tknStr.Value, claims)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Login)))
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	var tknStr AccessToken
	if err := json.NewDecoder(r.Body).Decode(&tknStr); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	claims := &Claims{}
	_, err := ParseToken(tknStr.Value, claims)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	token, err := GenerateToken(time.Now().Add(5 * time.Minute))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(token)
}
