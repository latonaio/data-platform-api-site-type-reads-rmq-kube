package requests

type Text struct {
	SiteType     		string  `json:"SiteType"`
	Language          	string  `json:"Language"`
	SiteTypeName		string  `json:"SiteTypeName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
