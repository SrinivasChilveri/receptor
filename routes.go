package receptor

import "github.com/tedsuo/rata"

const (
	// Tasks
	CreateTaskRoute          = "CreateTask"
	GetAllTasksRoute         = "GetAllTasks"
	GetAllTasksByDomainRoute = "GetAllTasksByDomain"
	GetTaskRoute             = "GetTask"
	DeleteTaskRoute          = "DeleteTask"

	// DesiredLRPs
	CreateDesiredLRPRoute          = "CreateDesiredLRP"
	GetDesiredLRPRoute             = "GetDesiredLRP"
	UpdateDesiredLRPRoute          = "UpdateDesiredLRP"
	DeleteDesiredLRPRoute          = "DeleteDesiredLRP"
	GetAllDesiredLRPsRoute         = "GetAllDesiredLRPs"
	GetAllDesiredLRPsByDomainRoute = "GetAllDesiredLRPsByDomain"

	// ActualLRPs
	GetAllActualLRPsRoute = "GetAllActualLRPs"
)

var Routes = rata.Routes{
	// Tasks
	{Path: "/tasks", Method: "POST", Name: CreateTaskRoute},
	{Path: "/tasks", Method: "GET", Name: GetAllTasksRoute},
	{Path: "/domains/:domain/tasks", Method: "GET", Name: GetAllTasksByDomainRoute},
	{Path: "/tasks/:task_guid", Method: "GET", Name: GetTaskRoute},
	{Path: "/tasks/:task_guid", Method: "DELETE", Name: DeleteTaskRoute},

	// DesiredLRPS
	{Path: "/desired_lrps", Method: "POST", Name: CreateDesiredLRPRoute},
	{Path: "/desired_lrps/:process_guid", Method: "GET", Name: GetDesiredLRPRoute},
	{Path: "/desired_lrps/:process_guid", Method: "PUT", Name: UpdateDesiredLRPRoute},
	{Path: "/desired_lrps/:process_guid", Method: "DELETE", Name: DeleteDesiredLRPRoute},
	{Path: "/desired_lrps", Method: "GET", Name: GetAllDesiredLRPsRoute},
	{Path: "/domains/:domain/desired_lrps", Method: "GET", Name: GetAllDesiredLRPsByDomainRoute},

	// ActualLRPS
	{Path: "/actual_lrps", Method: "GET", Name: GetAllActualLRPsRoute},
}
