-- ========================
-- DROP TRIGGERS
-- ========================
DROP TRIGGER IF EXISTS set_users_updated_at ON users;
DROP TRIGGER IF EXISTS set_user_profiles_updated_at ON user_profiles;
DROP TRIGGER IF EXISTS set_roles_updated_at ON roles;
DROP TRIGGER IF EXISTS set_permissions_updated_at ON permissions;
DROP TRIGGER IF EXISTS set_user_providers_updated_at ON user_providers;

-- ========================
-- DROP TRIGGER FUNCTION
-- ========================
DROP FUNCTION IF EXISTS update_updated_at_column;

-- ========================
-- DROP TABLES (reverse order of dependencies)
-- ========================
DROP TABLE IF EXISTS user_providers;
DROP TABLE IF EXISTS user_roles;
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS user_profiles;
DROP TABLE IF EXISTS users;

-- ========================
-- DROP EXTENSIONS
-- ========================
DROP EXTENSION IF EXISTS "uuid-ossp";
