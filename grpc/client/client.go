package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "https://github.com/Yous113/Distributed-systems-/blob/main/grpc/school" 
)

func main() {
    // Set up a connection to the server.
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    // Create a client for the CourseService
    client := pb.NewCourseServiceClient(conn)

    // Create a context with a timeout
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // Call GetCourse method
    courseID := int32(1)
    res, err := client.GetCourse(ctx, &pb.GetCourseRequest{CourseId: courseID})
    if err != nil {
        log.Fatalf("could not get course: %v", err)
    }
    log.Printf("Course Details: ID=%d, Title=%s, Description=%s", res.Course.Id, res.Course.Title, res.Course.Description)

    // Call CreateCourse method
    newCourse := &pb.Course{
        Id:          2,
        Title:       "Cloud Computing",
        Description: "Introduction to cloud computing",
        StudentIds:  []int32{4, 5},
    }
    createRes, err := client.CreateCourse(ctx, &pb.CreateCourseRequest{Course: newCourse})
    if err != nil {
        log.Fatalf("could not create course: %v", err)
    }
    log.Printf("Created Course: ID=%d, Title=%s", createRes.Course.Id, createRes.Course.Title)
}
