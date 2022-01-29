package handler

import (
	"context"
	grpc2 "github.com/densus/movie_service/delivery/grpc"
	"github.com/densus/movie_service/entity"
	"github.com/densus/movie_service/entity/dto"
	external_service "github.com/densus/movie_service/service/external-service"
	internal_service "github.com/densus/movie_service/service/internal-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"sync"
)

func NewMovieServerGrpc(gServer *grpc.Server, externalServ external_service.ExternalService, internalServ internal_service.InternalService) {
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
	//use wait group to allow process to be concurrent
	var wg sync.WaitGroup
	page := request.GetPagination()
	list := s.externalService.Search(request.GetSearchWord(), int(page))

	if list.Search == nil || len(list.Search) == 0 {
		return nil
	}

	arrMovie := make([]*grpc2.Movie, len(list.Search)) //make data type []*grpc2.Movie with a length of list.Search
	for i, a := range list.Search {
		//fmt.Println("a: ", a)
		wg.Add(i)
		a := a
		go func(int642 int64) {
			defer wg.Done()
			//fmt.Println("i", i)
			_dto := s.movieDtoToRPC(&a)
			arrMovie = append(arrMovie, _dto)
			resp := grpc2.SearchResult{Search: arrMovie}
			//fmt.Println("resp: ", resp)
			if err := searchServer.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
		}(int64(i))
	}

	wg.Wait() //
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

func (s *server) movieDtoToRPC(mv *dto.MovieDTO) *grpc2.Movie {
	res := &grpc2.Movie{
		Title:  mv.Title,
		Year:   mv.Year,
		ImdbID: mv.ImdbID,
		Type:   mv.Type,
		Poster: mv.Poster,
	}
	return res
}
