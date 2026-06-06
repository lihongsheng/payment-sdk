-- ============================================================
-- 拉卡拉 (lakala)
-- 旧: {"app_id","mch_id","rsa_private_key","rsa_private_number",
--      "rsa_public_key","term_no","api_host"}
-- 新: {"merchant":{"app_id","mch_id","term_no"},
--      "cert":{"rsa_private_key","rsa_private_number","rsa_public_key"},
--      "api":{"api_host"}}
-- ============================================================
UPDATE payment_account
SET channel_config = jsonb_build_object(
  'merchant', jsonb_build_object(
    'app_id',  (channel_config->>'app_id'),
    'mch_id',  (channel_config->>'mch_id'),
    'term_no', (channel_config->>'term_no')
  ),
  'cert', jsonb_build_object(
    'rsa_private_key',    (channel_config->>'rsa_private_key'),
    'rsa_private_number', (channel_config->>'rsa_private_number'),
    'rsa_public_key',     (channel_config->>'rsa_public_key')
  ),
  'api', jsonb_build_object(
    'api_host', (channel_config->>'api_host')
  )
)
WHERE channel = 'Lakala'
  AND channel_config IS NOT NULL
  AND channel_config != ''
  AND channel_config != '{}'
  AND channel_config->>'merchant' IS NULL;