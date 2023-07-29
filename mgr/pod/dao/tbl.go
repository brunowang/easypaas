package dao

const podTable = `
CREATE TABLE pod (
  id bigint NOT NULL AUTO_INCREMENT,
  pod_name varchar(255) DEFAULT NULL,
  pod_namespace varchar(255) DEFAULT NULL,
  pod_team_id varchar(255) DEFAULT NULL,
  pod_cpu_min double DEFAULT NULL,
  pod_cpu_max double DEFAULT NULL,
  pod_replicas int DEFAULT NULL,
  pod_memory_min double DEFAULT NULL,
  pod_memory_max double DEFAULT NULL,
  pod_pull_policy varchar(255) DEFAULT NULL,
  pod_restart varchar(255) DEFAULT NULL,
  pod_type varchar(255) DEFAULT NULL,
  pod_image varchar(255) DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY uix_pod_pod_name (pod_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
`

const podEnvTable = `
CREATE TABLE pod_env (
  id bigint NOT NULL AUTO_INCREMENT,
  pod_id bigint DEFAULT NULL,
  env_key varchar(255) DEFAULT NULL,
  env_value varchar(255) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
`

const podPortTable = `
CREATE TABLE pod_port (
  id bigint NOT NULL AUTO_INCREMENT,
  pod_id bigint DEFAULT NULL,
  container_port int DEFAULT NULL,
  protocol varchar(255) DEFAULT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
`
