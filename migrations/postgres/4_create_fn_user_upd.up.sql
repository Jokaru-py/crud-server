CREATE OR REPLACE FUNCTION public.fn_user_upd(arg_name text, arg_age text, arg_email text, arg_user_id text)
RETURNS TABLE(name character varying, age character varying, email character varying) AS
$BODY$
begin
	perform z_exception_if(arg_name is null or arg_name = '', 'Не передано ФИО');
	perform z_exception_if(arg_age is null or arg_age = '', 'Не передан возраст пользователя');
	perform z_exception_if(arg_email is null or arg_email = '', 'Не передан email пользователя');
	perform z_exception_if(arg_user_id is null or arg_user_id = '', 'Не передан email пользователя');
	
	-- Обновление записи
	update public.t_users t set
		name = arg_name,
		age = arg_age,
		email = arg_email
	where t.id = arg_user_id::int;
	
	if not found then
		raise exception 'Пользователь с таким ID не найден';
	end if;

	return query
    select
        t.name,
        t.age,
        t.email
    from public.t_users t
    where t.id = arg_user_id::int;	
end;
$BODY$
    LANGUAGE plpgsql VOLATILE
	COST 100;

ALTER FUNCTION public.fn_user_upd(text, text, text, text) OWNER TO postgres;
GRANT EXECUTE ON FUNCTION public.fn_user_upd(text, text, text, text) TO postgres;
GRANT EXECUTE ON FUNCTION public.fn_user_upd(text, text, text, text) TO public;
COMMENT ON FUNCTION public.fn_user_upd(text, text, text, text) IS 'Обновление данных пользователя';