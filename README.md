Create readonly account with access to stats DB


```
CREATE ROLE readonly;
GRANT USAGE ON SCHEMA public TO readonly;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO readonly;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO readonly;
CREATE USER gomastodonstats WITH PASSWORD 'superrandompassword';
GRANT readonly TO gomastodonstats;

CREATE DATABASE gomastodonstats WITH OWNER gomastodonstats;
\c gomastodonstats
CREATE TABLE IF NOT EXISTS statsdb (
	id INT NOT NULL,
	service VARCHAR(50) NOT NULL,
	metric_name VARCHAR(50) NOT NULL,
    metric_time INT NOT NULL,
	metric_value INT NOT NULL,
    PRIMARY KEY (id)
);
CREATE INDEX service_lookup ON statsdb USING btree (service,metric_name, metric_time);
```