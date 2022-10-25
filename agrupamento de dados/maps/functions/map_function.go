package functions

import "fmt"

func SaveUsers(user string, number int){
    users := map[string]int{"Gopher":7483545}
    if userExist, ok := users["Testeshhsfdj"]; !ok{
        fmt.Println("Number not found");
    }else{
        fmt.Println(userExist);
    }
    users[user] = number;
    fmt.Println(users);
    fmt.Println(users[user]);
}
func RangeMap(id int, name string){
      estoque := map[int]string{
        200: "uvas",
       }
    estoque[id] = name;
    for ids, names := range estoque{
         fmt.Println("Product:\t",ids,names);
     }
    delete(estoque,200);
    fmt.Println("Deleting uvas Product:\n .1%\n...3%\n........11%\n..............99%\n[Finish]100% Deleted");
}


