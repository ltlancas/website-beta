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
		Image: "article1.png",
		Title: "Ringberg Talk on Turbulent Stellar Wind Bubbles",
		Description: `I have broad research interests but my work in recent years has
			focused on the regulation of star
			formation within the
			dense, molecular clouds where stars form.<br><br>
			Specifically, my thesis has concentrated on the role that winds from massive stars play in this
			process.`,
		Link:         "/todo",
		CallToAction: "Watch virtual talk",
	},
	{
		Image:        "article2.png",
		Title:        "AAS Nova Article on Our Work on CN Cha",
		Description:  `The American Astronomical Society (AAS) has a public outreach website, Nova, which promotes work published in their journals for more general, public consumption. In March of 2021, Nova decided to feature our work on the Symbiotic Nova CN Cha. This work also piqued the interest of several amateur astronomers!`,
		Link:         "/todo",
		CallToAction: "Read article",
	},
	{
		Image:        "article3.png",
		Title:        "FDM on The Star Spot Podcast",
		Description:  `The American Astronomical Society (AAS) has a public outreach website, Nova, which promotes work published in their journals for more general, public consumption. In March of 2021, Nova decided to feature our work on the Symbiotic Nova CN Cha, which you can find the full paper on <a href="/todo">here</a>. This work also piqued the interest of several amateur astronomers!`,
		Link:         "/todo",
		CallToAction: "Listen to podcast",
	},
	{
		Image:        "article4.png",
		Title:        "ChuTalk at Churchill College, Cambridge",
		Description:  `The American Astronomical Society (AAS) has a public outreach website, Nova, which promotes work published in their journals for more general, public consumption. In March of 2021, Nova decided to feature our work on the Symbiotic Nova CN Cha. This work also piqued the interest of several amateur astronomers!`,
		Link:         "/todo",
		CallToAction: "Watch talk",
	},
}
