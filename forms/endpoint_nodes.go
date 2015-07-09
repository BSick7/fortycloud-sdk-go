package forms

import (
	"github.com/mdl/fortycloud-sdk-go/internal"
)

type NodesEndpoint struct {
	service *internal.JsonService
	url     string
}

type Node struct {
	Type               string `json:"_type"`
	Id                 int    `json:"id"`
	EnrollingTokenUUID string `json:"enrollingTokenUUID"`
	IsGW               bool   `json:"isGW"`
	Name               string `json:"name"`
	PublicIP           string `json:"publicIPD"`
	PrivateIP          string `json:"privateIP"`
	UserVPNIP          string `json:"userVPNIP"`
	LocalIPForUserVPN  string `json:"localIPForUserVPN"`
	Active             bool   `json:"active"`
	AgentRelease       string `json:"agentRelease"`
	OverlayAddress     string `json:"overlayAddress"`
	OverlayAddressLong int64  `json:"overlayAddressLong"`
}

func NewNodesEndpoint(service *internal.JsonService) *NodesEndpoint {
	return &NodesEndpoint{
		service: service,
		url:     "/EntityNode",
	}
}

type nodesAllRequest struct {
	Match   ConnectionsMatch `json:"match"`
	Offset  int              `json:"offset"`
	Order   string           `json:"order"`
	OrderBy string           `json:"orderBy"`
	Rows    int              `json:"rows"`
	Where   []FilterClause   `json:"where"`
}
type nodesAllResult struct {
	EntityAllResult
	Objects []*Node `json:"objects"`
}

func (endpoint *NodesEndpoint) All(filters []FilterClause) ([]*Node, error) {
	if filters == nil {
		filters = []FilterClause{}
	}
	body := &nodesAllRequest{
		Offset:  0,
		Order:   "DESC",
		OrderBy: "id",
		Rows:    100,
		Where:   filters,
	}

	var result nodesAllResult
	_, err := endpoint.service.Post(endpoint.url, body, &result)
	if err != nil {
		return nil, err
	}
	return result.Objects, nil
}

func (endpoint *NodesEndpoint) GetByPublicIP(publicIP string) (*Node, error) {
	nodes, err := endpoint.All([]FilterClause{
		NewFilterLike("public_ip", publicIP),
	})
	if err != nil {
		return nil, err
	}
	if len(nodes) <= 0 {
		return nil, nil
	}
	return nodes[0], nil
}