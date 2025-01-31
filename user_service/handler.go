package main

import (
	pb "go_projects/learngo/shippy_demo/user_service/proto/user"
	"golang.org/x/net/context"
)

type service struct {
	repo         Repository
	tokenService Authable
}

func (s *service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	if err := s.repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (s *service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := s.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (s *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (s *service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	_, err := s.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}
	res.Token = "testingabc"
	return nil
}

func (s *service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	return nil
}
