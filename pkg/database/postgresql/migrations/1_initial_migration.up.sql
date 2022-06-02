-- +migrate Up

CREATE TABLE public.person
(
    id           serial                      NOT NULL,
    first_name   varchar                     NOT NULL,
    last_name    varchar                     NOT NULL,
    CONSTRAINT offer_pkey PRIMARY KEY (id)
);