package Form

type Container struct {
	Name     string `json:"name"`     
	Password string `json:"password"` 
	Port     string `json:"port"`     
	Token    string `json:"token"`   
	User     string `json:"user"`    
	Type     string `json:"type"`
	ID string `json:"id"`     
}

type ContainerAddRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Type     string `json:"type"`
}