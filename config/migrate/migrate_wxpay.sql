-- ============================================================
-- 微信 (wxpay)
-- 旧: {"app_id":"...","mch_id":"...","app_secret":"...",
--      "api_secret":"...","rsa_private":"...","rsa_private_number":"...",
--      "rsa_public":"...","rsa_public_number":"...",
--      "score_service_id":"..."}
-- 新: {"merchant":{"app_id","mch_id","app_secret"},
--      "cert":{"api_secret","rsa_private","rsa_private_number",
--              "rsa_public","rsa_public_number"},
--      "service":{"score_service_id"}}
-- ============================================================
UPDATE payment_account
SET channel_config = jsonb_build_object(
  'merchant', jsonb_build_object(
    'app_id',     (channel_config->>'app_id'),
    'mch_id',     (channel_config->>'mch_id'),
    'app_secret', (channel_config->>'app_secret')
  ),
  'cert', jsonb_build_object(
    'api_secret',         (channel_config->>'api_secret'),
    'rsa_private',        (channel_config->>'rsa_private'),
    'rsa_private_number', (channel_config->>'rsa_private_number'),
    'rsa_public',         (channel_config->>'rsa_public'),
    'rsa_public_number',  (channel_config->>'rsa_public_number')
  ),
  'service', jsonb_build_object(
    'score_service_id', (channel_config->>'score_service_id')
  )
)
WHERE channel = 'Wechat'
  AND channel_config IS NOT NULL
  AND channel_config != ''
  AND channel_config != '{}'
  AND channel_config->>'merchant' IS NULL;