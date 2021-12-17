
CREATE TABLE grades (
    id integer PRIMARY KEY NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    image_id integer,
    user_id integer,
    grade integer CHECK (grade >= 0 AND grade <= 10)  -- See? It's a contraint, yo.
    -- some primary and foreign key constraints.
);
