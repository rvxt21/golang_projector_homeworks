package main

import "github.com/rs/zerolog/log"

type OrangeStorage interface {
	CreateOrange(o Orange) error
	GetAllOranges() []Orange
	GetAnalytics() map[string]int
}

type OrangeService struct {
	Storage OrangeStorage
}

type Orange struct {
	ID   int
	Size int
}

type OrangesAnalytics struct {
	Small  int
	Medium int
	Big    int
}

type OrangeEvent struct {
	OrangeID int
	Size     int
}

func (s *OrangeService) ConsumeOrangeEvent(oe OrangeEvent) error {
	const op = "consumer-service.consumer-service.ConsumeOrangeEvent"
	var orange Orange
	orange.ID, orange.Size = oe.OrangeID, oe.Size

	err := s.Storage.CreateOrange(orange)
	if err != nil {
		log.Info().Err(err).Msgf("%s: %s", op, err)
	}

	return nil
}

func (s *OrangeService) GetAnalytics() map[string]int {
	analytics := s.Storage.GetAnalytics()
	return analytics
}
