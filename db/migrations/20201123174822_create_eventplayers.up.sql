CREATE TABLE IF NOT EXISTS public.eventplayers (
  player_id int NOT NULL REFERENCES players (id),
  event_id int NOT NULL REFERENCES events (id),
  match_points int NOT NULL DEFAULT 0,
  game_points int NOT NULL DEFAULT 0,
  matches_played int NOT NULL DEFAULT 0,
  games_played int NOT NULL DEFAULT 0,
  match_win_perc float DEFAULT 0.0,
  opp_match_win_perc float DEFAULT 0.0,
  game_win_perc float DEFAULT 0.0,
  opp_game_win_perc float DEFAULT 0.0,
  prev_opponents int[] NOT NULL DEFAULT '{}',
  UNIQUE (player_id, event_id)
);
