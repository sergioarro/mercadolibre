package usecase

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"mercadolibre/config"
	"mercadolibre/internal/location"
	"mercadolibre/internal/models"
	"mercadolibre/pkg/httpErrors"
	"mercadolibre/pkg/logger"
	"mercadolibre/pkg/utils"
)

const (
	basePrefix    = "api-location:"
	cacheDuration = 3600
)

type locationUC struct {
	cfg          *config.Config
	locationRepo location.Repository
	logger       logger.Logger
}

var NoSolutionLocation = fmt.Errorf("No solution for localization.")
var CoordinatesError = fmt.Errorf("The number of coordinates to analyze is incorrect. It must be one, two, or three maximum coordinates.")
var NoSolutionMessages = fmt.Errorf("The message cannot be decrypted.")

var SatellCoordinates []utils.Point

// Location UseCase constructor
func NewLocationUseCase(cfg *config.Config, locationRepo location.Repository, logger logger.Logger) location.UseCase {
	return &locationUC{cfg: cfg, locationRepo: locationRepo, logger: logger}
}

func (u *locationUC) GetLocationBySatellites(ctx context.Context, satellites models.Request) (*models.Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "locationUC.GetLocationBySatellites")
	defer span.Finish()

	if len(satellites.RequestSatellites) > 3 {
		err := errors.New("Too many satellities")
		return nil, httpErrors.NewRestError(http.StatusPreconditionRequired, "Too many satellities", errors.Wrap(err, "locationUC.GetLocationBySatellites.ValidToManySatellities"))
	}

	distances := make([]float64, 0)
	var messages [][]string
	for _, satellit := range satellites.RequestSatellites {
		distances = append(distances, satellit.Distance)
		messages = append(messages, satellit.Message)
	}
	u.logger.Debug("distances : %s ", distances)
	u.logger.Debug("messages : %s ", messages)

	//var satellitesOperating []models.Satellite = u.getAllSatellitesInService()
	//SatellCoordinates = getPositionOfOperationalSatellites(satellites.RequestSatellites, satellitesOperating)

	x, y, err := GetLocation(distances...)
	if err != nil {
		return nil, err
	}

	position := models.Position{X: x, Y: y}
	n := &models.Response{
		Position: position,
		Message:  "este es un mensaje secreto",
	}

	//message, errmsg := GetMessage(messages...)
	/*n, err := u.locationRepo.GetLocationBySatellites(ctx, satellites)
	if err != nil {
		return nil, err
	}*/

	return n, err
}

func (u *locationUC) getKeyWithPrefix(newsID string) string {
	return fmt.Sprintf("%s: %s", basePrefix, newsID)
}

//get all the satellites that are operating
func (u *locationUC) getAllSatellitesInService() (SatellitesInService []models.Satellite) {
	SatellitesInService = append(SatellitesInService,
		models.Satellite{Name: "Kenobi", Position: models.Position{X: -500, Y: -200}},
		models.Satellite{Name: "Skywalker", Position: models.Position{X: 100, Y: -100}},
		models.Satellite{Name: "Sato", Position: models.Position{X: 500, Y: 100}})
	return
}

func getPositionOfOperationalSatellites(shipToSatellites []models.Satellite, posotionStellites []models.Satellite) (coordinates []utils.Point) {
	for _, satellite := range shipToSatellites {
		for _, satelliteOfTotal := range posotionStellites {
			if strings.ToUpper(satellite.Name) == strings.ToUpper(satelliteOfTotal.Name) {
				coordinates = append(coordinates, utils.Point{X: satelliteOfTotal.Position.X, Y: satelliteOfTotal.Position.Y})
			}
		}
	}
	return coordinates
}

func GetLocation(distances ...float64) (x, y float64, err error) {

	countCoordinates := len(SatellCoordinates)
	p1 := utils.Point{}
	p2 := utils.Point{}
	p3 := utils.Point{}

	switch countCoordinates {
	case 1:
		p1 = SatellCoordinates[0]
		p1.R = distances[0]
	case 2:
		p1 = SatellCoordinates[0]
		p1.R = distances[0]
		p2 = SatellCoordinates[1]
		p2.R = distances[1]
	case 3:
		p1 = SatellCoordinates[0]
		p1.R = distances[0]
		p2 = SatellCoordinates[1]
		p2.R = distances[1]
		p3 = SatellCoordinates[2]
		p3.R = distances[2]
	default:
		return 9999999999, 9999999999, httpErrors.NewRestError(http.StatusBadRequest, CoordinatesError.Error(), errors.Wrap(err, "locationUC.GetLocationBySatellites.GetLocation"))
	}

	ex := utils.Divide(utils.Subtract(p2, p1), utils.Normalize(utils.Subtract(p2, p1)))

	i := utils.Dot(ex, utils.Subtract(p3, p1))
	a := utils.Subtract(utils.Subtract(p3, p1), utils.Multiply(ex, i))
	ey := utils.Divide(a, utils.Normalize(a))
	d := utils.Normalize(utils.Subtract(p2, p1))
	j := utils.Dot(ey, utils.Subtract(p3, p1))

	//calculate X and Y for the coordinate
	x = ((p1.R * p1.R) - (p2.R * p2.R) + (d * d)) / (2 * d)
	y = (utils.Square(p1.R)-utils.Square(p3.R)+utils.Square(i)+utils.Square(j))/(2*j) - (i/j)*x

	if x == 0 && y == 0 {
		return 9999999999, 9999999999, NoSolutionLocation
	}
	location := utils.Add(p1, utils.Add(utils.Multiply(ex, x), utils.Multiply(ey, y)))
	location = utils.RoundUp(location, 3)

	return location.X, location.Y, nil
}

func GetMessage(messages ...[]string) (msg string) {

	return "msg"
}
