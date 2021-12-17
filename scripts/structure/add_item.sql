-- A function used to log insertions.
-- Lots of the fluff was added by pgAdmin.

-- FUNCTION: public.add_item(character, integer)

-- DROP FUNCTION IF EXISTS public.add_item(character, integer);

CREATE OR REPLACE FUNCTION public.add_item(
	tag character,
	record_id integer)
    RETURNS void
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
BEGIN
    INSERT INTO logs (type, stamp, record_id) SELECT "tag", now(), "record_id";
END;
$BODY$;

ALTER FUNCTION public.add_item(character, integer)
    OWNER TO postgres;
