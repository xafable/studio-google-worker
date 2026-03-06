package usecases

import (
	"github.com/xafable/studio-google-worker/internal/birthdays/entities"
	"github.com/xafable/studio-google-worker/internal/birthdays/interfaces"
)

type GetBirthdays struct {
	Repo   interfaces.BirthdayRepository
	Source interfaces.BirthdaySource
}

func NewBirthdaysServ(repo interfaces.BirthdayRepository, source interfaces.BirthdaySource) *GetBirthdays {
	return &GetBirthdays{
		Repo:   repo,
		Source: source,
	}
}

func (g GetBirthdays) Get() ([]*entities.Birthday, error) {
	b, err := g.Source.GetBirthdays()
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (g GetBirthdays) Save(b []*entities.Birthday) {
	g.Repo.SaveBirthdays(b)
}
