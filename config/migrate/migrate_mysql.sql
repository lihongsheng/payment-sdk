-- ============================================================
-- MySQL 版本（channel_config 是 TEXT/JSON 列）
-- 使用 JSON_OBJECT / JSON_EXTRACT / JSON_UNQUOTE
-- 表名按你项目实际改：payment_account
-- ============================================================

-- 支付宝
UPDATE payment_account
SET channel_config = JSON_OBJECT(
  'merchant', JSON_OBJECT(
    'app_id',         JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.app_id')),
    'app_auth_token', JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.app_auth_token'))
  ),
  'cert', JSON_OBJECT(
    'rsa_private_key',  JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.rsa_private_key')),
    'rsa_app_crt',      JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.rsa_app_crt')),
    'rsa_root_crt',     JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.rsa_root_crt')),
    'rsa_public_key',   JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.rsa_public_key')),
    'rsa_app_cert_sn',  JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.rsa_app_cert_sn')),
    'rsa_root_cert_sn', JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.rsa_root_cert_sn'))
  ),
  'proxy', JSON_OBJECT(
    'host',      JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.proxy.host')),
    'port',      JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.proxy.port')),
    'user_name', JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.proxy.user_name')),
    'password',  JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.proxy.password'))
  )
)
WHERE channel = 'Alipay'
  AND channel_config IS NOT NULL
  AND channel_config != ''
  AND channel_config != '{}'
  AND JSON_EXTRACT(channel_config, '$.merchant') IS NULL;

-- 微信
UPDATE payment_account
SET channel_config = JSON_OBJECT(
  'merchant', JSON_OBJECT(
    'app_id',     JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.app_id')),
    'mch_id',     JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.mch_id')),
    'app_secret', JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.app_secret'))
  ),
  'cert', JSON_OBJECT(
    'api_secret',         JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.api_secret')),
    'rsa_private',        JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.rsa_private')),
    'rsa_private_number', JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.rsa_private_number')),
    'rsa_public',         JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.rsa_public')),
    'rsa_public_number',  JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.rsa_public_number'))
  ),
  'service', JSON_OBJECT(
    'score_service_id', JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.score_service_id'))
  )
)
WHERE channel = 'Wechat'
  AND channel_config IS NOT NULL
  AND channel_config != ''
  AND channel_config != '{}'
  AND JSON_EXTRACT(channel_config, '$.merchant') IS NULL;

-- 拉卡拉
UPDATE payment_account
SET channel_config = JSON_OBJECT(
  'merchant', JSON_OBJECT(
    'app_id',  JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.app_id')),
    'mch_id',  JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.mch_id')),
    'term_no', JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.term_no'))
  ),
  'cert', JSON_OBJECT(
    'rsa_private_key',    JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.rsa_private_key')),
    'rsa_private_number', JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.rsa_private_number')),
    'rsa_public_key',     JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.rsa_public_key'))
  ),
  'api', JSON_OBJECT(
    'api_host', JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.api_host'))
  )
)
WHERE channel = 'Lakala'
  AND channel_config IS NOT NULL
  AND channel_config != ''
  AND channel_config != '{}'
  AND JSON_EXTRACT(channel_config, '$.merchant') IS NULL;

-- 富友
UPDATE payment_account
SET channel_config = JSON_OBJECT(
  'merchant', JSON_OBJECT(
    'mch_id',       JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.mch_id')),
    'api_secret',   JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.api_secret')),
    'order_prefix', JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.order_prefix'))
  ),
  'cert', JSON_OBJECT(
    'rsa_private_key', JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.rsa_private_key')),
    'rsa_public_key',  JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.rsa_public_key'))
  ),
  'api', JSON_OBJECT(
    'api_host', JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.api_host')),
    'version',  JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.version'))
  ),
  'wechat', JSON_OBJECT(
    'app_id',     JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.wechat_app_id')),
    'app_secret', JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.wechat_app_secret'))
  ),
  'alipay', JSON_OBJECT(
    'app_id',          JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.alipay_app_id')),
    'rsa_private_key', JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.alipay_rsa_private_key')),
    'rsa_root_crt',    JSON_UNQUOTE(JSON_EXTRACT(channel_config, '$.alipay_rsa_root_crt'))
  )
)
WHERE channel = 'Fuiou'
  AND channel_config IS NOT NULL
  AND channel_config != ''
  AND channel_config != '{}'
  AND JSON_EXTRACT(channel_config, '$.merchant') IS NULL;