-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS jobs (
	id varchar(10) NOT NULL,
    company_id varchar(18) NOT NULL,
	title varchar(64) NOT NULL,
    location varchar(64) NOT NULL,
    company varchar(64) NOT NULL,
    listing_date timestamptz NOT NULL,
    salary varchar(64) NOT NULL,
    role varchar(64) NOT NULL,
    job_description text NOT NULL,
    work_type varchar(48), 
    url varchar(48) NOT NULL,
    extra_info text,
    created_date timestamptz NOT NULL DEFAULT now(),
	modified_date timestamptz NOT NULL DEFAULT now()
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE jobs;