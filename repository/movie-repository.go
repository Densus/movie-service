package repository

import (
	"github.com/densus/movie_service/entity"
	"gorm.io/gorm"
)

type MovieRepository interface {
	Log(movie entity.Movie) entity.Movie
	GetByTitle(title string) []entity.Movie
	GetByImdbID(imdbID string) entity.Movie
}

type movieRepository struct {
	dbMovieConnection *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{
		dbMovieConnection: db,
	}
}

func (m *movieRepository) Log(movie entity.Movie) entity.Movie {
	m.dbMovieConnection.Save(&movie)
	return movie
}

func (m *movieRepository) GetByTitle(title string) []entity.Movie {
	var movie []entity.Movie
	_title := "%"+title+"%"
	m.dbMovieConnection.Where("title LIKE ?", _title).Find(&movie)
	return movie
}

func (m *movieRepository) GetByImdbID(imdbID string) entity.Movie {
	var movie entity.Movie
	m.dbMovieConnection.Find(&movie,"imdb_id = ?", imdbID)
	return movie
}
