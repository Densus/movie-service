package main

import (
	"github.com/densus/movie_service/config"
	"github.com/densus/movie_service/delivery/grpc/handler"
	"github.com/densus/movie_service/delivery/http"
	"github.com/densus/movie_service/repository"
	external_service "github.com/densus/movie_service/service/external-service"
	internal_service "github.com/densus/movie_service/service/internal-service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main()  {
	db := config.SetupDBConnection()
	defer config.CloseDBConnection(db)

	repo := repository.NewMovieRepository(db)

	internalService := internal_service.NewInternalService(repo)
	externalService := external_service.NewExternalService(repo)

	r := gin.Default()

	http.NewMovieController(r, externalService, internalService)

	startGRPCServer(externalService, internalService)
	r.Run(":8080")


}

//func (*server) search(ctx context.Context, r *movie.) (*basketBallPlayer.PlayerResponse, error) {
//	id := r.GetId()
//
//	playerVal, found := c.Get(id)
//	if found {
//		player := playerVal.(basketBallPlayer.Player)
//		return &basketBallPlayer.PlayerResponse{
//			Result: &player,
//		}, nil
//	}
//
//	return nil, fmt.Errorf("could not find player with id: %v", id)
//}

func startGRPCServer(externalService external_service.ExternalService, internalService internal_service.InternalService) {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("could not attach listener to port: %v", err)
	}

	s := grpc.NewServer()
	handler.NewMovieServerGrpc(s, externalService, internalService)
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("could not start grpc server: %v", err)
		}
	}()
}
