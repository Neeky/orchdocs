
-- topology:
--
-- 22293
-- + 22294
--   + 22295
-- + 22296
-- + 22297
--
DELETE FROM database_instance_maintenance;
DELETE FROM database_instance_downtime;
DELETE FROM database_instance;
DELETE FROM database_instance_tags;
DELETE FROM database_instance_stale_binlog_coordinates;

INSERT INTO database_instance (hostname, port, last_checked, last_attempted_check, last_seen, uptime, server_id, server_uuid, version, binlog_server, read_only, binlog_format, log_bin, log_slave_updates, binary_log_file, binary_log_pos, master_host, master_port, slave_sql_running, slave_io_running, has_replication_filters, oracle_gtid, master_uuid, ancestry_uuid, executed_gtid_set, gtid_purged, gtid_errant, supports_oracle_gtid, mariadb_gtid, pseudo_gtid, master_log_file, read_master_log_pos, relay_master_log_file, exec_master_log_pos, relay_log_file, relay_log_pos, last_sql_error, last_io_error, seconds_behind_master, slave_lag_seconds, sql_delay, allow_tls, num_slave_hosts, slave_hosts, cluster_name, suggested_cluster_alias, data_center, region, physical_environment, instance_alias, semi_sync_enforced, replication_depth, is_co_master, replication_credentials_available, has_replication_credentials, version_comment, major_version, last_check_partial_success, binlog_row_image, last_discovery_latency, semi_sync_master_enabled, semi_sync_replica_enabled, gtid_mode, replication_group_members) VALUES ('testhost',22293,'2017-02-02 08:29:57','2017-02-02 08:29:57','2017-02-02 08:29:57',670447,1,'00022293-1111-1111-1111-111111111111','5.6.28-log',0,0,'STATEMENT',1,1,'mysql-bin.000167',137086726,'',0,0,0,0,0,'','','','','',0,0,1,'',0,'',0,'',0,'','',NULL,NULL,0,0,3,'[{\"Hostname\":\"testhost\",\"Port\":22294},{\"Hostname\":\"testhost\",\"Port\":22297},{\"Hostname\":\"testhost\",\"Port\":22296}]','testhost:22293','testhost','ny','us-region','','',0,0,0,0,0,'MySQL Community Server (GPL)','5.6',0,'full',0,0,0,'OFF', '[]');
INSERT INTO database_instance (hostname, port, last_checked, last_attempted_check, last_seen, uptime, server_id, server_uuid, version, binlog_server, read_only, binlog_format, log_bin, log_slave_updates, binary_log_file, binary_log_pos, master_host, master_port, slave_sql_running, slave_io_running, has_replication_filters, oracle_gtid, master_uuid, ancestry_uuid, executed_gtid_set, gtid_purged, gtid_errant, supports_oracle_gtid, mariadb_gtid, pseudo_gtid, master_log_file, read_master_log_pos, relay_master_log_file, exec_master_log_pos, relay_log_file, relay_log_pos, last_sql_error, last_io_error, seconds_behind_master, slave_lag_seconds, sql_delay, allow_tls, num_slave_hosts, slave_hosts, cluster_name, suggested_cluster_alias, data_center, region, physical_environment, instance_alias, semi_sync_enforced, replication_depth, is_co_master, replication_credentials_available, has_replication_credentials, version_comment, major_version, last_check_partial_success, binlog_row_image, last_discovery_latency, semi_sync_master_enabled, semi_sync_replica_enabled, gtid_mode, replication_group_members) VALUES ('testhost',22294,'2017-02-02 08:29:57','2017-02-02 08:29:57','2017-02-02 08:29:57',72466,101,'9a996060-6b8f-11e6-903f-6191b3cde928','5.6.28-log',0,0,'STATEMENT',1,1,'mysql-bin.000132',15931747,'testhost',22293,1,1,0,0,'','','','','',0,0,1,'mysql-bin.000167',137086726,'mysql-bin.000167',137086726,'mysql-relay.000029',15978254,'\"\"','\"\"',0,0,0,0,1,'[{\"Hostname\":\"testhost\",\"Port\":22295}]','testhost:22293','','ny','us-region','','',0,1,0,0,1,'MySQL Community Server (GPL)','5.6',0,'full',0,0,0,'OFF', '[]');
INSERT INTO database_instance (hostname, port, last_checked, last_attempted_check, last_seen, uptime, server_id, server_uuid, version, binlog_server, read_only, binlog_format, log_bin, log_slave_updates, binary_log_file, binary_log_pos, master_host, master_port, slave_sql_running, slave_io_running, has_replication_filters, oracle_gtid, master_uuid, ancestry_uuid, executed_gtid_set, gtid_purged, gtid_errant, supports_oracle_gtid, mariadb_gtid, pseudo_gtid, master_log_file, read_master_log_pos, relay_master_log_file, exec_master_log_pos, relay_log_file, relay_log_pos, last_sql_error, last_io_error, seconds_behind_master, slave_lag_seconds, sql_delay, allow_tls, num_slave_hosts, slave_hosts, cluster_name, suggested_cluster_alias, data_center, region, physical_environment, instance_alias, semi_sync_enforced, replication_depth, is_co_master, replication_credentials_available, has_replication_credentials, version_comment, major_version, last_check_partial_success, binlog_row_image, last_discovery_latency, semi_sync_master_enabled, semi_sync_replica_enabled, gtid_mode, replication_group_members) VALUES ('testhost',22295,'2017-02-02 08:29:57','2017-02-02 08:29:57','2017-02-02 08:29:57',670442,102,'9dc85926-6b8f-11e6-903f-85211507e568','5.6.28-log',0,0,'STATEMENT',1,1,'mysql-bin.000129',136688950,'testhost',22294,1,1,0,0,'','','','','',0,0,1,'mysql-bin.000132',15931747,'mysql-bin.000132',15931747,'mysql-relay.000002',15868528,'\"\"','\"\"',0,0,0,0,0,'[]','testhost:22293','','ny','us-region','','',0,2,0,0,1,'MySQL Community Server (GPL)','5.6',0,'full',0,0,0,'OFF', '[]');
INSERT INTO database_instance (hostname, port, last_checked, last_attempted_check, last_seen, uptime, server_id, server_uuid, version, binlog_server, read_only, binlog_format, log_bin, log_slave_updates, binary_log_file, binary_log_pos, master_host, master_port, slave_sql_running, slave_io_running, has_replication_filters, oracle_gtid, master_uuid, ancestry_uuid, executed_gtid_set, gtid_purged, gtid_errant, supports_oracle_gtid, mariadb_gtid, pseudo_gtid, master_log_file, read_master_log_pos, relay_master_log_file, exec_master_log_pos, relay_log_file, relay_log_pos, last_sql_error, last_io_error, seconds_behind_master, slave_lag_seconds, sql_delay, allow_tls, num_slave_hosts, slave_hosts, cluster_name, suggested_cluster_alias, data_center, region, physical_environment, instance_alias, semi_sync_enforced, replication_depth, is_co_master, replication_credentials_available, has_replication_credentials, version_comment, major_version, last_check_partial_success, binlog_row_image, last_discovery_latency, semi_sync_master_enabled, semi_sync_replica_enabled, gtid_mode, replication_group_members) VALUES ('testhost',22296,'2017-02-02 08:29:57','2017-02-02 08:29:57','2017-02-02 08:29:57',670438,103,'00022296-4444-4444-4444-444444444444','5.6.28',0,0,'STATEMENT',0,1,'',0,'testhost',22293,1,1,0,0,'','','','','',0,0,1,'mysql-bin.000167',137086726,'mysql-bin.000167',137086726,'mysql-relay.000052',137086889,'\"\"','\"\"',0,0,0,0,0,'[]','testhost:22293','','ny','us-region','','',0,1,0,0,1,'MySQL Community Server (GPL)','5.6',0,'full',0,0,0,'OFF', '[]');
INSERT INTO database_instance (hostname, port, last_checked, last_attempted_check, last_seen, uptime, server_id, server_uuid, version, binlog_server, read_only, binlog_format, log_bin, log_slave_updates, binary_log_file, binary_log_pos, master_host, master_port, slave_sql_running, slave_io_running, has_replication_filters, oracle_gtid, master_uuid, ancestry_uuid, executed_gtid_set, gtid_purged, gtid_errant, supports_oracle_gtid, mariadb_gtid, pseudo_gtid, master_log_file, read_master_log_pos, relay_master_log_file, exec_master_log_pos, relay_log_file, relay_log_pos, last_sql_error, last_io_error, seconds_behind_master, slave_lag_seconds, sql_delay, allow_tls, num_slave_hosts, slave_hosts, cluster_name, suggested_cluster_alias, data_center, region, physical_environment, instance_alias, semi_sync_enforced, replication_depth, is_co_master, replication_credentials_available, has_replication_credentials, version_comment, major_version, last_check_partial_success, binlog_row_image, last_discovery_latency, semi_sync_master_enabled, semi_sync_replica_enabled, gtid_mode, replication_group_members) VALUES ('testhost',22297,'2017-02-02 08:29:57','2017-02-02 08:29:57','2017-02-02 08:29:57',670435,104,'00022297-5555-5555-5555-555555555555','5.6.28-log',0,0,'STATEMENT',1,1,'mysql-bin.001013',106518535,'testhost',22293,1,1,0,0,'','','','','',0,0,1,'mysql-bin.000167',137086726,'mysql-bin.000167',137086726,'mysql-relay.000041',86111934,'\"\"','\"\"',0,0,0,0,0,'[]','testhost:22293','','seattle','us-region','','',0,1,0,0,1,'MySQL Community Server (GPL)','5.6',0,'full',0,0,0,'OFF', '[]');

UPDATE database_instance SET replication_sql_thread_state=1 WHERE slave_sql_running=1;
UPDATE database_instance SET replication_io_thread_state=1  WHERE slave_io_running=1;
UPDATE database_instance SET replication_sql_thread_state=-1, replication_io_thread_state=-1 WHERE port=22293;
DELETE FROM candidate_database_instance;