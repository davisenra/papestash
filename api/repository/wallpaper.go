package repository

import (
	"database/sql"
	"errors"
	"time"
)

type Wallpaper struct {
	Id                int       `db:"id" json:"id"`
	Name              string    `db:"name" json:"name"`
	Path              string    `db:"path" json:"path"`
	ThumbnailPath     string    `db:"thumbnail_path" json:"thumbnailPath"`
	MostFrequentColor string    `db:"most_frequent_color" json:"mostFrequentColor"`
	Height            int       `db:"height" json:"height"`
	Width             int       `db:"width" json:"width"`
	AspectRatio       string    `db:"aspect_ratio" json:"aspectRatio"`
	SizeInBytes       int       `db:"size_in_bytes" json:"sizeInBytes"`
	CreatedAt         time.Time `db:"created_at" json:"createdAt"`
}

type WallpaperRepository struct {
	db *sql.DB
}

func NewWallpaperRepository(db *sql.DB) *WallpaperRepository {
	return &WallpaperRepository{db: db}
}

func (r *WallpaperRepository) GetAll() ([]Wallpaper, error) {
	rows, err := r.db.Query("SELECT id, name, path, thumbnail_path, most_frequent_color, height, width, aspect_ratio, size_in_bytes, created_at FROM wallpapers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallpapers []Wallpaper
	for rows.Next() {
		var w Wallpaper
		if err := rows.Scan(&w.Id, &w.Name, &w.Path, &w.ThumbnailPath, &w.MostFrequentColor, &w.Height, &w.Width, &w.AspectRatio, &w.SizeInBytes, &w.CreatedAt); err != nil {
			return nil, err
		}
		wallpapers = append(wallpapers, w)
	}

	if len(wallpapers) == 0 {
		return []Wallpaper{}, nil
	}

	return wallpapers, rows.Err()
}

func (r *WallpaperRepository) GetById(id int) (*Wallpaper, error) {
	var w Wallpaper
	err := r.db.QueryRow(`
		SELECT id, name, path, thumbnail_path, most_frequent_color, height, width, aspect_ratio, size_in_bytes, created_at
		FROM wallpapers WHERE id = ?`, id).
		Scan(&w.Id, &w.Name, &w.Path, &w.ThumbnailPath, &w.MostFrequentColor, &w.Height, &w.Width, &w.AspectRatio, &w.SizeInBytes, &w.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &w, nil
}

func (r *WallpaperRepository) Create(w Wallpaper) (int, error) {
	result, err := r.db.Exec(`
		INSERT INTO wallpapers (name, path, thumbnail_path, most_frequent_color, height, width, aspect_ratio, size_in_bytes, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		w.Name, w.Path, w.ThumbnailPath, w.MostFrequentColor, w.Height, w.Width, w.AspectRatio, w.SizeInBytes, w.CreatedAt)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	w.Id = int(id)

	return w.Id, nil
}

func (r *WallpaperRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM wallpapers WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}
