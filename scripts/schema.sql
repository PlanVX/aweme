CREATE TABLE `videos` (
    `id` bigint (20),
    `user_id` bigint (20) NOT NULL,
    `video_url` varchar(200) NOT NULL,
    `cover_url` varchar(200) NOT NULL,
    `title` varchar(200) NOT NULL,
    `created_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`),
    KEY `created_at` (`created_at` DESC),
    KEY `user_created` (`user_id`, `created_at` DESC)) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `users` (
    `id` bigint (20),
    `username` varchar(40) NOT NULL,
    `password` varchar(200) NOT NULL,
    `avatar` varchar(200) DEFAULT NULL,
    `background_image` varchar(200) DEFAULT NULL,
    `signature` varchar(200) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`)) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `comments` (
    `id` bigint (20),
    `content` text NOT NULL,
    `video_id` bigint (20) NOT NULL,
    `user_id` bigint (20) NOT NULL,
    `created_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`),
    KEY `video_index` (`video_id`, `created_at` DESC),
    KEY `video_user_index` (`video_id`, `user_id`)) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `likes` (
    `id` bigint (20),
    `video_id` bigint (20) NOT NULL,
    `user_id` bigint (20) NOT NULL,
    `created_at` datetime (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    # `created_at` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_index` (`video_id`, `user_id`),
    KEY `user_index` (`user_id`, `created_at` DESC, `video_id`)) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `relations` (
    `id` bigint (20),
    `user_id` bigint (20) NOT NULL,
    `follow_to` bigint (20) NOT NULL,
    `created_at` datetime NOT NULL,
    PRIMARY KEY (`id`),
    KEY `user_created` (`user_id`, `created_at` DESC),
    KEY `follow_to_created` (`follow_to`, `created_at` DESC)) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
