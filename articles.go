package main

type article struct {
	Image        string
	Title        string
	Description  string
	Link         string
	CallToAction string
}

var articles = []article{
	{
		Image:        "article1.png",
		Title:        "Ringberg Talk on Turbulent Stellar Wind Bubbles",
		Description:  `This is a virtual talk I gave as part of a series in connection to a planned conference at the Ringberg Castle in Germany that was moved to a virtual format due to the COVID-19 Pandemic.`,
		Link:         "https://www.youtube.com/embed/GWHoX8WaShg?start=1781",
		CallToAction: "Watch virtual talk",
	},
	{
		Image:        "article2.png",
		Title:        "AAS Nova Article on Our Work on CN Cha",
		Description:  `The American Astronomical Society (AAS) has a public outreach website, Nova, which promotes work published in their journals for more general, public consumption. In March of 2021, Nova decided to feature our work on the Symbiotic Nova CN Cha. This work also piqued the interest of several amateur astronomers!`,
		Link:         "https://aasnova.org/2021/03/17/discovery-of-a-mystery-hidden-in-chamaeleon/",
		CallToAction: "Read article",
	},
	{
		Image:        "article3.png",
		Title:        "FDM on The Star Spot Podcast",
		Description:  `I was invited on to The Star Spot Podacst hosted by Justin Trottier to discuss my work with Dr. Philip Mocz on Fuzzy Dark Matter (FDM). The interview covers the background of this strange flavor of Dark Matter and the several works spear-headed by Philip on how cosmological structure is formed in this Dark Matter paradigm. One of these works went on to win the third place in the 2020 Buchalter Prize in Cosmology.`,
		Link:         "https://podcastaddict.com/episode/108055327",
		CallToAction: "Listen to podcast",
	},
	{
		Image:        "article4.png",
		Title:        "ChuTalk at Churchill College, Cambridge",
		Description:  `This is a talk I gave to fellow graduate students at Churchill College in Cambridge, UK in February of 2017 while I was in a masters program there. The talk was aimed at trying to explain some of the common questions I heard from friends about “The Universe.”`,
		Link:         "https://www.youtube.com/embed/DYZxKs3oWLQ",
		CallToAction: "Watch talk",
	},
}
