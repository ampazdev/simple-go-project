package constant

const (
	ConfigProjectFilepath = "../../svc/userservice/config/userservice.development.yaml" // %s = environment
	//ConfigProjectFilepath     = "svc/userservice/config/userservice.development.yaml"   // %s = environment
	ConfigLinuxConfigFilepath = "/etc/goproject/userservice.%s.yaml" // %s = environment
	SecretJWT                 = "testSecret"
)
