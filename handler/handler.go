package handler

import (
	"encoding/json"
	"fmt"
	"gopls-workspace/libs"
	"log"
	"net/http"
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

type KeycloakResponse struct {
	Status         int
	keycloakDomain string
	namespace      string
}

// post requrst
// body { 'namespace': string, 'keycloakDomain' : string }
func (c *helmHandler) InstallHandler(w http.ResponseWriter, r *http.Request) {
	ns := r.FormValue("namespace")
	domain := r.FormValue("keycloakDomain")

	err := c.K8SClient.CreateNameSpace(ns)
	if err != nil {
		log.Println(err.Error())
		errorHandler(w, r, err)
		return
	}
	libs.UpdateEnvValue("domain", domain)
	libs.UpdateEnvValue("namespace", ns)
	err = c.HelmGo.Install(ns)
	if err != nil {
		log.Println(err.Error())
		errorHandler(w, r, err)
		return
	}

	keycloak := KeycloakResponse{http.StatusOK, domain, ns}
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
