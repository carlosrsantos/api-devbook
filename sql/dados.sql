insert into usuarios(nome, nick, email, senha) 
values
(1,"Usuario 1", "usr1","usuario1@mail.com", "$2a$10$xwjcz.8xr3wjz/WpLD8LgO.x4Eql25GHnuTREQioTy9SfELDurWZG"),
(2,"Usuario 2", "usr2","usuario2@mail.com", "$2a$10$HzuY2wicoE7NamZLC5LpR.nvYKC2o3/mLP/lcNnUJoxXveR1Jk4cC"),
(3,"Usuario 3", "usr3","usuario3@mail.com", "$2a$10$XtO8Httu5dH0uKQyFD9TZO2hNSYB9daARauYui2Pv1Gu4rgztlgb.");

insert into seguidores(usuario_id, seguidor_id) 
values
(1, 2),
(3, 1),
(1, 3);
