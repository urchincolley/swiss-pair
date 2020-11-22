CREATE TABLE IF NOT EXISTS public.events (
  id serial NOT NULL PRIMARY KEY,
  name text NOT NULL UNIQUE
);
