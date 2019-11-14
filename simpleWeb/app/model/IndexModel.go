
package model

import(
	"fmt"

	"simpleWeb/lib/base"
)

type CityModel struct{
	base.Model
}

type City struct {
	ID int
	Name string 
	CountryCode string 
	District string
	Population int
}

func (model *CityModel) GetCity()(string,error){
	city := new(City)
	row := model.Conn.QueryRow("SELECT ID,Name,CountryCode,District,Population FROM world.city LIMIT 1")
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&city.ID,&city.Name,&city.CountryCode,&city.District,&city.Population); err != nil{
		fmt.Printf("scan failed, err:%v",err)
		return "",err
	}
	return city.Name,nil
}
