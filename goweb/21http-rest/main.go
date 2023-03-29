package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

// Route Name表示 HTTP 方法的名称，Method表示 HTTP 方法类型，可以是GET、POST、PUT、DELETE等等，Pattern表示 URL 路径，HandlerFunc表示 HTTP 处理程序
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route{
		"getEmployees",
		"GET",
		"/employees",
		getEmployees,
	},
	Route{
		"getEmployee",
		"GET",
		"/employee/{id}",
		getEmployee,
	},
	Route{
		"addEmployee",
		"POST",
		"/employee/add",
		addEmployee,
	},
	Route{
		"updateEmployee",
		"PUT",
		"/employee/update",
		updateEmployee,
	},
	Route{
		"deleteEmployee",
		"DELETE",
		"/employee/delete",
		deleteEmployee,
	},
}

type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type Employees []Employee

var (
	employees   []Employee
	employeesV1 []Employee
	employeesV2 []Employee
)

func init() {
	employees = Employees{
		Employee{Id: "1", FirstName: "Foo", LastName: "Bar"},
		Employee{Id: "2", FirstName: "Baz", LastName: "Qux"},
	}
	employeesV1 = Employees{
		Employee{Id: "1", FirstName: "Foo", LastName: "Bar"},
		Employee{Id: "2", FirstName: "Baz", LastName: "Qux"},
	}
	employeesV2 = Employees{
		Employee{Id: "1", FirstName: "Baz", LastName: "Qux"},
		Employee{Id: "2", FirstName: "Quux", LastName: "Quuz"},
	}
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	// 修改了getEmployees处理程序，以检查 URL 中的前缀，并采取相应的行动
	if strings.HasPrefix(r.URL.Path, "/v1") {
		_ = json.NewEncoder(w).Encode(employeesV1)
	} else if strings.HasPrefix(r.URL.Path, "/v2") {
		_ = json.NewEncoder(w).Encode(employeesV2)
	} else {
		// 默认提供一个包含单个记录的列表，我们可以将其称为 REST 端点的默认或初始响应
		_ = json.NewEncoder(w).Encode(employees)
	}
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, employee := range employees {
		if employee.Id == id {
			if err := json.NewEncoder(w).Encode(employee); err != nil {
				log.Print("error getting requested employee :: ", err)
			}
		}
	}
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Print("error occurred while decoding employee data :: ", err)
		return
	}
	log.Printf("adding employee id :: %s with firstName as :: %s and lastName as :: %s ", employee.Id, employee.FirstName, employee.LastName)
	employees = append(employees, Employee{
		Id:        employee.Id,
		FirstName: employee.FirstName, LastName: employee.LastName,
	})
	_ = json.NewEncoder(w).Encode(employees)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Print("error occurred while decoding employee data :: ", err)
		return
	}
	isUpsert := true
	for idx, emp := range employees {
		if emp.Id == employee.Id {
			isUpsert = false
			log.Printf("updating employee id :: %s with firstName as :: %s and lastName as:: %s ", employee.Id, employee.FirstName, employee.LastName)
			employees[idx].FirstName = employee.FirstName
			employees[idx].LastName = employee.LastName
			break
		}
	}
	if isUpsert {
		log.Printf("upserting employee id :: %s with firstName as :: %s and lastName as:: %s ",
			employee.Id, employee.FirstName, employee.LastName)
		employees = append(employees, Employee{
			Id:        employee.Id,
			FirstName: employee.FirstName, LastName: employee.LastName,
		})
	}
	_ = json.NewEncoder(w).Encode(employees)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Print("error occurred while decoding employee data :: ", err)
		return
	}
	log.Printf("deleting employee id :: %s with firstName as :: %s and lastName as :: %s ", employee.Id,
		employee.FirstName, employee.LastName)
	index := GetIndex(employee.Id)
	employees = append(employees[:index], employees[index+1:]...)
	_ = json.NewEncoder(w).Encode(employees)
}

func GetIndex(id string) int {
	for i := 0; i < len(employees); i++ {
		if employees[i].Id == id {
			return i
		}
	}
	return -1
}

// AddRoutes 迭代我们定义的 routes 数组，将其添加到gorilla/mux路由器，并返回Router对象
func AddRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

func main() {
	// 密码StrictSlash定义了新路由的尾部斜线行为。默认值为false。
	// 当为true时，如果路由路径为“/path/”，则访问“/path”将执行到前者的重定向，反之亦然。换句话说，您的应用程序将始终看到路由中指定的路径。
	// 当为false时，如果路由路径为“/path”，则访问“/path/”将不匹配这条路线，反之亦然。
	// 重定向是HTTP 301（永久移动）。请注意，当使用非幂等方法（例如POST、PUT）为路由设置此选项时，大多数客户端将以GET的形式进行后续重定向请求。根据需要使用中间件或客户端设置来修改此行为。
	// 特殊情况：当路由使用PathPrefix（）方法设置路径前缀时，该路由的严格斜线被忽略，因为重定向行为不能可以仅根据前缀来确定。但是，从该路由创建的任何子例程都会继承原始的StrictSlash设置。
	muxRouter := mux.NewRouter().StrictSlash(true)
	router := AddRoutes(muxRouter)
	// 定义两个版本的相同 URL 路径，支持 HTTPGET方法，其中一个以v1为前缀，另一个以v2为前缀
	// v1
	AddRoutes(muxRouter.PathPrefix("/v1").Subrouter())
	// v2
	AddRoutes(muxRouter.PathPrefix("/v2").Subrouter())
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
