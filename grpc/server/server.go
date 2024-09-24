package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Yous113/Distributed-systems-/grpc/school" // Import the generated code
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCourseServiceServer  // Embedding the generated gRPC server for CourseService
	pb.UnimplementedStudentServiceServer // Embedding the generated gRPC server for StudentService
	pb.UnimplementedTeacherServiceServer // Embedding the generated gRPC server for TeacherService
}

// Implement methods for CourseService
func (s *server) GetCourse(ctx context.Context, req *pb.GetCourseRequest) (*pb.GetCourseResponse, error) {
	course := &pb.Course{
		Id:          req.CourseId,
		Title:       "Distributed Systems",
		Description: "An advanced course on distributed systems",
		StudentIds:  []int32{1, 2, 3},
	}
	return &pb.GetCourseResponse{Course: course}, nil
}

// Implement other services similarly...

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCourseServiceServer(s, &server{})  // Register CourseService
	pb.RegisterStudentServiceServer(s, &server{}) // Register StudentService
	pb.RegisterTeacherServiceServer(s, &server{}) // Register TeacherService

	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
