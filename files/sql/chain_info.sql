/*
 Navicat Premium Data Transfer

 Source Server         : localhost-pg-5432
 Source Server Type    : PostgreSQL
 Source Server Version : 140008 (140008)
 Source Host           : 127.0.0.1:5432
 Source Catalog        : mt_node
 Source Schema         : mt_auth

 Target Server Type    : PostgreSQL
 Target Server Version : 140008 (140008)
 File Encoding         : 65001

 Date: 08/06/2026 15:15:45
*/


-- ----------------------------
-- Table structure for chain_info
-- ----------------------------
DROP TABLE IF EXISTS "mt_auth"."chain_info";
CREATE TABLE "mt_auth"."chain_info" (
  "id" int8 NOT NULL DEFAULT nextval('"mt_auth".blockchain_networks_id_seq'::regclass),
  "chain_id" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "chain_id_decimal" int8 NOT NULL,
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "short_name" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "slug" varchar(50) COLLATE "pg_catalog"."default",
  "chain_type" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "environment" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "is_active" bool DEFAULT true,
  "rpc_urls" jsonb NOT NULL,
  "wss_urls" jsonb,
  "default_rpc_index" int4 DEFAULT 0,
  "explorer_url" varchar(500) COLLATE "pg_catalog"."default",
  "api_url" varchar(500) COLLATE "pg_catalog"."default",
  "native_currency" jsonb NOT NULL,
  "icon_url" varchar(500) COLLATE "pg_catalog"."default",
  "color" varchar(20) COLLATE "pg_catalog"."default",
  "is_testnet" bool DEFAULT false,
  "currency_symbol" varchar(20) COLLATE "pg_catalog"."default",
  "block_explorer_name" varchar(100) COLLATE "pg_catalog"."default",
  "display_order" int4 DEFAULT 0,
  "is_deprecated" bool DEFAULT false,
  "supports_eip1559" bool DEFAULT false,
  "supports_batch_requests" bool DEFAULT true,
  "created_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz(6) DEFAULT CURRENT_TIMESTAMP
)
;
ALTER TABLE "mt_auth"."chain_info" OWNER TO "postgres";
COMMENT ON COLUMN "mt_auth"."chain_info"."id" IS '自增主键';
COMMENT ON COLUMN "mt_auth"."chain_info"."chain_id" IS '链ID（字符串存储，兼容16进制和10进制）';
COMMENT ON COLUMN "mt_auth"."chain_info"."chain_id_decimal" IS '链ID（10进制数字，用于排序和比较）';
COMMENT ON COLUMN "mt_auth"."chain_info"."name" IS '区块链网络全称';
COMMENT ON COLUMN "mt_auth"."chain_info"."short_name" IS '链简称（用于内部标识）';
COMMENT ON COLUMN "mt_auth"."chain_info"."slug" IS 'URL友好的标识符';
COMMENT ON COLUMN "mt_auth"."chain_info"."chain_type" IS '链类型';
COMMENT ON COLUMN "mt_auth"."chain_info"."environment" IS '网络环境';
COMMENT ON COLUMN "mt_auth"."chain_info"."is_active" IS '是否启用';
COMMENT ON COLUMN "mt_auth"."chain_info"."rpc_urls" IS 'RPC URL列表（JSON数组）';
COMMENT ON COLUMN "mt_auth"."chain_info"."wss_urls" IS 'WebSocket URL列表（JSON数组）';
COMMENT ON COLUMN "mt_auth"."chain_info"."default_rpc_index" IS '默认RPC索引';
COMMENT ON COLUMN "mt_auth"."chain_info"."explorer_url" IS '区块浏览器URL';
COMMENT ON COLUMN "mt_auth"."chain_info"."api_url" IS '区块浏览器API URL';
COMMENT ON COLUMN "mt_auth"."chain_info"."native_currency" IS '原生代币信息';
COMMENT ON COLUMN "mt_auth"."chain_info"."icon_url" IS '图标URL';
COMMENT ON COLUMN "mt_auth"."chain_info"."color" IS '主题颜色（16进制颜色码）';
COMMENT ON COLUMN "mt_auth"."chain_info"."is_testnet" IS '是否是测试网';
COMMENT ON COLUMN "mt_auth"."chain_info"."currency_symbol" IS '货币符号（显示用）';
COMMENT ON COLUMN "mt_auth"."chain_info"."block_explorer_name" IS '区块浏览器名称';
COMMENT ON COLUMN "mt_auth"."chain_info"."display_order" IS '显示排序';
COMMENT ON COLUMN "mt_auth"."chain_info"."is_deprecated" IS '是否已弃用';
COMMENT ON COLUMN "mt_auth"."chain_info"."supports_eip1559" IS '是否支持EIP-1559';
COMMENT ON COLUMN "mt_auth"."chain_info"."supports_batch_requests" IS '是否支持批量请求';
COMMENT ON COLUMN "mt_auth"."chain_info"."created_at" IS '创建时间';
COMMENT ON COLUMN "mt_auth"."chain_info"."updated_at" IS '更新时间';
COMMENT ON TABLE "mt_auth"."chain_info" IS '区块链网络配置表';

-- ----------------------------
-- Indexes structure for table chain_info
-- ----------------------------
CREATE INDEX "idx_blockchain_networks_chain_metadata" ON "mt_auth"."chain_info" USING btree (
  "short_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "chain_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "is_testnet" "pg_catalog"."bool_ops" ASC NULLS LAST
);
CREATE INDEX "idx_blockchain_networks_chain_search" ON "mt_auth"."chain_info" USING btree (
  "environment" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "is_active" "pg_catalog"."bool_ops" ASC NULLS LAST,
  "chain_type" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "display_order" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "idx_blockchain_networks_chain_type" ON "mt_auth"."chain_info" USING btree (
  "chain_type" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_blockchain_networks_created_at" ON "mt_auth"."chain_info" USING btree (
  "created_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);
CREATE INDEX "idx_blockchain_networks_display_order" ON "mt_auth"."chain_info" USING btree (
  "display_order" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "idx_blockchain_networks_environment" ON "mt_auth"."chain_info" USING btree (
  "environment" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_blockchain_networks_is_active" ON "mt_auth"."chain_info" USING btree (
  "is_active" "pg_catalog"."bool_ops" ASC NULLS LAST
);
CREATE INDEX "idx_blockchain_networks_short_name" ON "mt_auth"."chain_info" USING btree (
  "short_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Triggers structure for table chain_info
-- ----------------------------
CREATE TRIGGER "update_blockchain_networks_updated_at" BEFORE UPDATE ON "mt_auth"."chain_info"
FOR EACH ROW
EXECUTE PROCEDURE "mt_auth"."update_updated_at_column"();

-- ----------------------------
-- Uniques structure for table chain_info
-- ----------------------------
ALTER TABLE "mt_auth"."chain_info" ADD CONSTRAINT "unique_chain_id" UNIQUE ("chain_id");
ALTER TABLE "mt_auth"."chain_info" ADD CONSTRAINT "unique_chain_id_decimal" UNIQUE ("chain_id_decimal");

-- ----------------------------
-- Checks structure for table chain_info
-- ----------------------------
ALTER TABLE "mt_auth"."chain_info" ADD CONSTRAINT "blockchain_networks_chain_type_check" CHECK (chain_type::text = ANY (ARRAY['EVM'::character varying, 'NON_EVM'::character varying, 'L2'::character varying, 'L3'::character varying]::text[]));
ALTER TABLE "mt_auth"."chain_info" ADD CONSTRAINT "blockchain_networks_environment_check" CHECK (environment::text = ANY (ARRAY['mainnet'::character varying, 'testnet'::character varying, 'devnet'::character varying]::text[]));

-- ----------------------------
-- Primary Key structure for table chain_info
-- ----------------------------
ALTER TABLE "mt_auth"."chain_info" ADD CONSTRAINT "blockchain_networks_pkey" PRIMARY KEY ("id");
