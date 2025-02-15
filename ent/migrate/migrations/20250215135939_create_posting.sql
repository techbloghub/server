-- Create "postings" table
CREATE TABLE "postings" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "title" character varying NOT NULL, "url" character varying NOT NULL, "published_at" timestamptz NOT NULL, "tags" text[] NULL, "company_postings" bigint NULL, PRIMARY KEY ("id"), CONSTRAINT "postings_companies_postings" FOREIGN KEY ("company_postings") REFERENCES "companies" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "postings_url_key" to table: "postings"
CREATE UNIQUE INDEX "postings_url_key" ON "postings" ("url");
