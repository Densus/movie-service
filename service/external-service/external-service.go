package external_service

import (
	"encoding/json"
	"fmt"
	"github.com/densus/movie_service/entity"
	"github.com/densus/movie_service/entity/dto"
	"github.com/densus/movie_service/repository"
	"github.com/mashingan/smapping"
	"log"
	"net/http"
	"os"
	"strconv"
)

type ExternalService interface {
	Search(search string, page int) dto.MovieResponse
}

type externalService struct {
	movieRepository repository.MovieRepository
}

func NewExternalService(movieRepo repository.MovieRepository) ExternalService {
	return &externalService{
		movieRepository: movieRepo,
	}
}

func (e *externalService) Search(search string, page int) dto.MovieResponse {
	var client = &http.Client{}
	var data dto.MovieResponse

	var dataFromDb []entity.Movie
	dataFromDb = e.movieRepository.GetByTitle(search, page, 5)
	fmt.Println(dataFromDb)
	if len(dataFromDb) == 0 { //check if dataFromDb == []
		url := os.Getenv("URL")
		apiKey := "apikey=" + os.Getenv("API_KEY")
		searchWord := "&s=" + search
		_page := strconv.Itoa(page)
		pagination := "&page=" + _page
		//pagination := "&page=" + string((page))

		req, err := http.NewRequest("GET", url+apiKey+searchWord+pagination, nil)
		if err != nil {
			panic(err)
		}
		//log.Println("Test: ", req)

		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		log.Println(res)
		defer res.Body.Close()

		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			panic(err)
		}

		for _, each := range data.Search {
			a := mapDTOtoEntity(each)
			e.movieRepository.Log(a)
		}
		return data
	} else {
		for _, each := range dataFromDb {
			a := mapEntityToDTO(each)
			data.Search = append(data.Search, a)
		}
		return data
	}
}

func mapDTOtoEntity(data dto.MovieDTO) entity.Movie {
	mapped := smapping.MapFields(&data)
	movieToCreate := entity.Movie{}
	err := smapping.FillStruct(&movieToCreate, mapped)
	if err != nil {
		panic(err)
	}

	return movieToCreate
}
func mapEntityToDTO(data entity.Movie) dto.MovieDTO {
	mapped := smapping.MapFields(&data)
	movieToView := dto.MovieDTO{}
	err := smapping.FillStruct(&movieToView, mapped)
	if err != nil {
		panic(err)
	}

	return movieToView
}
