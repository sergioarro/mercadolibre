DROP TABLE IF EXISTS satellite;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE TABLE satellite
(
    name              VARCHAR(30)                                   NOT NULL CHECK ( name <> '' ),
    message           TEXT[],
    distance          NUMERIC,
    position          JSON                   
);


