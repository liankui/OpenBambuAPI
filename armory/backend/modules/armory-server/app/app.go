package app

import (
	"context"

	"github.com/chaos-io/chaos/logs"
	"github.com/golang/protobuf/proto"
	http2 "github.com/liankui/OpenBambuAPI/armory/backend/infra/http"

	// this service api
	pb "github.com/liankui/OpenBambuAPI/armory/go/armory/v1"
)

const baseUrl = "https://api.bambulab.com"

type armoryServer struct {
	pb.UnimplementedArmoryServer

	client http2.IClient
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.ArmoryServer {
	hc := http2.NewClient()
	return armoryServer{client: hc}
}

// Login implements Interface.
func (s armoryServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	logs.Infow("Login", "request", in)

	reqParam := &http2.RequestParam{
		RequestURI: "",
		Method:     "",
		Header:     nil,
		Timeout:    0,
		Cluster:    nil,
		WithSD:     nil,
	}

	err := s.client.DoHTTPRequest(ctx, reqParam)
	if err != nil {
		return nil, logs.NewErrorw("failed to login", "error", err)
	}

	resp := &pb.LoginResponse{
		// AccessToken:
		// RefreshToken:
		// LoginType:
		// ExpiresIn:
	}

	if reqParam.Response != nil {
		proto.Merge(resp, reqParam.Response)
	}

	return resp, nil
}

// RefreshToken implements Interface.
func (s armoryServer) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	logs.Infow("RefreshToken", "request", in)

	resp := &pb.RefreshTokenResponse{
		// AccessToken:
		// RefreshToken:
		// ExpiresIn:
		// RefreshExpiresIn:
	}
	return resp, nil
}

// ListMyMessages implements Interface.
func (s armoryServer) ListMyMessages(ctx context.Context, in *pb.ListMyMessagesRequest) (*pb.ListMyMessagesResponse, error) {
	logs.Infow("ListMyMessages", "request", in)

	resp := &pb.ListMyMessagesResponse{
		// Hits:
	}
	return resp, nil
}

// ListMyTasks implements Interface.
func (s armoryServer) ListMyTasks(ctx context.Context, in *pb.ListMyTasksRequest) (*pb.ListMyTasksResponse, error) {
	logs.Infow("ListMyTasks", "request", in)

	resp := &pb.ListMyTasksResponse{
		// Total:
		// Hits:
	}
	return resp, nil
}

// GetSlicerResources implements Interface.
func (s armoryServer) GetSlicerResources(ctx context.Context, in *pb.GetSlicerResourcesRequest) (*pb.GetSlicerResourcesResponse, error) {
	logs.Infow("GetSlicerResources", "request", in)

	resp := &pb.GetSlicerResourcesResponse{
		// Base:
		// Software:
		// Resources:
	}
	return resp, nil
}

// ListUserDevices implements Interface.
func (s armoryServer) ListUserDevices(ctx context.Context, in *pb.ListUserDevicesRequest) (*pb.ListUserDevicesResponse, error) {
	logs.Infow("ListUserDevices", "request", in)

	resp := &pb.ListUserDevicesResponse{
		// Base:
		// Devices:
	}
	return resp, nil
}

// GetPrintStatus implements Interface.
func (s armoryServer) GetPrintStatus(ctx context.Context, in *pb.GetPrintStatusRequest) (*pb.GetPrintStatusResponse, error) {
	logs.Infow("GetPrintStatus", "request", in)

	resp := &pb.GetPrintStatusResponse{
		// Base:
		// Devices:
	}
	return resp, nil
}

// GetTtcode implements Interface.
func (s armoryServer) GetTtcode(ctx context.Context, in *pb.GetTTCodeRequest) (*pb.GetTTCodeResponse, error) {
	logs.Infow("GetTtcode", "request", in)

	resp := &pb.GetTTCodeResponse{
		// Base:
		// Ttcode:
		// Passwd:
		// Authkey:
	}
	return resp, nil
}
