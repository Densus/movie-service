package internal_service

import (
	"fmt"
	"github.com/densus/movie_service/entity"
	"github.com/densus/movie_service/repository"
)

type InternalService interface {
	GetByImdbID (imdbID string) entity.Movie
}

type internalService struct {
	movieRepository repository.MovieRepository
}

func NewInternalService (movieRepo repository.MovieRepository) InternalService {
	return &internalService{movieRepository: movieRepo}
}

func (i *internalService) GetByImdbID(imdbID string) entity.Movie {
	res := i.movieRepository.GetByImdbID(imdbID)
	fmt.Println("res: ", res)
	return res
}