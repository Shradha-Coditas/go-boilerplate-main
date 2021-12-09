package albums

import (
	MysqlWrapper "boilerplate/library/mysql"
	"boilerplate/util/logwrapper"
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//albumRequest - album post request structure.
type AlbumRequest struct {
	Title  string  `form:"title" json:"title" binding:"required"`   //title of Album
	Artist string  `form:"artist" json:"artist" binding:"required"` //Artist name of Album
	Price  float64 `form:"price" json:"price" binding:"required"`   //Price of Album
}

//AlbumUpdateRequest - album update request structure.
type AlbumUpdateRequest struct {
	ID     int64   `form:"id" json:"id" binding:"required"` //Id of Album
	Title  string  `form:"title" json:"title"`              //title of Album
	Artist string  `form:"artist" json:"artist"`            //Artist name of Album
	Price  float64 `form:"price" json:"price"`              //Price of Album
}

// AlbumResponse - Album Response structure.
type AlbumResponse struct {
	Message string `form:"message" json:"message"`
}

// AlbumResposeData - album AlbumResponseData response structure.
type AlbumsResponseData struct {
	ID        int64   `json:"id"`        //Id of Album
	Title     string  `json:"title"`     //title of Album
	Artist    string  `json:"artist"`    //Artist name of Album
	Price     float64 `json:"price"`     //Price of Album
	CreatedAt string  `json:"createdAt"` // time at row created
}

// ArtistDetails - album ArtistDetails response structure.
type ArtistDetails struct {
	ID        int64   `json:"id"`        //Id of Album
	Title     string  `json:"title"`     //title of Album
	Artist    string  `json:"artist"`    //Artist name of Album
	Price     float64 `json:"price"`     //Price of Album
	CreatedAt string  `json:"createdAt"` // time at row created
}

func createAlbumTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS albums(
		id int primary key auto_increment, 
		title text,
		artist text, 
        price int, 
		createdat datetime default CURRENT_TIMESTAMP
	)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	rows, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating album table", err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return nil
}

// Query Executer For Album Insert
func InsertAlbum(s AlbumRequest) (bool, error) {
	err := createAlbumTable(MysqlWrapper.Client)
	if err != nil {
		log.Printf("Create Albums table failed with error %s", err)
		return false, err
	}
	args := []interface{}{s.Title, s.Artist, s.Price}
	insert, err := MysqlWrapper.Client.Query("INSERT INTO albums(title,artist,price) VALUES(?,?,?);", args...)
	// if there is an error inserting, handle it

	if err != nil {
		logwrapper.Logger.Debugln(err)
		return false, err
	}
	defer insert.Close()
	return true, nil
}

// Query Executer For Album Update
func UpdateAlbum(s AlbumUpdateRequest) (int64, error) {
	Update, err := MysqlWrapper.Client.Exec("update albums set title = ?, artist = ?, price = ? where id = ?", s.Title, s.Artist, s.Price, s.ID)
	// if there is an error Updating, handle it
	if err != nil {
		return 0, err
	}
	rows, err := Update.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

// Query Executer For Album GET
func GetAlbum(offset int64, limit int64) ([]AlbumsResponseData, error) {
	GetDetails, err := MysqlWrapper.Client.Query("select * from albums limit ?, ?", offset, limit)
	// if there is an error Get data, handle it
	if err != nil {
		panic(err.Error())
	}
	albums := []AlbumsResponseData{}
	for GetDetails.Next() {
		var album AlbumsResponseData
		err = GetDetails.Scan(&album.ID, &album.Title, &album.Artist, &album.Price, &album.CreatedAt)
		if err != nil {
			logwrapper.Logger.Debugln(err)
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		albums = append(albums, album)
	}
	return albums, nil
}

// Query Executer For Album By Id
func GetAlbumById(id int64) (AlbumsResponseData, error) {
	GetDetails, err := MysqlWrapper.Client.Query("select * from albums WHERE id = ?", id)
	// if there is an error Get data, handle it
	if err != nil {
		panic(err.Error())
	}
	var album AlbumsResponseData
	for GetDetails.Next() {
		err = GetDetails.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
	return album, nil
}

// Query Executer For Album Delete
func DeletAlbum(id int64) (int64, error) {
	Delete, err := MysqlWrapper.Client.Exec("delete from albums where id = ?", id)
	// if there is an error Deleting, handle it
	if err != nil {
		// panic(err.Error())
		logwrapper.Logger.Debugln(err)
		return 0, err
	}
	rows, err := Delete.RowsAffected()
	if err != nil {
		return 0, err
	}
	logwrapper.Logger.Debugln(rows)
	return rows, nil
}

// Query Executer For Album Check Row exists
func RowExists(id int64) (int64, error) {
	count, err := MysqlWrapper.Client.Query("Select count() FROM albums where id = ?", id)
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	var count_album int64
	for count.Next() {
		count.Scan(&count_album)
	}
	return count_album, nil
}
