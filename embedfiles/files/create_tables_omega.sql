create table if not exists datas(
	id integer primary key default 0,
   ot integer default 0,
	alias integer default 0,
	setcode integer default 0,
	type integer default 0,
	atk integer default 0,
	def integer default 0,
	level integer default 0,
	race integer default 0,
	attribute integer default 0,
	category integer default 0,
	genre integer default 0,
	script blob,
	support integer default 0,
	ocgdate integer default 253402207200,
	tcgdate integer default 253402207200
);

create table if not exists setcodes(
	officialcode integer,
	betacode integer,
	name text unique,
	cardid integer default 0
);

create table if not exists texts(
	id integer primary key,
	name text,
	desc text,
	str1 text,
	str2 text,
	str3 text,
	str4 text,
	str5 text,
	str6 text,
	str7 text,
	str8 text,
	str9 text,
	str10 text,
	str11 text,
	str12 text,
	str13 text,
	str14 text,
	str15 text,
	str16 text
);
