Create readonly account with access to stats DB


```
CREATE USER gomastodonstats WITH PASSWORD 'superrandompassword';

\c pixelfed
GRANT SELECT ON ALL TABLES IN SCHEMA public TO gomastodonstats;

\c mtrx
GRANT SELECT ON ALL TABLES IN SCHEMA public TO gomastodonstats;

\c mastodon_production
GRANT SELECT ON ALL TABLES IN SCHEMA public TO gomastodonstats;

\c peertube_prod
GRANT SELECT ON ALL TABLES IN SCHEMA public TO gomastodonstats;

\c mobilizon_prod
GRANT SELECT ON ALL TABLES IN SCHEMA public TO gomastodonstats;

CREATE DATABASE gomastodonstats WITH OWNER gomastodonstats;

\c gomastodonstats
CREATE TABLE IF NOT EXISTS statsdb (
	id SERIAL,
	service VARCHAR(50) NOT NULL,
	metric_name VARCHAR(50) NOT NULL,
    metric_time TIMESTAMP NOT NULL,
	metric_value INT NOT NULL,
    PRIMARY KEY (id)
);
CREATE INDEX service_lookup ON statsdb USING btree (service,metric_name, metric_time);

GRANT ALL ON ALL TABLES IN SCHEMA public TO gomastodonstats;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO gomastodonstats;
```