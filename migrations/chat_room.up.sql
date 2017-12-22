create table if not exists chat_room (
    id integer primary key not null
);

create table if not exists user_to_chat_room (
    user_id integer not null,
    chat_room_id integer not null,
    CONSTRAINT user_to_chat_room_user_fkey FOREIGN KEY (user_id)
      REFERENCES auth_user(id),
    CONSTRAINT user_to_chat_room_chat_room_fkey FOREIGN KEY (chat_room_id)
      REFERENCES chat_room(id),
);

create table if not exists user_to_message (
    user_id integer not null,
    message_id integer not null,
    CONSTRAINT user_to_message_user_fkey FOREIGN KEY (user_id)
      REFERENCES auth_user(id),
    CONSTRAINT user_to_message_message_fkey FOREIGN KEY (message_id)
      REFERENCES message(message_id),
);

create table if not exists message (
    id integer primary key not null,
    author_id integer not null,
    chat_room_id integer not null,
    text text,
    created_at datetime,
    CONSTRAINT message_user_fkey FOREIGN KEY (author_id)
      REFERENCES auth_user(id),
    CONSTRAINT message_chat_room_fkey FOREIGN KEY (chat_room_id)
      REFERENCES chat_room(id),
);