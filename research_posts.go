package main

type research_post struct {
	ID          string
	Title       string
	Description string
	Date        string
	Link        string
}

var research_posts = []research_post{
	{
		ID:    "wandering-black-holes",
		Title: "Wandering Black Holes in the Milky Way",
		Description: `Working with Dr. Jenny Greene and a slew of other talented collaborators, 
			we set out to try to find “wandering” Intermediate Mass Black Holes (IMBHs) that used 
			to live at the center of dense star clusters that have since disrupted in the potential 
			of our Milky Way Galaxy. Past works had postulated that such wandering IMBHs would retain 
			several stars in their gravitational influence, remaining tightly bound around them in what 
			would appear to be an extremely dense star cluster.`,
		Date: "August 2, 2021",
		Link: "/research/wandering-black-holes",
	},
	{
		ID:    "stellar-wind-bubbles",
		Title: "Efficiently Cooled Stellar Wind Bubbles",
		Description: `In the middle of April, 2021 I released the first two major works of my thesis 
			outlining a new theory for the expansion of stellar wind bubbles from clusters of massive 
			stars in dense, turbulent molecular clouds (the sorts of places where these massive stars 
			form). They are now published in the Astrophysical Journal. You can find an abridged 
			explanation of these papers in a Twitter thread I wrote when I posted them.`,
		Date: "August 2, 2021",
		Link: "/research/stellar-wind-bubble",
	},
}
