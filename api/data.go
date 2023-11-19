package handler

import (
	"io"
	"net/http"
)

func DummyHandle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello :) \n")
}

var ChannelList = `
[
  {
    "name": "Somnath Temple",
    "link": "https://rogstream.fun//hls/tata/play.php?id=842",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-56389-kfdgngts-v3/imageContent-56389-kfdgngts-m3.png",
    "genre": "Spiritual"
  },
  {
    "name": "Shirdi Sai Baba",
    "link": "https://rogstream.fun//hls/tata/play.php?id=840",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-56386-kfc14w60-v4/imageContent-56386-kfc14w60-m4.png",
    "genre": "Spiritual"
  },
  {
    "name": "Tata Play Fitness",
    "link": "https://rogstream.fun//hls/tata/play.php?id=121",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-405-j5jr3sz4-v2/imageContent-405-j5jr3sz4-m2.png",
    "genre": "Lifestyle"
  },
  {
    "name": "Tata Play Cooking",
    "link": "https://rogstream.fun//hls/tata/play.php?id=641",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-23495-jf92iycg-v3/imageContent-23495-jf92iycg-m4.png",
    "genre": "Lifestyle"
  },
  {
    "name": "Tata Play Beauty",
    "link": "https://rogstream.fun//hls/tata/play.php?id=618",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12282-ja02jlp4-v2/imageContent-12282-ja02jlp4-m2.png",
    "genre": "Lifestyle"
  },
  {
    "name": "DD National",
    "link": "https://rogstream.fun//hls/tata/play.php?id=191",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-795-j5m7axrc-v1/imageContent-795-j5m7axrc-m1.png",
    "genre": "Entertainment"
  },
  {
    "name": "Tata Play Romance",
    "link": "https://rogstream.fun//hls/tata/play.php?id=959",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-67692-kwmxxw08-v2/imageContent-67692-kwmxxw08-m3.png",
    "genre": "Entertainment"
  },
  {
    "name": "SET HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=15",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-43-j5fca4k0-v3/imageContent-43-j5fca4k0-m4.png",
    "genre": "Entertainment"
  },
  {
    "name": "STAR Plus HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=8",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-25307-jhrhflww-v1/imageContent-25307-jhrhflww-m1.png",
    "genre": "Entertainment"
  },
  {
    "name": "Star Bharat HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=244",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-6895-j6vqhqnc-v2/imageContent-6895-j6vqhqnc-m2.png",
    "genre": "Entertainment"
  },
  {
    "name": "SONY SAB HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=48",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-141-j5fpeji0-v3/imageContent-141-j5fpeji0-m3.png",
    "genre": "Entertainment"
  },
  {
    "name": "SET",
    "link": "https://rogstream.fun//hls/tata/play.php?id=556",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12074-j9oat6qw-v6/imageContent-12074-j9oat6qw-m6.png",
    "genre": "Entertainment"
  },
  {
    "name": "\u0026TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=578",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12129-j9p7kluo-v2/imageContent-12129-j9p7kluo-m3.png",
    "genre": "Entertainment"
  },
  {
    "name": "SONY SAB",
    "link": "https://rogstream.fun//hls/tata/play.php?id=559",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12081-j9oc38xc-v8/imageContent-12081-j9oc38xc-m7.png",
    "genre": "Entertainment"
  },
  {
    "name": "\u0026tv HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=40",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-117-j5fl7440-v1/imageContent-117-j5fl7440-m1.png",
    "genre": "Entertainment"
  },
  {
    "name": "Zee TV HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=63",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11073-j95e7nyo-v1/imageContent-11073-j95e7nyo-m1.png",
    "genre": "Entertainment"
  },
  {
    "name": "Zee TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=557",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12077-j9ob4gm0-v2/imageContent-12077-j9ob4gm0-m2.png",
    "genre": "Entertainment"
  },
  {
    "name": "Colors HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=52",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-155-j5frd2uo-v1/imageContent-155-j5frd2uo-m1.png",
    "genre": "Entertainment"
  },
  {
    "name": "Colors",
    "link": "https://rogstream.fun//hls/tata/play.php?id=543",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12053-j9o3yhko-v1/imageContent-12053-j9o3yhko-m1.png",
    "genre": "Entertainment"
  },
  {
    "name": "Tata Play Javed Akhtar",
    "link": "https://rogstream.fun//hls/tata/play.php?id=95",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-304-j5ji9otc-v2/imageContent-304-j5ji9otc-m2.png",
    "genre": "Entertainment"
  },
  {
    "name": "Tata Play Videshi Kahaniyan",
    "link": "https://rogstream.fun//hls/tata/play.php?id=943",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-65443-kty2tkwo-v2/imageContent-65443-kty2tkwo-m2.png",
    "genre": "Entertainment"
  },
  {
    "name": "UTV Bindass",
    "link": "https://rogstream.fun//hls/tata/play.php?id=142",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-475-j5jx44s8-v1/imageContent-475-j5jx44s8-m1.png",
    "genre": "Entertainment"
  },
  {
    "name": "Tata Play Zindagi",
    "link": "https://rogstream.fun//hls/tata/play.php?id=986",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-70821-l3edt8uw-v1/imageContent-70821-l3edt8uw-m3.png",
    "genre": "null"
  },
  {
    "name": "Investigation Discovery",
    "link": "https://rogstream.fun//hls/tata/play.php?id=633",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-49002-k5cch1cg-v2/imageContent-49002-k5cch1cg-m5.png",
    "genre": "Entertainment"
  },
  {
    "name": "Naaptol",
    "link": "https://rogstream.fun//hls/tata/play.php?id=619",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12285-ja1e8580-v3/imageContent-12285-ja1e8580-m3.png",
    "genre": "Shopping"
  },
  {
    "name": "Ezmall",
    "link": "https://rogstream.fun//hls/tata/play.php?id=634",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-21303-je858doo-v2/imageContent-21303-je858doo-m2.png",
    "genre": "Shopping"
  },
  {
    "name": "STAR Utsav",
    "link": "https://rogstream.fun//hls/tata/play.php?id=551",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12060-j9o91vfc-v2/imageContent-12060-j9o91vfc-m2.png",
    "genre": "Entertainment"
  },
  {
    "name": "Zee Anmol",
    "link": "https://rogstream.fun//hls/tata/play.php?id=523",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11969-j9luigc0-v2/imageContent-11969-j9luigc0-m2.png",
    "genre": "Entertainment"
  },
  {
    "name": "Colors Rishtey",
    "link": "https://rogstream.fun//hls/tata/play.php?id=438",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/RIS_Thumbnail-v3/RIS_Thumbnail.png",
    "genre": "Entertainment"
  },
  {
    "name": "Sony Pal",
    "link": "https://rogstream.fun//hls/tata/play.php?id=554",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12071-j9oa4gbc-v5/imageContent-12071-j9oa4gbc-m4.png",
    "genre": "Entertainment"
  },
  {
    "name": "The Q",
    "link": "https://rogstream.fun//hls/tata/play.php?id=54",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-150-j5frd0jc-v3/imageContent-150-j5frd0jc-m5.png",
    "genre": "Entertainment"
  },
  {
    "name": "Big Magic",
    "link": "https://rogstream.fun//hls/tata/play.php?id=297",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11266-j9j2spmg-v1/imageContent-11266-j9j2spmg-m1.png",
    "genre": "Entertainment"
  },
  {
    "name": "Dangal",
    "link": "https://rogstream.fun//hls/tata/play.php?id=51",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-158-j5frd560-v2/imageContent-158-j5frd560-m3.png",
    "genre": "Entertainment"
  },
  {
    "name": "Anjan TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=180",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-765-j5m2mbjk-v2/imageContent-765-j5m2mbjk-m2.png",
    "genre": "Entertainment"
  },
  {
    "name": "Shemaroo TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=818",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-52945-ka1wdss8-v2/imageContent-52945-ka1wdss8-m2.png",
    "genre": "Entertainment"
  },
  {
    "name": "Ishara",
    "link": "https://rogstream.fun//hls/tata/play.php?id=888",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-60986-klq1vs5s-v2/imageContent-60986-klq1vs5s-m5.png",
    "genre": "Entertainment"
  },
  {
    "name": "Shemaroo UMANG",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1045",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/SHEUM_Thumbnail-v2/SHEUM_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "Atrangii",
    "link": "https://rogstream.fun//hls/tata/play.php?id=990",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/ATRAN_Thumbnail-v6/ATRAN_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "DD Kisan",
    "link": "https://rogstream.fun//hls/tata/play.php?id=326",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11350-j9jpk8pk-v1/imageContent-11350-j9jpk8pk-m1.png",
    "genre": "Knowledge"
  },
  {
    "name": "Tata Play Hits",
    "link": "https://rogstream.fun//hls/tata/play.php?id=808",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-49809-k7454ppk-v3/imageContent-49809-k7454ppk-m6.png",
    "genre": "Entertainment"
  },
  {
    "name": "Colors Infinity HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=187",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-1577-j5xrm7fc-v1/imageContent-1577-j5xrm7fc-m1.png",
    "genre": "Entertainment"
  },
  {
    "name": "Comedy Central HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=307",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11296-j9j5ux7s-v2/imageContent-11296-j9j5ux7s-m3.png",
    "genre": "Entertainment"
  },
  {
    "name": "Comedy Central",
    "link": "https://rogstream.fun//hls/tata/play.php?id=306",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11293-j9j5u4o0-v3/imageContent-11293-j9j5u4o0-m2.png",
    "genre": "Entertainment"
  },
  {
    "name": "Colors Infinity",
    "link": "https://rogstream.fun//hls/tata/play.php?id=544",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-18305-jcqmvp7s-v1/imageContent-18305-jcqmvp7s-m1.png",
    "genre": "Entertainment"
  },
  {
    "name": "Tata Play Hollywood Local",
    "link": "https://rogstream.fun//hls/tata/play.php?id=789",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-48748-k4jbpl00-v2/imageContent-48748-k4jbpl00-m3.png",
    "genre": "Movies"
  },
  {
    "name": "STAR GOLD HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=245",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/SGHD_Thumbnail-v2/SGHD_Thumbnail.png",
    "genre": "Movies"
  },
  {
    "name": "SONY MAX HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=80",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-243-j5fyl2f4-v3/imageContent-243-j5fyl2f4-m4.png",
    "genre": "Movies"
  },
  {
    "name": "SONY MAX",
    "link": "https://rogstream.fun//hls/tata/play.php?id=132",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-441-j5jt3mmg-v3/imageContent-441-j5jt3mmg-m3.png",
    "genre": "Movies"
  },
  {
    "name": "Zee Cinema HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=503",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11915-j9l5clzs-v1/imageContent-11915-j9l5clzs-m1.png",
    "genre": "Movies"
  },
  {
    "name": "Zee Cinema",
    "link": "https://rogstream.fun//hls/tata/play.php?id=123",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11027-j8yjr8k0-v1/imageContent-11027-j8yjr8k0-m1.png",
    "genre": "Movies"
  },
  {
    "name": "Zee Bollywood",
    "link": "https://rogstream.fun//hls/tata/play.php?id=175",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-31233-jli1wlvc-v1/imageContent-31233-jli1wlvc-m1.png",
    "genre": "Movies"
  },
  {
    "name": "Zee Classic",
    "link": "https://rogstream.fun//hls/tata/play.php?id=727",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-42199-jxh2zlj4-v1/imageContent-42199-jxh2zlj4-m1.png",
    "genre": "Movies"
  },
  {
    "name": "Tata Play Theatre HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=666",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-31419-jmb48u74-v3/imageContent-31419-jmb48u74-m4.png",
    "genre": "Movies"
  },
  {
    "name": "\u0026pictures HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=267",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11173-j9hth720-v1/imageContent-11173-j9hth720-m1.png",
    "genre": "Movies"
  },
  {
    "name": "\u0026pictures",
    "link": "https://rogstream.fun//hls/tata/play.php?id=148",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-499-j5jydp7s-v1/imageContent-499-j5jydp7s-m1.png",
    "genre": "Movies"
  },
  {
    "name": "Tata Play ShortsTV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=677",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-34170-jok0hf00-v7/imageContent-34170-jok0hf00-m8.PNG",
    "genre": "Movies"
  },
  {
    "name": "Colors Cineplex HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=61",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-183-j5fsqz5c-v2/imageContent-183-j5fsqz5c-m2.png",
    "genre": "Movies"
  },
  {
    "name": "Colors Cineplex",
    "link": "https://rogstream.fun//hls/tata/play.php?id=53",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-149-j5frd0jc-v2/imageContent-149-j5frd0jc-m2.png",
    "genre": "Movies"
  },
  {
    "name": "B4U Kadak",
    "link": "https://rogstream.fun//hls/tata/play.php?id=730",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-43104-jxzls8yw-v2/imageContent-43104-jxzls8yw-m2.png",
    "genre": "Movies"
  },
  {
    "name": "Zee Action",
    "link": "https://rogstream.fun//hls/tata/play.php?id=100",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11024-j8yix4g8-v1/imageContent-11024-j8yix4g8-m1.png",
    "genre": "Movies"
  },
  {
    "name": "Tata Play South Talkies",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1003",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-72254-l6dn8nco-v3/imageContent-72254-l6dn8nco-m4.png",
    "genre": "null"
  },
  {
    "name": "Sony Wah",
    "link": "https://rogstream.fun//hls/tata/play.php?id=56",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-198-j5fsr7mw-v2/imageContent-198-j5fsr7mw-m3.png",
    "genre": "Movies"
  },
  {
    "name": "B4U Movies",
    "link": "https://rogstream.fun//hls/tata/play.php?id=7",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-18-j5f9ui8g-v1/imageContent-18-j5f9ui8g-m1.png",
    "genre": "Movies"
  },
  {
    "name": "SONY MAX 2",
    "link": "https://rogstream.fun//hls/tata/play.php?id=120",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-402-j5jq1gko-v3/imageContent-402-j5jq1gko-m4.png",
    "genre": "Movies"
  },
  {
    "name": "Cinema TV India",
    "link": "https://rogstream.fun//hls/tata/play.php?id=182",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-771-j5m3ayw0-v4/imageContent-771-j5m3ayw0-m4.png",
    "genre": "Movies"
  },
  {
    "name": "WoW Cinema One",
    "link": "https://rogstream.fun//hls/tata/play.php?id=276",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11200-j9i4gzig-v3/imageContent-11200-j9i4gzig-m4.png",
    "genre": "Movies"
  },
  {
    "name": "Dangal 2",
    "link": "https://rogstream.fun//hls/tata/play.php?id=194",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-798-j5m7esnc-v4/imageContent-798-j5m7esnc-m11.png",
    "genre": "Movies"
  },
  {
    "name": "BFLIX",
    "link": "https://rogstream.fun//hls/tata/play.php?id=296",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11263-j9j1tch4-v2/imageContent-11263-j9j1tch4-m5.png",
    "genre": "Movies"
  },
  {
    "name": "Zee Anmol Cinema",
    "link": "https://rogstream.fun//hls/tata/play.php?id=64",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11090-j95hdh6o-v1/imageContent-11090-j95hdh6o-m1.png",
    "genre": "Movies"
  },
  {
    "name": "Goldmines",
    "link": "https://rogstream.fun//hls/tata/play.php?id=823",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-53943-kblsqraw-v3/imageContent-53943-kblsqraw-m4.png",
    "genre": "Marathi"
  },
  {
    "name": "Manoranjan TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=731",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-43106-jxzlt6x4-v2/imageContent-43106-jxzlt6x4-m7.png",
    "genre": "Movies"
  },
  {
    "name": "Manoranjan Grand",
    "link": "https://rogstream.fun//hls/tata/play.php?id=965",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-67799-kx7m7rhc-v1/imageContent-67799-kx7m7rhc-m2.png",
    "genre": "Movies"
  },
  {
    "name": "Movies Now",
    "link": "https://rogstream.fun//hls/tata/play.php?id=173",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-748-j5l5fnts-v1/imageContent-748-j5l5fnts-m1.png",
    "genre": "Movies"
  },
  {
    "name": "Colors Cineplex Bollywood",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1000",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-71752-l5hqgxso-v2/imageContent-71752-l5hqgxso-m1.png",
    "genre": "null"
  },
  {
    "name": "MNX",
    "link": "https://rogstream.fun//hls/tata/play.php?id=234",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-1017-j5ngpo2g-v1/imageContent-1017-j5ngpo2g-m1.png",
    "genre": "Movies"
  },
  {
    "name": "Dhamaka Movies B4U",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1037",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/DHAMB4U_Thumbnail-v2/DHAMB4U_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "SONY PIX HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=32",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-93-j5fjv614-v3/imageContent-93-j5fjv614-m3.png",
    "genre": "Movies"
  },
  {
    "name": "SONY PIX",
    "link": "https://rogstream.fun//hls/tata/play.php?id=558",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12080-j9obub60-v6/imageContent-12080-j9obub60-m6.png",
    "genre": "Movies"
  },
  {
    "name": "Movies Now HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=562",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12095-j9ooixfs-v1/imageContent-12095-j9ooixfs-m1.png",
    "genre": "Movies"
  },
  {
    "name": "MN+ HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=210",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-826-j5m9kx5c-v1/imageContent-826-j5m9kx5c-m1.png",
    "genre": "Movies"
  },
  {
    "name": "MNX HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=599",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12216-j9qm5hjs-v2/imageContent-12216-j9qm5hjs-m2.png",
    "genre": "Movies"
  },
  {
    "name": "Romedy Now",
    "link": "https://rogstream.fun//hls/tata/play.php?id=174",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-751-j5l5kb9k-v3/imageContent-751-j5l5kb9k-m3.png",
    "genre": "Movies"
  },
  {
    "name": "DD Sports",
    "link": "https://rogstream.fun//hls/tata/play.php?id=223",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-941-j5mhk5fc-v1/imageContent-941-j5mhk5fc-m1.png",
    "genre": "Sports"
  },
  {
    "name": "Star Sports Select 1 HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=246",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-7891-j75vq7k0-v1/imageContent-7891-j75vq7k0-m1.PNG",
    "genre": "Sports"
  },
  {
    "name": "Star Sports Select 2 HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=463",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11791-j9kyqpy8-v1/imageContent-11791-j9kyqpy8-m1.png",
    "genre": "Sports"
  },
  {
    "name": "SONY SPORTS TEN 1 HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=81",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-261-j5fz9bvk-v4/imageContent-261-j5fz9bvk-m4.png",
    "genre": "Sports"
  },
  {
    "name": "SONY SPORTS TEN 1",
    "link": "https://rogstream.fun//hls/tata/play.php?id=150",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-513-j5jzouqw-v5/imageContent-513-j5jzouqw-m7.png",
    "genre": "Sports"
  },
  {
    "name": "SONY SPORTS TEN 2 HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=462",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11784-j9kylfd4-v2/imageContent-11784-j9kylfd4-m3.png",
    "genre": "Sports"
  },
  {
    "name": "SONY SPORTS TEN 2",
    "link": "https://rogstream.fun//hls/tata/play.php?id=170",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-742-j5l50le0-v5/imageContent-742-j5l50le0-m7.png",
    "genre": "Sports"
  },
  {
    "name": "SONY SPORTS TEN 3 HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=464",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11787-j9kynnm0-v3/imageContent-11787-j9kynnm0-m3.png",
    "genre": "Sports"
  },
  {
    "name": "SONY SPORTS TEN 3",
    "link": "https://rogstream.fun//hls/tata/play.php?id=171",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-739-j5l4y6yw-v5/imageContent-739-j5l4y6yw-m5.png",
    "genre": "Sports"
  },
  {
    "name": "SONY SPORTS TEN 5 HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=35",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-108-j5fl5wwg-v3/imageContent-108-j5fl5wwg-m3.png",
    "genre": "Sports"
  },
  {
    "name": "Sports 18 - 1 HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1033",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-74232-la6z41u8-v3/imageContent-74232-la6z41u8-m5.png",
    "genre": "null"
  },
  {
    "name": "Sports 18 - 1",
    "link": "https://rogstream.fun//hls/tata/play.php?id=980",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-70046-l21cybww-v2/imageContent-70046-l21cybww-m5.png",
    "genre": "Sports"
  },
  {
    "name": "Eurosport",
    "link": "https://rogstream.fun//hls/tata/play.php?id=693",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-35309-jqxwmcq0-v4/imageContent-35309-jqxwmcq0-m4.png",
    "genre": "Sports"
  },
  {
    "name": "Eurosport HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=812",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-52842-k980h8vc-v1/imageContent-52842-k980h8vc-m2.png",
    "genre": "Sports"
  },
  {
    "name": "1sports",
    "link": "https://rogstream.fun//hls/tata/play.php?id=815",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-52856-k984c4vc-v1/imageContent-52856-k984c4vc-m1.png",
    "genre": "Sports"
  },
  {
    "name": "DD News HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=691",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-34886-jqc7dybs-v2/imageContent-34886-jqc7dybs-m2.png",
    "genre": "News"
  },
  {
    "name": "DD News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=322",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11333-j9joxlw8-v2/imageContent-11333-j9joxlw8-m2.png",
    "genre": "News"
  },
  {
    "name": "ABP News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=177",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-759-j5m13014-v2/imageContent-759-j5m13014-m3.png",
    "genre": "News"
  },
  {
    "name": "Tata Play Seniors",
    "link": "https://rogstream.fun//hls/tata/play.php?id=783",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-47647-k2lbp2q8-v2/imageContent-47647-k2lbp2q8-m2.png",
    "genre": "Knowledge"
  },
  {
    "name": "NDTV India",
    "link": "https://rogstream.fun//hls/tata/play.php?id=179",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-483-j5jx9x48-v1/imageContent-483-j5jx9x48-m1.png",
    "genre": "News"
  },
  {
    "name": "Aaj Tak HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=689",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-34652-jpmey8gw-v1/imageContent-34652-jpmey8gw-m1.png",
    "genre": "News"
  },
  {
    "name": "Aaj Tak",
    "link": "https://rogstream.fun//hls/tata/play.php?id=153",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-537-j5kwkd88-v1/imageContent-537-j5kwkd88-m1.png",
    "genre": "News"
  },
  {
    "name": "Tata Play Har Ghar Startup",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1056",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/HAGHAR_Thumbnail-v1/HAGHAR_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "Zee News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=259",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11110-j95ioyag-v1/imageContent-11110-j95ioyag-m1.png",
    "genre": "News"
  },
  {
    "name": "Tata Play Astro Duniya",
    "link": "https://rogstream.fun//hls/tata/play.php?id=872",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-58613-kj283ly8-v2/imageContent-58613-kj283ly8-m2.png",
    "genre": "Spiritual"
  },
  {
    "name": "India TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=104",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-8568-j7lnde80-v1/imageContent-8568-j7lnde80-m1.png",
    "genre": "News"
  },
  {
    "name": "News 24",
    "link": "https://rogstream.fun//hls/tata/play.php?id=209",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-849-j5m9tmf4-v1/imageContent-849-j5m9tmf4-m1.png",
    "genre": "News"
  },
  {
    "name": "Tata Play JEE Prep",
    "link": "https://rogstream.fun//hls/tata/play.php?id=929",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-64069-kra2o75k-v2/imageContent-64069-kra2o75k-m2.png",
    "genre": "Knowledge"
  },
  {
    "name": "Tata Play NEET Prep",
    "link": "https://rogstream.fun//hls/tata/play.php?id=928",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-64066-kra2j2qo-v2/imageContent-64066-kra2j2qo-m2.png",
    "genre": "Knowledge"
  },
  {
    "name": "News18 India",
    "link": "https://rogstream.fun//hls/tata/play.php?id=106",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-339-j5jm6kko-v1/imageContent-339-j5jm6kko-m1.png",
    "genre": "News"
  },
  {
    "name": "Zee Hindustan",
    "link": "https://rogstream.fun//hls/tata/play.php?id=514",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11978-j9lum7d4-v1/imageContent-11978-j9lum7d4-m1.png",
    "genre": "News"
  },
  {
    "name": "R Bharat",
    "link": "https://rogstream.fun//hls/tata/play.php?id=696",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-35599-jrm11j9c-v1/imageContent-35599-jrm11j9c-m1.PNG",
    "genre": "Hindi"
  },
  {
    "name": "India News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=183",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-774-j5m3h6nk-v1/imageContent-774-j5m3h6nk-m1.png",
    "genre": "News"
  },
  {
    "name": "News Nation",
    "link": "https://rogstream.fun//hls/tata/play.php?id=36",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-114-j5fl632o-v1/imageContent-114-j5fl632o-m1.png",
    "genre": "News"
  },
  {
    "name": "TV9 Bharatvarsh",
    "link": "https://rogstream.fun//hls/tata/play.php?id=706",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-67672-kwly0m4w-v1/imageContent-67672-kwly0m4w-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "TOTAL TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=485",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11859-j9l2n36w-v2/imageContent-11859-j9l2n36w-m2.png",
    "genre": "News"
  },
  {
    "name": "Times Now Navbharat HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=932",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-64606-krslbshs-v1/imageContent-64606-krslbshs-m2.png",
    "genre": "News"
  },
  {
    "name": "Times Now Navbharat",
    "link": "https://rogstream.fun//hls/tata/play.php?id=966",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-67796-kx7lokw8-v2/imageContent-67796-kx7lokw8-m4.png",
    "genre": "News"
  },
  {
    "name": "News India 24x7",
    "link": "https://rogstream.fun//hls/tata/play.php?id=831",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-55547-kdzeosow-v3/imageContent-55547-kdzeosow-m6.png",
    "genre": "News"
  },
  {
    "name": "Bharat 24",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1010",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/BH24_Thumbnail-v3/BH24_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "Surya Samachar",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1031",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-72980-l8fn3f28-v2/imageContent-72980-l8fn3f28-m4.png",
    "genre": "null"
  },
  {
    "name": "Sudarshan News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=475",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11831-j9kzkn40-v1/imageContent-11831-j9kzkn40-m1.png",
    "genre": "News"
  },
  {
    "name": "Bharat Express",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1051",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/BHEXP_Thumbnail-v2/BHEXP_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "India Daily Live",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1084",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/INDLive_Thumbnail-v1/INDLive_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "APN News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=286",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-17573-jclve4rc-v1/imageContent-17573-jclve4rc-m1.png",
    "genre": "News"
  },
  {
    "name": "ET Now Swadesh",
    "link": "https://rogstream.fun//hls/tata/play.php?id=950",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-66223-ku9gfjug-v1/imageContent-66223-ku9gfjug-m1.png",
    "genre": "News"
  },
  {
    "name": "India Voice",
    "link": "https://rogstream.fun//hls/tata/play.php?id=685",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-34516-jp9bj0qg-v4/imageContent-34516-jp9bj0qg-m4.png",
    "genre": "News"
  },
  {
    "name": "News 1 India",
    "link": "https://rogstream.fun//hls/tata/play.php?id=414",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11634-j9k0pbr4-v1/imageContent-11634-j9k0pbr4-m1.png",
    "genre": "News"
  },
  {
    "name": "Hindi Khabar",
    "link": "https://rogstream.fun//hls/tata/play.php?id=344",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11467-j9jr3k80-v2/imageContent-11467-j9jr3k80-m2.png",
    "genre": "News"
  },
  {
    "name": "Bharat Samachar",
    "link": "https://rogstream.fun//hls/tata/play.php?id=291",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11243-j9itj54w-v2/imageContent-11243-j9itj54w-m2.png",
    "genre": "News"
  },
  {
    "name": "Samay",
    "link": "https://rogstream.fun//hls/tata/play.php?id=441",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11719-j9ktg6ds-v2/imageContent-11719-j9ktg6ds-m2.png",
    "genre": "News"
  },
  {
    "name": "Network 10",
    "link": "https://rogstream.fun//hls/tata/play.php?id=848",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-56488-kfgbqzso-v3/imageContent-56488-kfgbqzso-m5.png",
    "genre": "News"
  },
  {
    "name": "Jantantra TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=701",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-37210-jswmlty0-v2/imageContent-37210-jswmlty0-m2.PNG",
    "genre": "Hindi"
  },
  {
    "name": "Prime News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=869",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-58484-kixj8h6o-v3/imageContent-58484-kixj8h6o-m2.png",
    "genre": "News"
  },
  {
    "name": "Har Khabar",
    "link": "https://rogstream.fun//hls/tata/play.php?id=988",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-70993-l3v760zc-v1/imageContent-70993-l3v760zc-m4.png",
    "genre": "null"
  },
  {
    "name": "FM News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=784",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/FNEWS_Thumbnail-v2/FNEWS_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "Zee Business",
    "link": "https://rogstream.fun//hls/tata/play.php?id=260",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11120-j95jyuao-v1/imageContent-11120-j95jyuao-m1.png",
    "genre": "News"
  },
  {
    "name": "Good News Today",
    "link": "https://rogstream.fun//hls/tata/play.php?id=278",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11206-j9icarsw-v2/imageContent-11206-j9icarsw-m2.png",
    "genre": "News"
  },
  {
    "name": "CNBC Awaaz",
    "link": "https://rogstream.fun//hls/tata/play.php?id=204",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-846-j5m9p55k-v1/imageContent-846-j5m9p55k-m1.png",
    "genre": "News"
  },
  {
    "name": "Sansad TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=213",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-825-j5m9c6c0-v2/imageContent-825-j5m9c6c0-m2.png",
    "genre": "News"
  },
  {
    "name": "Sansad TV - Rajya Sabha",
    "link": "https://rogstream.fun//hls/tata/play.php?id=212",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-858-j5ma94ag-v2/imageContent-858-j5ma94ag-m2.png",
    "genre": "News"
  },
  {
    "name": "NDTV 24x7",
    "link": "https://rogstream.fun//hls/tata/play.php?id=208",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-834-j5m9nrrs-v1/imageContent-834-j5m9nrrs-m1.png",
    "genre": "News"
  },
  {
    "name": "Times Now World",
    "link": "https://rogstream.fun//hls/tata/play.php?id=547",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11853-j9l1tvhs-v4/imageContent-11853-j9l1tvhs-m5.png",
    "genre": "News"
  },
  {
    "name": "TIMES NOW",
    "link": "https://rogstream.fun//hls/tata/play.php?id=90",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-303-j5ji8zco-v3/imageContent-303-j5ji8zco-m3.png",
    "genre": "News"
  },
  {
    "name": "INDIA TODAY",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-3-j5dkuhwo-v1/imageContent-3-j5dkuhwo-m2.png",
    "genre": "News"
  },
  {
    "name": "CNN News18",
    "link": "https://rogstream.fun//hls/tata/play.php?id=206",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-843-j5m9oyzc-v1/imageContent-843-j5m9oyzc-m1.png",
    "genre": "News"
  },
  {
    "name": "NewsX",
    "link": "https://rogstream.fun//hls/tata/play.php?id=189",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-8012-j7a2doxc-v1/imageContent-8012-j7a2doxc-m1.png",
    "genre": "English News"
  },
  {
    "name": "MIRROR NOW",
    "link": "https://rogstream.fun//hls/tata/play.php?id=591",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12189-j9plqlf4-v1/imageContent-12189-j9plqlf4-m1.png",
    "genre": "News"
  },
  {
    "name": "WION",
    "link": "https://rogstream.fun//hls/tata/play.php?id=255",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-9000-j7rkz0m0-v1/imageContent-9000-j7rkz0m0-m1.png",
    "genre": "News"
  },
  {
    "name": "REPUBLIC TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=72",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-231-j5fxh0ag-v1/imageContent-231-j5fxh0ag-m1.png",
    "genre": "News"
  },
  {
    "name": "CNBC TV18 Prime HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=304",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11284-j9j4sn4o-v2/imageContent-11284-j9j4sn4o-m2.png",
    "genre": "News"
  },
  {
    "name": "DD India",
    "link": "https://rogstream.fun//hls/tata/play.php?id=324",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11344-j9jpg7nc-v1/imageContent-11344-j9jpg7nc-m1.png",
    "genre": "News"
  },
  {
    "name": "NDTV Profit Prime",
    "link": "https://rogstream.fun//hls/tata/play.php?id=93",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-294-j5ji8n08-v1/imageContent-294-j5ji8n08-m1.png",
    "genre": "News"
  },
  {
    "name": "CNBC-TV18",
    "link": "https://rogstream.fun//hls/tata/play.php?id=168",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-731-j5l3muhs-v1/imageContent-731-j5l3muhs-m1.png",
    "genre": "News"
  },
  {
    "name": "ET NOW",
    "link": "https://rogstream.fun//hls/tata/play.php?id=88",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-6904-j6vreaps-v1/imageContent-6904-j6vreaps-m1.png",
    "genre": "News"
  },
  {
    "name": "CNN",
    "link": "https://rogstream.fun//hls/tata/play.php?id=243",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-4154-j69pdmvk-v1/imageContent-4154-j69pdmvk-m1.png",
    "genre": "News"
  },
  {
    "name": "BBC News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=188",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/BBCWNW_Thumbnail-v3/BBCWNW_Thumbnail.png",
    "genre": "News"
  },
  {
    "name": "Al Jazeera",
    "link": "https://rogstream.fun//hls/tata/play.php?id=190",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-789-j5m5m0uw-v1/imageContent-789-j5m5m0uw-m1.png",
    "genre": "News"
  },
  {
    "name": "Channel News Asia",
    "link": "https://rogstream.fun//hls/tata/play.php?id=157",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-548-j5kxtaig-v1/imageContent-548-j5kxtaig-m1.png",
    "genre": "News"
  },
  {
    "name": "France 24",
    "link": "https://rogstream.fun//hls/tata/play.php?id=131",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-435-j5jsyz6o-v1/imageContent-435-j5jsyz6o-m1.png",
    "genre": "News"
  },
  {
    "name": "TV5 Monde Asie",
    "link": "https://rogstream.fun//hls/tata/play.php?id=65",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-204-j5ftndtc-v1/imageContent-204-j5ftndtc-m1.png",
    "genre": "News"
  },
  {
    "name": "DW",
    "link": "https://rogstream.fun//hls/tata/play.php?id=60",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-186-j5fsr1go-v1/imageContent-186-j5fsr1go-m1.png",
    "genre": "News"
  },
  {
    "name": "Russia Today",
    "link": "https://rogstream.fun//hls/tata/play.php?id=98",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-312-j5jjsfhs-v1/imageContent-312-j5jjsfhs-m1.png",
    "genre": "News"
  },
  {
    "name": "NHK World Japan",
    "link": "https://rogstream.fun//hls/tata/play.php?id=976",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-69827-l1dh5izk-v1/imageContent-69827-l1dh5izk-m1.png",
    "genre": "News"
  },
  {
    "name": "Super Hungama",
    "link": "https://rogstream.fun//hls/tata/play.php?id=587",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-35365-jr3xz4uo-v5/imageContent-35365-jr3xz4uo-m6.png",
    "genre": "Kids"
  },
  {
    "name": "Disney",
    "link": "https://rogstream.fun//hls/tata/play.php?id=114",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-1356-j5tdnwmw-v1/imageContent-1356-j5tdnwmw-m1.png",
    "genre": "Kids"
  },
  {
    "name": "Australia TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=105",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-324-j5jl5hp4-v2/imageContent-324-j5jl5hp4-m2.png",
    "genre": "News"
  },
  {
    "name": "Tata Play English in Hindi",
    "link": "https://rogstream.fun//hls/tata/play.php?id=71",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-25721-jia4i70o-v2/imageContent-25721-jia4i70o-m3.PNG",
    "genre": "Knowledge"
  },
  {
    "name": "Nick",
    "link": "https://rogstream.fun//hls/tata/play.php?id=138",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-465-j5jw9va0-v1/imageContent-465-j5jw9va0-m1.png",
    "genre": "Kids"
  },
  {
    "name": "Nick HD+",
    "link": "https://rogstream.fun//hls/tata/play.php?id=433",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11658-j9k5lvgo-v1/imageContent-11658-j9k5lvgo-m1.png",
    "genre": "Kids"
  },
  {
    "name": "Tata Play Fun Learn Preschool",
    "link": "https://rogstream.fun//hls/tata/play.php?id=626",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-13094-jbdzlsfs-v2/imageContent-13094-jbdzlsfs-m2.png",
    "genre": "Kids"
  },
  {
    "name": "Tata Play Fun Learn Junior",
    "link": "https://rogstream.fun//hls/tata/play.php?id=627",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-13097-jbdzn9og-v2/imageContent-13097-jbdzn9og-m2.png",
    "genre": "Kids"
  },
  {
    "name": "Cartoon Network HD+",
    "link": "https://rogstream.fun//hls/tata/play.php?id=681",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-34164-jok06xoo-v1/imageContent-34164-jok06xoo-m1.png",
    "genre": "Kids"
  },
  {
    "name": "Cartoon Network",
    "link": "https://rogstream.fun//hls/tata/play.php?id=238",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11161-j99h67u8-v1/imageContent-11161-j99h67u8-m1.png",
    "genre": "Kids"
  },
  {
    "name": "Pogo",
    "link": "https://rogstream.fun//hls/tata/play.php?id=239",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-2694-j638bugw-v1/imageContent-2694-j638bugw-m1.png",
    "genre": "Kids"
  },
  {
    "name": "Discovery Kids",
    "link": "https://rogstream.fun//hls/tata/play.php?id=119",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-24724-jgvwokqw-v1/imageContent-24724-jgvwokqw-m1.png",
    "genre": "Kids"
  },
  {
    "name": "Sonic Nickelodeon",
    "link": "https://rogstream.fun//hls/tata/play.php?id=127",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-8018-j7a7n1a8-v1/imageContent-8018-j7a7n1a8-m1.png",
    "genre": "Kids"
  },
  {
    "name": "SONY YAY!",
    "link": "https://rogstream.fun//hls/tata/play.php?id=45",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-132-j5fpedbs-v4/imageContent-132-j5fpedbs-m6.png",
    "genre": "Kids"
  },
  {
    "name": "Tata Play Toons+",
    "link": "https://rogstream.fun//hls/tata/play.php?id=994",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-71550-l4zi50uw-v2/imageContent-71550-l4zi50uw-m5.png",
    "genre": "null"
  },
  {
    "name": "Gubbare",
    "link": "https://rogstream.fun//hls/tata/play.php?id=867",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-57985-khvdchrs-v2/imageContent-57985-khvdchrs-m2.png",
    "genre": "Kannada"
  },
  {
    "name": "ETV Bal Bharat",
    "link": "https://rogstream.fun//hls/tata/play.php?id=905",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-62236-kny7yzaw-v1/imageContent-62236-kny7yzaw-m1.png",
    "genre": "Kids"
  },
  {
    "name": "Nick Jr",
    "link": "https://rogstream.fun//hls/tata/play.php?id=118",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-390-j5jpps1s-v1/imageContent-390-j5jpps1s-m1.png",
    "genre": "Kids"
  },
  {
    "name": "Disney Junior",
    "link": "https://rogstream.fun//hls/tata/play.php?id=144",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-35368-jr3xz7xs-v1/imageContent-35368-jr3xz7xs-m1.png",
    "genre": "Kids"
  },
  {
    "name": "CBeeBies",
    "link": "https://rogstream.fun//hls/tata/play.php?id=816",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-52868-k9af4lxs-v1/imageContent-52868-k9af4lxs-m1.png",
    "genre": "Kids"
  },
  {
    "name": "Tata Play Smart Manager",
    "link": "https://rogstream.fun//hls/tata/play.php?id=111",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-352-j5jmrkqw-v2/imageContent-352-j5jmrkqw-m2.png",
    "genre": "Knowledge"
  },
  {
    "name": "Tata Play Vedic Maths",
    "link": "https://rogstream.fun//hls/tata/play.php?id=167",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-703-j5l1wgco-v2/imageContent-703-j5l1wgco-m2.png",
    "genre": "Knowledge"
  },
  {
    "name": "National Geographic HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=605",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12234-j9qoqc54-v1/imageContent-12234-j9qoqc54-m1.png",
    "genre": "Knowledge"
  },
  {
    "name": "National Geographic",
    "link": "https://rogstream.fun//hls/tata/play.php?id=137",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-462-j5jw44hk-v2/imageContent-462-j5jw44hk-m3.png",
    "genre": "Knowledge"
  },
  {
    "name": "Nat Geo Wild HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=413",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11631-j9k0ndjc-v1/imageContent-11631-j9k0ndjc-m1.png",
    "genre": "Knowledge"
  },
  {
    "name": "Nat Geo Wild",
    "link": "https://rogstream.fun//hls/tata/play.php?id=184",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-438-j5jt26xc-v2/imageContent-438-j5jt26xc-m2.png",
    "genre": "Knowledge"
  },
  {
    "name": "Discovery HD World",
    "link": "https://rogstream.fun//hls/tata/play.php?id=341",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11396-j9jq7jg0-v1/imageContent-11396-j9jq7jg0-m1.png",
    "genre": "Knowledge"
  },
  {
    "name": "Discovery Channel",
    "link": "https://rogstream.fun//hls/tata/play.php?id=219",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-876-j5mcolzc-v2/imageContent-876-j5mcolzc-m2.png",
    "genre": "Knowledge"
  },
  {
    "name": "Animal Planet HD World",
    "link": "https://rogstream.fun//hls/tata/play.php?id=287",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-44915-jzuyh3w0-v1/imageContent-44915-jzuyh3w0-m1.png",
    "genre": "Knowledge"
  },
  {
    "name": "Discovery Science",
    "link": "https://rogstream.fun//hls/tata/play.php?id=113",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-369-j5jo9i3s-v1/imageContent-369-j5jo9i3s-m1.png",
    "genre": "Knowledge"
  },
  {
    "name": "Animal Planet",
    "link": "https://rogstream.fun//hls/tata/play.php?id=130",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-432-j5jsx754-v4/imageContent-432-j5jsx754-m4.png",
    "genre": "Knowledge"
  },
  {
    "name": "History TV18 HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=616",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12267-j9qx2si0-v3/imageContent-12267-j9qx2si0-m3.png",
    "genre": "Knowledge"
  },
  {
    "name": "History TV18",
    "link": "https://rogstream.fun//hls/tata/play.php?id=172",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-745-j5l54bnc-v2/imageContent-745-j5l54bnc-m2.png",
    "genre": "Knowledge"
  },
  {
    "name": "SONY BBC Earth HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=460",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11779-j9kx7rk0-v3/imageContent-11779-j9kx7rk0-m4.png",
    "genre": "Knowledge"
  },
  {
    "name": "SONY BBC Earth",
    "link": "https://rogstream.fun//hls/tata/play.php?id=158",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-551-j5ky7ufk-v5/imageContent-551-j5ky7ufk-m7.png",
    "genre": "Knowledge"
  },
  {
    "name": "EPIC",
    "link": "https://rogstream.fun//hls/tata/play.php?id=126",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-8017-j7a4nf20-v1/imageContent-8017-j7a4nf20-m1.png",
    "genre": "Knowledge"
  },
  {
    "name": "FOX Life HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=367",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11429-j9jqu3y0-v1/imageContent-11429-j9jqu3y0-m1.png",
    "genre": "Lifestyle"
  },
  {
    "name": "DD Gyan Darshan",
    "link": "https://rogstream.fun//hls/tata/play.php?id=646",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-23855-jg9jt83c-v2/imageContent-23855-jg9jt83c-m8.png",
    "genre": "Knowledge"
  },
  {
    "name": "TLC HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=480",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11858-j9l1v6k8-v3/imageContent-11858-j9l1v6k8-m4.png",
    "genre": "Lifestyle"
  },
  {
    "name": "TLC",
    "link": "https://rogstream.fun//hls/tata/play.php?id=135",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-453-j5jtup5s-v2/imageContent-453-j5jtup5s-m2.png",
    "genre": "Lifestyle"
  },
  {
    "name": "GOOD TiMES",
    "link": "https://rogstream.fun//hls/tata/play.php?id=136",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-42020-jxb7qjw8-v1/imageContent-42020-jxb7qjw8-m1.png",
    "genre": "Lifestyle"
  },
  {
    "name": "Discovery Turbo",
    "link": "https://rogstream.fun//hls/tata/play.php?id=228",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-967-j5mr8tw8-v1/imageContent-967-j5mr8tw8-m1.png",
    "genre": "Lifestyle"
  },
  {
    "name": "Travelxp HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=484",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11862-j9l2p61c-v1/imageContent-11862-j9l2p61c-m1.png",
    "genre": "Lifestyle"
  },
  {
    "name": "Travelxp",
    "link": "https://rogstream.fun//hls/tata/play.php?id=55",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-157-j5frd4e8-v1/imageContent-157-j5frd4e8-m1.png",
    "genre": "Lifestyle"
  },
  {
    "name": "Food Food",
    "link": "https://rogstream.fun//hls/tata/play.php?id=117",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-384-j5jpj9hc-v1/imageContent-384-j5jpj9hc-m1.png",
    "genre": "Lifestyle"
  },
  {
    "name": "Fashion TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=227",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-959-j5mipsnk-v1/imageContent-959-j5mipsnk-m1.png",
    "genre": "Lifestyle"
  },
  {
    "name": "Tata Play Ibaadat",
    "link": "https://rogstream.fun//hls/tata/play.php?id=735",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-44321-jz2zb8pk-v3/imageContent-44321-jz2zb8pk-m3.png",
    "genre": "Spiritual"
  },
  {
    "name": "MTV HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=406",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11610-j9jzore8-v2/imageContent-11610-j9jzore8-m2.png",
    "genre": "Music"
  },
  {
    "name": "MTV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=103",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-330-j5jl5on4-v2/imageContent-330-j5jl5on4-m2.png",
    "genre": "Music"
  },
  {
    "name": "9XM",
    "link": "https://rogstream.fun//hls/tata/play.php?id=139",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-468-j5jwdbi8-v1/imageContent-468-j5jwdbi8-m1.png",
    "genre": "Music"
  },
  {
    "name": "Zoom",
    "link": "https://rogstream.fun//hls/tata/play.php?id=96",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-24696-jgrq09a8-v1/imageContent-24696-jgrq09a8-m1.png",
    "genre": "Music"
  },
  {
    "name": "E24",
    "link": "https://rogstream.fun//hls/tata/play.php?id=224",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-946-j5mip6a0-v1/imageContent-946-j5mip6a0-m1.png",
    "genre": "Music"
  },
  {
    "name": "MTV Beats HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=405",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11607-j9jznhvc-v1/imageContent-11607-j9jznhvc-m1.png",
    "genre": "Music"
  },
  {
    "name": "MTV Beats",
    "link": "https://rogstream.fun//hls/tata/play.php?id=236",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-1018-j5nhzme0-v1/imageContent-1018-j5nhzme0-m1.png",
    "genre": "Music"
  },
  {
    "name": "B4U Music",
    "link": "https://rogstream.fun//hls/tata/play.php?id=9",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-24-j5fb33y0-v1/imageContent-24-j5fb33y0-m1.png",
    "genre": "Music"
  },
  {
    "name": "Mastiii",
    "link": "https://rogstream.fun//hls/tata/play.php?id=17",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-57-j5fdr2f4-v1/imageContent-57-j5fdr2f4-m1.png",
    "genre": "Music"
  },
  {
    "name": "Zing",
    "link": "https://rogstream.fun//hls/tata/play.php?id=517",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11993-j9luua9c-v1/imageContent-11993-j9luua9c-m1.png",
    "genre": "Music"
  },
  {
    "name": "9X Jalwa",
    "link": "https://rogstream.fun//hls/tata/play.php?id=39",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-102-j5fl5pyg-v1/imageContent-102-j5fl5pyg-m1.png",
    "genre": "Music"
  },
  {
    "name": "Showbox",
    "link": "https://rogstream.fun//hls/tata/play.php?id=733",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-44228-jyzbtse0-v1/imageContent-44228-jyzbtse0-m1.png",
    "genre": "Music"
  },
  {
    "name": "VH1 HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=539",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12030-j9nyctjk-v3/imageContent-12030-j9nyctjk-m5.png",
    "genre": "Music"
  },
  {
    "name": "Tata Play Aradhana",
    "link": "https://rogstream.fun//hls/tata/play.php?id=742",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-47669-k2loqvjc-v5/imageContent-47669-k2loqvjc-m5.png",
    "genre": "Spiritual"
  },
  {
    "name": "Aastha",
    "link": "https://rogstream.fun//hls/tata/play.php?id=38",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-105-j5fl5t1k-v1/imageContent-105-j5fl5t1k-m1.png",
    "genre": "Spiritual"
  },
  {
    "name": "Sanskar",
    "link": "https://rogstream.fun//hls/tata/play.php?id=43",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-123-j5fnf1fc-v1/imageContent-123-j5fnf1fc-m1.png",
    "genre": "Spiritual"
  },
  {
    "name": "Paras Gold One",
    "link": "https://rogstream.fun//hls/tata/play.php?id=446",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11725-j9kw030w-v4/imageContent-11725-j9kw030w-m4.png",
    "genre": "Spiritual"
  },
  {
    "name": "MH One Shraddha",
    "link": "https://rogstream.fun//hls/tata/play.php?id=397",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11590-j9jyx4so-v1/imageContent-11590-j9jyx4so-m1.png",
    "genre": "Spiritual"
  },
  {
    "name": "Sadhna TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=447",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11746-j9kwfmfs-v2/imageContent-11746-j9kwfmfs-m2.png",
    "genre": "Spiritual"
  },
  {
    "name": "Sharnam TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=385",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11547-j9ju62m8-v3/imageContent-11547-j9ju62m8-m3.png",
    "genre": "Spiritual"
  },
  {
    "name": "Peace of Mind",
    "link": "https://rogstream.fun//hls/tata/play.php?id=420",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11664-j9kcouh4-v1/imageContent-11664-j9kcouh4-m1.png",
    "genre": "Spiritual"
  },
  {
    "name": "Jinvani Channel",
    "link": "https://rogstream.fun//hls/tata/play.php?id=373",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11507-j9jto5k0-v1/imageContent-11507-j9jto5k0-m1.png",
    "genre": "Spiritual"
  },
  {
    "name": "Dharma Sandesh",
    "link": "https://rogstream.fun//hls/tata/play.php?id=288",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/ARIH_Thumbnail-v2/ARIH_Thumbnail.png",
    "genre": "Spiritual"
  },
  {
    "name": "Ishwar TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=594",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12196-j9q3ez54-v1/imageContent-12196-j9q3ez54-m1.png",
    "genre": "Spiritual"
  },
  {
    "name": "Satsang TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=456",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11767-j9kwretk-v1/imageContent-11767-j9kwretk-m1.png",
    "genre": "Spiritual"
  },
  {
    "name": "Shubh TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=458",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11775-j9kwuk8w-v1/imageContent-11775-j9kwuk8w-m1.png",
    "genre": "Spiritual"
  },
  {
    "name": "Aastha Bhajan",
    "link": "https://rogstream.fun//hls/tata/play.php?id=283",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11212-j9icd8jc-v1/imageContent-11212-j9icd8jc-m1.png",
    "genre": "Spiritual"
  },
  {
    "name": "Vedic",
    "link": "https://rogstream.fun//hls/tata/play.php?id=500",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11903-j9l4wxy8-v1/imageContent-11903-j9l4wxy8-m1.png",
    "genre": "Spiritual"
  },
  {
    "name": "Subharti TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=728",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-43103-jxzlrb0o-v2/imageContent-43103-jxzlrb0o-m2.png",
    "genre": "Hindi"
  },
  {
    "name": "DIVYA",
    "link": "https://rogstream.fun//hls/tata/play.php?id=775",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11408-j9jqmlvs-v6/imageContent-11408-j9jqmlvs-m4.png",
    "genre": "Spiritual"
  },
  {
    "name": "Awakening",
    "link": "https://rogstream.fun//hls/tata/play.php?id=935",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-65100-kspm2814-v1/imageContent-65100-kspm2814-m2.png",
    "genre": "Spiritual"
  },
  {
    "name": "Santwani",
    "link": "https://rogstream.fun//hls/tata/play.php?id=911",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-62610-koxujv2g-v2/imageContent-62610-koxujv2g-m3.png",
    "genre": "Spiritual"
  },
  {
    "name": "Aadinath TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=969",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-69024-kzo68h3c-v1/imageContent-69024-kzo68h3c-m1.png",
    "genre": "Spiritual"
  },
  {
    "name": "Zee Ganga",
    "link": "https://rogstream.fun//hls/tata/play.php?id=12",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-42-j5fc9b8g-v3/imageContent-42-j5fc9b8g-m3.png",
    "genre": "Hindi"
  },
  {
    "name": "MH One Dil Se",
    "link": "https://rogstream.fun//hls/tata/play.php?id=970",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-69027-kzo69re0-v1/imageContent-69027-kzo69re0-m2.png",
    "genre": "Hindi"
  },
  {
    "name": "Zee Biskope",
    "link": "https://rogstream.fun//hls/tata/play.php?id=814",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-52853-k982u9ts-v1/imageContent-52853-k982u9ts-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "Filamchi Bhojpuri",
    "link": "https://rogstream.fun//hls/tata/play.php?id=830",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-55553-kdzetu0o-v3/imageContent-55553-kdzetu0o-m8.png",
    "genre": "Hindi"
  },
  {
    "name": "Bhojpuri Cinema",
    "link": "https://rogstream.fun//hls/tata/play.php?id=181",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-768-j5m31cfs-v1/imageContent-768-j5m31cfs-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "Oscar Movies Bhojpuri",
    "link": "https://rogstream.fun//hls/tata/play.php?id=431",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11653-j9k1kyn4-v1/imageContent-11653-j9k1kyn4-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "B4U Bhojpuri",
    "link": "https://rogstream.fun//hls/tata/play.php?id=729",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-43109-jxzlyrjc-v1/imageContent-43109-jxzlyrjc-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "Pasand",
    "link": "https://rogstream.fun//hls/tata/play.php?id=899",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-62150-knicubns-v2/imageContent-62150-knicubns-m2.png",
    "genre": "Movies"
  },
  {
    "name": "News18 Uttar Pradesh Uttarakhand",
    "link": "https://rogstream.fun//hls/tata/play.php?id=79",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-23223-jeyhnzfs-v1/imageContent-23223-jeyhnzfs-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "Kashish News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=838",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-56320-kf6aheeo-v2/imageContent-56320-kf6aheeo-m3.png",
    "genre": "Hindi"
  },
  {
    "name": "Zee Uttar Pradesh Uttarakhand",
    "link": "https://rogstream.fun//hls/tata/play.php?id=637",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-21315-je85i72w-v1/imageContent-21315-je85i72w-m5.png",
    "genre": "News"
  },
  {
    "name": "News18 Bihar Jharkhand",
    "link": "https://rogstream.fun//hls/tata/play.php?id=166",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-23206-jey39grs-v1/imageContent-23206-jey39grs-m2.png",
    "genre": "Hindi"
  },
  {
    "name": "Zee Bihar Jharkhand",
    "link": "https://rogstream.fun//hls/tata/play.php?id=20",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11140-j96syfuo-v1/imageContent-11140-j96syfuo-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "News State Bihar Jharkhand",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1034",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-74233-la71l5eg-v4/imageContent-74233-la71l5eg-m5.png",
    "genre": "null"
  },
  {
    "name": "India News UP UK",
    "link": "https://rogstream.fun//hls/tata/play.php?id=30",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-9007-j7ss7x80-v1/imageContent-9007-j7ss7x80-m1.png",
    "genre": "News"
  },
  {
    "name": "News State UP Uttarakhand",
    "link": "https://rogstream.fun//hls/tata/play.php?id=613",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12252-j9qrkgu8-v3/imageContent-12252-j9qrkgu8-m2.png",
    "genre": "News"
  },
  {
    "name": "Sadhna Plus News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=827",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-55124-kd89b0f4-v1/imageContent-55124-kd89b0f4-m4.png",
    "genre": "Hindi"
  },
  {
    "name": "Sahara Samay UP-UKD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=973",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-69315-l02nqy9k-v1/imageContent-69315-l02nqy9k-m6.png",
    "genre": "Hindi"
  },
  {
    "name": "DD Bharati",
    "link": "https://rogstream.fun//hls/tata/play.php?id=316",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11320-j9jlusp4-v1/imageContent-11320-j9jlusp4-m1.png",
    "genre": "Entertainment"
  },
  {
    "name": "DD Uttar Pradesh",
    "link": "https://rogstream.fun//hls/tata/play.php?id=339",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11389-j9jq2ybk-v1/imageContent-11389-j9jq2ybk-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "DD Bihar",
    "link": "https://rogstream.fun//hls/tata/play.php?id=317",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11321-j9jlxaz4-v1/imageContent-11321-j9jlxaz4-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "News18 Madhya Pradesh Chhattisgarh",
    "link": "https://rogstream.fun//hls/tata/play.php?id=203",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-23266-jez523so-v1/imageContent-23266-jez523so-m2.png",
    "genre": "Hindi"
  },
  {
    "name": "Zee Madhya Pradesh Chattisgarh",
    "link": "https://rogstream.fun//hls/tata/play.php?id=512",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11981-j9lunvjs-v1/imageContent-11981-j9lunvjs-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "IBC 24",
    "link": "https://rogstream.fun//hls/tata/play.php?id=27",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-78-j5fikkk8-v1/imageContent-78-j5fikkk8-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "Bansal News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=652",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-25494-ji1a6ym0-v2/imageContent-25494-ji1a6ym0-m2.png",
    "genre": "Hindi"
  },
  {
    "name": "INH 24X7",
    "link": "https://rogstream.fun//hls/tata/play.php?id=698",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-35622-jrr5mvig-v1/imageContent-35622-jrr5mvig-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "India News MP CG",
    "link": "https://rogstream.fun//hls/tata/play.php?id=222",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-8971-j7rffd40-v1/imageContent-8971-j7rffd40-m1.png",
    "genre": "News"
  },
  {
    "name": "News State MP CG",
    "link": "https://rogstream.fun//hls/tata/play.php?id=643",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-31191-jldnwhhs-v1/imageContent-31191-jldnwhhs-m1.png",
    "genre": "News"
  },
  {
    "name": "Swaraj Express SMBC",
    "link": "https://rogstream.fun//hls/tata/play.php?id=577",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12143-j9p7vx9c-v1/imageContent-12143-j9p7vx9c-m1.png",
    "genre": "News"
  },
  {
    "name": "Sadhna News MP CG",
    "link": "https://rogstream.fun//hls/tata/play.php?id=829",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-55112-kd88u9so-v1/imageContent-55112-kd88u9so-m2.png",
    "genre": "Hindi"
  },
  {
    "name": "News 24 Madhyapradesh Chattisgarh",
    "link": "https://rogstream.fun//hls/tata/play.php?id=902",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-62181-knrp5974-v1/imageContent-62181-knrp5974-m1.png",
    "genre": "News"
  },
  {
    "name": "Anaadi TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=937",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-65108-ksr7d028-v1/imageContent-65108-ksr7d028-m2.png",
    "genre": "Spiritual"
  },
  {
    "name": "News Hour",
    "link": "https://rogstream.fun//hls/tata/play.php?id=993",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-71503-l4su2elk-v2/imageContent-71503-l4su2elk-m7.png",
    "genre": "null"
  },
  {
    "name": "TV27 News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1009",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-72340-l6nkcw74-v1/imageContent-72340-l6nkcw74-m1.png",
    "genre": "null"
  },
  {
    "name": "DD Madhya Pradesh",
    "link": "https://rogstream.fun//hls/tata/play.php?id=330",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11351-j9jpmdvc-v1/imageContent-11351-j9jpmdvc-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "News18 Rajasthan",
    "link": "https://rogstream.fun//hls/tata/play.php?id=205",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-23226-jeyiac80-v1/imageContent-23226-jeyiac80-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "Zee Rajasthan News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=583",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12158-j9pd883k-v1/imageContent-12158-j9pd883k-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "First India Rajasthan",
    "link": "https://rogstream.fun//hls/tata/play.php?id=650",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-25048-jhir0kko-v1/imageContent-25048-jhir0kko-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "India News Rajasthan",
    "link": "https://rogstream.fun//hls/tata/play.php?id=14",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-9012-j7ssufeo-v1/imageContent-9012-j7ssufeo-m1.png",
    "genre": "News"
  },
  {
    "name": "Patrika TV Rajasthan",
    "link": "https://rogstream.fun//hls/tata/play.php?id=5",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-10-j5f7xqxc-v1/imageContent-10-j5f7xqxc-m1.png",
    "genre": "News"
  },
  {
    "name": "Jan TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=870",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-58492-kixkb2ug-v1/imageContent-58492-kixkb2ug-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "Sach Bedhadak",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1075",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/SACHB_Thumbnail-v2/SACHB_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "TV 24",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1074",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/TV24N_Thumbnail-v2/TV24N_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "DD Rajasthan",
    "link": "https://rogstream.fun//hls/tata/play.php?id=332",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11375-j9jpws3k-v1/imageContent-11375-j9jpws3k-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "News18 Jammu Kashmir Ladakh Himachal",
    "link": "https://rogstream.fun//hls/tata/play.php?id=354",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11426-j9jqror4-v3/imageContent-11426-j9jqror4-m3.png",
    "genre": "Hindi"
  },
  {
    "name": "Zee Salaam",
    "link": "https://rogstream.fun//hls/tata/play.php?id=518",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11984-j9lup2rc-v2/imageContent-11984-j9lup2rc-m3.png",
    "genre": "Hindi"
  },
  {
    "name": "Channel WIN",
    "link": "https://rogstream.fun//hls/tata/play.php?id=161",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-7906-j7684flk-v1/imageContent-7906-j7684flk-m1.png",
    "genre": "Hindi"
  },
  {
    "name": "Gulistan News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=365",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11455-j9jqz61k-v2/imageContent-11455-j9jqz61k-m4.png",
    "genre": "Hindi"
  },
  {
    "name": "DD Urdu",
    "link": "https://rogstream.fun//hls/tata/play.php?id=338",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11384-j9jq29mo-v2/imageContent-11384-j9jq29mo-m3.png",
    "genre": "Hindi"
  },
  {
    "name": "Colors Marathi HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=308",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-16764-jcix3xpc-v2/imageContent-16764-jcix3xpc-m2.PNG",
    "genre": "Marathi"
  },
  {
    "name": "Colors Marathi",
    "link": "https://rogstream.fun//hls/tata/play.php?id=134",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-450-j5jtumug-v3/imageContent-450-j5jtumug-m3.png",
    "genre": "Marathi"
  },
  {
    "name": "Zee Marathi HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=501",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11930-j9l5hv1c-v1/imageContent-11930-j9l5hv1c-m1.png",
    "genre": "Marathi"
  },
  {
    "name": "Zee Marathi",
    "link": "https://rogstream.fun//hls/tata/play.php?id=251",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-8846-j7pq3jmo-v2/imageContent-8846-j7pq3jmo-m3.png",
    "genre": "Marathi"
  },
  {
    "name": "Star Pravah HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=469",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11807-j9kyzle8-v1/imageContent-11807-j9kyzle8-m1.png",
    "genre": "Marathi"
  },
  {
    "name": "Zee YUVA",
    "link": "https://rogstream.fun//hls/tata/play.php?id=248",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11136-j96p4r28-v1/imageContent-11136-j96p4r28-m1.png",
    "genre": "Marathi"
  },
  {
    "name": "Sony Marathi",
    "link": "https://rogstream.fun//hls/tata/play.php?id=658",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-30959-jky1nic8-v2/imageContent-30959-jky1nic8-m3.png",
    "genre": "Marathi"
  },
  {
    "name": "Tata Play Toons+",
    "link": "https://rogstream.fun//hls/tata/play.php?id=999",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/TONSMAR_Thumbnail-v10/TONSMAR_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "Zee Talkies HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=515",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11987-j9lura8g-v1/imageContent-11987-j9lura8g-m1.png",
    "genre": "Marathi"
  },
  {
    "name": "Zee Talkies",
    "link": "https://rogstream.fun//hls/tata/play.php?id=249",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-8842-j7pp0me8-v2/imageContent-8842-j7pp0me8-m2.png",
    "genre": "Marathi"
  },
  {
    "name": "Fakt Marathi",
    "link": "https://rogstream.fun//hls/tata/play.php?id=192",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-801-j5m7jg34-v2/imageContent-801-j5m7jg34-m2.png",
    "genre": "Marathi"
  },
  {
    "name": "Shemaroo MarathiBana",
    "link": "https://rogstream.fun//hls/tata/play.php?id=800",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-49025-k5lxf84o-v2/imageContent-49025-k5lxf84o-m2.png",
    "genre": "Marathi"
  },
  {
    "name": "ABP Majha",
    "link": "https://rogstream.fun//hls/tata/play.php?id=155",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-542-j5kwwml4-v2/imageContent-542-j5kwwml4-m2.png",
    "genre": "Marathi"
  },
  {
    "name": "News18 Lokmat",
    "link": "https://rogstream.fun//hls/tata/play.php?id=140",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12140-j9p7vga8-v1/imageContent-12140-j9p7vga8-m1.png",
    "genre": "Marathi"
  },
  {
    "name": "Zee 24 Taas",
    "link": "https://rogstream.fun//hls/tata/play.php?id=261",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11156-j96xbrfk-v1/imageContent-11156-j96xbrfk-m1.png",
    "genre": "Marathi"
  },
  {
    "name": "Saam TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=546",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12056-j9o5hia8-v1/imageContent-12056-j9o5hia8-m1.png",
    "genre": "Marathi"
  },
  {
    "name": "Jai Maharashtra",
    "link": "https://rogstream.fun//hls/tata/play.php?id=233",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-1014-j5ngcjug-v2/imageContent-1014-j5ngcjug-m2.png",
    "genre": "Marathi"
  },
  {
    "name": "TV9 Marathi",
    "link": "https://rogstream.fun//hls/tata/play.php?id=97",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-315-j5jjsht4-v2/imageContent-315-j5jjsht4-m2.png",
    "genre": "Marathi"
  },
  {
    "name": "Lokshahi",
    "link": "https://rogstream.fun//hls/tata/play.php?id=826",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-54569-kcviihig-v3/imageContent-54569-kcviihig-m3.png",
    "genre": "Marathi"
  },
  {
    "name": "News State Maharashtra Goa",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1038",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/NSMG_Thumbnail-v2/NSMG_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "DD Sahyadri",
    "link": "https://rogstream.fun//hls/tata/play.php?id=336",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11378-j9jq0f9s-v1/imageContent-11378-j9jq0f9s-m1.png",
    "genre": "Marathi"
  },
  {
    "name": "Sangeet Marathi",
    "link": "https://rogstream.fun//hls/tata/play.php?id=217",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-870-j5mb6j80-v1/imageContent-870-j5mb6j80-m1.png",
    "genre": "Marathi"
  },
  {
    "name": "Colors Bangla HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=305",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11287-j9j4tl2w-v2/imageContent-11287-j9j4tl2w-m3.png",
    "genre": "Bengali"
  },
  {
    "name": "Colors Bangla",
    "link": "https://rogstream.fun//hls/tata/play.php?id=26",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-75-j5fikj0o-v1/imageContent-75-j5fikj0o-m1.png",
    "genre": "Bengali"
  },
  {
    "name": "Zee Bangla HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=522",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11970-j9luk5ag-v1/imageContent-11970-j9luk5ag-m1.png",
    "genre": "Bengali"
  },
  {
    "name": "Star Jalsha HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=468",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11813-j9kz0vow-v1/imageContent-11813-j9kz0vow-m1.png",
    "genre": "Bengali"
  },
  {
    "name": "SONY AATH",
    "link": "https://rogstream.fun//hls/tata/play.php?id=34",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-94-j5fjveio-v3/imageContent-94-j5fjveio-m2.png",
    "genre": "Bengali"
  },
  {
    "name": "Tata Play Toons+",
    "link": "https://rogstream.fun//hls/tata/play.php?id=995",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/TONSBEN_Thumbnail-v5/TONSBEN_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "Aakaash Aath",
    "link": "https://rogstream.fun//hls/tata/play.php?id=129",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-267-j5j8x17s-v1/imageContent-267-j5j8x17s-m2.png",
    "genre": "Bengali"
  },
  {
    "name": "Ruposhi Bangla",
    "link": "https://rogstream.fun//hls/tata/play.php?id=3",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-9-j5f6oezc-v1/imageContent-9-j5f6oezc-m1.png",
    "genre": "Bengali"
  },
  {
    "name": "Enterr10 Bangla",
    "link": "https://rogstream.fun//hls/tata/play.php?id=788",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-48494-k3zd1a80-v1/imageContent-48494-k3zd1a80-m4.png",
    "genre": "Bengali"
  },
  {
    "name": "Jalsha Movies HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=537",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11514-j9jtp5tk-v1/imageContent-11514-j9jtp5tk-m2.png",
    "genre": "Bengali"
  },
  {
    "name": "Zee Bangla Cinema",
    "link": "https://rogstream.fun//hls/tata/play.php?id=254",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11009-j8xbk7js-v1/imageContent-11009-j8xbk7js-m1.png",
    "genre": "Bengali"
  },
  {
    "name": "Colors Bangla Cinema",
    "link": "https://rogstream.fun//hls/tata/play.php?id=896",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/CLRBANG_Thumbnail-v2/CLRBANG_Thumbnail.png",
    "genre": "Bengali"
  },
  {
    "name": "Khushboo Bangla",
    "link": "https://rogstream.fun//hls/tata/play.php?id=379",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-27017-jjo92wyg-v1/imageContent-27017-jjo92wyg-m1.png",
    "genre": "Bengali"
  },
  {
    "name": "ABP Ananda",
    "link": "https://rogstream.fun//hls/tata/play.php?id=102",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-321-j5jl5g5k-v2/imageContent-321-j5jl5g5k-m2.png",
    "genre": "Bengali"
  },
  {
    "name": "Zee 24 Ghanta",
    "link": "https://rogstream.fun//hls/tata/play.php?id=258",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-9031-j7u3nz6w-v1/imageContent-9031-j7u3nz6w-m1.png",
    "genre": "Bengali"
  },
  {
    "name": "Kolkata TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=381",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11533-j9jtxqgo-v1/imageContent-11533-j9jtxqgo-m1.png",
    "genre": "Bengali"
  },
  {
    "name": "News18 Bangla",
    "link": "https://rogstream.fun//hls/tata/play.php?id=23",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-22186-jenuxy6g-v2/imageContent-22186-jenuxy6g-m5.png",
    "genre": "Bengali"
  },
  {
    "name": "TV9 Bangla",
    "link": "https://rogstream.fun//hls/tata/play.php?id=873",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-59401-kjudpa1c-v2/imageContent-59401-kjudpa1c-m3.png",
    "genre": "Bengali"
  },
  {
    "name": "Republic Bangla",
    "link": "https://rogstream.fun//hls/tata/play.php?id=890",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-61251-klylf4u0-v2/imageContent-61251-klylf4u0-m9.png",
    "genre": "News"
  },
  {
    "name": "News Time Bangla",
    "link": "https://rogstream.fun//hls/tata/play.php?id=207",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-852-j5m9xe80-v1/imageContent-852-j5m9xe80-m1.png",
    "genre": "Bengali"
  },
  {
    "name": "CTVN AKD Plus",
    "link": "https://rogstream.fun//hls/tata/play.php?id=311",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11299-j9j61p1k-v1/imageContent-11299-j9j61p1k-m1.png",
    "genre": "Bengali"
  },
  {
    "name": "Calcutta News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=648",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-24685-jgrl9bqg-v1/imageContent-24685-jgrl9bqg-m2.png",
    "genre": "Bengali"
  },
  {
    "name": "DD Bangla",
    "link": "https://rogstream.fun//hls/tata/play.php?id=314",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11314-j9j8kcxk-v1/imageContent-11314-j9j8kcxk-m1.png",
    "genre": "Bengali"
  },
  {
    "name": "Sangeet Bangla",
    "link": "https://rogstream.fun//hls/tata/play.php?id=215",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-864-j5mava4o-v1/imageContent-864-j5mava4o-m1.png",
    "genre": "Bengali"
  },
  {
    "name": "Zee Telugu HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=635",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-21304-je859a3c-v1/imageContent-21304-je859a3c-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "Zee Telugu",
    "link": "https://rogstream.fun//hls/tata/play.php?id=250",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-8845-j7ppmm28-v3/imageContent-8845-j7ppmm28-m7.png",
    "genre": "Telugu"
  },
  {
    "name": "Gemini TV HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=355",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11444-j9jqx6a8-v1/imageContent-11444-j9jqx6a8-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "STAR Maa HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=516",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-25361-jhsvr3nc-v1/imageContent-25361-jhsvr3nc-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "ETV HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=359",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11418-j9jqpvxs-v1/imageContent-11418-j9jqpvxs-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "ETV Telugu",
    "link": "https://rogstream.fun//hls/tata/play.php?id=145",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-492-j5jxvaeo-v1/imageContent-492-j5jxvaeo-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "Gemini Life",
    "link": "https://rogstream.fun//hls/tata/play.php?id=361",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11437-j9jqv7ao-v2/imageContent-11437-j9jqv7ao-m2.png",
    "genre": "Telugu"
  },
  {
    "name": "ETV Plus",
    "link": "https://rogstream.fun//hls/tata/play.php?id=352",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11423-j9jqqegg-v1/imageContent-11423-j9jqqegg-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "Tata Play Toons+",
    "link": "https://rogstream.fun//hls/tata/play.php?id=996",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/TONSTEL_Thumbnail-v4/TONSTEL_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "Telugu Naaptol",
    "link": "https://rogstream.fun//hls/tata/play.php?id=590",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12186-j9plqam8-v4/imageContent-12186-j9plqam8-m5.png",
    "genre": "Telugu"
  },
  {
    "name": "Vissa TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=585",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12164-j9phk5iw-v1/imageContent-12164-j9phk5iw-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "Tata Play English in Telugu",
    "link": "https://rogstream.fun//hls/tata/play.php?id=645",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-23721-jg0fhacg-v2/imageContent-23721-jg0fhacg-m2.PNG",
    "genre": "Knowledge"
  },
  {
    "name": "Gemini Movies HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=362",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11438-j9jqvyaw-v1/imageContent-11438-j9jqvyaw-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "Star Maa Movies HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=387",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-18348-jcqpx9ns-v1/imageContent-18348-jcqpx9ns-m1.PNG",
    "genre": "Telugu"
  },
  {
    "name": "Star Maa Gold",
    "link": "https://rogstream.fun//hls/tata/play.php?id=388",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11553-j9julvag-v3/imageContent-11553-j9julvag-m6.png",
    "genre": "Telugu"
  },
  {
    "name": "Zee Cinemalu HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=636",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-21307-je859sm0-v1/imageContent-21307-je859sm0-m2.png",
    "genre": "Telugu"
  },
  {
    "name": "Zee Cinemalu",
    "link": "https://rogstream.fun//hls/tata/play.php?id=252",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-8849-j7pqmypc-v2/imageContent-8849-j7pqmypc-m3.png",
    "genre": "Telugu"
  },
  {
    "name": "ETV Cinema",
    "link": "https://rogstream.fun//hls/tata/play.php?id=128",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-426-j5jsm1wo-v1/imageContent-426-j5jsm1wo-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "Tata Play Telugu Classics",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1073",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/TPTECL_Thumbnail-v1/TPTECL_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "ETV Life",
    "link": "https://rogstream.fun//hls/tata/play.php?id=268",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11182-j9i3o5pc-v2/imageContent-11182-j9i3o5pc-m2.png",
    "genre": "Telugu"
  },
  {
    "name": "ETV Abhiruchi",
    "link": "https://rogstream.fun//hls/tata/play.php?id=358",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11417-j9jqpcnc-v1/imageContent-11417-j9jqpcnc-m1.png",
    "genre": "Kannada"
  },
  {
    "name": "Sakshi TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=596",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12260-j9qv0lz4-v1/imageContent-12260-j9qv0lz4-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "ETV Andhra Pradesh",
    "link": "https://rogstream.fun//hls/tata/play.php?id=146",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-498-j5jy7i80-v1/imageContent-498-j5jy7i80-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "TV5 News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=320",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11329-j9jot000-v2/imageContent-11329-j9jot000-m2.png",
    "genre": "Telugu"
  },
  {
    "name": "TV9 Telugu",
    "link": "https://rogstream.fun//hls/tata/play.php?id=11",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-30-j5fc0v80-v2/imageContent-30-j5fc0v80-m2.png",
    "genre": "Telugu"
  },
  {
    "name": "NTV Telugu",
    "link": "https://rogstream.fun//hls/tata/play.php?id=160",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-586-j5kz3qkw-v1/imageContent-586-j5kz3qkw-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "ABN Andhra Jyothy",
    "link": "https://rogstream.fun//hls/tata/play.php?id=225",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-951-j5mipdzs-v1/imageContent-951-j5mipdzs-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "T News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=49",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-145-j5fpfoe8-v1/imageContent-145-j5fpfoe8-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "V6 Telugu",
    "link": "https://rogstream.fun//hls/tata/play.php?id=274",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11194-j9i4cz80-v1/imageContent-11194-j9i4cz80-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "ETV Telangana",
    "link": "https://rogstream.fun//hls/tata/play.php?id=83",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-255-j5fz95pc-v1/imageContent-255-j5fz95pc-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "HM TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=349",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11470-j9jr4uio-v1/imageContent-11470-j9jr4uio-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "Raj News Telugu",
    "link": "https://rogstream.fun//hls/tata/play.php?id=430",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11697-j9kjhxs0-v3/imageContent-11697-j9kjhxs0-m3.png",
    "genre": "Telugu"
  },
  {
    "name": "10 TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=266",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-17564-jclrk6pk-v1/imageContent-17564-jclrk6pk-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "Swatantra TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1027",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-73187-l8ou8wo8-v1/imageContent-73187-l8ou8wo8-m4.png",
    "genre": "null"
  },
  {
    "name": "4tv News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=774",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-47248-k1vn9ouw-v2/imageContent-47248-k1vn9ouw-m4.png",
    "genre": "Telugu"
  },
  {
    "name": "DD Saptagiri",
    "link": "https://rogstream.fun//hls/tata/play.php?id=337",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11381-j9jq1pkg-v1/imageContent-11381-j9jq1pkg-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "DD Yadagiri",
    "link": "https://rogstream.fun//hls/tata/play.php?id=548",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11392-j9jq3lgw-v3/imageContent-11392-j9jq3lgw-m2.png",
    "genre": "Telugu"
  },
  {
    "name": "Raj Musix Telugu",
    "link": "https://rogstream.fun//hls/tata/play.php?id=429",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11691-j9kj50i0-v1/imageContent-11691-j9kj50i0-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "Bhakti TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=290",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11242-j9irisyw-v2/imageContent-11242-j9irisyw-m2.png",
    "genre": "Telugu"
  },
  {
    "name": "Shubhavaarta TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=445",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11745-j9kwb568-v2/imageContent-11745-j9kwb568-m3.png",
    "genre": "Telugu"
  },
  {
    "name": "Aradana TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=285",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11224-j9iclykw-v1/imageContent-11224-j9iclykw-m2.png",
    "genre": "Telugu"
  },
  {
    "name": "Swara Sagar",
    "link": "https://rogstream.fun//hls/tata/play.php?id=351",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-33495-jnqw9k9c-v4/imageContent-33495-jnqw9k9c-m5.png",
    "genre": "Telugu"
  },
  {
    "name": "Hindu Dharmam",
    "link": "https://rogstream.fun//hls/tata/play.php?id=630",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-18536-jct505dc-v2/imageContent-18536-jct505dc-m2.png",
    "genre": "Telugu"
  },
  {
    "name": "Sun TV HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=521",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11960-j9lu9vow-v1/imageContent-11960-j9lu9vow-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Sun Life",
    "link": "https://rogstream.fun//hls/tata/play.php?id=474",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11837-j9kzms9s-v2/imageContent-11837-j9kzms9s-m3.png",
    "genre": "Tamil"
  },
  {
    "name": "Star Vijay HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=496",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11909-j9l52epk-v2/imageContent-11909-j9l52epk-m4.png",
    "genre": "Tamil"
  },
  {
    "name": "Zee Tamil HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=608",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11997-j9lux9ig-v1/imageContent-11997-j9lux9ig-m5.png",
    "genre": "Tamil"
  },
  {
    "name": "Zee Tamil",
    "link": "https://rogstream.fun//hls/tata/play.php?id=257",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-9006-j7rlwsns-v2/imageContent-9006-j7rlwsns-m4.png",
    "genre": "Tamil"
  },
  {
    "name": "Jaya TV HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=708",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-38738-juc4a1ig-v5/imageContent-38738-juc4a1ig-m8.png",
    "genre": "Telugu"
  },
  {
    "name": "Colors Tamil HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=674",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-33814-jo5jba9s-v3/imageContent-33814-jo5jba9s-m3.PNG",
    "genre": "Tamil"
  },
  {
    "name": "Kalaignar TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=200",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-819-j5m8fsfs-v1/imageContent-819-j5m8fsfs-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Colors Tamil",
    "link": "https://rogstream.fun//hls/tata/play.php?id=418",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-21152-je6orzfk-v1/imageContent-21152-je6orzfk-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Raj TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=439",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11700-j9ksbotk-v1/imageContent-11700-j9ksbotk-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "D Tamil",
    "link": "https://rogstream.fun//hls/tata/play.php?id=99",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-44904-jzu8ri4o-v1/imageContent-44904-jzu8ri4o-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Raj Digital Plus",
    "link": "https://rogstream.fun//hls/tata/play.php?id=426",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11679-j9khnqag-v1/imageContent-11679-j9khnqag-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Mega TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=37",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-111-j5fl5xo8-v1/imageContent-111-j5fl5xo8-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Tamil Naaptol",
    "link": "https://rogstream.fun//hls/tata/play.php?id=723",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-41267-jwjjvrt4-v5/imageContent-41267-jwjjvrt4-m7.png",
    "genre": "Tamil"
  },
  {
    "name": "Vasanth TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=499",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11900-j9l4vygg-v1/imageContent-11900-j9l4vygg-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Polimer TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=272",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11188-j9i480zc-v1/imageContent-11188-j9i480zc-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Makkal TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=392",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11565-j9jv0np4-v1/imageContent-11565-j9jv0np4-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Sirippoli",
    "link": "https://rogstream.fun//hls/tata/play.php?id=611",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12244-j9qq7k1c-v1/imageContent-12244-j9qq7k1c-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Vendhar TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=659",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-31187-jldgeg1k-v2/imageContent-31187-jldgeg1k-m2.png",
    "genre": "Tamil"
  },
  {
    "name": "KTV HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=380",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11536-j9jty5w8-v1/imageContent-11536-j9jty5w8-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "J Movies",
    "link": "https://rogstream.fun//hls/tata/play.php?id=201",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-816-j5m8dqd4-v1/imageContent-816-j5m8dqd4-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Zee Thirai",
    "link": "https://rogstream.fun//hls/tata/play.php?id=797",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-49016-k5lx1za0-v2/imageContent-49016-k5lx1za0-m2.png",
    "genre": "Tamil"
  },
  {
    "name": "Travelxp Tamil",
    "link": "https://rogstream.fun//hls/tata/play.php?id=444",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11724-j9kvwnkg-v1/imageContent-11724-j9kvwnkg-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Puthiya Thalaimurai",
    "link": "https://rogstream.fun//hls/tata/play.php?id=220",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-882-j5mdrg4o-v1/imageContent-882-j5mdrg4o-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Thanthi TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=524",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11966-j9luich4-v1/imageContent-11966-j9luich4-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Jaya Plus",
    "link": "https://rogstream.fun//hls/tata/play.php?id=198",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-810-j5m84h14-v1/imageContent-810-j5m84h14-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "News 7 Tamil",
    "link": "https://rogstream.fun//hls/tata/play.php?id=21",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-66-j5ff8qio-v1/imageContent-66-j5ff8qio-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Polimer News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=509",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11942-j9lthi34-v1/imageContent-11942-j9lthi34-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "News18 Tamil Nadu",
    "link": "https://rogstream.fun//hls/tata/play.php?id=58",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-192-j5fsr3s0-v1/imageContent-192-j5fsr3s0-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Seithigal TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=44",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-126-j5fnf4ig-v2/imageContent-126-j5fnf4ig-m2.png",
    "genre": "Tamil"
  },
  {
    "name": "Raj News Tamil",
    "link": "https://rogstream.fun//hls/tata/play.php?id=525",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12002-j9nrtf08-v2/imageContent-12002-j9nrtf08-m2.png",
    "genre": "Tamil"
  },
  {
    "name": "Sathiyam TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=455",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11764-j9kwqj6o-v1/imageContent-11764-j9kwqj6o-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "MalaiMurasu Seithigal",
    "link": "https://rogstream.fun//hls/tata/play.php?id=391",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/MAMUR_Thumbnail-v2/MAMUR_Thumbnail.png",
    "genre": "Tamil"
  },
  {
    "name": "News J",
    "link": "https://rogstream.fun//hls/tata/play.php?id=702",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-37207-jswmhmpk-v1/imageContent-37207-jswmhmpk-m1.PNG",
    "genre": "Tamil"
  },
  {
    "name": "News Tamil 24x7",
    "link": "https://rogstream.fun//hls/tata/play.php?id=968",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-69018-kzo65wi0-v1/imageContent-69018-kzo65wi0-m2.png",
    "genre": "Tamil"
  },
  {
    "name": "DD Podhigai",
    "link": "https://rogstream.fun//hls/tata/play.php?id=334",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11369-j9jpvzjs-v1/imageContent-11369-j9jpvzjs-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Jaya Max",
    "link": "https://rogstream.fun//hls/tata/play.php?id=199",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-807-j5m7zoyo-v1/imageContent-807-j5m7zoyo-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Raj Musix",
    "link": "https://rogstream.fun//hls/tata/play.php?id=425",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11682-j9khxjoo-v1/imageContent-11682-j9khxjoo-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Isaiaruvi",
    "link": "https://rogstream.fun//hls/tata/play.php?id=647",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-24578-jgotgry0-v1/imageContent-24578-jgotgry0-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Murasu TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=411",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11616-j9k04eo0-v1/imageContent-11616-j9k04eo0-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Mega Musiq",
    "link": "https://rogstream.fun//hls/tata/play.php?id=400",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11583-j9jype8w-v3/imageContent-11583-j9jype8w-m3.png",
    "genre": "Tamil"
  },
  {
    "name": "Jothi TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=434",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11712-j9kt9fbs-v2/imageContent-11712-j9kt9fbs-m2.png",
    "genre": "Tamil"
  },
  {
    "name": "Madha TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=390",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11562-j9juzmns-v1/imageContent-11562-j9juzmns-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Angel TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=282",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11230-j9icx1i0-v1/imageContent-11230-j9icx1i0-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Nambikkai TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=408",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11622-j9k0845k-v2/imageContent-11622-j9k0845k-m2.png",
    "genre": "Tamil"
  },
  {
    "name": "SVBC 2",
    "link": "https://rogstream.fun//hls/tata/play.php?id=490",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11882-j9l3qcrs-v1/imageContent-11882-j9l3qcrs-m1.png",
    "genre": "Telugu"
  },
  {
    "name": "Udaya TV HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=492",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11883-j9l3rjzc-v1/imageContent-11883-j9l3rjzc-m1.png",
    "genre": "Kannada"
  },
  {
    "name": "Colors Kannada HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=612",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12249-j9qr87hc-v5/imageContent-12249-j9qr87hc-m5.png",
    "genre": "Kannada"
  },
  {
    "name": "Colors Kannada",
    "link": "https://rogstream.fun//hls/tata/play.php?id=108",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-345-j5jmkws0-v2/imageContent-345-j5jmkws0-m4.png",
    "genre": "Kannada"
  },
  {
    "name": "Star Suvarna HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=467",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11816-j9kz224o-v1/imageContent-11816-j9kz224o-m1.png",
    "genre": "Kannada"
  },
  {
    "name": "Zee Kannada HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=675",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-33827-jo5kcr1c-v4/imageContent-33827-jo5kcr1c-m5.png",
    "genre": "Kannada"
  },
  {
    "name": "Zee Kannada",
    "link": "https://rogstream.fun//hls/tata/play.php?id=256",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11030-j8yk2g3s-v1/imageContent-11030-j8yk2g3s-m1.png",
    "genre": "Kannada"
  },
  {
    "name": "Jeevan TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=372",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12323-jaa5pzxs-v1/imageContent-12323-jaa5pzxs-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "Colors SUPER",
    "link": "https://rogstream.fun//hls/tata/play.php?id=533",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-18308-jcqnzl68-v1/imageContent-18308-jcqnzl68-m1.PNG",
    "genre": "Kannada"
  },
  {
    "name": "SIRIKANNADA Alltime",
    "link": "https://rogstream.fun//hls/tata/play.php?id=940",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-65190-ktai1xtk-v1/imageContent-65190-ktai1xtk-m1.png",
    "genre": "Kannada"
  },
  {
    "name": "Udaya Movies",
    "link": "https://rogstream.fun//hls/tata/play.php?id=231",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-1002-j5nf74c0-v1/imageContent-1002-j5nf74c0-m1.png",
    "genre": "Kannada"
  },
  {
    "name": "Star Suvarna Plus",
    "link": "https://rogstream.fun//hls/tata/play.php?id=540",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12038-j9o0554o-v2/imageContent-12038-j9o0554o-m2.png",
    "genre": "Kannada"
  },
  {
    "name": "Colors Kannada Cinema",
    "link": "https://rogstream.fun//hls/tata/play.php?id=667",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-31454-jmirfosw-v2/imageContent-31454-jmirfosw-m2.png",
    "genre": "Kannada"
  },
  {
    "name": "Zee Picchar",
    "link": "https://rogstream.fun//hls/tata/play.php?id=810",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-52310-k7bd8ixk-v2/imageContent-52310-k7bd8ixk-m4.png",
    "genre": "Kannada"
  },
  {
    "name": "Public Movies",
    "link": "https://rogstream.fun//hls/tata/play.php?id=661",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-31188-jldk0joo-v2/imageContent-31188-jldk0joo-m2.png",
    "genre": "Kannada"
  },
  {
    "name": "Kasthuri",
    "link": "https://rogstream.fun//hls/tata/play.php?id=196",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-822-j5m8ivjs-v1/imageContent-822-j5m8ivjs-m1.png",
    "genre": "Kannada"
  },
  {
    "name": "TV9 Kannada",
    "link": "https://rogstream.fun//hls/tata/play.php?id=152",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-531-j5kuelug-v2/imageContent-531-j5kuelug-m2.png",
    "genre": "Kannada"
  },
  {
    "name": "Asianet Suvarna News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=555",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11963-j9lui4rc-v3/imageContent-11963-j9lui4rc-m4.png",
    "genre": "Kannada"
  },
  {
    "name": "News18 Kannada",
    "link": "https://rogstream.fun//hls/tata/play.php?id=85",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-9824-j85wb88o-v1/imageContent-9824-j85wb88o-m1.png",
    "genre": "Kannada"
  },
  {
    "name": "TV5 Kannada",
    "link": "https://rogstream.fun//hls/tata/play.php?id=629",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-19066-jczvvocg-v1/imageContent-19066-jczvvocg-m1.png",
    "genre": "Kannada"
  },
  {
    "name": "Public TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=33",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-98-j5fjwu7s-v2/imageContent-98-j5fjwu7s-m3.png",
    "genre": "Kannada"
  },
  {
    "name": "Raj News Kannada",
    "link": "https://rogstream.fun//hls/tata/play.php?id=510",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11948-j9ltrfc8-v1/imageContent-11948-j9ltrfc8-m1.png",
    "genre": "Kannada"
  },
  {
    "name": "Dighvijay 24x7 News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=342",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11397-j9jq7rxk-v1/imageContent-11397-j9jq7rxk-m1.png",
    "genre": "Kannada"
  },
  {
    "name": "Power TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=824",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-53941-kblrui1c-v3/imageContent-53941-kblrui1c-m3.png",
    "genre": "Kannada"
  },
  {
    "name": "News 1st Kannada",
    "link": "https://rogstream.fun//hls/tata/play.php?id=913",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/NWS1ST_Thumbnail-v3/NWS1ST_Thumbnail.png",
    "genre": "Kannada"
  },
  {
    "name": "Vistara News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1077",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/VISNEWS_Thumbnail-v2/VISNEWS_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "DD Chandana",
    "link": "https://rogstream.fun//hls/tata/play.php?id=321",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11338-j9jozdxs-v1/imageContent-11338-j9jozdxs-m1.png",
    "genre": "Kannada"
  },
  {
    "name": "Raj Musix Kannada",
    "link": "https://rogstream.fun//hls/tata/play.php?id=427",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11685-j9kipky0-v1/imageContent-11685-j9kipky0-m2.png",
    "genre": "Kannada"
  },
  {
    "name": "Public Music",
    "link": "https://rogstream.fun//hls/tata/play.php?id=424",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11676-j9khi3cw-v1/imageContent-11676-j9khi3cw-m1.png",
    "genre": "Kannada"
  },
  {
    "name": "Sri Sankara TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=477",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11788-j9kypa94-v1/imageContent-11788-j9kypa94-m2.png",
    "genre": "Kannada"
  },
  {
    "name": "Ayush TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=660",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11257-j9iy8oj4-v2/imageContent-11257-j9iy8oj4-m2.png",
    "genre": "Kannada"
  },
  {
    "name": "SVBC 3",
    "link": "https://rogstream.fun//hls/tata/play.php?id=971",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-69020-kzo676so-v1/imageContent-69020-kzo676so-m1.png",
    "genre": "Kannada"
  },
  {
    "name": "Colors Gujarati",
    "link": "https://rogstream.fun//hls/tata/play.php?id=107",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-340-j5jmiei0-v1/imageContent-340-j5jmiei0-m1.png",
    "genre": "Gujarati"
  },
  {
    "name": "Tata Play Gujarati Cinema",
    "link": "https://rogstream.fun//hls/tata/play.php?id=772",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-47224-k1uh5c4g-v2/imageContent-47224-k1uh5c4g-m2.png",
    "genre": "Gujarati"
  },
  {
    "name": "Colors Gujarati Cinema",
    "link": "https://rogstream.fun//hls/tata/play.php?id=692",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-35306-jqxwmaeo-v1/imageContent-35306-jqxwmaeo-m2.png",
    "genre": "Gujarati"
  },
  {
    "name": "TV9 Gujarati",
    "link": "https://rogstream.fun//hls/tata/play.php?id=489",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11874-j9l376yg-v2/imageContent-11874-j9l376yg-m2.png",
    "genre": "Gujarati"
  },
  {
    "name": "Gujarat Samachar TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=552",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12068-j9o9t9jc-v1/imageContent-12068-j9o9t9jc-m1.png",
    "genre": "Gujarati"
  },
  {
    "name": "Sandesh News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=586",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12167-j9phnio0-v2/imageContent-12167-j9phnio0-m2.png",
    "genre": "Gujarati"
  },
  {
    "name": "CNBC Bajaar",
    "link": "https://rogstream.fun//hls/tata/play.php?id=303",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11281-j9j4lf3k-v1/imageContent-11281-j9j4lf3k-m1.png",
    "genre": "Gujarati"
  },
  {
    "name": "News18 Gujarati",
    "link": "https://rogstream.fun//hls/tata/play.php?id=6",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-15-j5f97j2o-v4/imageContent-15-j5f97j2o-m3.png",
    "genre": "Gujarati"
  },
  {
    "name": "ABP Asmita",
    "link": "https://rogstream.fun//hls/tata/play.php?id=73",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-228-j5fxgzio-v2/imageContent-228-j5fxgzio-m2.png",
    "genre": "Gujarati"
  },
  {
    "name": "VTV News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=497",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11912-j9l58n8w-v2/imageContent-11912-j9l58n8w-m4.png",
    "genre": "Gujarati"
  },
  {
    "name": "Zee 24 Kalak",
    "link": "https://rogstream.fun//hls/tata/play.php?id=507",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11924-j9l5fs6w-v3/imageContent-11924-j9l5fs6w-m3.png",
    "genre": "Gujarati"
  },
  {
    "name": "India News Gujarat",
    "link": "https://rogstream.fun//hls/tata/play.php?id=654",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-26589-jj9w5h54-v1/imageContent-26589-jj9w5h54-m2.png",
    "genre": "Gujarati"
  },
  {
    "name": "Mantavya News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=891",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-61443-kmbf79yw-v1/imageContent-61443-kmbf79yw-m1.png",
    "genre": "News"
  },
  {
    "name": "1st Gujarat",
    "link": "https://rogstream.fun//hls/tata/play.php?id=974",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/1STGUJ_Thumbnail-v2/1STGUJ_Thumbnail.png",
    "genre": "Gujarati"
  },
  {
    "name": "DD Girnar",
    "link": "https://rogstream.fun//hls/tata/play.php?id=323",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11341-j9jpdu00-v1/imageContent-11341-j9jpdu00-m1.png",
    "genre": "Gujarati"
  },
  {
    "name": "Colors Oriya",
    "link": "https://rogstream.fun//hls/tata/play.php?id=115",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-355-j5jndjn4-v1/imageContent-355-j5jndjn4-m2.png",
    "genre": "Odia"
  },
  {
    "name": "Zee Sarthak",
    "link": "https://rogstream.fun//hls/tata/play.php?id=597",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12207-j9qlt5vk-v2/imageContent-12207-j9qlt5vk-m2.png",
    "genre": "Odia"
  },
  {
    "name": "Sidharth TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=978",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-69926-l1ooghuw-v3/imageContent-69926-l1ooghuw-m3.png",
    "genre": "Odia"
  },
  {
    "name": "Peppers TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=421",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11667-j9kdkb6w-v1/imageContent-11667-j9kdkb6w-m1.png",
    "genre": "Tamil"
  },
  {
    "name": "Kanak News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=376",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12271-j9to48zc-v1/imageContent-12271-j9to48zc-m2.png",
    "genre": "Odia"
  },
  {
    "name": "Kalinga TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=375",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11513-j9jtooug-v1/imageContent-11513-j9jtooug-m1.png",
    "genre": "Odia"
  },
  {
    "name": "News18 Odia",
    "link": "https://rogstream.fun//hls/tata/play.php?id=41",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-13243-jbirprzs-v1/imageContent-13243-jbirprzs-m1.png",
    "genre": "Odia"
  },
  {
    "name": "News7",
    "link": "https://rogstream.fun//hls/tata/play.php?id=448",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11733-j9kw4100-v4/imageContent-11733-j9kw4100-m4.png",
    "genre": "Odia"
  },
  {
    "name": "Nandighosha TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=805",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-49597-k6k8ts8g-v2/imageContent-49597-k6k8ts8g-m2.png",
    "genre": "Odia"
  },
  {
    "name": "Argus News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=907",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-62573-kop7duow-v6/imageContent-62573-kop7duow-m9.png",
    "genre": "Odia"
  },
  {
    "name": "DD Odia",
    "link": "https://rogstream.fun//hls/tata/play.php?id=333",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11366-j9jpvkw0-v1/imageContent-11366-j9jpvkw0-m1.png",
    "genre": "Odia"
  },
  {
    "name": "Surya TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=230",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-1005-j5nfiy9c-v1/imageContent-1005-j5nfiy9c-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "Asianet HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=292",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11246-j9iw5g74-v3/imageContent-11246-j9iw5g74-m3.png",
    "genre": "Malayalam"
  },
  {
    "name": "Asianet Plus",
    "link": "https://rogstream.fun//hls/tata/play.php?id=294",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11254-j9ixwlcg-v1/imageContent-11254-j9ixwlcg-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "Zee Keralam HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=694",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-35312-jqy146c8-v3/imageContent-35312-jqy146c8-m8.png",
    "genre": "Malayalam"
  },
  {
    "name": "Zee Keralam",
    "link": "https://rogstream.fun//hls/tata/play.php?id=684",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-9003-j7rlh2aw-v2/imageContent-9003-j7rlh2aw-m2.png",
    "genre": "Malayalam"
  },
  {
    "name": "Flowers",
    "link": "https://rogstream.fun//hls/tata/play.php?id=10",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-27-j5fbpbbs-v1/imageContent-27-j5fbpbbs-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "Amrita TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=178",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-762-j5m21xqw-v1/imageContent-762-j5m21xqw-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "Mazhavil Manorama",
    "link": "https://rogstream.fun//hls/tata/play.php?id=31",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-90-j5fjv4hk-v1/imageContent-90-j5fjv4hk-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "Kairali TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=25",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-33-j5fc5fko-v1/imageContent-33-j5fc5fko-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "Mazhavil Manorama HD",
    "link": "https://rogstream.fun//hls/tata/play.php?id=395",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11577-j9jv8mqg-v1/imageContent-11577-j9jv8mqg-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "Surya Movies",
    "link": "https://rogstream.fun//hls/tata/play.php?id=229",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-1008-j5nfqjeo-v1/imageContent-1008-j5nfqjeo-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "Asianet Movies",
    "link": "https://rogstream.fun//hls/tata/play.php?id=293",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11251-j9ixmaz4-v1/imageContent-11251-j9ixmaz4-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "We TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=553",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12065-j9o9phqg-v1/imageContent-12065-j9o9phqg-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "Safari TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=436",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11709-j9kt6bg0-v1/imageContent-11709-j9kt6bg0-m1.png",
    "genre": "Knowledge"
  },
  {
    "name": "Asianet News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=532",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12020-j9ntf12w-v1/imageContent-12020-j9ntf12w-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "Kairali News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=211",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-855-j5ma1eig-v2/imageContent-855-j5ma1eig-m2.png",
    "genre": "Malayalam"
  },
  {
    "name": "Manorama News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=87",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-279-j5jcwqts-v1/imageContent-279-j5jcwqts-m2.png",
    "genre": "Malayalam"
  },
  {
    "name": "Mathrubhumi News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=394",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11572-j9jv7wi0-v1/imageContent-11572-j9jv7wi0-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "News18 Kerala",
    "link": "https://rogstream.fun//hls/tata/play.php?id=66",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12309-ja9ac8n4-v1/imageContent-12309-ja9ac8n4-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "Jaihind TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=370",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11495-j9jtg2ns-v2/imageContent-11495-j9jtg2ns-m2.png",
    "genre": "Malayalam"
  },
  {
    "name": "Media One",
    "link": "https://rogstream.fun//hls/tata/play.php?id=576",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12137-j9p7pmeo-v2/imageContent-12137-j9p7pmeo-m2.png",
    "genre": "Malayalam"
  },
  {
    "name": "Kaumudy TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=378",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11524-j9jtu07c-v2/imageContent-11524-j9jtu07c-m2.png",
    "genre": "Malayalam"
  },
  {
    "name": "Raj News Malayalam",
    "link": "https://rogstream.fun//hls/tata/play.php?id=428",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11694-j9kjc0tc-v1/imageContent-11694-j9kjc0tc-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "Janam TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=270",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11179-j9i3nhs8-v2/imageContent-11179-j9i3nhs8-m3.png",
    "genre": "Malayalam"
  },
  {
    "name": "Twenty Four",
    "link": "https://rogstream.fun//hls/tata/play.php?id=799",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-49022-k5lxarmw-v2/imageContent-49022-k5lxarmw-m2.png",
    "genre": "Malayalam"
  },
  {
    "name": "Raj Musix Malayalam",
    "link": "https://rogstream.fun//hls/tata/play.php?id=541",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11688-j9kiz1zs-v1/imageContent-11688-j9kiz1zs-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "Kappa TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=377",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11520-j9jtr0y8-v1/imageContent-11520-j9jtr0y8-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "Goodness",
    "link": "https://rogstream.fun//hls/tata/play.php?id=269",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11176-j9i3lv54-v2/imageContent-11176-j9i3lv54-m2.png",
    "genre": "Malayalam"
  },
  {
    "name": "God TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=356",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11449-j9jqxsns-v1/imageContent-11449-j9jqxsns-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "Shalom TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=457",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11772-j9kwsqns-v1/imageContent-11772-j9kwsqns-m1.png",
    "genre": "Malayalam"
  },
  {
    "name": "Harvest TV 24x7",
    "link": "https://rogstream.fun//hls/tata/play.php?id=348",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-35638-jruyxl0g-v4/imageContent-35638-jruyxl0g-m7.PNG",
    "genre": "Malayalam"
  },
  {
    "name": "Tata Play Punjab De Rang",
    "link": "https://rogstream.fun//hls/tata/play.php?id=571",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-68739-kywk7vq0-v1/imageContent-68739-kywk7vq0-m1.png",
    "genre": "Punjabi"
  },
  {
    "name": "Zee Punjabi",
    "link": "https://rogstream.fun//hls/tata/play.php?id=796",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-49009-k5g6nid4-v1/imageContent-49009-k5g6nid4-m1.png",
    "genre": "Punjabi"
  },
  {
    "name": "PTC Punjabi",
    "link": "https://rogstream.fun//hls/tata/play.php?id=122",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-408-j5jr93k8-v1/imageContent-408-j5jr93k8-m1.png",
    "genre": "Punjabi"
  },
  {
    "name": "MH One",
    "link": "https://rogstream.fun//hls/tata/play.php?id=399",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11586-j9jyt26w-v1/imageContent-11586-j9jyt26w-m1.png",
    "genre": "Music"
  },
  {
    "name": "TABBAR HITS",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1083",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/TABHIT_Thumbnail-v2/TABHIT_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "Pitaara TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=624",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12931-jb64agig-v1/imageContent-12931-jb64agig-m1.png",
    "genre": "Punjabi"
  },
  {
    "name": "PTC Punjabi Gold",
    "link": "https://rogstream.fun//hls/tata/play.php?id=794",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-48860-k4nlxy14-v2/imageContent-48860-k4nlxy14-m2.png",
    "genre": "Punjabi"
  },
  {
    "name": "Zee Punjab Haryana Himachal Pradesh",
    "link": "https://rogstream.fun//hls/tata/play.php?id=580",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12155-j9pd5ug8-v1/imageContent-12155-j9pd5ug8-m1.png",
    "genre": "Punjabi"
  },
  {
    "name": "PTC News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=91",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-300-j5ji8seo-v1/imageContent-300-j5ji8seo-m1.png",
    "genre": "Punjabi"
  },
  {
    "name": "News18 Punjab Haryana",
    "link": "https://rogstream.fun//hls/tata/play.php?id=232",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-1011-j5ng33kg-v2/imageContent-1011-j5ng33kg-m2.png",
    "genre": "News"
  },
  {
    "name": "India News Punjab",
    "link": "https://rogstream.fun//hls/tata/play.php?id=366",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11482-j9jsdwfk-v1/imageContent-11482-j9jsdwfk-m1.jpg",
    "genre": "Punjabi"
  },
  {
    "name": "India News Haryana",
    "link": "https://rogstream.fun//hls/tata/play.php?id=13",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-8968-j7rcix2w-v1/imageContent-8968-j7rcix2w-m1.png",
    "genre": "News"
  },
  {
    "name": "Janta TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=371",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11502-j9jtlmi8-v3/imageContent-11502-j9jtlmi8-m4.png",
    "genre": "News"
  },
  {
    "name": "MH One News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=401",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11587-j9jyvxl4-v1/imageContent-11587-j9jyvxl4-m1.png",
    "genre": "News"
  },
  {
    "name": "Living India News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=793",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-57445-kg7p5yk8-v2/imageContent-57445-kg7p5yk8-m3.png",
    "genre": "Punjabi"
  },
  {
    "name": "Khabar Fast",
    "link": "https://rogstream.fun//hls/tata/play.php?id=972",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-69238-kzt9ig3k-v1/imageContent-69238-kzt9ig3k-m2.png",
    "genre": "Hindi"
  },
  {
    "name": "Zee Delhi NCR Haryana",
    "link": "https://rogstream.fun//hls/tata/play.php?id=262",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-35649-jrz9mhq0-v2/imageContent-35649-jrz9mhq0-m2.PNG",
    "genre": "Odia"
  },
  {
    "name": "STV Haryana News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=989",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-71064-l427mi60-v2/imageContent-71064-l427mi60-m3.png",
    "genre": "null"
  },
  {
    "name": "Daily Post Punjab Haryana Himachal",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1078",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/DAILYP_Thumbnail-v3/DAILYP_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "WPN World Punjabi TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1085",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/WTVPunjabi_Thumbnail-v1/WTVPunjabi_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "DD Punjabi",
    "link": "https://rogstream.fun//hls/tata/play.php?id=335",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11374-j9jpwgiw-v1/imageContent-11374-j9jpwgiw-m1.png",
    "genre": "Punjabi"
  },
  {
    "name": "PTC Chak De",
    "link": "https://rogstream.fun//hls/tata/play.php?id=92",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-297-j5ji8qv4-v1/imageContent-297-j5ji8qv4-m1.png",
    "genre": "Punjabi"
  },
  {
    "name": "PTC Music",
    "link": "https://rogstream.fun//hls/tata/play.php?id=922",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-63498-kqq89g5k-v2/imageContent-63498-kqq89g5k-m3.png",
    "genre": "Punjabi"
  },
  {
    "name": "Chardikla Time TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=193",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-792-j5m75dww-v1/imageContent-792-j5m75dww-m1.png",
    "genre": "Punjabi"
  },
  {
    "name": "PTC Simran",
    "link": "https://rogstream.fun//hls/tata/play.php?id=773",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-47242-k1vmve74-v1/imageContent-47242-k1vmve74-m2.png",
    "genre": "Punjabi"
  },
  {
    "name": "Nepal 1",
    "link": "https://rogstream.fun//hls/tata/play.php?id=567",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12106-j9ooxpug-v2/imageContent-12106-j9ooxpug-m2.png",
    "genre": "Others"
  },
  {
    "name": "Rengoni TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=214",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-860-j5mao56o-v1/imageContent-860-j5mao56o-m1.png",
    "genre": "Others"
  },
  {
    "name": "Indradhanu",
    "link": "https://rogstream.fun//hls/tata/play.php?id=368",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11485-j9jsysqw-v1/imageContent-11485-j9jsysqw-m1.png",
    "genre": "Others"
  },
  {
    "name": "DY 365",
    "link": "https://rogstream.fun//hls/tata/play.php?id=353",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11411-j9jqocds-v1/imageContent-11411-j9jqocds-m1.png",
    "genre": "Others"
  },
  {
    "name": "News18 Assam North East",
    "link": "https://rogstream.fun//hls/tata/play.php?id=68",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-12390-jai3fgy0-v1/imageContent-12390-jai3fgy0-m1.png",
    "genre": "Others"
  },
  {
    "name": "NEWS LIVE",
    "link": "https://rogstream.fun//hls/tata/play.php?id=226",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-956-j5mipkxs-v1/imageContent-956-j5mipkxs-m1.png",
    "genre": "Others"
  },
  {
    "name": "Protidin Time",
    "link": "https://rogstream.fun//hls/tata/play.php?id=221",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-879-j5mdmlqw-v1/imageContent-879-j5mdmlqw-m1.png",
    "genre": "Others"
  },
  {
    "name": "Prag News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=423",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11673-j9kfe73s-v1/imageContent-11673-j9kfe73s-m1.png",
    "genre": "Others"
  },
  {
    "name": "North East Live",
    "link": "https://rogstream.fun//hls/tata/play.php?id=417",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11644-j9k1gjow-v1/imageContent-11644-j9k1gjow-m1.png",
    "genre": "Others"
  },
  {
    "name": "NKTV PLUS",
    "link": "https://rogstream.fun//hls/tata/play.php?id=894",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/NKTVPLUS_Thumbnail-v2/NKTVPLUS_Thumbnail.png",
    "genre": "Others"
  },
  {
    "name": "ND24-Newsdaily.in",
    "link": "https://rogstream.fun//hls/tata/play.php?id=977",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-69824-l1dfgvcg-v1/imageContent-69824-l1dfgvcg-m3.png",
    "genre": "Others"
  },
  {
    "name": "NB NEWS",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1076",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/NBNEWS_Thumbnail-v2/NBNEWS_Thumbnail.png",
    "genre": "null"
  },
  {
    "name": "DD North East",
    "link": "https://rogstream.fun//hls/tata/play.php?id=331",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11363-j9jprc40-v1/imageContent-11363-j9jprc40-m1.png",
    "genre": "Others"
  },
  {
    "name": "DD Kashir",
    "link": "https://rogstream.fun//hls/tata/play.php?id=325",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11345-j9jpidkw-v1/imageContent-11345-j9jpidkw-m1.png",
    "genre": "Others"
  },
  {
    "name": "Ramdhenu",
    "link": "https://rogstream.fun//hls/tata/play.php?id=449",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-11728-j9kw1zp4-v1/imageContent-11728-j9kw1zp4-m1.png",
    "genre": "Others"
  },
  {
    "name": "DD Meghalaya",
    "link": "https://rogstream.fun//hls/tata/play.php?id=991",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-71497-l4stx05k-v3/imageContent-71497-l4stx05k-m5.png",
    "genre": "null"
  },
  {
    "name": "Ek Onkar",
    "link": "https://rogstream.fun//hls/tata/play.php?id=790",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-48807-k4m7ud0o-v2/imageContent-48807-k4m7ud0o-m2.png",
    "genre": "Punjabi"
  },
  {
    "name": "Fateh TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=791",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-48804-k4m7nz2w-v1/imageContent-48804-k4m7nz2w-m1.png",
    "genre": "Punjabi"
  },
  {
    "name": "Haryana Beats",
    "link": "https://rogstream.fun//hls/tata/play.php?id=912",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-62613-koxw351c-v1/imageContent-62613-koxw351c-m1.png",
    "genre": "Music"
  },
  {
    "name": "In Goa 24x7",
    "link": "https://rogstream.fun//hls/tata/play.php?id=807",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-49633-k6lny7mw-v1/imageContent-49633-k6lny7mw-m1.png",
    "genre": "Marathi"
  },
  {
    "name": "Namma TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=782",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-47615-k2jwaqug-v1/imageContent-47615-k2jwaqug-m3.png",
    "genre": "Kannada"
  },
  {
    "name": "Bangla Bhakti",
    "link": "https://rogstream.fun//hls/tata/play.php?id=865",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-57966-khtydrm8-v1/imageContent-57966-khtydrm8-m1.png",
    "genre": "Bengali"
  },
  {
    "name": "Tara News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=866",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-57969-khu0zsnc-v1/imageContent-57969-khu0zsnc-m1.png",
    "genre": "Bengali"
  },
  {
    "name": "Garv Gurbani",
    "link": "https://rogstream.fun//hls/tata/play.php?id=887",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-60927-kljn2znc-v1/imageContent-60927-kljn2znc-m1.png",
    "genre": "Spiritual"
  },
  {
    "name": "Swar Shree",
    "link": "https://rogstream.fun//hls/tata/play.php?id=883",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-60933-kljn4c9c-v1/imageContent-60933-kljn4c9c-m1.png",
    "genre": "Spiritual"
  },
  {
    "name": "Valam TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=893",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-61633-kmlgwaww-v1/imageContent-61633-kmlgwaww-m2.png",
    "genre": "News"
  },
  {
    "name": "Mahakaleshwar Temple, Ujjain",
    "link": "https://rogstream.fun//hls/tata/play.php?id=841",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/MAHAK_Thumbnail-v3/MAHAK_Thumbnail.png",
    "genre": "Spiritual"
  },
  {
    "name": "Samara News",
    "link": "https://rogstream.fun//hls/tata/play.php?id=892",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-61629-kmlgw3yw-v2/imageContent-61629-kmlgw3yw-m5.png",
    "genre": "News"
  },
  {
    "name": "Live Iskcon Vrindavan",
    "link": "https://rogstream.fun//hls/tata/play.php?id=944",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-65966-ku6es6wo-v4/imageContent-65966-ku6es6wo-m4.png",
    "genre": "Spiritual"
  },
  {
    "name": "Live Patna Sahib Patna",
    "link": "https://rogstream.fun//hls/tata/play.php?id=947",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-65977-ku6g2hko-v3/imageContent-65977-ku6g2hko-m3.png",
    "genre": "Spiritual"
  },
  {
    "name": "Live Sri Naga Sai Mandir Coimbatore",
    "link": "https://rogstream.fun//hls/tata/play.php?id=949",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-65985-ku6gly6w-v3/imageContent-65985-ku6gly6w-m3.png",
    "genre": "Spiritual"
  },
  {
    "name": "Lakshya TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=961",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-67703-kwog27o0-v1/imageContent-67703-kwog27o0-m1.png",
    "genre": "Gujarati"
  },
  {
    "name": "Shri Ganga Aarti, Varanasi",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1004",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-72303-l6hqin9k-v1/imageContent-72303-l6hqin9k-m2.png",
    "genre": "null"
  },
  {
    "name": "Shri ISKCON Girgaon, Mumbai",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1008",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-72306-l6hr8log-v1/imageContent-72306-l6hr8log-m1.png",
    "genre": "null"
  },
  {
    "name": "Shri Ashtavinayak Mahaganpati, Ranjangaon",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1007",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-72309-l6hro3js-v1/imageContent-72309-l6hro3js-m2.png",
    "genre": "null"
  },
  {
    "name": "Shri Mahalaxmi Temple, Mumbai",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1006",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-72312-l6hsvue8-v1/imageContent-72312-l6hsvue8-m2.png",
    "genre": "null"
  },
  {
    "name": "Shri Babulnaath Temple, Mumbai",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1005",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-72315-l6htmt2o-v1/imageContent-72315-l6htmt2o-m2.png",
    "genre": "null"
  },
  {
    "name": "Daiji World TV",
    "link": "https://rogstream.fun//hls/tata/play.php?id=1023",
    "image": "https://ltsk-cdn.s3.eu-west-1.amazonaws.com/jumpstart/Temp_Live/cdn/HLS/Channel/imageContent-73014-l8gzok74-v2/imageContent-73014-l8gzok74-m4.png",
    "genre": "Knowledge"
  },
  {
    "name": "HBO",
    "link": "https://boomtv.co.in/ayna/player.php?c=hbo",
    "header": "https://mhdtvworld.me/",
    "mhd": true,
    "image": "https://mhdtvworld.me/wp-content/uploads/2023/06/HBO.webp",
    "genre": "Movies"
  },
  {
    "name": "Star Sports 1 HD",
    "link": "https://faketv.co.in/criclife/starsp1hindi.php",
    "header": "https://mhdtvworld.me/",
    "mhd": true,
    "image": "https://mhdtvworld.me/wp-content/uploads/2020/02/star1hd.webp",
    "genre": "Sports"
  }
]
`
