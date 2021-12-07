create table "user" (
	id int,
	name varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	password varchar(255) NOT NULL,
	age int NOT NULL,
	created_at DATE,
	updated_at DATE,
	
	constraint PK_User PRIMARY KEY (id)
);

create table "photo" (
	id int,
	title varchar(255) NOT NULL,
	caption varchar(255),
	photo_url varchar(255) NOT NULL,
	user_id int,
	created_at date,
	updated_at date,
	constraint PK_Photo PRIMARY KEY (id),
	constraint FK_UserPhoto FOREIGN KEY (user_id) REFERENCES "user"(id)
);

create table "comment"(
	id int,
	user_id int,
	photo_id int,
	message varchar(255) NOT NULL,
	created_at date,
	updated_at date,
	constraint PK_Comment PRIMARY KEY (id),
	constraint FK_UserComment FOREIGN KEY (user_id) REFERENCES "user"(id),
	constraint FK_PhotoComment FOREIGN KEY (photo_id) REFERENCES "photo"(id)
);

create table "socialmedia"(
	id int,
	name varchar(255) NOT NULL,
	social_media_url varchar(255) NOT NULL,
	user_id int,
	constraint PK_SocialMedia PRIMARY KEY (id),
	constraint FK_UserSocialMedia FOREIGN KEY (user_id) REFERENCES "user"(id)
);

