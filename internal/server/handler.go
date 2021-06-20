package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kaz/itms-api/itms"
)

var (
	token = os.Getenv("TOKEN")
)

func HandleControl(instrument string, action string) error {
	if instrument[0] == 'b' && action == "on" {
		return itms.TurnOnBath()
	} else if instrument[0] == 'b' && action == "off" {
		return itms.TurnOffBath()
	} else if instrument[0] == 'f' && action == "on" {
		return itms.TurnOnFloorHeating()
	} else if instrument[0] == 'f' && action == "off" {
		return itms.TurnOffFloorHeating()
	}
	return fmt.Errorf("no such instrument/action")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Printf("error: ParseForm: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if token == "" || token != r.Form["token"][0] {
		fmt.Printf("unauthorized access: remote_addr=%s", r.RemoteAddr)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err := HandleControl(r.Form["instrument"][0], r.Form["action"][0]); err != nil {
		fmt.Printf("error: HandleControl: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
