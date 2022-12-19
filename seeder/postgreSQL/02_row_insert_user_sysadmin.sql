-- +migrate Up
-- +migrate StatementBegin

-- password hash sha256 dari @Sysadmin37
INSERT INTO users (user_id, email, password, client_id, status, created_client)
VALUES (1, 'info.okami.project@gmail.com',
        '6fe88193540bbe2b9113b349b0eacbc50938ed19943696c9d568d68aa4ee55d5',
        '7cfb9e47614c4fe685e6e13dada2828b', 2,
        '7cfb9e47614c4fe685e6e13dada2828b');

-- +migrate StatementEnd
-- +migrate Down
