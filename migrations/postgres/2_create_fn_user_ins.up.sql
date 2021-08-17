CREATE OR REPLACE FUNCTION public.fn_user_ins(arg_name text, arg_age text, arg_email text)
RETURNS int AS
$BODY$
declare
	v_id int;
begin
	perform z_exception_if(arg_name is null or arg_name = '', 'Не передано ФИО');
	perform z_exception_if(arg_age is null or arg_age = '', 'Не передан возраст пользователя');
	perform z_exception_if(arg_email is null or arg_email = '', 'Не передан email пользователя');

	-- Добавление нового пользователя без проверки
	insert into t_users as t (
		name,
		age,
		email
	) values (
		arg_name,
		arg_age,
		arg_email
	) RETURNING id INTO v_id;
	
	return v_id;
end;
$BODY$
    LANGUAGE plpgsql VOLATILE
	COST 100;

ALTER FUNCTION public.fn_user_ins(text, text, text) OWNER TO postgres;
GRANT EXECUTE ON FUNCTION public.fn_user_ins(text, text, text) TO postgres;
GRANT EXECUTE ON FUNCTION public.fn_user_ins(text, text, text) TO public;
COMMENT ON FUNCTION public.fn_user_ins(text, text, text) IS 'Добавление нового пользователя';