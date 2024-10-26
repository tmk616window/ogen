-- Drop index "todos_priority_id_key" from table: "todos"
DROP INDEX "todos_priority_id_key";
-- Modify "todos" table
ALTER TABLE "todos" ADD COLUMN "status_id" bigint NOT NULL DEFAULT 1, ADD CONSTRAINT "todos_status_todo" FOREIGN KEY ("status_id") REFERENCES "status" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
