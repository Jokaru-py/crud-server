CREATE TABLE public.t_users (
	id      serial      NOT NULL,
	name  	varchar     NOT NULL,
	age  	varchar     NOT NULL,
	email 	varchar   	NOT NULL
);

ALTER TABLE public.t_users OWNER TO postgres;
GRANT ALL ON TABLE public.t_users TO postgres;
REVOKE ALL ON TABLE public.t_users FROM public;

COMMENT ON COLUMN public.t_users."id" IS 'Идентификатор';
COMMENT ON COLUMN public.t_users."name" IS 'Полное ФИО пользователя';
COMMENT ON COLUMN public.t_users."age" IS 'Возраст пользователя';
COMMENT ON COLUMN public.t_users.email IS 'Эл. адрес';

