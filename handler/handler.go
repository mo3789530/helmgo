package handler

import (
	"encoding/json"
	"fmt"
	"helmgo/libs"
	"log"
	"net/http"
	"os"
)

type HelmHandler interface {
	InstallHandler(w http.ResponseWriter, r *http.Request)
	HealthHandler(w http.ResponseWriter, r *http.Request)
}

type helmHandler struct {
	K8SClient libs.K8SClient
	HelmGo    libs.HelmGo
}

func NewHelmHandler(k8s libs.K8SClient, helm libs.HelmGo) HelmHandler {
	return &helmHandler{
		K8SClient: k8s,
		HelmGo:    helm,
	}
}

type KeycloakRequest struct {
	KeycloakDomain string `json:"keycloakDoamin"`
	Namespace      string `json:"namespace"`
}

type KeycloakResponse struct {
	Status         int
	KeycloakDomain string
	Namespace      string
}

// post requrst
// body { 'namespace': string, 'keycloakDomain' : string }
// curl  -X POST http://localhost:8080/api/helm/keycloak  -d '{ "namespace": "test", "keycloakDoamin" : "auth1" }'  -H "Content-Type: application/json"
func (c *helmHandler) InstallHandler(w http.ResponseWriter, r *http.Request) {

	var request KeycloakRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		errorHandler(w, r, err)
		return
	}
	url := request.KeycloakDomain + "." + os.Getenv("DOMAIN")
	err := c.HelmGo.Install(request.Namespace, url, nil)
	if err != nil {
		log.Println(err.Error())
		errorHandler(w, r, err)
		return
	}

	keycloak := KeycloakResponse{http.StatusOK, request.KeycloakDomain, request.Namespace}
	res, err := json.Marshal(keycloak)
	if err != nil {
		log.Println(err.Error())
		errorHandler(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = err.Error()
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s", err)
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write(jsonResp)
}

func (c *helmHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}
