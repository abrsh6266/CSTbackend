package model

type User struct {
	Id           int `json:"id"`
	Username     string `json:"username"`
	Phone        string `json:"phone"`
	Work         string `json:"work"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Token        string `json:"token"`
	Messages     []Message `json:"Messages"`
}
type Movie struct {
	Id           int 	`json:"id"`
	Title        string `json:"title"`
	Duration     string `json:"duration"`
	Description	 string `json:"description"`
	ReleasedYear string `json:"releasedYear"`
	Poster		 string `json:"poster"`
	Genre        string		 `json:"genre"`
	Country      string `json:"country"`
	Casts        []Cast   `json:"Casts"`
	DirectorMovieRels        []DirectorMovieRel      `json:"DirectorMovieRels"`
	Bookmarks    []Bookmark  `json:"Bookmarks"`
	Ratings      []Rating    `json:"Ratings"`
	Schedules    []Schedule  `json:"Schedules"`
	Ratings_aggregates Ratings_aggregate  `json:"Ratings_aggregate"`

}
type Ratings_aggregate  struct{
	Aggregate struct{
		Count int `json:"count"`
		Avg struct{
			Rate float32 `json:"rate"`
		} `json:"avg"`
	} `json:"aggregate"`
}
type Bookmark struct {
	UserId     	 int    `json:"userId"`
	Id       	 int    `json:"id"`
	MovieId 	 int    `json:"movieId"`
	Movies 	     Movie  `json:"Movie"`
}
type Star struct {
	Name     	 string		 `json:"name"`
	Id       	 int   		 `json:"id"`
	Casts        []Cast      `json:"Casts"`
}
type Cast struct {
	Id       	 int    `json:"id"`
	MovieId 	 int    `json:"movieId"`
	StarId       int    `json:"starId"`
	Movies       Movie  `json:"Movie"`
	Stars        Star   `json:"Star"`
}
type DirectorMovieRel struct {
	Id       	 	 int    	`json:"id"`
	MovieId 	 	 int    	`json:"movieId"`
	DirectorId       int    	`json:"directorId"`
	Movies           Movie  	`json:"Movie"`
	Directors        Director  	`json:"Director"`
}
type Rating struct {
	Rate         int    `json:"rate"`
	Id       	 int    `json:"id"`
	MovieId 	 int    `json:"movieId"`
	UserId   	 int    `json:"userId"`
}
type Director struct {
	Name     	 string 	 `json:"name"`
	Id       	 int    	 `json:"id"`
	DirectorMovieRels        []DirectorMovieRel      `json:"DirectorMovieRel"`
}
type Schedule struct {
	Day	     	 string 	 `json:"day"`
	Id       	 int  	     `json:"id"`
	MovieId 	 int   		 `json:"movieId"`
	StartTime 	 []string    `json:"startTime"`
	Movies       Movie       `json:"Movie"`
}
type Message struct{
	SentMessage	     	  string   `json:"sentMessage"`
	Id       	 		  int  	   `json:"id"`
	UserId	 		      int      `json:"userId"`
	UserEmail   	      string   `json:"userEmail"`
	RecievedMesssage 	  string   `json:"recievedMessage"`
	UserSeenStatus        bool     `json:"UserSeenStatus"`
	AdminSentStatus       bool     `json:"adminSentStatus"`
	Users                 User     `json:"User"`
}