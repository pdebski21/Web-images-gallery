CREATE OR REPLACE TRIGGER image_added AFTER INSERT ON images
FOR EACH ROW  -- There's an error when ROW is replaced STATEMENT... why? I though AFTER INSERT should be executed on statement by statement basis.
EXECUTE FUNCTION tf_image_added();