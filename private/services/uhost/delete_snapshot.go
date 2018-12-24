//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UHost DeleteSnapshot

package uhost

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DeleteSnapshotRequest is request schema for DeleteSnapshot action
type DeleteSnapshotRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"true"`

	// 快照Id
	SnapshotId *string `required:"true"`
}

// DeleteSnapshotResponse is response schema for DeleteSnapshot action
type DeleteSnapshotResponse struct {
	response.CommonBase

	// 快照Id
	SnapshotId string
}

// NewDeleteSnapshotRequest will create request of DeleteSnapshot action.
func (c *UHostClient) NewDeleteSnapshotRequest() *DeleteSnapshotRequest {
	req := &DeleteSnapshotRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteSnapshot - 删除快照
func (c *UHostClient) DeleteSnapshot(req *DeleteSnapshotRequest) (*DeleteSnapshotResponse, error) {
	var err error
	var res DeleteSnapshotResponse

	err = c.client.InvokeAction("DeleteSnapshot", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}