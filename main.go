package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Driver struct {
	gorm.Model
	Name    string
	License string
	Cars    []Car
}

type Car struct {
	gorm.Model
	Year      int
	Make      string
	ModelName string
	DriverID  int
}

/** our database types */

type User struct {
	ID       int `gorm:"primaryKey"`
	Username string
	Password string
}

type Image struct {
	ID          int `gorm:"primaryKey"` // This version of gorm does not support a basic Date SQL type. What the fuck. Also, tried adding gorm.Model, maybe we'll siphon the date type out. Nope. gorm.Model also uses timestamps.
	Name        string
	Description string
	AuthorID    int    `gorm:"foreignKey"`
	Value       string // Changed to string, because we'll be storing base64 representations anyway.
	Extension   string
	DateAdded   time.Time
}

type Comment struct {
	ID        int `gorm:"primaryKey"`
	ImageID   int `gorm:"foreignKey"`
	UserID    int `gorm:"foreignKey"`
	Text      string
	RepliesTo int `gorm:"foreignKey"` // replies to -> commentID
	DateAdded time.Time
}

type Favourite struct {
	ID      int `gorm:"primaryKey"`
	ImageID int `gorm:"foreignKey"`
	UserID  int `gorm:"foreignKey"`
}

type Grade struct {
	ID      int `gorm:"primaryKey"`
	ImageID int `gorm:"foreignKey"`
	UserID  int `gorm:"foreignKey"`
	Grade   int
}

var db *gorm.DB
var err error

var (
	// drivers = []Driver{
	// 	{Name: "Jimmy Johnson", License: "ABC123"},
	// 	{Name: "Howard Hills", License: "XYZ789"},
	// 	{Name: "Craig Colbin", License: "DEF333"},
	// }
	// cars = []Car{
	// 	{Year: 2000, Make: "Toyota", ModelName: "Tundra", DriverID: 1},
	// 	{Year: 2001, Make: "Honda", ModelName: "Accord", DriverID: 1},
	// 	{Year: 2002, Make: "Nissan", ModelName: "Sentra", DriverID: 2},
	// 	{Year: 2003, Make: "Ford", ModelName: "F-150", DriverID: 3},
	// }

	users = []User{
		{ID: 1, Username: "student", Password: "kocham piwo"},
		{ID: 2, Username: "student2", Password: "te?? kocham piwo"},
	}
	images = []Image{
		{
			ID:          1,
			Name:        "testo z pomaranczami",
			Description: "graphic design is my passion",
			AuthorID:    1,
			Value:       "iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==",
			Extension:   "png",
		},
	}
	comments = []Comment{
		{ID: 1, ImageID: 1, UserID: 1, Text: "masz pocz??stuj si??", RepliesTo: 0},
		{ID: 2, ImageID: 1, UserID: 2, Text: "nie dziekuje", RepliesTo: 1},
		{ID: 3, ImageID: 1, UserID: 1, Text: "nie dla psa", RepliesTo: 2},
	}
	favourites = []Favourite{
		{ID: 1, ImageID: 1, UserID: 1},
		{ID: 2, ImageID: 1, UserID: 2},
	}
	grades = []Grade{
		{ID: 1, ImageID: 1, UserID: 1, Grade: 5},
	}
)

/** database consts */
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "web_gallery"
)

func main() {

	/** image encoding to string and next to html */
	/*
		myImage := image.NewRGBA(image.Rect(0, 0, 10, 20))
		var buff bytes.Buffer
		png.Encode(&buff, myImage)
		encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())
		htmlImage := "<img src=\"data:image/png;base64," + encodedString + "\" />"
		fmt.Println(htmlImage)
	*/

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	//db.AutoMigrate(&Driver{})
	//db.AutoMigrate(&Car{})

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Image{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Favourite{})

	/*
		for index := range cars {
			db.Create(&cars[index])
		}

		for index := range drivers {
			db.Create(&drivers[index])
		}
	*/
	/** add row to database table */

	// We don't need to insert default shit no mo'.
	// for index := range users {
	// 	db.Create(&users[index])
	// }
	// for index := range images {
	// 	db.Create(&images[index])
	// }
	// for index := range comments {
	// 	db.Create(&comments[index])
	// }
	// for index := range favourites {
	// 	db.Create(&favourites[index])
	// }
	// for index := range grades {
	// 	db.Create(&grades[index])
	// }

	/** rest handlers */
	router := mux.NewRouter()
	// router.HandleFunc("/cars", GetCars).Methods("GET")
	// router.HandleFunc("/cars/{id}", GetCar).Methods("GET")
	// router.HandleFunc("/drivers/{id}", GetDriver).Methods("GET")
	// router.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")

	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	router.HandleFunc("/images", GetImages).Methods("GET")
	router.HandleFunc("/images/{id}", GetImage).Methods("GET")
	router.HandleFunc("/images/{id}", DeleteImage).Methods("DELETE")

	router.HandleFunc("/comments", GetComments).Methods("GET")
	router.HandleFunc("/comments/{id}", GetComment).Methods("GET")
	router.HandleFunc("/comments/{id}", DeleteComment).Methods("DELETE")

	router.HandleFunc("/favourites", GetFavourites).Methods("GET")
	router.HandleFunc("/favourites/{id}", GetFavourite).Methods("GET")
	router.HandleFunc("/favourites/{id}", DeleteFavourite).Methods("DELETE")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))
}

func GetCars(w http.ResponseWriter, r *http.Request) {
	var cars []Car
	db.Find(&cars)
	json.NewEncoder(w).Encode(&cars)
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var car Car
	db.First(&car, params["id"])
	json.NewEncoder(w).Encode(&car)
}

func GetDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var driver Driver
	var cars []Car
	db.First(&driver, params["id"])
	db.Model(&driver).Related(&cars)
	driver.Cars = cars
	json.NewEncoder(w).Encode(&driver)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var car Car
	db.First(&car, params["id"])
	db.Delete(&car)

	var cars []Car
	db.Find(&cars)
	json.NewEncoder(w).Encode(&cars)
}

/** Rest User handler functions */

/** Users handlers */

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)
	if err := json.NewEncoder(w).Encode(&users); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	db.First(&user, params["id"])
	json.NewEncoder(w).Encode(&user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	u := User{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		return
	}

	users = append(users, u)
	db.Create(&users[len(users)-1])

	response, err := json.Marshal(&u)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	db.First(&user, params["id"])
	db.Delete(&user)

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

/** Images handlers */

func GetImages(w http.ResponseWriter, r *http.Request) {
	var images []Image
	db.Find(&images)
	json.NewEncoder(w).Encode(&images)
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var image Image
	db.First(&image, params["id"])
	json.NewEncoder(w).Encode(&image)
}

func DeleteImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var image Image
	db.First(&image, params["id"])
	db.Delete(&image)

	var images []Image
	db.Find(&images)
	json.NewEncoder(w).Encode(&images)
}

/** Comments handlers */
func GetComments(w http.ResponseWriter, r *http.Request) {
	var comments []Image
	db.Find(&comments)
	json.NewEncoder(w).Encode(&comments)
}

func GetComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var comment Comment
	db.First(&comment, params["id"])
	json.NewEncoder(w).Encode(&comment)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var comment Comment
	db.First(&comment, params["id"])
	db.Delete(&comment)

	var comments []Comment
	db.Find(&comments)
	json.NewEncoder(w).Encode(&comments)
}

/** Favourites handlers */

func GetFavourites(w http.ResponseWriter, r *http.Request) {
	var favourites []Favourite
	db.Find(&favourites)
	json.NewEncoder(w).Encode(&favourites)
}

func GetFavourite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var favourite Favourite
	db.First(&favourite, params["id"])
	json.NewEncoder(w).Encode(&favourite)
}

func DeleteFavourite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var favourite Favourite
	db.First(&favourite, params["id"])
	db.Delete(&favourite)

	var favourites []Favourite
	db.Find(&favourites)
	json.NewEncoder(w).Encode(&favourites)
}
