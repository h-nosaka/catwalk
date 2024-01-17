CREATE DATABASE IF NOT EXISTS app CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

CREATE TABLE `app`.`accounts` (
	`id` uuid DEFAULT UUID() NOT NULL COMMENT 'ID',
	`email` varchar(256) NOT NULL COMMENT 'メールアドレス',
	`hashed_password` varchar(256) NOT NULL COMMENT 'ハッシュ化済みパスワード',
	`salt` varchar(8) NOT NULL COMMENT 'ソルト',
	`notification_id` uuid NOT NULL COMMENT 'notifications.id',
	`role` int unsigned NOT NULL COMMENT 'ロール',
	`state` tinyint unsigned NOT NULL COMMENT 'ステータス',
	`flags` bigint unsigned NOT NULL COMMENT 'フラグ',
	`freezed_at` timestamp NOT NULL COMMENT '凍結日',
	`deleted_at` timestamp NOT NULL COMMENT '削除日',
	`created_at` timestamp DEFAULT current_timestamp() NOT NULL COMMENT '作成日',
	`updated_at` timestamp DEFAULT current_timestamp() COMMENT '更新日'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='アカウントマスタ';

ALTER TABLE PRIMARY ADD CONSTRAINT `app`.`accounts` PRIMARY KEY (`id`);
CREATE INDEX accounts_email_IDX USING BTREE ON `app`.`accounts` (`email`);
CREATE INDEX accounts_code_IDX USING BTREE ON `app`.`accounts` (`email`);
CREATE UNIQUE INDEX accounts_multi_IDX USING BTREE ON `app`.`accounts` (`email`,`deleted_at`);

ALTER TABLE `app`.`accounts` ADD CONSTRAINT accounts_account_devices_FK FOREIGN KEY `app`.`accounts`(`id`) REFERENCES `app`.`account_devices`(`account_id`);

CREATE TABLE `app`.`pins` (
	`id` uuid DEFAULT UUID() NOT NULL COMMENT 'ID',
	`pin` varchar(128) NOT NULL COMMENT 'ピン',
	`usage` tinyint NOT NULL COMMENT '用途',
	`created_at` timestamp DEFAULT current_timestamp() NOT NULL COMMENT '作成日',
	`updated_at` timestamp DEFAULT current_timestamp() COMMENT '更新日'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='ピンコードマスタ';

ALTER TABLE PRIMARY ADD CONSTRAINT `app`.`pins` PRIMARY KEY (`id`);
CREATE INDEX pins_pin_IDX USING BTREE ON `app`.`pins` (`pin`);


CREATE TABLE `app`.`account_devices` (
	`id` uuid DEFAULT UUID() NOT NULL COMMENT 'ID',
	`account_id` uuid NOT NULL COMMENT 'accounts.id',
	`uuid` varchar(64) NOT NULL COMMENT 'デバイスID',
	`activated_at` timestamp COMMENT 'アクティベート日時',
	`last_login_at` timestamp COMMENT '最終ログイン日時',
	`created_at` timestamp DEFAULT current_timestamp() NOT NULL COMMENT '作成日',
	`updated_at` timestamp DEFAULT current_timestamp() COMMENT '更新日'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='デバイス管理マスタ';

ALTER TABLE PRIMARY ADD CONSTRAINT `app`.`account_devices` PRIMARY KEY (`id`);
CREATE INDEX account_devices_uuid_IDX USING BTREE ON `app`.`account_devices` (`uuid`);
CREATE INDEX account_devices_account_id_IDX USING BTREE ON `app`.`account_devices` (`account_id`);
CREATE INDEX account_devices_pin_id_IDX USING BTREE ON `app`.`account_devices` (`pin_id`);

ALTER TABLE `app`.`account_devices` ADD CONSTRAINT account_devices_accounts_FK FOREIGN KEY `app`.`account_devices`(`account_id`) REFERENCES `app`.`accounts`(`id`);
ALTER TABLE `app`.`account_devices` ADD CONSTRAINT account_devices_pins_FK FOREIGN KEY `app`.`account_devices`(`pin_id`) REFERENCES `app`.`pins`(`id`);

CREATE TABLE `app`.`account_pins` (
	`id` uuid DEFAULT UUID() NOT NULL COMMENT 'ID',
	`account_id` uuid NOT NULL COMMENT 'accounts.id',
	`pin_id` uuid NOT NULL COMMENT 'pins.id',
	`expired_at` timestamp NOT NULL COMMENT 'PIN有効期限日時',
	`deleted_at` timestamp COMMENT '使用済み日時',
	`created_at` timestamp DEFAULT current_timestamp() NOT NULL COMMENT '作成日',
	`updated_at` timestamp DEFAULT current_timestamp() COMMENT '更新日'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='アカウントとピンの紐付け';

ALTER TABLE PRIMARY ADD CONSTRAINT `app`.`account_pins` PRIMARY KEY (`id`);
CREATE INDEX account_pins_account_id_IDX USING BTREE ON `app`.`account_pins` (`account_id`);
CREATE INDEX account_pins_pin_id_IDX USING BTREE ON `app`.`account_pins` (`pin_id`);

ALTER TABLE `app`.`account_pins` ADD CONSTRAINT account_pins_accounts_FK FOREIGN KEY `app`.`account_pins`(`account_id`) REFERENCES `app`.`accounts`(`id`);
ALTER TABLE `app`.`account_pins` ADD CONSTRAINT account_pins_pins_FK FOREIGN KEY `app`.`account_pins`(`pin_id`) REFERENCES `app`.`pins`(`id`);

CREATE TABLE `app`.`action_logs` (
	`id` uuid DEFAULT UUID() NOT NULL COMMENT 'ID',
	`uuid` varchar(64) NOT NULL COMMENT 'UUID',
	`email` varchar(256) NOT NULL COMMENT 'メールアドレス',
	`action_type` smallint NOT NULL COMMENT 'タイプ',
	`log` mediumtext NOT NULL COMMENT 'メッセージ',
	`recorded_at` timestamp NOT NULL COMMENT '実行日時',
	`created_at` timestamp DEFAULT current_timestamp() NOT NULL COMMENT '作成日',
	`updated_at` timestamp DEFAULT current_timestamp() COMMENT '更新日'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='アクションログ';



CREATE TABLE `app`.`items` (
	`id` uuid DEFAULT UUID() NOT NULL COMMENT 'ID',
	`price` varchar(32) NOT NULL COMMENT '価格',
	`created_at` timestamp DEFAULT current_timestamp() NOT NULL COMMENT '作成日',
	`updated_at` timestamp DEFAULT current_timestamp() COMMENT '更新日'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='データバリエーション';

ALTER TABLE PRIMARY ADD CONSTRAINT `app`.`items` PRIMARY KEY (`id`);
CREATE INDEX items_price_IDX USING BTREE ON `app`.`items` (`price`);


