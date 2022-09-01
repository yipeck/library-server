CREATE TABLE IF NOT EXISTS author (
  Id integer PRIMARY KEY AUTOINCREMENT,
  Name text NOT NULL,
  CountryId integer NOT NULL,
  UserId integer NOT NULL,
  Avatar text NOT NULL DEFAULT '',
  Letter text NOT NULL,
  Created integer NOT NULL
);

CREATE TABLE IF NOT EXISTS book (
  Id integer PRIMARY KEY AUTOINCREMENT,
  ISBN text NOT NULL,
  AuthorId integer NOT NULL,
  CategoryId integer NOT NULL,
  PublisherId integer NOT NULL,
  UserId integer NOT NULL,
  Title text NOT NULL,
  Picture text NOT NULL DEFAULT '',
  PublishYear integer NOT NULL,
  PublishMonth integer NOT NULL,
  ShopYear integer NOT NULL,
  ShopMonth integer NOT NULL,
  ShopDay integer NOT NULL,
  ReadStatus integer NOT NULL DEFAULT -1, -- -1 表示暂未阅读，0 表示正在阅读，1 表示完成阅读
  Letter text NOT NULL,
  Created integer NOT NULL
);

CREATE TABLE IF NOT EXISTS category (
  Id integer PRIMARY KEY AUTOINCREMENT,
  UserId integer NOT NULL,
  Title text NOT NULL,
  Created integer NOT NULL
);

CREATE TABLE IF NOT EXISTS country (
  Id integer PRIMARY KEY AUTOINCREMENT,
  UserId integer NOT NULL,
  Name text NOT NULL
);

CREATE TABLE IF NOT EXISTS publisher (
  Id integer PRIMARY KEY AUTOINCREMENT,
  UserId integer NOT NULL,
  Name text NOT NULL
);

CREATE TABLE IF NOT EXISTS user (
  Id integer PRIMARY KEY AUTOINCREMENT,
  Email text NOT NULL,
  Password text NOT NULL,
  NickName text NOT NULL DEFAULT '',
  Slogan text NOT NULL DEFAULT '',
  Avatar text NOT NULL DEFAULT '',
  Created integer NOT NULL
);
