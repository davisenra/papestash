CREATE TABLE wallpapers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    path TEXT NOT NULL,
    thumbnail_path TEXT,
    height INTEGER NOT NULL,
    width INTEGER NOT NULL,
    aspect_ratio VARCHAR NOT NULL,
    size_in_bytes INTEGER NOT NULL,
    most_frequent_color VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
