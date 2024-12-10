CREATE TABLE wallpapers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    path TEXT NOT NULL,
    thumbnail_path TEXT,
    height INTEGER NOT NULL,
    width INTEGER NOT NULL,
    size_in_bytes INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);