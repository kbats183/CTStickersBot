DROP TABLE IF EXISTS sticker;
DROP SEQUENCE IF EXISTS sticker_id_seq;

CREATE SEQUENCE sticker_id_seq;
CREATE TABLE sticker
(
    id           INTEGER NOT NULL DEFAULT nextval('sticker_id_seq'),
    tg_set_name  VARCHAR NOT NULL,
    tg_file_id   VARCHAR NOT NULL,
    text_content VARCHAR NOT NULL,
    UNIQUE (id),
    UNIQUE (tg_file_id)
);
ALTER SEQUENCE sticker_id_seq OWNED BY sticker.id;

CREATE INDEX idx_sticker_id ON sticker (id);

DROP TABLE IF EXISTS request;
DROP TABLE IF EXISTS users;
DROP SEQUENCE IF EXISTS users_id_seq;
DROP SEQUENCE IF EXISTS request_id_seq;

CREATE SEQUENCE users_id_seq;
CREATE TABLE users
(
    id    INTEGER NOT NULL DEFAULT nextval('users_id_seq'),
    tg_id INTEGER NOT NULL,
    login VARCHAR NOT NULL,
    UNIQUE (id),
    UNIQUE (tg_id),
    UNIQUE (login)
);
ALTER SEQUENCE users_id_seq OWNED BY users.id;

CREATE SEQUENCE request_id_seq;
CREATE TABLE request
(
    id      INTEGER   NOT NULL DEFAULT nextval('request_id_seq'),
    user_id INTEGER   NOT NULL REFERENCES users (id),
    text    VARCHAR   NOT NULL,
    time    TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE (id)
);
ALTER SEQUENCE request_id_seq OWNED BY request.id;

DROP TABLE IF EXISTS admins;
DROP SEQUENCE IF EXISTS admins_id_seq;

CREATE SEQUENCE admins_id_seq;
CREATE TABLE admins
(
    id       INTEGER NOT NULL DEFAULT nextval('admins_id_seq'),
    tg_id    INTEGER NOT NULL,
    tg_login VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    active   BOOLEAN NOT NULL,
    UNIQUE (tg_id),
    UNIQUE (id)
);
ALTER SEQUENCE admins_id_seq OWNED BY admins.id;

INSERT INTO public.admins (tg_id, tg_login, password, active)
VALUES (316671439, 'kbats183', '', true);

INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMcYSVqItKfOHkJxy7pYHa0byrLHGgAAlUAA88F4BKLRD6LRXJpFyAE', 'Коля Лаврентьев
У нас на этой неделе нет домашки?
Pavel Mavrin
Блин
Щас
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMdYSVqKGlxLwnwgi0-jaAVtphvaz0AAlYAA88F4BIYMJR820KB-CAE', 'Pavel Mavrin
казалось бы больше задач - больше
баллов
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMeYSVqKiP_cMoGZsMpQgO6gvZCs1gAAlcAA88F4BJlZVv5zZ450SAE', 'Артем Юсупов
как бан получим, так и узнаем
че гадать-то
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMfYSVqLBxGHtK7XeM3s4AB-mReH3wAAlgAA88F4BJwJLdwCN91kSAE', 'DaniI Bulanov
Я сделал конструкторы
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMgYSVqLRd-W3UrY7ZXOPtJP7DJENwAAlkAA88F4BKWpDlmsFktRSAE', 'Vladislav kovalchuk
Никакого в личку
Все вопросы через старост
А, так он и есть староста
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMhYSVqL0tJn1Zsw-Po8jU7T_g-Y0gAAloAA88F4BI-MjnB4tTiOyAE', 'Artem Vasilyev
религию новых заданий
комментировать не могу, но они
выложены
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMiYSVqMfgh6uQwaXMMRBavkHqHoY4AAlsAA88F4BI8LITZmEd1OCAE', 'Pavel Mavrin
Блин забыл
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMjYSVqM_ND0UKBxhV-94eAYDqXZZ8AAlwAA88F4BLXacGw9TIG_SAE', '6(д4)
Vitaly Aksenov
да, заебало уже
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMkYSVqNTUKLzVZM90f5_jQyU2xwloAAl0AA88F4BKEBCMaUMWwaiAE', 'Pavel Mavrin
ну вы пингуйте по кд. в какой-то
момент выложу)
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMlYSVqNx5Ae1S_mdf6eoXv8ImBcVAAAl4AA88F4BI66O2SQfkjriAE', 'Artem konkov
выходят и заходят
зачем они это делают?
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMmYSVqOYJuVEhPQ1LwldfBJKSuq8kAAl8AA88F4BLCWLB13_gIuSAE', 'GIeb Shanshin
как же хочется домашку, невинную,
никем не тронутую, с латеховскими
шрифтами....
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMnYSVqOyNd514tZDNJ_6j67uunk94AAmAAA88F4BKbb7ITrh-J1CAE', 'А
Alexander Trifanov
Пересдача будет:))
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMoYSVqPNT7k-yJAAERWYTE-xt6Qr1HAAJhAAPPBeASRUuEFBIkNUcgBA', 'Vitaly Aksenov
я - блядь
кому тут отсосать за 3300?
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMpYSVqPv4Nzvs_alc5ot-n62po9IMAAmIAA88F4BL1RunAlNQPvSAE', 'Vladislav kovalchuk
От контролируй до не контролируй
один шаг
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMqYSVqPzqTzFxnRZG4EQzpFx9OfD4AAmMAA88F4BKnoK7OLCGMoSAE', 'Darya Grechishkina
в какой-то момент познаешь дзен
(если дзен познан недостаточно
хорошо после этого тебя
отчисляют)
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMrYSVqQndOxMFxupCp6sQjYvAa1XsAAmQAA88F4BIax5PHu2UkDyAE', ',bTf
AIexander Dyuvenzhi
мастерство в том чтобы
балансировать между тем чтобы
отчислиться и быть отчисленным...
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMsYSVqROS-usxDAoIrFJ5WXm3avyAAAmUAA88F4BKR-9lgs6-KBiAE', 'Andrew Stankevich
че че это че
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMtYSVqRhGCPdfSkEUSLBFIw3AGsiIAAmYAA88F4BIpYfYcoMwMeCAE', 'А
Alexander Trifanov
Пересдача будет:))
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMuYSVqSMD-LZve39SZyE-ibNEfw-0AAmgAA88F4BKuyErg88q6RSAE', 'с
cpud36
У Вас ошибка
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMvYSVqS-NwLpjEE3GntOh0vjX7MiYAAmkAA88F4BI53TG4_hTHQSAE', 'Igor Podtsepko
Господи, какой же я тупой.
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMwYSVqT0PAHgXZ0e5o6m4lWsZf8soAAmoAA88F4BKYqqebg-3CDCAE', 'Igor Podtsepko
Страшно ДЗ8 сдавать.
Георгий Каданцев
Страшно сдавать ДЗ
Константин Бац
КБ
Страшно сдавать
Anatoly kochenyuk
Страшно
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMxYSVqUQT1mM6xJbgmiFiFSDFXa9kAAmsAA88F4BI2J2dDccSceSAE', 'Vladislav PADORU kuznetsov
радуйся нахуй
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMyYSVqVMwFEqbScfQN6YqFtLSVpTIAAmwAA88F4BIBJ4L-JDBPYCAE', 'N
Georgiy korneev
Требуется доброволец на практику
по транзакциям
Nikolay Vedernikov
это не чат БД
Georgiy korneev
Прошу прощения, это
действительно не чат по БД
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAMzYSVqV-1ggr1pJ2LEFeiUH7oew4YAAm0AA88F4BIED4i2_z8XIyAE', 'Andrew Stankevich
знаете почему не получили?
потому что бессмысленный вопрос
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAM0YSVqXMQS1gyb5oOXJtM5HGud6rgAAm4AA88F4BKkz5Yvlj2Q9SAE', 'Georgiy korneev
Спасибо
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAM1YSVqXwR-797CYEy2ydH_RRFMZHAAAm8AA88F4BIkGbtzHtVWZSAE', 'Georgiy korneev
И вам — спасибо
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAM2YSVqYfeHfw8jnXJgWx28ONfI6DgAAnAAA88F4BIP3akVRLP57SAE', 'Кирилл Пешков
Там мемная задача. Советую 10 раз
перечитать задание
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAM3YSVqY8XxH90GU0tDoENhKliG_JsAAnEAA88F4BIwXH3jfPgDbCAE', 'Roman Melnikov
');
INSERT INTO public.sticker (tg_set_name, tg_file_id, text_content)
VALUES ('CT_y2020_M3136_37', 'CAACAgIAAxkBAAM4YSVqZhgLT7LaDR3NG5AIahZfMp0AAnIAA88F4BKv-2VoX3ee7CAE', 'PaveI Mavrin
напиши на массиве и не
выпендривайся
');