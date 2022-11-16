CREATE TABLE account_activates (
	id bigint(20) unsigned auto_increment NOT NULL COMMENT 'primary key',
	account_id bigint(20) unsigned NOT NULL COMMENT 'accounts.id',
	uuid varchar(64) NOT NULL COMMENT 'UUID',
	pincode_id bigint(20) unsigned NOT NULL COMMENT 'pincodes.id',
	expired_at timestamp DEFAULT NULL NULL COMMENT 'PIN有効期限日時',
	activated_at timestamp DEFAULT NULL NULL COMMENT 'アクティベート日時',
	last_login_at timestamp DEFAULT NULL NULL COMMENT '最終ログイン日時',
	created_at timestamp DEFAULT current_timestamp() NULL COMMENT '作成日',
	updated_at timestamp DEFAULT current_timestamp() NULL COMMENT '更新日',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='UUID管理マスタ';

CREATE UNIQUE INDEX account_activates_uuid_IDX USING BTREE ON account_activates (`uuid`);

CREATE INDEX account_activates_account_id_IDX USING BTREE ON account_activates (account_id);

CREATE INDEX account_activates_pincode_id_IDX USING BTREE ON account_activates (pincode_id);

CREATE TABLE pincodes (
	id bigint(20) unsigned auto_increment NOT NULL COMMENT 'primary key',
	pin varchar(6) NOT NULL COMMENT 'ピン',
	created_at timestamp DEFAULT current_timestamp() NULL COMMENT '作成日',
	updated_at timestamp DEFAULT current_timestamp() NULL COMMENT '更新日',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='ピンコードマスタ';

CREATE UNIQUE INDEX pincodes_pin_IDX USING BTREE ON pincodes (`pin`);

CREATE TABLE accounts (
	id bigint(20) unsigned auto_increment NOT NULL COMMENT 'primary key',
	email varchar(256) NOT NULL COMMENT 'メールアドレス',
	hashed_password varchar(256) NOT NULL COMMENT 'ハッシュ化済みパスワード',
	salt varchar(8) NOT NULL COMMENT 'ソルト',
	code varchar(64) NOT NULL COMMENT '表示ID',
	notification_id bigint(20) unsigned DEFAULT NULL NULL COMMENT 'notifications.id',
	role tinyint(3) unsigned DEFAULT NULL NULL COMMENT 'ロール',
	status tinyint(3) unsigned DEFAULT NULL NULL COMMENT 'ステータス',
	flags int(10) unsigned DEFAULT NULL NULL COMMENT 'フラグ',
	freezed_at timestamp DEFAULT NULL NULL COMMENT '削除日',
	deleted_at timestamp DEFAULT NULL NULL COMMENT '削除日',
	created_at timestamp DEFAULT current_timestamp() NULL COMMENT '作成日',
	updated_at timestamp DEFAULT current_timestamp() NULL COMMENT '更新日',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='アカウントマスタ';

CREATE UNIQUE INDEX accounts_email_IDX USING BTREE ON accounts (`email`);

CREATE UNIQUE INDEX accounts_code_IDX USING BTREE ON accounts (`code`);

CREATE TABLE account_pincodes (
	id bigint(20) unsigned auto_increment NOT NULL COMMENT 'primary key',
	account_id bigint(20) unsigned NOT NULL COMMENT 'accounts.id',
	pincode_id bigint(20) unsigned NOT NULL COMMENT 'pincodes.id',
	expired_at timestamp DEFAULT NULL NULL COMMENT 'PIN有効期限日時',
	deleted_at timestamp DEFAULT NULL NULL COMMENT '使用済み日時',
	created_at timestamp DEFAULT current_timestamp() NULL COMMENT '作成日',
	updated_at timestamp DEFAULT current_timestamp() NULL COMMENT '更新日',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='アカウントとピンコードの紐付け';

CREATE UNIQUE INDEX account_pincodes_account_id_IDX USING BTREE ON account_pincodes (`account_id`);

CREATE UNIQUE INDEX account_pincodes_pincode_id_IDX USING BTREE ON account_pincodes (`pincode_id`);

