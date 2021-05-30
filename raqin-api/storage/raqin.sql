CREATE TABLE `user` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` ENUM ('ADMIN', 'RAQIN', 'VOLUNTEER') NOT NULL,
  `gender` varchar(1) DEFAULT "m",
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT (now()),
  `deleted_at` timestamp
);

CREATE TABLE `book_initiater` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `book_id` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT (now()),
  `deleted_at` timestamp
);

CREATE TABLE `book` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `path` varchar(255) NOT NULL,
  `note` text,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT (now()),
  `deleted_at` timestamp
);

CREATE TABLE `author` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT (now()),
  `deleted_at` timestamp
);

CREATE TABLE `book_author` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `book_id` int NOT NULL,
  `author_id` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT (now()),
  `deleted_at` timestamp
);

CREATE TABLE `book_category` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `book_id` int NOT NULL,
  `category_id` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT (now()),
  `deleted_at` timestamp
);

CREATE TABLE `category` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT (now()),
  `deleted_at` timestamp
);

CREATE TABLE `page` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `book_id` int NOT NULL,
  `path` varchar(255) NOT NULL,
  `number` int NOT NULL,
  `stage` ENUM ('NONE', 'INIT', 'REV', 'DONE') NOT NULL DEFAULT "NONE",
  `approved_revision` int,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT (now()),
  `deleted_at` timestamp
);

CREATE TABLE `page_revision` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `reviewer_id` int NOT NULL,
  `page_id` int NOT NULL,
  `page_text` text,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT (now()),
  `deleted_at` timestamp
);

CREATE TABLE `page_revision_reaction` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `page_revision_id` int NOT NULL,
  `reactor_id` int NOT NULL,
  `reaction` ENUM ('NONE', 'APPROVE', 'DISAPPROVE') NOT NULL DEFAULT "NONE",
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT (now()),
  `deleted_at` timestamp
);

CREATE TABLE `page_revision_comment` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `page_revision_id` int NOT NULL,
  `commenter_id` int NOT NULL,
  `comment` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT (now()),
  `deleted_at` timestamp
);

CREATE TABLE `activity` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `resource` ENUM ('BOOK', 'LINE', 'USER', 'CATEGORY', 'AUTHOR') NOT NULL,
  `action` ENUM ('CREATE', 'UPDATE', 'DELETE') NOT NULL,
  `user_id` int NOT NULL,
  `value` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT (now()),
  `deleted_at` timestamp
);

ALTER TABLE `book_initiater` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `book_initiater` ADD FOREIGN KEY (`book_id`) REFERENCES `book` (`id`);

ALTER TABLE `book_author` ADD FOREIGN KEY (`book_id`) REFERENCES `book` (`id`);

ALTER TABLE `book_author` ADD FOREIGN KEY (`author_id`) REFERENCES `author` (`id`);

ALTER TABLE `book_category` ADD FOREIGN KEY (`book_id`) REFERENCES `book` (`id`);

ALTER TABLE `book_category` ADD FOREIGN KEY (`category_id`) REFERENCES `category` (`id`);

ALTER TABLE `page` ADD FOREIGN KEY (`book_id`) REFERENCES `book` (`id`);

ALTER TABLE `page` ADD FOREIGN KEY (`approved_revision`) REFERENCES `page_revision` (`id`);

ALTER TABLE `page_revision` ADD FOREIGN KEY (`reviewer_id`) REFERENCES `user` (`id`);

ALTER TABLE `page_revision` ADD FOREIGN KEY (`page_id`) REFERENCES `page` (`id`);

ALTER TABLE `page_revision_reaction` ADD FOREIGN KEY (`page_revision_id`) REFERENCES `page_revision` (`id`);

ALTER TABLE `page_revision_reaction` ADD FOREIGN KEY (`reactor_id`) REFERENCES `user` (`id`);

ALTER TABLE `page_revision_comment` ADD FOREIGN KEY (`page_revision_id`) REFERENCES `page_revision` (`id`);

ALTER TABLE `page_revision_comment` ADD FOREIGN KEY (`commenter_id`) REFERENCES `user` (`id`);

ALTER TABLE `activity` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

CREATE UNIQUE INDEX `book_initiater_index_0` ON `book_initiater` (`user_id`, `book_id`);

CREATE UNIQUE INDEX `book_author_index_1` ON `book_author` (`author_id`, `book_id`);

CREATE UNIQUE INDEX `book_category_index_2` ON `book_category` (`category_id`, `book_id`);

CREATE UNIQUE INDEX `page_revision_index_3` ON `page_revision` (`reviewer_id`, `page_id`, `deleted_at`);

CREATE UNIQUE INDEX `page_revision_reaction_index_4` ON `page_revision_reaction` (`reactor_id`, `page_revision_id`);
