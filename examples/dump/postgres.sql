CREATE SCHEMA IF NOT EXISTS public;

CREATE TABLE "public"."accounts" (
	"id" uuid DEFAULT uuid_generate_v4() NOT NULL,
	"email" varchar(256) NOT NULL,
	"hashed_password" varchar(256) NOT NULL,
	"salt" varchar(8) NOT NULL,
	"notification_id" uuid NOT NULL,
	"role" int4 NOT NULL,
	"state" int2 NOT NULL,
	"flags" int8 NOT NULL,
	"freezed_at" timestamp NULL,
	"deleted_at" timestamp NULL,
	"created_at" timestamp DEFAULT (now())::timestamp(0) without time zone NOT NULL,
	"updated_at" timestamp DEFAULT (now())::timestamp(0) without time zone NULL
);
COMMENT ON TABLE "public"."accounts" IS 'アカウントマスタ';

COMMENT ON COLUMN "public"."accounts"."id" IS 'ID';
COMMENT ON COLUMN "public"."accounts"."email" IS 'メールアドレス';
COMMENT ON COLUMN "public"."accounts"."hashed_password" IS 'ハッシュ化済みパスワード';
COMMENT ON COLUMN "public"."accounts"."salt" IS 'ソルト';
COMMENT ON COLUMN "public"."accounts"."notification_id" IS 'notifications.id';
COMMENT ON COLUMN "public"."accounts"."role" IS 'ロール';
COMMENT ON COLUMN "public"."accounts"."state" IS 'ステータス';
COMMENT ON COLUMN "public"."accounts"."flags" IS 'フラグ';
COMMENT ON COLUMN "public"."accounts"."freezed_at" IS '凍結日';
COMMENT ON COLUMN "public"."accounts"."deleted_at" IS '削除日';
COMMENT ON COLUMN "public"."accounts"."created_at" IS '作成日';
COMMENT ON COLUMN "public"."accounts"."updated_at" IS '更新日';
CREATE INDEX IF NOT EXISTS accounts_email_IDX ON "public"."accounts" ("email");
CREATE UNIQUE INDEX IF NOT EXISTS accounts_multi_IDX ON "public"."accounts" ("email","deleted_at");
ALTER TABLE "public"."accounts" ADD CONSTRAINT accounts_primary_IDX PRIMARY KEY ("id");
CREATE INDEX IF NOT EXISTS accounts_state_IDX ON "public"."accounts" ("state");

ALTER TABLE ONLY "public"."accounts" ADD CONSTRAINT "accounts_account_devices_FK" FOREIGN KEY ("id") REFERENCES "public"."account_devices"(account_id);

CREATE TABLE "public"."pins" (
	"id" uuid DEFAULT uuid_generate_v4() NOT NULL,
	"pin" varchar(128) NOT NULL,
	"usage" smallint NOT NULL,
	"created_at" timestamp DEFAULT (now())::timestamp(0) without time zone NOT NULL,
	"updated_at" timestamp DEFAULT (now())::timestamp(0) without time zone NULL
);
COMMENT ON TABLE "public"."pins" IS 'ピンコードマスタ';

COMMENT ON COLUMN "public"."pins"."id" IS 'ID';
COMMENT ON COLUMN "public"."pins"."pin" IS 'ピン';
COMMENT ON COLUMN "public"."pins"."usage" IS '用途';
COMMENT ON COLUMN "public"."pins"."created_at" IS '作成日';
COMMENT ON COLUMN "public"."pins"."updated_at" IS '更新日';
CREATE INDEX IF NOT EXISTS pins_pin_IDX ON "public"."pins" ("pin");
ALTER TABLE "public"."pins" ADD CONSTRAINT pins_primary_IDX PRIMARY KEY ("id");


CREATE TABLE "public"."account_devices" (
	"id" uuid DEFAULT uuid_generate_v4() NOT NULL,
	"account_id" uuid NOT NULL,
	"uuid" varchar(64) NOT NULL,
	"activated_at" timestamp NULL,
	"last_login_at" timestamp NULL,
	"created_at" timestamp DEFAULT (now())::timestamp(0) without time zone NOT NULL,
	"updated_at" timestamp DEFAULT (now())::timestamp(0) without time zone NULL
);
COMMENT ON TABLE "public"."account_devices" IS 'デバイス管理マスタ';

COMMENT ON COLUMN "public"."account_devices"."id" IS 'ID';
COMMENT ON COLUMN "public"."account_devices"."account_id" IS 'accounts.id';
COMMENT ON COLUMN "public"."account_devices"."uuid" IS 'デバイスID';
COMMENT ON COLUMN "public"."account_devices"."activated_at" IS 'アクティベート日時';
COMMENT ON COLUMN "public"."account_devices"."last_login_at" IS '最終ログイン日時';
COMMENT ON COLUMN "public"."account_devices"."created_at" IS '作成日';
COMMENT ON COLUMN "public"."account_devices"."updated_at" IS '更新日';
CREATE INDEX IF NOT EXISTS account_devices_account_id_IDX ON "public"."account_devices" ("account_id");
ALTER TABLE "public"."account_devices" ADD CONSTRAINT account_devices_primary_IDX PRIMARY KEY ("id");
CREATE INDEX IF NOT EXISTS account_devices_uuid_IDX ON "public"."account_devices" ("uuid");

ALTER TABLE ONLY "public"."account_devices" ADD CONSTRAINT "account_devices_accounts_FK" FOREIGN KEY ("account_id") REFERENCES "public"."accounts"(id);

CREATE TABLE "public"."account_pins" (
	"id" uuid DEFAULT uuid_generate_v4() NOT NULL,
	"account_id" uuid NOT NULL,
	"pin_id" uuid NOT NULL,
	"expired_at" timestamp DEFAULT "1900-01-01T00:00:00" NOT NULL,
	"deleted_at" timestamp NULL,
	"created_at" timestamp DEFAULT (now())::timestamp(0) without time zone NOT NULL,
	"updated_at" timestamp DEFAULT (now())::timestamp(0) without time zone NULL
);
COMMENT ON TABLE "public"."account_pins" IS 'アカウントとピンの紐付け';

COMMENT ON COLUMN "public"."account_pins"."id" IS 'ID';
COMMENT ON COLUMN "public"."account_pins"."account_id" IS 'accounts.id';
COMMENT ON COLUMN "public"."account_pins"."pin_id" IS 'pins.id';
COMMENT ON COLUMN "public"."account_pins"."expired_at" IS 'PIN有効期限日時';
COMMENT ON COLUMN "public"."account_pins"."deleted_at" IS '使用済み日時';
COMMENT ON COLUMN "public"."account_pins"."created_at" IS '作成日';
COMMENT ON COLUMN "public"."account_pins"."updated_at" IS '更新日';
CREATE INDEX IF NOT EXISTS account_pins_account_id_IDX ON "public"."account_pins" ("account_id");
CREATE INDEX IF NOT EXISTS account_pins_pin_id_IDX ON "public"."account_pins" ("pin_id");
ALTER TABLE "public"."account_pins" ADD CONSTRAINT account_pins_primary_IDX PRIMARY KEY ("id");

ALTER TABLE ONLY "public"."account_pins" ADD CONSTRAINT "account_pins_accounts_FK" FOREIGN KEY ("account_id") REFERENCES "public"."accounts"(id);
ALTER TABLE ONLY "public"."account_pins" ADD CONSTRAINT "account_pins_pins_FK" FOREIGN KEY ("pin_id") REFERENCES "public"."pins"(id);

CREATE TABLE "public"."action_logs" (
	"id" uuid DEFAULT uuid_generate_v4() NOT NULL,
	"uuid" varchar(64) NOT NULL,
	"email" varchar(256) NOT NULL,
	"action_type" smallint NOT NULL,
	"log" text NOT NULL,
	"recorded_at" timestamp DEFAULT current_timestamp() NOT NULL,
	"created_at" timestamp DEFAULT (now())::timestamp(0) without time zone NOT NULL,
	"updated_at" timestamp DEFAULT (now())::timestamp(0) without time zone NULL
);
COMMENT ON TABLE "public"."action_logs" IS 'アクションログ';

COMMENT ON COLUMN "public"."action_logs"."id" IS 'ID';
COMMENT ON COLUMN "public"."action_logs"."uuid" IS 'UUID';
COMMENT ON COLUMN "public"."action_logs"."email" IS 'メールアドレス';
COMMENT ON COLUMN "public"."action_logs"."action_type" IS 'タイプ';
COMMENT ON COLUMN "public"."action_logs"."log" IS 'メッセージ';
COMMENT ON COLUMN "public"."action_logs"."recorded_at" IS '実行日時';
COMMENT ON COLUMN "public"."action_logs"."created_at" IS '作成日';
COMMENT ON COLUMN "public"."action_logs"."updated_at" IS '更新日';


CREATE TABLE "public"."items" (
	"id" uuid DEFAULT uuid_generate_v4() NOT NULL,
	"price" varchar(32) NOT NULL,
	"created_at" timestamp DEFAULT (now())::timestamp(0) without time zone NOT NULL,
	"updated_at" timestamp DEFAULT (now())::timestamp(0) without time zone NULL
);
COMMENT ON TABLE "public"."items" IS 'データバリエーション';

COMMENT ON COLUMN "public"."items"."id" IS 'ID';
COMMENT ON COLUMN "public"."items"."price" IS '価格';
COMMENT ON COLUMN "public"."items"."created_at" IS '作成日';
COMMENT ON COLUMN "public"."items"."updated_at" IS '更新日';
CREATE INDEX IF NOT EXISTS items_price_IDX ON "public"."items" ("price");
ALTER TABLE "public"."items" ADD CONSTRAINT items_primary_IDX PRIMARY KEY ("id");


