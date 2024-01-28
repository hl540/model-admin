-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2024-01-29 00:34:20
-- 服务器版本： 5.7.26
-- PHP 版本： 7.3.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `test`
--

-- --------------------------------------------------------

--
-- 表的结构 `role`
--

CREATE TABLE `role` (
                        `id` int(11) NOT NULL,
                        `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT ''
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `role`
--

INSERT INTO `role` (`id`, `name`) VALUES
                                      (1, '学员'),
                                      (2, '教师');

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE `user` (
                        `id` int(11) NOT NULL,
                        `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
                        `age` int(11) NOT NULL,
                        `sex` tinyint(4) NOT NULL,
                        `shot` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
                        `role_id` int(11) NOT NULL,
                        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `user`
--

INSERT INTO `user` (`id`, `name`, `age`, `sex`, `shot`, `role_id`, `created_at`) VALUES
                                                                                     (1, '张三', 25, 1, 'https://srv.carbonads.net/static/30242/0cb80bb72aaa688ad3b9fa0e955e4313260d52e3', 1, '2024-01-28 13:01:01'),
                                                                                     (3, '李四', 20, 1, 'https://srv.carbonads.net/static/30242/3913207e72f710f619a5492de1f4271d876d8c0f', 1, '2024-01-28 13:01:08'),
                                                                                     (5, '阿花', 18, 0, 'https://srv.carbonads.net/static/30242/d73f1601fd4c38caa238b885f3c610c8cbee3169', 2, '2024-01-28 13:01:10'),
                                                                                     (7, '小明', 15, 1, 'https://srv.carbonads.net/static/30242/e9e159e560c126ca858d8494d70bd0bf7375e539', 2, '2024-01-28 13:01:11');

--
-- 转储表的索引
--

--
-- 表的索引 `role`
--
ALTER TABLE `role`
    ADD PRIMARY KEY (`id`);

--
-- 表的索引 `user`
--
ALTER TABLE `user`
    ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `role`
--
ALTER TABLE `role`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `user`
--
ALTER TABLE `user`
    MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
