/*
Navicat MySQL Data Transfer

Source Server         : 禅道库
Source Server Version : 50622
Source Host           : 10.44.200.24:3306
Source Database       : cicdadmin

Target Server Type    : MYSQL
Target Server Version : 50622
File Encoding         : 65001

Date: 2020-09-07 15:21:40
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/calendar', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/calendar', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/calendar/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/calendar/:id', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/config', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/config', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/config/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/config/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/configKey/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/configList', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/dept', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/dept', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/dept/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/dept/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/deptList', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/deptTree', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/dict/data', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/dict/data/', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/dict/data/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/dict/data/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/dict/databytype/', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/dict/databytype/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/dict/datalist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/dict/type', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/dict/type', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/dict/type/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/dict/type/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/dict/typelist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/dict/typeoptionselect', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/doctor', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/doctor', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/doctor/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/doctor/:id', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/doctor/pic/', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/getinfo', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/loginlog/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/loginloglist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/logout', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/member', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/member', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/member/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/member/:id', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/menu', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/menu', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/menu/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/menu/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/menuids', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/menulist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/menurole', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/menuTreeselect', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/operloglist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/orders', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/orders', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/orders/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/orders/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/orders/:id', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/post', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/post', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/post/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/post/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/postlist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/role', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/role', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/role/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/role/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/roledatascope', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/roleDeptTreeselect/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/rolelist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/rolemenu', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/rolemenu', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/rolemenu/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/roleMenuTreeselect/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/user/avatar', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'Tester', '/api/v1/user/pwd', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menulist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menu', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/databytype/', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menu', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menu/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/sysUserList', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/sysUser/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/sysUser/', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/sysUser', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/sysUser', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/sysUser/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/user/profile', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/rolelist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/role/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/role', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/role', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/role/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/configList', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/config/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/configKey/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/config', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/config', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/config/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menurole', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/roleMenuTreeselect/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menuTreeselect', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/rolemenu', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/rolemenu', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/rolemenu/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/deptList', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dept/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dept', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dept', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dept/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/datalist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/data/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/databytype/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/data', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/data/', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/data/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/typelist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/type/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/type', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/type', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/type/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/postlist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/post/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/post', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/post', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/post/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menulist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menu/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menu', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menu/:id', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menu/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/menuids', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/loginloglist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/loginlog/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/operloglist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/getinfo', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/roledatascope', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/roleDeptTreeselect/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/deptTree', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/configKey/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/logout', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/user/avatar', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/user/pwd', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/dict/typeoptionselect', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releasecicd', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releasecicd', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releasecicd/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releasecicdList', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releasecicd/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releaseridList', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releaserid', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releaserid', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releaserid/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releaserid/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releasetodo', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releasetodo', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releasetodo/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releasetodoList', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releasetodo/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releasechart/deploybyday', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releasechart/deploybyhour', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/releasestgList', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menulist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menu', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dict/databytype/', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menu', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menu/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/sysUserList', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/sysUser/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/sysUser/', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/sysUser', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/sysUser', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/sysUser/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/user/profile', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/rolelist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/role/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/role', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/role', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/role/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/configKey/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menurole', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/roleMenuTreeselect/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menuTreeselect', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/rolemenu', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/rolemenu', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/rolemenu/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/deptList', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dept/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dept', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dept', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dept/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dict/databytype/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/postlist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/post/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/post', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/post', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/post/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menulist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menu/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menu', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menu/:id', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menu/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/menuids', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/loginloglist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/loginlog/:id', 'DELETE', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/operloglist', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/getinfo', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/roledatascope', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/roleDeptTreeselect/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/deptTree', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/configKey/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/logout', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/user/avatar', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/user/pwd', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/dict/typeoptionselect', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/releasecicdList', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/releasecicd/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/releaseridList', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/releaserid/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/releasetodoList', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/releasetodo/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'common', '/api/v1/releasestgList', 'GET', null, null, null);

-- ----------------------------
-- Table structure for releasecicd
-- ----------------------------
DROP TABLE IF EXISTS `releasecicd`;
CREATE TABLE `releasecicd` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `productname` varchar(250) NOT NULL DEFAULT '' COMMENT '产品线名',
  `rid` varchar(250) NOT NULL DEFAULT '' COMMENT '计划ID',
  `fid` varchar(25) NOT NULL DEFAULT '' COMMENT '需求ID',
  `bid` varchar(25) NOT NULL DEFAULT '' COMMENT 'bugID',
  `rnotes` text NOT NULL COMMENT 'Release Notes',
  `sure_status` enum('需求中','待开发排期','开发中','待测试排期','测试中','已上线','') NOT NULL DEFAULT '' COMMENT '认领状态',
  `sure_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '开始处理日期',
  `sure_reasion` varchar(500) NOT NULL DEFAULT '' COMMENT '认领原因',
  `appname` varchar(25) NOT NULL DEFAULT '' COMMENT '应用名',
  `civersion` text NOT NULL COMMENT 'CICD版本',
  `commitid` varchar(50) NOT NULL DEFAULT '' COMMENT 'CommitID',
  `prdversion` varchar(70) DEFAULT NULL,
  `openedby` varchar(25) NOT NULL DEFAULT '' COMMENT '产品发起人',
  `openeddate` datetime DEFAULT NULL,
  `cicdtime` datetime DEFAULT NULL,
  `num` int(3) NOT NULL DEFAULT '0' COMMENT '测试发布次数',
  `deploytime` datetime DEFAULT NULL,
  `fidstatus` varchar(25) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `assign_by` varchar(100) DEFAULT NULL COMMENT '指派人',
  `assign_to` varchar(100) DEFAULT NULL COMMENT '被指派人',
  `assign_time` timestamp NULL DEFAULT NULL COMMENT '指派时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3065 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of releasecicd
-- ----------------------------
INSERT INTO `releasecicd` VALUES ('1', '业务X', '业务_0710', '1136', '', 'xxx的业务类型', '', '2020-08-11 10:04:48', '', 'cTask', 'irr20200605.2348bdaf74425864fa58568a6fa9e6e571dc19ac.build.2020-07-10_18:15:35.1;irr20200605.3c520dc8d70e2a25095de8c59c99f1ac355c525a.build.2020-07-10_18:50:27.2', '3c520dc8d70e2a25095de8c59c99f1ac355c525a', null, '产品李某', '2020-07-10 16:29:36', '2020-07-10 18:00:00', '2', '2020-07-13 15:42:16', '已上线', null, '1', null, '2020-07-13 05:37:59', '2020-07-13 06:56:35', '21212', '1212121', '2020-08-01 15:41:38');
 
-- ----------------------------
-- Table structure for releasenotes
-- ----------------------------
DROP TABLE IF EXISTS `releasenotes`;
CREATE TABLE `releasenotes` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `rid` varchar(250) NOT NULL DEFAULT '' COMMENT 'release ID',
  `bid` varchar(25) NOT NULL DEFAULT '' COMMENT 'bugID',
  `appname` varchar(25) NOT NULL DEFAULT '' COMMENT '应用名',
  `civersion` text NOT NULL COMMENT 'CI版本',
  `rnotes` text NOT NULL COMMENT 'release notes',
  `developer` varchar(25) NOT NULL DEFAULT '' COMMENT '开发',
  `commitid` varchar(50) DEFAULT NULL,
  `status` varchar(5) DEFAULT NULL,
  `deploytime` date DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2215 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of releasenotes
-- ----------------------------
INSERT INTO `releasenotes` VALUES ('1', '会员系统_2.0.0.1', '', ' jjvccc', '\nmaster.cb5486954eaeb58935723cc184ab0b3fd5290093.build.2020-06-12_10:25:51.36', '\n1026/【会员系统】更新调整会员等级规则文案/jiaopengru\n1027/【会员系统】更改「普通会员」为「普通用户」/产品张某\n1028/【会员系统】权益重复领取的弹框文案更改为「该权益您已领取，请勿重复操作！」/产品张某', '', '', '已发布生产', '2020-07-09');
 

-- ----------------------------
-- Table structure for releasetodo
-- ----------------------------
DROP TABLE IF EXISTS `releasetodo`;
CREATE TABLE `releasetodo` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `rid` varchar(250) NOT NULL DEFAULT '' COMMENT '计划ID',
  `fid` varchar(25) NOT NULL DEFAULT '' COMMENT '需求ID',
  `bid` varchar(25) NOT NULL DEFAULT '' COMMENT 'bugID',
  `rnotes` text NOT NULL COMMENT 'release notes',
  `openedby` varchar(25) NOT NULL DEFAULT '' COMMENT '发起人',
  `openeddate` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '发起时间',
  `status` varchar(5) NOT NULL DEFAULT '' COMMENT '上线状态',
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=128 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of releasetodo
-- ----------------------------
INSERT INTO `releasetodo` VALUES ('1', 'XX功能_2.4.1.0_7.02', '1100', '', '好友规则', '令狐', '2020-07-01 16:28:51', '', null, null, null, null, null);
INSERT INTO `releasetodo` VALUES ('2', 'YY功能_1.1.4', '1101', '', '公告字数限制', '张倩倩', '2020-07-01 18:47:47', '', null, null, null, null, null);

-- ----------------------------
-- Table structure for sys_columns
-- ----------------------------
DROP TABLE IF EXISTS `sys_columns`;
CREATE TABLE `sys_columns` (
  `column_id` int(11) NOT NULL AUTO_INCREMENT,
  `table_id` int(11) DEFAULT NULL,
  `column_name` varchar(128) DEFAULT NULL,
  `column_comment` varchar(128) DEFAULT NULL,
  `column_type` varchar(128) DEFAULT NULL,
  `go_type` varchar(128) DEFAULT NULL,
  `go_field` varchar(128) DEFAULT NULL,
  `json_field` varchar(128) DEFAULT NULL,
  `is_pk` char(4) DEFAULT NULL,
  `is_increment` char(4) DEFAULT NULL,
  `is_required` char(4) DEFAULT NULL,
  `is_insert` char(4) DEFAULT NULL,
  `is_edit` char(4) DEFAULT NULL,
  `is_list` char(4) DEFAULT NULL,
  `is_query` char(4) DEFAULT NULL,
  `query_type` varchar(128) DEFAULT NULL,
  `html_type` varchar(128) DEFAULT NULL,
  `dict_type` varchar(128) DEFAULT NULL,
  `sort` int(4) DEFAULT NULL,
  `list` char(1) DEFAULT NULL,
  `pk` char(1) DEFAULT NULL,
  `required` char(1) DEFAULT NULL,
  `super_column` char(1) DEFAULT NULL,
  `usable_column` char(1) DEFAULT NULL,
  `increment` char(1) DEFAULT NULL,
  `insert` char(1) DEFAULT NULL,
  `edit` char(1) DEFAULT NULL,
  `query` char(1) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_By` varchar(128) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`column_id`),
  KEY `idx_sys_columns_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_columns
-- ----------------------------
INSERT INTO `sys_columns` VALUES ('1', '1', 'id', '自增ID', 'int(11) unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '0', '1', '', 'EQ', 'input', '', '1', '', '1', '1', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:39', '2020-07-22 18:13:08', null);
INSERT INTO `sys_columns` VALUES ('2', '1', 'rid', '计划ID', 'varchar(250)', 'string', 'Rid', 'rid', '0', '', '1', '1', '1', '1', '1', 'EQ', 'input', '', '2', '', '0', '1', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:39', '2020-07-22 18:13:09', null);
INSERT INTO `sys_columns` VALUES ('3', '1', 'fid', '需求ID', 'varchar(25)', 'string', 'Fid', 'fid', '0', '', '1', '1', '1', '1', '1', 'EQ', 'input', '', '3', '', '0', '1', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:39', '2020-07-22 18:13:09', null);
INSERT INTO `sys_columns` VALUES ('4', '1', 'bid', 'bugID', 'varchar(25)', 'string', 'Bid', 'bid', '0', '', '1', '1', '1', '1', '1', 'EQ', 'input', '', '4', '', '0', '1', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:39', '2020-07-22 18:13:09', null);
INSERT INTO `sys_columns` VALUES ('5', '1', 'rnotes', 'RNotes', 'text', 'string', 'Rnotes', 'rnotes', '0', '', '1', '1', '1', '1', '', 'EQ', 'input', '', '5', '', '0', '1', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:39', '2020-07-22 18:13:09', null);
INSERT INTO `sys_columns` VALUES ('6', '1', 'appname', '应用名', 'varchar(25)', 'string', 'Appname', 'appname', '0', '', '1', '1', '0', '1', '1', 'EQ', 'input', '', '6', '', '0', '1', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:39', '2020-07-22 18:13:09', null);
INSERT INTO `sys_columns` VALUES ('7', '1', 'civersion', 'STG版本', 'text', 'string', 'Civersion', 'civersion', '0', '', '1', '1', '', '1', '0', 'EQ', 'select', '', '7', '', '0', '1', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:39', '2020-07-22 18:13:09', null);
INSERT INTO `sys_columns` VALUES ('8', '1', 'commitid', 'CommitID', 'varchar(50)', 'string', 'Commitid', 'commitid', '0', '', '1', '1', '', '1', '1', 'EQ', 'input', '', '8', '', '0', '1', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:39', '2020-07-22 18:13:10', null);
INSERT INTO `sys_columns` VALUES ('9', '1', 'prdversion', 'PRD版本', 'varchar(70)', 'string', 'Prdversion', 'prdversion', '0', '', '1', '1', '', '1', '1', 'LIKE', 'input', '', '9', '', '0', '0', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:40', '2020-07-22 18:13:10', null);
INSERT INTO `sys_columns` VALUES ('10', '1', 'openedby', '发起人', 'varchar(25)', 'string', 'Openedby', 'openedby', '0', '', '1', '1', '', '1', '1', 'EQ', 'input', '', '10', '', '0', '1', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:40', '2020-07-22 18:13:10', null);
INSERT INTO `sys_columns` VALUES ('11', '1', 'openeddate', '发起时间', 'datetime', 'string', 'Openeddate', 'openeddate', '0', '', '1', '1', '', '1', '1', 'GTE', 'input', '', '11', '', '0', '0', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:40', '2020-07-22 18:13:10', null);
INSERT INTO `sys_columns` VALUES ('12', '1', 'cicdstatus', 'STG状态', 'char(5)', 'string', 'Cicdstatus', 'cicdstatus', '0', '', '1', '1', '', '1', '1', 'EQ', 'input', '', '12', '', '0', '1', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:40', '2020-07-22 18:13:10', null);
INSERT INTO `sys_columns` VALUES ('13', '1', 'cicdtime', '提测时间', 'datetime', 'string', 'Cicdtime', 'cicdtime', '0', '', '1', '1', '', '1', '1', 'GTE', 'input', '', '13', '', '0', '0', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:40', '2020-07-22 18:13:10', null);
INSERT INTO `sys_columns` VALUES ('14', '1', 'num', 'STG次数', 'int(3)', 'int', 'Num', 'num', '0', '', '1', '1', '', '1', '', 'EQ', 'input', '', '14', '', '0', '1', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:40', '2020-07-22 18:13:10', null);
INSERT INTO `sys_columns` VALUES ('15', '1', 'deploystatus', 'PRD状态', 'char(5)', 'string', 'Deploystatus', 'deploystatus', '0', '', '1', '1', '', '1', '1', 'EQ', 'input', '', '15', '', '0', '1', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:40', '2020-07-22 18:13:11', null);
INSERT INTO `sys_columns` VALUES ('16', '1', 'deploytime', 'PRD时间', 'datetime', 'string', 'Deploytime', 'deploytime', '0', '', '1', '1', '', '1', '1', 'EQ', 'input', '', '16', '', '0', '0', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:40', '2020-07-22 18:13:11', null);
INSERT INTO `sys_columns` VALUES ('17', '1', 'devtime', '开发耗时', 'varchar(25)', 'string', 'Devtime', 'devtime', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', '17', '', '0', '0', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:41', '2020-07-22 18:13:11', null);
INSERT INTO `sys_columns` VALUES ('18', '1', 'stgtime', '测试耗时', 'varchar(25)', 'string', 'Stgtime', 'stgtime', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', '18', '', '0', '0', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:41', '2020-07-22 18:13:11', null);
INSERT INTO `sys_columns` VALUES ('19', '1', 'CreateBy', '创建人', 'varchar(128)', 'string', 'CreateBy', 'createBy', '0', '', '1', '1', '', '1', '', 'EQ', 'input', '', '19', '', '0', '1', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:41', '2020-07-22 18:13:11', null);
INSERT INTO `sys_columns` VALUES ('20', '1', 'UpdateBy', '更新者', 'varchar(128)', 'string', 'UpdateBy', 'updateBy', '0', '', '1', '1', '', '1', '', 'EQ', 'input', '', '20', '', '0', '1', '0', '0', '0', '1', '0', '0', '', '', '', '2020-07-09 17:07:41', '2020-07-22 18:13:11', null);

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config` (
  `config_id` int(11) NOT NULL AUTO_INCREMENT,
  `config_name` varchar(128) DEFAULT NULL,
  `config_key` varchar(128) DEFAULT NULL,
  `config_value` varchar(255) DEFAULT NULL,
  `config_type` varchar(64) DEFAULT NULL,
  `remark` varchar(128) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`config_id`),
  KEY `idx_sys_config_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES ('1', '主框架页-默认皮肤样式名称', 'sys_index_skinName', 'skin-blue', 'Y', '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow', '1', '1', '2020-02-29 10:37:48', '2020-04-08 13:03:21', null);
INSERT INTO `sys_config` VALUES ('2', '用户管理-账号初始密码', 'sys.user.initPassword', '123456', 'Y', '初始化密码 123456', '1', '1', '2020-02-29 10:38:23', '2020-03-11 00:35:28', null);
INSERT INTO `sys_config` VALUES ('3', '主框架页-侧边栏主题', 'sys_index_sideTheme', 'theme-dark', 'Y', '深色主题theme-dark，浅色主题theme-light', '1', '1', '2020-02-29 10:39:01', '2020-04-07 23:21:50', null);

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `dept_id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) DEFAULT NULL,
  `dept_path` varchar(255) DEFAULT NULL,
  `dept_name` varchar(128) DEFAULT NULL,
  `sort` int(4) DEFAULT NULL,
  `leader` varchar(128) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `email` varchar(64) DEFAULT NULL,
  `status` int(1) DEFAULT NULL,
  `create_by` varchar(64) DEFAULT NULL,
  `update_by` varchar(64) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`dept_id`),
  KEY `idx_sys_dept_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES ('1', '0', '/0/1', 'NCF', '0', 'aituo', '13918924470', 'atuo@aituo.com', '0', '1', '1', '2020-02-27 15:30:19', '2020-07-19 22:15:54', null);
INSERT INTO `sys_dept` VALUES ('7', '1', '/0/1/7', '研发', '1', '', '', '', '0', '1', '1', '2020-03-08 23:10:59', '2020-07-25 09:14:54', null);
INSERT INTO `sys_dept` VALUES ('8', '1', '/0/1/8', '运维', '0', '', '', '', '0', '1', '1', '2020-03-08 23:11:08', '2020-07-25 09:14:45', null);
INSERT INTO `sys_dept` VALUES ('9', '1', '/0/1/9', '产品', '0', '', '', '', '0', '1', '1', '2020-03-08 23:11:15', '2020-07-25 09:14:36', null);
INSERT INTO `sys_dept` VALUES ('10', '1', '/0/1/10', '人力资源', '3', '张三', '', '', '1', '1', '1', '2020-04-07 23:48:38', '2020-04-07 23:48:46', '2020-07-19 22:16:04');
INSERT INTO `sys_dept` VALUES ('11', '0', '/0/11', '测试部', '0', '', '', '', '0', '1', '', '2020-07-19 22:16:35', '2020-07-19 22:16:35', '2020-07-19 22:16:40');
INSERT INTO `sys_dept` VALUES ('12', '1', '/0/1/12', '测试', '0', '', '', '', '0', '1', '', '2020-07-25 09:14:27', '2020-07-25 09:14:27', null);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data` (
  `dict_code` int(11) NOT NULL AUTO_INCREMENT,
  `dict_sort` int(4) DEFAULT NULL,
  `dict_label` varchar(128) DEFAULT NULL,
  `dict_value` varchar(255) DEFAULT NULL,
  `dict_type` varchar(64) DEFAULT NULL,
  `css_class` varchar(128) DEFAULT NULL,
  `list_class` varchar(128) DEFAULT NULL,
  `is_default` varchar(8) DEFAULT NULL,
  `status` int(1) DEFAULT NULL,
  `default` varchar(8) DEFAULT NULL,
  `create_by` varchar(64) DEFAULT NULL,
  `update_by` varchar(64) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`dict_code`),
  KEY `idx_sys_dict_data_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES ('1', '0', '正常', '0', 'sys_normal_disable', '', '', '', '0', '', '1', '', '系统正常', '2020-02-28 20:55:34', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('2', '0', '停用', '1', 'sys_normal_disable', '', '', '', '0', '', '1', '', '系统停用', '2020-02-28 21:10:41', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('3', '0', '男', '0', 'sys_user_sex', '', '', '', '0', '', '1', '', '性别男', '2020-02-28 21:37:28', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('4', '0', '女', '1', 'sys_user_sex', '', '', '', '0', '', '1', '', '性别女', '2020-02-28 21:37:40', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('5', '0', '未知', '2', 'sys_user_sex', '', '', '', '0', '', '1', '', '性别未知', '2020-02-28 21:38:05', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('6', '0', '显示', '0', 'sys_show_hide', '', '', '', '0', '', '1', '', '显示菜单', '2020-02-28 21:38:36', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('7', '0', '隐藏', '1', 'sys_show_hide', '', '', '', '0', '', '1', '', '隐藏菜单', '2020-02-28 21:38:50', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('8', '0', '是', 'Y', 'sys_yes_no', '', '', '', '0', '', '1', '', '系统默认是', '2020-02-28 21:39:40', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('9', '0', '否', 'N', 'sys_yes_no', '', '', '', '0', '', '1', '', '系统默认否', '2020-02-28 21:40:06', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('10', '0', '正常', '0', 'sys_job_status', '', '', '', '0', '', '1', '', '正常状态', '2020-02-28 21:41:02', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('11', '0', '停用', '1', 'sys_job_status', '', '', '', '0', '', '1', '', '停用状态', '2020-02-28 21:41:15', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('12', '0', '默认', 'DEFAULT', 'sys_job_group', '', '', '', '0', '', '1', '', '默认分组', '2020-02-28 21:41:48', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('13', '0', '系统', 'SYSTEM', 'sys_job_group', '', '', '', '0', '', '1', '', '系统分组', '2020-02-28 21:42:02', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('14', '0', '通知', '1', 'sys_notice_type', '', '', '', '0', '', '1', '', '通知', '2020-02-28 21:42:43', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('15', '0', '公告', '2', 'sys_notice_type', '', '', '', '0', '', '1', '', '公告', '2020-02-28 21:42:53', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('16', '0', '正常', '0', 'sys_common_status', '', '', '', '0', '', '1', '', '正常状态', '2020-02-28 21:43:21', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('17', '0', '关闭', '1', 'sys_common_status', '', '', '', '0', '', '1', '', '关闭状态', '2020-02-28 21:43:31', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('18', '0', '新增', '1', 'sys_oper_type', '', '', '', '0', '', '1', '', '新增操作', '2020-02-28 21:44:14', '2020-02-28 22:00:22', null);
INSERT INTO `sys_dict_data` VALUES ('19', '0', '修改', '2', 'sys_oper_type', '', '', '', '0', '', '1', '', '修改操作', '2020-02-28 21:44:34', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('20', '0', '删除', '3', 'sys_oper_type', '', '', '', '0', '', '1', '', '删除操作', '2020-02-28 21:44:52', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('21', '0', '授权', '4', 'sys_oper_type', '', '', '', '0', '', '1', '', '授权操作', '2020-02-28 21:45:18', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('22', '0', '导出', '5', 'sys_oper_type', '', '', '', '0', '', '1', '', '导出操作', '2020-02-28 21:45:44', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('23', '0', '导入', '6', 'sys_oper_type', '', '', '', '0', '', '1', '', '导入操作', '2020-02-28 21:46:02', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('24', '0', '强退', '7', 'sys_oper_type', '', '', '', '0', '', '1', '', '强退操作', '2020-02-28 21:46:25', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('25', '0', '生成代码', '8', 'sys_oper_type', '', '', '', '0', '', '1', '', '生成操作', '2020-02-28 21:46:53', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('26', '0', '清空数据', '9', 'sys_oper_type', '', '', '', '0', '', '1', '', '清空操作', '2020-02-28 21:47:15', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('27', '0', '成功', '0', 'sys_notice_status', '', '', '', '0', '', '1', '', '成功状态', '2020-02-28 22:03:24', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('28', '0', '失败', '1', 'sys_notice_status', '', '', '', '0', '', '1', '', '失败状态', '2020-02-28 22:03:39', '2020-03-08 23:11:15', null);
INSERT INTO `sys_dict_data` VALUES ('29', '0', '登录', '10', 'sys_oper_type', '', '', '', '0', '', '1', '1', '登录操作', '2020-03-15 18:35:23', '2020-03-15 18:39:30', null);
INSERT INTO `sys_dict_data` VALUES ('30', '0', '退出', '11', 'sys_oper_type', '', '', '', '0', '', '1', '', '', '2020-03-15 18:35:39', '2020-03-15 18:35:39', null);
INSERT INTO `sys_dict_data` VALUES ('31', '0', '获取验证码', '12', 'sys_oper_type', '', '', '', '0', '', '1', '', '获取验证码', '2020-03-15 18:38:42', '2020-03-15 18:35:39', null);

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type` (
  `dict_id` int(11) NOT NULL AUTO_INCREMENT,
  `dict_name` varchar(128) DEFAULT NULL,
  `dict_type` varchar(128) DEFAULT NULL,
  `status` int(1) DEFAULT NULL,
  `create_by` varchar(11) DEFAULT NULL,
  `update_by` varchar(11) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`dict_id`),
  KEY `idx_sys_dict_type_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES ('1', '系统开关', 'sys_normal_disable', '0', '1', '1', '系统开关列表', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_dict_type` VALUES ('2', '用户性别', 'sys_user_sex', '0', '1', '', '用户性别列表', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_dict_type` VALUES ('3', '菜单状态', 'sys_show_hide', '0', '1', '', '菜单状态列表', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_dict_type` VALUES ('4', '系统是否', 'sys_yes_no', '0', '1', '', '系统是否列表', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_dict_type` VALUES ('5', '任务状态', 'sys_job_status', '0', '1', '', '任务状态列表', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_dict_type` VALUES ('6', '任务分组', 'sys_job_group', '0', '1', '', '任务分组列表', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_dict_type` VALUES ('7', '通知类型', 'sys_notice_type', '0', '1', '', '通知类型列表', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_dict_type` VALUES ('8', '系统状态', 'sys_common_status', '0', '1', '', '登录状态列表', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_dict_type` VALUES ('9', '操作类型', 'sys_oper_type', '0', '1', '', '操作类型列表', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_dict_type` VALUES ('10', '通知状态', 'sys_notice_status', '0', '1', '', '通知状态列表', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_dict_type` VALUES ('11', '1', '1', '1', '1', '1', '1', '2020-04-11 15:52:48', null, null);

-- ----------------------------
-- Table structure for sys_loginlog
-- ----------------------------
DROP TABLE IF EXISTS `sys_loginlog`;
CREATE TABLE `sys_loginlog` (
  `info_id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(128) DEFAULT NULL,
  `status` int(1) DEFAULT NULL,
  `ipaddr` varchar(255) DEFAULT NULL,
  `login_location` varchar(255) DEFAULT NULL,
  `browser` varchar(255) DEFAULT NULL,
  `os` varchar(255) DEFAULT NULL,
  `platform` varchar(255) DEFAULT NULL,
  `login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `msg` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`info_id`),
  KEY `idx_sys_loginlog_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=734 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_loginlog
-- ----------------------------
 
-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `menu_id` int(11) NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(11) DEFAULT NULL,
  `title` varchar(64) DEFAULT NULL,
  `icon` varchar(128) DEFAULT NULL,
  `path` varchar(128) DEFAULT NULL,
  `paths` varchar(128) DEFAULT NULL,
  `menu_type` varchar(1) DEFAULT NULL,
  `action` varchar(16) DEFAULT NULL,
  `permission` varchar(32) DEFAULT NULL,
  `parent_id` int(11) DEFAULT NULL,
  `no_cache` char(1) DEFAULT NULL,
  `breadcrumb` varchar(255) DEFAULT NULL,
  `component` varchar(255) DEFAULT NULL,
  `sort` int(4) DEFAULT NULL,
  `visible` char(1) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `is_frame` int(1) DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`menu_id`),
  KEY `idx_sys_menu_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=307 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES ('2', 'Upms', '系统管理', 'example', '/upms', '/0/2', 'M', '无', '', '0', '1', '', 'Layout', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('3', 'Sysuser', '用户管理', 'user', 'sysuser', '/0/2/3', 'C', '无', 'system:sysuser:list', '2', null, null, '/sysuser/index', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:10:42', null);
INSERT INTO `sys_menu` VALUES ('43', null, '新增用户', null, '/api/v1/sysuser', '/0/2/3/43', 'F', 'POST', 'system:sysuser:add', '3', null, null, null, '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('44', null, '查询用户', null, '/api/v1/sysuser', '/0/2/3/44', 'F', 'GET', 'system:sysuser:query', '3', null, null, null, '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('45', null, '修改用户', null, '/api/v1/sysuser/', '/0/2/3/45', 'F', 'PUT', 'system:sysuser:edit', '3', null, null, null, '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('46', null, '删除用户', null, '/api/v1/sysuser/', '/0/2/3/46', 'F', 'DELETE', 'system:sysuser:remove', '3', null, null, null, '0', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 15:32:45', null);
INSERT INTO `sys_menu` VALUES ('51', 'Menu', '菜单管理', 'tree-table', 'menu', '/0/2/51', 'C', '无', 'system:sysmenu:list', '2', '1', '', '/menu/index', '3', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-07-26 22:02:35', null);
INSERT INTO `sys_menu` VALUES ('52', 'Role', '角色管理', 'peoples', 'role', '/0/2/52', 'C', '无', 'system:sysrole:list', '2', '1', '', '/role/index', '2', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:16:19', null);
INSERT INTO `sys_menu` VALUES ('56', 'Dept', '部门管理', 'tree', 'dept', '/0/2/56', 'C', '无', 'system:sysdept:list', '2', '0', '', '/dept/index', '4', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:16:47', null);
INSERT INTO `sys_menu` VALUES ('57', 'post', '岗位管理', 'pass', 'post', '/0/2/57', 'C', '无', 'system:syspost:list', '2', '0', '', '/post/index', '5', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:16:53', null);
INSERT INTO `sys_menu` VALUES ('58', 'Dict', '字典管理', 'education', 'dict', '/0/2/58', 'C', '无', 'system:sysdicttype:list', '2', '0', '', '/dict/index', '6', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:17:04', '2020-07-25 09:07:24');
INSERT INTO `sys_menu` VALUES ('59', 'DictData', '字典数据', 'education', 'dict/data/:dictId', '/0/2/59', 'C', '无', 'system:sysdictdata:list', '2', '0', '', '/dict/data', '100', '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:17:25', '2020-07-25 09:10:20');
INSERT INTO `sys_menu` VALUES ('60', 'Tools', '系统接口', 'component', '/tools', '/0/60', 'M', '无', '', '0', '0', '', 'Layout', '3', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-07-27 07:04:16', null);
INSERT INTO `sys_menu` VALUES ('61', 'Swagger', '系统API', 'guide', 'swagger', '/0/60/61', 'C', '无', '', '60', '0', '', '/tools/swagger/index', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-07-26 21:59:53', null);
INSERT INTO `sys_menu` VALUES ('62', 'Config', '参数设置', 'list', '/config', '/0/2/62', 'C', '无', 'system:sysconfig:list', '2', '0', '', '/config/index', '7', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:17:16', '2020-07-25 09:07:56');
INSERT INTO `sys_menu` VALUES ('63', '', '接口权限', 'bug', '', '/0/63', 'M', '', '', '0', '0', '', '', '99', '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 16:39:32', null);
INSERT INTO `sys_menu` VALUES ('64', '', '用户管理', 'user', '', '/0/63/64', 'M', '', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('66', '', '菜单管理', 'tree-table', '', '/0/63/66', 'C', '', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('67', '', '菜单列表', 'tree-table', '/api/v1/menulist', '/0/63/66/67', 'A', 'GET', '', '66', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('68', '', '新建菜单', 'tree', '/api/v1/menu', '/0/63/66/68', 'A', 'POST', '', '66', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('69', '', '字典', 'dict', '', '/0/63/69', 'M', '', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, '2020-07-26 22:01:06');
INSERT INTO `sys_menu` VALUES ('70', '', '类型', 'dict', '', '/0/63/69/70', 'C', '', '', '69', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('71', '', '字典类型获取', 'tree', '/api/v1/dict/databytype/', '/0/63/256/71', 'A', 'GET', '', '256', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, '2020-07-26 22:41:32');
INSERT INTO `sys_menu` VALUES ('72', '', '修改菜单', 'bug', '/api/v1/menu', '/0/63/66/72', 'A', 'PUT', '', '66', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('73', '', '删除菜单', 'bug', '/api/v1/menu/:id', '/0/63/66/73', 'A', 'DELETE', '', '66', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('74', null, '管理员列表', 'bug', '/api/v1/sysUserList', '/0/63/64/74', 'A', 'GET', null, '64', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('75', null, '根据id获取管理员', 'bug', '/api/v1/sysUser/:id', '/0/63/64/75', 'A', 'GET', null, '64', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('76', null, '获取管理员', 'bug', '/api/v1/sysUser/', '/0/63/64/76', 'A', 'GET', null, '64', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('77', null, '创建管理员', 'bug', '/api/v1/sysUser', '/0/63/64/77', 'A', 'POST', null, '64', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('78', null, '修改管理员', 'bug', '/api/v1/sysUser', '/0/63/64/78', 'A', 'PUT', null, '64', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('79', null, '删除管理员', 'bug', '/api/v1/sysUser/:id', '/0/63/64/79', 'A', 'DELETE', null, '64', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('80', null, '当前用户个人信息', 'bug', '/api/v1/user/profile', '/0/63/256/267/80', 'A', 'GET', null, '267', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-05-03 20:50:40', null);
INSERT INTO `sys_menu` VALUES ('81', null, '角色列表', 'bug', '/api/v1/rolelist', '/0/63/201/81', 'A', 'GET', null, '201', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('82', null, '获取角色信息', 'bug', '/api/v1/role/:id', '/0/63/201/82', 'A', 'GET', null, '201', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('83', null, '创建角色', 'bug', '/api/v1/role', '/0/63/201/83', 'A', 'POST', null, '201', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('84', null, '修改角色', 'bug', '/api/v1/role', '/0/63/201/84', 'A', 'PUT', null, '201', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('85', null, '删除角色', 'bug', '/api/v1/role/:id', '/0/63/201/85', 'A', 'DELETE', null, '201', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('86', null, '参数列表', 'bug', '/api/v1/configList', '/0/63/202/86', 'A', 'GET', null, '202', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('87', null, '根据id获取参数', 'bug', '/api/v1/config/:id', '/0/63/202/87', 'A', 'GET', null, '202', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('88', null, '根据key获取参数', 'bug', '/api/v1/configKey/:id', '/0/63/256/88', 'A', 'GET', null, '256', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-05-03 20:54:37', '2020-07-26 22:41:47');
INSERT INTO `sys_menu` VALUES ('89', null, '创建参数', 'bug', '/api/v1/config', '/0/63/202/89', 'A', 'POST', null, '202', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('90', null, '修改参数', 'bug', '/api/v1/config', '/0/63/202/90', 'A', 'PUT', null, '202', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('91', null, '删除参数', 'bug', '/api/v1/config/:id', '/0/63/202/91', 'A', 'DELETE', null, '202', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('92', null, '获取角色菜单', 'bug', '/api/v1/menurole', '/0/63/201/92', 'A', 'GET', null, '201', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('93', null, '根据角色id获取角色', 'bug', '/api/v1/roleMenuTreeselect/:id', '/0/63/201/93', 'A', 'GET', null, '201', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('94', null, '获取菜单树', 'bug', '/api/v1/menuTreeselect', '/0/63/256/94', 'A', 'GET', null, '256', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-05-03 20:52:11', null);
INSERT INTO `sys_menu` VALUES ('95', null, '获取角色菜单', 'bug', '/api/v1/rolemenu', '/0/63/205/95', 'A', 'GET', null, '205', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('96', null, '创建角色菜单', 'bug', '/api/v1/rolemenu', '/0/63/205/96', 'A', 'POST', null, '205', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('97', null, '删除用户菜单数据', 'bug', '/api/v1/rolemenu/:id', '/0/63/205/97', 'A', 'DELETE', null, '205', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('103', null, '部门菜单列表', 'bug', '/api/v1/deptList', '/0/63/203/103', 'A', 'GET', null, '203', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('104', null, '根据id获取部门信息', 'bug', '/api/v1/dept/:id', '/0/63/203/104', 'A', 'GET', null, '203', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('105', null, '创建部门', 'bug', '/api/v1/dept', '/0/63/203/105', 'A', 'POST', null, '203', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('106', null, '修改部门', 'bug', '/api/v1/dept', '/0/63/203/106', 'A', 'PUT', null, '203', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('107', null, '删除部门', 'bug', '/api/v1/dept/:id', '/0/63/203/107', 'A', 'DELETE', null, '203', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('108', null, '字典数据列表', 'bug', '/api/v1/dict/datalist', '/0/63/69/206/108', 'A', 'GET', null, '206', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('109', null, '通过编码获取字典数据', 'bug', '/api/v1/dict/data/:id', '/0/63/69/206/109', 'A', 'GET', null, '206', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('110', null, '通过字典类型获取字典数据', 'bug', '/api/v1/dict/databytype/:id', '/0/63/256/110', 'A', 'GET', null, '256', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, '2020-07-26 22:41:09');
INSERT INTO `sys_menu` VALUES ('111', null, '创建字典数据', 'bug', '/api/v1/dict/data', '/0/63/69/206/111', 'A', 'POST', null, '206', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('112', null, '修改字典数据', 'bug', '/api/v1/dict/data/', '/0/63/69/206/112', 'A', 'PUT', null, '206', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('113', null, '删除字典数据', 'bug', '/api/v1/dict/data/:id', '/0/63/69/206/113', 'A', 'DELETE', null, '206', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('114', null, '字典类型列表', 'bug', '/api/v1/dict/typelist', '/0/63/69/70/114', 'A', 'GET', null, '70', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('115', null, '通过id获取字典类型', 'bug', '/api/v1/dict/type/:id', '/0/63/69/70/115', 'A', 'GET', null, '70', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('116', null, '创建字典类型', 'bug', '/api/v1/dict/type', '/0/63/69/70/116', 'A', 'POST', null, '70', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('117', null, '修改字典类型', 'bug', '/api/v1/dict/type', '/0/63/69/70/117', 'A', 'PUT', null, '70', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('118', null, '删除字典类型', 'bug', '/api/v1/dict/type/:id', '/0/63/69/70/118', 'A', 'DELETE', null, '70', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('119', null, '获取岗位列表', 'bug', '/api/v1/postlist', '/0/63/204/119', 'A', 'GET', null, '204', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('120', null, '通过id获取岗位信息', 'bug', '/api/v1/post/:id', '/0/63/204/120', 'A', 'GET', null, '204', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('121', null, '创建岗位', 'bug', '/api/v1/post', '/0/63/204/121', 'A', 'POST', null, '204', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('122', null, '修改岗位', 'bug', '/api/v1/post', '/0/63/204/122', 'A', 'PUT', null, '204', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('123', null, '删除岗位', 'bug', '/api/v1/post/:id', '/0/63/204/123', 'A', 'DELETE', null, '204', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('137', null, '菜单列表', 'bug', '/api/v1/menulist', '/0/63/66/137', 'A', 'GET', null, '66', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('138', null, '获取根据id菜单信息', 'bug', '/api/v1/menu/:id', '/0/63/66/138', 'A', 'GET', null, '66', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('139', null, '创建菜单', 'bug', '/api/v1/menu', '/0/63/66/139', 'A', 'POST', null, '66', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('140', null, '修改菜单', 'bug', '/api/v1/menu/:id', '/0/63/66/140', 'A', 'PUT', null, '66', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('141', null, '删除菜单', 'bug', '/api/v1/menu/:id', '', 'A', 'DELETE', null, '66', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('142', null, '获取角色对应的菜单id数组', 'bug', '/api/v1/menuids', '/0/63/256/142', 'A', 'GET', null, '256', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('201', '', '角色管理', 'peoples', '', '/0/63/201', 'C', 'GET', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('202', '', '参数设置', 'list', '', '/0/63/202', 'C', 'DELETE', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, '2020-07-26 22:01:22');
INSERT INTO `sys_menu` VALUES ('203', '', '部门管理', 'tree', '', '/0/63/203', 'C', 'POST', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('204', '', '岗位管理', 'pass', '', '/0/63/204', 'C', '', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('205', '', '角色菜单管理', 'nested', '', '/0/63/205', 'C', '', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('206', '', '数据', '', '', '/0/63/69/206', 'C', 'PUT', '', '69', '0', '', '', '2', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('211', 'Log', '日志管理', 'log', '/log', '/0/2/211', 'M', '', '', '2', '0', '', '/log/index', '8', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:15:32', null);
INSERT INTO `sys_menu` VALUES ('212', 'LoginLog', '登录日志', 'logininfor', '/loginlog', '/0/2/211/212', 'C', '', 'system:sysloginlog:list', '211', '0', '', '/loginlog/index', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('213', null, '获取登录日志', 'bug', '/api/v1/loginloglist', '/0/63/214/213', 'A', 'GET', null, '214', null, null, null, '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('214', '', '日志管理', 'log', '', '/0/63/214', 'M', 'GET', '', '63', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('215', '', '删除日志', 'bug', '/api/v1/loginlog/:id', '/0/63/214/215', 'A', 'DELETE', '', '214', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('216', 'OperLog', '操作日志', 'skill', '/operlog', '/0/2/211/216', 'C', '', 'system:sysoperlog:list', '211', '0', '', '/operlog/index', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('217', '', '获取操作日志', 'bug', '/api/v1/operloglist', '/0/63/214/217', 'A', 'GET', '', '214', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('220', '', '新增菜单', '', '', '/0/2/51/220', 'F', '', 'system:sysmenu:add', '51', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('221', '', '修改菜单', 'edit', '', '/0/2/51/221', 'F', '', 'system:sysmenu:edit', '51', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('222', '', '查询菜单', 'search', '', '/0/2/51/222', 'F', '', 'system:sysmenu:query', '51', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('223', '', '删除菜单', '', '', '/0/2/51/223', 'F', '', 'system:sysmenu:remove', '51', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('224', '', '新增角色', '', '', '/0/2/52/224', 'F', '', 'system:sysrole:add', '52', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('225', '', '查询角色', '', '', '/0/2/52/225', 'F', '', 'system:sysrole:query', '52', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('226', '', '修改角色', '', '', '/0/2/52/226', 'F', '', 'system:sysrole:edit', '52', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('227', '', '删除角色', '', '', '/0/2/52/227', 'F', '', 'system:sysrole:remove', '52', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('228', '', '查询部门', '', '', '/0/2/56/228', 'F', '', 'system:sysdept:query', '56', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('229', '', '新增部门', '', '', '/0/2/56/229', 'F', '', 'system:sysdept:add', '56', '0', '', '', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('230', '', '修改部门', '', '', '/0/2/56/230', 'F', '', 'system:sysdept:edit', '56', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('231', '', '删除部门', '', '', '/0/2/56/231', 'F', '', 'system:sysdept:remove', '56', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('232', '', '查询岗位', '', '', '/0/2/57/232', 'F', '', 'system:syspost:query', '57', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('233', '', '新增岗位', '', '', '/0/2/57/233', 'F', '', 'system:syspost:add', '57', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('234', '', '修改岗位', '', '', '/0/2/57/234', 'F', '', 'system:syspost:edit', '57', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('235', '', '删除岗位', '', '', '/0/2/57/235', 'F', '', 'system:syspost:remove', '57', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('236', '', '字典查询', '', '', '/0/2/58/236', 'F', '', 'system:sysdicttype:query', '58', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('237', '', '新增类型', '', '', '/0/2/58/237', 'F', '', 'system:sysdicttype:add', '58', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('238', '', '修改类型', '', '', '/0/2/58/238', 'F', '', 'system:sysdicttype:edit', '58', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('239', '', '删除类型', '', '', '/0/2/58/239', 'F', '', 'system:sysdicttype:remove', '58', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('240', '', '查询数据', '', '', '/0/2/59/240', 'F', '', 'system:sysdictdata:query', '59', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('241', '', '新增数据', '', '', '/0/2/59/241', 'F', '', 'system:sysdictdata:add', '59', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('242', '', '修改数据', '', '', '/0/2/59/242', 'F', '', 'system:sysdictdata:edit', '59', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('243', '', '删除数据', '', '', '/0/2/59/243', 'F', '', 'system:sysdictdata:remove', '59', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('244', '', '查询参数', '', '', '/0/2/62/244', 'F', '', 'system:sysconfig:query', '62', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('245', '', '新增参数', '', '', '/0/2/62/245', 'F', '', 'system:sysconfig:add', '62', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('246', '', '修改参数', '', '', '/0/2/62/246', 'F', '', 'system:sysconfig:edit', '62', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('247', '', '删除参数', '', '', '/0/2/62/247', 'F', '', 'system:sysconfig:remove', '62', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('248', '', '查询登录日志', '', '', '/0/2/211/212/248', 'F', '', 'system:sysloginlog:query', '212', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('249', '', '删除登录日志', '', '', '/0/2/211/212/249', 'F', '', 'system:sysloginlog:remove', '212', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('250', '', '查询操作日志', '', '', '/0/2/211/216/250', 'F', '', 'system:sysoperlog:query', '216', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('251', '', '删除操作日志', '', '', '/0/2/211/216/251', 'F', '', 'system:sysoperlog:remove', '216', '0', '', '', '0', '0', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('252', '', '获取登录用户信息', '', '/api/v1/getinfo', '/0/63/256/252', 'A', 'GET', '', '256', '0', '', '', '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('253', '', '角色数据权限', '', '/api/v1/roledatascope', '/0/63/201/253', 'A', 'PUT', '', '201', '0', '', '', '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('254', '', '部门树接口【数据权限】', '', '/api/v1/roleDeptTreeselect/:id', '/0/63/256/254', 'A', 'GET', '', '256', '0', '', '', '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('255', '', '部门树【用户列表】', '', '/api/v1/deptTree', '/0/63/256/255', 'A', 'GET', '', '256', '0', '', '', '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('256', '', '必开接口', '', '', '/0/63/256', 'M', 'GET', '', '63', '0', '', '', '0', '1', '1', '', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('257', '', '通过key获取参数', 'bug', '/api/v1/configKey/:id', '/0/63/256/257', 'A', 'GET', '', '256', '0', '', '', '1', '1', '1', '1', '0', '2020-04-11 15:52:48', null, '2020-07-26 22:41:27');
INSERT INTO `sys_menu` VALUES ('258', '', '退出登录', '', '/api/v1/logout', '/0/63/256/258', 'A', 'POST', '', '256', '0', '', '', '0', '1', '1', '1', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('259', '', '头像上传', '', '/api/v1/user/avatar', '/0/63/256/267/259', 'A', 'POST', '', '267', '0', '', '', '0', '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-05-03 20:50:05', null);
INSERT INTO `sys_menu` VALUES ('260', '', '修改密码', '', '/api/v1/user/pwd', '/0/63/256/260', 'A', 'PUT', '', '256', '0', '', '', '0', '1', '1', '', '0', '2020-04-11 15:52:48', null, null);
INSERT INTO `sys_menu` VALUES ('261', 'Gen', '代码生成', 'code', 'gen', '/0/60/261', 'C', '', '', '60', '0', '', '/tools/gen/index', '2', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:18:12', '2020-07-26 21:59:25');
INSERT INTO `sys_menu` VALUES ('262', 'EditTable', '代码生成修改', 'build', 'editTable', '/0/60/262', 'C', '', '', '60', '0', '', '/tools/gen/editTable', '100', '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-05-03 20:38:36', '2020-07-26 21:59:32');
INSERT INTO `sys_menu` VALUES ('263', '', '字典类型下拉框【生成功能】', '', '/api/v1/dict/typeoptionselect', '/0/63/256/263', 'A', 'GET', '', '256', '0', '', '', '0', '1', '1', '', '0', '2020-04-11 15:52:48', null, '2020-07-26 22:40:13');
INSERT INTO `sys_menu` VALUES ('264', 'Build', '表单构建', 'build', 'build', '/0/60/264', 'C', '', '', '60', '0', '', '/tools/build/index', '1', '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:18:05', '2020-07-26 21:59:19');
INSERT INTO `sys_menu` VALUES ('267', '', '个人中心', '', '', '/0/63/256/267', 'M', '', '', '256', '0', '', '', '0', '1', '1', '', '1', '2020-05-03 20:49:39', '2020-05-03 20:49:39', null);
INSERT INTO `sys_menu` VALUES ('268', '', '持续集成', 'calendar', '/releasecicd', '/0/268', 'M', '', '', '0', '0', '', 'Layout', '0', '0', '1', '1', '1', '2020-07-09 18:14:07', '2020-07-17 05:49:54', null);
INSERT INTO `sys_menu` VALUES ('269', 'releasecicd', '已上线', 'documentation', '/releasecicd/index', '/0/268/269', 'C', '', 'releasecicd:releasecicd:list', '268', '0', '', '/releasecicd/index', '3', '0', '1', '1', '1', '2020-07-09 18:16:59', '2020-07-27 10:57:55', null);
INSERT INTO `sys_menu` VALUES ('270', '', '持续集成', 'bug', '', '/0/63/270', 'M', '', '', '63', '0', '', '', '0', '1', '1', '1', '1', '2020-07-09 18:18:24', '2020-07-14 17:36:58', null);
INSERT INTO `sys_menu` VALUES ('271', '', '已发生产', '', '', '/0/63/270/271', 'M', '', '', '270', '0', '', '', '0', '1', '1', '1', '1', '2020-07-09 18:19:29', '2020-07-22 17:50:15', null);
INSERT INTO `sys_menu` VALUES ('272', '', '录入发布', 'bug', '/api/v1/releasecicd', '/0/63/270/271/272', 'A', 'POST', '', '271', '0', '', '', '0', '1', '1', '', '1', '2020-07-09 18:20:29', '2020-07-09 18:20:29', null);
INSERT INTO `sys_menu` VALUES ('273', '', '修改发布', 'bug', '/api/v1/releasecicd', '/0/63/270/271/273', 'A', 'PUT', '', '271', '0', '', '', '0', '1', '1', '1', '1', '2020-07-09 18:21:24', '2020-07-09 18:23:13', null);
INSERT INTO `sys_menu` VALUES ('274', '', '删除发布', 'bug', '/api/v1/releasecicd/:id', '/0/63/270/271/274', 'A', 'DELETE', '', '271', '0', '', '', '0', '1', '1', '', '1', '2020-07-09 18:22:21', '2020-07-09 18:22:21', null);
INSERT INTO `sys_menu` VALUES ('275', '', '发布列表', 'bug', '/api/v1/releasecicdList', '/0/63/270/271/275', 'A', 'GET', '', '271', '0', '', '', '0', '1', '1', '1', '1', '2020-07-09 18:24:26', '2020-07-09 18:24:36', null);
INSERT INTO `sys_menu` VALUES ('276', '', '获取发布', 'bug', '/api/v1/releasecicd/:id', '/0/63/270/271/276', 'A', 'GET', '', '271', '0', '', '', '0', '1', '1', '', '1', '2020-07-09 18:25:24', '2020-07-09 18:25:24', null);
INSERT INTO `sys_menu` VALUES ('278', '', '发布概览', '', '', '/0/63/270/278', 'M', '', '', '270', '0', '', '', '0', '1', '1', '1', '1', '2020-07-14 17:38:54', '2020-07-14 17:40:06', '2020-07-14 18:57:08');
INSERT INTO `sys_menu` VALUES ('279', '', '发布概览', '', '', '/0/63/270/279', 'M', '', '', '270', '0', '', '', '0', '1', '1', '', '1', '2020-07-14 17:39:42', '2020-07-14 17:39:42', '2020-07-14 17:39:55');
INSERT INTO `sys_menu` VALUES ('280', '', 'Rid概览', 'bug', '/api/v1/releaseridList', '/0/63/270/278/280', 'A', 'GET', '', '278', '0', '', '', '0', '1', '1', '1', '1', '2020-07-14 17:41:54', '2020-07-14 18:21:45', null);
INSERT INTO `sys_menu` VALUES ('281', '', '发布概览', 'documentation', '/releaserid', '/0/281', 'M', '', '', '0', '0', '', 'Layout', '0', '0', '1', '', '1', '2020-07-14 18:43:45', '2020-07-14 18:43:45', '2020-07-17 06:10:05');
INSERT INTO `sys_menu` VALUES ('282', 'releaserid', '计划概览', 'documentation', '/releaserid/index', '/0/268/282', 'C', '', 'releaserid:releaserid:list', '268', '0', '', '/releaserid/index', '4', '0', '1', '1', '1', '2020-07-14 18:45:35', '2020-07-22 18:20:00', null);
INSERT INTO `sys_menu` VALUES ('283', '', '发布概览', 'bug', '', '/0/63/283', 'M', '', '', '63', '0', '', '', '0', '1', '1', '', '1', '2020-07-14 18:51:48', '2020-07-14 18:51:48', null);
INSERT INTO `sys_menu` VALUES ('284', '', 'RID概览', '', '', '/0/63/283/284', 'M', '', '', '283', '0', '', '', '0', '1', '1', '', '1', '2020-07-14 18:53:37', '2020-07-14 18:53:37', null);
INSERT INTO `sys_menu` VALUES ('285', '', 'RID列表', 'bug', '/api/v1/releaseridList', '/0/63/283/284/285', 'A', 'GET', '', '284', '0', '', '', '0', '1', '1', '', '1', '2020-07-14 18:54:56', '2020-07-14 18:54:56', null);
INSERT INTO `sys_menu` VALUES ('286', '', '录入rid', 'bug', '/api/v1/releaserid', '/0/63/283/284/286', 'A', 'POST', '', '284', '0', '', '', '0', '1', '1', '', '1', '2020-07-15 10:41:03', '2020-07-15 10:41:03', null);
INSERT INTO `sys_menu` VALUES ('287', '', '修改rid', 'bug', '/api/v1/releaserid', '/0/63/283/284/287', 'A', 'PUT', '', '284', '0', '', '', '0', '1', '1', '', '1', '2020-07-15 10:41:50', '2020-07-15 10:41:50', null);
INSERT INTO `sys_menu` VALUES ('288', '', '获取rid', 'bug', '/api/v1/releaserid/:id', '/0/63/283/284/288', 'A', 'GET', '', '284', '0', '', '', '0', '1', '1', '', '1', '2020-07-15 10:42:42', '2020-07-15 10:42:42', null);
INSERT INTO `sys_menu` VALUES ('289', '', '删除rid', 'bug', '/api/v1/releaserid/:id', '/0/63/283/284/289', 'A', 'DELETE', '', '284', '0', '', '', '0', '1', '1', '', '1', '2020-07-15 10:43:28', '2020-07-15 10:43:28', null);
INSERT INTO `sys_menu` VALUES ('290', 'releasetodo', '未提测', 'documentation', '/releasetodo/index', '/0/268/290', 'C', '', 'releasetodo:releasetodo:list', '268', '0', '', '/releasetodo/index', '1', '0', '1', '1', '1', '2020-07-17 17:26:18', '2020-07-22 17:49:39', null);
INSERT INTO `sys_menu` VALUES ('291', '', '未提测', '', '', '/0/63/291', 'M', '', '', '63', '0', '', '', '0', '1', '1', '1', '1', '2020-07-17 17:28:02', '2020-07-22 17:50:33', null);
INSERT INTO `sys_menu` VALUES ('292', '', '录入需求', 'bug', '/api/v1/releasetodo', '/0/63/291/292', 'A', 'POST', '', '291', '0', '', '', '0', '1', '1', '', '1', '2020-07-17 17:29:06', '2020-07-17 17:29:06', null);
INSERT INTO `sys_menu` VALUES ('293', '', '修改需求', 'bug', '/api/v1/releasetodo', '/0/63/291/293', 'A', 'PUT', '/api/v1/releasetodo', '291', '0', '', '', '0', '1', '1', '1', '1', '2020-07-17 17:32:27', '2020-07-17 17:35:05', '2020-07-17 17:35:23');
INSERT INTO `sys_menu` VALUES ('294', '', '修改需求', 'bug', '/api/v1/releasetodo', '/0/63/291/294', 'A', 'PUT', '', '291', '0', '', '', '0', '1', '1', '', '1', '2020-07-17 17:35:58', '2020-07-17 17:35:58', null);
INSERT INTO `sys_menu` VALUES ('295', '', '删除需求', 'bug', '/api/v1/releasetodo/:id', '/0/63/291/295', 'A', 'PUT', '', '291', '0', '', '', '0', '1', '1', '', '1', '2020-07-17 17:37:08', '2020-07-17 17:37:08', '2020-07-17 17:37:46');
INSERT INTO `sys_menu` VALUES ('296', '', '删除需求', 'bug', '/api/v1/releasetodo/:id', '/0/63/291/296', 'A', 'DELETE', '', '291', '0', '', '', '0', '1', '1', '', '1', '2020-07-17 17:38:13', '2020-07-17 17:38:13', null);
INSERT INTO `sys_menu` VALUES ('297', '', '需求列表', 'bug', '/api/v1/releasetodoList', '/0/63/291/297', 'A', 'GET', '', '291', '0', '', '', '0', '1', '1', '', '1', '2020-07-17 17:39:09', '2020-07-17 17:39:09', null);
INSERT INTO `sys_menu` VALUES ('298', '', '获取需求', 'bug', '/api/v1/releasetodo/:id', '/0/63/291/298', 'A', 'GET', '', '291', '0', '', '', '0', '1', '1', '', '1', '2020-07-17 17:39:49', '2020-07-17 17:39:49', null);
INSERT INTO `sys_menu` VALUES ('299', '', '日发布图', 'bug', '/api/v1/releasechart/deploybyday', '/0/63/270/271/299', 'A', 'GET', '', '271', '0', '', '', '0', '1', '1', '1', '1', '2020-07-17 23:25:30', '2020-07-19 19:02:41', null);
INSERT INTO `sys_menu` VALUES ('300', '', '时发布图', '', '/api/v1/releasechart/deploybyhour', '/0/63/270/271/299/300', 'A', 'GET', '', '299', '0', '', '', '0', '1', '1', '', '1', '2020-07-19 19:03:10', '2020-07-19 19:03:10', '2020-07-19 19:04:34');
INSERT INTO `sys_menu` VALUES ('301', '', '时发布图', 'bug', '/api/v1/releasechart/deploybyhour', '/0/63/270/271/301', 'A', 'GET', '', '271', '0', '', '', '0', '1', '1', '', '1', '2020-07-19 19:05:06', '2020-07-19 19:05:06', null);
INSERT INTO `sys_menu` VALUES ('302', 'releasestg', '已发测试', 'documentation', '/api/v1/releasestg', '/0/302', 'M', '', '', '0', '0', '', '/api/v1/releasestg', '0', '0', '1', '', '1', '2020-07-22 17:46:01', '2020-07-22 17:46:01', '2020-07-22 17:46:22');
INSERT INTO `sys_menu` VALUES ('303', 'releasestg', '测试中', 'documentation', '/releasestg/index', '/0/268/303', 'C', '', 'releasestg:releasestg:list', '268', '0', '', '/releasestg/index', '2', '0', '1', '1', '1', '2020-07-22 17:47:05', '2020-07-27 10:57:44', null);
INSERT INTO `sys_menu` VALUES ('304', '', '已发测试', '', '', '/0/63/270/304', 'M', '', '', '270', '0', '', '', '0', '1', '1', '', '1', '2020-07-22 17:51:20', '2020-07-22 17:51:20', '2020-07-22 17:52:07');
INSERT INTO `sys_menu` VALUES ('305', '', '已发测试', '', '', '/0/63/305', 'M', '', '', '63', '0', '', '', '0', '1', '1', '', '1', '2020-07-22 17:53:48', '2020-07-22 17:53:48', null);
INSERT INTO `sys_menu` VALUES ('306', '', 'STG列表', 'bug', '/api/v1/releasestgList', '/0/63/305/306', 'A', 'GET', '', '305', '0', '', '', '0', '1', '1', '', '1', '2020-07-22 17:54:43', '2020-07-22 17:54:43', null);

-- ----------------------------
-- Table structure for sys_operlog
-- ----------------------------
DROP TABLE IF EXISTS `sys_operlog`;
CREATE TABLE `sys_operlog` (
  `oper_id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `business_type` varchar(128) DEFAULT NULL,
  `business_types` varchar(128) DEFAULT NULL,
  `method` varchar(128) DEFAULT NULL,
  `request_method` varchar(128) DEFAULT NULL,
  `operator_type` varchar(128) DEFAULT NULL,
  `oper_name` varchar(128) DEFAULT NULL,
  `dept_name` varchar(128) DEFAULT NULL,
  `oper_url` varchar(255) DEFAULT NULL,
  `oper_ip` varchar(128) DEFAULT NULL,
  `oper_location` varchar(128) DEFAULT NULL,
  `oper_param` varchar(255) DEFAULT NULL,
  `status` int(1) DEFAULT NULL,
  `oper_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `json_result` varchar(255) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `latency_time` varchar(128) DEFAULT NULL,
  `user_agent` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`oper_id`),
  KEY `idx_sys_operlog_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1062 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_operlog
-- ----------------------------

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `post_id` int(11) NOT NULL AUTO_INCREMENT,
  `post_name` varchar(128) DEFAULT NULL,
  `post_code` varchar(128) DEFAULT NULL,
  `sort` int(4) DEFAULT NULL,
  `status` int(1) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`post_id`),
  KEY `idx_sys_post_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES ('1', 'JAVA开发工程师', 'DEV', '0', '0', 'JAVA开发工程师', '1', '1', '2020-04-11 15:52:48', '2020-07-25 09:17:04', null);
INSERT INTO `sys_post` VALUES ('2', '测试工程师', 'STG', '2', '0', '测试工程师', '1', '1', '2020-04-11 15:52:48', '2020-07-25 09:17:16', null);
INSERT INTO `sys_post` VALUES ('3', '产品经理', 'PM', '3', '0', '产品经理', '1', '1', '2020-04-11 15:52:48', '2020-07-25 09:17:41', null);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `role_id` int(11) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(128) DEFAULT NULL,
  `status` int(1) DEFAULT NULL,
  `role_key` varchar(128) DEFAULT NULL,
  `role_sort` int(4) DEFAULT NULL,
  `flag` varchar(128) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `admin` char(1) DEFAULT NULL,
  `data_scope` varchar(128) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`role_id`),
  KEY `idx_sys_role_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES ('1', '系统管理员', '0', 'admin', '1', null, '1', null, null, '0', null, '2020-04-11 15:52:48', '2020-07-22 18:03:40', null);
INSERT INTO `sys_role` VALUES ('2', '普通角色', '0', 'common', '2', null, '1', null, null, '0', '2', '2020-04-11 15:52:48', '2020-07-26 22:05:00', null);
INSERT INTO `sys_role` VALUES ('3', '测试角色', '0', 'Tester', '3', '', '1', null, null, '0', null, '2020-04-11 15:52:48', '2020-04-12 14:10:52', null);

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept` (
  `role_id` int(11) DEFAULT NULL,
  `dept_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `role_id` int(11) DEFAULT NULL,
  `menu_id` int(11) DEFAULT NULL,
  `role_name` varchar(128) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES ('1', '2', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '3', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '43', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '44', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '45', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '46', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '51', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '52', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '56', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '57', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '58', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '59', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '60', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '61', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '62', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '63', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '64', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '66', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '67', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '68', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '69', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '70', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '71', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '72', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '73', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '74', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '75', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '76', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '77', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '78', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '79', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '80', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '81', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '82', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '83', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '84', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '85', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '86', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '87', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '88', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '89', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '90', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '91', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '92', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '93', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '94', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '95', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '96', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '97', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '103', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '104', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '105', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '106', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '107', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '108', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '109', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '110', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '111', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '112', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '113', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '114', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '115', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '116', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '117', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '118', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '119', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '120', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '121', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '122', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '123', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '137', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '138', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '139', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '140', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '141', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '142', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '201', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '202', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '203', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '204', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '205', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '206', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '211', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '212', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '213', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '214', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '215', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '216', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '217', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '220', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '221', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '222', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '223', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '224', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '225', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '226', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '227', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '228', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '229', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '230', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '231', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '232', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '233', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '234', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '235', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '236', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '237', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '238', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '239', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '240', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '241', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '242', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '243', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '244', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '245', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '246', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '247', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '248', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '249', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '250', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '251', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '252', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '253', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '254', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '255', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '256', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '257', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '258', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '259', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '260', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '261', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '262', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '263', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '264', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '267', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '268', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '269', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '270', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '271', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '272', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '273', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '274', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '275', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '276', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '282', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '283', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '284', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '285', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '286', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '287', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '288', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '289', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '290', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '291', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '292', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '294', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '296', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '297', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '298', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '299', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '301', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '303', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '305', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('1', '306', 'admin', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '2', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '3', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '43', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '44', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '45', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '46', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '51', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '52', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '56', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '57', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '60', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '61', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '63', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '64', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '66', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '67', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '68', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '71', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '72', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '73', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '74', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '75', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '76', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '77', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '78', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '79', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '80', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '81', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '82', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '83', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '84', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '85', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '88', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '92', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '93', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '94', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '95', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '96', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '97', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '103', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '104', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '105', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '106', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '107', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '110', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '119', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '120', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '121', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '122', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '123', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '137', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '138', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '139', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '140', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '141', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '142', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '201', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '203', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '204', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '205', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '211', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '212', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '213', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '214', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '215', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '216', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '217', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '220', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '221', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '222', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '223', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '224', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '225', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '226', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '227', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '228', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '229', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '230', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '231', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '232', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '233', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '234', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '235', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '248', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '249', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '250', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '251', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '252', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '253', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '254', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '255', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '256', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '257', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '258', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '259', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '260', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '263', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '267', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '268', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '269', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '270', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '271', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '275', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '276', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '282', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '283', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '284', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '285', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '288', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '290', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '291', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '297', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '298', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '303', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '305', 'common', null, null);
INSERT INTO `sys_role_menu` VALUES ('2', '306', 'common', null, null);

-- ----------------------------
-- Table structure for sys_tables
-- ----------------------------
DROP TABLE IF EXISTS `sys_tables`;
CREATE TABLE `sys_tables` (
  `table_id` int(11) NOT NULL AUTO_INCREMENT,
  `table_name` varchar(255) DEFAULT NULL,
  `table_comment` varchar(255) DEFAULT NULL,
  `class_name` varchar(255) DEFAULT NULL,
  `tpl_category` varchar(255) DEFAULT NULL,
  `package_name` varchar(255) DEFAULT NULL,
  `module_name` varchar(255) DEFAULT NULL,
  `business_name` varchar(255) DEFAULT NULL,
  `function_name` varchar(255) DEFAULT NULL,
  `function_author` varchar(255) DEFAULT NULL,
  `pk_column` varchar(255) DEFAULT NULL,
  `pk_go_field` varchar(255) DEFAULT NULL,
  `pk_json_field` varchar(255) DEFAULT NULL,
  `options` varchar(255) DEFAULT NULL,
  `tree_code` varchar(255) DEFAULT NULL,
  `tree_parent_code` varchar(255) DEFAULT NULL,
  `tree_name` varchar(255) DEFAULT NULL,
  `tree` char(1) DEFAULT NULL,
  `crud` char(1) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `is_logical_delete` char(1) DEFAULT NULL,
  `logical_delete` char(1) DEFAULT NULL,
  `logical_delete_column` varchar(128) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`table_id`),
  KEY `idx_sys_tables_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_tables
-- ----------------------------
INSERT INTO `sys_tables` VALUES ('1', 'releasecicd', 'Releasecicd', 'Releasecicd', 'crud', 'releasecicd', 'releasecicd', 'releasecicd', 'Releasecicd', 'wenjianzhang', 'id', 'Id', 'id', '', '', '', '', '0', '1', '', '1', '1', 'is_del', '', '', '2020-07-09 17:07:39', '2020-07-22 18:13:08', null);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `nick_name` varchar(128) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  `salt` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `sex` varchar(255) DEFAULT NULL,
  `email` varchar(128) DEFAULT NULL,
  `dept_id` int(11) DEFAULT NULL,
  `post_id` int(11) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `status` int(1) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `username` varchar(64) DEFAULT NULL,
  `password` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  KEY `idx_sys_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('1', 'linghu', '13918924470', '1', null, '', '0', '405493744@qq.com', '1', '1', '1', '1', null, '0', '2019-11-10 14:05:55', '2020-07-15 10:45:32', null, 'admin', '$2a$10$cKFFTCzGOvaIHHJY2K45Zuwt8TD6oPzYi4s5MzYIBAWCLL6ZhouP2');