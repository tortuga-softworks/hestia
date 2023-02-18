package registration

import (
	"context"
	"errors"
	"reflect"

	"github.com/tortuga-softworks/hestia/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegistrationServer struct {
	proto.UnimplementedRegistrationServer
	registrationService *RegistrationService
}

func NewRegistrationServer(registrationService *RegistrationService) (*RegistrationServer, error) {
	if registrationService == nil {
		return nil, errors.New("could not create a registration server: registration service is nil")
	}

	return &RegistrationServer{registrationService: registrationService}, nil
}

func (rs *RegistrationServer) SignUp(ctx context.Context, in *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	email := in.Email
	password := in.Password

	userID, err := rs.registrationService.SignUp(ctx, email, password)

	if err != nil {
		switch err.(type) {
		case *EmailFormatError:
			return nil, status.Error(codes.InvalidArgument, "email")
		case *PasswordFormatError:
			return nil, status.Error(codes.InvalidArgument, "password")
		case *EmailAlreadyExistsError:
			return nil, status.Error(codes.AlreadyExists, "email")
		default:
			return nil, status.Errorf(codes.Internal, "%v: %v", reflect.TypeOf(err), err)
		}
	} else {
		return &proto.SignUpResponse{UserId: userID}, nil
	}
}
