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

func (repo *Repository) AddSegment(segmentName string) (int64, error) {
	sb := squirrel.StatementBuilder.
		Insert("Segment").
		Columns("slug").
		PlaceholderFormat(squirrel.Dollar).
		RunWith(repo.db)
	sb = sb.Values(segmentName)
	result, err := sb.Exec()
	if err != nil {
		return 0, errors.Wrap(err, "sb.exec")
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "lastInsertID")
	}
	return lastInsertID, nil
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

func (repo *Repository) GetUserSegments(userID uint) ([]string, error) {
	rowCount, err := repo.getUserSegmentsCount(userID)
	if err != nil {
		return nil, errors.Wrap(err, "getUserSegmentsCount")
	}

	segments := make([]string, 0, rowCount)
	result, err := repo.db.Query("SELECT Segment.slug FROM Segment JOIN UsersSegments ON UsersSegments.segmentId = Segment.id WHERE UsersSegments.userId = $1", userID)
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

func (repo *Repository) AddUserToSegments(userID uint, segments []string) error {
	sb := squirrel.StatementBuilder.
		Insert("UsersSegments").
		Columns("userId", "segmentId").
		PlaceholderFormat(squirrel.Dollar).
		RunWith(repo.db)

	for _, segment := range segments {
		segmentID, err := repo.getSegmentIDBySlug(segment)
		if err != nil {
			return errors.Wrap(err, "getSegmentIDBySlug")
		}
		sb = sb.Values(
			userID,
			segmentID,
		)
	}
	if _, err := sb.Exec(); err != nil {
		return errors.Wrap(err, "sb.exec")
	}
	return nil
}

func (repo *Repository) DeleteUserFromSegments(userID uint, segments []string) error {
	for _, segment := range segments {
		segmentID, err := repo.getSegmentIDBySlug(segment)
		if err != nil {
			return errors.Wrap(err, "getSegmentIDBySlug")
		}

		sb := squirrel.StatementBuilder.
			Delete("UsersSegments").
			Where(squirrel.Eq{"userId": userID, "segmentId": segmentID}).
			PlaceholderFormat(squirrel.Dollar).
			RunWith(repo.db)

		if _, err := sb.Exec(); err != nil {
			return errors.Wrap(err, "sb.exec")
		}
	}
	return nil
}

func (repo *Repository) getUserSegmentsCount(userID uint) (int, error) {
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

func (repo *Repository) getSegmentIDBySlug(segmentSlug string) (uint, error) {
	var segmentID uint
	query := squirrel.Select("id").
		From("Segment").
		Where(squirrel.Eq{"slug": segmentSlug}).
		PlaceholderFormat(squirrel.Dollar).
		RunWith(repo.db)

	if err := query.QueryRow().Scan(&segmentID); err != nil {
		return 0, errors.Wrap(err, "scan")
	}
	return segmentID, nil
}
