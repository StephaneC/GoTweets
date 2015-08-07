package main

type Tweet struct {
	Timestamp        int64  `json:"timestamp_ms"`
	Id          string  `json:"id_str"`
	text         string  `json:"text"`
	Description string  `json:"description"`

}

/*


  source: '<a href="http://twitter.com/download/iphone" rel="nofollow">Twitter for iPhone</a>',
  truncated: false,
  in_reply_to_status_id: null,
  in_reply_to_status_id_str: null,
  in_reply_to_user_id: null,
  in_reply_to_user_id_str: null,
  in_reply_to_screen_name: null,
  user:
   { id: 1550405756,
     id_str: '1550405756',
     name: 'EDENHAZARD™⚡️',
     screen_name: 'ArturZakharyan7',
     location: '',
     url: null,
     description: 'jaime pas me melangerrrr',
     protected: false,
     verified: false,
     followers_count: 328,
     friends_count: 93,
     listed_count: 2,
     favourites_count: 1250,
     statuses_count: 8782,
     created_at: 'Thu Jun 27 11:25:04 +0000 2013',
     utc_offset: null,
     time_zone: null,
     geo_enabled: true,
     lang: 'fr',
     contributors_enabled: false,
     is_translator: false,
     profile_background_color: 'C0DEED',
     profile_background_image_url: 'http://abs.twimg.com/images/themes/theme1/bg.png',
     profile_background_image_url_https: 'https://abs.twimg.com/images/themes/theme1/bg.png',
     profile_background_tile: false,
     profile_link_color: '0084B4',
     profile_sidebar_border_color: 'C0DEED',
     profile_sidebar_fill_color: 'DDEEF6',
     profile_text_color: '333333',
     profile_use_background_image: true,
     profile_image_url: 'http://pbs.twimg.com/profile_images/623644822457466880/b89ihS49_normal.jpg',
     profile_image_url_https: 'https://pbs.twimg.com/profile_images/623644822457466880/b89ihS49_normal.jpg',
     profile_banner_url: 'https://pbs.twimg.com/profile_banners/1550405756/1435183747',
     default_profile: true,
     default_profile_image: false,
     following: null,
     follow_request_sent: null,
     notifications: null },
  geo: null,
  coordinates: null,
  place: null,
  contributors: null,
  retweeted_status:
   { created_at: 'Wed Jul 22 13:06:44 +0000 2015',
     id: 623841609457750000,
     id_str: '623841609457750016',
     text: 'Quand t\'es fier de ton fiston 🙈 #Benzema #AnaliciaChaves http://t.co/mcf6xqICno',
     source: '<a href="http://twitter.com/download/iphone" rel="nofollow">Twitter for iPhone</a>',
     truncated: false,
     in_reply_to_status_id: null,
     in_reply_to_status_id_str: null,
     in_reply_to_user_id: null,
     in_reply_to_user_id_str: null,
     in_reply_to_screen_name: null,
     user:
      { id: 60797084,
        id_str: '60797084',
        name: 'Bill',
        screen_name: 'Heisenberg_19',
        location: 'charming',
        url: 'http://www.savewalterwhite.com/',
        description: 'Be Water',
        protected: false,
        verified: false,
        followers_count: 550,
        friends_count: 243,
        listed_count: 18,
        favourites_count: 3551,
        statuses_count: 64105,
        created_at: 'Tue Jul 28 03:24:29 +0000 2009',
        utc_offset: 7200,
        time_zone: 'Paris',
        geo_enabled: true,
        lang: 'fr',
        contributors_enabled: false,
        is_translator: false,
        profile_background_color: '0099B9',
        profile_background_image_url: 'http://pbs.twimg.com/profile_background_images/550637444119273473/IFCt7_wq.jpeg',
        profile_background_image_url_https: 'https://pbs.twimg.com/profile_background_images/550637444119273473/IFCt7_wq.jpeg',
        profile_background_tile: false,
        profile_link_color: '0099B9',
        profile_sidebar_border_color: 'FFFFFF',
        profile_sidebar_fill_color: 'DEDEDE',
        profile_text_color: '2B0C13',
        profile_use_background_image: true,
        profile_image_url: 'http://pbs.twimg.com/profile_images/621747026678673408/PXBKJXiH_normal.jpg',
        profile_image_url_https: 'https://pbs.twimg.com/profile_images/621747026678673408/PXBKJXiH_normal.jpg',
        profile_banner_url: 'https://pbs.twimg.com/profile_banners/60797084/1432249816',
        default_profile: false,
        default_profile_image: false,
        following: null,
        follow_request_sent: null,
        notifications: null },
     geo: null,
     coordinates: null,
     place: null,
     contributors: null,
     retweet_count: 4,
     favorite_count: 0,
     entities:
      { hashtags: [Object],
        trends: [],
        urls: [],
        user_mentions: [],
        symbols: [],
        media: [Object] },
     extended_entities: { media: [Object] },
     favorited: false,
     retweeted: false,
     possibly_sensitive: false,
     filter_level: 'low',
     lang: 'fr' },
  retweet_count: 0,
  favorite_count: 0,
  entities:
   { hashtags: [ [Object], [Object] ],
     trends: [],
     urls: [],
     user_mentions: [ [Object] ],
     symbols: [],
     media: [ [Object] ] },
  extended_entities: { media: [ [Object] ] },

*/