package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type movie struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Genre       string   `json:"genre"`
	Director    string   `json:"director"`
	Actors      []string `json:"actors"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	Poster      string   `json:"posterUri"`
}

var movies = []movie{
	{
		Name:        "Avatar",
		Genre:       "Action, Adventure, Fantasy",
		Director:    "James Cameron",
		Actors:      []string{"Sam Worthington, Zoe Saldana, Sigourney Weaver, Stephen Lang"},
		Description: "A paraplegic marine dispatched to the moon Pandora on a unique mission becomes torn between following his orders and protecting the world he feels is his home.",
		Poster:      "https://images-na.ssl-images-amazon.com/images/M/MV5BMjEyOTYyMzUxNl5BMl5BanBnXkFtZTcwNTg0MTUzNA@@._V1_SX1500_CR0,0,1500,999_AL_.jpg",
		Id:          "tt0499549",
		Type:        "movie",
	},
	{
		Name:        "I Am Legend",
		Genre:       "Drama, Horror, Sci-Fi",
		Director:    "Francis Lawrence",
		Actors:      []string{"Will Smith, Alice Braga, Charlie Tahan, Salli Richardson-Whitfield"},
		Description: "Years after a plague kills most of humanity and transforms the rest into monsters, the sole survivor in New York City struggles valiantly to find a cure.",
		Poster:      "https://images-na.ssl-images-amazon.com/images/M/MV5BMTI0NTI4NjE3NV5BMl5BanBnXkFtZTYwMDA0Nzc4._V1_.jpg",
		Id:          "tt0480249",
		Type:        "movie",
	},
	{
		Name:        "300",
		Genre:       "Action, Drama, Fantasy",
		Director:    "Zack Snyder",
		Actors:      []string{"Gerard Butler, Lena Headey, Dominic West, David Wenham"},
		Description: "King Leonidas of Sparta and a force of 300 men fight the Persians at Thermopylae in 480 B.C.",
		Poster:      "https://images-na.ssl-images-amazon.com/images/M/MV5BMTMwNTg5MzMwMV5BMl5BanBnXkFtZTcwMzA2NTIyMw@@._V1_SX1777_CR0,0,1777,937_AL_.jpg",
		Id:          "tt0416449",
		Type:        "movie",
	},
	{
		Name:        "The Avengers",
		Genre:       "Action, Sci-Fi, Thriller",
		Director:    "Joss Whedon",
		Actors:      []string{"Robert Downey Jr., Chris Evans, Mark Ruffalo, Chris Hemsworth"},
		Description: "Earth's mightiest heroes must come together and learn to fight as a team if they are to stop the mischievous Loki and his alien army from enslaving humanity.",
		Poster:      "https://images-na.ssl-images-amazon.com/images/M/MV5BMTA0NjY0NzE4OTReQTJeQWpwZ15BbWU3MDczODg2Nzc@._V1_SX1777_CR0,0,1777,999_AL_.jpg",
		Id:          "tt0848228",
		Type:        "movie",
	},
	{
		Name:        "The Wolf of Wall Street",
		Genre:       "Biography, Comedy, Crime",
		Director:    "Martin Scorsese",
		Actors:      []string{"Leonardo DiCaprio, Jonah Hill, Margot Robbie, Matthew McConaughey"},
		Description: "Based on the true story of Jordan Belfort, from his rise to a wealthy stock-broker living the high life to his fall involving crime, corruption and the federal government.",
		Poster:      "https://images-na.ssl-images-amazon.com/images/M/MV5BNDIwMDIxNzk3Ml5BMl5BanBnXkFtZTgwMTg0MzQ4MDE@._V1_SX1500_CR0,0,1500,999_AL_.jpg",
		Id:          "tt0993846",
		Type:        "movie",
	},
	{
		Name:        "Interstellar",
		Genre:       "Adventure, Drama, Sci-Fi",
		Director:    "Christopher Nolan",
		Actors:      []string{"Ellen Burstyn, Matthew McConaughey, Mackenzie Foy, John Lithgow"},
		Description: "A team of explorers travel through a wormhole in space in an attempt to ensure humanity's survival.",
		Poster:      "https://images-na.ssl-images-amazon.com/images/M/MV5BMjA3NTEwOTMxMV5BMl5BanBnXkFtZTgwMjMyODgxMzE@._V1_SX1500_CR0,0,1500,999_AL_.jpg",
		Id:          "tt0816692",
		Type:        "movie",
	},
	{
		Name:        "Game of Thrones",
		Genre:       "Adventure, Drama, Fantasy",
		Director:    "N/A",
		Actors:      []string{"Peter Dinklage, Lena Headey, Emilia Clarke, Kit Harington"},
		Description: "While a civil war brews between several noble families in Westeros, the children of the former rulers of the land attempt to rise up to power. Meanwhile a forgotten race, bent on destruction, plans to return after thousands of years in the North.",
		Poster:      "https://images-na.ssl-images-amazon.com/images/M/MV5BNDc1MGUyNzItNWRkOC00MjM1LWJjNjMtZTZlYWIxMGRmYzVlXkEyXkFqcGdeQXVyMzU3MDEyNjk@._V1_SX1777_CR0,0,1777,999_AL_.jpg",
		Id:          "tt0944947",
		Type:        "series",
	},
	{
		Name:        "Vikings",
		Genre:       "Action, Drama, History",
		Director:    "N/A",
		Actors:      []string{"Travis Fimmel, Clive Standen, Gustaf Skarsgård, Katheryn Winnick"},
		Description: "The world of the Vikings is brought to life through the journey of Ragnar Lothbrok, the first Viking to emerge from Norse legend and onto the pages of history - a man on the edge of myth.",
		Poster:      "https://images-na.ssl-images-amazon.com/images/M/MV5BMjM5MTM1ODUxNV5BMl5BanBnXkFtZTgwNTAzOTI2ODE@._V1_.jpg",
		Id:          "tt2306299",
		Type:        "series",
	},
	{
		Name:        "Gotham",
		Genre:       "Action, Crime, Drama",
		Director:    "N/A",
		Actors:      []string{"Ben McKenzie, Donal Logue, David Mazouz, Sean Pertwee"},
		Description: "The story behind Detective James Gordon's rise to prominence in Gotham City in the years before Batman's arrival.",
		Poster:      "https://images-na.ssl-images-amazon.com/images/M/MV5BNDI3ODYyODY4OV5BMl5BanBnXkFtZTgwNjE5NDMwMDI@._V1_SY1000_SX1500_AL_.jpg",
		Id:          "tt3749900",
		Type:        "series",
	},
	{
		Name:        "Power",
		Genre:       "Crime, Drama",
		Director:    "N/A",
		Actors:      []string{"Omari Hardwick, Joseph Sikora, Andy Bean, Lela Loren"},
		Description: "James \"Ghost\" St. Patrick, a wealthy New York night club owner who has it all, catering for the city's elite and dreaming big, lives a double life as a drug kingpin.",
		Poster:      "https://images-na.ssl-images-amazon.com/images/M/MV5BMTc2ODg0MzMzM15BMl5BanBnXkFtZTgwODYxODA5NTE@._V1_SY1000_SX1500_AL_.jpg",
		Id:          "tt3281796",
		Type:        "series",
	},
	{
		Name:        "Narcos",
		Genre:       "Biography, Crime, Drama",
		Director:    "N/A",
		Actors:      []string{"Wagner Moura, Boyd Holbrook, Pedro Pascal, Joanna Christie"},
		Description: "A chronicled look at the criminal exploits of Colombian drug lord Pablo Escobar.",
		Poster:      "https://images-na.ssl-images-amazon.com/images/M/MV5BMTk2MDMzMTc0MF5BMl5BanBnXkFtZTgwMTAyMzA1OTE@._V1_SX1500_CR0,0,1500,999_AL_.jpg",
		Id:          "tt2707408",
		Type:        "series",
	},
	{
		Name:        "Breaking Bad",
		Genre:       "Crime, Drama, Thriller",
		Director:    "N/A",
		Actors:      []string{"Bryan Cranston, Anna Gunn, Aaron Paul, Dean Norris"},
		Description: "A high school chemistry teacher diagnosed with inoperable lung cancer turns to manufacturing and selling methamphetamine in order to secure his family's financial future.",
		Poster:      "https://images-na.ssl-images-amazon.com/images/M/MV5BMTgyMzI5NDc5Nl5BMl5BanBnXkFtZTgwMjM0MTI2MDE@._V1_SY1000_CR0,0,1498,1000_AL_.jpg",
		Id:          "tt0903747",
		Type:        "series",
	},
	{
		Name:        "Doctor Strange",
		Genre:       "Action, Adventure, Fantasy",
		Director:    "Scott Derrickson",
		Actors:      []string{"Rachel McAdams, Benedict Cumberbatch, Mads Mikkelsen, Tilda Swinton"},
		Description: "After his career is destroyed, a brilliant but arrogant and conceited surgeon gets a new lease on life when a sorcerer takes him under her wing and trains him to defend the world against evil.",
		Poster:      "https://images-na.ssl-images-amazon.com/images/M/MV5BMjM3ODc1ODI5Ml5BMl5BanBnXkFtZTgwODMzMDY3OTE@._V1_.jpg",
		Id:          "tt1211837",
		Type:        "movie",
	},
	{
		Name:        "Rogue One: A Star Wars Story",
		Genre:       "Action, Adventure, Sci-Fi",
		Director:    "Gareth Edwards",
		Actors:      []string{"Felicity Jones, Riz Ahmed, Mads Mikkelsen, Ben Mendelsohn"},
		Description: "The Rebellion makes a risky move to steal the plans to the Death Star, setting up the epic saga to follow.",
		Poster:      "https://images-na.ssl-images-amazon.com/images/M/MV5BMjE3MzA4Nzk3NV5BMl5BanBnXkFtZTgwNjAxMTc1ODE@._V1_SX1777_CR0,0,1777,744_AL_.jpg",
		Id:          "tt3748528",
		Type:        "movie",
	},
	{
		Name:        "Assassin's Creed",
		Genre:       "Action, Adventure, Fantasy",
		Director:    "Justin Kurzel",
		Actors:      []string{"Michael Fassbender, Michael Kenneth Williams, Marion Cotillard, Jeremy Irons"},
		Description: "When Callum Lynch explores the memories of his ancestor Aguilar and gains the skills of a Master Assassin, he discovers he is a descendant of the secret Assassins society.",
		Poster:      "https://images-na.ssl-images-amazon.com/images/M/MV5BN2EyYzgyOWEtNTY2NS00NjRjLWJiNDYtMWViMjg5MWZjYjgzXkEyXkFqcGdeQXVyNjUwNzk3NDc@._V1_.jpg",
		Id:          "tt2094766",
		Type:        "movie",
	},
	{
		Name:        "Luke Cage",
		Genre:       "Action, Crime, Drama",
		Director:    "N/A",
		Actors:      []string{"Mahershala Ali, Mike Colter, Frankie Faison, Erik LaRay Harvey"},
		Description: "Given superstrength and durability by a sabotaged experiment, a wrongly accused man escapes prison to become a superhero for hire.",
		Poster:      "https://images-na.ssl-images-amazon.com/images/M/MV5BMjMxNjc1NjI0NV5BMl5BanBnXkFtZTgwNzA0NzY0ODE@._V1_SY1000_CR0,0,1497,1000_AL_.jpg",
		Id:          "tt3322314",
		Type:        "series",
	},
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/movies", getMovies)
	router.GET("/movies/:id", getMovieById)
	router.POST("/movies", postMovie)
	router.PUT("/movies/:id", putMovie)

	router.Run("localhost:8080")
}

func getMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}

func postMovie(c *gin.Context) {
	var newMovie movie

	if err := c.BindJSON(&newMovie); err != nil {
		return
	}

	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusCreated, newMovie)
}

func getMovieById(c *gin.Context) {
	id := c.Param("id")

	for _, m := range movies {
		if m.Id == id {
			c.IndentedJSON(http.StatusOK, m)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "movie not found"})
}

func putMovie(c *gin.Context) {
	id := c.Param("id")

	var updatedMovie movie
	if err := c.BindJSON(&updatedMovie); err != nil {
		return
	}

	for i, m := range movies {
		if m.Id == id {
			movies[i] = updatedMovie
		}
	}
}
