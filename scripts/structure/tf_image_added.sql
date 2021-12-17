CREATE OR REPLACE FUNCTION public.tf_image_added()
    RETURNS trigger
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE NOT LEAKPROOF
AS $BODY$
BEGIN
    PERFORM add_item('I', NEW.id);
    RETURN NEW;
END;
$BODY$;

ALTER FUNCTION public.tf_image_added()
    OWNER TO postgres;

GRANT EXECUTE ON FUNCTION public.tf_image_added() TO postgres WITH GRANT OPTION;

GRANT EXECUTE ON FUNCTION public.tf_image_added() TO PUBLIC;
