package postgres

import (
	"errors"
	"strings"

	"private-conda-repo/domain/entity"
)

func (p *Postgres) CreateChannel(channel, password, email string) (*entity.Channel, error) {
	chn := entity.NewChannel(channel, password, email)

	if errs := p.db.Create(chn).GetErrors(); len(errs) > 0 {
		return nil, joinErrors(errs)
	}

	return chn, nil
}

func (p *Postgres) GetAllChannels() ([]*entity.Channel, error) {
	var channels []*entity.Channel
	if errs := p.db.Find(&channels).GetErrors(); len(errs) > 0 {
		return nil, joinErrors(errs)
	}
	return channels, nil
}

func (p *Postgres) GetChannel(channel string) (*entity.Channel, error) {
	var chn entity.Channel
	if errs := p.db.
		Where("channel = ?", strings.ToLower(channel)).
		First(&chn).
		GetErrors(); len(errs) > 0 {
		return nil, joinErrors(errs)
	}

	return &chn, nil
}

func (p *Postgres) RemoveChannel(channel, password string) error {
	var chn entity.Channel
	if errs := p.db.
		Where("channel = ?", strings.ToLower(channel)).
		First(&chn).
		GetErrors(); len(errs) > 0 {
		return joinErrors(errs)
	}

	if !chn.HasValidPassword(password) {
		return errors.New("incorrect credentials supplied to delete chn")
	}

	if errs := p.db.Delete(&chn).GetErrors(); len(errs) > 0 {
		return joinErrors(errs)
	}

	return nil
}
