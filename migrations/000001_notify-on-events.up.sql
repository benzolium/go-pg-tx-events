begin;

create table if not exists _events(
    id bigserial PRIMARY KEY,
    created_at timestamp with time zone NOT NULL default now(),
    object_id text not null,
    object_routing_key text not null,
    data json not null
);

create index _events_rk_id on _events(object_routing_key, object_id);

CREATE OR REPLACE FUNCTION events_notify() RETURNS trigger AS $$
BEGIN
  PERFORM pg_notify('events', row_to_json(NEW.*)::text);
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER events_notify_insert AFTER INSERT ON _events FOR EACH ROW EXECUTE PROCEDURE events_notify();

commit;
