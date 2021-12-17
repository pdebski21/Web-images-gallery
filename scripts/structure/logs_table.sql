-- Table: public.logs

-- DROP TABLE IF EXISTS public.logs;

CREATE TABLE IF NOT EXISTS public.logs
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    type character(1) COLLATE pg_catalog."default" NOT NULL,
    stamp timestamp with time zone NOT NULL,
    record_id integer NOT NULL,
    CONSTRAINT logs_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.logs
    OWNER to postgres;