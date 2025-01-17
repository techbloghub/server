-- Create "companies" table
CREATE TABLE "companies" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "create_time" timestamptz NOT NULL, "update_time" timestamptz NOT NULL, "name" character varying NOT NULL, "logo_url" character varying NOT NULL, "blog_url" character varying NOT NULL, "rss_url" character varying NOT NULL, PRIMARY KEY ("id"));
