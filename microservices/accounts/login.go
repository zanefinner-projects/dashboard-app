package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type credentials struct {
	Username string
	Password string
	Token    string
}

func Login(w http.ResponseWriter, r *http.Request) {
	defer fmt.Println("POST: [/accounts/login]")
	var cr credentials
	err := json.NewDecoder(r.Body).Decode(&cr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	go fmt.Println("Retrieved", cr)
	cr.Token = "69abc-765"
	go fmt.Println("Sent Back token")
	fmt.Fprintf(w, cr.Token)

}
