CREATE TABLE IF NOT EXISTS public.players (
  id serial NOT NULL PRIMARY KEY,
  first_name text NOT NULL,
  last_name text NOT NULL,
  email text NOT NULL UNIQUE
);
