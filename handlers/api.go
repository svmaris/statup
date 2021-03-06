// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"os"
)

type ApiResponse struct {
	Object string `json:"type"`
	Method string `json:"method"`
	Id     int64  `json:"id"`
	Status string `json:"status"`
}

func ApiIndexHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	out := core.CoreApp
	json.NewEncoder(w).Encode(out)
}

func ApiRenewHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	var err error
	core.CoreApp.ApiKey = utils.NewSHA1Hash(40)
	core.CoreApp.ApiSecret = utils.NewSHA1Hash(40)
	core.CoreApp, err = core.UpdateCore(core.CoreApp)
	if err != nil {
		utils.Log(3, err)
	}
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

func ApiCheckinHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	checkin := core.FindCheckin(vars["api"])
	//checkin.Receivehit()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(checkin)
}

func ApiServiceHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	if service == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(service)
}

func ApiCreateServiceHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	var service *types.Service
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&service)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	newService := core.ReturnService(service)
	_, err = newService.Create()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(service)
}

func ApiServiceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	if service == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	var updatedService *types.Service
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&updatedService)
	service = core.ReturnService(updatedService)
	err := service.Update(true)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(service)
}

func ApiServiceDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	if service == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	err := service.Delete()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	output := ApiResponse{
		Object: "service",
		Method: "delete",
		Id:     service.Id,
		Status: "success",
	}
	json.NewEncoder(w).Encode(output)
}

func ApiAllServicesHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	services := core.CoreApp.Services()
	json.NewEncoder(w).Encode(services)
}

func ApiUserHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	user, err := core.SelectUser(utils.StringInt(vars["id"]))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func ApiUserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	user, err := core.SelectUser(utils.StringInt(vars["id"]))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	var updateUser *types.User
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&updateUser)
	user = core.ReturnUser(updateUser)
	err = user.Update()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func ApiUserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	user, err := core.SelectUser(utils.StringInt(vars["id"]))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	err = user.Delete()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	output := ApiResponse{
		Object: "user",
		Method: "delete",
		Id:     user.Id,
		Status: "success",
	}
	json.NewEncoder(w).Encode(output)
}

func ApiAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	users, _ := core.SelectAllUsers()
	json.NewEncoder(w).Encode(users)
}

func ApiCreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	var user *types.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	newUser := core.ReturnUser(user)
	uId, err := newUser.Create()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	output := ApiResponse{
		Object: "user",
		Method: "create",
		Id:     uId,
		Status: "success",
	}
	json.NewEncoder(w).Encode(output)
}

func isAPIAuthorized(r *http.Request) bool {
	if os.Getenv("GO_ENV") == "test" {
		return true
	}
	if IsAuthenticated(r) {
		return true
	}
	if isAuthorized(r) {
		return true
	}
	return false
}
