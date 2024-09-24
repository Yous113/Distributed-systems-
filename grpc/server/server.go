package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Yous113/Distributed-systems-/GRPC/school"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSchoolServer
}

func (s *server) GetCourse(ctx context.Context, req *pb.GetCourseRequest) (*pb.GetCourseResponse, error) {
	course := &pb.Course{
		Id:          req.CourseId,
		Title:       "Distributed Systems",
		Description: "An advanced course on distributed systems",
		StudentIds:  []int32{1, 2, 3},
	}
	return &pb.GetCourseResponse{Course: course}, nil
}

func (s *server) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (*pb.CreateCourseResponse, error) {
	return &pb.CreateCourseResponse{Course: req.Course}, nil
}

func (s *server) UpdateCourse(ctx context.Context, req *pb.UpdateCourseRequest) (*pb.UpdateCourseResponse, error) {
	return &pb.UpdateCourseResponse{Course: req.Course}, nil
}

func (s *server) DeleteCourse(ctx context.Context, req *pb.DeleteCourseRequest) (*pb.DeleteCourseResponse, error) {
	return &pb.DeleteCourseResponse{Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSchoolServer(s, &server{})

	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
