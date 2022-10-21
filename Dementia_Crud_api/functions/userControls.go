package functions

import (
	"dementia/go_crud/initilizers"
	"dementia/go_crud/models"
    "github.com/gin-gonic/gin"
)


func UserCreate(c *gin.Context){
	
	//get data of  req body
	var user struct{
		Email    string
	    Password string
	    Name     string
	}
	c.Bind(&user)
	
	//create a user
	newUser := models.User{Email:user.Email, Name:user.Name,Password:user.Password} //json
	//check unique email or not
    initilizers.DB.Where("email = ?",newUser.Email).Find(&newUser)
	
	//response
	if  newUser.Id == 0 {
       
		// pass pointer of data to Create
	result := initilizers.DB.Create(&newUser) 
	if result.Error !=nil{
		c.Status(400)
		return
	}
	
	//REturn it
	
			c.JSON(200, gin.H{
				"newUser": newUser,
			})
	
     }else{
		c.JSON(406,gin.H{
			"message":"email already exist",
		})
	 }


	
		
	}

func ShowUser(c *gin.Context){
	//POST method
	//get details
	var user struct{
		Email    string
	    Password string
	}
	c.Bind(&user)

    var usr models.User
	initilizers.DB.Where("email = ? AND password = ?",user.Email,user.Password).Find(&usr)
	
	//response
	if usr.Id==0{
	c.Status(400)
}else{
	c.IndentedJSON(200, gin.H{
		"user": usr,
	})
}
}

func GetUser(c *gin.Context){
	//get details
	email:=c.Param("email")
	password:=c.Param("password")
    var user models.User
	initilizers.DB.Where("email = ? AND password = ?",email,password).Find(&user)
	
	//response
	if user.Id==0{
		c.JSON(404,gin.H{
			"meaa":"error",
		})
	}else{
	c.JSON(200, gin.H{
		"user": user,
	})
   }
}