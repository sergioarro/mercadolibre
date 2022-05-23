package usecase

import (
	"context"
	"math"
	"net/http"
	"strings"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"mercadolibre/config"
	"mercadolibre/internal/location"
	"mercadolibre/internal/models"
	"mercadolibre/pkg/httpErrors"
	"mercadolibre/pkg/logger"
)

type locationUC struct {
	cfg          *config.Config
	locationRepo location.Repository
	logger       logger.Logger
}

var satellitesPositions []models.Position
var satellitePosition models.Position

// Location UseCase constructor
func NewLocationUseCase(cfg *config.Config, locationRepo location.Repository, logger logger.Logger) location.UseCase {
	return &locationUC{cfg: cfg, locationRepo: locationRepo, logger: logger}
}

func (u *locationUC) PostTopSecretSplit(ctx context.Context, satelliteName string, satellite models.Satellite) (*models.Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "locationUC.PostTopSecretSplit")
	defer span.Finish()
	satellite.Name = satelliteName
	if satellite.Distance == 0 || len(satellite.Message) == 0 {
		err := errors.New("Request not valid")
		return nil, httpErrors.NewRestError(http.StatusPreconditionRequired, "Request not valid", errors.Wrap(err, "locationUC.PostTopSecretSplit"))
	}

	var satellitesOperating []models.Satellite = u.getAllSatellitesInService()
	satellitePosition = u.searchSatellite(satellite, satellitesOperating)
	u.logger.Debug("searchSatellite position : %s ", satellitePosition)
	if satellitePosition.X == 0 {
		err := errors.New("Satellite does not exist")
		return nil, httpErrors.NewRestError(http.StatusPreconditionRequired, "Satellite does not exist", errors.Wrap(err, "locationUC.PostTopSecretSplit"))
	}

	satellite.Position = satellitePosition
	u.logger.Debug("PostTopSecretSplit satellite : ", satellite)
	countSatellite, err := u.locationRepo.FindSatelliteByName(ctx, satelliteName)
	if err != nil {
		return nil, err
	}
	u.logger.Debug("FindSatelliteByName countSatellite : ", countSatellite)

	if countSatellite == 0 {
		err := u.locationRepo.Create(ctx, satellite)
		if err != nil {
			return nil, err
		}

	}

	var satellites []models.Satellite
	satellites, err = u.locationRepo.GetAllSatellites(ctx)

	if err != nil {
		return nil, err
	}

	var modelRequest models.Request
	modelRequest.RequestSatellites = satellites
	u.logger.Debug("PostTopSecretSplit RequestSatellites %+v : ", modelRequest.RequestSatellites)
	shipPositionAndMessage, err := u.CalculatePosition(modelRequest)
	if err != nil {
		return nil, err
	}

	return shipPositionAndMessage, nil
}

func (u *locationUC) addSatellite(satellite models.Satellite) error {

	return nil
}

func (u *locationUC) GetLocationBySatellites(ctx context.Context, satellites models.Request) (*models.Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "locationUC.GetLocationBySatellites")
	defer span.Finish()

	n, err := u.CalculatePosition(satellites)
	if err != nil {
		return nil, err
	}

	return n, nil
}

func (u *locationUC) CalculatePosition(satellites models.Request) (*models.Response, error) {

	if len(satellites.RequestSatellites) != 3 {
		err := errors.New("Not enough data")
		return nil, httpErrors.NewRestError(http.StatusNotFound, "Not enough data", errors.Wrap(err, "locationUC.GetLocationBySatellites.ValidToManySatellities"))
	}

	distances := make([]float64, 0)
	var messages [][]string
	for _, satellit := range satellites.RequestSatellites {
		distances = append(distances, satellit.Distance)
		messages = append(messages, satellit.Message)
	}
	u.logger.Debug("distances : %s ", distances)
	u.logger.Debug("messages : %s ", messages)

	var satellitesOperating []models.Satellite = u.getAllSatellitesInService()
	satellitesPositions = u.getPositionOfOperationalSatellites(satellites.RequestSatellites, satellitesOperating)

	x, y, err := GetLocation(distances...)
	if err != nil {
		return nil, err
	}

	message := GetMessage(messages...)
	position := models.Position{X: x, Y: y}
	n := &models.Response{
		Position: position,
		Message:  message,
	}

	return n, nil
}

//get all the satellites that are operating
func (u *locationUC) getAllSatellitesInService() (SatellitesInService []models.Satellite) {
	SatellitesInService = append(SatellitesInService,
		models.Satellite{Name: "Kenobi", Position: models.Position{X: -500, Y: -200}},
		models.Satellite{Name: "Skywalker", Position: models.Position{X: 100, Y: -100}},
		models.Satellite{Name: "Sato", Position: models.Position{X: 500, Y: 100}})
	return
}

func (u *locationUC) getPositionOfOperationalSatellites(shipToSatellites []models.Satellite, posotionStellites []models.Satellite) (coordinates []models.Position) {
	for _, satellite := range shipToSatellites {
		for _, satelliteOfTotal := range posotionStellites {
			if strings.ToUpper(satellite.Name) == strings.ToUpper(satelliteOfTotal.Name) {
				coordinates = append(coordinates, models.Position{X: satelliteOfTotal.Position.X, Y: satelliteOfTotal.Position.Y})
			}
		}
	}
	return coordinates
}

func (u *locationUC) searchSatellite(satellite models.Satellite, posotionStellites []models.Satellite) (coordinate models.Position) {
	for _, satelliteOfTotal := range posotionStellites {
		if strings.ToUpper(satellite.Name) == strings.ToUpper(satelliteOfTotal.Name) {
			coordinate = models.Position{X: satelliteOfTotal.Position.X, Y: satelliteOfTotal.Position.Y}
		}
	}
	return coordinate
}

// Función que devuelve la localizacion del emisor en sus coordenadas
func GetLocation(distances ...float64) (x, y float64, err error) {
	// Variables a utilizar para uso del algoritmo
	d1 := distances[0]
	d2 := distances[1]
	d3 := distances[2]
	// Posiciones actuales de los satelites
	i1 := satellitesPositions[0].X
	i2 := satellitesPositions[1].X
	i3 := satellitesPositions[2].X
	j1 := satellitesPositions[0].Y
	j2 := satellitesPositions[1].Y
	j3 := satellitesPositions[2].Y

	// Se calcula el valor de las coordenadas x,y según algoritmo de Trilateración
	x = (((math.Pow(d1, 2)-math.Pow(d2, 2))+(math.Pow(i2, 2)-math.Pow(i1, 2))+(math.Pow(j2, 2)-math.Pow(j1, 2)))*(2*j3-2*j2) - ((math.Pow(d2, 2)-math.Pow(d3, 2))+(math.Pow(i3, 2)-math.Pow(i2, 2))+(math.Pow(j3, 2)-math.Pow(j2, 2)))*(2*j2-2*j1)) / ((2*i2-2*i3)*(2*j2-2*j1) - (2*i1-2*i2)*(2*j3-2*j2))

	y = ((math.Pow(d1, 2) - math.Pow(d2, 2)) + (math.Pow(i2, 2) - math.Pow(i1, 2)) + (math.Pow(j2, 2) - math.Pow(j1, 2)) + x*(2*i1-2*i2)) / (2*j2 - 2*j1)

	return x, y, nil
}

func GetMessage(messages ...[]string) (mssg string) {
	var msg []string
	for i := 0; i < len(messages[0]); i++ {
		for j := 0; j < len(messages); j++ {
			// control longitud mensajes
			if len(messages[0]) != len(messages[j]) {
				return ""
			}
			// control de mensajes vacios y repetidos
			if messages[j][i] != "" && (len(msg) == 0 || msg[len(msg)-1] != messages[j][i]) {
				msg = append(msg, messages[j][i])
			}
		}
	}
	// control mensaje vacio
	if len(msg) > 0 {
		return strings.Join(msg, " ")
	} else {
		return ""
	}
}
