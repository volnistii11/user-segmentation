package database

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) AddSegment(segmentName string) error {
	sb := squirrel.StatementBuilder.
		Insert("Segment").
		Columns("slug").
		PlaceholderFormat(squirrel.Dollar).
		RunWith(repo.db)
	sb = sb.Values(segmentName)
	if _, err := sb.Exec(); err != nil {
		return errors.Wrap(err, "sb.exec")
	}
	return nil
}

func (repo *Repository) DeleteSegment(segmentName string) error {
	sb := squirrel.StatementBuilder.
		Delete("Segment").
		Where(squirrel.Eq{"slug": segmentName}).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(repo.db)
	if _, err := sb.Exec(); err != nil {
		return errors.Wrap(err, "sb.exec")
	}
	return nil
}

func (repo *Repository) GetUserSegments(userID int) ([]string, error) {
	rowCount, err := repo.getUserSegmentsCount(userID)
	if err != nil {
		return nil, errors.Wrap(err, "getUserSegmentsCount")
	}

	segments := make([]string, 0, rowCount)
	result, err := repo.db.Query("SELECT Segment.slug FROM Segment JOIN UsersSegments ON UsersSegments.segmentId = Segment.id WHERE UsersSegments.userId = ?", userID)
	if err != nil {
		return nil, errors.Wrap(err, "select user segments")
	}
	for result.Next() {
		var segment string
		if err = result.Scan(&segment); err != nil {
			return nil, errors.Wrap(err, "result.Scan")
		}
		segments = append(segments, segment)
	}

	return segments, nil
}

func (repo *Repository) getUserSegmentsCount(userID int) (int, error) {
	var rowCount int
	query := squirrel.Select("COUNT(*)").
		From("UsersSegments").
		Where(squirrel.Eq{"userId": userID}).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(repo.db)

	if err := query.QueryRow().Scan(&rowCount); err != nil {
		return 0, errors.Wrap(err, "scan")
	}
	return rowCount, nil
}
