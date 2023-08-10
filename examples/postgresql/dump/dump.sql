CREATE SEQUENCE accounts_seq START WITH 1 INCREMENT BY 1 NO MINVALUE MAXVALUE 2147483647 CACHE 1;

CREATE TABLE accounts (
	id uuid DEFAULT "uuid_generate_v4()" NOT NULL,
	email character varying(256) NOT NULL,
	hashed_password character varying(256) NOT NULL,
	salt character varying(8) NOT NULL,
	code character varying(64) NOT NULL,
	notification_id int8,
	role int2,
	status int2,
	flags int4,
	freezed_at timestamp without time zone,
	deleted_at timestamp without time zone,
	created_at timestamp without time zone DEFAULT (now())::timestamp(0) without time zone,
	updated_at timestamp without time zone DEFAULT (now())::timestamp(0) without time zone
);

COMMENT ON TABLE accounts IS 'アカウントマスタ';

COMMENT ON COLUMN accounts.id IS 'primary key';

COMMENT ON COLUMN accounts.email IS 'メールアドレス';

COMMENT ON COLUMN accounts.hashed_password IS 'ハッシュ化済みパスワード';

COMMENT ON COLUMN accounts.salt IS 'ソルト';

COMMENT ON COLUMN accounts.code IS '表示ID';

COMMENT ON COLUMN accounts.notification_id IS 'notifications.id';

COMMENT ON COLUMN accounts.role IS 'ロール';

COMMENT ON COLUMN accounts.status IS 'ステータス';

COMMENT ON COLUMN accounts.flags IS 'フラグ';

COMMENT ON COLUMN accounts.freezed_at IS '凍結日';

COMMENT ON COLUMN accounts.deleted_at IS '削除日';

COMMENT ON COLUMN accounts.created_at IS '作成日';

COMMENT ON COLUMN accounts.updated_at IS '更新日';

ALTER TABLE accounts ADD CONSTRAINT accounts_primary_idx PRIMARY KEY (id);

CREATE INDEX accounts_email_IDX ON accounts (email);

CREATE INDEX accounts_code_IDX ON accounts (code);

CREATE SEQUENCE pincodes_seq START WITH 1 INCREMENT BY 1 NO MINVALUE MAXVALUE 2147483647 CACHE 1;

CREATE TABLE pincodes (
	id uuid DEFAULT "uuid_generate_v4()" NOT NULL,
	pin character varying(6) NOT NULL,
	created_at timestamp without time zone DEFAULT (now())::timestamp(0) without time zone,
	updated_at timestamp without time zone DEFAULT (now())::timestamp(0) without time zone
);

COMMENT ON TABLE pincodes IS 'ピンコードマスタ';

COMMENT ON COLUMN pincodes.id IS 'primary key';

COMMENT ON COLUMN pincodes.pin IS 'ピン';

COMMENT ON COLUMN pincodes.created_at IS '作成日';

COMMENT ON COLUMN pincodes.updated_at IS '更新日';

ALTER TABLE pincodes ADD CONSTRAINT pincodes_primary_idx PRIMARY KEY (id);

CREATE INDEX pincodes_pin_IDX ON pincodes (pin);

CREATE SEQUENCE account_activates_seq START WITH 1 INCREMENT BY 1 NO MINVALUE MAXVALUE 2147483647 CACHE 1;

CREATE TABLE account_activates (
	id uuid DEFAULT "uuid_generate_v4()" NOT NULL,
	account_id int8 NOT NULL,
	uuid character varying(64) NOT NULL,
	pincode_id int8 NOT NULL,
	expired_at timestamp without time zone,
	activated_at timestamp without time zone,
	last_login_at timestamp without time zone,
	created_at timestamp without time zone DEFAULT (now())::timestamp(0) without time zone,
	updated_at timestamp without time zone DEFAULT (now())::timestamp(0) without time zone
);

COMMENT ON TABLE account_activates IS 'UUID管理マスタ';

COMMENT ON COLUMN account_activates.id IS 'primary key';

COMMENT ON COLUMN account_activates.account_id IS 'accounts.id';

COMMENT ON COLUMN account_activates.uuid IS 'UUID';

COMMENT ON COLUMN account_activates.pincode_id IS 'pincodes.id';

COMMENT ON COLUMN account_activates.expired_at IS 'PIN有効期限日時';

COMMENT ON COLUMN account_activates.activated_at IS 'アクティベート日時';

COMMENT ON COLUMN account_activates.last_login_at IS '最終ログイン日時';

COMMENT ON COLUMN account_activates.created_at IS '作成日';

COMMENT ON COLUMN account_activates.updated_at IS '更新日';

ALTER TABLE account_activates ADD CONSTRAINT account_activates_primary_idx PRIMARY KEY (id);

CREATE INDEX account_activates_uuid_IDX ON account_activates (uuid);

ALTER TABLE account_activates ADD CONSTRAINT account_activates_account_id_IDX UNIQUE (account_id);

ALTER TABLE account_activates ADD CONSTRAINT account_activates_pincode_id_IDX UNIQUE (pincode_id);

ALTER TABLE ONLY account_activates ADD CONSTRAINT account_activates_accounts_FK FOREIGN KEY (account_id) REFERENCES accounts(id);

ALTER TABLE ONLY account_activates ADD CONSTRAINT account_activates_pincodes_FK FOREIGN KEY (pincode_id) REFERENCES pincodes(id);

CREATE SEQUENCE account_pincodes_seq START WITH 1 INCREMENT BY 1 NO MINVALUE MAXVALUE 2147483647 CACHE 1;

CREATE TABLE account_pincodes (
	id uuid DEFAULT "uuid_generate_v4()" NOT NULL,
	account_id int8 NOT NULL,
	pincode_id int8 NOT NULL,
	expired_at timestamp without time zone,
	deleted_at timestamp without time zone,
	created_at timestamp without time zone DEFAULT (now())::timestamp(0) without time zone,
	updated_at timestamp without time zone DEFAULT (now())::timestamp(0) without time zone
);

COMMENT ON TABLE account_pincodes IS 'アカウントとピンコードの紐付け';

COMMENT ON COLUMN account_pincodes.id IS 'primary key';

COMMENT ON COLUMN account_pincodes.account_id IS 'accounts.id';

COMMENT ON COLUMN account_pincodes.pincode_id IS 'pincodes.id';

COMMENT ON COLUMN account_pincodes.expired_at IS 'PIN有効期限日時';

COMMENT ON COLUMN account_pincodes.deleted_at IS '使用済み日時';

COMMENT ON COLUMN account_pincodes.created_at IS '作成日';

COMMENT ON COLUMN account_pincodes.updated_at IS '更新日';

ALTER TABLE account_pincodes ADD CONSTRAINT account_pincodes_primary_idx PRIMARY KEY (id);

CREATE INDEX account_pincodes_account_id_IDX ON account_pincodes (account_id);

CREATE INDEX account_pincodes_pincode_id_IDX ON account_pincodes (pincode_id);

ALTER TABLE ONLY account_pincodes ADD CONSTRAINT account_pincodes_accounts_FK FOREIGN KEY (account_id) REFERENCES accounts(id);

ALTER TABLE ONLY account_pincodes ADD CONSTRAINT account_pincodes_pincodes_FK FOREIGN KEY (pincode_id) REFERENCES pincodes(id);

CREATE SEQUENCE action_logs_seq START WITH 1 INCREMENT BY 1 NO MINVALUE MAXVALUE 2147483647 CACHE 1;

CREATE TABLE action_logs (
	_id character varying(64) NOT NULL,
	uuid character varying(64) NOT NULL,
	email character varying(256),
	action_type int2,
	message text NOT NULL,
	recorded_at timestamp without time zone NOT NULL,
	created_at timestamp without time zone DEFAULT (now())::timestamp(0) without time zone,
	updated_at timestamp without time zone DEFAULT (now())::timestamp(0) without time zone
);

COMMENT ON TABLE action_logs IS 'アクションログ ESIDX';

COMMENT ON COLUMN action_logs._id IS 'ID';

COMMENT ON COLUMN action_logs.uuid IS 'UUID';

COMMENT ON COLUMN action_logs.email IS 'メールアドレス';

COMMENT ON COLUMN action_logs.action_type IS 'タイプ';

COMMENT ON COLUMN action_logs.message IS 'メッセージ';

COMMENT ON COLUMN action_logs.recorded_at IS '実行日時';

COMMENT ON COLUMN action_logs.created_at IS '作成日';

COMMENT ON COLUMN action_logs.updated_at IS '更新日';

