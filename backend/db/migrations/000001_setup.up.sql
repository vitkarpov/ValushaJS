CREATE TABLE comments
(
  id serial,
  parent_id serial,
  author_name character varying(100) NOT NULL,
  comment text NOT NULL,
  created_at time without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);