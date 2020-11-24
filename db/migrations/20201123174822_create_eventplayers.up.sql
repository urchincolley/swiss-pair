CREATE TABLE IF NOT EXISTS public.eventplayers (
  player_id int NOT NULL REFERENCES players (id),
  event_id int NOT NULL REFERENCES events (id),
  UNIQUE (player_id, event_id)
);
