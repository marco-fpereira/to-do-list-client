package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	inputAdapter "github.com/marco-fpereira/to-do-list-client/app/adapters/input"
	outputAdapter "github.com/marco-fpereira/to-do-list-client/app/adapters/output"
	pb "github.com/marco-fpereira/to-do-list-client/app/config/grpc"
	domain "github.com/marco-fpereira/to-do-list-client/app/domain/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	grpcPort := os.Getenv("GRPC_PORT")

	conn, err := grpc.NewClient(
		(":" + grpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("error starting grpc client. Details: %v", err)
	}
	defer conn.Close()
	accountClient := pb.NewAccountClient(conn)
	taskClient := pb.NewTaskClient(conn)

	outputPort := outputAdapter.NewToDoServerAdapter(accountClient, taskClient)
	authenticatorPort := outputAdapter.NewJwtAuthenticationAdapter()
	accountPort := domain.NewAccountUseCase(outputPort, authenticatorPort)
	taskPort := domain.NewTaskUseCase(outputPort, authenticatorPort)
	accountAdapter := inputAdapter.NewAccountAdapter(accountPort)
	taskAdapter := inputAdapter.NewTaskAdapter(taskPort)

	r := initRouter(accountAdapter, taskAdapter)
	err = r.Run(":" + httpPort)
	if err != nil {
		log.Fatalf("error starting http server. Details: %v", err)
	}
}

func initRouter(
	accountAdapter inputAdapter.AccountAdapterInterafece,
	taskAdapter inputAdapter.TaskAdapterInterface,
) *gin.Engine {
	r := gin.Default()
	r.POST("/signup", accountAdapter.Signup)
	r.POST("/login", accountAdapter.Login)
	r.GET("/users/:userId/tasks", taskAdapter.GetAllTasks)
	r.POST("/users/:userId/tasks", taskAdapter.CreateTask)
	r.PUT("/users/:userId/tasks/:taskId", taskAdapter.UpdateTaskMessage)
	r.PATCH("/users/:userId/tasks/:taskId/completeness", taskAdapter.UpdateTaskCompleteness)
	r.DELETE("/users/:userId/tasks/:taskId", taskAdapter.DeleteTask)
	return r
}
