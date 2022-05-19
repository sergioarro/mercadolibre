package usecase

import (
	"context"
	"fmt"
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

var satellitesPositions []models.Position

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

	var satellitesOperating []models.Satellite = u.getAllSatellitesInService()
	satellitesPositions = getPositionOfOperationalSatellites(satellites.RequestSatellites, satellitesOperating)

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

func getPositionOfOperationalSatellites(shipToSatellites []models.Satellite, posotionStellites []models.Satellite) (coordinates []models.Position) {
	for _, satellite := range shipToSatellites {
		for _, satelliteOfTotal := range posotionStellites {
			if strings.ToUpper(satellite.Name) == strings.ToUpper(satelliteOfTotal.Name) {
				coordinates = append(coordinates, models.Position{X: satelliteOfTotal.Position.X, Y: satelliteOfTotal.Position.Y})
			}
		}
	}
	return coordinates
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
