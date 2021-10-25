package handler

import (
	"context"
	"fmt"
	grpc2 "github.com/densus/movie_service/delivery/grpc"
	"github.com/densus/movie_service/entity"
	"github.com/densus/movie_service/entity/dto"
	external_service "github.com/densus/movie_service/service/external-service"
	internal_service "github.com/densus/movie_service/service/internal-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewMovieServerGrpc(gServer *grpc.Server, externalServ external_service.ExternalService, internalServ internal_service.InternalService)  {
	handler := &server{
		externalService: externalServ,
		internalService: internalServ,
	}
	grpc2.RegisterMovieHandlerServer(gServer, handler)

	reflection.Register(gServer) //exposes all the publicly accessible gRPC services on a gRPC server.
}

type server struct {
	externalService external_service.ExternalService
	internalService internal_service.InternalService
}

func (s *server) Search(request *grpc2.SearchRequest, searchServer grpc2.MovieHandler_SearchServer) error {
	//list := s.externalService.Search(request.GetSearchWord(), request.GetPagination())
	//arrMovie := make([]*grpc2.SearchRequest, len(list.Search))
	//for i, a := range list.Search {
	//	ar := s.dTOToRPC(&a)
	//	arrMovie[i] = ar
	//}
	fmt.Println("test")
	return nil
}

func (s *server) GetMovie(ctx context.Context, request *grpc2.SingleRequest) (*grpc2.Movie, error) {
	id := request.ImdbID
	movie := s.internalService.GetByImdbID(id)
	res := s.movieToRPC(&movie)
	return res, nil
}

func (s *server) movieToRPC(mv *entity.Movie) *grpc2.Movie {
	res := &grpc2.Movie{
		Title:  mv.Title,
		Year:   mv.Year,
		ImdbID: mv.ImdbID,
		Type:   mv.Type,
		Poster: mv.Poster,
	}
	return res
}

func (s *server) dTOToRPC(mv *dto.MovieDTO) *grpc2.Movie {
	res := &grpc2.Movie{
		Title:  mv.Title,
		Year:   mv.Year,
		ImdbID: mv.ImdbID,
		Type:   mv.Type,
		Poster: mv.Poster,
	}
	return res
}

func (s *server) rPCtoMovie(mv *grpc2.Movie) *entity.Movie {
	res := &entity.Movie{
		Title:  mv.Title,
		Year:   mv.Year,
		ImdbID: mv.ImdbID,
		Type:   mv.Type,
		Poster: mv.Poster,
	}
	return res
}

