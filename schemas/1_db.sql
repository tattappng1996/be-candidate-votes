-- users definition

-- Drop table

-- DROP TABLE users;

CREATE TABLE users (
    id serial PRIMARY KEY,
    username varchar(50) NOT NULL,
    password text NOT NULL,
    is_active boolean DEFAULT true,
    created_at timestamp(6) DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(6) DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6)
);

-- candidates definition

-- Drop table

-- DROP TABLE candidates;

CREATE TABLE candidates (
    id serial PRIMARY KEY,
    name varchar(30) NOT NULL,
	description text NOT NULL,
	created_at timestamp(6) DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(6) DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6)
);

-- votes definition

-- Drop table

-- DROP TABLE votes;

CREATE TABLE votes (
    id serial PRIMARY KEY,
	user_id integer NOT NULL,
    candidate_id integer NOT NULL,
    created_at timestamp(6) DEFAULT CURRENT_TIMESTAMP
);

-- vote_statuses definition

-- Drop table

-- DROP TABLE vote_statuses;

CREATE TABLE vote_statuses (
	is_active boolean DEFAULT true,
    created_at timestamp(6) DEFAULT CURRENT_TIMESTAMP
);