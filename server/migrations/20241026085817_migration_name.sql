-- Modify "priorities" table
ALTER TABLE "priorities" ADD COLUMN "created_at" timestamptz NOT NULL DEFAULT NOW(), ADD COLUMN "updated_at" timestamptz NOT NULL DEFAULT NOW();
-- Modify "status" table
ALTER TABLE "status" ADD COLUMN "created_at" timestamptz NOT NULL DEFAULT NOW(), ADD COLUMN "updated_at" timestamptz NOT NULL DEFAULT NOW();
-- Modify "todos" table
ALTER TABLE "todos" ADD COLUMN "created_at" timestamptz NOT NULL DEFAULT NOW(), ADD COLUMN "updated_at" timestamptz NOT NULL DEFAULT NOW();
