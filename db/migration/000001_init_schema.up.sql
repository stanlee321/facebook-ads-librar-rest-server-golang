CREATE TABLE "FacebookAd" (
  "id" BIGSERIAL PRIMARY KEY,
  "ad_id" bigint,
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
  "impressions_min" varchar,
  "spend_min" bigint,
  "spend_max" bigint,
  "currency" varchar,
  "ad_url" varchar,
  "social_media_facebook" varchar,
  "social_media_instagram" varchar,
  "social_media_whatsapp" varchar,
  "search_terms" varchar,
  "created_at" datetime
);

CREATE TABLE "FacebookDemos" (
  "id" BIGSERIAL PRIMARY KEY,
  "ad_id" bigint,
  "page_id" bigint,
  "age" int,
  "gender" varchar,
  "percentage" number,
  "ad_delivery_start_time" varchar,
  "created_at" datetime
);

CREATE TABLE "FacebookRegions" (
  "id" BIGSERIAL PRIMARY KEY,
  "ad_id" bigint,
  "page_id" bigint,
  "region" int,
  "percentage" varchar,
  "ad_delivery_start_time" varchar,
  "created_at" datetime
);

ALTER TABLE "FacebookDemos" ADD FOREIGN KEY ("ad_id") REFERENCES "FacebookAd" ("id");

ALTER TABLE "FacebookRegions" ADD FOREIGN KEY ("ad_id") REFERENCES "FacebookAd" ("id");

CREATE INDEX ON "FacebookAd" ("ad_id");

CREATE INDEX ON "FacebookAd" ("page_id");

CREATE INDEX ON "FacebookAd" ("page_name");

CREATE INDEX ON "FacebookDemos" ("ad_id");

CREATE INDEX ON "FacebookDemos" ("page_id");

CREATE INDEX ON "FacebookRegions" ("ad_id");

CREATE INDEX ON "FacebookRegions" ("page_id");