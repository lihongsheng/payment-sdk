-- ============================================================
-- 支付宝 (alipay)
-- 旧: {"app_id":"...", "rsa_private_key":"...", "rsa_app_crt":"...",
--       "rsa_root_crt":"...", "app_auth_token":"..."}
-- 新: {"merchant":{"app_id":"...", "app_auth_token":"..."},
--       "cert":{"rsa_private_key":"...", "rsa_app_crt":"...",
--               "rsa_root_crt":"...", "rsa_public_key":"...",
--               "rsa_app_cert_sn":"...", "rsa_root_cert_sn":"..."}}
--       "proxy":{...}}
--
-- 注意: rsa_public_key/rsa_app_cert_sn/rsa_root_cert_sn 之前由
--       initConfig() 从 crt 中自动提取，此处不做重复；如果旧数据中已有，
--       同样保留。
-- ============================================================
UPDATE payment_account
SET channel_config = jsonb_build_object(
  'merchant', jsonb_build_object(
    'app_id',         (channel_config->>'app_id'),
    'app_auth_token', (channel_config->>'app_auth_token')
  ),
  'cert', jsonb_build_object(
    'rsa_private_key',  (channel_config->>'rsa_private_key'),
    'rsa_app_crt',      (channel_config->>'rsa_app_crt'),
    'rsa_root_crt',     (channel_config->>'rsa_root_crt'),
    'rsa_public_key',   (channel_config->>'rsa_public_key'),
    'rsa_app_cert_sn',  (channel_config->>'rsa_app_cert_sn'),
    'rsa_root_cert_sn', (channel_config->>'rsa_root_cert_sn')
  ),
  'proxy', jsonb_build_object(
    'host',      (channel_config->'proxy'->>'host'),
    'port',      (channel_config->'proxy'->>'port'),
    'user_name', (channel_config->'proxy'->>'user_name'),
    'password',  (channel_config->'proxy'->>'password')
  )
)
WHERE channel = 'Alipay'
  AND channel_config IS NOT NULL
  AND channel_config != ''
  AND channel_config != '{}'
  AND channel_config->>'merchant' IS NULL;  -- 跳过已迁移的记录