package main
import  ("github.com/julienschmidt/httprouter"
		"fmt"
		"net/http"
		"strconv"
		"encoding/json"
		"strings"
		"sort")


type key_val struct{
	Key int	`json:"key,omitempty"`
	Value string	`json:"value,omitempty"`
} 


var s1,s2,s3 [] key_val
var index1,index2,index3 int
type ByKey []key_val
func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }


func Get_Keys(rw http.ResponseWriter, request *http.Request,p httprouter.Params){
	port := strings.Split(request.Host,":")
	if(port[1]=="3000"){
		sort.Sort(ByKey(s1))
		result,_:= json.Marshal(s1)
		fmt.Fprintln(rw,string(result))
	}else if(port[1]=="3001"){
		sort.Sort(ByKey(s2))
		result,_:= json.Marshal(s2)
		fmt.Fprintln(rw,string(result))
	}else{
		sort.Sort(ByKey(s3))
		result,_:= json.Marshal(s3)
		fmt.Fprintln(rw,string(result))
	}
}

func Put_Keys(rw http.ResponseWriter, request *http.Request,p httprouter.Params){
	port := strings.Split(request.Host,":")
	key,_ := strconv.Atoi(p.ByName("key_id"))
	if(port[1]=="3000"){
		s1 = append(s1,key_val{key,p.ByName("value")})
		index1++
	}else if(port[1]=="3001"){
		s2 = append(s2,key_val{key,p.ByName("value")})
		index2++
	}else{
		s3 = append(s3,key_val{key,p.ByName("value")})
		index3++
	}	
}

func Get_Key(rw http.ResponseWriter, request *http.Request,p httprouter.Params){	
	out := s1
	ind := index1
	port := strings.Split(request.Host,":")
	if(port[1]=="3001"){
		out = s2 
		ind = index2
	}else if(port[1]=="3002"){
		out = s3
		ind = index3
	}	
	key,_ := strconv.Atoi(p.ByName("key_id"))
	for i:=0 ; i< ind ;i++{
		if(out[i].Key==key){
			result,_:= json.Marshal(out[i])
			fmt.Fprintln(rw,string(result))
		}
	}
}



func main(){
	index1 = 0
	index2 = 0
	index3 = 0
	mux := httprouter.New()
    mux.GET("/keys",Get_Keys)
    mux.GET("/keys/:key_id",Get_Key)
    mux.PUT("/keys/:key_id/:value",Put_Keys)
    go http.ListenAndServe(":3000",mux)
    go http.ListenAndServe(":3001",mux)
    go http.ListenAndServe(":3002",mux)
    select {}
}