package forms

import (
	"errors"
	"fmt"
	"github.com/mdl/fortycloud-sdk-go/internal"
)

type ConnectionsEndpoint struct {
	service *internal.JsonService
	url     string
}

type Connection struct {
	Type string `json:"_type"`
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	PeerA Peer `json:"peerA"`
	PeerANetwork string `json:"peerAnetwork"`
	PeerB Peer `json:"peerB`
	PeerBNetwork string `json:"peerBnetwork"`
	Active bool `json:"active"`
	Version int `json:"version"`
	KeyLifetime string `json:"keyLifetime"`
	ConnectionPolicy string `json:"connectionPolicy"`
	ConnectionState string `json:"connectionState"`
	DpdTimeout int `json:"dpdtimeout"`
	ForceNATT bool `json:"forceNATT"`
	Pfs bool `json:"pfs"`
	Phase2Lifetime string `json:"phase2Lifetime"`
	UpdownScript string `json:"updownScript"`
	BytesIncomingA string `json:"bytesIncomingA"`
	BytesIncomingB string `json:"bytesIncomingB"`
	BytesOutgoingA string `json:"bytesOutgoingA"`
	BytesOutgoingB string `json:"bytesOutgoingB"`
	State ConnectionState `json:"state"`
}
type Peer struct {
	Id int `json:"id"`
	Name string `json:"name"`
}
type ConnectionState struct {
	Id int `json:"id"`
	EncryptionAState string `json:"encryptionAstate"`
	EncryptionBState string `json:"encryptionBstate"`
	LastUpdateTimeA string `json:"lastUpdateTimeA"`
	LastUpdateTimeB string `json:"lastUpdateTimeB"`
	OverlayAstate string `json:"overlayAstate"`
	OverlayBstate string `json:"overlayBstate"`
	Version int `json:"version"`
}

func NewConnectionsEndpoint(service *internal.JsonService) *ConnectionsEndpoint {
	return &ConnectionsEndpoint{
		service: service,
		url:     "/EntityConnection",
	}
}

type ConnectionsMatch struct {
	PeerAId int `json:"peerA.id,omitempty"`
	PeerBId int `json:"peerB.id,omitempty"`
}
type connectionsAllRequest struct {
	Match ConnectionsMatch `json:"match"`
	Offset int `json:"offset"`
	Order string `json:"order"`
	OrderBy string `json:"orderBy"`
	Rows int `json:"rows"`
	Where []FilterClause `json:"where"`
}
type connectionsAllResult struct {
	EntityAllResult
	Objects []*Connection`json:"objects"`
}
func (endpoint *ConnectionsEndpoint) All(peerAId int, peerBId int, filters []FilterClause) ([]*Connection, error) {
	if filters == nil {
		filters = []FilterClause{}
	}
	body := &connectionsAllRequest{
		Match: ConnectionsMatch {
			PeerAId: peerAId,
			PeerBId: peerBId,
		},
		Offset: 0,
		Order: "DESC",
		OrderBy: "id",
		Rows: 100,
		Where: filters,
	}
	
	var result connectionsAllResult
	_, err := endpoint.service.Post(endpoint.url, body, &result)
	if err != nil {
		return nil, err
	}
	return result.Objects, nil
}

func (endpoint *ConnectionsEndpoint) Create(connection *Connection) (*Connection, error) {
	err := endpoint.put(&connectionsEndpointPutRequest{
		PeerAId: connection.PeerA.Id,
		PeerBId: connection.PeerB.Id,
	})
	if err != nil {
		return nil, err
	}
	
	conns, err2 := endpoint.All(connection.PeerA.Id, connection.PeerB.Id, nil)
	if err2 != nil {
		return nil, err2
	}
	
	if len(conns) <= 0 {
		return nil, errors.New("Could not get created connection.")
	}
	
	return conns[0], nil
}

type connectionDeleteResult struct {
	Result string `json:"result"`
	Total int `json:"total"`
}
func (endpoint *ConnectionsEndpoint) Delete(id int) error {
	var result connectionDeleteResult
	_, err := endpoint.service.Delete(endpoint.url, []int{id}, &result)
	if err != nil {
		return err
	}
	if result.Result != "OK" {
		return errors.New(fmt.Sprintf("Failed connection delete: %s", result.Result))
	}
	return nil
}

type connectionsEndpointPutRequest struct {
	PeerAId int `json:"peerA.id"`
	PeerBId int `json:"peerB.id"`
}
type connectionsEndpointPutResult struct {
	Result string `json:"result"`
	Total int `json:"total"`
}
func (endpoint *ConnectionsEndpoint) put(connection *connectionsEndpointPutRequest) error {
	var result connectionsEndpointPutResult
	_, err := endpoint.service.Put(endpoint.url, "", connection, &result)
	if err != nil {
		return err
	}
	if result.Result != "OK" {
		return errors.New(fmt.Sprintf("Failed connections put: %s", result.Result))
	}
	return nil
}