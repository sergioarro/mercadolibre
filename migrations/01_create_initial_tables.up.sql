DROP TABLE IF EXISTS satellite;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE TABLE satellite
(
    satellite_id      UUID PRIMARY KEY                              DEFAULT uuid_generate_v4(),
    name              VARCHAR(30)                                   NOT NULL CHECK ( name <> '' ),
    message           TEXT[],
    distance          NUMERIC,
    position          JSON,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP   NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP                     
);
