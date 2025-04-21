-- Drop indexes
DROP INDEX IF EXISTS idx_banners_position;
DROP INDEX IF EXISTS idx_product_downloads_product_id;

-- Drop tables in reverse order to handle dependencies
DROP TABLE IF EXISTS banners;
DROP TABLE IF EXISTS product_downloads;
DROP TABLE IF EXISTS topic_product;
DROP TABLE IF EXISTS topics;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;