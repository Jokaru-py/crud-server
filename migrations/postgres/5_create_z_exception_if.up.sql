CREATE OR REPLACE FUNCTION public.z_exception_if(
    arg_condition boolean,
    arg_message text,
    VARIADIC arg_params text[] DEFAULT ARRAY[]::text[]
)
RETURNS void AS
$BODY$
begin
    if arg_condition then
        raise exception 'USER_ERROR %', format(arg_message, variadic arg_params);
    end if;
end;
$BODY$
  LANGUAGE plpgsql IMMUTABLE
  COST 100;

ALTER FUNCTION public.z_exception_if(boolean, text, text[]) OWNER TO postgres;
GRANT EXECUTE ON FUNCTION public.z_exception_if(boolean, text, text[]) TO postgres;
GRANT EXECUTE ON FUNCTION public.z_exception_if(boolean, text, text[]) TO public;
COMMENT ON FUNCTION public.z_exception_if(boolean, text, text[]) IS 'Выбрасывает исключение по условию';