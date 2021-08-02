package answers

import (
	"libria/models"

	uuid "github.com/nu7hatch/gouuid"
)

type Service struct {
	answerRepo Repository
}

func NewService(answerRepository Repository) Service {
	return Service{answerRepo: answerRepository}
}

func (s *Service) GetAll() ([]models.Answer, error) {
	return s.answerRepo.GetAll()
}

func (s *Service) GetAllByTopic(topicId string) ([]models.Answer, error) {
	return s.answerRepo.GetAllByTopic(topicId)
}

func (s *Service) GetById(id string) (models.Answer, error) {
	return s.answerRepo.GetById(id)
}

func (s *Service) Post(answer *models.Answer) (*models.Answer, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return answer, err
	}
	answer.ID = id.String()
	return s.answerRepo.Post(answer)
}

func (s *Service) Update(id string, answer *models.Answer) (models.Answer, error) {
	answer.ID = id
	return s.answerRepo.Update(answer)
}

func (s *Service) Delete(id string) (models.Answer, error) {
	var answer models.Answer
	answer.ID = id
	return s.answerRepo.Delete(answer)
}
