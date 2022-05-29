package UserMod

import (
	"time"
)

type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func ConsultaUser() {

	// db := Connection.Conectar()

	// consulta, err := db.Query("SELECT * FROM mydb.TBL_Usuarios;")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// return consulta
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

func (this *Usuario) altaUsuario(id int, nombre string, fechaalta time.Time, status bool) {
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.Status = status
}
