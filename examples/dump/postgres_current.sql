CREATE SCHEMA IF NOT EXISTS app;

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
CREATE INDEX IF NOT EXISTS accounts_email_idx ON "public"."accounts" ("email");
CREATE INDEX IF NOT EXISTS accounts_multi_idx ON "public"."accounts" ("email","deleted_at");
ALTER TABLE "public"."accounts" ADD CONSTRAINT accounts_primary_idx PRIMARY KEY ("id");
CREATE INDEX IF NOT EXISTS accounts_state_idx ON "public"."accounts" ("state");


CREATE TABLE "public"."pins" (
	"id" uuid DEFAULT uuid_generate_v4() NOT NULL,
	"pin" varchar(128) NOT NULL,
	"usage" int2 NOT NULL,
	"created_at" timestamp DEFAULT (now())::timestamp(0) without time zone NOT NULL,
	"updated_at" timestamp DEFAULT (now())::timestamp(0) without time zone NULL
);
COMMENT ON TABLE "public"."pins" IS 'ピンコードマスタ';

COMMENT ON COLUMN "public"."pins"."id" IS 'ID';
COMMENT ON COLUMN "public"."pins"."pin" IS 'ピン';
COMMENT ON COLUMN "public"."pins"."usage" IS '用途';
COMMENT ON COLUMN "public"."pins"."created_at" IS '作成日';
COMMENT ON COLUMN "public"."pins"."updated_at" IS '更新日';
CREATE INDEX IF NOT EXISTS pins_pin_idx ON "public"."pins" ("pin");
ALTER TABLE "public"."pins" ADD CONSTRAINT pins_primary_idx PRIMARY KEY ("id");


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
CREATE INDEX IF NOT EXISTS account_devices_account_id_idx ON "public"."account_devices" ("account_id");
ALTER TABLE "public"."account_devices" ADD CONSTRAINT account_devices_primary_idx PRIMARY KEY ("id");
CREATE INDEX IF NOT EXISTS account_devices_uuid_idx ON "public"."account_devices" ("uuid");


