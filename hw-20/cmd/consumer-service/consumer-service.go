package main

import "github.com/rs/zerolog/log"

type OrangeStorage interface {
	CreateOrange(o Orange) error
	GetAllOranges() []Orange
}

type OrangeService struct {
	Storage   OrangeStorage
	Analytics OrangesAnalytics
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

func NewOrangesAnalytics() *OrangesAnalytics {
	return &OrangesAnalytics{}
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

	switch {
	case oe.Size >= 0 && oe.Size <= 100:
		s.Analytics.Small++
	case oe.Size > 100 && oe.Size <= 200:
		s.Analytics.Medium++
	case oe.Size > 200 && oe.Size <= 300:
		s.Analytics.Big++
	}
	return nil
}

func (s *OrangeService) GetAnalytics() map[string]int {
	analytics := make(map[string]int, 3)
	analytics["small"] = s.Analytics.Small
	analytics["medium"] = s.Analytics.Medium
	analytics["big"] = s.Analytics.Big
	return analytics
}
