package main

type research_post struct {
	ID          string
	Title       string
	Description string
	Date        string
	Link        string
}

var research_posts = []research_post{
	/*{
		ID:    "wandering-black-holes",
		Title: "Wandering Black Holes in the Milky Way",
		Description: `Working with Dr. Jenny Greene and a slew of other talented collaborators,
			we set out to try to find “wandering” Intermediate Mass Black Holes (IMBHs) that used
			to live at the center of dense star clusters that have since disrupted in the potential
			of our Milky Way Galaxy. Past works had postulated that such wandering IMBHs would retain
			several stars in their gravitational influence, remaining tightly bound around them in what
			would appear to be an extremely dense star cluster.`,
		Date: "August, 2021",
		Link: "https://ui.adsabs.harvard.edu/abs/2021ApJ...917...17G/abstract",
	},*/
	{
		ID:    "stellar-wind-bubbles",
		Title: "Efficiently Cooled Stellar Wind Bubbles",
		Description: `In the middle of April, 2021 I posted the first two major works of my thesis 
			outlining a new theory for the expansion of stellar wind bubbles from clusters of massive 
			stars in dense, turbulent molecular clouds (the sorts of places where these massive stars 
			form). They are now published in the Astrophysical Journal. You can find an abridged 
			explanation of these papers in a Twitter thread I wrote when I posted them.`,
		Date: "June, 2021",
		Link: "https://ui.adsabs.harvard.edu/abs/2021ApJ...914...89L/abstract",
	},
	{
		ID:    "cncha",
		Title: "A Mystery in Chamaeleon",
		Description: `In March of 2019 I was observing at Las Campanas Observatory in Chile. I was there 
			part of a research plan looking for exceptionally strange objects. When you go looking for 
			strange things, you end up finding some, not necessarily the type you were looking for.`,
		Date: "September, 2020",
		Link: "https://ui.adsabs.harvard.edu/abs/2020AJ....160..125L/abstract",
	},
	{
		ID:    "fdm-df",
		Title: "Dynamical Friction when Dark Matter is Fuzzy",
		Description: `Dynamical Friction is the strange process by which a massive body, moving through 
			a 'sea' of lighter particles in space, creates a 'wake' of the lighter particles. That wake, 
			formed through gravity, then gravitationally pulls on the massive body, slowing it down. We 
			investigated how this process worked in a universe where dark matter behaves quantum 
			mechanically.`,
		Date: "January, 2020",
		Link: "https://ui.adsabs.harvard.edu/abs/2020JCAP...01..001L/abstract",
	},
	{
		ID:    "beta-halo",
		Title: "Orbit's in the Milky Way's Stellar Halo",
		Description: `The European Space Agency's Gaia Satellite has given us the first real picture of 
			our Galaxy in motion. We used this excellent data to reveal the Milky Way's complicated past 
			of galactic cannibalism.`,
		Date: "June, 2019",
		Link: "https://ui.adsabs.harvard.edu/abs/2019MNRAS.486..378L/abstract",
	},
	{
		ID:    "tpcf-halo",
		Title: "Correlation Functions and Hidden Structure",
		Description: `When looking for small signals of the much-obscured past of the Milky Way's formation,
			it helps to use statistics. We investigate what correlation functions can tell us about the 
			hidden structure in the stellar halo of our Galaxy.`,
		Date: "April, 2019",
		Link: "https://ui.adsabs.harvard.edu/abs/2019MNRAS.484.2556L/abstract",
	},
}
