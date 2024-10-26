INSERT INTO priorities (id, name)
VALUES 
    (1, 'low'),
    (2, 'middle'),
    (3, 'high')
ON CONFLICT (id) 
DO UPDATE SET 
    name = EXCLUDED.name;

INSERT INTO todos (id, title, description, name, priority_id)
VALUES
    (1, 'Sample Title 1', 'Sample description 1', 'Name 1', 1),
    (2, 'Sample Title 2', 'Sample description 2', 'Name 2', 2)
ON CONFLICT (id)
DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    name = EXCLUDED.name,
    priority_id = EXCLUDED.priority_id;

INSERT INTO status (id, value)
VALUES 
    (1, '未着手'),
    (2, '着手'),
    (3, '完了')
ON CONFLICT (id)
DO UPDATE SET
    value = EXCLUDED.value;
