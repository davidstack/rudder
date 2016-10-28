package controller

import (
	"github.com/AcalephStorage/rudder/client"
	log "github.com/Sirupsen/logrus"
	tiller "k8s.io/helm/pkg/proto/hapi/services"
)

type ReleaseController struct {
	tillerClient *client.TillerClient
}

func NewReleaseController(tillerClient *client.TillerClient) *ReleaseController {
	return &ReleaseController{tillerClient: tillerClient}
}

func (rc *ReleaseController) ListReleases(req *tiller.ListReleasesRequest) (*tiller.ListReleasesResponse, error) {
	res, err := rc.tillerClient.ListReleases(req)
	if err != nil {
		log.WithError(err).Error("unable to get list of releases from tiller")
		return nil, err
	}
	return res, nil
}