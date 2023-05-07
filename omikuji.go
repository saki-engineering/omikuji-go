package omikuji

import (
	"errors"

	"github.com/saki-engineering/omikuji-go/internal"
)

type Omikuji interface {
	DrawStatus() (string, error)
	DrawPaper() (*Paper, error)
	Finish()
}

type kuji struct {
	ck internal.OmikujiItf
}

var _ Omikuji = &kuji{}

func New() Omikuji {
	return &kuji{
		ck: internal.OmikujiItfCreateBox(),
	}
}

func (k *kuji) DrawPaper() (*Paper, error) {
	paper := internal.NewPaper()
	defer internal.DeletePaper(paper)

	result := k.ck.DrawPaper(paper)
	if err := convertErrCode(result); err != nil {
		return nil, err
	}
	return convertPaper(paper), nil
}

func (k *kuji) DrawStatus() (string, error) {
	status := internal.New_strp()
	defer internal.Delete_strp(status)

	result := k.ck.DrawStatus(status)
	if err := convertErrCode(result); err != nil {
		return "", err
	}

	return internal.Strp_value(status), nil
}

func (k *kuji) Finish() {
	internal.DeleteOmikujiItf(k.ck)
	k.ck = nil
}

type Paper struct {
	Status      string
	Description string
}

func convertPaper(p internal.Paper) *Paper {
	return &Paper{
		Status:      p.GetStatus(),
		Description: p.GetDescription(),
	}
}

var (
	ErrUnknown  = errors.New("unknown err occurred")
	ErrBreakBox = errors.New("oops, Omikuji box is broken")
	ErrStuckked = errors.New("omikuji is stucked")
)

func convertErrCode(code internal.OmikujiError) error {
	switch code {
	case internal.Error_NONE:
		return nil
	case internal.Error_BREAK_BOX:
		return ErrBreakBox
	case internal.Error_STUCKKED:
		return ErrStuckked
	default:
		return ErrUnknown
	}
}
