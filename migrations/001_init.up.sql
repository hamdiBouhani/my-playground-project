SET search_path TO public;
DROP EXTENSION IF EXISTS "uuid-ossp";

CREATE EXTENSION "uuid-ossp" SCHEMA public;


CREATE TABLE object_store (
    id                 bigserial                            NOT NULL,
    uuid               UUID      DEFAULT public.uuid_generate_v4() NOT NULL,

    name               character varying(50),
    description        character varying(5000),
    type               character varying(255),
    link               character varying(255),

    created_date       timestamp DEFAULT now(),
    changed_date       timestamp,
    deleted_date       timestamp,
    
    CONSTRAINT t_pk PRIMARY KEY (id)
);

