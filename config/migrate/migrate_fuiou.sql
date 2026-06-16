-- ============================================================
-- 富友 (fuiou)
-- 旧: {"mch_id","api_secret","order_prefix","api_host","version",
--      "rsa_private_key","rsa_public_key",
--      "wechat_app_id","wechat_app_secret",
--      "alipay_app_id","alipay_rsa_private_key","alipay_rsa_root_crt"}
-- 新: {"merchant":{"mch_id","api_secret","order_prefix"},
--      "cert":{"rsa_private_key","rsa_public_key"},
--      "api":{"api_host","version"},
--      "wechat":{"app_id","app_secret"},
--      "alipay":{"app_id","rsa_private_key","rsa_root_crt"}}
-- ============================================================
UPDATE payment_account
SET channel_config = jsonb_build_object(
  'merchant', jsonb_build_object(
    'mch_id',       (channel_config->>'mch_id'),
    'api_secret',   (channel_config->>'api_secret'),
    'order_prefix', (channel_config->>'order_prefix')
  ),
  'cert', jsonb_build_object(
    'rsa_private_key', (channel_config->>'rsa_private_key'),
    'rsa_public_key',  (channel_config->>'rsa_public_key')
  ),
  'api', jsonb_build_object(
    'api_host', (channel_config->>'api_host'),
    'version',  (channel_config->>'version')
  ),
  'wechat', jsonb_build_object(
    'app_id',     (channel_config->>'wechat_app_id'),
    'app_secret', (channel_config->>'wechat_app_secret')
  ),
  'alipay', jsonb_build_object(
    'app_id',          (channel_config->>'alipay_app_id'),
    'rsa_private_key', (channel_config->>'alipay_rsa_private_key'),
    'rsa_root_crt',    (channel_config->>'alipay_rsa_root_crt')
  )
)
WHERE channel = 'Fuiou'
  AND channel_config IS NOT NULL
  AND channel_config != ''
  AND channel_config != '{}'
  AND channel_config->>'merchant' IS NULL;