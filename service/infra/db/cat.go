package db

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/keitaro1020/lambda-golang-slf-example/service/domain"
	"github.com/keitaro1020/lambda-golang-slf-example/service/infra/db/models"
)

type catRepository struct {
	config *Config
}

func NewCatRepository(config *Config) domain.CatRepository {
	return &catRepository{
		config: config,
	}
}

func (re *catRepository) Get(ctx context.Context, id domain.CatID) (*domain.Cat, error) {
	db, err := connectDB(re.config)
	if err != nil {
		return nil, err
	}

	cat, err := models.Cats(models.CatWhere.ID.EQ(string(id))).One(ctx, db)
	if err != nil {
		return nil, err
	}
	return re.toDomain(cat), nil
}

func (re *catRepository) GetAll(ctx context.Context) (domain.Cats, error) {
	db, err := connectDB(re.config)
	if err != nil {
		return nil, err
	}

	cats, err := models.Cats().All(ctx, db)
	if err != nil {
		return nil, err
	}

	dcats := make(domain.Cats, len(cats))
	for i := range cats {
		dcats[i] = re.toDomain(cats[i])
	}
	return dcats, nil
}

func (re *catRepository) Create(ctx context.Context, cat *domain.Cat) (*domain.Cat, error) {
	db, err := connectDB(re.config)
	if err != nil {
		return nil, err
	}

	mcat := &models.Cat{
		ID:     string(cat.ID),
		URL:    cat.URL,
		Width:  int(cat.Width),
		Height: int(cat.Height),
	}
	if err := mcat.Upsert(ctx, db, boil.Infer(), boil.Infer()); err != nil {
		return nil, err
	}

	return re.toDomain(mcat), nil
}

func (re *catRepository) toDomain(cat *models.Cat) *domain.Cat {
	return &domain.Cat{
		ID:     domain.CatID(cat.ID),
		URL:    cat.URL,
		Width:  int64(cat.Width),
		Height: int64(cat.Height),
	}
}