package sensorUsecase

import (
	"github.com/KumKeeHyun/gin-clean-arch/app/domain/model"
	"github.com/KumKeeHyun/gin-clean-arch/app/domain/repository"
)

type sensorUsecase struct {
	sr repository.SensorRepository
}

func NewSensorUsecase(sr repository.SensorRepository) *sensorUsecase {
	return &sensorUsecase{
		sr: sr,
	}
}

func (su *sensorUsecase) GetAllSensors() ([]model.Sensor, error) {
	sensors, err := su.sr.GetAll()
	if err != nil {
		return nil, err
	}
	for i := range sensors {
		sensors[i].ValueList, err = su.sr.GetValuesByUUID(sensors[i].UUID)
		if err != nil {
			return nil, err
		}
	}
	return sensors, nil
}

func (su *sensorUsecase) GetRegister() ([]model.Sensor, error) {
	sensors, err := su.GetAllSensors()
	return sensors, err
}

func (su *sensorUsecase) RegisterSensor(s *model.Sensor) error {
	newSensor := model.NewSensor(s.Name)
	if err := su.sr.Create(&newSensor); err != nil {
		return err
	}
	for i, v := range s.ValueList {
		v.SensorUUID = newSensor.UUID
		v.Index = i
		if err := su.sr.CreateValue(&v); err != nil {
			return err
		}
	}
	return nil
}
