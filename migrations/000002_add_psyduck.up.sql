-- Create categories table
CREATE TABLE IF NOT EXISTS categories (
    id          VARCHAR(7)   NOT NULL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL UNIQUE,
    slug        VARCHAR(255) NOT NULL UNIQUE,
    description VARCHAR(255),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create products table
CREATE TABLE IF NOT EXISTS products (
    id          VARCHAR(7)   NOT NULL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    slug        VARCHAR(255) NOT NULL UNIQUE,
    content     TEXT,
    category_id VARCHAR(7) REFERENCES categories ON DELETE SET NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    description VARCHAR(255),
    thumbnail   VARCHAR(255)
);

-- Create topics table
CREATE TABLE IF NOT EXISTS topics (
    id         VARCHAR(7)   NOT NULL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL UNIQUE,
    slug       VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create topic_product junction table
CREATE TABLE IF NOT EXISTS topic_product (
    id         UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    product_id VARCHAR(7) NOT NULL REFERENCES products ON DELETE CASCADE,
    topic_id   VARCHAR(7) NOT NULL REFERENCES topics ON DELETE CASCADE,
    UNIQUE (product_id, topic_id)
);

-- Create product_downloads table
CREATE TABLE IF NOT EXISTS product_downloads (
    id             UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    product_id     VARCHAR(7)   NOT NULL REFERENCES products ON DELETE CASCADE,
    title          VARCHAR(255) NOT NULL,
    url            TEXT         NOT NULL,
    file_size      BIGINT,
    file_type      VARCHAR(50),
    is_active      BOOLEAN DEFAULT TRUE,
    download_count INTEGER DEFAULT 0,
    created_at     TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_product_downloads_product_id ON product_downloads (product_id);

-- Create banners table
CREATE TABLE IF NOT EXISTS banners (
    id          UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    image_url   VARCHAR(255) NOT NULL,
    link_url    VARCHAR(255),
    position    VARCHAR(50)  NOT NULL,
    priority    INTEGER DEFAULT 0,
    is_active   BOOLEAN DEFAULT TRUE,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_banners_position ON banners (position);