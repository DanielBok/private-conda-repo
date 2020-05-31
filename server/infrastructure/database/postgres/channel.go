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

func (p *Postgres) RemoveChannel(id int) error {
	if id <= 0 {
		return errors.New("invalid channel id")
	}

	var channel entity.Channel
	err := p.db.First(&channel, id).Error
	if err != nil {
		return err
	}

	errs := p.db.Delete(&channel).GetErrors()
	if len(errs) > 0 {
		return joinErrors(errs)
	}

	return nil
}
