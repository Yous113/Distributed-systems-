package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Yous113/Distributed-systems-/grpc/school"
	"google.golang.org/grpc"
)

type server struct {
	pb.unimplementedSchoolServer
}

func (s *courseServer) GetCourse(ctx context.Context, req *pb.GetCourseRequest) (*pb.GetCourseResponse, error) {
	// Simulate getting course data
	course := &pb.Course{
		Id:          req.CourseId,
		Title:       "Distributed Systems",
		Description: "An advanced course on distributed systems",
		StudentIds:  []int32{1, 2, 3},
	}
	return &pb.GetCourseResponse{Course: course}, nil
}

// Implement the CreateCourse method
func (s *courseServer) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (*pb.CreateCourseResponse, error) {
	return &pb.CreateCourseResponse{Course: req.Course}, nil
}

// Implement the UpdateCourse method
func (s *courseServer) UpdateCourse(ctx context.Context, req *pb.UpdateCourseRequest) (*pb.UpdateCourseResponse, error) {
	return &pb.UpdateCourseResponse{Course: req.Course}, nil
}

// Implement the DeleteCourse method
func (s *courseServer) DeleteCourse(ctx context.Context, req *pb.DeleteCourseRequest) (*pb.DeleteCourseResponse, error) {
	return &pb.DeleteCourseResponse{Success: true}, nil
}

func main() {
	// Create a listener on TCP port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server
	s := grpc.NewServer()

	// Register the CourseService with the server
	pb.RegisterCourseServiceServer(s, &courseServer{})

	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
