INSERT INTO usuarios (nome, nick, email, senha)
VALUES
("Usuario 1", "usuario_1", "usuario1@gmail.com", "$2a$10$nqBrBQgv1O/E5NHJ09VMaO7WsOtZMY8Kc1RtL8nMWZobzbX8EXyty"),
("Usuario 2", "usuario_2", "usuario2@gmail.com", "$2a$10$nqBrBQgv1O/E5NHJ09VMaO7WsOtZMY8Kc1RtL8nMWZobzbX8EXyty"),
("Usuario 3", "usuario_3", "usuario3@gmail.com", "$2a$10$nqBrBQgv1O/E5NHJ09VMaO7WsOtZMY8Kc1RtL8nMWZobzbX8EXyty");

INSERT INTO seguidores (usuario_id, seguidor_id)
VALUES
(1,2),
(3,1),
(1,3);

INSERT INTO publicacoes(titulo, conteudo, autor_id)
VALUES
("Publicação do usuário 1", "Essa é a publicação do usuário 1! Oba!", 1);
("Publicação do usuário 2", "Essa é a publicação do usuário 2! Oba!", 2);
("Publicação do usuário 3", "Essa é a publicação do usuário 3! Oba!", 3);