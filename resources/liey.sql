-- ----------------------------
-- Table structure for admins
-- ----------------------------
CREATE TABLE `admins` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `email` varchar(64) NOT NULL COMMENT '邮箱',
  `password` longtext NOT NULL COMMENT '密码',
  `avatar` varchar(191) NOT NULL DEFAULT '' COMMENT '头像',
  `is_active` tinyint NOT NULL DEFAULT '1' COMMENT '状态 0：禁用 1：启用',
  `register_at` datetime(3) DEFAULT NULL COMMENT '用户注册时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB;

-- ----------------------------
-- Table structure for menus
-- ----------------------------
CREATE TABLE `menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `parent_id` bigint unsigned NOT NULL COMMENT '父菜单ID',
  `path` longtext NOT NULL COMMENT '路由path',
  `title` longtext NOT NULL COMMENT '标题',
  `name` longtext NOT NULL COMMENT '路由name',
  `icon` longtext NOT NULL COMMENT '图标',
  `requires_auth` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否需要权限',
  `hidden` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否在菜单隐藏',
  `sort_num` bigint NOT NULL DEFAULT '1' COMMENT '排序标记',
  `keepalive` tinyint(1) NOT NULL DEFAULT '1' COMMENT '缓存',
  `is_default` tinyint(1) NOT NULL DEFAULT '0' COMMENT '默认菜单',
  PRIMARY KEY (`id`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB;

-- ----------------------------
-- Table structure for permissions
-- ----------------------------
CREATE TABLE `permissions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `name` varchar(64) NOT NULL COMMENT '名称',
  `slug` varchar(64) NOT NULL COMMENT '标识',
  `rules` longtext NOT NULL COMMENT '权限规则',
  `parent_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '父级权限',
  `sort_num` bigint NOT NULL DEFAULT '0' COMMENT '排序值',
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB;

-- ----------------------------
-- Table structure for role_has_permissions
-- ----------------------------
CREATE TABLE `role_has_permissions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `permission_id` bigint unsigned NOT NULL COMMENT '权限ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_role_permission` (`role_id`,`permission_id`)
) ENGINE=InnoDB;

-- ----------------------------
-- Table structure for roles
-- ----------------------------
CREATE TABLE `roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `name` varchar(64) NOT NULL COMMENT '角色名称',
  `slug` varchar(32) NOT NULL COMMENT '角色标识',
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB;

-- ----------------------------
-- Table structure for users
-- ----------------------------
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `username` longtext,
  `password` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB;

INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `title`, `name`, `icon`, `requires_auth`, `hidden`, `sort_num`, `keepalive`, `is_default`) VALUES (1, '2022-06-10 09:26:12.000', '2022-06-10 09:26:15.000', 0, '/dashboard', '控制台', 'dashboard', 'carbon:dashboard', 1, 0, 1, 1, 0);
INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `title`, `name`, `icon`, `requires_auth`, `hidden`, `sort_num`, `keepalive`, `is_default`) VALUES (2, '2022-06-10 09:27:44.000', '2022-06-10 09:27:47.000', 1, '/dashboard/analysis', '分析页', 'dashboard_analysis', 'icon-park-outline:analysis', 1, 0, 1, 1, 0);
INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `title`, `name`, `icon`, `requires_auth`, `hidden`, `sort_num`, `keepalive`, `is_default`) VALUES (3, '2022-06-10 09:28:19.000', '2022-06-10 09:28:21.000', 1, '/dashboard/workbench', '工作台', 'dashboard_workbench', 'icon-park-outline:workbench', 1, 0, 1, 1, 0);
INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `title`, `name`, `icon`, `requires_auth`, `hidden`, `sort_num`, `keepalive`, `is_default`) VALUES (4, '2022-06-10 09:29:50.000', '2022-06-10 09:29:51.000', 0, '/system', '系统管理', 'system', 'icon-park-outline:system', 1, 0, 1, 1, 0);
INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `title`, `name`, `icon`, `requires_auth`, `hidden`, `sort_num`, `keepalive`, `is_default`) VALUES (5, '2022-06-10 09:30:34.000', '2022-06-10 09:30:36.000', 4, '/system/menu', '菜单管理', 'system_menu', 'bx:food-menu', 1, 0, 1, 1, 0);
