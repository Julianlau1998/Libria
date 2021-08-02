package topics

import (
	"fmt"
	"libria/models"
	"math/rand"

	uuid "github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	topicRepo Repository
}

func NewService(topicRepository Repository) Service {
	return Service{topicRepo: topicRepository}
}

func (s *Service) GetAll() ([]models.Topic, error) {
	return s.topicRepo.GetAll()
}

func (s *Service) GetById(id string) (models.Topic, error) {
	return s.topicRepo.GetById(id)
}

func (s *Service) GetRandom() (models.Topic, error) {
	topics, err := s.GetAll()
	if err != nil {
		log.Warnf("topicsService GetAll(), could not load toptcs: %s", err)
	}
	randomIndex := rand.Intn(len(topics))
	randomTopic := topics[randomIndex]
	return randomTopic, err
}

func (s *Service) Post(topic *models.Topic) (*models.Topic, error) {
	id, err := uuid.NewV4()
	if err != nil {
		fmt.Print(err)
		return topic, err
	}
	topic.ID = id.String()
	answer, err := s.topicRepo.Post(topic)
	if err != nil {
		fmt.Print(err)
	}
	return answer, err
}

func (s *Service) Update(id string, topic *models.Topic) (models.Topic, error) {
	topic.ID = id
	return s.topicRepo.Update(topic)
}

func (s *Service) Delete(id string) (models.Topic, error) {
	var topic models.Topic
	topic.ID = id
	return s.topicRepo.Delete(topic)
}
