CREATE TABLE "FacebookAd" (
  "ad_id" bigint PRIMARY KEY NOT NULL,
  "job_id" bigint,
  "page_id" bigint,
  "page_name" varchar,
  "ad_snapshot_url" varchar,
  "ad_creative_body" varchar,
  "ad_creative_link_caption" varchar,
  "ad_creative_link_description" varchar,
  "ad_creative_link_title" varchar,
  "ad_delivery_start_time" varchar,
  "ad_delivery_stop_time" varchar,
  "funding_entity" varchar,
  "impressions_min" int,
  "impressions_max" int,
  "spend_min" int,
  "spend_max" int,
  "currency" varchar,
  "ad_url" varchar,
  "social_media_facebook" varchar,
  "social_media_instagram" varchar,
  "social_media_whatsapp" varchar,
  "search_terms" varchar,
  "ad_creation_time" varchar,
  "potential_reach_max" int,
  "potential_reach_min" int,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "FacebookDemos" (
  "id" BIGSERIAL PRIMARY KEY,
  "job_id" bigint,
  "ad_id" bigint,
  "page_id" bigint,
  "age" varchar,
  "gender" varchar,
  "percentage" decimal,
  "ad_delivery_start_time" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "FacebookRegions" (
  "id" BIGSERIAL PRIMARY KEY,
  "job_id" bigint,
  "ad_id" bigint,
  "page_id" bigint,
  "region" varchar,
  "percentage" decimal,
  "ad_delivery_start_time" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "FacebookJob" (
  "id" BIGSERIAL PRIMARY KEY,
  "search_terms" varchar,
  "access_token" varchar,
  "page_total" bigint,
  "search_total" bigint,
  "ad_active_status" varchar,
  "ad_delivery_date_max" varchar,
  "ad_delivery_date_min" varchar,
  "ad_reached_countries" varchar,
  "total_found_ads" bigint,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "FacebookAd" ADD FOREIGN KEY ("job_id") REFERENCES "FacebookJob" ("id");

ALTER TABLE "FacebookDemos" ADD FOREIGN KEY ("job_id") REFERENCES "FacebookJob" ("id");

ALTER TABLE "FacebookDemos" ADD FOREIGN KEY ("ad_id") REFERENCES "FacebookAd" ("ad_id");

ALTER TABLE "FacebookRegions" ADD FOREIGN KEY ("job_id") REFERENCES "FacebookJob" ("id");

ALTER TABLE "FacebookRegions" ADD FOREIGN KEY ("ad_id") REFERENCES "FacebookAd" ("ad_id");

CREATE INDEX ON "FacebookAd" ("ad_id");

CREATE INDEX ON "FacebookAd" ("job_id");

CREATE INDEX ON "FacebookJob" ("search_terms", "page_total", "search_total", "ad_active_status", "ad_delivery_date_max", "ad_delivery_date_min", "ad_reached_countries");