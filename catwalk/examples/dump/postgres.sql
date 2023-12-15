CREATE SCHEMA IF NOT EXISTS app;

CREATE TABLE "app"."accounts" (
	"id" uuid DEFAULT UUID() NOT NULL,
	"email" varchar(256) NOT NULL,
	"hashed_password" varchar(256) NOT NULL,
	"salt" varchar(8) NOT NULL,
	"notification_id" uuid NOT NULL,
	"role" int4 NOT NULL,
	"state" int2 NOT NULL,
	"flags" int8 NOT NULL,
	"freezed_at" timestamp NOT NULL,
	"deleted_at" timestamp NOT NULL,
	"created_at" timestamp DEFAULT current_timestamp() NOT NULL,
	"updated_at" timestamp DEFAULT current_timestamp()
);
COMMENT ON TABLE "app"."accounts" IS 'アカウントマスタ';

COMMENT ON COLUMN "app"."accounts"."アカウントマスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."accounts"."アカウントマスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."accounts"."アカウントマスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."accounts"."アカウントマスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."accounts"."アカウントマスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."accounts"."アカウントマスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."accounts"."アカウントマスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."accounts"."アカウントマスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."accounts"."アカウントマスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."accounts"."アカウントマスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."accounts"."アカウントマスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."accounts"."アカウントマスタ" IS '%!s(MISSING)';
ALTER TABLE "app"."accounts" ADD CONSTRAINT PRIMARY PRIMARY KEY ("id");
CREATE INDEX accounts_email_IDX ON "app"."accounts" ("email");
CREATE INDEX accounts_code_IDX ON "app"."accounts" ("email");
ALTER TABLE "app"."accounts" ADD CONSTRAINT accounts_multi_IDX UNIQUE ("email","deleted_at");

ALTER TABLE ONLY "app"."accounts" ADD CONSTRAINT "accounts_account_devices_FK" FOREIGN KEY ("id") REFERENCES "app"."account_devices"(account_id);

CREATE TABLE "app"."pins" (
	"id" uuid DEFAULT UUID() NOT NULL,
	"pin" varchar(128) NOT NULL,
	"usage" smallint NOT NULL,
	"created_at" timestamp DEFAULT current_timestamp() NOT NULL,
	"updated_at" timestamp DEFAULT current_timestamp()
);
COMMENT ON TABLE "app"."pins" IS 'ピンコードマスタ';

COMMENT ON COLUMN "app"."pins"."ピンコードマスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."pins"."ピンコードマスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."pins"."ピンコードマスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."pins"."ピンコードマスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."pins"."ピンコードマスタ" IS '%!s(MISSING)';
ALTER TABLE "app"."pins" ADD CONSTRAINT PRIMARY PRIMARY KEY ("id");
CREATE INDEX pins_pin_IDX ON "app"."pins" ("pin");


CREATE TABLE "app"."account_devices" (
	"id" uuid DEFAULT UUID() NOT NULL,
	"account_id" uuid NOT NULL,
	"uuid" varchar(64) NOT NULL,
	"activated_at" timestamp,
	"last_login_at" timestamp,
	"created_at" timestamp DEFAULT current_timestamp() NOT NULL,
	"updated_at" timestamp DEFAULT current_timestamp()
);
COMMENT ON TABLE "app"."account_devices" IS 'デバイス管理マスタ';

COMMENT ON COLUMN "app"."account_devices"."デバイス管理マスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."account_devices"."デバイス管理マスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."account_devices"."デバイス管理マスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."account_devices"."デバイス管理マスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."account_devices"."デバイス管理マスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."account_devices"."デバイス管理マスタ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."account_devices"."デバイス管理マスタ" IS '%!s(MISSING)';
ALTER TABLE "app"."account_devices" ADD CONSTRAINT PRIMARY PRIMARY KEY ("id");
CREATE INDEX account_devices_uuid_IDX ON "app"."account_devices" ("uuid");
CREATE INDEX account_devices_account_id_IDX ON "app"."account_devices" ("account_id");
CREATE INDEX account_devices_pin_id_IDX ON "app"."account_devices" ("pin_id");

ALTER TABLE ONLY "app"."account_devices" ADD CONSTRAINT "account_devices_accounts_FK" FOREIGN KEY ("account_id") REFERENCES "app"."accounts"(id);
ALTER TABLE ONLY "app"."account_devices" ADD CONSTRAINT "account_devices_pins_FK" FOREIGN KEY ("pin_id") REFERENCES "app"."pins"(id);

CREATE TABLE "app"."account_pins" (
	"id" uuid DEFAULT UUID() NOT NULL,
	"account_id" uuid NOT NULL,
	"pin_id" uuid NOT NULL,
	"expired_at" timestamp NOT NULL,
	"deleted_at" timestamp,
	"created_at" timestamp DEFAULT current_timestamp() NOT NULL,
	"updated_at" timestamp DEFAULT current_timestamp()
);
COMMENT ON TABLE "app"."account_pins" IS 'アカウントとピンの紐付け';

COMMENT ON COLUMN "app"."account_pins"."アカウントとピンの紐付け" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."account_pins"."アカウントとピンの紐付け" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."account_pins"."アカウントとピンの紐付け" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."account_pins"."アカウントとピンの紐付け" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."account_pins"."アカウントとピンの紐付け" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."account_pins"."アカウントとピンの紐付け" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."account_pins"."アカウントとピンの紐付け" IS '%!s(MISSING)';
ALTER TABLE "app"."account_pins" ADD CONSTRAINT PRIMARY PRIMARY KEY ("id");
CREATE INDEX account_pins_account_id_IDX ON "app"."account_pins" ("account_id");
CREATE INDEX account_pins_pin_id_IDX ON "app"."account_pins" ("pin_id");

ALTER TABLE ONLY "app"."account_pins" ADD CONSTRAINT "account_pins_accounts_FK" FOREIGN KEY ("account_id") REFERENCES "app"."accounts"(id);
ALTER TABLE ONLY "app"."account_pins" ADD CONSTRAINT "account_pins_pins_FK" FOREIGN KEY ("pin_id") REFERENCES "app"."pins"(id);

CREATE TABLE "app"."action_logs" (
	"id" uuid DEFAULT UUID() NOT NULL,
	"uuid" varchar(64) NOT NULL,
	"email" varchar(256) NOT NULL,
	"action_type" smallint NOT NULL,
	"log" text NOT NULL,
	"recorded_at" timestamp NOT NULL,
	"created_at" timestamp DEFAULT current_timestamp() NOT NULL,
	"updated_at" timestamp DEFAULT current_timestamp()
);
COMMENT ON TABLE "app"."action_logs" IS 'アクションログ';

COMMENT ON COLUMN "app"."action_logs"."アクションログ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."action_logs"."アクションログ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."action_logs"."アクションログ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."action_logs"."アクションログ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."action_logs"."アクションログ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."action_logs"."アクションログ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."action_logs"."アクションログ" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."action_logs"."アクションログ" IS '%!s(MISSING)';


CREATE TABLE "app"."items" (
	"id" uuid DEFAULT UUID() NOT NULL,
	"price" varchar(32) NOT NULL,
	"created_at" timestamp DEFAULT current_timestamp() NOT NULL,
	"updated_at" timestamp DEFAULT current_timestamp()
);
COMMENT ON TABLE "app"."items" IS 'データバリエーション';

COMMENT ON COLUMN "app"."items"."データバリエーション" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."items"."データバリエーション" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."items"."データバリエーション" IS '%!s(MISSING)';
COMMENT ON COLUMN "app"."items"."データバリエーション" IS '%!s(MISSING)';
ALTER TABLE "app"."items" ADD CONSTRAINT PRIMARY PRIMARY KEY ("id");
CREATE INDEX items_price_IDX ON "app"."items" ("price");


