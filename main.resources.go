package main

/*Resource bla*/
type Resource struct {
	Interviews []Interview
	Artefacts  []Artefact
}

/*Interview bla*/
type Interview struct {
	Key         string
	Title       string
	Description string
	Link        string
	Flag        string
}

/*Interviews bla*/
var Interviews = []Interview{
	{
		"spain-moonlanding",
		"Spanish perception of the landing on the moon",
		"In this video,  a women and a men, from differents cities of Spain, explains their point of view about the landing on the moon.",
		"qoD9r_agYVs",
		"Flag_of_Spain.svg",
	},
	{
		"spain-franco-freedom",
		"Lack of freedom in Spain during Franco's dictatorship",
		"Three persons tells us how Franco’s dictatorship suppress freedom and rights of spanish people. They explain their own life and experiences.",
		"Dy1HKOXQ134",
		"Flag_of_Spain.svg",
	},
	{
		"spain-franco-school",
		"Going to school in Francoist Spain",
		"Three grandmother explain how was their daily life, specially in the school. They explain how was school during Franco´s dictatorship and what they learnt. We are able to see the differences between Franco´s school in Spain and the one of nowadays.",
		"gRyD_t1slJg",
		"Flag_of_Spain.svg",
	},
	{
		"colonial-conflicts-morocco",
		"Colonial conflicts during the Cold War Sidi Ifni Morocco",
		"In this video, a spanish men tells us  about his experience in the conflict in Morocco from his own point of view as a soldier. In Spain, they hadn’t any notice about this war.",
		"4j7p1epi_Dw",
		"Flag_of_Spain.svg",
	},
	{
		"coup-1981",
		"Failed coup d'etat in 1981 23 F",
		"A woman talks about the failed coup d’etat, her feelings and how Spain reacted to it.",
		"c0C1_hyDsJA",
		"Flag_of_Spain.svg",
	},
	{
		"spain-franco-child-3",
		"Being a child in Spain during the first years of Franco's dictatorship III",
		"A man tells us about his daily life and the differences between the past and the actual educative system.",
		"Atoj2myC0q0",
		"Flag_of_Spain.svg",
	},
	{
		"spain-franco-child-2",
		"Being a child in Spain during the first years of Franco's dictatorship II",
		"A spanish grandmother talks about her life and the difficulties that she had to face, specially during her childhood. She also tells us about the perception of religion she had.",
		"pq1Re5h-_IU",
		"Flag_of_Spain.svg",
	},
	{
		"spain-franco-child-1",
		"Being a child in Spain during the first years of Franco's dictatorship I",
		"A Spanish grandfather tells us what his school yeras were like, and how he became a high rank military man during Franco´s dicatorship.",
		"g9rFUFraBYY",
		"Flag_of_Spain.svg",
	},
	{
		"haffke-candels",
		"Dr. J. Haffke - Candels",
		"Dr. Jürgen Haffke, former history teacher and currently an author, recounts his memories during the German partition and how the people in West German showed compassion for the people living in East Germany.",
		"VNMymHyqgNQ",
		"Flag_of_Germany.svg",
	},
	{
		"haffke-cuba",
		"Dr. J. Haffke - Cuba Crisis",
		"Dr. Jürgen Haffke, former history teacher and currently an author, recounts his memories during the Cuba Crisis and explains protection measures in this time.",
		"gTP9rE6U1Ng",
		"Flag_of_Germany.svg",
	},
	{
		"haffke-sputnik",
		"Dr. J. Haffke - Sputnik",
		"Dr. Jürgen Haffke, former history teacher and currently an author, recounts his first experiences with the Cold War at the age of four.",
		"ZcspzbTc8yQ",
		"Flag_of_Germany.svg",
	},
	{
		"ewald-full",
		"German Interviews: Ewald (Full Version)",
		"Mr. Ewald, a former boarder patrol officer of the Federal Republic of Germany, talks about his time working on the boarder to the German Democratic Republic during the Cold War. (Full Interview)",
		"RGYfdVMCkEw",
		"Flag_of_Germany.svg",
	},
	{
		"beckmann",
		"German Interview with Mrs. Beckmann",
		"Mrs Beckmann talks about her experience in the GDR.",
		"XbmqcL4mzv4",
		"Flag_of_Germany.svg",
	},
	{
		"ewald-cut",
		"German Interviews: Ewald (Cut Version)",
		"Mr. Ewald, a former boarder patrol officer of the Federal Republic of Germany, talks about his time working on the boarder to the German Democratic Republic during the Cold War. (Teaser)",
		"hW2lSJngFYk",
		"Flag_of_Germany.svg",
	},
	{
		"zak",
		"Žák",
		"Mr Žák talks about religion during the Cold War. He mentions how people suffered for being Christians and reading religious literature. ",
		"cI-29USjpM4",
		"Flag_of_Slovakia.svg",
	},
	{
		"muzikova",
		"Muzikova",
		"Mrs Muziková talks about her experience of being found guilty and imprisoned during the socialist era and the suppression of freedom of religion.",
		"5yRlDIDvHvE",
		"Flag_of_Slovakia.svg",
	},
	{
		"mikloskova",
		"An interview with Mrs Mikloskova",
		"Mrs Miklošková talks about her attendance at Candle Manifestation (a demonstration leading to the Velvet Revolution) which took place in Slovakia on March 25th, 1988.",
		"Qrh-_4WP4ZM",
		"Flag_of_Slovakia.svg",
	},
	{
		"zaborska",
		"An interview with Mrs Anna Zaborska",
		"Mrs Záborská, an ex-MEP, describes a fearful life of a dissident in Czechoslovakia during the Cold war. She experienced imprisoning of her father as a result of the lack of freedom of speech.",
		"K6kqpzMcpMM",
		"Flag_of_Slovakia.svg",
	},
	{
		"hilton",
		"Views of an Ex British serviceman stationed in Germany",
		"Mr Hilton, a former UK army engineer who served in Germany for many years during the Cold War talks about exercises and preparedness for a possible Soviet nuclear attack.",
		"DB5Z25jWxBs",
		"Flag_of_Wales.svg",
	},
}

/*Link bla*/
type Link struct {
	Name    string
	Pointer string
}

/*Artefact bla*/
type Artefact struct {
	Key      string
	Name     string
	Text     string
	Photo    string
	Source   string
	AddPhoto []string
	Links    []Link
}

/*Artefacts bla*/
var Artefacts = []Artefact{
	{
		"mantilla",
		"Mantilla",
		`The Spanish “mantilla” is a piece of cloth to cover women’s hair when going to church. It was mandatory at the times of Franco’s regime and a symbol of religious oppression. It is also a typical garment in Spanish culture. Historically, it has been used as a religious symbol, in processions and other festivities such as Holy Week. Moreover, it is also used in different events like bullfighting games. At present, this piece of cloth is a part of a lot of regional costumes from Spain.`,
		"https://img2.rtve.es/imagenes/carcel-especial-calzada-oropesa/1536310343441.jpg",
		"",
		[]string{"https://live.staticflickr.com/7606/16485895504_9b9e780b9d_b.jpg"},
		[]Link{
			{
				"Article about religion",
				"/article/social/religion",
			},
		},
	},
	{
		"franco-flag",
		"Francoist Flag",
		`This flag started to be used in 1938, during the Spanish Civil War. Apart from the traditional Spanish national coat of arms, and the emblem “Una, Grande y Libre” (“United, Great and Free”), there are some elements that were included to link Franco’s regime to the Catholic Monarchs’ glory and ideals. These elements are St. John’s Eagle and the yoke and bundle of arrows.`,
		"https://upload.wikimedia.org/wikipedia/commons/3/33/Flag_of_Spain_%281945%E2%80%931977%29.svg",
		"",
		[]string{},
		[]Link{},
	},
	{
		"naranjito",
		"Naranjito",
		`The "Naranjito" was the mascot of the Football World Cup held in Spain in 1982. The character represented an orange, which is a typical fruit in Spain. It was wearing the national team equipment and a football on its left hand.`,
		"http://img2.rtve.es/i/?w=1600&i=1389699712798.jpg",
		"",
		[]string{},
		[]Link{
			{
				"Articel about Sport",
				"/article/social/sports",
			},
		},
	},
	{
		"seat",
		"Seat 600",
		`The Seat 600 was a car produced by the Spanish manufacturer SEAT between 1957-1973. It is probably one of the most emblematic and valued cars in the whole SEAT production because it was the first car which everyone could afford. It is a symbol of the Spanish economic development at that time.`,
		"https://upload.wikimedia.org/wikipedia/commons/7/7d/Seat_600_%281%29.jpg",
		"",
		[]string{},
		[]Link{},
	},
	{
		"profumo-scandal",
		"The Profumo Scandal",
		`The Profumo scandal refers to the affair between John Profumo Britain’s Secretary of State for War and a 19-year-old model Christine Keeler. Such a sex scandal in itself was scandalous especially in the 1960s. However, Profumo initially denied any impropriety and later it emerged that Miss Keeler was also involved with a Russian naval officer Yevgeny Ivanov.
		 This fact meant that a real possibility existed that state secrets could be passed from Profumo to Keeler to Ivanov and onto the Soviet Government. Given the timing of the affair with tensions in Suez, Korea and Berlin much on the minds of the public, the possibility that the Russians were setting was may now be called a “honey trap” was a distinct possibility.
		 The scandal grew when Stephen Ward committed suicide as he was the one who had introduced the parties and had been charged with living off immoral earnings (essentially organising prostitution). The scandal therefore had many elements and threatened the British Government. As the years have passed the parties are now deceased but all denied any wrongdoing, even the charges against Ward may have been trumped up according to some conspiracy theorists.
		 The artefact in this case is the headline from the Daily Mail newspaper. As one of the most read newspapers in the UK it may be seen to be quite influential in helping to shape public opinion and could help to harden anti Soviet feelings thus allowing Government rhetoric on the Cold War to have more weight and help to harden the UKs stance against Communism.`,
		"https://i.dailymail.co.uk/i/pix/2010/07/11/article-0-000EF24500000258-359_468x286.jpg",
		"https://www.dailymail.co.uk/travel/article-1294019/Profumo-affair-scandal-tour-offered-Cliveden-House.html",
		[]string{},
		[]Link{},
	},
	{
		"atomic-bunker",
		"Nuclear fallout bunker",
		`As part of the trip to Wales, our group visited the so-called “secret nuclear bunker” near Nantwich in Cheshire, https://hackgreen.co.uk/
		 The British Government wished to ensure that should nuclear war were to break out, there would still be a Government in operation that could carry out key roles.
		 These would include coordinating defence, distributing food, ensuring law and order prevail. Each region had a fallout shelter set up where they could coordinate local and regional government from.
		 The shelter is set underground over several floors. Radiation suits, radio and listening devices, dormitories, kitchens and everything else needed to survive is present. The site is surprisingly large and allows for many people to be kept safe until radiation levels fell to a safer level.
		 Hack Green also served as a secret radar detection site and so was off limits to the public for most of its operational life.`,
		"/assets/img/atomic-bunker.png",
		"",
		[]string{},
		[]Link{},
	},
	{
		"russia-love-poster",
		"From Russia with Love Movie Poster 1963",
		`This iconic movie poster from the 1963 movie of the same name demonstrates how Russia was seen as a hostile enemy in the 1960's. With Sean Connery in the lead role, we are led to believe that the Soviets are planning espionage whereas in fact there is a positive message as the Soviet spy Romanova ends up working with the Mi6 agent Bond to work against Spectre who realise that they can use distrust between the sides in the Cold War for their own gain. By working together, the east & West could defeat the plot of Spectre. Despite this feel good message, the overall tone is that despite one good agent, the rest of the Soviets were not to be trusted and that they were quite a cold and inhumane culture especially the KGB. Using popular culture to perpetuate and strengthen a view of the Soviets as a culture not to be trusted helps to solidify a view held by western citizens and no doubt helps to give force to a western governmental view that an arms race was fully justified.`,
		"https://upload.wikimedia.org/wikipedia/en/a/ad/From_Russia_with_Love_%E2%80%93_UK_cinema_poster.jpg",
		"",
		[]string{},
		[]Link{},
	},
	{
		"watchtower",
		"BT6 GDR Watchtower in Erna",
		`BT6 GDR Watchtower in Erna-Berger-Straße, Berlin. 302 watchtowers in total secured the 155km long border to West-Berlin. This specific concrete tower was built in 1966. It was manned at all times by two border patrol soldiers working 8h shifts watching the area in front of the border. Nowadays it is the last still existing tower of its kind and an official historical monument since 2001.`,
		"/assets/img/wachturm.jpg",
		"",
		[]string{
			"/assets/img/wachturm-1.jpg",
		},
		[]Link{},
	},
	{
		"vw-bus",
		"VW Type 2 (VW Bus)",
		`The shown car is a modified version of the „Volkswagen Type 2“, which is known as the „VW Bus” in
		Germany. The Type 2 was first produced in 1949 and was designed as a transporter. The shown Type
		2 was build and modified in 1966, a few years later it was shipped over to California.
		It was designed as a sign of the upcoming counterculture and freedom movement in the 1960s and
		1970s. The white peace sign and the rainbow are some signature signs of this anti-establishment
		movement other popular symbols from this time are for example bright colours like pink, green or
		yellow as well as the marihuana plant which is probably the plant you can see in the cars side
		windows. Overall plants and flowers have been a symbol too.
		The main goal of the counterculture and freedom movement was to protest against the current
		political situation, where you encountered wars and conflicts on an everyday base in the news. The
		best example for this is probably the Vietnam War.
		Besides this Type 2 transporter. There were also many other each individually redesigned Type 2
		transporters during this time period. That’s the reason why the car itself transformed into a sign of
		the Hippie movement, which is still well known today.`,
		"/assets/img/vw-bus.png",
		"",
		[]string{},
		[]Link{},
	},
	{
		"border-security",
		"Border Patrol training",
		`Training of the prospective border patrol officers of the FRD / West Germany. Part of the training was to practice street fight scenario.  case Russia / the Soviet Union invaded West Germany, the border patrol was supposed to tie the enemy until the Federal Armed Forces (German Army) arrived.`,
		"/assets/img/border-security.jpg",
		"",
		[]string{},
		[]Link{},
	},
	{
		"babetta",
		"Babetta",
		`Babetta (also known as Jawa in other countries) was a moped built in Czechoslovakia. It was manufactured in Považská Bystrica on the West and Kolárovo on the East of Slovakia. The vehicle was named after a popular Czech song „Babeta“ written by Jiří Suchý. The first Babettas appeared in 1970. Slovak engineers changed the construction of the moped after unexpectedly low incomes in 1973. After these changes, Babetta´s popularity skyrocketed not only in Czechoslovakia but in Germany, Netherlands and the USA, too. The manufactory stopped in 1997 after Dissolution od Czechoslovakia and a massive privatization, following.`,
		"/assets/img/artefacts/babetta.png",
		"",
		[]string{
			"/assets/img/artefacts/babetta-1.jpg",
		},
		[]Link{},
	},
}
