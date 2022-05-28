package main

import (

	// "./Connection"
	"./test"
)

// type User struct {
// 	Id            string
// 	Nombre        string
// 	Correo        string
// 	NumeroCel     string
// 	FechaCreacion string
// 	Pass          string
// 	Status        string
// }

// func home(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "./index.hCml")
// }

func main() {

	// fmt.Println(Connection.Connec())
	test.Prueba()

	// http.HandleFunc("/", home)
	// http.ListenAndServe(":3000", nil)

	// cfg := mysql.Config{
	// 	User:                 "root",
	// 	Passwd:               "example",
	// 	Net:                  "tcp",
	// 	Addr:                 "127.0.0.1:3306",
	// 	DBName:               "",
	// 	AllowNativePasswords: true,
	// }

	// var err error
	// db, err = sql.Open("mysql", cfg.FormatDSN())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// pingErr := db.Ping()
	// if pingErr != nil {
	// 	log.Fatal(pingErr)
	// }
	// fmt.Println("Conected!!!")

	// consulta, err := db.Query("SELECT * FROM mydb.TBL_Usuarios;")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// for consulta.Next() {
	// 	var user User
	// 	err = consulta.Scan(&user.Id, &user.Nombre, &user.Correo, &user.NumeroCel, &user.FechaCreacion, &user.Pass, &user.Status)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	log.Printf(user.Nombre, user.Correo)
	// }

	// insert, err := db.Exec("INSERT INTO mydb.TBL_Usuarios(usr_names, usr_mail, usr_num_verification, date_regist, usr_pass, usr_state ) VALUES ('2','dpfddjj','f','f','f','f')")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// fmt.Println(insert)

}
