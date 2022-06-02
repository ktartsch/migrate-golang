-- +migrate Up

ALTER TABLE public.person ADD COLUMN email varchar NULL;
