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
	Flag		string
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
		"Flag of Slovakia.svg",
	},
	{
		"muzikova",
		"Muzikova",
		"Mrs Muziková talks about her experience of being found guilty and imprisoned during the socialist era and the suppression of freedom of religion.",
		"5yRlDIDvHvE",
		"Flag of Slovakia.svg",
	},
	{
		"mikloskova",
		"An interview with Mrs Mikloskova",
		"Mrs Miklošková talks about her attendance at Candle Manifestation (a demonstration leading to the Velvet Revolution) which took place in Slovakia on March 25th, 1988.",
		"Qrh-_4WP4ZM",
		"Flag of Slovakia.svg",
	},
	{
		"hilton",
		"Views of an Ex British serviceman stationed in Germany",
		"Mr Hilton, a former UK army engineer who served in Germany for many years during the Cold War talks about exercises and preparedness for a possible Soviet nuclear attack."
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
		[]string{},
		[]Link{},
	},
	{
		"naranjito",
		"Naranjito",
		`The "Naranjito" was the mascot of the Football World Cup held in Spain in 1982. The character represented an orange, which is a typical fruit in Spain. It was wearing the national team equipment and a football on its left hand.`,
		"http://img2.rtve.es/i/?w=1600&i=1389699712798.jpg",
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
		[]string{},
		[]Link{},
	},
}
