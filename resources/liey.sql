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

INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `sort_num`, `keepalive`, `is_default`, `title`, `icon`, `requires_auth`, `deleted_at`) VALUES (1, '2022-05-21 10:42:32.000', '2022-05-29 13:40:13.064', 0, '/system', 'system', 0, 2, 1, 0, '系统管理', 'icon-park-outline:system', 0, NULL);
INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `sort_num`, `keepalive`, `is_default`, `title`, `icon`, `requires_auth`, `deleted_at`) VALUES (2, '2022-05-21 11:13:00.000', '2022-05-23 23:57:08.313', 1, '/system/menu', 'system_menu', 0, 1, 1, 1, '菜单管理', 'carbon:layers', 0, NULL);
INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `sort_num`, `keepalive`, `is_default`, `title`, `icon`, `requires_auth`, `deleted_at`) VALUES (3, '2022-05-22 18:20:28.000', '2022-05-23 23:45:34.827', 0, '/dashboard', 'dashboard', 0, 11311, 1, 0, '仪表盘', 'carbon:logo-google', 0, NULL);
INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `sort_num`, `keepalive`, `is_default`, `title`, `icon`, `requires_auth`, `deleted_at`) VALUES (4, '2022-05-22 18:22:21.000', '2022-05-23 23:45:58.017', 3, '/dashboard/analysis', 'dashboard_analysis', 0, 500, 1, 0, '分析页', 'carbon:download', 0, NULL);
INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `sort_num`, `keepalive`, `is_default`, `title`, `icon`, `requires_auth`, `deleted_at`) VALUES (5, '2022-05-22 18:22:54.000', '2022-05-23 23:46:08.282', 3, '/dashboard/workbench', 'dashboard_workbench', 0, 641, 1, 0, '工作台', 'carbon:cube', 0, NULL);
INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `sort_num`, `keepalive`, `is_default`, `title`, `icon`, `requires_auth`, `deleted_at`) VALUES (6, '2022-05-23 23:55:30.000', '2022-05-29 11:52:35.463', 1, '/system/admin', 'system_admin', 0, 98, 1, 0, '管理员', 'ion:logo-octocat', 1, NULL);
INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `sort_num`, `keepalive`, `is_default`, `title`, `icon`, `requires_auth`, `deleted_at`) VALUES (7, '2022-05-25 18:24:50.544', '2022-05-25 18:24:50.544', 0, '/component', 'component', 0, 1, 0, 0, '组件', 'carbon:application-web', 1, NULL);
INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `sort_num`, `keepalive`, `is_default`, `title`, `icon`, `requires_auth`, `deleted_at`) VALUES (8, '2022-05-25 18:33:59.670', '2022-05-25 18:33:59.670', 7, '/component/form', 'component_form', 0, 1, 0, 0, 'Form', 'fluent:form-48-regular', 1, NULL);
INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `sort_num`, `keepalive`, `is_default`, `title`, `icon`, `requires_auth`, `deleted_at`) VALUES (9, '2022-05-25 18:39:56.561', '2022-05-25 18:39:56.561', 8, '/component/form/cropper', 'component_form_cropper', 0, 1, 1, 0, '图片裁剪', 'gg:image', 1, NULL);
INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `sort_num`, `keepalive`, `is_default`, `title`, `icon`, `requires_auth`, `deleted_at`) VALUES (10, '2022-05-29 11:35:20.000', '2022-05-29 13:37:12.119', 1, '/system/role', 'system_role', 0, 72, 1, 0, '角色管理', 'la:users-cog', 1, NULL);
INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `sort_num`, `keepalive`, `is_default`, `title`, `icon`, `requires_auth`, `deleted_at`) VALUES (11, '2022-05-29 13:36:49.000', '2022-05-29 13:40:29.670', 1, '/system/permission', 'system_permission', 0, 70, 1, 0, '权限管理', 'octicon:shield-lock-16', 1, NULL);
INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `sort_num`, `keepalive`, `is_default`, `title`, `icon`, `requires_auth`, `deleted_at`) VALUES (15, '2022-05-30 23:16:01.000', '2022-05-30 23:22:43.255', 0, '/tools', 'tools', 0, 1, 1, 0, '开发工具', 'fa-solid:tools', 1, NULL);
INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `sort_num`, `keepalive`, `is_default`, `title`, `icon`, `requires_auth`, `deleted_at`) VALUES (16, '2022-05-30 23:21:20.920', '2022-05-30 23:21:20.920', 15, '/tools/scaffold', 'tools_scaffold', 0, 1, 1, 0, '脚手架', 'fa-solid:toolbox', 1, NULL);
INSERT INTO `liey`.`menus` (`id`, `created_at`, `updated_at`, `parent_id`, `path`, `name`, `hidden`, `sort_num`, `keepalive`, `is_default`, `title`, `icon`, `requires_auth`, `deleted_at`) VALUES (17, '2022-05-31 00:33:41.304', '2022-05-31 00:33:41.304', 1, '/system/role/:id', 'system_role_detail', 0, 1, 1, 0, '角色详情', NULL, 1, NULL);