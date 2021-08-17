CREATE OR REPLACE FUNCTION public.fn_user_del(arg_id text)
RETURNS void AS
$BODY$
declare
	v_expire_time timestamp;
    v_path text;
begin
    perform z_exception_if(arg_id is null or arg_id = '', 'Не передан ID пользователя для удаления');

	delete from public.t_users t
    where t.id = arg_id::int;
end;
$BODY$
    LANGUAGE plpgsql VOLATILE
	COST 100;

ALTER FUNCTION public.fn_user_del(text) OWNER TO postgres;
GRANT EXECUTE ON FUNCTION public.fn_user_del(text) TO postgres;
GRANT EXECUTE ON FUNCTION public.fn_user_del(text) TO public;
COMMENT ON FUNCTION public.fn_user_del(text) IS 'Удаление пользователя по ID';
