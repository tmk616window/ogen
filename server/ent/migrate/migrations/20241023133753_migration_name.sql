-- Create "priorities" table
CREATE TABLE "priorities" ("id" character varying NOT NULL, "name" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "priority_todos" table
CREATE TABLE "priority_todos" ("priority_id" character varying NOT NULL, "todo_id" character varying NOT NULL, PRIMARY KEY ("priority_id", "todo_id"), CONSTRAINT "priority_todos_priority_id" FOREIGN KEY ("priority_id") REFERENCES "priorities" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "priority_todos_todo_id" FOREIGN KEY ("todo_id") REFERENCES "todos" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
