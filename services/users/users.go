package users

import (
	"context"

	"gitea.dancheg97.ru/dancheg97/regen/gen/pb"
	"gitea.dancheg97.ru/dancheg97/regen/gen/sqlc"
	"gitea.dancheg97.ru/dancheg97/regen/postgres"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrUnknown = status.Error(codes.NotFound, `unknown error`)

type Handlers struct {
	*postgres.Postgres
}

func (s Handlers) Create(ctx context.Context, in *pb.User) (*pb.User, error) {
	id, err := s.Postgres.InsertUser(ctx, sqlc.InsertUserParams{
		Description: in.Description,
		Name:        in.Name,
		Age:         in.Age,
	})
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:          id,
		Name:        in.Name,
		Age:         in.Age,
		Description: in.Description,
	}, nil
}

func (s Handlers) List(in *pb.Empty, str pb.UserStorage_ListServer) error {
	users, err := s.Postgres.SelectUsers(str.Context())
	if err != nil {
		return ErrUnknown
	}
	for _, sur := range users {
		err = str.Send(&pb.User{
			Id:          sur.ID,
			Name:        sur.Name,
			Age:         sur.Age,
			Description: sur.Description,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Handlers) Remove(ctx context.Context, in *pb.Id) (*pb.Empty, error) {
	err := s.Postgres.DeleteUser(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s Handlers) Update(ctx context.Context, in *pb.User) (*pb.User, error) {
	err := s.Postgres.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:          in.Id,
		Name:        in.Name,
		Age:         in.Age,
		Description: in.Description,
	})
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:          in.Id,
		Name:        in.Name,
		Age:         in.Age,
		Description: in.Description,
	}, nil
}
