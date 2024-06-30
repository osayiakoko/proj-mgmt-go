-- Create index "tasks_priority_idx" to table: "tasks"
CREATE INDEX "tasks_priority_idx" ON "tasks" USING hash ("priority");
-- Create index "tasks_status_idx" to table: "tasks"
CREATE INDEX "tasks_status_idx" ON "tasks" USING hash ("status");
-- Create index "tasks_title_idx" to table: "tasks"
CREATE INDEX "tasks_title_idx" ON "tasks" USING gin ((to_tsvector('simple'::regconfig, title)));
