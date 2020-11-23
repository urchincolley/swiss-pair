package models

type Model interface {
	WithId(int)
	GetById(context.Context, *application.Application) error
}

func (m *Model) GetById(ctx context.Context, app *application.Application) error {
	return errors.New("not implemented")
}
