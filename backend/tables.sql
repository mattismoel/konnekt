CREATE TABLE member (
  id INTEGER PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  password_hash TEXT NOT NULL,
  profile_picture_url TEXT,
  active BOOLEAN NOT NULL DEFAULT "FALSE"
);

CREATE TABLE session (
  id TEXT PRIMARY KEY,
  member_id INTEGER NOT NULL,
  expires_at TIMESTAMP NOT NULL
);

CREATE TABLE team (
  id INTEGER PRIMARY KEY,
  name TEXT UNIQUE NOT NULL,
  display_name TEXT NOT NULL,
  description TEXT
);

CREATE TABLE members_teams (
  member_id INTEGER NOT NULL,
  team_id INTEGER NOT NULL,
  PRIMARY KEY (member_id, team_id)
);

CREATE TABLE permission (
  id INTEGER PRIMARY KEY,
  name TEXT UNIQUE NOT NULL,
  display_name TEXT NOT NULL,
  description TEXT
);

CREATE TABLE teams_permissions (
  team_id INTEGER NOT NULL,
  permission_id INTEGER NOT NULL,
  PRIMARY KEY (team_id, permission_id)
);

CREATE TABLE event (
  id INTEGER PRIMARY KEY,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  ticket_url TEXT NOT NULL,
  image_url TEXT NOT NULL,
  venue_id INTEGER NOT NULL,

  FOREIGN KEY (venue_id) REFERENCES venue (id)
);

CREATE TABLE concert (
  id INTEGER PRIMARY KEY,
  from_date TIMESTAMP NOT NULL,
  to_date TIMESTAMP NOT NULL,
  event_id INTEGER NOT NULL,
  artist_id INTEGER NOT NULL,

  FOREIGN KEY (event_id) REFERENCES event (id),
  FOREIGN KEY (artist_id) REFERENCES artist (id)
);

CREATE TABLE venue (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  country_code TEXT NOT NULL,
  city TEXT NOT NULL
);

CREATE TABLE artist (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  image_url TEXT NOT NULL,
  preview_url TEXT NOT NULL,
  description TEXT NOT NULL
);

CREATE TABLE social (
  id INTEGER PRIMARY KEY,
  url TEXT NOT NULL
);

CREATE TABLE genre (
  id INTEGER PRIMARY KEY,
  name TEXT UNIQUE NOT NULL
);

CREATE TABLE artists_socials (
  artist_id INTEGER NOT NULL,
  social_id INTEGER NOT NULL,
  PRIMARY KEY (artist_id, social_id)
);

CREATE TABLE artists_genres (
  artist_id INTEGER NOT NULL,
  genre_id INTEGER NOT NULL,

  PRIMARY KEY (artist_id, genre_id)
);
