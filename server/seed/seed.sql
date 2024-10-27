INSERT INTO priorities (id, name, created_at, updated_at)
VALUES 
    (1, 'low', NOW(), NOW()),
    (2, 'middle', NOW(), NOW()),
    (3, 'high', NOW(), NOW())
ON CONFLICT (id) 
DO UPDATE SET 
    name = EXCLUDED.name,
    created_at = EXCLUDED.created_at,
    updated_at = EXCLUDED.updated_at;
SELECT pg_catalog.setval('public.priorities_id_seq', (select max(id) from priorities), true);

INSERT INTO todos (id, title, description, priority_id, created_at, updated_at)
VALUES
    (1, 'Sample Title 1', 'Sample description 1', 1, NOW(), NOW()),
    (2, 'Sample Title 2', 'Sample description 2', 2, NOW(), NOW())
ON CONFLICT (id)
DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    priority_id = EXCLUDED.priority_id,
    created_at = EXCLUDED.created_at,
    updated_at = EXCLUDED.updated_at;
SELECT pg_catalog.setval('public.todos_id_seq', (select max(id) from todos), true);

INSERT INTO status (id, value, created_at, updated_at)
VALUES 
    (1, '未着手', NOW(), NOW()),
    (2, '着手', NOW(), NOW()),
    (3, '完了', NOW(), NOW())
ON CONFLICT (id)
DO UPDATE SET
    value = EXCLUDED.value,
    created_at = EXCLUDED.created_at,
    updated_at = EXCLUDED.updated_at;
SELECT pg_catalog.setval('public.status_id_seq', (select max(id) from status), true);

INSERT INTO labels (id, value, created_at, updated_at)
VALUES 
    (1, 'Label1', NOW(), NOW()),
    (2, 'Label2', NOW(), NOW())
ON CONFLICT (id)
DO UPDATE SET
    value = EXCLUDED.value,
    created_at = EXCLUDED.created_at,
    updated_at = EXCLUDED.updated_at;
SELECT pg_catalog.setval('public.labels_id_seq', (select max(id) from labels), true);

INSERT INTO label_todos (label_id, todo_id)
VALUES
    (1, 1),
    (2, 1),
    (2, 2)
ON CONFLICT (label_id, todo_id)
DO UPDATE SET
    label_id = EXCLUDED.label_id,
    todo_id = EXCLUDED.todo_id;
    
