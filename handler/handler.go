package handler

import (
	"gopls-workspace/libs"
	"net/http"
)

type HelmHandler interface {
	InstallHandler(w http.ResponseWriter, r *http.Request) error
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

func (c *helmHandler) InstallHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}
