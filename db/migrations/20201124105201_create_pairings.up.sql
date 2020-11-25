CREATE TABLE IF NOT EXISTS public.pairings (
  event_id int NOT NULL REFERENCES events (id),
  rnd int NOT NULL,
  tble int NOT NULL,
  first_player int NOT NULL REFERENCES players (id),
  second_player int NOT NULL REFERENCES players (id),

  first_player_wins int NOT NULL DEFAULT 0,
  second_player_wins int NOT NULL DEFAULT 0,
  draws int NOT NULL DEFAULT 0,

  locked boolean NOT NULL DEFAULT FALSE,
  UNIQUE (event_id, rnd, tble)
);
