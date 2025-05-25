
```
SELECT id, version, name, endpoint_ip, endpoint_ipv6, public_key, listen_port, mtu, persistent_keepalive, is_current, last_modified, created_at, data
	FROM public.hosts  WHERE is_current = true ORDER BY name ASC;
```
SELECT * FROM networks WHERE is_current = true ORDER BY last_modified DESC;

SELECT * FROM nodes WHERE is_current = true ORDER BY last_modified DESC;

SELECT * FROM ext_clients WHERE is_current = true ORDER BY last_modified DESC;

SELECT * FROM dns WHERE is_current = true ORDER BY last_modified DESC;

SELECT * FROM acls WHERE is_current = true ORDER BY last_modified DESC;

SELECT * FROM sync_history ORDER BY created_at DESC LIMIT 10;


SELECT 'networks' as table_name, COUNT(*) as total_records, COUNT(*) FILTER (WHERE is_current = true) as current_records FROM networks
UNION ALL
SELECT 'nodes', COUNT(*), COUNT(*) FILTER (WHERE is_current = true) FROM nodes
UNION ALL
SELECT 'acls', COUNT(*), COUNT(*) FILTER (WHERE is_current = true) FROM acls
UNION ALL
SELECT 'hosts', COUNT(*), COUNT(*) FILTER (WHERE is_current = true) FROM hosts
UNION ALL
SELECT 'ext_clients', COUNT(*), COUNT(*) FILTER (WHERE is_current = true) FROM ext_clients;


