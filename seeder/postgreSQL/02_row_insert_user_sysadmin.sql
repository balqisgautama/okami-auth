-- +migrate Up
-- +migrate StatementBegin

-- password hash sha256 dari @Sysadmin37
INSERT INTO users (user_id, email, password, client_id, status, created_client, sysadmin)
VALUES (1, 'info.okami.project@gmail.com',
        '6fe88193540bbe2b9113b349b0eacbc50938ed19943696c9d568d68aa4ee55d5',
        '3be5760b-c160-4efc-b450-1974cd0b8788', 2,
        '3be5760b-c160-4efc-b450-1974cd0b8788', 1);

-- +migrate StatementEnd
-- +migrate Down
