package main

import 
(
    "fmt"
    "os"
    "log"
    "math/rand"
    "strconv"
)

var states = []string {
        "Washingston",
        "Oregon",
        "California",
        "Nevada",
        "Idaho",
        "Montana",
        "Wyoming",
        "Utah",
        "Arizona",
        "Colorado",
        "New Mexico",
        "Texas",
        "North Dakota",
        "South Dakota",
        "Nebraska",
        "Kansas",
        "Oklahoma",
        "Minnesota",
        "Iowa",
        "Missouri",
        "Arkansas",
        "Louisiana",
        "Wisconsin",
        "Illinois",
        "Michigan",
        "Indiana",
        "Kentucky",
        "Tennessee",
        "Alabama",
        "Mississippi",
        "Georgia",
        "Florida",
        "South Carolina",
        "North Carolina",
        "Virginia",
        "Ohio",
        "West Virginia",
        "Washington D.C.",
        "Maryland",
        "Delaware",
        "New Jersey",
        "New York",
        "Connecticut",
        "Rhode Island",
        "Maine",
        "New Hampshire",
        "Vermont",
        "Massachusetts",
        "Alaska",
        "Hawaii",
        "Pennyslvania",
    };

var street = []string{
        "1st",
        "2nd",
        "3rd",
        "4th",
        "5th",
        "6th",
        "7th",
        "8th",
        "9th",
        "10th",
        "Oak",
        "Jackson",
        "Park",
        "Pine",
        "Main",
        "Willow",
        "Birch",
        "Sunset",
        "Apache",
        "Elm",
        "Maple",
        "Redwood",
        "Sunset",
        "Lake",
        "Meadow",
        "West",
        "East",
        "North",
        "South",
        "Hillside",
        "Evergreen",
        "Ridge",
    };

var users = []string{
    "Damian",
    "Nicole",
    "Jared",
    "Becca",
    "James",
    "Ronald",
    "John",
    "Robert",
    "Michael",
    "Mark",
    "Donald",
    "Anthony",
    "Jason",
    "Steven",
    "Mary",
    "Lisa",
    "Laura",
    "Sarah",
    "Karen",
    "Betty",
    "Linda",
    "Jennifer",
    "Helen",
    "Sandra",
    "Donna",
    "Susan",
    };

var topics = []string{
    "Basketball",
    "Tennis",
    "Basketball",
    "Volleyball",
    "Pool",
    "Swimming",
    "Skiing",
    "Beach",
    "Park",
    "Haunted",
    "Food",
    "Carnivals",
    "Cafes",
    "Amusement Park",
    "Holiday",
    "Romantic",
    "Zoos",
    "Hiking",
    "Monuments",
    "Muesems",
    "Art Gallery",
    "Ponds",
    "Piers",
    "Fishing",
    "Shopping Mall",
    "Bowling",
    "Concert",
    "Feeding Animals",
    "Apple Picking",
    "Boxing",
    }



func createFile_Age(){
    file, err := os.Create("age.csv")
    if(err != nil){
        log.Fatal("Cannot Create Age", err)
    }

    defer file.Close()

    for i:= 0; i < len(topics); i++ {
        line := topics[i] + ", "
        ageMin := (18 + rand.Intn(50))
        ageMax := (18 + rand.Intn(50))
        
        if(ageMin > ageMax){
            temp := ageMin
            ageMin = ageMax
            ageMax = temp
        }

        line += (strconv.Itoa(ageMin) + "-" + strconv.Itoa(ageMax) + "\n")
        write, err := file.WriteString(line)
        fmt.Printf("Age Index: %d wrote %d bytes\n", i, write)
        if(err != nil){
            log.Fatal("Error Writing To File ", err)
        }
        file.Sync()
        line = ""
    }
}

func createFile_Group(){
    file, err := os.Create("group.csv")
    if(err != nil){
        log.Fatal("Cannot Create Group", err)
    }

    defer file.Close()

    for i:= 0; i < len(topics); i++ {
        line := topics[i] + ", "
        group := (1 + rand.Intn(10))
        
        line += (strconv.Itoa(group) + "\n")
        write, err := file.WriteString(line)
        fmt.Printf("Group Index: %d wrote %d bytes\n", i, write)
        if(err != nil){
            log.Fatal("Error Writing To File ", err)
        }
        file.Sync()
        line = ""
    }
}

func createFile_Seasonal(){
    file, err := os.Create("seasonal.csv")
    if(err != nil){
        log.Fatal("Cannot Create seasonal", err)
    }

    defer file.Close()

    for i:= 0; i < len(topics); i++ {
        line := topics[i] + ", "
        season := (rand.Intn(4))
        if(season == 0){
            line += "winter, "
        }else if(season == 1){
            line += "spring, "
        }else if(season == 2){
            line += "fall, "
        }else if(season == 3){
            line += "summer, "
        }
        
        line += states[rand.Intn(len(states))] + "\n"

        write, err := file.WriteString(line)
        fmt.Printf("Seasonal Index: %d wrote %d bytes\n", i, write)
        if(err != nil){
            log.Fatal("Error Writing To File ", err)
        }
        file.Sync()
        line = ""
    }
}

func createFile_Cost(){
    file, err := os.Create("cost.csv")
    if(err != nil){
        log.Fatal("Cannot Create seasonal", err)
    }

    defer file.Close()

    line := ""
    for i:= 0; i <= 1000; i++ {
        address := ""
        address = strconv.Itoa(1 + rand.Intn(100))
        address += (" " + street[rand.Intn(len(street))])
        streetType := rand.Intn(1)
        if(streetType == 1){
            address += " Street, "
        }else{
            address += " Avenue, "
        }
        line += address
        line += strconv.Itoa(rand.Intn(99999)) + ", "
        line += states[rand.Intn(len(states))] + ", "
        line += topics[rand.Intn(len(topics))] + ", "
        line += strconv.Itoa(rand.Intn(100)) + "\n"

        write, err := file.WriteString(line)
        fmt.Printf("Cost Index: %d wrote %d bytes\n", i, write)
        if(err != nil){
            log.Fatal("Error Writing To File ", err)
        }
        file.Sync()
        line = ""
    }
}

func createFile_TimeofDay(){
    file, err := os.Create("time_of_day.csv")
    if(err != nil){
        log.Fatal("Cannot Create Time of Day", err)
    }

    defer file.Close()

    for i:= 0; i < len(topics); i++ {
        line := topics[i] + ", "
       
        ihour := strconv.Itoa((rand.Intn(24)))
        imin := strconv.Itoa((1 + rand.Intn(60)))
        fhour := strconv.Itoa((rand.Intn(24)))
        fmin := strconv.Itoa(1 + rand.Intn(60))

        line += ((ihour + ":" + imin) + " - " + (fhour + ":" + fmin) + "\n")
        write, err := file.WriteString(line)
        fmt.Printf("Time of Day Index: %d wrote %d bytes\n", i, write)
        if(err != nil){
            log.Fatal("Error Writing To File ", err)
        }
        file.Sync()
        line = ""
    }

}

func createFile_Point_of_Interest(){
    file, err := os.Create("point_of_interest.csv")

    if(err != nil){
        log.Fatal("Cannot Create Point of Interest", err)
    }

    defer file.Close()

    line := ""
    for i:=0; i <= 1000; i++ {
        address := ""
        address = strconv.Itoa(1 + rand.Intn(100))
        address += (" " + street[rand.Intn(len(street))])
        streetType := rand.Intn(1)
        if(streetType == 1){
            address += " Street, "
        }else{
            address += " Avenue, "
        }
        line += address
        line += strconv.Itoa(rand.Intn(99999)) + ", "
        line += states[rand.Intn(len(states))] + ", "
        line += topics[rand.Intn(len(topics))] + "\n"
    
        write, err := file.WriteString(line)
        fmt.Printf("Point of Interest Index: %d wrote %d bytes\n", i, write)
        if(err != nil){
            log.Fatal("Error Writing To File ", err)
        }
        file.Sync()
        line = ""
    }
}

func createFile_User(){
    file, err := os.Create("user.csv")
    
    if err != nil{
        log.Fatal("Cannot Create User Table", err) 
    }
    
    defer file.Close()
    
    line := ""
    for i:=0; i <= 1000; i++{
        number := strconv.Itoa(rand.Intn(9999))
        
        line += users[rand.Intn(len(users))] + number + ", "
        
        age := strconv.Itoa(18 + rand.Intn(50))
        line += age + ", "

        gender := rand.Intn(1)
        if(gender == 1){
            line += "female, " 
        }else{
            line += "male, "
        }

        address := ""
        address = strconv.Itoa(1 + rand.Intn(100))
        address += (" " + street[rand.Intn(len(street))])
        streetType := rand.Intn(1)
        if(streetType == 1){
            address += " Street, "
        }else{
            address += " Avenue, "
        }

        line += address

        line += strconv.Itoa(rand.Intn(99999)) + ", "

        line += states[rand.Intn(len(states))] + "\n"

        write, err := file.WriteString(line)
        fmt.Printf("User Index: %d wrote %d bytes\n", i, write)
        if(err != nil){
            log.Fatal("Error Writing To File ", err)
        }
        file.Sync()
        line = ""
    }
}

func main(){
    //createFile_User()
    //createFile_Point_of_Interest()
    //createFile_Age()
    //createFile_Group()
    //createFile_Seasonal()
    //createFile_TimeofDay()
    createFile_Cost()
}
