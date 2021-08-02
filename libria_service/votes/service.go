package votes

import (
	"fmt"
	"libria/models"

	uuid "github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	voteRepo Repository
}

func NewService(voteRepository Repository) Service {
	return Service{voteRepo: voteRepository}
}

func (s *Service) GetAllByAnswer(answerID string) ([]models.Vote, error) {
	votes, err := s.voteRepo.GetAllByAnswer(answerID)
	return votes, err
}

func (s *Service) Post(vote *models.Vote) (*models.Vote, error) {
	id, err := uuid.NewV4()
	if err != nil {
		log.Warnf("VoteService.Post() Could not create new uuid: %s", err)
		return vote, err
	}
	vote.ID = id.String()
	fmt.Print(vote.AnswerID)
	vote, err = s.voteRepo.Post(vote)
	if err != nil {
		log.Warnf("VoteService.Post() Could not Post Vote: %s", err)
		return vote, err
	}
	return vote, err
}

func (s *Service) Update(id string, vote *models.Vote) (models.Vote, error) {
	vote.ID = id
	updatedVote, err := s.voteRepo.Update(vote)
	return updatedVote, err
}
